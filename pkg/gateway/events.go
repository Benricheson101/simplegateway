package gateway

import (
	"bytes"
	"reflect"
)

// will the workflow work?

// The ready event is dispatched when a client has completed the initial handshake with the gateway (for new sessions). The ready event can be the largest and most complex event the gateway will send, as it contains all the state required for a client to begin interacting with the rest of the platform
type Ready struct {
	// gateway version
	Version int `json:"v"`
	// information about the user including email
	User User `json:"user"`
	// the guilds the user is in
	Guilds []UnavailableGuild `json:"guilds"`
	// used for resuming connections
	SessionID string `json:"session_id"`
	// the shard information associated with this session, if sent when identifying
	Shard *GatewayShard `json:"shard,omitempty"`
	// contains id and flags
	Application Application `json:"application"`
}

// Sent when a new guild channel is created, relevant to the current user
type ChannelCreate struct {
	Channel
}

// Sent when a channel is updated
type ChannelUpdate struct {
	Channel
}

// Sent when a channel is deleted
type ChannelDelete struct {
	Channel
}

// Sent when a message is pinned or unpinned in a text channel. This is not sent when a pinned message is deleted
type ChannelPinsUpdate struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the time at which the most recent pinned message was pinned
	LastPinTimestamp string `json:"last_pin_timestamp"`
}

// Sent when a thread is created, relevant to the current user, or when the current user is added to a thread
type ThreadCreate struct {
	Channel
	NewlyCreated bool `json:"newly_created"`
}

// Sent when a thread is updated
type ThreadUpdate struct {
	Channel
}

// Sent when a thread relevant to the current user is deleted
type ThreadDelete struct {
	// the id of this channel
	ID Snowflake `json:"id"`
	// the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	GuildID Snowflake `json:"guild_id"`
	// for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	ParentID Snowflake `json:"parent_id"`
	// the type of channel
	Type ChannelType `json:"type"`
}

// Sent when the current user gains access to a channel
type ThreadListSync struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the parent channel ids whose threads are being synced. If omitted, then threads were synced for the entire guild. This array may contain channel_ids that have no active threads as well, so you know to clear that data.
	ChannelIDs []Snowflake `json:"channel_ids"`
	// all active threads in the given channels that the current user can access
	Threads []Channel `json:"threads"`
	// all thread member objects from the synced threads for the current user, indicating which threads the current user has been added to
	Members []ThreadMember `json:"members"`
}

// Sent when the thread member object for the current user is updated
type ThreadMemberUpdate struct {
	ThreadMember
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when anyone is added to or removed from a thread

type ThreadMembersUpdate struct {
	// the id of the thread
	ID Snowflake `json:"id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the approximate number of members in the thread, capped at 50
	MemberCount int `json:"member_count"`
	// the users who were added to the thread
	AddedMembers []ThreadMember `json:"added_members*"`
	// the id of the users who were removed from the thread
	RemovedMemberIDs []Snowflake `json:"removed_member_ids"`
}

// This event can be sent in three different scenarios:
// 1. When a user is initially connecting, to lazily load and backfill information for all unavailable guilds sent in the Ready event. Guilds that are unavailable due to an outage will send a Guild Delete event.
// 2. When a Guild becomes available again to the client.
// 3. When the current user joins a new Guild.
type GuildCreate struct {
	Guild
}

// Sent when a guild is updated
type GuildUpdate struct {
	Guild
}

// Sent when a guild becomes or was already unavailable due to an outage, or when the user leaves or is removed from a guild
type GuildDelete struct {
	UnavailableGuild
}

// Sent when a user is banned from a guild
type GuildBanAdd struct {
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the banned user
	User User `json:"user"`
}

// Sent when a user is unbanned from a guild
type GuildBanRemove struct {
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the unbanned user
	User User `json:"user"`
}

// Sent when a guild's emojis have been updated
type GuildEmojisUpdate struct {
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// array of emojis
	Emojis []Emoji `json:" emojis"`
}

// Sent when a guild's stickers have been updated
type GuildStickersUpdate struct {
	// id of the guild
	GuildID Snowflake `json:" guild_id"`
	// array of stickers
	Stickers []Sticker `json:"stickers"`
}

// Sent when a guild integration is updated
type GuildIntegrationsUpdate struct {
	// id of the guild whose integrations were updated
	GuildID Snowflake `json:"guild_id"`
}

// Sent when a new user joins a guild
type GuildMemberAdd struct {
	GuildMember
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when a user is removed from a guild (leave/kick/ban)
type GuildMemberRemove struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the user who was removed
	User User `json:"user"`
}

// Sent when a guild member is updated
type GuildMemberUpdate struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// user role ids
	Roles []Snowflake `json:"roles"`
	// the user
	User User `json:"user"`
	// nickname of the user in the guild
	Nick string `json:"nick"`
	// the member's guild avatar hash
	Avatar string `json:"avatar"`
	// when the user joined the guild
	JoinedAt string `json:"joined_at"`
	// when the user starting boosting the guild
	PremiumSince string `json:"premium_since"`
	// whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`
	// whether the user is muted in voice channels
	Mute bool `json:"mute"`
	// whether the user has not yet passed the guild's Membership Screening requirements
	Pending bool `json:"pending"`
	// when the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
	CommunicationDisabledUntil string `json:"communication_disabled_until"`
}

// Sent in response to Guild Request Members
type GuildMembersChunk struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// set of guild members
	Members []GuildMember `json:"members"`
	// the chunk index in the expected chunks for this response (0 <= chunk_index < chunk_count)
	ChunkIndex int `json:"chunk_index"`
	// the total number of expected chunks for this response
	ChunkCount int `json:"chunk_count"`
	// if passing an invalid id to REQUEST_GUILD_MEMBERS, it will be returned here
	NotFound []Snowflake `json:"not_found"`
	// if passing true to REQUEST_GUILD_MEMBERS, presences of the returned members will be here
	Presences []Presence `json:"presences"`
	// the nonce used in the Guild Members Request
	Nonce string `json:"nonce"`
}

// Sent when a guild role is created
type GuildRoleCreate struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the role created
	Role Role `json:"role"`
}

// Sent when a guild role is updated
type GuildRoleUpdate struct {
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
	// the role updated
	Role Role `json:"role"`
}

// Sent when a guild role is deleted
type GuildRoleDelete struct {
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// id of the role
	RoleID Snowflake `json:"role_id"`
}

// Sent when a guild scheduled event is created
type GuildScheduledEventCreate struct {
	GuildScheduledEvent
}

// Sent when a guild scheduled event is updated
type GuildScheduledEventUpdate struct {
	GuildScheduledEvent
}

// Sent when a guild scheduled event is deleted
type GuildScheduledEventDelete struct {
	GuildScheduledEvent
}

// Sent when a user has subscribed to a guild scheduled event
type GuildScheduledEventUserAdd struct {
	// id of the guild scheduled event
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	// id of the user
	UserID Snowflake `json:"user_id"`
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when a user has unsubscribed from a guild scheduled event
type GuildScheduledEventUserRemove struct {
	// id of the guild scheduled event
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	// id of the user
	UserID Snowflake `json:"user_id"`
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when an integration is created
type IntegrationCreate struct {
	Integration
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when an integration is updated
type IntegrationUpdate struct {
	Integration
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when an integration is deleted
type IntegrationDelete struct {
	// integration id
	ID Snowflake `json:"id"`
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// id of the bot/OAuth2 application for this discord integration
	ApplicationID Snowflake `json:"application_id"`
}

// Sent when a new invite to a channel is created
type InviteCreate struct {
	// the channel the invite is for
	ChannelID Snowflake `json:"channel_id"`
	// the unique invite code
	Code string `json:"code"`
	// the time at which the invite was created
	CreatedAt string `json:"created_at"`
	// the guild of the invite
	GuildID Snowflake `json:"guild_id"`
	// the user that created the invite
	Inviter User `json:"inviter"`
	// how long the invite is valid for (in seconds)
	MaxAge int `json:"max_age"`
	// the maximum number of times the invite can be used
	MaxUses int `json:"max_uses"`
	// the type of target for this voice channel invite
	TargetType InviteTargetType `json:"target_type"`
	// the user whose stream to display for this voice channel stream invite
	TargetUser User `json:"target_user"`
	// the embedded application to open for this voice channel embedded application invite
	TargetApplication Application `json:"target_application"`
	// whether or not the invite is temporary (invited users will be kicked on disconnect unless they're assigned a role)
	Temporary bool `json:"temporary"`
	// how many times the invite has been used (always will be 0)
	Uses int `json:"uses"`
}

// Sent when an invite is deleted
type InviteDelete struct {
	// the channel of the invite
	ChannelID Snowflake `json:"channel_id"`
	// the guild of the invite
	GuildID Snowflake `json:"guild_id"`
	// the unique invite code
	Code string `json:"code"`
}

// Sent when a message is created
type MessageCreate struct {
	Message
}

// Sent when a message is updated
type MessageUpdate struct {
	Message
}

// Sent when a message is deleted
type MessageDelete struct {
	// the id of the message
	ID Snowflake `json:"id"`
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id"`
}

// Sent when multiple messages are deleted at once
type MessageBulkDelete struct {
	// the ids of the messages
	IDs []Snowflake `json:"ids"`
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// Sent when a user adds a reaction to a message
type MessageReactionAdd struct {
	// the id of the user
	UserID Snowflake `json:"user_id"`
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the id of the message
	MessageID Snowflake `json:"message_id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id,omitempty"`
	// the member who reacted if this happened in a guild
	Member GuildMember `json:"member,omitempty"`
	// the emoji used to react - example
	Emoji Emoji `json:"emoji"`
}

// Sent when a user removes a reaction from a message
type MessageReactionRemove struct {
	// the id of the user
	UserID Snowflake `json:"user_id"`
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the id of the message
	MessageID Snowflake `json:"message_id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id,omitempty"`
	// the emoji used to react
	Emoji Emoji `json:"emoji"`
}

// Sent when a user explicitly removes all reactions from a message
type MessageReactionRemoveAll struct {
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the id of the message
	MessageID Snowflake `json:"message_id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// Sent when a bot removes all instances of a given emoji from the reactions of a message
type MessageReactionRemoveEmoji struct {
	// the id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// the id of the guild
	GuildID Snowflake `json:"guild_id,omitempty"`
	// the id of the message
	MessageID Snowflake `json:"message_id"`
	// the emoji that was removed
	Emoji Emoji `json:"emoji"`
}

// Sent when a user's presence or info, such as name or avatar, is updated
type PresenceUpdate struct {
	Presence
}

// Sent when a user starts typing in a channel
type TypingStart struct {
	// id of the channel
	ChannelID Snowflake `json:"channel_id"`
	// id of the guild
	GuildID Snowflake `json:"guild_id,omitempty"`
	// id of the user
	UserID Snowflake `json:"user_id"`
	// unix time (in seconds) of when the user started typing
	Timestamp int `json:"timestamp"`
	// the member who started typing if this happened in a guild
	Member *GuildMember `json:"member,omitempty"`
}

// Sent when properties about the user change
type UserUpdate struct {
	User
}

// Sent when someone joins/leaves/moves voice channels
type VoiceStateUpdate struct {
	VoiceState
}

// Sent when a guild's voice server is updated. This is sent when initially connecting to voice, and when the current voice instance fails over to a new server
type VoiceServerUpdate struct {
	// voice connection token
	Token string `json:"token"`
	// the guild this voice server update is for
	GuildID Snowflake `json:"guild_id"`
	// the voice server host
	Endpoint *string `json:"endpoint"`
}

// Sent when a guild channel's webhook is created, updated, or deleted
type WebhooksUpdate struct {
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// id of the channel
	ChannelID Snowflake `json:"channel_id"`
}

// TODO: type InteractionCreate struct{}

// Sent when a Stage instance is created (i.e. the Stage is now "live")
type StageInstanceCreate struct {
	StageInstance
}

// Sent when a Stage instance has been updated
type StageInstanceUpdate struct {
	StageInstance
}

// Sent when a Stage instance has been deleted (i.e. the Stage has been closed)
type StageInstanceDelete struct {
	StageInstance
}

func eventHandlerToEventName(fn interface{}) string {
	typ := reflect.TypeOf(fn)
	if typ.Kind() != reflect.Func {
		panic("type is not func")
	}

	ins := typ.NumIn()
	if ins != 2 {
		panic("wrong number of parameters")
	}

	snd := typ.In(1)
	if snd.Kind() != reflect.Ptr {
		panic("second param must be a ptr")
	}

	sndInd := reflect.Indirect(reflect.ValueOf(snd))
	if sndInd.Kind() != reflect.Struct {
		panic("second param must be of type `struct`")
	}

	return camelToScreamingSnake(snd.Elem().Name())
}

func camelToScreamingSnake(s string) string {
	var buf bytes.Buffer

	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c)
		} else {
			buf.WriteRune(c + 'A' - 'a')
		}
	}

	return buf.String()
}
