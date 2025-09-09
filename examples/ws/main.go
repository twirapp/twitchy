package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kvizyx/twitchy/eventsub"
)

func main() {
	manager := eventsub.New()

	ws := manager.Websocket()

	ctx := context.TODO()

	ws.OnWelcome(func(message eventsub.WebsocketWelcomeMessage) {
		// you can for example make api request to update subscription list
	})

	ws.OnChannelPointsAutomaticRewardRedemptionAddV2(func(v2 eventsub.ChannelPointsAutomaticRewardRedemptionAddEventV2, metadata eventsub.WebsocketNotificationMetadata) {
		//someone used redemption
	})

	go func() {
		for {
			// since it is infinite loop it will auto-reconnect
			if err := ws.Connect(ctx); err != nil {
				fmt.Printf("disconnected: %v\n", err)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	select {}
}
