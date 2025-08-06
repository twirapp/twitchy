package eventsub

import (
	"errors"
	"fmt"

	"github.com/kvizyx/twitchy/internal/json"
)

// ErrUndefinedEventType indicates that eventsub server sent event that is not defined in library and user does not set
// callback handler for undefined events.
var ErrUndefinedEventType = errors.New("undefined event type")

func (c *callback[Metadata]) runEventCallback(eventType EventType, eventVersion string, payload []byte, metadata Metadata) error {
	switch eventType {
	case EventTypeAutomodMessageHold:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageHold, payload, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageHoldV2, payload, metadata)
		}
	case EventTypeAutomodMessageUpdate:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onAutomodMessageUpdate, payload, metadata)
		case "2":
			return runCallbackHandler(c.onAutomodMessageUpdateV2, payload, metadata)
		}
	case EventTypeAutomodSettingsUpdate:
		return runCallbackHandler(c.onAutomodSettingsUpdate, payload, metadata)
	case EventTypeAutomodTermsUpdate:
		return runCallbackHandler(c.onAutomodTermsUpdate, payload, metadata)
	case EventTypeChannelBitsUse:
		return runCallbackHandler(c.onChannelBitsUse, payload, metadata)
	case EventTypeChannelUpdate:
		return runCallbackHandler(c.onChannelUpdate, payload, metadata)
	case EventTypeChannelFollow:
		return runCallbackHandler(c.onChannelFollow, payload, metadata)
	case EventTypeChannelAdBreakBegin:
		return runCallbackHandler(c.onChannelAdBreakBegin, payload, metadata)
	case EventTypeChannelChatClear:
		return runCallbackHandler(c.onChannelChatClear, payload, metadata)
	case EventTypeChannelChatClearUserMessages:
		return runCallbackHandler(c.onChannelChatClearUserMessages, payload, metadata)
	case EventTypeChannelChatMessage:
		return runCallbackHandler(c.onChannelChatMessage, payload, metadata)
	case EventTypeConduitShardDisabled:
		return runCallbackHandler(c.onConduitShardDisabled, payload, metadata)
	case EventTypeChannelBan:
		return runCallbackHandler(c.onChannelBan, payload, metadata)
	case EventTypeChannelChatNotification:
		return runCallbackHandler(c.onChannelChatNotification, payload, metadata)
	case EventTypeChannelModeratorAdd:
		return runCallbackHandler(c.onChannelModeratorAdd, payload, metadata)
	case EventTypeChannelModeratorRemove:
		return runCallbackHandler(c.onChannelModeratorRemove, payload, metadata)
	case EventTypeChannelPollBegin:
		return runCallbackHandler(c.onChannelPollBegin, payload, metadata)
	case EventTypeChannelPollProgress:
		return runCallbackHandler(c.onChannelPollProgress, payload, metadata)
	case EventTypeChannelPollEnd:
		return runCallbackHandler(c.onChannelPollEnd, payload, metadata)
	case EventTypeChannelPredictionBegin:
		return runCallbackHandler(c.onChannelPredictionBegin, payload, metadata)
	case EventTypeChannelPredictionProgress:
		return runCallbackHandler(c.onChannelPredictionProgress, payload, metadata)
	case EventTypeChannelPredictionLock:
		return runCallbackHandler(c.onChannelPredictionLock, payload, metadata)
	case EventTypeChannelPredictionEnd:
		return runCallbackHandler(c.onChannelPredictionEnd, payload, metadata)
	case EventTypeChannelRaid:
		return runCallbackHandler(c.onChannelRaid, payload, metadata)
	case EventTypeChannelPointsCustomRewardRedemptionAdd:
		return runCallbackHandler(c.onChannelPointsCustomRewardRedemptionAdd, payload, metadata)
	case EventTypeChannelPointsCustomRewardRedemptionUpdate:
		return runCallbackHandler(c.onChannelPointsCustomRewardRedemptionUpdate, payload, metadata)
	case EventTypeChannelPointsAutomaticRewardRedemptionAdd:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onChannelPointsAutomaticRewardRedemptionAdd, payload, metadata)
		case "2":
			return runCallbackHandler(c.onChannelPointsAutomaticRewardRedemptionAddV2, payload, metadata)
		}
	case EventTypeUserAuthorizationRevoke:
		return runCallbackHandler(c.onUserAuthorizationRevoke, payload, metadata)
	case EventTypeChannelPointsRewardAdd:
		return runCallbackHandler(c.onChannelPointsCustomRewardAdd, payload, metadata)
	case EventTypeChannelPointsRewardUpdate:
		return runCallbackHandler(c.onChannelPointsCustomRewardUpdate, payload, metadata)
	case EventTypeChannelPointsRewardRemove:
		return runCallbackHandler(c.onChannelPointsCustomRewardRemove, payload, metadata)
	case EventTypeStreamOffline:
		return runCallbackHandler(c.onStreamOffline, payload, metadata)
	case EventTypeStreamOnline:
		return runCallbackHandler(c.onStreamOnline, payload, metadata)
	case EventTypeChannelSubscribe:
		return runCallbackHandler(c.onChannelSubscribe, payload, metadata)
	case EventTypeChannelSubscriptionEnd:
		return runCallbackHandler(c.onChannelSubscriptionEnd, payload, metadata)
	case EventTypeChannelSubscriptionMessage:
		return runCallbackHandler(c.onChannelSubscriptionMessage, payload, metadata)
	case EventTypeChannelSubscriptionGift:
		return runCallbackHandler(c.onChannelSubscriptionGift, payload, metadata)
	case EventTypeChannelUnbanRequestCreate:
		return runCallbackHandler(c.onChannelUnbanRequestCreate, payload, metadata)
	case EventTypeChannelUnbanRequestResolve:
		return runCallbackHandler(c.onChannelUnbanRequestResolve, payload, metadata)
	case EventTypeUserUpdate:
		return runCallbackHandler(c.onUserUpdate, payload, metadata)
	case EventTypeChannelVipAdd:
		return runCallbackHandler(c.onChannelVipAdd, payload, metadata)
	case EventTypeChannelVipRemove:
		return runCallbackHandler(c.onChannelVipRemove, payload, metadata)
	case EventTypeChannelMessageDelete:
		return runCallbackHandler(c.onChannelChatMessageDelete, payload, metadata)
	default:
		return ErrUndefinedEventType
	}

	return nil
}

// runCallbackHandler parses JSON payload and runs handler in a separate go-routine with parsed payload if it's not nil.
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
