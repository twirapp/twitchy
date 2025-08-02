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
