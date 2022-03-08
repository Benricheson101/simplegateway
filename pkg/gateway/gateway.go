package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benricheson101/simplegateway/pkg/gateway/intents"
	"github.com/benricheson101/simplegateway/pkg/rest"
	"github.com/gorilla/websocket"
)

// TODO: better error messages
// TODO: should there be any deadline set on WS reads/writes?

const (
	// The number of heartbeat acks allowed to be skipped before
	// automatically reconnecting. This is a strong indicator
	// that the connection between the client and Discord gateway
	// has been severed.
	HEARTBEAT_ACK_FAILURE_BEFORE_RECONNECT = 5

	// The maximum number of seconds to wait before reconnecting
	// to the gateway. Backoff is exponential until reaching
	// this duration.
	MAX_BACKOFF_SECS = 300 * time.Second
)

type GatewayStatus int

const (
	DOWN GatewayStatus = iota
	WAITING_FOR_HELLO
	HELLO_RECEIVED
	IDENTIFYING
	WAITING_FOR_READY
	RESUMING
	WAITING_FOR_RESUMED
	RECONNECTING
	UP
)

type RawHandlerFunc func(*Gateway, *RawGatewayPayload)

type Gateway struct {
	// Used to prevent large operations (Up, Down, Resume) from
	// happening concurrently
	sync.Mutex

	// Used to prevent concurrent R/W to websocket
	wsMu sync.Mutex

	// Used to prevent race conditions when code depends on Gateway::Status
	statusMu sync.Mutex

	// Prevents concurrent writes to the handler maps
	handlerMu sync.RWMutex

	conn *websocket.Conn
	rest *rest.RestClient

	token string

	// How often to send a heartbeat
	heartbeatInterval time.Duration
	// The time of the last heartbeat ack message from Discord
	lastHeartbeatAck time.Time
	// The time the last heartbeat was sent to Discord
	lastHeartbeatSent time.Time

	// Whether or not the gateway connection should be reestablushed
	// if it gets disconnected
	reconnectOnDisconnect bool

	SessionID string
	Sequence  int64

	Status GatewayStatus

	identifyPayload IdentifyPayload
	gateway         rest.GatewayBot

	// Used to stop all goroutines
	cancel context.CancelFunc
	// Stop the heartbeat goroutine
	cancelHB context.CancelFunc
	// Stop the message processor loop
	cancelMP context.CancelFunc

	handlers    map[string][]interface{}
	rawHandlers []RawHandlerFunc
}

// -- public facing API --

// Creates a new Gateway
func New(token string, ops ...GatewayOption) *Gateway {
	var (
		defaultStatus                = DOWN
		defaultGateway               = rest.GatewayBot{}
		defaultRest                  = rest.New(token)
		defaultHandlers              = make(map[string][]interface{})
		defaultRawHandlers           = []RawHandlerFunc{}
		defaultReconnectOnDisconnect = true

		defaultIdentifyPayload = IdentifyPayload{
			Op: IDENTIFY,
			Data: IdentifyPayloadData{
				Token:    token,
				Intents:  int(intents.AllWithoutPrivileged()),
				Presence: &IdentifyPayloadDataPresence{},
				Properties: IdentifyPayloadDataProperties{
					OS:      runtime.GOOS,
					Browser: "simplegateway",
					Device:  "simplegateway",
				},
			},
		}
	)

	gw := &Gateway{
		Status:                defaultStatus,
		gateway:               defaultGateway,
		handlers:              defaultHandlers,
		identifyPayload:       defaultIdentifyPayload,
		rawHandlers:           defaultRawHandlers,
		reconnectOnDisconnect: defaultReconnectOnDisconnect,
		rest:                  defaultRest,
		token:                 token,
	}

	for _, opt := range ops {
		opt(gw)
	}

	return gw
}

// Connects to Discord's gateway. If SessionID and Sequence are both provided, it will first attempt to resume. If this fails, it will automatically re-identify.
func (gw *Gateway) Up(ctx context.Context) error {
	gw.Lock()
	defer func() {
		gw.Unlock()
	}()

	gw.statusMu.Lock()
	if gw.Status == DOWN && gw.conn != nil {
		gw.cleanup(false)
	} else if gw.Status == RECONNECTING {
		// sets the last sent and last ack to now so reconnect doesn't get triggered again
		// when checking for 5 failed heartbeat acks
		gw.lastHeartbeatAck = time.Now()
		gw.lastHeartbeatSent = time.Now()
	}
	gw.statusMu.Unlock()

	ctx, cancel := context.WithCancel(ctx)
	gw.cancel = cancel

	err := gw.openConn()
	if err != nil {
		return err
	}

	err = nil

	if gw.Sequence != 0 && gw.SessionID != "" {
		err = gw.resume(ctx)
	} else {
		err = gw.identify(ctx)
	}

	if err != nil {
		if e, ok := err.(*websocket.CloseError); ok {
			if e.Code == INVALID_SESSION {
				gw.cleanup(true)

				// TODO: crying
				gw.Unlock()
				er := gw.Up(ctx)
				gw.Lock()
				return er
			}
		}

		return err
	}

	// This should never happen, but on the off chance that it does, it'll throw a generic error because cause is unknown
	if gw.Status != UP {
		return ErrCouldNotConnect
	}

	// TODO: is there a way to use the `ctx` that was passed into the function, but have it drain ctx.Done() so it doesn't
	hbCtx, hbCancel := context.WithCancel(context.Background())
	gw.cancelHB = hbCancel
	go gw.startHeartbeatLoop(hbCtx)

	mpCtx, mpCancel := context.WithCancel(context.Background())
	gw.cancelMP = mpCancel
	go gw.startMessageProcessorLoop(mpCtx)

	return nil
}

// Attempts to resume the connection to the gateway. Unlike Up(), if the connection cannot be resumed, it will NOT identify, and will return an error.
func (gw *Gateway) TryResume(ctx context.Context) error {
	gw.Lock()
	defer gw.Unlock()

	err := gw.openConn()
	if err != nil {
		return err
	}

	err = gw.resume(ctx)
	if err != nil {
		return err
	}

	// TODO: move the next block to a func and use it in Up()

	// This should never happen, but on the off chance that it does,
	// it'll throw a generic error because cause is unknown
	if gw.Status != UP {
		return ErrCouldNotConnect
	}

	hbCtx, hbCancel := context.WithCancel(ctx)
	gw.cancelHB = hbCancel
	go gw.startHeartbeatLoop(hbCtx)

	mpCtx, mpCancel := context.WithCancel(ctx)
	gw.cancelMP = mpCancel
	go gw.startMessageProcessorLoop(mpCtx)

	return nil
}

// TODO: down should probably be behind the main Gateay mutex

// Closes the gateway connection with close code `1000`
func (gw *Gateway) Down(ctx context.Context) error {
	return gw.DownWithCode(ctx, websocket.CloseNormalClosure)
}

// Closes the gateway connection with the supplied code
func (gw *Gateway) DownWithCode(ctx context.Context, code int) error {
	gw.cancel()
	gw.Sequence = 0
	gw.SessionID = ""

	gw.statusMu.Lock()
	gw.Status = DOWN
	gw.statusMu.Unlock()

	var err error
	if gw.conn != nil {
		err = gw.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, ""))
		gw.conn = nil
	}

	return err
}

// Adds a function that runs when a specific event is received.
func (gw *Gateway) AddHandlerFunc(fn interface{}) {
	gw.handlerMu.Lock()
	defer gw.handlerMu.Unlock()

	eventName := eventHandlerToEventName(fn)
	if !eventNameIsValid(eventName) {
		fmt.Fprintf(os.Stderr, "handler for event %v could not be added because %v is not a valid event name\n", eventName, eventName)
		return
	}

	handlers, ok := gw.handlers[eventName]
	if ok {
		handlers = append(handlers, fn)
	} else {
		handlers = []interface{}{fn}
	}

	gw.handlers[eventName] = handlers
}

func (gw *Gateway) AddRawHandlerFunc(fn RawHandlerFunc) {
	gw.handlerMu.Lock()
	defer gw.handlerMu.Unlock()

	gw.rawHandlers = append(gw.rawHandlers, fn)
}

// -- internal API --

func (gw *Gateway) openConn() error {
	if gw.conn != nil {
		return ErrAlreadyOpenConnection
	}

	if gw.gateway.URL == "" {
		g, err := gw.rest.GetGatewayAuthed()
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not get gateway info: %v\n", err)
			return err
		}

		gw.gateway = *g
	}

	gw.wsMu.Lock()
	defer gw.wsMu.Unlock()

	var headers http.Header
	conn, _, err := websocket.DefaultDialer.Dial(gw.gateway.URL, headers)
	if err != nil {
		return err
	}

	gw.conn = conn

	gw.statusMu.Lock()
	gw.Status = WAITING_FOR_HELLO
	gw.statusMu.Unlock()

	var helloPayload HelloPayload
	err = gw.conn.ReadJSON(&helloPayload)
	if err != nil || helloPayload.Op != HELLO {
		gw.statusMu.Lock()
		gw.Status = DOWN
		gw.statusMu.Unlock()
		return err
	}

	gw.statusMu.Lock()
	gw.Status = HELLO_RECEIVED
	gw.statusMu.Unlock()

	gw.heartbeatInterval = time.Duration(helloPayload.Data.HeartbeatInterval) * time.Millisecond

	return nil
}

func (gw *Gateway) identify(ctx context.Context) error {
	if gw.conn == nil {
		return ErrNilConnection
	}

	gw.wsMu.Lock()
	defer gw.wsMu.Unlock()

	gw.statusMu.Lock()
	gw.Status = IDENTIFYING
	gw.statusMu.Unlock()

	err := gw.conn.WriteJSON(gw.identifyPayload)
	if err != nil {
		gw.statusMu.Lock()
		gw.Status = DOWN
		gw.statusMu.Unlock()
		return nil
	}

	gw.statusMu.Lock()
	gw.Status = WAITING_FOR_READY
	gw.statusMu.Unlock()

	var readyResp RawGatewayPayload
	err = gw.conn.ReadJSON(&readyResp)
	// TODO: give "not ready" its own error
	if err != nil {
		gw.statusMu.Lock()
		gw.Status = DOWN
		gw.statusMu.Unlock()
		return err
	}

	if readyResp.Type != "READY" {
		return fmt.Errorf("%w: %v", ErrNotReady, readyResp.Type)
	}

	gw.statusMu.Lock()
	gw.Status = UP
	gw.statusMu.Unlock()

	go gw.processMessage(ctx, readyResp)

	return nil
}

func (gw *Gateway) resume(ctx context.Context) error {
	if gw.conn == nil {
		return ErrNilConnection
	}

	gw.wsMu.Lock()
	defer gw.wsMu.Unlock()

	gw.statusMu.Lock()
	gw.Status = RESUMING
	gw.statusMu.Unlock()

	resumePayload := ResumePayload{
		Op: RESUME,
		Data: ResumePayloadData{
			Token:     gw.token,
			SessionID: gw.SessionID,
			Sequence:  gw.Sequence,
		},
	}

	err := gw.conn.WriteJSON(resumePayload)
	if err != nil {
		return err
	}

	var resumedResp RawGatewayPayload
	err = gw.conn.ReadJSON(&resumedResp)
	if err != nil {
		gw.statusMu.Lock()
		gw.Status = DOWN
		gw.statusMu.Unlock()
		return err
	}

	if resumedResp.Op == INVALID_SESSION {
		gw.statusMu.Lock()
		gw.Status = DOWN
		gw.statusMu.Unlock()
		// TODO: this is bad. how can i do it better?
		return &websocket.CloseError{
			Code: INVALID_SESSION,
			Text: "invalid session",
		}
	}

	gw.statusMu.Lock()
	gw.Status = UP
	gw.statusMu.Unlock()

	go gw.processMessage(ctx, resumedResp)

	return nil
}

// TODO: add a way to add a custom reconnect function. it would allow for coordinating ratelimits across multiple shard/clusters to prevent getting rate limited

// reconnects to the gateway
//
// reconnect() uses exponential backoff for reconnect attempts so the user doesn't burn through their daily identifies accidentally
// backoff starts at 1 second and increases by the power of two for each failed attempt, until reaching 300 secinds.
// timer resets once a successful connection has been established.
func (gw *Gateway) reconnect(ctx context.Context) error {
	wait := 1 * time.Second

	// TODO: is this a bad idea lol
	gw.statusMu.Lock()
	if gw.Status == RECONNECTING {
		return errors.New("already reconnecting")
	}

	gw.Status = RECONNECTING
	gw.statusMu.Unlock()

	for {
		gw.cleanup(false)

		err := gw.Up(ctx)
		if err == nil {
			return nil
		}

		if e, ok := err.(*websocket.CloseError); ok {
			if !CloseCode(e.Code).CanReconnect() {

				return err
			}

			if e.Code == INVALID_SESSION {
				gw.cleanup(true)
			}
		}

		wait *= 2
		if wait >= MAX_BACKOFF_SECS {
			wait = MAX_BACKOFF_SECS
		}

		log.Printf("Reconnect failed. Retrying in %v\n", wait)
		<-time.After(wait)
	}
}

func (gw *Gateway) heartbeat() error {
	if gw.conn == nil {
		return ErrNilConnection
	}

	gw.wsMu.Lock()
	defer gw.wsMu.Unlock()

	seq := atomic.LoadInt64(&gw.Sequence)

	heartbeatPayload := HeartbeatPayload{
		Op:   HEARTBEAT,
		Data: seq,
	}

	err := gw.conn.WriteJSON(heartbeatPayload)
	if err != nil {
		return err
	}

	gw.lastHeartbeatSent = time.Now()

	return nil
}

func (gw *Gateway) startHeartbeatLoop(ctx context.Context) error {
	if gw.heartbeatInterval == 0 {
		return ErrNilHeartbeatInterval
	}

	t := time.NewTicker(gw.heartbeatInterval)

	for {

		err := gw.heartbeat()
		if err != nil {
			if errors.Is(err, ErrNilConnection) {
				return ErrNilConnection
			}
		}

		// The gateway had failed to ack the last 5 heartbeats,
		// which means the connection probably dropped and a
		// reconnect should be attempted
		if !gw.lastHeartbeatAck.IsZero() &&
			time.Now().Sub(gw.lastHeartbeatAck) > gw.heartbeatInterval*HEARTBEAT_ACK_FAILURE_BEFORE_RECONNECT {
			gw.reconnect(ctx)
			return nil
		}

		select {
		case <-t.C:
		case <-ctx.Done():
			return nil
		}
	}
}

func (gw *Gateway) processMessage(ctx context.Context, msg RawGatewayPayload) error {
	if msg.Sequence != 0 {
		atomic.StoreInt64(&gw.Sequence, msg.Sequence)
	}

	gw.handlerMu.RLock()
	rawHandlers := gw.rawHandlers
	gw.handlerMu.RUnlock()

	for _, rawHandlerFn := range rawHandlers {
		go rawHandlerFn(gw, &msg)
	}

	switch msg.Op {
	case DISPATCH:
		eventPaylpad := eventNameToPayload(msg.Type)
		if eventPaylpad == nil {
			fmt.Fprintf(os.Stderr, "Got event %v but did not have a corresponding struct to deserialize into\n", msg.Type)
			return ErrNoDispatchEventStruct
		}

		err := json.Unmarshal(msg.Data, &eventPaylpad)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to unmarshal payload for event %v: %v\n", msg.Type, err)
			return err
		}

		gw.handlerMu.RLock()
		handlers := gw.handlers[msg.Type]
		gw.handlerMu.RUnlock()

		go execHandlerFunc(gw, handlers, msg.Type, eventPaylpad)

	case HEARTBEAT:
		seq := atomic.LoadInt64(&gw.Sequence)
		gw.wsMu.Lock()
		gw.conn.WriteJSON(HeartbeatPayload{HEARTBEAT, seq})
		gw.wsMu.Unlock()

	case RECONNECT:
		return gw.reconnect(ctx)

	case INVALID_SESSION:
		gw.SessionID = ""
		gw.Sequence = 0
		return gw.reconnect(ctx)

	case HEARTBEAT_ACK:
		gw.lastHeartbeatAck = time.Now()
	}

	return nil
}

func (gw *Gateway) startMessageProcessorLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if gw.conn == nil {
			return ErrNilConnection
		}

		var msg RawGatewayPayload
		err := gw.conn.ReadJSON(&msg)

		if err != nil {
			closeErr, ok := err.(*websocket.CloseError)
			if ok {
				switch code := CloseCode(closeErr.Code); code {
				case UNKNOWN,
					UNKNOWN_OPCODE,
					DECODE_ERROR,
					NOT_AUTHENTICATED,
					ALREADY_AUTHENTICATED,
					RATE_LIMITED,
					SESSION_TIMEOUT:

					fmt.Fprintf(os.Stderr, "[WARN] websocket closed with code %v and will attempt to reconnect", code)
					return gw.reconnect(ctx)

				case INVALID_SESSION:
					fmt.Fprintln(os.Stderr, "[WARN] invalid session, reidentifying...")
					return gw.reconnect(ctx)

				default:
					fmt.Fprintf(os.Stderr, "[FATAL] websocket closed with code %v and CANNOT reconnect", code)
					gw.Down(ctx)
					return err
				}
			} else {
				fmt.Fprintf(os.Stderr, "[???] There was an error reading a message from the gateway: %v\n", err)
				return gw.reconnect(ctx)
			}
		}

		go gw.processMessage(ctx, msg)
	}
}

func (gw *Gateway) cleanup(clearSidSeq bool) {
	if clearSidSeq {
		gw.SessionID = ""
		gw.Sequence = 0
	}

	if gw.cancelHB != nil {
		gw.cancelHB()
	}

	if gw.cancelMP != nil {
		gw.cancelMP()
	}

	// TODO: is this bad to do?
	gw.conn = nil
}
