package eventsub

type WebsocketSession struct {
	Id                      string       `json:"id"`
	Status                  string       `json:"status"`
	KeepaliveTimeoutSeconds int          `json:"keepalive_timeout_seconds"`
	ReconnectURL            string       `json:"reconnect_url"`
	ConnectedAt             TimestampUTC `json:"connected_at"`
}

type WebsocketTransport struct {
	Method    string `json:"method"`
	SessionId string `json:"session_id"`
}

type WebsocketMessage[Metadata, Payload any] struct {
	Metadata Metadata `json:"metadata"`
	Payload  Payload  `json:"payload"`
}

type (
	WebsocketWelcomeMetadata struct {
		MessageId        string       `json:"message_id"`
		MessageType      string       `json:"message_type"`
		MessageTimestamp TimestampUTC `json:"message_timestamp"`
	}

	WebsocketWelcomePayload struct {
		Session WebsocketSession `json:"session"`
	}

	WebsocketWelcomeMessage WebsocketMessage[WebsocketWelcomeMetadata, WebsocketWelcomePayload]
)

type (
	WebsocketKeepaliveMetadata struct {
		MessageId        string       `json:"message_id"`
		MessageType      string       `json:"message_type"`
		MessageTimestamp TimestampUTC `json:"message_timestamp"`
	}

	WebsocketKeepalivePayload struct{}

	WebsocketKeepaliveMessage WebsocketMessage[WebsocketKeepaliveMetadata, WebsocketKeepalivePayload]
)

type (
	WebsocketNotificationMetadata struct {
		MessageId           string       `json:"message_id"`
		MessageType         string       `json:"message_type"`
		MessageTimestamp    TimestampUTC `json:"message_timestamp"`
		SubscriptionType    string       `json:"subscription_type"`
		SubscriptionVersion string       `json:"subscription_version"`
	}

	WebsocketNotificationPayload[Condition any] struct {
		Subscription Subscription[Condition, WebsocketTransport] `json:"subscription"`
	}

	WebsocketNotificationMessage[Event, Condition any] struct {
		WebsocketMessage[WebsocketNotificationMetadata, WebsocketNotificationPayload[Condition]]
		Event Event `json:"event"`
	}
)

type (
	WebsocketReconnectMetadata struct {
		MessageId        string       `json:"message_id"`
		MessageType      string       `json:"message_type"`
		MessageTimestamp TimestampUTC `json:"message_timestamp"`
	}

	WebsocketReconnectPayload struct {
		Session WebsocketSession `json:"session"`
	}

	WebsocketReconnectMessage WebsocketMessage[WebsocketReconnectMetadata, WebsocketReconnectPayload]
)

type (
	WebsocketRevocationMetadata struct {
		MessageId           string       `json:"message_id"`
		MessageType         string       `json:"message_type"`
		MessageTimestamp    TimestampUTC `json:"message_timestamp"`
		SubscriptionType    string       `json:"subscription_type"`
		SubscriptionVersion string       `json:"subscription_version"`
	}

	WebsocketRevocationPayload[Condition any] struct {
		Subscription Subscription[Condition, WebsocketTransport] `json:"subscription"`
	}

	WebsocketRevocationMessage[Condition any] struct {
		WebsocketMessage[WebsocketRevocationMetadata, WebsocketRevocationPayload[Condition]]
	}
)
