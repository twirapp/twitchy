package eventsub

type (
	onRevocation   = func(WebhookRevocationNotification)
	onVerification = func(WebhookCallbackVerificationNotification)
	onDuplicate    = func(MessageID)
)

// OnDuplicate invokes when duplicate message is caught.
func (wh *Webhook) OnDuplicate(onDuplicate onDuplicate) {
	wh.onDuplicate = onDuplicate
}

// OnRevocation invokes when webhook subscription revocation notification is caught (however, this message is processed
// automatically).
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#revoking-your-subscription.
func (wh *Webhook) OnRevocation(onRevocation onRevocation) {
	wh.onRevocation = onRevocation
}

// OnVerification invokes when webhook verification notification is caught (however, this message is processed
// automatically).
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#verifying-the-event-message.
func (wh *Webhook) OnVerification(onVerification onVerification) {
	wh.onVerification = onVerification
}
