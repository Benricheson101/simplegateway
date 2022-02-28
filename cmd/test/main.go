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
		Intents: 583,
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

	ctx := context.Background()
	if err := gw.Up(ctx); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("sid = %v seq = %v\n", gw.SessionID, gw.Sequence)

	select {}
}
