package gateway

import "encoding/json"

type Opcode int

const (
	// An event was dispatched.
	DISPATCH Opcode = 0
	// Fired periodically by the client to keep the connection alive.
	HEARTBEAT = 1
	// Starts a new session during the initial handshake.
	IDENTIFY = 2
	// Update the client's presence.
	PRESENCE_UPDATE = 3
	// Used to join/leave or move between voice channels.
	VOICE_STATUS_UPDATE = 4
	// Resume a previous session that was disconnected.
	RESUME = 6
	// You should attempt to reconnect and resume immediately.
	RECONNECT = 7
	// Request information about offline guild members in a large guild.
	REQUEST_GUILD_MEMBERS = 8
	// The session has been invalidated. You should reconnect and identify/resume accordingly.
	INVALID_SESSION = 9
	// Sent immediately after connecting, contains the heartbeat_interval to use.
	HELLO = 10
	// Sent in response to receiving a heartbeat to acknowledge that it has been received.
	HEARTBEAT_ACK = 11
)

type CloseCode int

const (
	// Unknown error	We're not sure what went wrong. Try reconnecting?	true
	UNKNOWN CloseCode = 4000
	// You sent an invalid Gateway opcode or an invalid payload for an opcode. Don't do that!	true
	UNKNOWN_OPCODE = 4001
	// You sent an invalid payload to us. Don't do that!	true
	DECODE_ERROR = 4002
	// You sent us a payload prior to identifying.	true
	NOT_AUTHENTICATED = 4003
	// The account token sent with your identify payload is incorrect.	false
	AUTHENTICATION_FAILED = 4004
	// You sent more than one identify payload. Don't do that!	true
	ALREADY_AUTHENTICATED = 4005
	// The sequence sent when resuming the session was invalid. Reconnect and start a new session.	true
	INVALID_SEQUENCE = 4007
	// Woah nelly! You're sending payloads to us too quickly. Slow it down! You will be disconnected on receiving this.	true
	RATE_LIMITED = 4008
	// Your session timed out. Reconnect and start a new one.	true
	SESSION_TIMEOUT = 4009
	// You sent us an invalid shard when identifying.	false
	INVALID_SHARD = 4010
	// The session would have handled too many guilds - you are required to shard your connection in order to connect.	false
	SHARDING_REQUIRED = 4011
	// You sent an invalid version for the gateway.	false
	INVALID_API_VERSION = 4012
	// You sent an invalid intent for a Gateway Intent. You may have incorrectly calculated the bitwise value.	false
	INVALID_INTENTS = 4013
	// You sent a disallowed intent for a Gateway Intent. You may have tried to specify an intent that you have not enabled or are not approved for.	false
	DISALLOWED_INTENTS = 4014
)

type UserStatus string

const (
	ONLINE         UserStatus = "online"
	DO_NOT_DISTURB            = "dnd"
	IDLE                      = "idle"
	INVISIBLE                 = "invisible"
	OFFLINE                   = "offline"
)

type ActivityType int

const (
	PLAYING   ActivityType = 0
	STREAMING              = 1
	LISTENING              = 2
	WATCHING               = 3
	CUSTOM                 = 4
	COMPETING              = 5
)

type PresenceActivity struct {
	Name string       `json:"name"`
	Type ActivityType `json:"type"`
	// Only valid when Type = ActivityType.STREAMING (1)
	URL string `json:"url,omitempty"`
}

// --- 0 | DISPATCH ---

type GenericDispatchPayload struct {
	Op       Opcode          `json:"op"`
	Data     json.RawMessage `json:"d"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
}

// --- 1 | HEARTBEAT ---

type HeartbeatPayload struct {
	// 1
	Op Opcode `json:"op"`
	// The last sequence received
	Data int64 `json:"d"`
}

// --- 2 | IDENTIFY ---

type IdentifyPayload struct {
	// 2
	Op   Opcode              `json:"op"`
	Data IdentifyPayloadData `json:"d"`
}

type GatewayShard [2]int

type IdentifyPayloadData struct {
	Token          string                        `json:"token"`
	Intents        int                           `json:"intents"`
	Properties     IdentifyPayloadDataProperties `json:"properties"`
	Presence       *IdentifyPayloadDataPresence  `json:"presence,omitempty"`
	Shard          *GatewayShard                 `json:"shard,omitempty"`
	Compress       bool                          `json:"compress,omitempty"`
	LargeThreshold int                           `json:"large_threshold,omitempty"`
}

type IdentifyPayloadDataProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
}

type IdentifyPayloadDataPresence struct {
	Since      int                 `json:"since,omitempty"`
	Activities *[]PresenceActivity `json:"activities,omitempty"`
	Status     UserStatus          `json:"status,omitempty"`
	AFT        bool                `json:"afk,omitempty"`
}

type ResumePayload struct {
	Op   Opcode            `json:"op"`
	Data ResumePayloadData `json:"d"`
}

type ResumePayloadData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int64  `json:"seq"`
}

// --- 10 | HELLO ---

type HelloPayload struct {
	// 10
	Op   Opcode           `json:"op"`
	Data HelloPayloadData `json:"d"`
}

type HelloPayloadData struct {
	HeartbeatInterval int      `json:"heartbeat_interval"`
	Trace             []string `json:"_trace"`
}
