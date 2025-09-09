package eventsub

// OnRevocation invokes when webhook subscription revocation notification is caught.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#revoking-your-subscription.
func (wh *Webhook) OnRevocation(onRevocation func(WebhookRevocationNotification)) {
	wh.onRevocation = onRevocation
}

// OnVerification invokes when webhook verification notification is caught.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#verifying-the-event-message.
func (wh *Webhook) OnVerification(onVerification func(WebhookCallbackVerificationNotification)) {
	wh.onVerification = onVerification
}
