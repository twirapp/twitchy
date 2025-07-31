package eventsub

import (
	"context"
	"encoding/json"
	"fmt"
)

type websocketRawMessage struct {
	Metadata struct {
		MessageId           string       `json:"message_id"`
		MessageType         string       `json:"message_type"`
		MessageTimestamp    TimestampUTC `json:"message_timestamp"`
		SubscriptionType    string       `json:"subscription_type"`
		SubscriptionVersion string       `json:"subscription_version"`
	} `json:"metadata"`
	Payload json.RawMessage `json:"payload"`
	Event   json.RawMessage `json:"event,omitempty"`
}

func (ws *Websocket) handleMessage(ctx context.Context, message websocketRawMessage) error {
	metadata := message.Metadata

	switch metadata.MessageType {
	case "session_welcome":
		var payload WebsocketWelcomePayload

		if err := json.Unmarshal(message.Payload, &payload); err != nil {
			return fmt.Errorf("unmsrshal payload: %w", err)
		}

		welcome := WebsocketWelcomeMessage{
			Metadata: WebsocketWelcomeMetadata{
				MessageId:        metadata.MessageId,
				MessageType:      metadata.MessageType,
				MessageTimestamp: metadata.MessageTimestamp,
			},
			Payload: payload,
		}

		select {
		case ws.welcome <- struct{}{}:
		default:
		}

		if ws.onWelcome != nil {
			go ws.onWelcome(welcome)
		}
	case "session_keepalive":
		keepalive := WebsocketKeepaliveMessage{
			Metadata: WebsocketKeepaliveMetadata{
				MessageId:        metadata.MessageId,
				MessageType:      metadata.MessageType,
				MessageTimestamp: metadata.MessageTimestamp,
			},
			Payload: struct{}{},
		}

		ws.setKeepalive(keepalive.Metadata.MessageTimestamp)

		if ws.onKeepalive != nil {
			go ws.onKeepalive(keepalive)
		}
	case "session_reconnect":
		var payload WebsocketReconnectPayload

		if err := json.Unmarshal(message.Payload, &payload); err != nil {
			return fmt.Errorf("unmsrshal payload: %w", err)
		}

		reconnect := WebsocketReconnectMessage{
			Metadata: WebsocketReconnectMetadata{
				MessageId:        metadata.MessageId,
				MessageType:      metadata.MessageType,
				MessageTimestamp: metadata.MessageTimestamp,
			},
			Payload: payload,
		}

		go func() {
			if err := ws.reconnect(ctx, reconnect.Payload.Session.ReconnectURL); err != nil {
				if ws.onReconnectError != nil {
					go ws.onReconnectError(err)
				}
			}
		}()

		if ws.onReconnect != nil {
			go ws.onReconnect(reconnect)
		}
	case "notification":
		if ws.eventTracker != nil {
			isDuplicate, err := ws.eventTracker.Track(ctx, metadata.MessageId)
			if err != nil {
				return fmt.Errorf("track event: %w", err)
			}

			if isDuplicate {
				if ws.onDuplicate != nil {
					go ws.onDuplicate(metadata.MessageId)
				}
				return nil
			}
		}

		nm := WebsocketNotificationMetadata{
			MessageId:           metadata.MessageId,
			MessageType:         metadata.MessageType,
			MessageTimestamp:    metadata.MessageTimestamp,
			SubscriptionType:    metadata.SubscriptionType,
			SubscriptionVersion: metadata.SubscriptionVersion,
		}

		return ws.callback.runEventCallback(metadata.SubscriptionType, metadata.SubscriptionVersion, message.Event, nm)
	case "revocation":
		// TODO: implement me
	}

	return nil
}
