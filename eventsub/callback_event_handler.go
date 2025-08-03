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
	case EventTypeChannelBan:
		return runCallbackHandler(c.onChannelBan, body, metadata)
	case EventTypeChannelChatNotification:
		return runCallbackHandler(c.onChannelChatNotification, body, metadata)
	case EventTypeChannelModeratorAdd:
		return runCallbackHandler(c.onChannelModeratorAdd, body, metadata)
	case EventTypeChannelModeratorRemove:
		return runCallbackHandler(c.onChannelModeratorRemove, body, metadata)
	case EventTypeChannelPollBegin:
		return runCallbackHandler(c.onChannelPollBegin, body, metadata)
	case EventTypeChannelPollProgress:
		return runCallbackHandler(c.onChannelPollProgress, body, metadata)
	case EventTypeChannelPollEnd:
		return runCallbackHandler(c.onChannelPollEnd, body, metadata)
	case EventTypeChannelPredictionBegin:
		return runCallbackHandler(c.onChannelPredictionBegin, body, metadata)
	case EventTypeChannelPredictionProgress:
		return runCallbackHandler(c.onChannelPredictionProgress, body, metadata)
	case EventTypeChannelPredictionLock:
		return runCallbackHandler(c.onChannelPredictionLock, body, metadata)
	case EventTypeChannelPredictionEnd:
		return runCallbackHandler(c.onChannelPredictionEnd, body, metadata)
	case EventTypeChannelRaid:
		return runCallbackHandler(c.onChannelRaid, body, metadata)
	case EventTypeChannelPointsCustomRewardRedemptionAdd:
		return runCallbackHandler(c.onChannelPointsCustomRewardRedemptionAdd, body, metadata)
	case EventTypeChannelPointsCustomRewardRedemptionUpdate:
		return runCallbackHandler(c.onChannelPointsCustomRewardRedemptionUpdate, body, metadata)
	case EventTypeChannelPointsAutomaticRewardRedemptionAdd:
		switch eventVersion {
		case "1":
			return runCallbackHandler(c.onChannelPointsAutomaticRewardRedemptionAdd, body, metadata)
		case "2":
			return runCallbackHandler(c.onChannelPointsAutomaticRewardRedemptionAddV2, body, metadata)
		}
	case EventTypeUserAuthorizationRevoke:
		return runCallbackHandler(c.onUserAuthorizationRevoke, body, metadata)
	case EventTypeChannelPointsRewardAdd:
		return runCallbackHandler(c.onChannelPointsRewardAdd, body, metadata)
	case EventTypeChannelPointsRewardUpdate:
		return runCallbackHandler(c.onChannelPointsRewardUpdate, body, metadata)
	case EventTypeChannelPointsRewardRemove:
		return runCallbackHandler(c.onChannelPointsRewardRemove, body, metadata)
	case EventTypeStreamOffline:
		return runCallbackHandler(c.onStreamOffline, body, metadata)
	case EventTypeStreamOnline:
		return runCallbackHandler(c.onStreamOnline, body, metadata)
	case EventTypeChannelSubscribe:
		return runCallbackHandler(c.onChannelSubscribe, body, metadata)
	case EventTypeChannelSubscriptionEnd:
		return runCallbackHandler(c.onChannelSubscriptionEnd, body, metadata)
	case EventTypeChannelSubscriptionMessage:
		return runCallbackHandler(c.onChannelSubscriptionMessage, body, metadata)
	case EventTypeChannelSubscriptionGift:
		return runCallbackHandler(c.onChannelSubscriptionGift, body, metadata)
	case EventTypeChannelUnbanRequestCreate:
		return runCallbackHandler(c.onChannelUnbanRequestCreate, body, metadata)
	case EventTypeChannelUnbanRequestResolve:
		return runCallbackHandler(c.onChannelUnbanRequestResolve, body, metadata)
	case EventTypeUserUpdate:
		return runCallbackHandler(c.onUserUpdate, body, metadata)
	case EventTypeChannelVipAdd:
		return runCallbackHandler(c.onChannelVipAdd, body, metadata)
	case EventTypeChannelVipRemove:
		return runCallbackHandler(c.onChannelVipRemove, body, metadata)
	case EventTypeChannelMessageDelete:
		return runCallbackHandler(c.onChannelChatMessageDelete, body, metadata)
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
