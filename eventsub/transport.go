package eventsub

// Transport is a marker for transport.
type Transport interface {
	transport()
}

type WebhookTransport struct {
	Transport
	Method   string `json:"method"`
	Callback string `json:"callback"`
}

type ConduitTransport struct {
	Transport
	Method    string `json:"method"`
	ConduitId string `json:"conduit_id"`
}

type WebsocketTransport struct {
	Transport
	Method    string `json:"method"`
	SessionId string `json:"session_id"`
}
