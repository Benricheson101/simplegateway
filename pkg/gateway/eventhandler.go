// DO NOT EDIT
// Generated at 2022-03-02T19:50:00Z

package gateway

var EventNames = []string{
	"GUILD_MEMBER_UPDATE",
	"GUILD_BAN_REMOVE",
	"GUILD_ROLE_UPDATE",
	"GUILD_BAN_ADD",
	"GUILD_MEMBERS_CHUNK",
	"INTEGRATION_CREATE",
	"MESSAGE_REACTION_REMOVE",
	"PRESENCE_UPDATE",
	"STAGE_INSTANCE_CREATE",
	"CHANNEL_UPDATE",
	"GUILD_MEMBER_REMOVE",
	"VOICE_SERVER_UPDATE",
	"THREAD_MEMBER_UPDATE",
	"GUILD_SCHEDULED_EVENT_USER_REMOVE",
	"MESSAGE_UPDATE",
	"MESSAGE_DELETE",
	"MESSAGE_REACTION_REMOVE_EMOJI",
	"THREAD_CREATE",
	"GUILD_STICKERS_UPDATE",
	"STAGE_INSTANCE_UPDATE",
	"CHANNEL_PINS_UPDATE",
	"GUILD_SCHEDULED_EVENT_CREATE",
	"READY",
	"STAGE_INSTANCE_DELETE",
	"WEBHOOKS_UPDATE",
	"GUILD_SCHEDULED_EVENT_USER_ADD",
	"MESSAGE_REACTION_REMOVE_ALL",
	"THREAD_MEMBERS_UPDATE",
	"INTEGRATION_UPDATE",
	"INVITE_CREATE",
	"GUILD_INTEGRATIONS_UPDATE",
	"GUILD_ROLE_DELETE",
	"GUILD_ROLE_CREATE",
	"GUILD_DELETE",
	"TYPING_START",
	"VOICE_STATE_UPDATE",
	"GUILD_UPDATE",
	"CHANNEL_DELETE",
	"THREAD_LIST_SYNC",
	"GUILD_CREATE",
	"GUILD_MEMBER_ADD",
	"INVITE_DELETE",
	"MESSAGE_CREATE",
	"CHANNEL_CREATE",
	"MESSAGE_BULK_DELETE",
	"THREAD_UPDATE",
	"GUILD_EMOJIS_UPDATE",
	"GUILD_SCHEDULED_EVENT_DELETE",
	"INTEGRATION_DELETE",
	"USER_UPDATE",
	"GUILD_SCHEDULED_EVENT_UPDATE",
	"MESSAGE_REACTION_ADD",
	"THREAD_DELETE",
}

func execHandlerFunc(gw *Gateway, handlers []interface{}, t string, payload interface{}) {
	switch t {
	case "GUILD_MEMBER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberUpdate))(gw, payload.(*GuildMemberUpdate))
		}
	case "GUILD_BAN_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildBanRemove))(gw, payload.(*GuildBanRemove))
		}
	case "GUILD_ROLE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleUpdate))(gw, payload.(*GuildRoleUpdate))
		}
	case "GUILD_BAN_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildBanAdd))(gw, payload.(*GuildBanAdd))
		}
	case "GUILD_MEMBERS_CHUNK":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMembersChunk))(gw, payload.(*GuildMembersChunk))
		}
	case "INTEGRATION_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationCreate))(gw, payload.(*IntegrationCreate))
		}
	case "MESSAGE_REACTION_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemove))(gw, payload.(*MessageReactionRemove))
		}
	case "PRESENCE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *PresenceUpdate))(gw, payload.(*PresenceUpdate))
		}
	case "STAGE_INSTANCE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceCreate))(gw, payload.(*StageInstanceCreate))
		}
	case "CHANNEL_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelUpdate))(gw, payload.(*ChannelUpdate))
		}
	case "GUILD_MEMBER_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberRemove))(gw, payload.(*GuildMemberRemove))
		}
	case "VOICE_SERVER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *VoiceServerUpdate))(gw, payload.(*VoiceServerUpdate))
		}
	case "THREAD_MEMBER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadMemberUpdate))(gw, payload.(*ThreadMemberUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUserRemove))(gw, payload.(*GuildScheduledEventUserRemove))
		}
	case "MESSAGE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageUpdate))(gw, payload.(*MessageUpdate))
		}
	case "MESSAGE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageDelete))(gw, payload.(*MessageDelete))
		}
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemoveEmoji))(gw, payload.(*MessageReactionRemoveEmoji))
		}
	case "THREAD_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadCreate))(gw, payload.(*ThreadCreate))
		}
	case "GUILD_STICKERS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildStickersUpdate))(gw, payload.(*GuildStickersUpdate))
		}
	case "STAGE_INSTANCE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceUpdate))(gw, payload.(*StageInstanceUpdate))
		}
	case "CHANNEL_PINS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelPinsUpdate))(gw, payload.(*ChannelPinsUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventCreate))(gw, payload.(*GuildScheduledEventCreate))
		}
	case "READY":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *Ready))(gw, payload.(*Ready))
		}
	case "STAGE_INSTANCE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceDelete))(gw, payload.(*StageInstanceDelete))
		}
	case "WEBHOOKS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *WebhooksUpdate))(gw, payload.(*WebhooksUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_USER_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUserAdd))(gw, payload.(*GuildScheduledEventUserAdd))
		}
	case "MESSAGE_REACTION_REMOVE_ALL":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemoveAll))(gw, payload.(*MessageReactionRemoveAll))
		}
	case "THREAD_MEMBERS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadMembersUpdate))(gw, payload.(*ThreadMembersUpdate))
		}
	case "INTEGRATION_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationUpdate))(gw, payload.(*IntegrationUpdate))
		}
	case "INVITE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *InviteCreate))(gw, payload.(*InviteCreate))
		}
	case "GUILD_INTEGRATIONS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildIntegrationsUpdate))(gw, payload.(*GuildIntegrationsUpdate))
		}
	case "GUILD_ROLE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleDelete))(gw, payload.(*GuildRoleDelete))
		}
	case "GUILD_ROLE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleCreate))(gw, payload.(*GuildRoleCreate))
		}
	case "GUILD_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildDelete))(gw, payload.(*GuildDelete))
		}
	case "TYPING_START":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *TypingStart))(gw, payload.(*TypingStart))
		}
	case "VOICE_STATE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *VoiceStateUpdate))(gw, payload.(*VoiceStateUpdate))
		}
	case "GUILD_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildUpdate))(gw, payload.(*GuildUpdate))
		}
	case "CHANNEL_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelDelete))(gw, payload.(*ChannelDelete))
		}
	case "THREAD_LIST_SYNC":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadListSync))(gw, payload.(*ThreadListSync))
		}
	case "GUILD_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildCreate))(gw, payload.(*GuildCreate))
		}
	case "GUILD_MEMBER_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberAdd))(gw, payload.(*GuildMemberAdd))
		}
	case "INVITE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *InviteDelete))(gw, payload.(*InviteDelete))
		}
	case "MESSAGE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageCreate))(gw, payload.(*MessageCreate))
		}
	case "CHANNEL_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelCreate))(gw, payload.(*ChannelCreate))
		}
	case "MESSAGE_BULK_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageBulkDelete))(gw, payload.(*MessageBulkDelete))
		}
	case "THREAD_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadUpdate))(gw, payload.(*ThreadUpdate))
		}
	case "GUILD_EMOJIS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildEmojisUpdate))(gw, payload.(*GuildEmojisUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventDelete))(gw, payload.(*GuildScheduledEventDelete))
		}
	case "INTEGRATION_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationDelete))(gw, payload.(*IntegrationDelete))
		}
	case "USER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *UserUpdate))(gw, payload.(*UserUpdate))
		}
	case "GUILD_SCHEDULED_EVENT_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUpdate))(gw, payload.(*GuildScheduledEventUpdate))
		}
	case "MESSAGE_REACTION_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionAdd))(gw, payload.(*MessageReactionAdd))
		}
	case "THREAD_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadDelete))(gw, payload.(*ThreadDelete))
		}

	}
}

func eventNameToPayload(t string) interface{} {
	switch t {
	case "GUILD_MEMBER_UPDATE":
		return &GuildMemberUpdate{}
	case "GUILD_BAN_REMOVE":
		return &GuildBanRemove{}
	case "GUILD_ROLE_UPDATE":
		return &GuildRoleUpdate{}
	case "GUILD_BAN_ADD":
		return &GuildBanAdd{}
	case "GUILD_MEMBERS_CHUNK":
		return &GuildMembersChunk{}
	case "INTEGRATION_CREATE":
		return &IntegrationCreate{}
	case "MESSAGE_REACTION_REMOVE":
		return &MessageReactionRemove{}
	case "PRESENCE_UPDATE":
		return &PresenceUpdate{}
	case "STAGE_INSTANCE_CREATE":
		return &StageInstanceCreate{}
	case "CHANNEL_UPDATE":
		return &ChannelUpdate{}
	case "GUILD_MEMBER_REMOVE":
		return &GuildMemberRemove{}
	case "VOICE_SERVER_UPDATE":
		return &VoiceServerUpdate{}
	case "THREAD_MEMBER_UPDATE":
		return &ThreadMemberUpdate{}
	case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
		return &GuildScheduledEventUserRemove{}
	case "MESSAGE_UPDATE":
		return &MessageUpdate{}
	case "MESSAGE_DELETE":
		return &MessageDelete{}
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		return &MessageReactionRemoveEmoji{}
	case "THREAD_CREATE":
		return &ThreadCreate{}
	case "GUILD_STICKERS_UPDATE":
		return &GuildStickersUpdate{}
	case "STAGE_INSTANCE_UPDATE":
		return &StageInstanceUpdate{}
	case "CHANNEL_PINS_UPDATE":
		return &ChannelPinsUpdate{}
	case "GUILD_SCHEDULED_EVENT_CREATE":
		return &GuildScheduledEventCreate{}
	case "READY":
		return &Ready{}
	case "STAGE_INSTANCE_DELETE":
		return &StageInstanceDelete{}
	case "WEBHOOKS_UPDATE":
		return &WebhooksUpdate{}
	case "GUILD_SCHEDULED_EVENT_USER_ADD":
		return &GuildScheduledEventUserAdd{}
	case "MESSAGE_REACTION_REMOVE_ALL":
		return &MessageReactionRemoveAll{}
	case "THREAD_MEMBERS_UPDATE":
		return &ThreadMembersUpdate{}
	case "INTEGRATION_UPDATE":
		return &IntegrationUpdate{}
	case "INVITE_CREATE":
		return &InviteCreate{}
	case "GUILD_INTEGRATIONS_UPDATE":
		return &GuildIntegrationsUpdate{}
	case "GUILD_ROLE_DELETE":
		return &GuildRoleDelete{}
	case "GUILD_ROLE_CREATE":
		return &GuildRoleCreate{}
	case "GUILD_DELETE":
		return &GuildDelete{}
	case "TYPING_START":
		return &TypingStart{}
	case "VOICE_STATE_UPDATE":
		return &VoiceStateUpdate{}
	case "GUILD_UPDATE":
		return &GuildUpdate{}
	case "CHANNEL_DELETE":
		return &ChannelDelete{}
	case "THREAD_LIST_SYNC":
		return &ThreadListSync{}
	case "GUILD_CREATE":
		return &GuildCreate{}
	case "GUILD_MEMBER_ADD":
		return &GuildMemberAdd{}
	case "INVITE_DELETE":
		return &InviteDelete{}
	case "MESSAGE_CREATE":
		return &MessageCreate{}
	case "CHANNEL_CREATE":
		return &ChannelCreate{}
	case "MESSAGE_BULK_DELETE":
		return &MessageBulkDelete{}
	case "THREAD_UPDATE":
		return &ThreadUpdate{}
	case "GUILD_EMOJIS_UPDATE":
		return &GuildEmojisUpdate{}
	case "GUILD_SCHEDULED_EVENT_DELETE":
		return &GuildScheduledEventDelete{}
	case "INTEGRATION_DELETE":
		return &IntegrationDelete{}
	case "USER_UPDATE":
		return &UserUpdate{}
	case "GUILD_SCHEDULED_EVENT_UPDATE":
		return &GuildScheduledEventUpdate{}
	case "MESSAGE_REACTION_ADD":
		return &MessageReactionAdd{}
	case "THREAD_DELETE":
		return &ThreadDelete{}
	default:
		return nil
	}
}
