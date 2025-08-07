package eventsub

import (
	"errors"
	"fmt"

	"github.com/kvizyx/twitchy/internal/json"
)

// ErrUndefinedEventType indicates that eventsub server sent event that is not defined in library and user does not set
// callback handler for undefined events.
var ErrUndefinedEventType = errors.New("undefined event type")

// RawEvent is an event sent from eventsub server that contains payload in raw format.
type RawEvent struct {
	Subscription []byte
	Event        []byte
}

func (c *callback[Metadata]) runEventCallback(
	eventType EventType,
	eventVersion string,
	rawEvent RawEvent,
	metadata Metadata,
) error {
	event := rawEvent.Event

	switch eventType {
	case EventTypeAutomodMessageHold:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageHold, event, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageHoldV2, event, metadata)
		}
	case EventTypeAutomodMessageUpdate:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageUpdate, event, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageUpdateV2, event, metadata)
		}
	case EventTypeAutomodSettingsUpdate:
		return runCallbackHandler(c.onAutomodSettingsUpdate, event, metadata)
	case EventTypeAutomodTermsUpdate:
		return runCallbackHandler(c.onAutomodTermsUpdate, event, metadata)
	case EventTypeChannelBitsUse:
		return runCallbackHandler(c.onChannelBitsUse, event, metadata)
	case EventTypeChannelUpdate:
		return runCallbackHandler(c.onChannelUpdate, event, metadata)
	case EventTypeChannelFollow:
		return runCallbackHandler(c.onChannelFollow, event, metadata)
	case EventTypeChannelAdBreakBegin:
		return runCallbackHandler(c.onChannelAdBreakBegin, event, metadata)
	case EventTypeChannelChatClear:
		return runCallbackHandler(c.onChannelChatClear, event, metadata)
	case EventTypeChannelChatClearUserMessages:
		return runCallbackHandler(c.onChannelChatClearUserMessages, event, metadata)
	case EventTypeChannelChatMessage:
		return runCallbackHandler(c.onChannelChatMessage, event, metadata)
	case EventTypeConduitShardDisabled:
		return runCallbackHandler(c.onConduitShardDisabled, event, metadata)
	case EventTypeChannelBan:
		return runCallbackHandler(c.onChannelBan, event, metadata)
	case EventTypeChannelChatNotification:
		return runCallbackHandler(c.onChannelChatNotification, event, metadata)
	case EventTypeChannelModeratorAdd:
		return runCallbackHandler(c.onChannelModeratorAdd, event, metadata)
	case EventTypeChannelModeratorRemove:
		return runCallbackHandler(c.onChannelModeratorRemove, event, metadata)
	case EventTypeChannelPollBegin:
		return runCallbackHandler(c.onChannelPollBegin, event, metadata)
	case EventTypeChannelPollProgress:
		return runCallbackHandler(c.onChannelPollProgress, event, metadata)
	case EventTypeChannelPollEnd:
		return runCallbackHandler(c.onChannelPollEnd, event, metadata)
	case EventTypeChannelPredictionBegin:
		return runCallbackHandler(c.onChannelPredictionBegin, event, metadata)
	case EventTypeChannelPredictionProgress:
		return runCallbackHandler(c.onChannelPredictionProgress, event, metadata)
	case EventTypeChannelPredictionLock:
		return runCallbackHandler(c.onChannelPredictionLock, event, metadata)
	case EventTypeChannelPredictionEnd:
		return runCallbackHandler(c.onChannelPredictionEnd, event, metadata)
	case EventTypeChannelRaid:
		return runCallbackHandler(c.onChannelRaid, event, metadata)
	case EventTypeChannelPointsCustomRewardRedemptionAdd:
		return runCallbackHandler(c.onChannelPointsCustomRewardRedemptionAdd, event, metadata)
	case EventTypeChannelPointsCustomRewardRedemptionUpdate:
		return runCallbackHandler(c.onChannelPointsCustomRewardRedemptionUpdate, event, metadata)
	case EventTypeChannelPointsAutomaticRewardRedemptionAdd:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onChannelPointsAutomaticRewardRedemptionAdd, event, metadata)
		case "2":
			return runCallbackHandler(c.onChannelPointsAutomaticRewardRedemptionAddV2, event, metadata)
		}
	case EventTypeUserAuthorizationRevoke:
		return runCallbackHandler(c.onUserAuthorizationRevoke, event, metadata)
	case EventTypeChannelPointsRewardAdd:
		return runCallbackHandler(c.onChannelPointsCustomRewardAdd, event, metadata)
	case EventTypeChannelPointsRewardUpdate:
		return runCallbackHandler(c.onChannelPointsCustomRewardUpdate, event, metadata)
	case EventTypeChannelPointsRewardRemove:
		return runCallbackHandler(c.onChannelPointsCustomRewardRemove, event, metadata)
	case EventTypeStreamOffline:
		return runCallbackHandler(c.onStreamOffline, event, metadata)
	case EventTypeStreamOnline:
		return runCallbackHandler(c.onStreamOnline, event, metadata)
	case EventTypeChannelSubscribe:
		return runCallbackHandler(c.onChannelSubscribe, event, metadata)
	case EventTypeChannelSubscriptionEnd:
		return runCallbackHandler(c.onChannelSubscriptionEnd, event, metadata)
	case EventTypeChannelSubscriptionMessage:
		return runCallbackHandler(c.onChannelSubscriptionMessage, event, metadata)
	case EventTypeChannelSubscriptionGift:
		return runCallbackHandler(c.onChannelSubscriptionGift, event, metadata)
	case EventTypeChannelUnbanRequestCreate:
		return runCallbackHandler(c.onChannelUnbanRequestCreate, event, metadata)
	case EventTypeChannelUnbanRequestResolve:
		return runCallbackHandler(c.onChannelUnbanRequestResolve, event, metadata)
	case EventTypeUserUpdate:
		return runCallbackHandler(c.onUserUpdate, event, metadata)
	case EventTypeChannelVipAdd:
		return runCallbackHandler(c.onChannelVipAdd, event, metadata)
	case EventTypeChannelVipRemove:
		return runCallbackHandler(c.onChannelVipRemove, event, metadata)
	case EventTypeChannelMessageDelete:
		return runCallbackHandler(c.onChannelChatMessageDelete, event, metadata)
	default:
		if c.onUndefinedEvent != nil {
			go c.onUndefinedEvent(rawEvent, metadata)
		} else {
			return ErrUndefinedEventType
		}
	}

	return nil
}

// runCallbackHandler parses JSON event payload and runs handler in a separate go-routine with parsed payload if it's not nil.
func runCallbackHandler[Event, Metadata any](handler Handler[Event, Metadata], payload []byte, metadata Metadata) error {
	var event Event

	if handler != nil {
		if err := json.Unmarshal(payload, &event); err != nil {
			return fmt.Errorf("unmarshal event body: %w", err)
		}

		go handler(event, metadata)
	}

	return nil
}
