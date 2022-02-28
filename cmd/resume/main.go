package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/benricheson101/simplegateway/pkg/gateway"
)

func main() {
	gw := gateway.New(os.Getenv("DISCORD_TOKEN"))

	sid := os.Args[1]
	seq, _ := strconv.ParseInt(os.Args[2], 10, 0)

	gw.SessionID = sid
	gw.Sequence = seq

	ctx := context.Background()
	err := gw.Up(ctx)
	if err != nil {
		fmt.Println("couldnt resume session:", err)
		os.Exit(1)
	}
	fmt.Printf("sid = %v seq = %v\n", gw.SessionID, gw.Sequence)

	select {}
}
