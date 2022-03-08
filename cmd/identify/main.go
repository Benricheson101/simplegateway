package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/benricheson101/simplegateway/pkg/gateway"
	"github.com/benricheson101/simplegateway/pkg/gateway/intents"
	"github.com/gorilla/websocket"
)

var (
	seq int64
	sid string
)

func main() {
	gw := gateway.New(
		os.Getenv("DISCORD_TOKEN"),
		gateway.WithIntents(intents.AllWithoutPrivileged()),
	)

	gw.AddHandlerFunc(func(gw *gateway.Gateway, r *gateway.Ready) {
		sid = r.SessionID
	})

	gw.AddRawHandlerFunc(func(gw *gateway.Gateway, pl *gateway.RawGatewayPayload) {
		if t := pl.Type; t != "" {
			fmt.Println("<-", t)
		}

		if s := pl.Sequence; s != 0 {
			seq = s
		}
	})

	err := gw.Up(context.Background())
	if err != nil {
		if e, ok := err.(*websocket.CloseError); ok {
			fmt.Printf("CloseError code=%v text=%v\n", e.Code, e.Text)
		} else {
			fmt.Printf("not CloseError, err=%v\n", err)
		}

		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	gw.DownWithCode(context.Background(), websocket.CloseAbnormalClosure)

	fmt.Println("resume command:")
	fmt.Printf("go run cmd/resume/main.go -sid %v -seq %v\n", sid, seq)
}
