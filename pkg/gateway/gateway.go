package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benricheson101/simplegateway/pkg/rest"
	"github.com/gorilla/websocket"
)

var (
	ErrOpenConnection   = errors.New("there is already an open websocket connection")
	ErrDialWebsocket    = errors.New("failed to dial websocket")
	ErrReadMessageFail  = errors.New("failed to read message")
	ErrWriteMessageFail = errors.New("failed to write message")
	ErrResumeFail       = errors.New("failed to resume session")
)

type EventHandlerFunc func(gw *Gateway, ev *GenericDispatchPayload)

type Gateway struct {
	sync.RWMutex

	rc                *rest.RestClient
	wsMu              sync.Mutex
	conn              *websocket.Conn
	heartbeatInterval time.Duration
	heartbeatTicker   *time.Ticker
	lastHeartbeatAck  time.Time

	MaxReconnectAttempts int

	// TODO: is it a bad idea to pass around context like this?
	context context.Context
	cancel  context.CancelFunc

	Identify IdentifyPayloadData

	token string

	Sequence  int64
	SessionID string

	handlerFunc EventHandlerFunc
}

type Shard [2]int

func New(token string) *Gateway {
	return &Gateway{
		token:                token,
		rc:                   rest.New(token),
		MaxReconnectAttempts: 3,
	}
}

func (gw *Gateway) Up(ctx context.Context) error {
	gw.Lock()
	defer gw.Unlock()

	return gw.connect(ctx)
}

func (gw *Gateway) resume() error {
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

	var m GenericDispatchPayload
	err = json.Unmarshal(msg, &m)
	if err != nil {
		return ErrReadMessageFail
	}

	if m.Op == 9 {
		return ErrResumeFail
	} else {
		gw.onMessage(msg)
	}

	go gw.startHeartbeat(gw.context)
	go gw.readMessages(gw.context)

	return nil
}

func (gw *Gateway) TryResume(ctx context.Context) error {
	gw.Lock()
	defer gw.Unlock()

	if gw.conn != nil {
		return ErrOpenConnection
	}

	ctx, cancel := context.WithCancel(ctx)
	gw.context = ctx
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

	return gw.resume()
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
					gw.connect(ctx)

				case INVALID_SESSION:
					gw.Sequence = 0
					gw.SessionID = ""
					gw.connect(ctx)

				default:
					fmt.Fprintf(os.Stderr, "[FATAL] websocket closed with code %v and CANNOT reconnect", code)
					gw.Down()
					return
				}
			}

			gw.connect(ctx)
			return
		}

		go gw.onMessage(msg)
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
	// TODO: is this a bad idea?
	gw.conn = nil
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
	} else {
		err := gw.resume()
		// if the connection cannot be resumed, re-identify
		if err != nil {
			gw.Sequence = 0
			gw.SessionID = ""

			return gw.connect(ctx)
		}
	}

	go gw.startHeartbeat(ctx)
	go gw.readMessages(ctx)

	return nil
}

func (gw *Gateway) startHeartbeat(ctx context.Context) {
	t := time.NewTicker(gw.heartbeatInterval)
	defer t.Stop()

	for {
		seq := atomic.LoadInt64(&gw.Sequence)
		gw.wsMu.Lock()
		gw.conn.WriteJSON(HeartbeatPayload{Op: HEARTBEAT, Data: seq})
		gw.wsMu.Unlock()

		select {
		case <-t.C:
		case <-ctx.Done():
			return
		}
	}
}

func (gw *Gateway) onMessage(msg []byte) {
	var event GenericDispatchPayload
	if err := json.Unmarshal(msg, &event); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse JSON from event: %v\n", err)
		return
	}

	if event.Sequence != 0 {
		atomic.StoreInt64(&gw.Sequence, event.Sequence)
	}

	switch event.Op {
	case DISPATCH:
		if event.Type == "READY" {
			var data ReadyPayloadData
			if err := json.Unmarshal(event.Data, &data); err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse READY body: %v\n", err)
				return
			}

			gw.SessionID = data.SessionID
		}

		go gw.handlerFunc(gw, &event)

	case HEARTBEAT:
		seq := atomic.LoadInt64(&gw.Sequence)
		gw.wsMu.Lock()
		gw.conn.WriteJSON(HeartbeatPayload{Op: HEARTBEAT, Data: seq})
		gw.wsMu.Unlock()

	case RECONNECT:
		gw.connect(gw.context)

	case INVALID_SESSION:
		gw.SessionID = ""
		gw.Sequence = 0
		gw.Up(gw.context)

	case HEARTBEAT_ACK:
		gw.lastHeartbeatAck = time.Now().UTC()
	}
}

func (gw *Gateway) HandleFunc(h EventHandlerFunc) {
	gw.handlerFunc = h
}
