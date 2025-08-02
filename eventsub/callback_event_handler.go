package eventsub

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrUndefinedEventType = errors.New("undefined event type")

func (c *callback[Metadata]) runEventCallback(eventType, eventVersion string, body []byte, metadata Metadata) error {
	switch eventType {
	case "automod.message.hold":
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageHold, body, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageHoldV2, body, metadata)
		}
	case "automod.message.update":
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageUpdate, body, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageUpdateV2, body, metadata)
		}
	case "automod.settings.update":
		return runCallbackHandler(c.onAutomodSettingsUpdate, body, metadata)
	case "automod.terms.update":
		return runCallbackHandler(c.onAutomodTermsUpdate, body, metadata)
	case "channel.bits.use":
		return runCallbackHandler(c.onChannelBitsUse, body, metadata)
	case "channel.update":
		return runCallbackHandler(c.onChannelUpdate, body, metadata)
	case "channel.follow":
		return runCallbackHandler(c.onChannelFollow, body, metadata)
	case "channel.ad_break.begin":
		return runCallbackHandler(c.onChannelAdBreakBegin, body, metadata)
	case "channel.chat.clear":
		return runCallbackHandler(c.onChannelChatClear, body, metadata)
	case "channel.chat.clear_user_messages":
		return runCallbackHandler(c.onChannelChatClearUserMessages, body, metadata)
	case "channel.chat.message":
		return runCallbackHandler(c.onChannelChatMessage, body, metadata)
	case "conduit.shard.disabled":
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
