package eventsub

type Handler[Event any, Metadata any] func(Event, Metadata)

// callback is a store for all EventSub subscription callbacks that are general (available on all transports).
type callback[Metadata any] struct {
	onAutomodMessageHold           Handler[AutomodMessageHoldEvent, Metadata]
	onAutomodMessageHoldV2         Handler[AutomodMessageHoldEventV2, Metadata]
	onAutomodMessageUpdate         Handler[AutomodMessageUpdateEvent, Metadata]
	onAutomodMessageUpdateV2       Handler[AutomodMessageUpdateEventV2, Metadata]
	onAutomodSettingsUpdate        Handler[AutomodSettingsUpdateEvent, Metadata]
	onAutomodTermsUpdate           Handler[AutomodTermsUpdateEvent, Metadata]
	onChannelBitsUse               Handler[ChannelBitsUseEvent, Metadata]
	onChannelUpdate                Handler[ChannelUpdateEvent, Metadata]
	onChannelFollow                Handler[ChannelFollowEvent, Metadata]
	onChannelAdBreakBegin          Handler[ChannelAdBreakBeginEvent, Metadata]
	onChannelChatClear             Handler[ChannelChatClearEvent, Metadata]
	onChannelChatClearUserMessages Handler[ChannelChatClearUserMessagesEvent, Metadata]
	onChannelChatMessage           Handler[ChannelChatMessageEvent, Metadata]
}

// OnAutomodMessageHold invokes when message is caught by automod for review.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#automodmessagehold.
func (c *callback[Metadata]) OnAutomodMessageHold(onAutomodMessageHold Handler[AutomodMessageHoldEvent, Metadata]) {
	c.onAutomodMessageHold = onAutomodMessageHold
}

// OnAutomodMessageHoldV2 invokes when message is caught by automod for review.
// Only public blocked terms trigger notifications, not private ones.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#automodmessagehold-v2.
func (c *callback[Metadata]) OnAutomodMessageHoldV2(onAutomodMessageHoldV2 Handler[AutomodMessageHoldEventV2, Metadata]) {
	c.onAutomodMessageHoldV2 = onAutomodMessageHoldV2
}

// OnAutomodMessageUpdate invokes when a message in the automod queue had its status changed.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#automodmessageupdate.
func (c *callback[Metadata]) OnAutomodMessageUpdate(onAutomodMessageUpdate Handler[AutomodMessageUpdateEvent, Metadata]) {
	c.onAutomodMessageUpdate = onAutomodMessageUpdate
}

// OnAutomodMessageUpdateV2 invokes when a message in the automod queue had its status changed.
// Only public blocked terms trigger notifications, not private ones.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#automodmessageupdate-v2.
func (c *callback[Metadata]) OnAutomodMessageUpdateV2(onAutomodMessageUpdateV2 Handler[AutomodMessageUpdateEventV2, Metadata]) {
	c.onAutomodMessageUpdateV2 = onAutomodMessageUpdateV2
}

// OnAutomodSettingsUpdate invokes when  a broadcaster’s automod settings are updated.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#automodsettingsupdate.
func (c *callback[Metadata]) OnAutomodSettingsUpdate(onAutomodSettingsUpdate Handler[AutomodSettingsUpdateEvent, Metadata]) {
	c.onAutomodSettingsUpdate = onAutomodSettingsUpdate
}

// OnAutomodTermsUpdate invokes when a broadcaster’s automod terms are updated. Changes to private terms are not sent.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#automodtermsupdate.
func (c *callback[Metadata]) OnAutomodTermsUpdate(onAutomodTermsUpdate Handler[AutomodTermsUpdateEvent, Metadata]) {
	c.onAutomodTermsUpdate = onAutomodTermsUpdate
}

// OnChannelBitsUse invokes when bits are used on a channel.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelbitsuse.
func (c *callback[Metadata]) OnChannelBitsUse(onChannelBitsUse Handler[ChannelBitsUseEvent, Metadata]) {
	c.onChannelBitsUse = onChannelBitsUse
}

// OnChannelUpdate invokes when a broadcaster updates the category, title, content classification labels, or broadcast
// language for their channel.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelupdate.
func (c *callback[Metadata]) OnChannelUpdate(onChannelUpdate Handler[ChannelUpdateEvent, Metadata]) {
	c.onChannelUpdate = onChannelUpdate
}

// OnChannelFollow invokes when a specified channel receives a follow.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelfollow
func (c *callback[Metadata]) OnChannelFollow(onChannelFollow Handler[ChannelFollowEvent, Metadata]) {
	c.onChannelFollow = onChannelFollow
}

// OnChannelAdBreakBegin invokes when a user runs a midroll commercial break, either manually or automatically via ads manager.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelad_breakbegin.
func (c *callback[Metadata]) OnChannelAdBreakBegin(onChannelAdBreakBegin Handler[ChannelAdBreakBeginEvent, Metadata]) {
	c.onChannelAdBreakBegin = onChannelAdBreakBegin
}

// OnChannelChatClear invokes when a moderator or bot clears all messages from the chat room.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelchatclear.
func (c *callback[Metadata]) OnChannelChatClear(onChannelChatClear Handler[ChannelChatClearEvent, Metadata]) {
	c.onChannelChatClear = onChannelChatClear
}

// OnChannelChatClearUserMessages invokes when a moderator or bot clears all messages for a specific user.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelchatclear_user_messages.
func (c *callback[Metadata]) OnChannelChatClearUserMessages(onChannelChatClearUserMessages Handler[ChannelChatClearUserMessagesEvent, Metadata]) {
	c.onChannelChatClearUserMessages = onChannelChatClearUserMessages
}

// OnChannelChatMessage invokes when any user sends a message to a channel’s chat room.
//
// Reference: https://dev.twitch.tv/dogc/eventsub/eventsub-subscription-types/#channelchatmessage.
func (c *callback[Metadata]) OnChannelChatMessage(onChannelChatMessage Handler[ChannelChatMessageEvent, Metadata]) {
	c.onChannelChatMessage = onChannelChatMessage
}
