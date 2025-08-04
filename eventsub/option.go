package eventsub

type Option func(*EventSub)

func WithEventTracker(eventTracker EventTracker) Option {
	return func(eventSub *EventSub) {
		eventSub.eventTracker = eventTracker
	}
}
