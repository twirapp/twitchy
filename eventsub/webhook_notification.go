package eventsub

type WebhookNotificationMetadata struct {
	MessageId           string
	MessageRetry        int
	MessageType         string
	MessageSignature    string
	MessageTimestamp    TimestampUTC
	SubscriptionType    EventType
	SubscriptionVersion string
}

type WebhookNotificationCondition struct {
	Condition
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type WebhookCallbackVerificationNotification struct {
	Challenge    string                                                       `json:"challenge"`
	Subscription Subscription[WebhookNotificationCondition, WebhookTransport] `json:"subscription"`
}

type WebhookRevocationNotification struct {
	Subscription Subscription[WebhookNotificationCondition, WebhookTransport] `json:"subscription"`
}
