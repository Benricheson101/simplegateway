// DO NOT EDIT
// Generated at 2022-03-02T19:23:43Z

package gateway

var EventNames = []string{
	"GUILD_SCHEDULED_EVENT_DELETE",
	"INTEGRATION_CREATE",
	"READY",
	"GUILD_SCHEDULED_EVENT_CREATE",
	"GUILD_SCHEDULED_EVENT_UPDATE",
	"GUILD_ROLE_DELETE",
	"MESSAGE_UPDATE",
	"MESSAGE_REACTION_REMOVE_EMOJI",
	"GUILD_UPDATE",
	"VOICE_STATE_UPDATE",
	"THREAD_UPDATE",
	"GUILD_DELETE",
	"INVITE_CREATE",
	"CHANNEL_UPDATE",
	"THREAD_DELETE",
	"MESSAGE_REACTION_REMOVE_ALL",
	"PRESENCE_UPDATE",
	"STAGE_INSTANCE_CREATE",
	"THREAD_LIST_SYNC",
	"GUILD_ROLE_UPDATE",
	"MESSAGE_DELETE",
	"GUILD_SCHEDULED_EVENT_USER_ADD",
	"GUILD_MEMBER_REMOVE",
	"GUILD_BAN_REMOVE",
	"CHANNEL_DELETE",
	"THREAD_MEMBER_UPDATE",
	"THREAD_MEMBERS_UPDATE",
	"USER_UPDATE",
	"GUILD_STICKERS_UPDATE",
	"INVITE_DELETE",
	"CHANNEL_PINS_UPDATE",
	"THREAD_CREATE",
	"MESSAGE_BULK_DELETE",
	"MESSAGE_CREATE",
	"MESSAGE_REACTION_REMOVE",
	"GUILD_MEMBERS_CHUNK",
	"GUILD_SCHEDULED_EVENT_USER_REMOVE",
	"INTEGRATION_UPDATE",
	"WEBHOOKS_UPDATE",
	"GUILD_ROLE_CREATE",
	"INTEGRATION_DELETE",
	"TYPING_START",
	"GUILD_CREATE",
	"MESSAGE_REACTION_ADD",
	"STAGE_INSTANCE_UPDATE",
	"GUILD_INTEGRATIONS_UPDATE",
	"GUILD_MEMBER_ADD",
	"VOICE_SERVER_UPDATE",
	"STAGE_INSTANCE_DELETE",
	"CHANNEL_CREATE",
	"GUILD_BAN_ADD",
	"GUILD_EMOJIS_UPDATE",
	"GUILD_MEMBER_UPDATE",
}

func execHandlerFunc(gw *Gateway, handlers []interface{}, t string, payload interface{}) {
	switch t {
	case "GUILD_SCHEDULED_EVENT_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventDelete))(gw, payload.(*GuildScheduledEventDelete))
		}
	case "INTEGRATION_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationCreate))(gw, payload.(*IntegrationCreate))
		}
	case "READY":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *Ready))(gw, payload.(*Ready))
		}
	case "GUILD_SCHEDULED_EVENT_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventCreate))(gw, payload.(*GuildScheduledEventCreate))
		}
	case "GUILD_SCHEDULED_EVENT_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUpdate))(gw, payload.(*GuildScheduledEventUpdate))
		}
	case "GUILD_ROLE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleDelete))(gw, payload.(*GuildRoleDelete))
		}
	case "MESSAGE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageUpdate))(gw, payload.(*MessageUpdate))
		}
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemoveEmoji))(gw, payload.(*MessageReactionRemoveEmoji))
		}
	case "GUILD_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildUpdate))(gw, payload.(*GuildUpdate))
		}
	case "VOICE_STATE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *VoiceStateUpdate))(gw, payload.(*VoiceStateUpdate))
		}
	case "THREAD_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadUpdate))(gw, payload.(*ThreadUpdate))
		}
	case "GUILD_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildDelete))(gw, payload.(*GuildDelete))
		}
	case "INVITE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *InviteCreate))(gw, payload.(*InviteCreate))
		}
	case "CHANNEL_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelUpdate))(gw, payload.(*ChannelUpdate))
		}
	case "THREAD_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadDelete))(gw, payload.(*ThreadDelete))
		}
	case "MESSAGE_REACTION_REMOVE_ALL":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemoveAll))(gw, payload.(*MessageReactionRemoveAll))
		}
	case "PRESENCE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *PresenceUpdate))(gw, payload.(*PresenceUpdate))
		}
	case "STAGE_INSTANCE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceCreate))(gw, payload.(*StageInstanceCreate))
		}
	case "THREAD_LIST_SYNC":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadListSync))(gw, payload.(*ThreadListSync))
		}
	case "GUILD_ROLE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleUpdate))(gw, payload.(*GuildRoleUpdate))
		}
	case "MESSAGE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageDelete))(gw, payload.(*MessageDelete))
		}
	case "GUILD_SCHEDULED_EVENT_USER_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUserAdd))(gw, payload.(*GuildScheduledEventUserAdd))
		}
	case "GUILD_MEMBER_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberRemove))(gw, payload.(*GuildMemberRemove))
		}
	case "GUILD_BAN_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildBanRemove))(gw, payload.(*GuildBanRemove))
		}
	case "CHANNEL_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelDelete))(gw, payload.(*ChannelDelete))
		}
	case "THREAD_MEMBER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadMemberUpdate))(gw, payload.(*ThreadMemberUpdate))
		}
	case "THREAD_MEMBERS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadMembersUpdate))(gw, payload.(*ThreadMembersUpdate))
		}
	case "USER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *UserUpdate))(gw, payload.(*UserUpdate))
		}
	case "GUILD_STICKERS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildStickersUpdate))(gw, payload.(*GuildStickersUpdate))
		}
	case "INVITE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *InviteDelete))(gw, payload.(*InviteDelete))
		}
	case "CHANNEL_PINS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelPinsUpdate))(gw, payload.(*ChannelPinsUpdate))
		}
	case "THREAD_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ThreadCreate))(gw, payload.(*ThreadCreate))
		}
	case "MESSAGE_BULK_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageBulkDelete))(gw, payload.(*MessageBulkDelete))
		}
	case "MESSAGE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageCreate))(gw, payload.(*MessageCreate))
		}
	case "MESSAGE_REACTION_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionRemove))(gw, payload.(*MessageReactionRemove))
		}
	case "GUILD_MEMBERS_CHUNK":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMembersChunk))(gw, payload.(*GuildMembersChunk))
		}
	case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildScheduledEventUserRemove))(gw, payload.(*GuildScheduledEventUserRemove))
		}
	case "INTEGRATION_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationUpdate))(gw, payload.(*IntegrationUpdate))
		}
	case "WEBHOOKS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *WebhooksUpdate))(gw, payload.(*WebhooksUpdate))
		}
	case "GUILD_ROLE_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildRoleCreate))(gw, payload.(*GuildRoleCreate))
		}
	case "INTEGRATION_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *IntegrationDelete))(gw, payload.(*IntegrationDelete))
		}
	case "TYPING_START":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *TypingStart))(gw, payload.(*TypingStart))
		}
	case "GUILD_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildCreate))(gw, payload.(*GuildCreate))
		}
	case "MESSAGE_REACTION_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *MessageReactionAdd))(gw, payload.(*MessageReactionAdd))
		}
	case "STAGE_INSTANCE_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceUpdate))(gw, payload.(*StageInstanceUpdate))
		}
	case "GUILD_INTEGRATIONS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildIntegrationsUpdate))(gw, payload.(*GuildIntegrationsUpdate))
		}
	case "GUILD_MEMBER_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberAdd))(gw, payload.(*GuildMemberAdd))
		}
	case "VOICE_SERVER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *VoiceServerUpdate))(gw, payload.(*VoiceServerUpdate))
		}
	case "STAGE_INSTANCE_DELETE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *StageInstanceDelete))(gw, payload.(*StageInstanceDelete))
		}
	case "CHANNEL_CREATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *ChannelCreate))(gw, payload.(*ChannelCreate))
		}
	case "GUILD_BAN_ADD":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildBanAdd))(gw, payload.(*GuildBanAdd))
		}
	case "GUILD_EMOJIS_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildEmojisUpdate))(gw, payload.(*GuildEmojisUpdate))
		}
	case "GUILD_MEMBER_UPDATE":
		for _, fn := range handlers {
			go fn.(func(*Gateway, *GuildMemberUpdate))(gw, payload.(*GuildMemberUpdate))
		}

	}
}

func eventNameToPayload(t string) interface{} {
	switch t {
	case "GUILD_SCHEDULED_EVENT_DELETE":
		return &GuildScheduledEventDelete{}
	case "INTEGRATION_CREATE":
		return &IntegrationCreate{}
	case "READY":
		return &Ready{}
	case "GUILD_SCHEDULED_EVENT_CREATE":
		return &GuildScheduledEventCreate{}
	case "GUILD_SCHEDULED_EVENT_UPDATE":
		return &GuildScheduledEventUpdate{}
	case "GUILD_ROLE_DELETE":
		return &GuildRoleDelete{}
	case "MESSAGE_UPDATE":
		return &MessageUpdate{}
	case "MESSAGE_REACTION_REMOVE_EMOJI":
		return &MessageReactionRemoveEmoji{}
	case "GUILD_UPDATE":
		return &GuildUpdate{}
	case "VOICE_STATE_UPDATE":
		return &VoiceStateUpdate{}
	case "THREAD_UPDATE":
		return &ThreadUpdate{}
	case "GUILD_DELETE":
		return &GuildDelete{}
	case "INVITE_CREATE":
		return &InviteCreate{}
	case "CHANNEL_UPDATE":
		return &ChannelUpdate{}
	case "THREAD_DELETE":
		return &ThreadDelete{}
	case "MESSAGE_REACTION_REMOVE_ALL":
		return &MessageReactionRemoveAll{}
	case "PRESENCE_UPDATE":
		return &PresenceUpdate{}
	case "STAGE_INSTANCE_CREATE":
		return &StageInstanceCreate{}
	case "THREAD_LIST_SYNC":
		return &ThreadListSync{}
	case "GUILD_ROLE_UPDATE":
		return &GuildRoleUpdate{}
	case "MESSAGE_DELETE":
		return &MessageDelete{}
	case "GUILD_SCHEDULED_EVENT_USER_ADD":
		return &GuildScheduledEventUserAdd{}
	case "GUILD_MEMBER_REMOVE":
		return &GuildMemberRemove{}
	case "GUILD_BAN_REMOVE":
		return &GuildBanRemove{}
	case "CHANNEL_DELETE":
		return &ChannelDelete{}
	case "THREAD_MEMBER_UPDATE":
		return &ThreadMemberUpdate{}
	case "THREAD_MEMBERS_UPDATE":
		return &ThreadMembersUpdate{}
	case "USER_UPDATE":
		return &UserUpdate{}
	case "GUILD_STICKERS_UPDATE":
		return &GuildStickersUpdate{}
	case "INVITE_DELETE":
		return &InviteDelete{}
	case "CHANNEL_PINS_UPDATE":
		return &ChannelPinsUpdate{}
	case "THREAD_CREATE":
		return &ThreadCreate{}
	case "MESSAGE_BULK_DELETE":
		return &MessageBulkDelete{}
	case "MESSAGE_CREATE":
		return &MessageCreate{}
	case "MESSAGE_REACTION_REMOVE":
		return &MessageReactionRemove{}
	case "GUILD_MEMBERS_CHUNK":
		return &GuildMembersChunk{}
	case "GUILD_SCHEDULED_EVENT_USER_REMOVE":
		return &GuildScheduledEventUserRemove{}
	case "INTEGRATION_UPDATE":
		return &IntegrationUpdate{}
	case "WEBHOOKS_UPDATE":
		return &WebhooksUpdate{}
	case "GUILD_ROLE_CREATE":
		return &GuildRoleCreate{}
	case "INTEGRATION_DELETE":
		return &IntegrationDelete{}
	case "TYPING_START":
		return &TypingStart{}
	case "GUILD_CREATE":
		return &GuildCreate{}
	case "MESSAGE_REACTION_ADD":
		return &MessageReactionAdd{}
	case "STAGE_INSTANCE_UPDATE":
		return &StageInstanceUpdate{}
	case "GUILD_INTEGRATIONS_UPDATE":
		return &GuildIntegrationsUpdate{}
	case "GUILD_MEMBER_ADD":
		return &GuildMemberAdd{}
	case "VOICE_SERVER_UPDATE":
		return &VoiceServerUpdate{}
	case "STAGE_INSTANCE_DELETE":
		return &StageInstanceDelete{}
	case "CHANNEL_CREATE":
		return &ChannelCreate{}
	case "GUILD_BAN_ADD":
		return &GuildBanAdd{}
	case "GUILD_EMOJIS_UPDATE":
		return &GuildEmojisUpdate{}
	case "GUILD_MEMBER_UPDATE":
		return &GuildMemberUpdate{}
	default:
		return nil
	}
}
