package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/benricheson101/simplegateway/pkg/gateway"
	"github.com/benricheson101/simplegateway/pkg/gateway/intents"
)

func main() {
	gw := gateway.New(
		os.Getenv("DISCORD_TOKEN"),
		gateway.WithIntents(intents.New(
			intents.GUILDS,
			intents.GUILD_MESSAGES,
		)),
		gateway.WithShard(0, 1),
		gateway.WithPresence(gateway.IdentifyPayloadDataPresence{
			Status: gateway.ONLINE,
			Activities: &[]gateway.PresenceActivity{
				{
					Name: "I work!",
					Type: gateway.PLAYING,
				},
			},
		}),
	)

	gw.AddHandlerFunc(func(gw *gateway.Gateway, r *gateway.Ready) {
		fmt.Printf("Ready as %v#%v\n", r.User.Username, r.User.Discriminator)
	})

	ctx := context.Background()
	err := gw.Up(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open gateway connection: %v\n", err)
		os.Exit(1)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig

	gw.Down(ctx)
}
