package eventsub

type Subscription[C Condition, T Transport] struct {
	Id        string       `json:"id"`
	Status    string       `json:"status"`
	Type      string       `json:"type"`
	Version   string       `json:"version"`
	Cost      int          `json:"cost"`
	Condition C            `json:"condition"`
	Transport T            `json:"transport"`
	CreatedAt TimestampUTC `json:"created_at"`
}
