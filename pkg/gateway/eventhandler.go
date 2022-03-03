// DO NOT EDIT
// Generated at 2022-03-03T04:58:13Z

package gateway

var EventNames = []string{
	"INTEGRATION_DELETE",
	"MESSAGE_REACTION_REMOVE",
	"MESSAGE_REACTION_REMOVE_ALL",
	"PRESENCE_UPDATE",
	"STAGE_INSTANCE_UPDATE",
	"CHANNEL_UPDATE",
	"CHANNEL_DELETE",
	"GUILD_ROLE_UPDATE",
	"MESSAGE_BULK_DELETE",
	"USER_UPDATE",
	"VOICE_STATE_UPDATE",
	"THREAD_CREATE",
	"GUILD_ROLE_CREATE",
	"TYPING_START",
	"GUILD_SCHEDULED_EVENT_USER_ADD",
	"CHANNEL_PINS_UPDATE",
	"THREAD_DELETE",
	"GUILD_STICKERS_UPDATE",
	"INTEGRATION_CREATE",
	"THREAD_UPDATE",
	"GUILD_SCHEDULED_EVENT_DELETE",
	"MESSAGE_CREATE",
	"GUILD_MEMBER_ADD",
	"THREAD_MEMBERS_UPDATE",
	"MESSAGE_DELETE",
	"CHANNEL_CREATE",
	"GUILD_BAN_REMOVE",
	"MESSAGE_UPDATE",
	"THREAD_MEMBER_UPDATE",
	"MESSAGE_REACTION_ADD",
	"VOICE_SERVER_UPDATE",
	"STAGE_INSTANCE_DELETE",
	"GUILD_ROLE_DELETE",
	"MESSAGE_REACTION_REMOVE_EMOJI",
	"GUILD_MEMBER_REMOVE",
	"READY",
	"RESUMED",
	"GUILD_BAN_ADD",
	"INVITE_DELETE",
	"GUILD_DELETE",
	"GUILD_SCHEDULED_EVENT_UPDATE",
	"INVITE_CREATE",
	"STAGE_INSTANCE_CREATE",
	"GUILD_MEMBER_UPDATE",
	"GUILD_MEMBERS_CHUNK",
	"INTEGRATION_UPDATE",
	"THREAD_LIST_SYNC",
	"GUILD_CREATE",
	"GUILD_UPDATE",
	"GUILD_EMOJIS_UPDATE",
	"GUILD_INTEGRATIONS_UPDATE",
	"GUILD_SCHEDULED_EVENT_CREATE",
	"GUILD_SCHEDULED_EVENT_USER_REMOVE",
	"WEBHOOKS_UPDATE",
}

func execHandlerFunc(gw *Gateway, handlers []interface{}, t string, payload interface{}) {
	switch t {
	case "INTEGRATION_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationDelete))(gw, payload.(*IntegrationDelete))
		}
	case "MESSAGE_REACTION_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemove))(gw, payload.(*MessageReactionRemove))
		}
	case "MESSAGE_REACTION_REMOVE_ALL":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemoveAll))(gw, payload.(*MessageReactionRemoveAll))
		}
	case "PRESENCE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *PresenceUpdate))(gw, payload.(*PresenceUpdate))
		}
	case "STAGE_INSTANCE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceUpdate))(gw, payload.(*StageInstanceUpdate))
		}
	case "CHANNEL_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelUpdate))(gw, payload.(*ChannelUpdate))
		}
	case "CHANNEL_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelDelete))(gw, payload.(*ChannelDelete))
		}
	case "GUILD_ROLE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleUpdate))(gw, payload.(*GuildRoleUpdate))
		}
	case "MESSAGE_BULK_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageBulkDelete))(gw, payload.(*MessageBulkDelete))
		}
	case "USER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *UserUpdate))(gw, payload.(*UserUpdate))
		}
	case "VOICE_STATE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *VoiceStateUpdate))(gw, payload.(*VoiceStateUpdate))
		}
	case "THREAD_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadCreate))(gw, payload.(*ThreadCreate))
		}
	case "GUILD_ROLE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleCreate))(gw, payload.(*GuildRoleCreate))
		}
	case "TYPING_START":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *TypingStart))(gw, payload.(*TypingStart))
		}
	case "GUILD_SCHEDULED_EVENT_USER_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUserAdd))(gw, payload.(*GuildScheduledEventUserAdd))
		}
	case "CHANNEL_PINS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelPinsUpdate))(gw, payload.(*ChannelPinsUpdate))
		}
	case "THREAD_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadDelete))(gw, payload.(*ThreadDelete))
		}
	case "GUILD_STICKERS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildStickersUpdate))(gw, payload.(*GuildStickersUpdate))
		}
	case "INTEGRATION_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationCreate))(gw, payload.(*IntegrationCreate))
		}
	case "THREAD_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadUpdate))(gw, payload.(*ThreadUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventDelete))(gw, payload.(*GuildScheduledEventDelete))
		}
	case "MESSAGE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageCreate))(gw, payload.(*MessageCreate))
		}
	case "GUILD_MEMBER_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberAdd))(gw, payload.(*GuildMemberAdd))
		}
	case "THREAD_MEMBERS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadMembersUpdate))(gw, payload.(*ThreadMembersUpdate))
		}
	case "MESSAGE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageDelete))(gw, payload.(*MessageDelete))
		}
	case "CHANNEL_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelCreate))(gw, payload.(*ChannelCreate))
		}
	case "GUILD_BAN_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildBanRemove))(gw, payload.(*GuildBanRemove))
		}
	case "MESSAGE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageUpdate))(gw, payload.(*MessageUpdate))
		}
	case "THREAD_MEMBER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadMemberUpdate))(gw, payload.(*ThreadMemberUpdate))
		}
	case "MESSAGE_REACTION_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionAdd))(gw, payload.(*MessageReactionAdd))
		}
	case "VOICE_SERVER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *VoiceServerUpdate))(gw, payload.(*VoiceServerUpdate))
		}
	case "STAGE_INSTANCE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceDelete))(gw, payload.(*StageInstanceDelete))
		}
	case "GUILD_ROLE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleDelete))(gw, payload.(*GuildRoleDelete))
		}
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemoveEmoji))(gw, payload.(*MessageReactionRemoveEmoji))
		}
	case "GUILD_MEMBER_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberRemove))(gw, payload.(*GuildMemberRemove))
		}
	case "READY":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *Ready))(gw, payload.(*Ready))
		}
	case "RESUMED":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *Resumed))(gw, payload.(*Resumed))
		}
	case "GUILD_BAN_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildBanAdd))(gw, payload.(*GuildBanAdd))
		}
	case "INVITE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *InviteDelete))(gw, payload.(*InviteDelete))
		}
	case "GUILD_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildDelete))(gw, payload.(*GuildDelete))
		}
	case "GUILD_SCHEDULED_EVENT_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUpdate))(gw, payload.(*GuildScheduledEventUpdate))
		}
	case "INVITE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *InviteCreate))(gw, payload.(*InviteCreate))
		}
	case "STAGE_INSTANCE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceCreate))(gw, payload.(*StageInstanceCreate))
		}
	case "GUILD_MEMBER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberUpdate))(gw, payload.(*GuildMemberUpdate))
		}
	case "GUILD_MEMBERS_CHUNK":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMembersChunk))(gw, payload.(*GuildMembersChunk))
		}
	case "INTEGRATION_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationUpdate))(gw, payload.(*IntegrationUpdate))
		}
	case "THREAD_LIST_SYNC":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadListSync))(gw, payload.(*ThreadListSync))
		}
	case "GUILD_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildCreate))(gw, payload.(*GuildCreate))
		}
	case "GUILD_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildUpdate))(gw, payload.(*GuildUpdate))
		}
	case "GUILD_EMOJIS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildEmojisUpdate))(gw, payload.(*GuildEmojisUpdate))
		}
	case "GUILD_INTEGRATIONS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildIntegrationsUpdate))(gw, payload.(*GuildIntegrationsUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventCreate))(gw, payload.(*GuildScheduledEventCreate))
		}
	case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUserRemove))(gw, payload.(*GuildScheduledEventUserRemove))
		}
	case "WEBHOOKS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *WebhooksUpdate))(gw, payload.(*WebhooksUpdate))
		}

	}
}

func eventNameToPayload(t string) interface{} {
	switch t {
	case "INTEGRATION_DELETE":
		return &IntegrationDelete{}
	case "MESSAGE_REACTION_REMOVE":
		return &MessageReactionRemove{}
	case "MESSAGE_REACTION_REMOVE_ALL":
		return &MessageReactionRemoveAll{}
	case "PRESENCE_UPDATE":
		return &PresenceUpdate{}
	case "STAGE_INSTANCE_UPDATE":
		return &StageInstanceUpdate{}
	case "CHANNEL_UPDATE":
		return &ChannelUpdate{}
	case "CHANNEL_DELETE":
		return &ChannelDelete{}
	case "GUILD_ROLE_UPDATE":
		return &GuildRoleUpdate{}
	case "MESSAGE_BULK_DELETE":
		return &MessageBulkDelete{}
	case "USER_UPDATE":
		return &UserUpdate{}
	case "VOICE_STATE_UPDATE":
		return &VoiceStateUpdate{}
	case "THREAD_CREATE":
		return &ThreadCreate{}
	case "GUILD_ROLE_CREATE":
		return &GuildRoleCreate{}
	case "TYPING_START":
		return &TypingStart{}
	case "GUILD_SCHEDULED_EVENT_USER_ADD":
		return &GuildScheduledEventUserAdd{}
	case "CHANNEL_PINS_UPDATE":
		return &ChannelPinsUpdate{}
	case "THREAD_DELETE":
		return &ThreadDelete{}
	case "GUILD_STICKERS_UPDATE":
		return &GuildStickersUpdate{}
	case "INTEGRATION_CREATE":
		return &IntegrationCreate{}
	case "THREAD_UPDATE":
		return &ThreadUpdate{}
	case "GUILD_SCHEDULED_EVENT_DELETE":
		return &GuildScheduledEventDelete{}
	case "MESSAGE_CREATE":
		return &MessageCreate{}
	case "GUILD_MEMBER_ADD":
		return &GuildMemberAdd{}
	case "THREAD_MEMBERS_UPDATE":
		return &ThreadMembersUpdate{}
	case "MESSAGE_DELETE":
		return &MessageDelete{}
	case "CHANNEL_CREATE":
		return &ChannelCreate{}
	case "GUILD_BAN_REMOVE":
		return &GuildBanRemove{}
	case "MESSAGE_UPDATE":
		return &MessageUpdate{}
	case "THREAD_MEMBER_UPDATE":
		return &ThreadMemberUpdate{}
	case "MESSAGE_REACTION_ADD":
		return &MessageReactionAdd{}
	case "VOICE_SERVER_UPDATE":
		return &VoiceServerUpdate{}
	case "STAGE_INSTANCE_DELETE":
		return &StageInstanceDelete{}
	case "GUILD_ROLE_DELETE":
		return &GuildRoleDelete{}
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		return &MessageReactionRemoveEmoji{}
	case "GUILD_MEMBER_REMOVE":
		return &GuildMemberRemove{}
	case "READY":
		return &Ready{}
	case "RESUMED":
		return &Resumed{}
	case "GUILD_BAN_ADD":
		return &GuildBanAdd{}
	case "INVITE_DELETE":
		return &InviteDelete{}
	case "GUILD_DELETE":
		return &GuildDelete{}
	case "GUILD_SCHEDULED_EVENT_UPDATE":
		return &GuildScheduledEventUpdate{}
	case "INVITE_CREATE":
		return &InviteCreate{}
	case "STAGE_INSTANCE_CREATE":
		return &StageInstanceCreate{}
	case "GUILD_MEMBER_UPDATE":
		return &GuildMemberUpdate{}
	case "GUILD_MEMBERS_CHUNK":
		return &GuildMembersChunk{}
	case "INTEGRATION_UPDATE":
		return &IntegrationUpdate{}
	case "THREAD_LIST_SYNC":
		return &ThreadListSync{}
	case "GUILD_CREATE":
		return &GuildCreate{}
	case "GUILD_UPDATE":
		return &GuildUpdate{}
	case "GUILD_EMOJIS_UPDATE":
		return &GuildEmojisUpdate{}
	case "GUILD_INTEGRATIONS_UPDATE":
		return &GuildIntegrationsUpdate{}
	case "GUILD_SCHEDULED_EVENT_CREATE":
		return &GuildScheduledEventCreate{}
	case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
		return &GuildScheduledEventUserRemove{}
	case "WEBHOOKS_UPDATE":
		return &WebhooksUpdate{}
	default:
		return nil
	}
}
