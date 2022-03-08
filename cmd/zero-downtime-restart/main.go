package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/benricheson101/simplegateway/pkg/gateway"
	"github.com/benricheson101/simplegateway/pkg/gateway/intents"
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
	var state *ShardSessions
	s := getState()
	if s == nil {
		state = &ShardSessions{
			Shards: make(map[int]*ShardResumeState),
		}
	} else {
		state = s
	}

	sessionID := ""
	sequence := int64(0)

	if sh0, ok := state.Shards[0]; ok && sh0.Sequence != 0 && sh0.SessionID != "" {
		sessionID = sh0.SessionID
		sequence = sh0.Sequence
	}

	gw := gateway.New(
		os.Getenv("DISCORD_TOKEN"),
		gateway.WithIntents(intents.AllWithoutPrivileged().Add(intents.GUILD_MESSAGES)),
		gateway.WithShard(0, 1),
		gateway.WithSession(sessionID, sequence),
		gateway.WithStatus(gateway.DO_NOT_DISTURB),
		gateway.WithActivity(gateway.WATCHING, "i work :D"),
	)

	gw.AddRawHandlerFunc(func(g *gateway.Gateway, pl *gateway.RawGatewayPayload) {
		if pl.Type != "READY" && pl.Sequence != 0 {
			state.Lock()
			s := state.Shards[0]
			s.Sequence = pl.Sequence
			state.Unlock()
		}

		if pl.Type != "" {
			fmt.Println(pl.Type)
		}
	})

	gw.AddHandlerFunc(func(gw *gateway.Gateway, r *gateway.Ready) {
		state.Lock()
		state.Shards[0] = &ShardResumeState{
			SessionID: r.SessionID,
			Sequence:  0,
		}
		state.Unlock()
	})

	ctx := context.Background()

	if gw.SessionID != "" && gw.Sequence != 0 {
		fmt.Println("stored sessionID and sequence present. attempting to resume")
		err := gw.TryResume(ctx)
		if err != nil {
			fmt.Println("resume failed, attempting to identify")
			gw.SessionID = ""
			gw.Sequence = 0
			err = gw.Up(ctx)
			if err != nil {
				fmt.Println("failed to resume or become ready:", err)
				os.Exit(1)
			}
		}
	} else {
		err := gw.Up(ctx)
		if err != nil {
			fmt.Println("couldnt start bot:", err)
			os.Exit(1)
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
