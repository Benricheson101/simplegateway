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

	gw.HandleFunc(func(gw *gateway.Gateway, ev *gateway.GenericDispatchPayload) {
		if ev.Type == "READY" {
			fmt.Printf("sid = %v seq = %v\n", gw.SessionID, gw.Sequence)
		}

		fmt.Println(string(ev.Data))
	})

	ctx := context.Background()
	if err := gw.Up(ctx); err != nil {
		log.Fatalln(err)
	}

	select {}
}
