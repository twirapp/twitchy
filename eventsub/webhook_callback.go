package eventsub

// callbackWebhook is a store for all EventSub webhook-specific callbacks.
type callbackWebhook[Metadata any] struct {
	onRevocation   func(WebhookRevocationNotification)
	onVerification func(WebhookCallbackVerificationNotification)
}

// OnRevocation invokes when webhook subscription revocation notification is caught (however, this message is processed
// automatically).
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#revoking-your-subscription.
func (cw *callbackWebhook[Metadata]) OnRevocation(onRevocation func(WebhookRevocationNotification)) {
	cw.onRevocation = onRevocation
}

// OnVerification invokes when webhook verification notification is caught (however, this message is processed
// automatically).
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#verifying-the-event-message.
func (cw *callbackWebhook[Metadata]) OnVerification(onVerification func(WebhookCallbackVerificationNotification)) {
	cw.onVerification = onVerification
}
