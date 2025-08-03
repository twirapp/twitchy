package eventsub

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	websocketRawMessage struct {
		Metadata websocketRawMessageMetadata `json:"metadata"`
		Payload  json.RawMessage             `json:"payload"`
	}

	websocketRawMessageMetadata struct {
		MessageId           string       `json:"message_id"`
		MessageType         string       `json:"message_type"`
		MessageTimestamp    TimestampUTC `json:"message_timestamp"`
		SubscriptionType    EventType    `json:"subscription_type"`
		SubscriptionVersion string       `json:"subscription_version"`
	}

	websocketRawMessagePayload struct {
		Event json.RawMessage `json:"event,omitempty"`
	}
)

func (ws *Websocket) handleMessage(ctx context.Context, message websocketRawMessage) error {
	metadata := message.Metadata

	switch metadata.MessageType {
	case "session_welcome":
		if err := ws.handleWelcomeMessage(metadata, message.Payload); err != nil {
			return fmt.Errorf("handle welcome message: %w", err)
		}
	case "session_keepalive":
		ws.handleKeepaliveMessage(metadata)
	case "session_reconnect":
		if err := ws.handleReconnectMessage(ctx, metadata, message.Payload); err != nil {
			return fmt.Errorf("handle reconnect message: %w", err)
		}
	case "notification":
		if err := ws.handleNotificationMessage(ctx, metadata, message.Payload); err != nil {
			return fmt.Errorf("handle notification message: %w", err)
		}
	case "revocation":
		// TODO: implement me
	}

	return nil
}

func (ws *Websocket) handleWelcomeMessage(metadata websocketRawMessageMetadata, messagePayload json.RawMessage) error {
	var payload WebsocketWelcomePayload

	if err := json.Unmarshal(messagePayload, &payload); err != nil {
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

	return nil
}

func (ws *Websocket) handleKeepaliveMessage(metadata websocketRawMessageMetadata) {
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
}

func (ws *Websocket) handleReconnectMessage(
	ctx context.Context,
	metadata websocketRawMessageMetadata,
	messagePayload json.RawMessage,
) error {
	var payload WebsocketReconnectPayload

	if err := json.Unmarshal(messagePayload, &payload); err != nil {
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

	return nil
}

func (ws *Websocket) handleNotificationMessage(
	ctx context.Context,
	metadata websocketRawMessageMetadata,
	messagePayload json.RawMessage,
) error {
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

	var payload websocketRawMessagePayload

	if err := json.Unmarshal(messagePayload, &payload); err != nil {
		return fmt.Errorf("unmsrshal payload: %w", err)
	}

	return ws.callback.runEventCallback(metadata.SubscriptionType, metadata.SubscriptionVersion, payload.Event, nm)
}
