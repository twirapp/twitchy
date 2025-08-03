package eventsub

// callbackWebsocket is a store for all EventSub websocket-specific callbacks.
type callbackWebsocket struct {
	onWelcome        func(WebsocketWelcomeMessage)
	onKeepalive      func(WebsocketKeepaliveMessage)
	onPing           func()
	onReconnect      func(WebsocketReconnectMessage)
	onReconnectError func(error)
	onDisconnect     func()
}

// OnWelcome invokes when eventsub sends welcome message to let you subscribe to events.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#welcome-message.
func (cw *callbackWebsocket) OnWelcome(onWelcome func(WebsocketWelcomeMessage)) {
	cw.onWelcome = onWelcome
}

// OnKeepalive invokes when eventsub sends keepalive message which indicates that the websocket connection is healthy.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#keepalive-message.
func (cw *callbackWebsocket) OnKeepalive(onKeepalive func(WebsocketKeepaliveMessage)) {
	cw.onKeepalive = onKeepalive
}

// OnPing invokes when eventsub sends websocket ping frame (ping-pong cycle).
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#ping-message.
func (cw *callbackWebsocket) OnPing(onPing func()) {
	cw.onPing = onPing
}

// OnReconnect invokes when eventsub sends reconnect message which indicates that the edge server that the client is connected
// to needs to be swapped.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#reconnect-message.
func (cw *callbackWebsocket) OnReconnect(onReconnect func(WebsocketReconnectMessage)) {
	cw.onReconnect = onReconnect
}

// OnReconnectError invokes when reconnection failed and error is returned.
func (cw *callbackWebsocket) OnReconnectError(onReconnectError func(error)) {
	cw.onReconnectError = onReconnectError
}

// OnDisconnect invokes when client instance is being disconnected from eventsub server.
func (cw *callbackWebsocket) OnDisconnect(onDisconnect func()) {
	cw.onDisconnect = onDisconnect
}
