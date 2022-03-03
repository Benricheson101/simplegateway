package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/benricheson101/simplegateway/pkg/gateway"
)

func main() {
	gw := gateway.New(os.Getenv("DISCORD_TOKEN"))
	gw.Identify = gateway.IdentifyPayloadData{
		Shard:   &gateway.GatewayShard{0, 1},
		Intents: 32511,
		Properties: gateway.IdentifyPayloadDataProperties{
			OS:      "MacOS",
			Browser: "simplegateway",
			Device:  "simplegateway",
		},
		Presence: &gateway.IdentifyPayloadDataPresence{
			Status: gateway.DO_NOT_DISTURB,
			Activities: &[]gateway.PresenceActivity{
				{Name: "i work :D", Type: gateway.PLAYING},
			},
		},
	}

	gw.AddHandleFunc(func(gw *gateway.Gateway, r *gateway.Ready) {
		fmt.Printf("sid = %v seq = %v\n", gw.SessionID, gw.Sequence)
	})

	gw.AddHandleFunc(func(gw *gateway.Gateway, mc *gateway.MessageCreate) {
		fmt.Printf("[MESSAGE_CREATE] %v\n", mc.Content)
	})

	ctx := context.Background()
	if err := gw.Up(ctx); err != nil {
		log.Fatalln(err)
	}

	select {}
}
