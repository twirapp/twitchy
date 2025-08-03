package eventsub

// callbackWebhook is a store for all EventSub webhook-specific callbacks.
type callbackWebhook[Metadata any] struct {
	onRevocation   func(WebhookRevocationNotification)
	onVerification func(WebhookCallbackVerificationNotification)

	onUserAuthorizationRevoke Handler[UserAuthorizationRevokeEvent, Metadata]
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

// OnUserAuthorizationRevoke invokes when a user revokes authorization for an application.
//
// NOTE: This subscription type is only supported by webhooks, and cannot be used with WebSockets.
//
// Reference: https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#userauthorizationrevoke.
func (cw *callbackWebhook[Metadata]) OnUserAuthorizationRevoke(onUserAuthorizationRevoke Handler[UserAuthorizationRevokeEvent, Metadata]) {
	cw.onUserAuthorizationRevoke = onUserAuthorizationRevoke
}
