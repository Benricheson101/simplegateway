package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benricheson101/simplegateway/pkg/rest"
	"github.com/gorilla/websocket"
)

// The number of heartbeat acks allowed to be missed before reconnecting
const HEARTBEAT_FAILURES = 5

var (
	ErrOpenConnection   = errors.New("there is already an open websocket connection")
	ErrDialWebsocket    = errors.New("failed to dial websocket")
	ErrReadMessageFail  = errors.New("failed to read message")
	ErrWriteMessageFail = errors.New("failed to write message")
	ErrResumeFail       = errors.New("failed to resume session")
)

type RawHandlerFunc func(*Gateway, *RawGatewayPayload)

type Gateway struct {
	sync.Mutex

	rc        *rest.RestClient
	wsMu      sync.Mutex
	handlerMu sync.Mutex

	conn              *websocket.Conn
	heartbeatInterval time.Duration
	heartbeatTicker   *time.Ticker
	lastHeartbeatAck  time.Time
	lastHeartbeatSent time.Time

	MaxReconnectAttempts int

	cancel   context.CancelFunc
	cancelHB context.CancelFunc
	cancelRM context.CancelFunc

	Identify IdentifyPayloadData

	token string

	Sequence  int64
	SessionID string

	handlers    map[string][]interface{}
	rawHandlers []RawHandlerFunc
}

// type Shard [2]int

func New(token string) *Gateway {
	return &Gateway{
		handlers:             make(map[string][]interface{}),
		token:                token,
		rc:                   rest.New(token),
		MaxReconnectAttempts: 3,
		lastHeartbeatAck:     time.Now(),
		lastHeartbeatSent:    time.Now(),
	}
}

func (gw *Gateway) Up(ctx context.Context) error {
	gw.Lock()
	defer gw.Unlock()

	return gw.connect(ctx)
}

func (gw *Gateway) resume(ctx context.Context) error {
	resumePayload := ResumePayload{
		Op: RESUME,
		Data: ResumePayloadData{
			Token:     gw.token,
			SessionID: gw.SessionID,
			Sequence:  gw.Sequence,
		},
	}

	gw.wsMu.Lock()
	err := gw.conn.WriteJSON(&resumePayload)
	gw.wsMu.Unlock()

	if err != nil {
		fmt.Println("failed to send resume packet")
		return err
	}

	_, msg, err := gw.conn.ReadMessage()
	if err != nil {
		return ErrReadMessageFail
	}

	var m RawGatewayPayload
	err = json.Unmarshal(msg, &m)
	if err != nil {
		return ErrReadMessageFail
	}

	if m.Op == 9 {
		return ErrResumeFail
	} else {
		gw.onMessage(ctx, msg)
	}

	go gw.startHeartbeat(ctx)
	go gw.readMessages(ctx)

	return nil
}

func (gw *Gateway) TryResume(ctx context.Context) error {
	gw.Lock()
	defer gw.Unlock()

	if gw.conn != nil {
		return ErrOpenConnection
	}

	ctx, cancel := context.WithCancel(ctx)
	gw.cancel = cancel

	var header http.Header
	conn, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg", header)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDialWebsocket, err)
	}

	gw.conn = conn

	var hello HelloPayload
	err = gw.conn.ReadJSON(&hello)
	if err != nil {
		return fmt.Errorf("%s, %v", ErrReadMessageFail, err)
	}
	gw.heartbeatInterval = time.Millisecond * time.Duration(hello.Data.HeartbeatInterval)

	return gw.resume(ctx)
}

func (gw *Gateway) readMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		_, msg, err := gw.conn.ReadMessage()
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
					gw.reconnect(ctx)

				case INVALID_SESSION:
					gw.Sequence = 0
					gw.SessionID = ""
					gw.reconnect(ctx)

				default:
					fmt.Fprintf(os.Stderr, "[FATAL] websocket closed with code %v and CANNOT reconnect", code)
					gw.Down()
					return
				}
			}

			fmt.Println("error reading message:", err)
			// gw.cancelHB()
			return
		}

		go gw.onMessage(ctx, msg)
	}
}

func (gw *Gateway) Down() error {
	return gw.DownWithCode(websocket.CloseNormalClosure)
}

func (gw *Gateway) DownWithCode(code int) error {
	gw.Lock()
	defer gw.Unlock()

	gw.cancel()
	return gw.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, ""))
}

func (gw *Gateway) connect(ctx context.Context) error {

	var header http.Header
	conn, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg", header)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDialWebsocket, err)
	}
	gw.conn = conn
	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("websocket close handler code=%v text=%v\n", code, text)
		return nil
	})

	var hello HelloPayload
	err = gw.conn.ReadJSON(&hello)
	if err != nil {
		return fmt.Errorf("%s, %v", ErrReadMessageFail, err)
	}
	gw.heartbeatInterval = time.Millisecond * time.Duration(hello.Data.HeartbeatInterval)

	if gw.SessionID == "" || gw.Sequence == 0 {
		gw.Identify.Token = gw.token
		gw.wsMu.Lock()
		err := gw.conn.WriteJSON(IdentifyPayload{
			Op:   IDENTIFY,
			Data: gw.Identify,
		})
		gw.wsMu.Unlock()

		if err != nil {
			return fmt.Errorf("%s: %v", ErrWriteMessageFail, err)
		}

		hbCtx, cancel := context.WithCancel(ctx)
		gw.cancelHB = cancel
		go gw.startHeartbeat(hbCtx)

		rmCtx, cancel := context.WithCancel(ctx)
		gw.cancelRM = cancel
		go gw.readMessages(rmCtx)
	} else {
		err := gw.resume(ctx)
		// if the connection cannot be resumed, re-identify
		if err != nil {
			gw.Sequence = 0
			gw.SessionID = ""

			return gw.connect(ctx)
		}
	}

	return nil
}

func (gw *Gateway) reconnect(ctx context.Context) {
	// TODO: move reconnect logic here
	// TODO: set max reconnect?

	fmt.Println("pretend the bot is reconnecting now")
}

func (gw *Gateway) startHeartbeat(ctx context.Context) {
	t := time.NewTicker(gw.heartbeatInterval)
	defer t.Stop()

	for {
		seq := atomic.LoadInt64(&gw.Sequence)
		gw.wsMu.Lock()
		err := gw.conn.WriteJSON(HeartbeatPayload{Op: HEARTBEAT, Data: seq})
		gw.wsMu.Unlock()

		if time.Now().Sub(gw.lastHeartbeatAck) > (gw.heartbeatInterval * time.Duration(HEARTBEAT_FAILURES) * time.Millisecond) {
			fmt.Printf("discord failed to ack the last %v heartbeats. reconnecting...\n", HEARTBEAT_FAILURES)
			gw.DownWithCode(websocket.CloseServiceRestart)
		} else if err != nil {
			log.Println("failed to send heartbeat packet:", err)
		} else {
			// log.Println("successful heartbeat")
			gw.lastHeartbeatSent = time.Now()
		}

		select {
		case <-t.C:
		case <-ctx.Done():
			return
		}
	}
}

func (gw *Gateway) onMessage(ctx context.Context, msg []byte) {
	var event RawGatewayPayload
	if err := json.Unmarshal(msg, &event); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse JSON from event: %v\n", err)
		return
	}

	if event.Sequence != 0 {
		atomic.StoreInt64(&gw.Sequence, event.Sequence)
	}

	for _, rawHandlerFn := range gw.rawHandlers {
		go rawHandlerFn(gw, &event)
	}

	switch event.Op {
	case DISPATCH:
		eventPayload := eventNameToPayload(event.Type)
		if eventPayload == nil {
			fmt.Fprintf(os.Stderr, "Got event %v but did not have a corresponding struct to deserialize into\n", event.Type)
			return
		}

		err := json.Unmarshal(event.Data, &eventPayload)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to unmarshal payload for event %v: %v\n", event.Type, err)
			return
		}

		if event.Type == "READY" {
			gw.SessionID = eventPayload.(*Ready).SessionID
		}

		// TODO: is it possible for this to deadlock?
		gw.handlerMu.Lock()
		handlers := gw.handlers[event.Type]
		gw.handlerMu.Unlock()

		execHandlerFunc(gw, handlers, event.Type, eventPayload)

	case HEARTBEAT:
		seq := atomic.LoadInt64(&gw.Sequence)
		gw.wsMu.Lock()
		gw.conn.WriteJSON(HeartbeatPayload{Op: HEARTBEAT, Data: seq})
		gw.wsMu.Unlock()

	case RECONNECT:
		gw.connect(ctx)

	case INVALID_SESSION:
		gw.SessionID = ""
		gw.Sequence = 0
		gw.Up(ctx)

	case HEARTBEAT_ACK:
		gw.lastHeartbeatAck = time.Now().UTC()
	}
}

func (gw *Gateway) AddHandleFunc(fn interface{}) {
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
