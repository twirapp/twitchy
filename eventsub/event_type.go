package eventsub

type EventType string

const (
	EventTypeAutomodMessageUpdate         EventType = "automod.message.update"
	EventTypeAutomodMessageHold           EventType = "automod.message.hold"
	EventTypeAutomodSettingsUpdate        EventType = "automod.settings.update"
	EventTypeAutomodTermsUpdate           EventType = "automod.terms.update"
	EventTypeChannelBitsUse               EventType = "channel.bits.use"
	EventTypeChannelUpdate                EventType = "channel.update"
	EventTypeChannelFollow                EventType = "channel.follow"
	EventTypeChannelAdBreakBegin          EventType = "channel.ad_break.begin"
	EventTypeChannelChatClear             EventType = "channel.chat.clear"
	EventTypeChannelChatClearUserMessages EventType = "channel.chat.clear_user_messages"
	EventTypeChannelChatMessage           EventType = "channel.chat.message"
	EventTypeConduitShardDisabled         EventType = "conduit.shard.disabled"
)

func (et EventType) String() string {
	return string(et)
}
