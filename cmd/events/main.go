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
		fmt.Printf("Ready as %v#%v with %d guilds\n", r.User.Username, r.User.Discriminator, len(r.Guilds))
	})

	gw.AddHandleFunc(func(gw *gateway.Gateway, r *gateway.Ready) {
		fmt.Println("Second READY handler")
	})

	gw.AddHandleFunc(func(gw *gateway.Gateway, ev *gateway.MessageCreate) {
		fmt.Println("MESSAGE CREATE")
	})

	gw.AddHandleFunc(func(gw *gateway.Gateway, gc *gateway.GuildCreate) {
		fmt.Printf("name = %v, ownerID = %v\n", gc.Name, gc.OwnerID)
	})

	ctx := context.Background()
	if err := gw.Up(ctx); err != nil {
		log.Fatalln(err)
	}

	select {}
}
