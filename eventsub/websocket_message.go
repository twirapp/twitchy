package eventsub

type WebsocketSession struct {
	Id                      string       `json:"id"`
	Status                  string       `json:"status"`
	KeepaliveTimeoutSeconds int          `json:"keepalive_timeout_seconds"`
	ReconnectURL            string       `json:"reconnect_url"`
	ConnectedAt             TimestampUTC `json:"connected_at"`
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

	WebsocketNotificationPayload[C Condition] struct {
		Subscription Subscription[C, WebsocketTransport] `json:"subscription"`
	}

	WebsocketNotificationMessage[Event any, C Condition] struct {
		WebsocketMessage[WebsocketNotificationMetadata, WebsocketNotificationPayload[C]]
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

	WebsocketRevocationPayload[C Condition] struct {
		Subscription Subscription[C, WebsocketTransport] `json:"subscription"`
	}

	WebsocketRevocationMessage[C Condition] struct {
		WebsocketMessage[WebsocketRevocationMetadata, WebsocketRevocationPayload[C]]
	}
)
