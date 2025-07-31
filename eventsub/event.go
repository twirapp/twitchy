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
