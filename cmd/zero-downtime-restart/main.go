package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/benricheson101/simplegateway/pkg/gateway"
)

const SHARD_STATE_FILE = "shard_state.json"

type ShardResumeState struct {
	SessionID string `json:"session_id"`
	Sequence  int64  `json:"sequence"`
}

type ShardSessions struct {
	sync.Mutex `json:"-"`

	Shards map[int]*ShardResumeState `json:"shards"`
}

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

	var state *ShardSessions
	s := getState()
	if s == nil {
		state = &ShardSessions{
			Shards: make(map[int]*ShardResumeState),
		}
	} else {
		state = s

		a := state.Shards[0]
		gw.Sequence = a.Sequence
		gw.SessionID = a.SessionID
	}

	gw.AddRawHandlerFunc(func(g *gateway.Gateway, rgp *gateway.RawGatewayPayload) {
		if rgp.Type != "READY" && rgp.Sequence != 0 {
			state.Lock()
			s := state.Shards[0]
			s.Sequence = rgp.Sequence
			state.Unlock()
		}

		if rgp.Type != "" {
			fmt.Println(rgp.Type)
		}
	})

	gw.AddHandleFunc(func(gw *gateway.Gateway, r *gateway.Ready) {
		// fmt.Printf("sid = %v seq = %v\n", gw.SessionID, gw.Sequence)

		state.Lock()
		state.Shards[0] = &ShardResumeState{
			SessionID: r.SessionID,
			Sequence:  0,
		}
		state.Unlock()
	})

	gw.AddHandleFunc(func(gw *gateway.Gateway, mc *gateway.MessageCreate) {
		if mc.Content == "!save" && mc.Author.ID == "255834596766253057" {
			saveState(state)
		}
	})

	ctx := context.Background()

	if gw.SessionID != "" && gw.Sequence != 0 {
		fmt.Println("Attempting to resume session")
		if err := gw.TryResume(ctx); err != nil {
			if errors.Is(err, gateway.ErrResumeFail) {
				fmt.Println("failed to resume session. IDENTIFYing")

				gw.Sequence = 0
				gw.SessionID = ""
				if err = gw.Up(ctx); err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		fmt.Println("No stored sessionID or sequence. IDENTIFYing")
		if err := gw.Up(ctx); err != nil {
			log.Fatalln(err)
		}
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	<-sig
	saveState(state)
}

func saveState(state *ShardSessions) error {
	state.Lock()
	j, _ := json.MarshalIndent(&state, "", "  ")
	state.Unlock()
	err := os.WriteFile(SHARD_STATE_FILE, j, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write state file: %v\n", err)
		return err
	}

	return nil
}

func getState() (ret *ShardSessions) {
	_, err := os.Stat(SHARD_STATE_FILE)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "stat on state file %v failed: %v\n", SHARD_STATE_FILE, err)
		return nil
	}

	f, err := os.ReadFile(SHARD_STATE_FILE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open state file %v: %v\n", SHARD_STATE_FILE, err)
		return nil
	}

	if err = json.Unmarshal(f, &ret); err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal state file content: %v\n", err)
		return nil
	}

	return
}
