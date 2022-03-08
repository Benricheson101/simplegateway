package gateway

import (
	"sync/atomic"

	"github.com/benricheson101/simplegateway/pkg/gateway/intents"
	"github.com/benricheson101/simplegateway/pkg/rest"
)

type GatewayOption func(*Gateway)

// Sets the SessionID and Sequence. Useful if you would like to attempt to
// resume the connection instead of re-identifying
func WithSession(sessionID string, sequence int64) GatewayOption {
	return func(gw *Gateway) {
		gw.SessionID = sessionID
		atomic.StoreInt64(&gw.Sequence, sequence)
	}
}

// Sets the presence of the bot in the identify payload
func WithPresence(presence IdentifyPayloadDataPresence) GatewayOption {
	return func(gw *Gateway) {
		gw.identifyPayload.Data.Presence = &presence
	}
}

// Sets the status of the bot user
func WithStatus(status UserStatus) GatewayOption {
	return func(gw *Gateway) {
		gw.identifyPayload.Data.Presence.Status = status
	}
}

// Sets the activity of the user
func WithActivity(kind ActivityType, name string) GatewayOption {
	return func(gw *Gateway) {
		gw.identifyPayload.Data.Presence.Activities = &[]PresenceActivity{
			{
				Name: name,
				Type: kind,
			},
		}
	}
}

// Assigns the created client to a certain shard
func WithShard(shardID, totalShards int) GatewayOption {
	return func(gw *Gateway) {
		gw.identifyPayload.Data.Shard = &GatewayShard{shardID, totalShards}
	}
}

// Sets the information about the gateway, like
// the websocket URL that the bot will use
func WithGateway(gateway rest.GatewayBot) GatewayOption {
	return func(gw *Gateway) {
		gw.gateway = gateway
	}
}

// Sets intents in the identify payload
func WithIntents(intents intents.GatewayIntent) GatewayOption {
	return func(gw *Gateway) {
		gw.identifyPayload.Data.Intents = int(intents)
	}
}

// TODO: func WithLogger() GatewayOption {}

// TODO: custom locks. this would allow for distributed locking and synchronizing shard startup across multiple machines
// TODO: func WithCustomLock(lock <some-lock-interface>) GatewayOption {}
