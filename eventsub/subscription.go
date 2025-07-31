package eventsub

type Subscription[Condition, Transport any] struct {
	Id        string       `json:"id"`
	Status    string       `json:"status"`
	Type      string       `json:"type"`
	Version   string       `json:"version"`
	Cost      int          `json:"cost"`
	Condition Condition    `json:"condition"`
	Transport Transport    `json:"transport"`
	CreatedAt TimestampUTC `json:"created_at"`
}
