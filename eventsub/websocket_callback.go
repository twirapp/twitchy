package eventsub

type (
	onWebsocketWelcome        = func(WebsocketWelcomeMessage)
	onWebsocketKeepalive      = func(WebsocketKeepaliveMessage)
	onWebsocketPing           = func()
	onWebsocketReconnect      = func(WebsocketReconnectMessage)
	onWebsocketReconnectError = func(error)
	onWebsocketDuplicate      = func(MessageID)
	onWebsocketDisconnect     = func()
)

// OnWelcome invokes when eventsub sends welcome message to let you subscribe to events.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#welcome-message.
func (ws *Websocket) OnWelcome(onWelcome onWebsocketWelcome) {
	ws.onWelcome = onWelcome
}

// OnKeepalive invokes when eventsub sends keepalive message which indicates that the websocket connection is healthy.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#keepalive-message.
func (ws *Websocket) OnKeepalive(onKeepalive onWebsocketKeepalive) {
	ws.onKeepalive = onKeepalive
}

// OnPing invokes when eventsub sends websocket ping frame (ping-pong cycle).
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#ping-message.
func (ws *Websocket) OnPing(onPing onWebsocketPing) {
	ws.onPing = onPing
}

// OnReconnect invokes when eventsub sends reconnect message which indicates that the edge server that the client is connected
// to needs to be swapped.
//
// Reference: https://dev.twitch.tv/docs/eventsub/handling-websocket-events/#reconnect-message.
func (ws *Websocket) OnReconnect(onReconnect onWebsocketReconnect) {
	ws.onReconnect = onReconnect
}

// OnReconnectError invokes when reconnection failed and error is returned.
func (ws *Websocket) OnReconnectError(onReconnectError onWebsocketReconnectError) {
	ws.onReconnectError = onReconnectError
}

// OnDuplicate invokes when duplicate message is caught.
func (ws *Websocket) OnDuplicate(onDuplicate onWebsocketDuplicate) {
	ws.onDuplicate = onDuplicate
}

// OnDisconnect invokes when client instance is being disconnected from eventsub server.
func (ws *Websocket) OnDisconnect(onWebsocketDisconnect onWebsocketDisconnect) {
	ws.onDisconnect = onWebsocketDisconnect
}
