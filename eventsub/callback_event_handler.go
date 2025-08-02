package eventsub

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrUndefinedEventType = errors.New("undefined event type")

func (c *callback[Metadata]) runEventCallback(eventType EventType, eventVersion string, body []byte, metadata Metadata) error {
	switch eventType {
	case EventTypeAutomodMessageHold:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageHold, body, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageHoldV2, body, metadata)
		}
	case EventTypeAutomodMessageUpdate:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageUpdate, body, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageUpdateV2, body, metadata)
		}
	case EventTypeAutomodSettingsUpdate:
		return runCallbackHandler(c.onAutomodSettingsUpdate, body, metadata)
	case EventTypeAutomodTermsUpdate:
		return runCallbackHandler(c.onAutomodTermsUpdate, body, metadata)
	case EventTypeChannelBitsUse:
		return runCallbackHandler(c.onChannelBitsUse, body, metadata)
	case EventTypeChannelUpdate:
		return runCallbackHandler(c.onChannelUpdate, body, metadata)
	case EventTypeChannelFollow:
		return runCallbackHandler(c.onChannelFollow, body, metadata)
	case EventTypeChannelAdBreakBegin:
		return runCallbackHandler(c.onChannelAdBreakBegin, body, metadata)
	case EventTypeChannelChatClear:
		return runCallbackHandler(c.onChannelChatClear, body, metadata)
	case EventTypeChannelChatClearUserMessages:
		return runCallbackHandler(c.onChannelChatClearUserMessages, body, metadata)
	case EventTypeChannelChatMessage:
		return runCallbackHandler(c.onChannelChatMessage, body, metadata)
	case EventTypeConduitShardDisabled:
		return runCallbackHandler(c.onConduitShardDisabled, body, metadata)
	default:
		return ErrUndefinedEventType
	}

	return nil
}

// runCallbackHandler parses payload from JSON body and runs handler in a separate go-routine with parsed payload if it's
// not nil.
func runCallbackHandler[Event, Metadata any](handler Handler[Event, Metadata], body []byte, metadata Metadata) error {
	var event Event

	if handler != nil {
		if err := json.Unmarshal(body, &event); err != nil {
			return fmt.Errorf("unmarshal event body: %w", err)
		}

		go handler(event, metadata)
	}

	return nil
}
