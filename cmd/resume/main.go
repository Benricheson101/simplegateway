package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/benricheson101/simplegateway/pkg/gateway"
)

var (
	seq int64
	sid string
)

func init() {
	flag.StringVar(&sid, "sid", "", "sessionID")
	flag.Int64Var(&seq, "seq", 0, "sequence")
	flag.Parse()
}

func main() {
	gw := gateway.New(
		os.Getenv("DISCORD_TOKEN"),
		gateway.WithSession(sid, seq),
	)

	gw.AddRawHandlerFunc(func(gw *gateway.Gateway, pl *gateway.RawGatewayPayload) {
		fmt.Println("<-", pl.Op, pl.Type)
	})

	ctx := context.Background()
	err := gw.Up(ctx)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	// err := gw.TryResume(ctx)
	// if err != nil {
	// 	fmt.Println("failed to resume:", err)
	// 	os.Exit(1)
	// }

	fmt.Println("waiting for ctrl c")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	gw.Down(context.Background())
}
