package eventsub

type AutomodSeverityLevel int

const (
	AutomodSeverityFirst AutomodSeverityLevel = iota + 1
	AutomodSeveritySecond
	AutomodSeverityThird
	AutomodSeverityFourth
)

type AutomodHoldReason string

const (
	AutomodHoldReasonAutomod     = "automod"
	AutomodHoldReasonBlockedTerm = "blocked_term"
)

type AutomodMessageStatus string

const (
	AutomodMessageStatusApproved AutomodMessageStatus = "approved"
	AutomodMessageStatusDenied   AutomodMessageStatus = "denied"
	AutomodMessageStatusExpired  AutomodMessageStatus = "expired"
)

type AutomodTermsAction string

const (
	AutomodTermsAddPermitted    AutomodTermsAction = "add_permitted"
	AutomodTermsRemovePermitted AutomodTermsAction = "remove_permitted"
	AutomodTermsAddBlocked      AutomodTermsAction = "add_blocked"
	AutomodTermsRemoveBlocked   AutomodTermsAction = "remove_blocked"
)

type FragmentType string

const (
	FragmentText      FragmentType = "text"
	FragmentEmote     FragmentType = "emote"
	FragmentCheermote FragmentType = "cheermote"
	FragmentMention   FragmentType = "mention"
)

type Mention struct {
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLogin string `json:"user_login"`
}

type (
	Fragment struct {
		Text      string     `json:"text"`
		Emote     *Emote     `json:"emote,omitempty"`
		Cheermote *Cheermote `json:"cheermote,omitempty"`
	}

	FragmentV2 struct {
		Type      FragmentType `json:"type"`
		Text      string       `json:"text"`
		Emote     *Emote       `json:"emote,omitempty"`
		Cheermote *Cheermote   `json:"cheermote,omitempty"`
	}

	FragmentV3 struct {
		Type      FragmentType `json:"type"`
		Text      string       `json:"text"`
		Emote     *EmoteV2     `json:"emote,omitempty"`
		Cheermote *Cheermote   `json:"cheermote,omitempty"`
	}

	FragmentV4 struct {
		Type      FragmentType `json:"type"`
		Text      string       `json:"text"`
		Emote     *EmoteV2     `json:"emote,omitempty"`
		Cheermote *Cheermote   `json:"cheermote,omitempty"`
		Mention   *Mention     `json:"mention"`
	}
)

type EmoteFormat string

const (
	EmoteFormatAnimated EmoteFormat = "animated"
	EmoteFormatStatic   EmoteFormat = "static"
)

type (
	Emote struct {
		Id         string `json:"id"`
		EmoteSetId string `json:"emote_set_id"`
	}

	EmoteV2 struct {
		Id         string        `json:"id"`
		EmoteSetId string        `json:"emote_set_id"`
		OwnerId    string        `json:"owner_id"`
		Format     []EmoteFormat `json:"format"`
	}
)

type RewardEmote struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cheermote struct {
	Prefix string `json:"prefix"`
	Bits   int    `json:"bits"`
	Tier   int    `json:"tier"`
}

type (
	Message struct {
		Text      string     `json:"text"`
		Fragments []Fragment `json:"fragments"`
	}

	MessageV2 struct {
		Text      string       `json:"text"`
		Fragments []FragmentV2 `json:"fragments"`
	}

	MessageV3 struct {
		Text      string       `json:"text"`
		Fragments []FragmentV3 `json:"fragments"`
	}

	MessageV4 struct {
		Text      string       `json:"text"`
		Fragments []FragmentV4 `json:"fragments"`
	}
)

type MessageEmote struct {
	// The emote ID.
	Id string `json:"id"`
	// The index of where the Emote starts in the text.
	Begin int `json:"begin"`
	// The index of where the Emote ends in the text.
	End int `json:"end"`
}

type Automod struct {
	Category   string               `json:"category"`
	Level      AutomodSeverityLevel `json:"level"`
	Boundaries []AutomodBoundary    `json:"boundaries"`
}

type AutomodBoundary struct {
	StartPosition int `json:"start_position"`
	EndPosition   int `json:"end_position"`
}

type Term struct {
	TermId                    string          `json:"term_id"`
	Boundary                  AutomodBoundary `json:"boundary"`
	OwnerBroadcasterUserId    string          `json:"owner_broadcaster_user_id"`
	OwnerBroadcasterUserLogin string          `json:"owner_broadcaster_user_login"`
	OwnerBroadcasterUserName  string          `json:"owner_broadcaster_user_name"`
}

type BlockedTerm struct {
	TermsFound []Term `json:"terms_found"`
}

type BitsUseType string

const (
	BitsUseCheer   BitsUseType = "cheer"
	BitsUsePowerUp BitsUseType = "power_up"
)

type BitsPowerUpType string

const (
	BitsPowerUpMessageEffect    BitsPowerUpType = "message_effect"
	BitsPowerUpCelebration      BitsPowerUpType = "celebration"
	BitsPowerUpGigantifyAnEmote BitsPowerUpType = "gigantify_an_emote"
)

type PowerUp struct {
	Type            BitsPowerUpType `json:"type"`
	Emote           *RewardEmote    `json:"emote,omitempty"`
	MessageEffectId string          `json:"message_effect_id,omitempty"`
}

type MessageType string

const (
	MessageText                     MessageType = "text"
	MessageChannelPointsHighlighted MessageType = "channel_points_highlighted"
	MessageChannelPointsSubOnly     MessageType = "channel_points_sub_only"
	MessageUserIntro                MessageType = "user_intro"
	MessagePowerUpsMessageEffect    MessageType = "power_ups_message_effect"
	MessagePowerUpsGigantifiedEmote MessageType = "power_ups_gigantified_emote"
)

type Badge struct {
	Id    string `json:"id"`
	SetId string `json:"set_id"`
	Info  string `json:"info"`
}

type Cheer struct {
	Bits int `json:"bits"`
}

type Reply struct {
	ParentMessageId   string `json:"parent_message_id"`
	ParentMessageBody string `json:"parent_message_body"`
	ParentUserId      string `json:"parent_user_id"`
	ParentUserName    string `json:"parent_user_name"`
	ParentUserLogin   string `json:"parent_user_login"`
	ThreadMessageId   string `json:"thread_message_id"`
	ThreadUserId      string `json:"thread_user_id"`
	ThreadUserName    string `json:"thread_user_name"`
	ThreadUserLogin   string `json:"thread_user_login"`
}

type ConduitShardDisabledEventTransport struct {
	// Method websocket or webhook
	Method string `json:"method"`
	// CallBack Optional. Webhook callback URL. Empty if method is set to websocket.
	CallBack string `json:"callback"`
	// SessionId Optional. WebSocket session ID. Empty if  method is set to webhook.
	SessionId string `json:"session_id"`
	// ConnectedAt Optional. Time that the WebSocket session connected. Empty if method is set to webhook.
	ConnectedAt TimestampUTC `json:"connected_at"`
	// DisconnectAt Optional. Time that the WebSocket session disconnected. Empty if method is set to webhook.
	DisconnectAt TimestampUTC `json:"disconnected_at"`
}

// ChatNotificationEventMessage represents a chat notification event.
type ChatNotificationEventMessage struct {
	// The chat message in plain text.
	Text string `json:"text"`
	// Ordered list of chat message fragments.
	Fragments []ChatNotificationEventMessageFragment `json:"fragments"`
}

// ChatNotificationEventMessageFragment represents metadata pertaining to the cheermote.
type ChatNotificationEventMessageFragment struct {
	// The type of message fragment. Possible values:
	//  - text
	//  - cheermote
	//  - emote
	//  - mention
	Type string `json:"type"`
	// Message text in fragment
	Text string `json:"text"`
	// Optional. Metadata pertaining to the cheermote.
	Cheermote *ChatNotificationEventMessageFragmentCheermote `json:"cheermote,omitempty"`
	// Optional. Metadata pertaining to the emote.
	Emote *ChatNotificationEventMessageFragmentEmote `json:"emote,omitempty"`
	// Optional. Metadata pertaining to the mention.
	Mention *ChatNotificationEventMessageFragmentMention `json:"mention,omitempty"`
}

// ChatNotificationEventMessageFragmentCheermote represents metadata pertaining to the emote.
type ChatNotificationEventMessageFragmentCheermote struct {
	// The name portion of the ChatNotificationMessageFragmentCheermote string that you use in chat to cheer Bits. The full ChatNotificationMessageFragmentCheermote string is the concatenation of {prefix} + {number of Bits}. For example, if the prefix is “Cheer” and you want to cheer 100 Bits, the full ChatNotificationMessageFragmentCheermote string is Cheer100. When the ChatNotificationMessageFragmentCheermote string is entered in chat, Twitch converts it to the image associated with the Bits tier that was cheered.
	Prefix string `json:"prefix"`
	// The amount of bits cheered.
	Bits int `json:"bits"`
	// The tier level of the cheermote.
	Tier int `json:"tier"`
}

// ChatNotificationEventMessageFragmentEmote represents metadata pertaining to an emote.
type ChatNotificationEventMessageFragmentEmote struct {
	// An ID that uniquely identifies this emote.
	Id string `json:"id"`
	// An ID that identifies the emote set that the emote belongs to.
	EmoteSetId string `json:"emote_set_id"`
	// The ID of the broadcaster who owns the emote.
	OwnerId string `json:"owner_id"`
	// The formats that the emote is available in. For example, if the emote is available only as a static PNG, the array contains only static. But if the emote is available as a static PNG and an animated GIF, the array contains static and animated. The possible formats are:
	//  - animated — An animated GIF is available for this emote.
	//  - static — A static PNG file is available for this emote.
	Format []string `json:"format"`
}

// ChatNotificationEventMessageFragmentMention represents metadata pertaining to a mention.
type ChatNotificationEventMessageFragmentMention struct {
	// The user ID of the mentioned user.
	UserId string `json:"user_id"`
	// The user name of the mentioned user.
	UserName string `json:"user_name"`
	// The user login of the mentioned user.
	UserLogin string `json:"user_login"`
}

// ChatNotificationEventBadge represents a chat badge.
type ChatNotificationEventBadge struct {
	// An ID that identifies this set of chat badges. For example, Bits or Subscriber.
	SetId string `json:"set_id"`
	// An ID that identifies this version of the badge. The ID can be any value. For example, for Bits, the ID is the Bits tier level, but for World of Warcraft, it could be Alliance or Horde.
	Id string `json:"id"`
	// Contains metadata related to the chat badges in the badges tag. Currently, this tag contains metadata only for subscriber badges, to indicate the number of months the user has been a subscriber.
	Info string `json:"info"`
}

// ChatNotificationEventSubEvent represents information about the sub event.
type ChatNotificationEventSubEvent struct {
	// The type of subscription plan being used. Possible values are:
	//  - 1000 — First level of paid or Prime subscription
	//  - 2000 — Second level of paid subscription
	//  - 3000 — Third level of paid subscription
	SubTier string `json:"sub_tier"`

	// Indicates if the subscription was obtained through Amazon Prime.
	IsPrime bool `json:"is_prime"`

	// The number of months the subscription is for.
	DurationMonths int `json:"duration_months"`
}

// ChatNotificationEventResubEvent represents information about the resub event.
type ChatNotificationEventResubEvent struct {
	// The total number of months the user has subscribed.
	CumulativeMonths int `json:"cumulative_months"`

	// The number of months the subscription is for.
	DurationMonths int `json:"duration_months"`

	// Optional. The number of consecutive months the user has subscribed.
	StreakMonths int `json:"streak_months,omitempty"`

	// The type of subscription plan being used. Possible values are:
	//  - 1000 — First level of paid or Prime subscription
	//  - 2000 — Second level of paid subscription
	//  - 3000 — Third level of paid subscription
	SubTier string `json:"sub_tier"`

	// Indicates if the resub was obtained through Amazon Prime.
	IsPrime bool `json:"is_prime"`

	// Whether or not the resub was a result of a gift.
	IsGift bool `json:"is_gift"`

	// Optional. Whether or not the gift was anonymous.
	GifterIsAnonymous bool `json:"gifter_is_anonymous,omitempty"`

	// Optional. The user ID of the subscription gifter. Null if anonymous.
	GifterUserId string `json:"gifter_user_id,omitempty"`

	// Optional. The user name of the subscription gifter. Null if anonymous.
	GifterUserName string `json:"gifter_user_name,omitempty"`

	// Optional. The user login of the subscription gifter. Null if anonymous.
	GifterUserLogin string `json:"gifter_user_login,omitempty"`
}

// ChatNotificationEventSubGiftEvent represents information about the gift sub event.
type ChatNotificationEventSubGiftEvent struct {
	// The number of months the subscription is for.
	DurationMonths int `json:"duration_months"`

	// Optional. The amount of gifts the gifter has given in this channel. Null if anonymous.
	CumulativeTotal int `json:"cumulative_total,omitempty"`

	// The user ID of the subscription gift recipient.
	RecipientUserId string `json:"recipient_user_id"`

	// The user name of the subscription gift recipient.
	RecipientUserName string `json:"recipient_user_name"`

	// The user login of the subscription gift recipient.
	RecipientUserLogin string `json:"recipient_user_login"`

	// The type of subscription plan being used. Possible values are:
	//  - 1000 — First level of paid subscription
	//  - 2000 — Second level of paid subscription
	//  - 3000 — Third level of paid subscription
	SubTier string `json:"sub_tier"`

	// Optional. The ID of the associated community gift. Null if not associated with a community gift.
	CommunityGiftId string `json:"community_gift_id,omitempty"`
}

// ChatNotificationEventCommunitySubGiftEvent represents information about the community gift sub event.
type ChatNotificationEventCommunitySubGiftEvent struct {
	// The ID of the associated community gift.
	Id string `json:"id"`

	// Number of subscriptions being gifted.
	Total int `json:"total"`

	// The type of subscription plan being used. Possible values are:
	//  - 1000 — First level of paid subscription
	//  - 2000 — Second level of paid subscription
	//  - 3000 — Third level of paid subscription
	SubTier string `json:"sub_tier"`

	// Optional. The amount of gifts the gifter has given in this channel. Null if anonymous.
	CumulativeTotal int `json:"cumulative_total,omitempty"`
}

// ChatNotificationEventGiftPaidUpgradeEvent represents information about the Prime gift paid upgrade event.
type ChatNotificationEventGiftPaidUpgradeEvent struct {
	// Whether the gift was given anonymously.
	GifterIsAnonymous bool `json:"gifter_is_anonymous"`

	// Optional. The user ID of the user who gifted the subscription. Null if anonymous.
	GifterUserId string `json:"gifter_user_id,omitempty"`

	// Optional. The user name of the user who gifted the subscription. Null if anonymous.
	GifterUserName string `json:"gifter_user_name,omitempty"`

	// Optional. The user login of the user who gifted the subscription. Null if anonymous.
	GifterUserLogin string `json:"gifter_user_login,omitempty"`
}

// ChatNotificationEventRaidEvent represents information about the raid event.
type ChatNotificationEventRaidEvent struct {
	// The user ID of the broadcaster raiding this channel.
	UserId string `json:"user_id"`

	// The user name of the broadcaster raiding this channel.
	UserName string `json:"user_name"`

	// The login name of the broadcaster raiding this channel.
	UserLogin string `json:"user_login"`

	// The number of viewers raiding this channel from the broadcaster’s channel.
	ViewerCount int `json:"viewer_count"`

	// Profile image URL of the broadcaster raiding this channel.
	ProfileImageURL string `json:"profile_image_url"`
}

// ChatNotificationEventUnraidEvent represents an empty payload for the unraid event.
type ChatNotificationEventUnraidEvent struct{}

// ChatNotificationEventPayItForwardEvent represents information about the pay it forward event.
type ChatNotificationEventPayItForwardEvent struct {
	// Whether the gift was given anonymously.
	GifterIsAnonymous bool `json:"gifter_is_anonymous"`

	// Optional. The user ID of the user who gifted the subscription. Null if anonymous.
	GifterUserId string `json:"gifter_user_id,omitempty"`

	// Optional. The user name of the user who gifted the subscription. Null if anonymous.
	GifterUserName string `json:"gifter_user_name,omitempty"`

	// Optional. The user login of the user who gifted the subscription. Null if anonymous.
	GifterUserLogin string `json:"gifter_user_login,omitempty"`
}

// ChatNotificationEventAnnouncementEvent represents information about the announcement event.
type ChatNotificationEventAnnouncementEvent struct {
	// Color of the announcement.
	Color string `json:"color"`
}

// ChatNotificationEventCharityDonationEvent represents information about the charity donation event.
type ChatNotificationEventCharityDonationEvent struct {
	// Name of the charity.
	CharityName string `json:"charity_name"`

	// An object that contains the amount of money that the user paid.
	Amount ChatNotificationEventCharityDonationEventDonationAmount `json:"amount"`
}

// ChatNotificationEventCharityDonationEventDonationAmount represents the amount of money that the user paid.
type ChatNotificationEventCharityDonationEventDonationAmount struct {
	// The monetary amount. The amount is specified in the currency’s minor unit.
	// For example, the minor units for USD is cents, so if the amount is $5.50 USD, value is set to 550.
	Value int `json:"value"`

	// The number of decimal places used by the currency.
	// For example, USD uses two decimal places.
	DecimalPlaces int `json:"decimal_places"`

	// The ISO-4217 three-letter currency code that identifies the type of currency in value.
	Currency string `json:"currency"`
}

// ChatNotificationEventBitsBadgeTierEvent represents information about the bits badge tier event.
type ChatNotificationEventBitsBadgeTierEvent struct {
	// The tier of the Bits badge the user just earned. For example, 100, 1000, or 10000.
	Tier int `json:"tier"`
}

type ChatNotificationEventNoticeType string

const (
	ChatNotificationEventNoticeTypeSub                        ChatNotificationEventNoticeType = "sub"
	ChatNotificationEventNoticeTypeResub                      ChatNotificationEventNoticeType = "resub"
	ChatNotificationEventNoticeTypeSubGift                    ChatNotificationEventNoticeType = "sub_gift"
	ChatNotificationEventNoticeTypeCommunitySubGift           ChatNotificationEventNoticeType = "community_sub_gift"
	ChatNotificationEventNoticeTypeGiftPaidUpgrade            ChatNotificationEventNoticeType = "gift_paid_upgrade"
	ChatNotificationEventNoticeTypePrimePaidUpgrade           ChatNotificationEventNoticeType = "prime_paid_upgrade"
	ChatNotificationEventNoticeTypeRaid                       ChatNotificationEventNoticeType = "raid"
	ChatNotificationEventNoticeTypeUnraid                     ChatNotificationEventNoticeType = "unraid"
	ChatNotificationEventNoticeTypePayItForward               ChatNotificationEventNoticeType = "pay_it_forward"
	ChatNotificationEventNoticeTypeAnnouncement               ChatNotificationEventNoticeType = "announcement"
	ChatNotificationEventNoticeTypeBitsBadgeTier              ChatNotificationEventNoticeType = "bits_badge_tier"
	ChatNotificationEventNoticeTypeCharityDonation            ChatNotificationEventNoticeType = "charity_donation"
	ChatNotificationEventNoticeTypeSharedChatSub              ChatNotificationEventNoticeType = "shared_chat_sub"
	ChatNotificationEventNoticeTypeSharedChatResub            ChatNotificationEventNoticeType = "shared_chat_resub"
	ChatNotificationEventNoticeTypeSharedChatSubGift          ChatNotificationEventNoticeType = "shared_chat_sub_gift"
	ChatNotificationEventNoticeTypeSharedChatCommunitySubGift ChatNotificationEventNoticeType = "shared_chat_community_sub_gift"
	ChatNotificationEventNoticeTypeSharedChatGiftPaidUpgrade  ChatNotificationEventNoticeType = "shared_chat_gift_paid_upgrade"
	ChatNotificationEventNoticeTypeSharedChatPrimePaidUpgrade ChatNotificationEventNoticeType = "shared_chat_prime_paid_upgrade"
	ChatNotificationEventNoticeTypeSharedChatRaid             ChatNotificationEventNoticeType = "shared_chat_raid"
	ChatNotificationEventNoticeTypeSharedChatPayItForward     ChatNotificationEventNoticeType = "shared_chat_pay_it_forward"
	ChatNotificationEventNoticeTypeSharedChatAnnouncement     ChatNotificationEventNoticeType = "shared_chat_announcement"
)

func (c ChatNotificationEventNoticeType) String() string {
	return string(c)
}

type ChannelPollEventChoice struct {
	// ID for the choice.
	Id string `json:"id"`
	// Text displayed for the choice.
	Title string `json:"title"`
	// Number of votes received via Bits.
	BitsVotes int `json:"bits_votes"`
	// Number of votes received via Channel Points.
	ChannelPointsVotes int `json:"channel_points_votes"`
	// Total number of votes received for the choice across all methods of voting.
	Votes int `json:"votes"`
}

type ChannelPollEventBitsVoting struct {
	// Indicates if Bits can be used for voting.
	IsEnabled bool `json:"is_enabled"`
	// Number of Bits required to vote once with Bits.
	AmountPerVote int `json:"amount_per_vote"`
}

type ChannelPollEventChannelPointsVoting struct {
	// Indicates if Channel Points can be used for voting.
	IsEnabled bool `json:"is_enabled"`
	// Number of Channel Points required to vote once with Channel Points.
	AmountPerVote int `json:"amount_per_vote"`
}

type ChannelPollEndEventStatus string

const (
	ChannelPollEndEventStatusCompleted  ChannelPollEndEventStatus = "completed"
	ChannelPollEndEventStatusArchived   ChannelPollEndEventStatus = "archived"
	ChannelPollEndEventStatusTerminated ChannelPollEndEventStatus = "terminated"
)

func (c ChannelPollEndEventStatus) String() string {
	return string(c)
}

type ChannelPredictionEventOutcome struct {
	// The outcome ID.
	Id string `json:"id"`
	// The outcome title.
	Title string `json:"title"`
	// The color for the outcome. Valid values are pink and blue.
	Color string `json:"color"`
	// The number of users who used Channel Points on this outcome.
	Users int `json:"users"`
	// The total number of Channel Points used on this outcome.
	ChannelPoints int `json:"channel_points"`
	// An array of users who used the most Channel Points on this outcome.
	TopPredictors []ChannelPredictionEventOutcomeTopPredictor `json:"top_predictors"`
}

type ChannelPredictionEventOutcomeTopPredictor struct {
	// The ID of the user.
	UserId string `json:"user_id"`
	// The login of the user.
	UserLogin string `json:"user_login"`
	// The display name of the user.
	UserName string `json:"user_name"`
	// The number of Channel Points won.
	// This value is always null in the event payload for Prediction progress and Prediction lock.
	// This value is 0 if the outcome did not win or if the Prediction was canceled and Channel Points were refunded.
	ChannelPointsWon int `json:"channel_points_won"`
	// The number of Channel Points used to participate in the Prediction.
	ChannelPointsUsed int `json:"channel_points_used"`
}

type ChannelPredictionEndEventStatus string

const (
	ChannelPredictionEndStatusResolved ChannelPredictionEndEventStatus = "resolved"
	ChannelPredictionEndStatusCanceled ChannelPredictionEndEventStatus = "canceled"
)

type ChannelPointsCustomEventReward struct {
	// The reward identifier.
	Id string `json:"id"`
	// The reward name.
	Title string `json:"title"`
	// The reward cost.
	Cost int `json:"cost"`
	// The reward description.
	Prompt string `json:"prompt"`
}

type ChannelPointsAutomaticRewardEventRewardType string

const (
	ChannelPointsAutomaticRewardEventRewardTypeSingleMessageBypassSubMode   ChannelPointsAutomaticRewardEventRewardType = "single_message_bypass_sub_mode"
	ChannelPointsAutomaticRewardEventRewardTypeSendHighlightedMessage       ChannelPointsAutomaticRewardEventRewardType = "send_highlighted_message"
	ChannelPointsAutomaticRewardEventRewardTypeRandomSubEmoteUnlock         ChannelPointsAutomaticRewardEventRewardType = "random_sub_emote_unlock"
	ChannelPointsAutomaticRewardEventRewardTypeChosenSubEmoteUnlock         ChannelPointsAutomaticRewardEventRewardType = "chosen_sub_emote_unlock"
	ChannelPointsAutomaticRewardEventRewardTypeChosenModifiedSubEmoteUnlock ChannelPointsAutomaticRewardEventRewardType = "chosen_modified_sub_emote_unlock"
	ChannelPointsAutomaticRewardEventRewardTypeMessageEffect                ChannelPointsAutomaticRewardEventRewardType = "message_effect"
	ChannelPointsAutomaticRewardEventRewardTypeGigantifyAnEmote             ChannelPointsAutomaticRewardEventRewardType = "gigantify_an_emote"
	ChannelPointsAutomaticRewardEventRewardTypeCelebration                  ChannelPointsAutomaticRewardEventRewardType = "celebration"
)

type ChannelPointsAutomaticRewardEventReward struct {
	// The type of reward. One of:
	// - single_message_bypass_sub_mode
	// - send_highlighted_message
	// - random_sub_emote_unlock
	// - chosen_sub_emote_unlock
	// - chosen_modified_sub_emote_unlock
	// - message_effect
	// - gigantify_an_emote
	// - celebration
	Type ChannelPointsAutomaticRewardEventRewardType `json:"type"`
	// The reward cost.
	Cost int
	// Optional. Emote that was unlocked.
	UnlockedEmote *ChannelPointsAutomaticRewardEventRewardUnlockedEmote `json:"unlocked_emote,omitempty"`
}

type ChannelPointsAutomaticRewardEventRewardUnlockedEmote struct {
	// The emote ID.
	Id string `json:"id"`
	// The human readable emote token.
	Name string `json:"name"`
}

type ChannelPointsAutomaticRewardEventRewardMessage struct {
	// The text of the chat message.
	Text string `json:"text"`
	// An array that includes the emote ID and start and end positions for where the emote appears in the text.
	Emotes []MessageEmote `json:"emotes,omitempty"`
}

type ChannelPointsAutomaticRewardEventRewardTypeV2 string

const (
	ChannelPointsAutomaticRewardEventRewardTypeV2SingleMessageBypassSubMode   ChannelPointsAutomaticRewardEventRewardTypeV2 = "single_message_bypass_sub_mode"
	ChannelPointsAutomaticRewardEventRewardTypeV2SendHighlightedMessage       ChannelPointsAutomaticRewardEventRewardTypeV2 = "send_highlighted_message"
	ChannelPointsAutomaticRewardEventRewardTypeV2RandomSubEmoteUnlock         ChannelPointsAutomaticRewardEventRewardTypeV2 = "random_sub_emote_unlock"
	ChannelPointsAutomaticRewardEventRewardTypeV2ChosenSubEmoteUnlock         ChannelPointsAutomaticRewardEventRewardTypeV2 = "chosen_sub_emote_unlock"
	ChannelPointsAutomaticRewardEventRewardTypeV2ChosenModifiedSubEmoteUnlock ChannelPointsAutomaticRewardEventRewardTypeV2 = "chosen_modified_sub_emote_unlock"
)

func (c ChannelPointsAutomaticRewardEventRewardTypeV2) String() string {
	return string(c)
}

type ChannelPointsAutomaticRewardEventRewardV2 struct {
	// The type of reward. One of:
	// - single_message_bypass_sub_mode
	// - send_highlighted_message
	// - random_sub_emote_unlock
	// - chosen_sub_emote_unlock
	// - chosen_modified_sub_emote_unlock
	Type ChannelPointsAutomaticRewardEventRewardTypeV2 `json:"type"`
	// Number of channel points used.
	ChannelPoints int `json:"channel_points"`
	// Optional. Emote associated with the reward.
	Emote *ChannelPointsAutomaticRewardEventRewardV2Emote `json:"emote""`
}

type ChannelPointsAutomaticRewardEventRewardV2Emote struct {
	// The emote ID.
	Id string `json:"id"`
	// The human readable emote token.
	Name string `json:"name"`
}

type ChannelPointsAutomaticRewardEventRewardMessageV2 struct {
	// The chat message in plain text.
	Text string `json:"text"`
	// The ordered list of chat message fragments.
	Fragments []ChannelPointsAutomaticRewardEventRewardMessageFragmentV2 `json:"fragments,omitempty"`
}

type ChannelPointsAutomaticRewardEventRewardMessageFragmentV2Type string

const (
	ChannelPointsAutomaticRewardEventRewardMessageFragmentV2TypeText  ChannelPointsAutomaticRewardEventRewardMessageFragmentV2Type = "text"
	ChannelPointsAutomaticRewardEventRewardMessageFragmentV2TypeEmote ChannelPointsAutomaticRewardEventRewardMessageFragmentV2Type = "emote"
)

type ChannelPointsAutomaticRewardEventRewardMessageFragmentV2 struct {
	// The message text in fragment.
	Text string `json:"text"`
	// The type of message fragment. Possible values are:
	// - text
	// - emote
	Type ChannelPointsAutomaticRewardEventRewardMessageFragmentV2Type `json:"type"`
	// Optional. The metadata pertaining to the emote.
	Emote *ChannelPointsAutomaticRewardEventRewardMessageFragmentV2Emote `json:"emote,omitempty"`
}

type ChannelPointsAutomaticRewardEventRewardMessageFragmentV2Emote struct {
	// The ID that uniquely identifies this emote.
	Id string `json:"id"`
}

type ChannelPointsRewardEventMaxPerStream struct {
	// Is the setting enabled.
	IsEnabled bool `json:"is_enabled"`
	// The max per stream limit.
	Value int `json:"value"`
}

type ChannelPointsRewardEventMaxPerUserPerStream struct {
	// Is the setting enabled.
	IsEnabled bool `json:"is_enabled"`
	// The max per user per stream limit.
	Value int `json:"value"`
}

type ChannelPointsRewardEventGlobalCooldown struct {
	// Is the setting enabled.
	IsEnabled bool `json:"is_enabled"`
	// The global cooldown in seconds.
	Seconds int `json:"seconds"` // The global cooldown in seconds.
}

type ChannelPointsRewardEventImage struct {
	// URL for the image at 1x size.
	Url1x string `json:"url_1x"`
	// URL for the image at 2x size.
	Url2x string `json:"url_2x"`
	// URL for the image at 4x size.
	Url4x string `json:"url_4x"`
}

type ChannelSubscriptionMessageEventMessage struct {
	// The text of the resubscription chat message.
	Text string `json:"text"`
	// An array that includes the emote ID and start and end positions for where the emote appears in the text.
	Emotes []MessageEmote `json:"emotes,omitempty"`
}

type ChannelUnbanRequestResolveEventStatus string

const (
	ChannelUnbanRequestResolveEventStatusApproved ChannelUnbanRequestResolveEventStatus = "approved"
	ChannelUnbanRequestResolveEventStatusCanceled ChannelUnbanRequestResolveEventStatus = "canceled"
	ChannelUnbanRequestResolveEventStatusDenied   ChannelUnbanRequestResolveEventStatus = "denied"
)

func (c ChannelUnbanRequestResolveEventStatus) String() string {
	return string(c)
}

type StreamOfflineEventType string

const (
	StreamOfflineEventTypeLive       StreamOfflineEventType = "live"
	StreamOfflineEventTypePlaylist   StreamOfflineEventType = "playlist"
	StreamOfflineEventTypeWatchParty StreamOfflineEventType = "watch_party"
	StreamOfflineEventTypePremiere   StreamOfflineEventType = "premiere"
	StreamOfflineEventTypeRerun      StreamOfflineEventType = "rerun"
)

func (c StreamOfflineEventType) String() string {
	return string(c)
}
