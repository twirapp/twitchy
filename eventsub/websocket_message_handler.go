package eventsub

import (
	"context"
	"fmt"

	"github.com/kvizyx/twitchy/internal/json"
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
		Subscription json.RawMessage `json:"subscription"`
		Event        json.RawMessage `json:"event"`
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
	_ context.Context,
	rawMetadata websocketRawMessageMetadata,
	rawPayload json.RawMessage,
) error {
	var payload websocketRawMessagePayload

	if err := json.Unmarshal(rawPayload, &payload); err != nil {
		return fmt.Errorf("unmsrshal payload: %w", err)
	}

	var (
		event = RawEvent{
			Subscription: payload.Subscription,
			Event:        payload.Event,
		}

		metadata = WebsocketNotificationMetadata{
			MessageId:           rawMetadata.MessageId,
			MessageType:         rawMetadata.MessageType,
			MessageTimestamp:    rawMetadata.MessageTimestamp,
			SubscriptionType:    rawMetadata.SubscriptionType,
			SubscriptionVersion: rawMetadata.SubscriptionVersion,
		}
	)

	if err := ws.callback.runEventCallback(metadata.SubscriptionType, metadata.SubscriptionVersion, event, metadata); err != nil {
		return fmt.Errorf("run event callback: %w", err)
	}

	return nil
}
