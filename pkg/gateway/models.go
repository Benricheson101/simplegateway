package gateway

type Snowflake string

type ChannelType int

const (
	// a text channel within a server
	GUILD_TEXT ChannelType = 0
	// a direct message between users
	DM = 1
	// a voice channel within a server
	GUILD_VOICE = 2
	// a direct message between multiple users
	GROUP_DM = 3
	// an organizational category that contains up to 50 channels
	GUILD_CATEGORY = 4
	// a channel that users can follow and crosspost into their own server
	GUILD_NEWS = 5
	// a channel in which game developers can sell their game on Discord
	GUILD_STORE = 6
	// a temporary sub-channel within a GUILD_NEWS channel
	GUILD_NEWS_THREAD = 10
	// a temporary sub-channel within a GUILD_TEXT channel
	GUILD_PUBLIC_THREAD = 11
	// a temporary sub-channel within a GUILD_TEXT channel that is only viewable by those invited and those with the MANAGE_THREADS permission
	GUILD_PRIVATE_THREAD = 12
	// a voice channel for hosting events with an audience
	GUILD_STAGE_VOICE = 13
)

// TODO: add omitempty in relevant places
type Channel struct {
	// the id of this channel
	ID Snowflake `json:"id"`
	// the type of channel
	Type ChannelType `json:"type"`
	// the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	GuildID Snowflake `json:"guild_id"`
	// sorting position of the channel
	Position int `json:"position"`
	// explicit permission overwrites for members and roles
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	// the name of the channel (1-100 characters)
	Name string `json:"name"`
	// the channel topic (0-1024 characters)
	Topic string `json:"topic"`
	// whether the channel is nsfw
	Nsfw bool `json:"nsfw"`
	// the id of the last message sent in this channel (may not point to an existing or valid message)
	Last_message_id bool `json:"last_message_id"`
	// the bitrate (in bits) of the voice channel
	Bitrate int `json:"bitrate"`
	// the user limit of the voice channel
	UserLimit int `json:"user_limit"`
	// amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages or manage_channel, are unaffected
	RateLimitPerUser int `json:"rate_limit_per_user*"`
	// the recipients of the DM
	Recipients []User `json:"recipients"`
	// icon hash of the group DM
	Icon string `json:"icon"`
	// id of the creator of the group DM or thread
	OwnerID Snowflake `json:"owner_id"`
	// application id of the group DM creator if it is bot-created
	ApplicationID Snowflake `json:"application_id"`
	// for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	ParentID Snowflake `json:"parent_id"`
	// when the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned.
	LastPinTimestamp string `json:"last_pin_timestamp"`
	// voice region id for the voice channel, automatic when set to null
	RTCRegion string `json:"rtc_region"`
	// the camera video quality mode of the voice channel, 1 when not present
	VideoQualityMode int `json:"video_quality_mode"`
	// an approximate count of messages in a thread, stops counting at 50
	ThreadMessageCount int `json:"message_count"`
	// an approximate count of users in a thread, stops counting at 50
	ThreadMemberCount int `json:"member_count"`
	// thread-specific fields not needed by other channels
	ThreadMetadata ThreadMetadata `json:"thread_metadata"`
	// thread member object for the current user, if they have joined the thread, only included on certain API endpoints
	ThreadMember ThreadMember `json:"member"`
	// default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration"`
	// computed permissions for the invoking user in the channel, including overwrites, only included when part of the resolved data received on a slash command interaction
	Permissions string `json:"permissions"`
}

type PermissionOverwrite struct {
	// role or user id
	ID Snowflake `json:"id"`
	// either 0 (role) or 1 (member)
	Type int `json:"type"`
	// permission bit set
	Allow string `json:"allow"`
	// permission bit set
	Deny string `json:"deny"`
}

type User struct {
	// the user's id
	ID Snowflake `json:"id"`
	// the user's username, not unique across the platform
	Username string `json:"username"`
	// the user's 4-digit discord-tag
	Discriminator string `json:"discriminator"`
	// the user's avatar hash
	Avatar string `json:"avatar"`
	// whether the user belongs to an OAuth2 application
	Bot bool `json:"bot"`
	// whether the user is an Official Discord System user (part of the urgent message system)
	System bool `json:"system"`
	// whether the user has two factor enabled on their account
	MFAEnabled bool `json:"mfa_enabled"`
	// the user's banner hash
	Banner string `json:"banner"`
	// the user's banner color encoded as an int representation of hexadecimal color code
	AccentColor int `json:"accent_color"`
	// the user's chosen language option
	Locale string `json:"locale"`
	// whether the email on this account has been verified
	Verified bool `json:"verified"`
	// the user's email
	Email string `json:"email"`
	// the flags on a user's account
	Flags int `json:"flags"`
	// the type of Nitro subscription on a user's account
	PremiumType int `json:"premium_type"`
	// the public flags on a user's account
	PublicFlags int `json:"public_flags"`
}

type ThreadMetadata struct {
	// whether the thread is archived
	Archived bool `json:"archived"`
	// duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration"`
	// timestamp when the thread's archive status was last changed, used for calculating recent activity
	ArchiveTimestamp string `json:"archive_timestamp"`
	// whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
	Locked bool `json:"locked"`
	// whether non-moderators can add other non-moderators to a thread; only available on private threads
	Invitable bool `json:"invitable"`
	// timestamp when the thread was created; only populated for threads created after 2022-01-09
	CreateTimestamp string `json:"create_timestamp"`
}

type ThreadMember struct {
	// the id of the thread
	ID Snowflake `json:"id"`
	// the id of the user
	UserID Snowflake `json:"user_id"`
	// the time the current user last joined the thread
	JoinTimestamp string `json:"join_timestamp"`
	// any user-thread settings, currently only used for notifications
	Flags int `json:"flags"`
}

type UnavailableGuild struct {
	// guild id
	ID Snowflake `json:"id"`
	// true if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable"`
}

type Guild struct {
	// guild id
	ID Snowflake `json:"id"`
	// guild name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`
	// icon hash
	Icon string `json:"icon"`
	// icon hash, returned when in the template object
	Icon_hash string `json:"icon_hash"`
	// splash hash
	Splash string `json:"splash"`
	// discovery splash hash; only present for guilds with the "DISCOVERABLE"feature
	Discovery_splash string `json:"discovery_splash"`
	// true if the user is the owner of the guild
	Owner bool `json:"owner **"`
	// id of owner
	OwnerID Snowflake `json:"owner_id"`
	// total permissions for the user in the guild (excludes overwrites)
	Permissions string `json:"permissions"`
	// voice region id for the guild (deprecated)
	Region string `json:"region"`
	// id of afk channel
	AFKChannelID Snowflake `json:"afk_channel_id"`
	// afk timeout in seconds
	AFKTimeout int `json:"afk_timeout"`
	// true if the server widget is enabled
	WidgetEnabled bool `json:"widget_enabled"`
	// the channel id that the widget will generate an invite to, or null if set to no invite
	WidgetChannelID Snowflake `json:"widget_channel_id"`
	// verification level required for the guild
	VerificationLevel int `json:"verification_level"`
	// default message notifications level
	DefaultMessageNotifications int `json:"default_message_notifications"`
	// explicit content filter level
	ExplicitContentFilter int `json:"explicit_content_filter"`
	// roles in the guild
	Roles []Role `json:"roles"`
	// custom guild emojis
	Emojis []Emoji `json:"emojis"`
	// enabled guild features
	Features []string `json:"features"`
	// required MFA level for the guild
	MFALevel int `json:"mfa_level"`
	// application id of the guild creator if it is bot-created
	ApplicationID Snowflake `json:"application_id"`
	// the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelID Snowflake `json:"system_channel_id"`
	// system channel flags
	SystemChannelFlags int `json:"system_channel_flags"`
	// the id of the channel where Community guilds can display rules and/or guidelines
	RulesChannelID Snowflake `json:"rules_channel_id"`
	// when this guild was joined at
	JoinedAt string `json:"joined_at *"`
	// true if this is considered a large guild
	Large bool `json:"large"`
	// true if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable"`
	// total number of members in this guild
	MemberCount int `json:"member_count"`
	// states of members currently in voice channels; lacks the guild_id key
	VoiceStates []VoiceState `json:"voice_states *"`
	// users in the guild
	Members []GuildMember `json:"members *"`
	// channels in the guild
	Channels []Channel `json:"channels *"`
	// all active threads in the guild that current user has permission to view
	Threads []Channel `json:"threads *"`
	// presences of the members in the guild, will only include non-offline members if the size is greater than large threshold
	Presences []PresenceUpdate `json:"presences *"`
	// the maximum number of presences for the guild (null is always returned, apart from the largest of guilds)
	MaxPresences int `json:"max_presences"`
	// the maximum number of members for the guild
	MaxMembers int `json:"max_members"`
	// the vanity url code for the guild
	VanityURLCode string `json:"vanity_url_code"`
	// the description of a Community guild
	Description string `json:"description"`
	// banner hash
	Banner string `json:"banner"`
	// premium tier (Server Boost level)
	PremiumTier int `json:"premium_tier"`
	// the number of boosts this guild currently has
	PremiumSubscriptionCount int `json:"premium_subscription_count"`
	// the preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	PreferredLocale string `json:"preferred_locale"`
	// the id of the channel where admins and moderators of Community guilds receive notices from Discord
	PublicUpdatesChannelID Snowflake `json:"public_updates_channel_id"`
	// the maximum amount of users in a video channel
	MaxVideoChannelUsers int `json:"max_video_channel_users"`
	// approximate number of members in this guild, returned from the GET /guilds/<id> endpoint when with_counts is true
	ApproximateMemberCount int `json:"approximate_member_count"`
	// approximate number of non-offline members in this guild, returned from the GET /guilds/<id> endpoint when with_counts is true
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// the welcome screen of a Community guild, shown to new members, returned in an Invite's guild object
	WelcomeScreen WelcomeScreen `json:"welcome_screen"`
	// guild NSFW level
	NSFWLevel int `json:"nsfw_level"`
	// Stage instances in the guild
	StageInstances []StageInstance `json:"stage_instances *"`
	// custom guild stickers
	Stickers []Sticker `json:"stickers"`
	// the scheduled events in the guild
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events *"`
	// whether the guild has the boost progress bar enabled
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
}

type Role struct {
	// role id
	ID Snowflake `json:"id"`
	// role name
	Name string `json:"name"`
	// int representation of hexadecimal color code
	Color int `json:"color"`
	// if this role is pinned in the user listing
	Hoist bool `json:"hoist"`
	// role icon hash
	Icon string `json:"icon"`
	// role unicode emoji
	UnicodeEmoji string `json:"unicode_emoji"`
	// position of this role
	Position int `json:"position"`
	// permission bit set
	Permissions string `json:"permissions"`
	// whether this role is managed by an integration
	Managed bool `json:"managed"`
	// whether this role is mentionable
	Mentionable bool `json:"mentionable"`
	// the tags this role has
	Tags []RoleTag `json:"tags"`
}

type RoleTag struct {
	// the id of the bot this role belongs to
	BotID Snowflake `json:"bot_id"`
	// the id of the integration this role belongs to
	IntegrationID Snowflake `json:"integration_id"`
	// whether this is the guild's premium subscriber role
	PremiumSubscriber Snowflake `json:"premium_subscriber"`
}

type Emoji struct {
	// emoji id
	ID Snowflake `json:"id"`
	// emoji name
	Name string `json:"name"`
	// roles allowed to use this emoji
	Roles []Snowflake `json:"roles"`
	// user that created this emoji
	User User `json:"user"`
	// whether this emoji must be wrapped in colons
	RequireColons bool `json:"require_colons"`
	// whether this emoji is managed
	Managed bool `json:"managed"`
	// whether this emoji is animated
	Animated bool `json:"animated"`
	// whether this emoji can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
}

type VoiceState struct {
	// the guild id this voice state is for
	GuildID Snowflake `json:"guild_id"`
	// the channel id this user is connected to
	ChannelID Snowflake `json:"channel_id"`
	// the user id this voice state is for
	UserID Snowflake `json:"user_id"`
	// the guild member this voice state is for
	Member GuildMember `json:"member"`
	// the session id for this voice state
	SessionID string `json:"session_id"`
	// whether this user is deafened by the server
	Deaf bool `json:"deaf"`
	// whether this user is muted by the server
	Mute bool `json:"mute"`
	// whether this user is locally deafened
	SelfDeaf bool `json:"self_deaf"`
	// whether this user is locally muted
	SelfMute bool `json:"self_mute"`
	// whether this user is streaming using "Go Live"
	SelfStream bool `json:"self_stream"`
	// whether this user's camera is enabled
	SelfVideo bool `json:"self_video"`
	// whether this user is muted by the current user
	Suppress bool `json:"suppress"`
	// the time at which the user requested to speak
	RequestToSpeakTimestamp string `json:"request_to_speak_timestamp"`
}

type GuildMember struct {
	// the user this guild member represents
	User User `json:"user"`
	// this user's guild nickname
	Nick string `json:"nick"`
	// the member's guild avatar hash
	Avatar string `json:"avatar"`
	// array of role object ids
	Roles []Snowflake `json:"roles"`
	// when the user joined the guild
	JoinedAt string `json:"joined_at"`
	// when the user started boosting the guild
	PremiumSince string `json:"premium_since"`
	// whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`
	// whether the user is muted in voice channels
	Mute bool `json:"mute"`
	// whether the user has not yet passed the guild's Membership Screening requirements
	Pending bool `json:"pending"`
	// total permissions of the member in the channel, including overwrites, returned when in the interaction object
	Permissions string `json:"permissions"`
	// when the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
	CommunicationDisabledUntil string `json:"communication_disabled_until"`
}

type PresenceUpdate struct {
	// the user presence is being updated for
	User User `json:"user"`
	// id of the guild
	GuildID Snowflake `json:"guild_id"`
	// either "idle", "dnd", "online", or "offline"
	Status string `json:"status"`
	// user's current activities
	Activities []Activity `json:"activities"`
	// user's platform-dependent status
	ClientStatus ClientStatus `json:"client_status"`
}

type ClientStatus struct {
	// the user's status set for an active desktop (Windows, Linux, Mac) application session
	Desktop string `json:"desktop"`
	// the user's status set for an active mobile (iOS, Android) application session
	Mobile string `json:"mobile"`
	// the user's status set for an active web (browser, bot account) application session
	Web string `json:"web"`
}

type Activity struct {
	// the activity's name
	Name string `json:"name"`
	// activity type
	Type int `json:"type"`
	// stream url, is validated when type is 1
	URL string `json:"url"`
	// unix timestamp (in milliseconds) of when the activity was added to the user's session
	CreatedAT int `json:"created_at"`
	// unix timestamps for start and/or end of the game
	Timestamps ActivityTimestamp `json:"timestamps"`
	// application id for the game
	ApplicationID Snowflake `json:"application_id"`
	// what the player is currently doing
	Details string `json:"details"`
	// the user's current party status
	State string `json:"state"`
	// the emoji used for a custom status
	Emoji ActivityEmoji `json:"emoji"`
	// information for the current party of the player
	Party ActivityParty `json:"party"`
	// images for the presence and their hover texts
	Assets ActivityAssets `json:"assets"`
	// secrets for Rich Presence joining and spectating
	Secrets ActivitySecrets `json:"secrets"`
	// whether or not the activity is an instanced game session
	Instance bool `json:"instance"`
	// activity flags ORd together, describes what the payload includes
	Flags int `json:"flags"`
	// the custom buttons shown in the Rich Presence (max 2)
	Buttons []ActivityButton `json:"buttons"`
}

type ActivityTimestamp struct {
	// unix time (in milliseconds) of when the activity started
	Start int `json:"start"`
	// unix time (in milliseconds) of when the activity ends
	End int `json:"end"`
}

type ActivityEmoji struct {
	// the name of the emoji
	Name string `json:"name"`
	// the id of the emoji
	ID Snowflake `json:"id"`
	// whether this emoji is animated
	Animated bool `json:"animated"`
}

type ActivityParty struct {
	// the id of the party
	ID string `json:"id"`
	// used to show the party's current and maximum size
	Size [2]int `json:"size"`
}

type ActivityAssets struct {
	// see Activity Asset Image
	LargeImage string `json:"large_image"`
	// text displayed when hovering over the large image of the activity
	LargeText string `json:"large_text"`
	// see Activity Asset Image
	SmallImage string `json:"small_image"`
	// text displayed when hovering over the small image of the activity
	SmallText string `json:"small_text"`
}

type ActivitySecrets struct {
	// the secret for joining a party
	Join string `json:"join"`
	// the secret for spectating a game
	Spectate string `json:"spectate"`
	// the secret for a specific instanced match
	Match string `json:"match"`
}

type ActivityButton struct {
	// the text shown on the button (1-32 characters)
	Label string `json:" label"`
	// the url opened when clicking the button (1-512 characters)
	URL string `json:" url"`
}

type WelcomeScreen struct {
	// the server description shown in the welcome screen
	Description string `json:"description"`
	// the channels shown in the welcome screen, up to 5
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	// the channel's id
	ChannelID Snowflake `json:"channel_id"`
	// the description shown for the channel
	Description string `json:"description"`
	// the emoji id, if the emoji is custom
	EmojiID Snowflake `json:"emoji_id"`
	// the emoji name if custom, the unicode character if standard, or null if no emoji is set
	EmojiName string `json:"emoji_name"`
}

type StageInstance struct {
	// The id of this Stage instance
	ID Snowflake `json:"id"`
	// The guild id of the associated Stage channel
	GuildID Snowflake `json:"guild_id"`
	// The id of the associated Stage channel
	ChannelID Snowflake `json:"channel_id"`
	// The topic of the Stage instance (1-120 characters)
	Topic string `json:"topic"`
	// The privacy level of the Stage instance
	PrivacyLevel StagePrivacyLevel `json:"privacy_level"`
	// Whether or not Stage Discovery is disabled (deprecated)
	DiscoverableDisabled bool `json:"discoverable_disabled"`
}

type StagePrivacyLevel int

const (
	PUBLIC     StagePrivacyLevel = 1
	GUILD_ONLY                   = 2
)

type Sticker struct {
	// id of the sticker
	ID Snowflake `json:"id"`
	// for standard stickers, id of the pack the sticker is from
	PackID Snowflake `json:"pack_id"`
	// name of the sticker
	Name string `json:"name"`
	// description of the sticker
	Description string `json:"description"`
	// autocomplete/suggestion tags for the sticker (max 200 characters)
	Tags *string `json:"tags*"`
	// Deprecated previously the sticker asset hash, now an empty string
	Asset string `json:"asset"`
	// type of sticker
	Type StickerType `json:"type"`
	// type of sticker format
	FormatType StickerFormat `json:"format_type"`
	// whether this guild sticker can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
	// id of the guild that owns this sticker
	GuildID Snowflake `json:"guild_id"`
	// the user that uploaded the guild sticker
	User User `json:"user"`
	// the standard sticker's sort order within its pack
	SortValue int `json:"sort_value"`
}

type StickerType int

const (
	// an official sticker in a pack, part of Nitro or in a removed purchasable pack
	STANDARD StickerType = 1
	// a sticker uploaded to a Boosted guild for the guild's members
	GUILD = 2
)

type GuildScheduledEvent struct {
	// the id of the scheduled event
	ID Snowflake `json:"id"`
	// the guild id which the scheduled event belongs to
	GuildID Snowflake `json:"guild_id"`
	// the channel id in which the scheduled event will be hosted, or null if scheduled entity type is EXTERNAL
	ChannelID Snowflake `json:"channel_id"`
	// the id of the user that created the scheduled event *
	CreatorID Snowflake `json:"creator_id"`
	// the name of the scheduled event (1-100 characters)
	Name string `json:"name"`
	// the description of the scheduled event (1-1000 characters)
	Description string `json:"description"`
	// the time the scheduled event will start
	ScheduledStartTime string `json:"scheduled_start_time"`
	// the time the scheduled event will end, required if entity_type is EXTERNAL
	ScheduledEndTime string `json:"scheduled_end_time **"`
	// the privacy level of the scheduled event
	PrivacyLevel GuildScheduledEventPrivacyLevel `json:"privacy_level"`
	// the status of the scheduled event
	Status GuildScheduledEventStatus `json:"status"`
	// the type of the scheduled event
	EntityType GuildScheduledEventEntityType `json:"entity_type"`
	// the id of an entity associated with a guild scheduled event
	EntityID Snowflake `json:"entity_id"`
	// additional metadata for the guild scheduled event
	EntityMetadata GuildScheduledEventEntityMetadata `json:"entity_metadata **"`
	// the user that created the scheduled event
	Creator User `json:"creator"`
	// the number of users subscribed to the scheduled event
	UserCount int `json:"user_count"`
	// the cover image hash of the scheduled event
	Image string `json:"image"`
}

type GuildScheduledEventPrivacyLevel int

// the scheduled event is only accessible to guild members
const SCHEDULED_EVENT_GUILD_ONLY GuildScheduledEventPrivacyLevel = 2

type GuildScheduledEventStatus int

const (
	SCHEDULED GuildScheduledEventStatus = 1
	ACTIVE                              = 2
	COMPLETED                           = 3
	CANCELED                            = 4
)

type GuildScheduledEventEntityType int

const (
	STAGE_INSTANCE GuildScheduledEventEntityType = 1
	VOICE                                        = 2
	EXTERNAL                                     = 3
)

type GuildScheduledEventEntityMetadata struct {
	// location of the event (1-100 characters)
	Location string `json:"location"`
}

type Integration struct {
	// integration id
	ID Snowflake `json:"id"`
	// integration name
	Name string `json:"name"`
	// integration type (twitch, youtube, or discord)
	Type string `json:"type"`
	// is this integration enabled
	Enabled bool `json:"enabled"`
	// is this integration syncing
	Syncing bool `json:"syncing"`
	// id that this integration uses for "subscribers"
	RoleID Snowflake `json:"role_id"`
	// whether emoticons should be synced for this integration (twitch only currently)
	EnableEmoticons bool `json:"enable_emoticons"`
	// the behavior of expiring subscribers
	ExpireBehavior IntegrationExpireBehavior `json:"expire_behavior *"`
	// the grace period (in days) before expiring subscribers
	ExpireGracePeriod int `json:"expire_grace_period"`
	// user for this integration
	User User `json:"user *"`
	// integration account information
	Account IntegrationAccount `json:"account"`
	// when this integration was last synced
	SyncedAt string `json:"synced_at *"`
	// how many subscribers this integration has
	SubscriberCount int `json:"subscriber_count *"`
	// has this integration been revoked
	Revoked bool `json:"revoked"`
	// The bot/OAuth2 application for discord integrations
	Application IntegrationApplication `json:"application"`
}

type IntegrationExpireBehavior int

const (
	REMOVE_ROLE IntegrationExpireBehavior = 0
	KICK                                  = 1
)

type IntegrationAccount struct {
	// id of the account
	ID string `json:"id"`
	// name of the account
	Name string `json:"name"`
}

type IntegrationApplication struct {
	// the id of the app
	ID Snowflake `json:"id"`
	// the name of the app
	Name string `json:"name"`
	// the icon hash of the app
	Icon string `json:"icon"`
	// the description of the app
	Description string `json:"description"`
	// the summary of the app
	Summary string `json:"summary"`
	// the bot associated with this application
	Bot User `json:"bot"`
}

type InviteTargetType int

const (
	STREAM               InviteTargetType = 1
	EMBEDDED_APPLICATION                  = 2
)

type Application struct {
	// the id of the app
	ID Snowflake `json:"id"`
	// the name of the app
	Name string `json:"name"`
	// the icon hash of the app
	Icon string `json:"icon"`
	// the description of the app
	Description string `json:"description"`
	// an array of rpc origin urls, if rpc is enabled
	RPCOrigins []string `json:"rpc_origins"`
	// when false only app owner can join the app's bot to guilds
	BotPublic bool `json:"bot_public"`
	// when true the app's bot will only join upon completion of the full oauth2 code grant flow
	BotRequireCodeGrant bool `json:"bot_require_code_grant"`
	// the url of the app's terms of service
	TermsOfServiceURL string `json:"terms_of_service_url"`
	// the url of the app's privacy policy
	PrivacyPolicyURL string `json:"privacy_policy_url"`
	// partial user object containing info on the owner of the application
	Owner User `json:"owner"`
	// if this application is a game sold on Discord, this field will be the summary field for the store page of its primary sku
	Summary string `json:"summary"`
	// the hex encoded key for verification in interactions and the GameSDK's GetTicket
	VerifyKey string `json:"verify_key"`
	// if the application belongs to a team, this will be a list of the members of that team
	Team Team `json:"team"`
	// if this application is a game sold on Discord, this field will be the guild to which it has been linked
	GuildID Snowflake `json:"guild_id"`
	// if this application is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	PrimarySKUID Snowflake `json:"primary_sku_id"`
	// if this application is a game sold on Discord, this field will be the URL slug that links to the store page
	Slug string `json:"slug"`
	// the application's default rich presence invite cover image hash
	CoverImage string `json:"cover_image"`
	// the application's public flags
	Flags int `json:"flags"`
}

type Team struct {
	// a hash of the image of the team's icon
	Icon string `json:"icon"`
	// the unique id of the team
	ID Snowflake
	// the members of the team
	Members []TeamMember
	// the name of the team
	Name string
	// the user id of the current team owner
	OwnerUserID Snowflake
}

type TeamMember struct {
	// the user's membership state on the team
	MembershipState TeamMembershipState `json:"membership_state"`
	// will always be ["*"]
	Permissions []string `json:"permissions"`
	// the id of the parent team of which they are a member
	TeamID Snowflake `json:"team_id"`
	// the avatar, discriminator, id, and username of the user
	User User `json:"user"`
}

type TeamMembershipState int

const (
	INVITED  TeamMembershipState = 1
	ACCEPTED                     = 2
)

type Message struct {
	// id of the message
	ID Snowflake `json:"id"`
	// id of the channel the message was sent in
	ChannelID Snowflake `json:"channel_id"`
	// id of the guild the message was sent in
	GuildID Snowflake `json:"guild_id"`
	// the author of this message (not guaranteed to be a valid user, see below)
	Author User `json:"author*"`
	// member properties for this message's author
	Member GuildMember `json:"member**"`
	// contents of the message
	Content string `json:"content"`
	// when this message was sent
	Timestamp string `json:"timestamp"`
	// when this message was edited (or null if never)
	EditedTimestamp string `json:"edited_timestamp"`
	// whether this was a TTS message
	TTS bool `json:"tts"`
	// whether this message mentions everyone
	MentionEveryone bool `json:"mention_everyone"`
	// TODO: is this the right type
	// users specifically mentioned in the message
	Mentions []User `json:"mentions"`
	// roles specifically mentioned in this message
	MentionRoles []Snowflake `json:"mention_roles"`
	// channels specifically mentioned in this message
	MentionChannels []ChannelMention `json:"mention_channels****"`
	// any attached files
	Attachments []Attachment `json:"attachments"`
	// any embedded content
	Embeds []Embed `json:"embeds"`
	// reactions to the message
	Reactions []Reaction `json:"reactions"`
	// used for validating a message was sent
	Nonce string `json:"nonce"`
	// whether this message is pinned
	Pinned bool `json:"pinned"`
	// if the message is generated by a webhook, this is the webhook's id
	WebhookID Snowflake `json:"webhook_id"`
	// type of message
	Type int `json:"type"`
	// sent with Rich Presence-related chat embeds
	Activity MessageActivity `json:"activity"`
	// sent with Rich Presence-related chat embeds
	Application Application `json:"application"`
	// if the message is an Interaction or application-owned webhook, this is the id of the application
	ApplicationID Snowflake `json:"application_id"`
	// data showing the source of a crosspost, channel follow add, pin, or reply message
	MessageReference MessageReference `json:"message_reference"`
	// message flags combined as a bitfield
	Flags int `json:"flags"`
	// the message associated with the message_reference
	ReferencedMessage *Message `json:"referenced_message*****"`
	// sent if the message is a response to an Interaction
	// TODO: interactions
	// Interaction message interaction object `json:"interaction"`
	// the thread that was started from this message, includes thread member object
	Thread Channel `json:"thread"`
	// sent if the message contains components like buttons, action rows, or other interactive components
	// TODO: components
	// Components Array of message components `json:"components"`
	// sent if the message contains stickers
	StickerItems []StickerItem `json:"sticker_items"`
	// Deprecated the stickers sent with the message
	Stickers []Sticker `json:"stickers"`
}

type ChannelMention struct {
	// id of the channel
	ID Snowflake `json:"id"`
	// id of the guild containing the channel
	GuildID Snowflake `json:"guild_id"`
	// the type of channel
	Type int `json:"type"`
	// the name of the channel
	Name string `json:"name"`
}

type Attachment struct {
	// attachment id
	ID Snowflake `json:"id"`
	// name of file attached
	Filename string `json:"filename"`
	// description for the file
	Description string `json:"description"`
	// the attachment's media type
	Content_type string `json:"content_type"`
	// size of file in bytes
	Size int `json:"size"`
	// source url of file
	URL string `json:"url"`
	// a proxied url of file
	ProxyURL string
	// height of file (if image)
	Height int
	// width of file (if image)
	Width int
	// whether this attachment is ephemeral
	Ephemeral bool
}

type Embed struct {
	// title of embed
	Title string `json:"title"`
	// type of embed (always "rich" for webhook embeds)
	Type string `json:"type"`
	// description of embed
	Description string `json:"description"`
	// url of embed
	URL string `json:"url"`
	// timestamp of embed content
	Timestamp string `json:" timestamp"`
	// color code of the embed
	Color int `json:" color"`
	// footer information
	Footer EmbedFooter `json:" footer"`
	// image information
	Image EmbedImage `json:" image"`
	// thumbnail information
	Thumbnail EmbedThumbnail `json:" thumbnail"`
	// video information
	Video EmbedVideo `json:" video"`
	// provider information
	Provider EmbedProvider `json:" provider"`
	// author information
	Author EmbedAuthor `json:" author"`
	// fields information
	Fields []EmbedField `json:" fields"`
}

type EmbedFooter struct {
	// footer text
	Text string `json:"text"`
	// url of footer icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url"`
	// a proxied url of footer icon
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedImage struct {
	// source url of image (only supports http(s) and attachments)
	URL string `json:"url"`
	// a proxied url of the image
	ProxyURL string `json:"proxy_url"`
	// height of image
	Height int `json:"height"`
	// width of image
	Width int `json:"width"`
}

type EmbedVideo struct {
	// source url of video
	URL string `json:"url"`
	// a proxied url of the video
	ProxyURL string `json:"proxy_url"`
	// height of video
	Height int `json:"height"`
	// width of video
	Width int `json:"width"`
}

type EmbedProvider struct {
	// name of provider
	Name string `json:"name"`
	// url of provider
	URL string `json:"url"`
}

type EmbedAuthor struct {
	// name of author
	Name string `json:"name"`
	// url of author
	URL string `json:"url"`
	// url of author icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url"`
	// a proxied url of author icon
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedThumbnail struct {
	// source url of thumbnail (only supports http(s) and attachments)
	URL string `json:"url"`
	//  a proxied url of the thumbnail
	ProxyURL string
	// height of thumbnail
	Height int `json:"height"`
	// width of thumbnail
	Width int `json:"width"`
}

type EmbedField struct {
	// name of the field
	Name string `json:"name"`
	// value of the field
	Value string `json:"value"`
	// whether or not this field should display inline
	Inline bool `json:"inline"`
}

type Reaction struct {
	// times this emoji has been used to react
	Count int `json:"count"`
	// whether the current user reacted using this emoji
	Me bool `json:"me"`
	// emoji information
	Emoji Emoji `json:"emoji"`
}

type MessageActivity struct {
	// type of message activity
	Type MessageActivityType `json:"type"`
	// party_id from a Rich Presence event
	PartyID string `json:"party_id"`
}

type MessageActivityType int

const (
	JOIN         MessageActivityType = 1
	SPECTATE                         = 2
	LISTEN                           = 3
	JOIN_REQUEST                     = 5
)

type MessageReference struct {
	// id of the originating message
	MessageID Snowflake `json:"message_id"`
	// id of the originating message's channel
	ChannelID Snowflake `json:"channel_id *"`
	// id of the originating message's guild
	GuildID Snowflake `json:"guild_id"`
	// when sending, whether to error if the referenced message doesn't exist instead of sending as a normal (non-reply) message, default true
	FailIfNotExists bool `json:"fail_if_not_exists"`
}

type StickerItem struct {
	// id of the sticker
	ID Snowflake `json:"id"`
	// name of the sticker
	Name string `json:"name"`
	// type of sticker format
	FormatType StickerFormat `json:"format_type"`
}

type StickerFormat int

const (
	PNG    StickerFormat = 1
	APNG                 = 2
	LOTTIE               = 3
)
