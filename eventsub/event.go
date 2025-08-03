package eventsub

type AutomodMessageHoldEvent struct {
	// BroadcasterUserId is an ID of the broadcaster specified in the request.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the broadcaster specified in the request.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a username of the broadcaster specified in the request.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// UserId is a message sender’s user ID.
	UserId string `json:"user_id"`
	// UserLogin is a message sender’s login name.
	UserLogin string `json:"user_login"`
	// UserName is a message sender’s display name.
	UserName string `json:"user_name"`
	// MessageId is an ID of the message that was flagged by automod.
	MessageId string `json:"message_id"`
	// Message is a body of the message.
	Message Message `json:"message"`
	// Category is a category of the message.
	Category string `json:"category"`
	// Level is a level of severity.
	Level AutomodSeverityLevel `json:"level"`
	// HeldAt is a timestamp of when automod saved the message.
	HeldAt Timestamp `json:"held_at"`
}

type AutomodMessageHoldEventV2 struct {
	// BroadcasterUserId is an ID of the broadcaster specified in the request.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the broadcaster specified in the request.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a username of the broadcaster specified in the request.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// UserId is a message sender’s user ID.
	UserId string `json:"user_id"`
	// UserLogin is a message sender’s login name.
	UserLogin string `json:"user_login"`
	// UserName is a message sender’s display name.
	UserName string `json:"user_name"`
	// MessageId is an ID of the message that was flagged by automod.
	MessageId string `json:"message_id"`
	// Message is a body of the message.
	Message MessageV2 `json:"message"`
	// HeldAt is a timestamp of when automod saved the message.
	HeldAt Timestamp `json:"held_at"`
	// Reason is a reason why automod hold this message.
	Reason AutomodHoldReason `json:"reason"`
	// Automod is a metadata if message was caught by automod.
	Automod *Automod `json:"automod"`
	// Automod is a metadata if message was caught by blocked term.
	BlockedTerm *BlockedTerm `json:"blocked_term"`
}

type AutomodMessageUpdateEvent struct {
	// BroadcasterUserId is an ID of the broadcaster specified in the request.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the broadcaster specified in the request.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a username of the broadcaster specified in the request.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// UserId is a message sender’s user ID.
	UserId string `json:"user_id"`
	// UserLogin is a message sender’s login name.
	UserLogin string `json:"user_login"`
	// UserName is a message sender’s display name.
	UserName string `json:"user_name"`
	// ModeratorUserId is an ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
	// ModeratorUserName is a moderator’s username.
	ModeratorUserName string `json:"moderator_user_name"`
	// ModeratorUserLogin is a login of the moderator.
	ModeratorUserLogin string `json:"moderator_user_login"`
	// MessageId is an ID of the message that was flagged by automod.
	MessageId string `json:"message_id"`
	// Message is a body of the message.
	Message Message `json:"message"`
	// Category is a category of the message.
	Category string `json:"category"`
	// Level is a level of severity.
	Level AutomodSeverityLevel `json:"level"`
	// Status is a message’s status.
	Status AutomodMessageStatus `json:"status"`
	// HeldAt is a timestamp of when automod saved the message.
	HeldAt Timestamp `json:"held_at"`
}

type AutomodMessageUpdateEventV2 struct {
	// BroadcasterUserId is an ID of the broadcaster specified in the request.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the broadcaster specified in the request.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a username of the broadcaster specified in the request.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// UserId is a message sender’s user ID.
	UserId string `json:"user_id"`
	// UserLogin is a message sender’s login name.
	UserLogin string `json:"user_login"`
	// UserName is a message sender’s display name.
	UserName string `json:"user_name"`
	// ModeratorUserId is an ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
	// ModeratorUserName is a moderator’s username.
	ModeratorUserName string `json:"moderator_user_name"`
	// ModeratorUserLogin is a login of the moderator.
	ModeratorUserLogin string `json:"moderator_user_login"`
	// MessageId is an ID of the message that was flagged by automod.
	MessageId string `json:"message_id"`
	// Message is a body of the message.
	Message MessageV2 `json:"message"`
	// Status is a message’s status.
	Status AutomodMessageStatus `json:"status"`
	// HeldAt is a timestamp of when automod saved the message.
	HeldAt Timestamp `json:"held_at"`
	// Reason is a reason why automod hold this message.
	Reason AutomodHoldReason `json:"reason"`
	// Automod is a metadata if message was caught by automod.
	Automod *Automod `json:"automod"`
	// Automod is a metadata if message was caught by blocked term.
	BlockedTerm *BlockedTerm `json:"blocked_term"`
}

type AutomodSettingsUpdateEvent struct {
	// BroadcasterUserId is an ID of the broadcaster specified in the request.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the broadcaster specified in the request.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a username of the broadcaster specified in the request.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// ModeratorUserId is an ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
	// ModeratorUserName is a moderator’s username.
	ModeratorUserName string `json:"moderator_user_name"`
	// ModeratorUserLogin is a login of the moderator.
	ModeratorUserLogin string `json:"moderator_user_login"`
	// Bullying is an Automod level for hostility involving name-calling or insults.
	Bullying int `json:"bullying"`
	// OverallLevel is a default AutoMod level for the broadcaster.
	// This field is empty if the broadcaster has set one or more of the individual settings.
	OverallLevel int `json:"overall_level,omitempty"`
	// Disability is an Automod level for discrimination against disability.
	Disability int `json:"disability"`
	// RaceEthnicityOrReligion is an Automod level for racial discrimination.
	RaceEthnicityOrReligion int `json:"race_ethnicity_or_religion"`
	// Misogyny is an Automod level for discrimination against women.
	Misogyny int `json:"misogyny"`
	// SexualitySexOrGender is an AutoMod level for discrimination based on sexuality, sex, or gender.
	SexualitySexOrGender int `json:"sexuality_sex_or_gender"`
	// Aggression is an Automod level for hostility involving aggression.
	Aggression int `json:"aggression"`
	// SexBasedTerms is an Automod level for sexual content.
	SexBasedTerms int `json:"sex_based_terms"`
	// Swearing is an Automod level for profanity.
	Swearing int `json:"swearing"`
}

type AutomodTermsUpdateEvent struct {
	// BroadcasterUserId is an ID of the broadcaster specified in the request.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the broadcaster specified in the request.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a username of the broadcaster specified in the request.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// ModeratorUserId is an ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
	// ModeratorUserName is a moderator’s username.
	ModeratorUserName string `json:"moderator_user_name"`
	// ModeratorUserLogin is a login of the moderator.
	ModeratorUserLogin string `json:"moderator_user_login"`
	// Action is a status change applied to the terms.
	Action AutomodTermsAction `json:"action"`
	// FromAutomod indicates whether this term was added due to an Automod message approve/deny action.
	FromAutomod bool `json:"from_automod"`
	// Terms is a list of terms that had a status change.
	Terms []string `json:"terms"`
}

type ChannelBitsUseEvent struct {
	// BroadcasterUserId is a user ID of the channel where the bits were redeemed.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a login of the channel where the bits were used.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a display name of the channel where the bits were used.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// UserId is a user ID of the redeeming user.
	UserId string `json:"user_id"`
	// UserLogin is a login name of the redeeming user.
	UserLogin string `json:"user_login"`
	// UserName is a display name of the redeeming user.
	UserName string `json:"user_name"`
	// Bits is a number if bits used.
	Bits int `json:"bits"`
	// Type is a type of bits usage.
	Type BitsUseType `json:"type"`
	// Message is an object that contains the user message and emote information needed to recreate the message.
	Message *MessageV3 `json:"message,omitempty"`
	// PowerUp is a data about power-up.
	PowerUp *PowerUp `json:"power_up,omitempty"`
}

type ChannelUpdateEvent struct {
	// BroadcasterUserId is a broadcaster’s user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Title is a channel’s stream title.
	Title string `json:"title"`
	// Language is a channel’s broadcast language.
	Language string `json:"language"`
	// CategoryId is a channel’s category ID.
	CategoryId string `json:"category_id"`
	// CategoryName is a category name.
	CategoryName string `json:"category_name"`
	// ContentClassificationLabels is a content classification label IDs currently applied on the channel. To retrieve
	// a list of all possible IDs, use the Get Content Classification Labels API endpoint.
	ContentClassificationLabels []string `json:"content_classification_labels"`
}

type ChannelFollowEvent struct {
	// UserId is a user ID for the user now following the specified channel.
	UserId string `json:"user_id"`
	// UserLogin is a user login for the user now following the specified channel.
	UserLogin string `json:"user_login"`
	// UserName is a user display name for the user now following the specified channel.
	UserName string `json:"user_name"`
	// BroadcasterUserId is a broadcaster’s user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// FollowedAt is an RFC3339 timestamp of when the follow occurred.
	FollowedAt Timestamp `json:"followed_at"`
}

type ChannelAdBreakBeginEvent struct {
	// DurationSeconds is a length in seconds of the mid-roll ad break requested.
	DurationSeconds int `json:"duration_seconds"`
	// StartedAt is a UTC timestamp of when the ad break began, in RFC3339 format.
	// Note that there is potential delay between this event, when the streamer requested the ad break, and when
	// the viewers will see ads.
	StartedAt TimestampUTC `json:"started_at"`
	// IsAutomatic indicates if the ad was automatically scheduled via ads manager.
	IsAutomatic bool `json:"is_automatic"`
	// BroadcasterUserId is a broadcaster’s user ID for the channel the ad was run on.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a broadcaster’s user login for the channel the ad was run on.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a broadcaster’s user display name for the channel the ad was run on.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// BroadcasterUserId is an ID of the user that requested the ad. For automatic ads, this will be the ID of the broadcaster.
	RequesterUserId string `json:"requester_user_id"`
	// BroadcasterUserLogin is a login of the user that requested the ad.
	RequesterUserLogin string `json:"requester_user_login"`
	// BroadcasterUserName is a display name of the user that requested the ad.
	RequesterUserName string `json:"requester_user_name"`
}

type ChannelChatClearEvent struct {
	// BroadcasterUserId is a broadcaster’s user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
}

type ChannelChatClearUserMessagesEvent struct {
	// BroadcasterUserId is a broadcaster’s user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// TargetUserId is an ID of the user that was banned or put in a timeout. All of their messages are deleted.
	TargetUserId string `json:"target_user_id"`
	// TargetUserLogin is a user login of the user that was banned or put in a timeout.
	TargetUserLogin string `json:"target_user_login"`
	// TargetUserName is a username of the user that was banned or put in a timeout.
	TargetUserName string `json:"target_user_name"`
}

type ChannelChatMessageEvent struct {
	// BroadcasterUserId is a broadcaster’s user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// BroadcasterUserLogin is a broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// BroadcasterUserName is a broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// ChatterUserId is a user ID of the user that sent the message.
	ChatterUserId string `json:"chatter_user_id"`
	// ChatterUserLogin is a username of the user that sent the message.
	ChatterUserName string `json:"chatter_user_login"`
	// ChatterUserName is a user login of the user that sent the message.
	ChatterUserLogin string `json:"chatter_user_name"`
	// MessageId is a UUID that identifies the message.
	MessageId string `json:"message_id"`
	// Message is a structured chat message.
	Message MessageV4 `json:"message"`
	// Type is a type of message.
	Type MessageType `json:"message_type"`
	// Badges is a list of chat badges.
	Badges []Badge `json:"badges"`
	// Cheer is a metadata if this message is a chee
	Cheer *Cheer `json:"cheer,omitempty"`
	// Color is a color of the user’s name in the chat room. This is a hexadecimal RGB color code in the form, #&lt;RGB&gt;.
	// This tag may be empty if it is never set.
	Color string `json:"color"`
	// Reply is a metadata if this message is a reply.
	Reply *Reply `json:"reply,omitempty"`
	// ChannelPointsCustomRewardId is an ID of a channel points custom reward that was redeemed.
	ChannelPointsCustomRewardId string `json:"channel_points_custom_reward_id,omitempty"`
	// SourceBroadcasterUserId is a broadcaster user ID of the channel the message was sent from. Is null when the message
	// happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action happens
	// in the channel of a participant other than the broadcaster.
	SourceBroadcasterUserId string `json:"source_broadcaster_user_id,omitempty"`
	// SourceBroadcasterUserName is a username of the broadcaster of the channel the message was sent from. Is null when
	// the message happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action
	// happens in the channel of a participant other than the broadcaster.
	SourceBroadcasterUserName string `json:"source_broadcaster_user_name,omitempty"`
	// SourceBroadcasterUserLogin is a login of the broadcaster of the channel the message was sent from. Is null when
	// the message happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action
	// happens in the channel of a participant other than the broadcaster.
	SourceBroadcasterUserLogin string `json:"source_broadcaster_user_login,omitempty"`
	// SourceMessageId is a UUID that identifies the source message from the channel the message was sent from.
	// Is null when the message happens in the same channel as the broadcaster. Is not null when in a shared chat session,
	// and the action happens in the channel of a participant other than the broadcaster.
	SourceMessageId string `json:"source_message_id,omitempty"`
	// SourceBadges is a list of chat badges for the chatter in the channel the message was sent from.
	// Is null when the message happens in the same channel as the broadcaster. Is not null when in a shared chat session,
	// and the action happens in the channel of a participant other than the broadcaster.
	SourceBadges []Badge `json:"source_badges,omitempty"`
	// IsSourceOnly determines if a message delivered during a shared chat session is only sent to the source channel.
	// Has no effect if the message is not sent during a shared chat session.
	IsSourceOnly bool `json:"is_source_only"`
}

type ConduitShardDisabledEvent struct {
	// ConduitId The ID of the conduit.
	ConduitId string `json:"conduit_id"`
	// ShardId The ID of the disabled shard.
	ShardId string `json:"shard_id"`
	// Status The new status of the transport.
	Status string `json:"status"`
	// Transport The disabled transport.
	Transport ConduitShardDisabledEventTransport `json:"transport"`
}

type ChannelBanEvent struct {
	// The user ID for the user who was banned on the specified channel.
	UserId string `json:"user_id"`
	// The user login for the user who was banned on the specified channel.
	UserLogin string `json:"user_login"`
	// The user display name for the user who was banned on the specified channel.
	UserName string `json:"user_name"`
	// The requested broadcaster ID.
	BroadcasterUserID string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The user ID of the issuer of the ban.
	ModeratorUserID string `json:"moderator_user_id"`
	// The user login of the issuer of the ban.
	ModeratorUserLogin string `json:"moderator_user_login"`
	// The user name of the issuer of the ban.
	ModeratorUserName string `json:"moderator_user_name"`
	// The reason behind the ban.
	Reason string `json:"reason"`
	// Will be null if permanent ban. If it is a timeout, this field shows when the timeout will end.
	EndsAt TimestampUTC `json:"ends_at"`
	// Indicates whether the ban is permanent (true) or a timeout (false). If true, ends_at will be null.
	IsPermanent bool `json:"is_permanent"`
}

type ChannelChatNotificationEvent struct {
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// Optional. The broadcaster user ID of the channel the message was sent from. Is null when the message notification happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action happens in the channel of a participant other than the broadcaster.
	SourceBroadcasterUserId string `json:"source_broadcaster_user_id,omitempty"`
	// Optional. The user name of the broadcaster of the channel the message was sent from. Is null when the message notification happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action happens in the channel of a participant other than the broadcaster.
	SourceBroadcasterUserName string `json:"source_broadcaster_user_name,omitempty"`
	// Optional. The login of the broadcaster of the channel the message was sent from. Is null when the message notification happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action happens in the channel of a participant other than the broadcaster.
	SourceBroadcasterUserLogin string `json:"source_broadcaster_user_login,omitempty"`
	// Optional. The UUID that identifies the source message from the channel the message was sent from. Is null when the message happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action happens in the channel of a participant other than the broadcaster.
	SourceMessageId string `json:"source_message_id,omitempty"`

	// Optional. The list of chat badges for the chatter in the channel the message was sent from. Is null when the message happens in the same channel as the broadcaster. Is not null when in a shared chat session, and the action happens in the channel of a participant other than the broadcaster.
	SourceBadges []ChatNotificationEventBadge `json:"source_badges,omitempty"`
	// The user ID of the user that sent the message.
	ChatterUserId string `json:"chatter_user_id"`
	// The user name of the user that sent the message.
	ChatterUserName string `json:"chatter_user_name"`
	// The user login of the user that sent the message.
	ChatterUserLogin string `json:"chatter_user_login"`
	// Whether or not the chatter is anonymous.
	ChatterIsAnonymous bool `json:"chatter_is_anonymous"`
	// The color of the user’s name in the chat room.
	Color string `json:"color"`
	// List of chat badges.
	Badges []ChatNotificationEventBadge `json:"badges"`
	// The message Twitch shows in the chat room for this notice.
	SystemMessage string `json:"system_message"`
	// A UUID that identifies the message.
	MessageId string `json:"message_id"`
	// The structured chat message
	Message ChatNotificationEventMessage `json:"message"`
	// The type of notice. Possible values are:
	// - sub
	// - resub
	// - sub_gift
	// - community_sub_gift
	// - gift_paid_upgrade
	// - prime_paid_upgrade
	// - raid
	// - unraid
	// - pay_it_forward
	// - announcement
	// - bits_badge_tier
	// - charity_donation
	// - shared_chat_sub
	// - shared_chat_resub
	// - shared_chat_sub_gift
	// - shared_chat_community_sub_gift
	// - shared_chat_gift_paid_upgrade
	// - shared_chat_prime_paid_upgrade
	// - shared_chat_raid
	// - shared_chat_pay_it_forward
	// - shared_chat_announcement
	NoticeType ChatNotificationEventNoticeType `json:"notice_type"`
	// Information about the sub event. Null if notice_type is not sub.
	Sub *ChatNotificationEventSubEvent `json:"sub,omitempty"`
	// Information about the resub event. Null if notice_type is not resub.
	Resub *ChatNotificationEventResubEvent `json:"resub,omitempty"`
	// Information about the gift sub event. Null if notice_type is not sub_gift.
	SubGift *ChatNotificationEventSubGiftEvent `json:"sub_gift,omitempty"`
	// Information about the community gift sub event. Null if notice_type is not community_sub_gift.
	CommunitySubGift *ChatNotificationEventCommunitySubGiftEvent `json:"community_sub_gift,omitempty"`
	// Information about the community gift paid upgrade event. Null if notice_type is not gift_paid_upgrade.
	GiftPaidUpgrade *ChatNotificationEventGiftPaidUpgradeEvent `json:"gift_paid_upgrade,omitempty"`
	// Information about the Prime gift paid upgrade event. Null if notice_type is not prime_paid_upgrade.
	PrimePaidUpgrade *ChatNotificationEventGiftPaidUpgradeEvent `json:"prime_paid_upgrade,omitempty"`
	// Information about the raid event. Null if notice_type is not raid.
	Raid *ChatNotificationEventRaidEvent `json:"raid,omitempty"`
	// Returns an empty payload if notice_type is unraid, otherwise returns null.
	Unraid *ChatNotificationEventUnraidEvent `json:"unraid,omitempty"`
	// Information about the pay it forward event. Null if notice_type is not pay_it_forward.
	PayItForward *ChatNotificationEventPayItForwardEvent `json:"pay_it_forward,omitempty"`
	// Information about the announcement event. Null if notice_type is not announcement.
	Announcement *ChatNotificationEventAnnouncementEvent `json:"announcement,omitempty"`
	// Information about the charity donation event. Null if notice_type is not charity_donation.
	CharityDonation *ChatNotificationEventCharityDonationEvent `json:"charity_donation,omitempty"`
	// Information about the bits badge tier event. Null if notice_type is not bits_badge_tier.
	BitsBadgeTier *ChatNotificationEventBitsBadgeTierEvent `json:"bits_badge_tier,omitempty"`
	// Information about the shared_chat_sub event. Is null if notice_type is not shared_chat_sub. This field has the same information as the sub field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatSub *ChatNotificationEventSubEvent `json:"shared_chat_sub,omitempty"`
	// Information about the shared_chat_resub event. Is null if notice_type is not shared_chat_resub. This field has the same information as the resub field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatResub *ChatNotificationEventResubEvent `json:"shared_chat_resub,omitempty"`
	// Information about the shared_chat_sub_gift event. Is null if notice_type is not shared_chat_sub_gift. This field has the same information as the chat_sub_gift field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatSubGift *ChatNotificationEventSubGiftEvent `json:"shared_chat_sub_gift,omitempty"`
	// Information about the shared_chat_community_sub_gift event. Is null if notice_type is not shared_chat_community_sub_gift. This field has the same information as the community_sub_gift field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatCommunitySubGift *ChatNotificationEventCommunitySubGiftEvent `json:"shared_chat_community_sub_gift,omitempty"`
	// Information about the shared_chat_gift_paid_upgrade event. Is null if notice_type is not shared_chat_gift_paid_upgrade. This field has the same information as the gift_paid_upgrade field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatGiftPaidUpgrade *ChatNotificationEventGiftPaidUpgradeEvent `json:"shared_chat_gift_paid_upgrade,omitempty"`
	// Information about the shared_chat_chat_prime_paid_upgrade event. Is null if notice_type is not shared_chat_prime_paid_upgrade. This field has the same information as the prime_paid_upgrade field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatPrimePaidUpgrade *ChatNotificationEventGiftPaidUpgradeEvent `json:"shared_chat_prime_paid_upgrade,omitempty"`
	// Information about the shared_chat_pay_it_forward event. Is null if notice_type is not shared_chat_pay_it_forward. This field has the same information as the pay_it_forward field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatPayItForward *ChatNotificationEventPayItForwardEvent `json:"shared_chat_pay_it_forward,omitempty"`
	// Information about the shared_chat_raid event. Is null if notice_type is not shared_chat_raid. This field has the same information as the raid field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatRaid *ChatNotificationEventRaidEvent `json:"shared_chat_raid,omitempty"`
	// Information about the shared_chat_announcement event. Is null if notice_type is not shared_chat_announcement. This field has the same information as the announcement field but for a notice that happened for a channel in a shared chat session other than the broadcaster in the subscription condition.
	SharedChatAnnouncement *ChatNotificationEventAnnouncementEvent `json:"shared_chat_announcement,omitempty"`
}

type ChannelModeratorAddEvent struct {
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The user ID of the new moderator.
	UserId string `json:"user_id"`
	// The user login of the new moderator.
	UserLogin string `json:"user_login"`
	// The display name of the new moderator.
	UserName string `json:"user_name"`
}

type ChannelModeratorRemoveEvent struct {
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The user ID of the removed moderator.
	UserId string `json:"user_id"`
	// The user login of the removed moderator.
	UserLogin string `json:"user_login"`
	// The display name of the removed moderator.
	UserName string `json:"user_name"`
}

type ChannelPollBeginEvent struct {
	// ID of the poll.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Question displayed for the poll.
	Title string `json:"title"`
	// An array of choices for the poll.
	Choices []ChannelPollEventChoice `json:"choices"`
	// The Bits voting settings for the poll.
	BitsVoting ChannelPollEventBitsVoting `json:"bits_voting"`
	// The Channel Points voting settings for the poll.
	ChannelPointsVoting ChannelPollEventChannelPointsVoting `json:"channel_points_voting"`
	// The time the poll started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the poll will end.
	EndsAt TimestampUTC `json:"ends_at"`
}

type ChannelPollProgressEvent struct {
	// ID of the poll.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Question displayed for the poll.
	Title string `json:"title"`
	// An array of choices for the poll. Includes vote counts.
	Choices []ChannelPollEventChoice `json:"choices"`
	// The Bits voting settings for the poll.
	BitsVoting ChannelPollEventBitsVoting `json:"bits_voting"`
	// The Channel Points voting settings for the poll.
	ChannelPointsVoting ChannelPollEventChannelPointsVoting `json:"channel_points_voting"`
	// The time the poll started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the poll will end.
	EndsAt TimestampUTC `json:"ends_at"`
}

type ChannelPollEndEvent struct {
	// ID of the poll.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Question displayed for the poll.
	Title string `json:"title"`
	// An array of choices for the poll. Includes vote counts.
	Choices []ChannelPollEventChoice `json:"choices"`
	// The Bits voting settings for the poll.
	BitsVoting ChannelPollEventBitsVoting `json:"bits_voting"`
	// The Channel Points voting settings for the poll.
	ChannelPointsVoting ChannelPollEventChannelPointsVoting `json:"channel_points_voting"`
	// The status of the poll. Valid values are completed, archived, and terminated.
	Status ChannelPollEndEventStatus `json:"status"`
	// The time the poll started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the poll ended.
	EndedAt TimestampUTC `json:"ended_at"`
}

type ChannelPredictionBeginEvent struct {
	// Channel Points Prediction ID.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Title for the Channel Points Prediction.
	Title string `json:"title"`
	// An array of outcomes for the Channel Points Prediction.
	Outcomes []ChannelPredictionEventOutcome `json:"outcomes"`
	// The time the Channel Points Prediction started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the Channel Points Prediction will automatically lock.
	LocksAt TimestampUTC `json:"locks_at"`
}

type ChannelPredictionProgressEvent struct {
	// Channel Points Prediction ID.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Title for the Channel Points Prediction.
	Title string `json:"title"`
	// An array of outcomes for the Channel Points Prediction. Includes top_predictors.
	Outcomes []ChannelPredictionEventOutcome `json:"outcomes"`
	// The time the Channel Points Prediction started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the Channel Points Prediction will automatically lock.
	LocksAt TimestampUTC `json:"locks_at"`
}

type ChannelPredictionLockEvent struct {
	// Channel Points Prediction ID.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Title for the Channel Points Prediction.
	Title string `json:"title"`
	// An array of outcomes for the Channel Points Prediction. Includes top_predictors.
	Outcomes []ChannelPredictionEventOutcome `json:"outcomes"`
	// The time the Channel Points Prediction started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the Channel Points Prediction was locked.
	LockedAt TimestampUTC `json:"locked_at"`
}

type ChannelPredictionEndEvent struct {
	// Channel Points Prediction ID.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Title for the Channel Points Prediction.
	Title string `json:"title"`
	// ID of the winning outcome.
	WinningOutcomeId string `json:"winning_outcome_id"`
	// An array of outcomes for the Channel Points Prediction. Includes top_predictors.
	Outcomes []ChannelPredictionEventOutcome `json:"outcomes"`
	// The status of the Channel Points Prediction. Valid values are resolved and canceled.
	Status ChannelPredictionEndEventStatus `json:"status"`
	// The time the Channel Points Prediction started.
	StartedAt TimestampUTC `json:"started_at"`
	// The time the Channel Points Prediction ended.
	EndedAt TimestampUTC `json:"ended_at"`
}

type ChannelRaidEvent struct {
	// The broadcaster ID that created the raid.
	FromBroadcasterUserId string `json:"from_broadcaster_user_id"`
	// The broadcaster login that created the raid.
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	// The broadcaster display name that created the raid.
	FromBroadcasterUserName string `json:"from_broadcaster_user_name"`
	// The broadcaster ID that received the raid.
	ToBroadcasterUserId string `json:"to_broadcaster_user_id"`
	// The broadcaster login that received the raid.
	ToBroadcasterUserLogin string `json:"to_broadcaster_user_login"`
	// The broadcaster display name that received the raid.
	ToBroadcasterUserName string `json:"to_broadcaster_user_name"`
	// The number of viewers in the raid.
	Viewers int `json:"viewers"`
}

type ChannelPointsCustomRewardRedemptionAddEvent struct {
	// The redemption identifier.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// User ID of the user that redeemed the reward.
	UserId string `json:"user_id"`
	// Login of the user that redeemed the reward.
	UserLogin string `json:"user_login"`
	// Display name of the user that redeemed the reward.
	UserName string `json:"user_name"`
	// The user input provided. Empty string if not provided.
	UserInput string `json:"user_input"`
	// Defaults to unfulfilled. Possible values are unknown, unfulfilled, fulfilled, and canceled.
	Status string `json:"status"`
	// Basic information about the reward that was redeemed, at the time it was redeemed.
	Reward ChannelPointsCustomEventReward `json:"reward"`
	// RFC3339 timestamp of when the reward was redeemed.
	RedeemedAt TimestampUTC `json:"redeemed_at"`
}

type ChannelPointsCustomRewardRedemptionUpdateEvent struct {
	// The redemption identifier.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// User ID of the user that redeemed the reward.
	UserId string `json:"user_id"`
	// Login of the user that redeemed the reward.
	UserLogin string `json:"user_login"`
	// Display name of the user that redeemed the reward.
	UserName string `json:"user_name"`
	// The user input provided. Empty string if not provided.
	UserInput string `json:"user_input"`
	// Will be fulfilled or canceled. Possible values are unknown, unfulfilled, fulfilled, and canceled.
	Status string `json:"status"`
	// Basic information about the reward that was redeemed, at the time it was redeemed.
	Reward ChannelPointsCustomEventReward `json:"reward"`
	// RFC3339 timestamp of when the reward was redeemed.
	RedeemedAt TimestampUTC `json:"redeemed_at"`
}

type ChannelPointsAutomaticRewardRedemptionAddEvent struct {
	// The ID of the channel where the reward was redeemed.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The login of the channel where the reward was redeemed.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The display name of the channel where the reward was redeemed.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The ID of the redeeming user.
	UserId string `json:"user_id"`
	// The login of the redeeming user.
	UserLogin string `json:"user_login"`
	// The display name of the redeeming user.
	UserName string `json:"user_name"`
	// The ID of the Redemption.
	Id string `json:"id"`
	// An object that contains the reward information.
	Reward ChannelPointsAutomaticRewardEventReward `json:"reward"`
	// An object that contains the user message and emote information needed to recreate the message.
	Message ChannelPointsAutomaticRewardEventRewardMessage `json:"message"`
	// Optional. A string that the user entered if the reward requires input.
	UserInput string `json:"user_input,omitempty"`
	// The UTC date and time (in RFC3339 format) of when the reward was
	RedeemedAt TimestampUTC `json:"redeemed_at"`
}

type ChannelPointsAutomaticRewardRedemptionAddEventV2 struct {
	// The ID of the channel where the reward was redeemed.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The login of the channel where the reward was redeemed.
	BroadcasterUserLogin string `json:"broadcaster_user~_login"`
	// The display name of the channel where the reward was redeemed.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The ID of the redeeming user.
	UserId string `json:"user_id"`
	// The login of the redeeming user.
	UserLogin string `json:"user_login"`
	// The display name of the redeeming user.
	UserName string `json:"user_name"`
	// The ID of the Redemption.
	Id string `json:"id"`
	// An object that contains the reward information.
	Reward ChannelPointsAutomaticRewardEventRewardV2 `json:"reward"`
	// An object that contains the user message and emote information needed to recreate the message.
	Message ChannelPointsAutomaticRewardEventRewardMessageV2 `json:"message"`
	// The UTC date and time (in RFC3339 format) of when the reward was redeemed.
	RedeemedAt TimestampUTC `json:"redeemed_at"`
}

type UserAuthorizationRevokeEvent struct {
	// The client_id of the application with revoked user access.
	ClientId string `json:"client_id"`
	// The user id for the user who has revoked authorization for your client id.
	UserId string `json:"user_id"`
	// The user login for the user who has revoked authorization for your client id. This is null if the user no longer exists.
	UserLogin string `json:"user_login"`
	// The user display name for the user who has revoked authorization for your client id. This is null if the user no longer exists.
	UserName string `json:"user_name"`
}

type ChannelPointsCustomRewardAddEvent struct {
	// The reward identifier.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Is the reward currently enabled. If false, the reward won’t show up to viewers.
	IsEnabled bool `json:"is_enabled"`
	// Is the reward currently paused. If true, viewers can’t redeem.
	IsPaused bool `json:"is_paused"`
	// Is the reward currently in stock. If false, viewers can’t redeem.
	IsInStock bool `json:"is_in_stock"`
	// The reward title.
	Title string `json:"title"`
	// The reward cost.
	Cost int `json:"cost"`
	// The reward description.
	Prompt string `json:"prompt"`
	// Does the viewer need to enter information when redeeming the reward.
	IsUserInputRequired bool `json:"is_user_input_required"`
	// Should redemptions be set to fulfilled status immediately when redeemed and skip the request queue instead of the normal unfulfilled status.
	ShouldRedemptionsSkipRequestQueue bool `json:"should_redemptions_skip_request_queue"`
	// Whether a maximum per stream is enabled and what the maximum is.
	MaxPerStream ChannelPointsRewardEventMaxPerStream `json:"max_per_stream"`
	// Whether a maximum per user per stream is enabled and what the maximum is.
	MaxPerUserPerStream ChannelPointsRewardEventMaxPerUserPerStream `json:"max_per_user_per_stream"`
	// Custom background color for the reward. Format: Hex with # prefix. Example: #FA1ED2.
	BackgroundColor string `json:"background_color"`
	// Set of custom images of 1x, 2x and 4x sizes for the reward. Can be null if no images have been uploaded.
	Image *ChannelPointsRewardEventImage `json:"image"`
	// Set of default images of 1x, 2x and 4x sizes for the reward.
	DefaultImage ChannelPointsRewardEventImage `json:"default_image"`
	// Whether a cooldown is enabled and what the cooldown is in seconds.
	GlobalCooldown ChannelPointsRewardEventGlobalCooldown `json:"global_cooldown"`
	// Timestamp of the cooldown expiration. null if the reward isn’t on cooldown.
	CooldownExpiresAt TimestampUTC `json:"cooldown_expires_at,omitempty"`
	// The number of redemptions redeemed during the current live stream. Counts against the max_per_stream limit.
	// null if the broadcasters stream isn’t live or max_per_stream isn’t enabled.
	RedemptionsRedeemedCurrentStream int `json:"redemptions_redeemed_current_stream"`
}

type ChannelPointsCustomRewardUpdateEvent struct {
	// The reward identifier.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Is the reward currently enabled. If false, the reward won’t show up to viewers.
	IsEnabled bool `json:"is_enabled"`
	// Is the reward currently paused. If true, viewers can’t redeem.
	IsPaused bool `json:"is_paused"`
	// Is the reward currently in stock. If false, viewers can’t redeem.
	IsInStock bool `json:"is_in_stock"`
	// The reward title.
	Title string `json:"title"`
	// The reward cost.
	Cost int `json:"cost"`
	// The reward description.
	Prompt string `json:"prompt"`
	// Does the viewer need to enter information when redeeming the reward.
	IsUserInputRequired bool `json:"is_user_input_required"`
	// Should redemptions be set to fulfilled status immediately when redeemed and skip the request queue instead of the normal unfulfilled status.
	ShouldRedemptionsSkipRequestQueue bool `json:"should_redemptions_skip_request_queue"`
	// Whether a maximum per stream is enabled and what the maximum is.
	MaxPerStream ChannelPointsRewardEventMaxPerStream `json:"max_per_stream"`
	// Whether a maximum per user per stream is enabled and what the maximum is.
	MaxPerUserPerStream ChannelPointsRewardEventMaxPerUserPerStream `json:"max_per_user_per_stream"`
	// Custom background color for the reward. Format: Hex with # prefix. Example: #FA1ED2.
	BackgroundColor string `json:"background_color"`
	// Set of custom images of 1x, 2x and 4x sizes for the reward. Can be null if no images have been uploaded.
	Image *ChannelPointsRewardEventImage `json:"image"`
	// Set of default images of 1x, 2x and 4x sizes for the reward.
	DefaultImage ChannelPointsRewardEventImage `json:"default_image"`
	// Whether a cooldown is enabled and what the cooldown is in seconds.
	GlobalCooldown ChannelPointsRewardEventGlobalCooldown `json:"global_cooldown"`
	// Timestamp of the cooldown expiration. null if the reward isn’t on cooldown.
	CooldownExpiresAt TimestampUTC `json:"cooldown_expires_at,omitempty"`
	// The number of redemptions redeemed during the current live stream. Counts against the max_per_stream limit.
	// null if the broadcasters stream isn’t live or max_per_stream isn’t enabled.
	RedemptionsRedeemedCurrentStream int `json:"redemptions_redeemed_current_stream,omitempty"`
}

type ChannelPointsCustomRewardRemoveEvent struct {
	// The reward identifier.
	Id string `json:"id"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Is the reward currently enabled. If false, the reward won’t show up to viewers.
	IsEnabled bool `json:"is_enabled"`
	// Is the reward currently paused. If true, viewers can’t redeem.
	IsPaused bool `json:"is_paused"`
	// Is the reward currently in stock. If false, viewers can’t redeem.
	IsInStock bool `json:"is_in_stock"`
	// The reward title.
	Title string `json:"title"`
	// The reward cost.
	Cost int `json:"cost"`
	// The reward description.
	Prompt string `json:"prompt"`
	// Does the viewer need to enter information when redeeming the reward.
	IsUserInputRequired bool `json:"is_user_input_required"`
	// Should redemptions be set to fulfilled status immediately when redeemed and skip the request queue instead of the normal unfulfilled status.
	ShouldRedemptionsSkipRequestQueue bool `json:"should_redemptions_skip_request_queue"`
	// Whether a maximum per stream is enabled and what the maximum is.
	MaxPerStream ChannelPointsRewardEventMaxPerStream `json:"max_per_stream"`
	// Whether a maximum per user per stream is enabled and what the maximum is.
	MaxPerUserPerStream ChannelPointsRewardEventMaxPerUserPerStream `json:"max_per_user_per_stream"`
	// Custom background color for the reward. Format: Hex with # prefix. Example: #FA1ED2.
	BackgroundColor string `json:"background_color"`
	// Set of custom images of 1x, 2x and 4x sizes for the reward. Can be null if no images have been uploaded.
	Image *ChannelPointsRewardEventImage `json:"image"`
	// Set of default images of 1x, 2x and 4x sizes for the reward.
	DefaultImage ChannelPointsRewardEventImage `json:"default_image"`
	// Whether a cooldown is enabled and what the cooldown is in seconds.
	GlobalCooldown ChannelPointsRewardEventGlobalCooldown `json:"global_cooldown"`
	// Timestamp of the cooldown expiration. null if the reward isn’t on cooldown.
	CooldownExpiresAt TimestampUTC `json:"cooldown_expires_at,omitempty"`
	// The number of redemptions redeemed during the current live stream. Counts against the max_per_stream limit.
	// null if the broadcasters stream isn’t live or max_per_stream isn’t enabled.
	RedemptionsRedeemedCurrentStream int `json:"redemptions_redeemed_current_stream,omitempty"`
}

type StreamOnlineEvent struct {
	// The id of the stream.
	Id string `json:"id"`
	// The broadcaster’s user id.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The stream type. Valid values are: live, playlist, watch_party, premiere, rerun.
	Type StreamOnlineEventType `json:"type"`
	// The timestamp at which the stream went online at.
	StartedAt TimestampUTC `json:"started_at"`
}

type StreamOfflineEvent struct {
	// The broadcaster’s user id.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster’s user login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster’s user display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
}

// ChannelSubscribeEvent A notification is sent when a specified channel receives a subscriber. This does not include resubscribes.
type ChannelSubscribeEvent struct {
	// The user ID for the user who subscribed to the specified channel.
	UserId string `json:"user_id"`
	// The user login for the user who subscribed to the specified channel.
	UserLogin string `json:"user_login"`
	// The user display name for the user who subscribed to the specified channel.
	UserName string `json:"user_name"`
	// The requested broadcaster ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The requested broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The requested broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The tier of the subscription. Valid values are 1000, 2000, and 3000.
	Tier string `json:"tier"`
	// Whether the subscription is a gift.
	IsGift bool `json:"is_gift"`
}

// ChannelSubscriptionMessageEvent A notification when a user sends a resubscription chat message in a specific channel.
type ChannelSubscriptionMessageEvent struct {
	// The user ID of the user who sent a resubscription chat message.
	UserId string `json:"user_id"`
	// The user login of the user who sent a resubscription chat message.
	UserLogin string `json:"user_login"`
	// The user display name of the user who a resubscription chat message.
	UserName string `json:"user_name"`
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The tier of the user’s subscription.
	Tier string `json:"tier"`
	// An object that contains the resubscription message and emote information needed to recreate the message.
	Message ChannelSubscriptionMessageEventMessage `json:"message"`
	// The total number of months the user has been subscribed to the channel.
	CumulativeTotal int `json:"cumulative_total"`
	// The number of consecutive months the user’s current subscription has been active.
	// This value is null if the user has opted out of sharing this information.
	StreakMonths int `json:"streak_months"`
	// The month duration of the subscription.
	DurationMonths int `json:"duration_months"`
}

// ChannelSubscriptionEndEvent A notification when a subscription to the specified channel ends.
type ChannelSubscriptionEndEvent struct {
	// The user ID for the user whose subscription ended.
	UserId string `json:"user_id"`
	// The user login for the user whose subscription ended.
	UserLogin string `json:"user_login"`
	// The user display name for the user whose subscription ended.
	UserName string `json:"user_name"`
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The tier of the subscription that ended. Valid values are 1000, 2000, and 3000.
	Tier string `json:"tier"`
	// Whether the subscription was a gift.
	IsGift bool `json:"is_gift"`
}

// ChannelSubscriptionGiftEvent A notification when a viewer gives a gift subscription to one or more users in the specified channel.
type ChannelSubscriptionGiftEvent struct {
	// The user ID of the user who sent the subscription gift. Set to null if it was an anonymous subscription gift.
	UserId string `json:"user_id"`
	// The user login of the user who sent the gift. Set to null if it was an anonymous subscription gift.
	UserLogin string `json:"user_login"`
	// The user display name of the user who sent the gift. Set to null if it was an anonymous subscription gift.
	UserName string `json:"user_name"`
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The number of subscriptions in the subscription gift.
	Total int `json:"total"`
	// The tier of subscriptions in the subscription gift.
	Tier string `json:"tier"`
	// The number of subscriptions gifted by this user in the channel.
	// This value is null for anonymous gifts or if the gifter has opted out of sharing this information.
	CumulativeTotal int `json:"cumulative_total"`
	// Whether the subscription gift was anonymous.
	IsAnonymous bool `json:"is_anonymous"`
}

// ChannelUnbanRequestCreateEvent A user creates an unban request.
type ChannelUnbanRequestCreateEvent struct {
	// The ID of the unban request.
	Id string `json:"id"`
	// The broadcaster’s user ID for the channel the unban request was created for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster’s login name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster’s display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// User ID of user that is requesting to be unbanned.
	UserId string `json:"user_id"`
	// The user’s login name.
	UserLogin string `json:"user_login"`
	// The user’s display name.
	UserName string `json:"user_name"`
	// Message sent in the unban request.
	Text string `json:"text"`
	// The UTC timestamp (in RFC3339 format) of when the unban request was created.
	CreatedAt TimestampUTC `json:"created_at"`
}

// ChannelUnbanRequestResolveEvent An unban request has been resolved.
type ChannelUnbanRequestResolveEvent struct {
	// The ID of the unban request.
	Id string `json:"id"`
	// The broadcaster’s user ID for the channel the unban request was created for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster’s login name.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster’s display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// Optional. User ID of moderator who approved/denied the request.
	ModeratorUserId string `json:"moderator_user_id,omitempty"`
	// Optional. The moderator’s login name
	ModeratorUserLogin string `json:"moderator_user_login,omitempty"`
	// Optional. The moderator’s display name
	ModeratorUserName string `json:"moderator_user_name,omitempty"`
	// User ID of user that requested to be unbanned.
	UserId string `json:"user_id"`
	// 	The user’s login name.
	UserLogin string `json:"user_login"`
	// The user’s display name.
	UserName string `json:"user_name"`
	// Optional. Resolution text supplied by the mod/broadcaster upon approval/denial of the request.
	ResolutionText string `json:"resolution_text,omitempty"`
	// Dictates whether the unban request was approved or denied.
	// Can be the following: approved, canceled, denied
	Status ChannelUnbanRequestResolveEventStatus `json:"status"`
}

// EventUserUpdate A user has updated their account.
type UserUpdateEvent struct {
	// The user’s user id.
	UserId string `json:"user_id"`
	// The user’s user login.
	UserLogin string `json:"user_login"`
	// The user’s user display name.
	UserName string `json:"user_name"`
	// The user’s email. Only included if you have the user:read:email scope for the user.
	Email string `json:"email"`
	// The user’s description.
	Description string `json:"description"`
}

type ChannelVipAddEvent struct {
	// The user ID of the user that was added as a VIP.
	UserId string `json:"user_id"`
	// The user login of the user that was added as a VIP.
	UserLogin string `json:"user_login"`
	// The user display name of the user that was added as a VIP.
	UserName string `json:"user_name"`
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
}

type ChannelVipRemoveEvent struct {
	// The user ID of the user that was removed as a VIP.
	UserId string `json:"user_id"`
	// The user login of the user that was removed as a VIP.
	UserLogin string `json:"user_login"`
	// The user display name of the user that was removed as a VIP.
	UserName string `json:"user_name"`
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
}

type ChannelChatMessageDeleteEvent struct {
	// The broadcaster user ID.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// The broadcaster display name.
	BroadcasterUserName string `json:"broadcaster_user_name"`
	// The broadcaster login.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	// The ID of the user whose message was deleted.
	TargetUserId string `json:"target_user_id"`
	// The user name of the user whose message was deleted.
	TargetUserName string `json:"target_user_name"`
	// The user login of the user whose message was deleted.
	TargetUserLogin string `json:"target_user_login"`
	// A UUID that identifies the message that was removed.
	MessageId string `json:"message_id"`
}
