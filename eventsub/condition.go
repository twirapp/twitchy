package eventsub

type ConditionMarker interface {
	isCondition()
}

type AutomodMessageHoldCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster (channel).
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type AutomodMessageUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster (channel), maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type AutomodSettingsUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster (channel), maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type AutomodTermsUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster (channel), maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelAdBreakBeginCondition struct {
	ConditionMarker
	// BroadcasterId is an id of the broadcaster that you want to get channel ad break begin notifications for, maximum 1.
	BroadcasterId string `json:"broadcaster_user_id"`
}

type ChannelBanCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get ban notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelBitsUseCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel broadcaster, maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelChatClearCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat clear events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatClearUserMessagesCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat clear user messages events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatMessageCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat message events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatMessageDeleteCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat message delete events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatNotificationCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat notification events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatSettingsUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat settings update events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatUserMessageHoldCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat message events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelChatUserMessageUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive chat message events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// UserId is a user ID to read chat as.
	UserId string `json:"user_id"`
}

type ChannelSubscribeCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get subscribe notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSubscriptionEndCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get subscription end notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSubscriptionGiftCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get subscription gift notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSubscriptionMessageCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get resubscription chat message notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelCheerCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get cheer notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get updates for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelFollowCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get follow notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is an ID of the moderator of the channel you want to get follow notifications for.
	// If you have authorization from the broadcaster rather than a moderator, specify the broadcaster’s user ID here.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelUnbanCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get unban notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelUnbanRequestCreateCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the broadcaster you want to get chat unban request notifications for, maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is an ID of the user that has permission to moderate the broadcaster’s channel and has granted
	// your app permission to subscribe to this subscription type.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelUnbanRequestResolveCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the broadcaster you want to get unban request resolution notifications for, maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is an ID of the user that has permission to moderate the broadcaster’s channel and has granted
	// your app permission to subscribe to this subscription type.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelRaidCondition struct {
	ConditionMarker
	// FromBroadcasterUserId is a broadcaster user ID that created the channel raid you want to get notifications for.
	// Use this parameter if you want to know when a specific broadcaster raids another broadcaster. The channel raid
	// condition must include either from_broadcaster_user_id or to_broadcaster_user_id.
	FromBroadcasterUserId string `json:"from_broadcaster_user_id"`
	// ToBroadcasterUserId is a broadcaster user ID that received the channel raid you want to get notifications for.
	// Use this parameter if you want to know when a specific broadcaster is raided by another broadcaster. The channel
	// raid condition must include either from_broadcaster_user_id or to_broadcaster_user_id.
	ToBroadcasterUserId string `json:"to_broadcaster_user_id"`
}

type ChannelModerateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelModerateV2Condition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelModeratorAddCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get moderator addition notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelModeratorRemoveCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to get moderator removal notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelGuestStarSessionBeginCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel hosting the guest star session.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator or broadcaster of the specified channel.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelGuestStarSessionEndCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel hosting the guest star session.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator or broadcaster of the specified channel.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelGuestStarGuestUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel hosting the guest star session.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator or broadcaster of the specified channel.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelGuestStarSettingsUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel hosting the guest star session.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator or broadcaster of the specified channel.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelPointsAutomaticRewardRedemptionAddCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points reward add notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPointsAutomaticRewardRedemptionAddV2Condition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points reward add notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPointsCustomRewardAddCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points custom reward add notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPointsCustomRewardUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points custom reward update notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// RewardId is a reward id to only receive notifications for a specific reward, optional.
	RewardId string `json:"reward_id,omitempty"`
}

type ChannelPointsCustomRewardRemoveCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points custom reward remove notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// RewardId is a reward id to only receive notifications for a specific reward, optional.
	RewardId string `json:"reward_id,omitempty"`
}

type ChannelPointsCustomRewardRedemptionAddCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points custom reward redemption add notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// RewardId is a reward id to only receive notifications for a specific reward, optional.
	RewardId string `json:"reward_id,omitempty"`
}

type ChannelPointsCustomRewardRedemptionUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID for the channel you want to receive channel points custom reward redemption update notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// RewardId is a reward id to only receive notifications for a specific reward, optional.
	RewardId string `json:"reward_id,omitempty"`
}

type ChannelPollBeginCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “poll begin” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPollProgressCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “poll progress” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPollEndCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “poll end” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPredictionBeginCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “prediction begin” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPredictionProgressCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “prediction progress” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPredictionLockCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “prediction lock” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelPredictionEndCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID of the channel for which “prediction end” notifications will be received.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSharedChatSessionBeginCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive shared chat session begin events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSharedChatSessionUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive shared chat session update events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSharedChatSessionEndCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the channel to receive shared chat session end events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelSuspiciousUserMessageCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the channel to receive chat message events for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is an ID of a user that has permission to moderate the broadcaster’s channel and has granted your
	// app permission to subscribe to this subscription type.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelSuspiciousUserUpdateCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster you want to get chat unban request notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is an ID of a user that has permission to moderate the broadcaster’s channel and has granted your
	// app permission to subscribe to this subscription type.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelVIPAddCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster (channel), maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelVIPRemoveCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster (channel), maximum 1.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type ChannelWarningAcknowledgeCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ChannelWarningSendCondition struct {
	ConditionMarker
	// BroadcasterUserId is a user ID of the broadcaster.
	BroadcasterUserId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a user ID of the moderator.
	ModeratorUserId string `json:"moderator_user_id"`
}

type ConduitShardDisabledCondition struct {
	ConditionMarker
	// ClientId is your application’s client id. The provided client_id must match the client ID in the application access token.
	ClientId string `json:"broadcaster_user_id"`
	// ModeratorUserId is a conduit ID to receive events for. If omitted, events for all of this client’s conduits are sent.
	ConduitId string `json:"moderator_user_id"`
}

type DropEntitlementGrantCondition struct {
	ConditionMarker
	// OrganizationId is an organization ID of the organization that owns the game on the developer portal.
	OrganizationId string `json:"organization_id"`
	// CategoryId is a category (or game) ID of the game for which entitlement notifications will be received.
	CategoryId string `json:"category_id"`
	// CampaignId is a campaign ID for a specific campaign for which entitlement notifications will be received.
	CampaignId string `json:"campaign_id"`
}

type ExtensionBitsTransactionCreateCondition struct {
	ConditionMarker
	// ExtensionClientId is a client ID of the extension.
	ExtensionClientId string `json:"extension_client_id"`
}

type GoalsCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the broadcaster to get notified about. The ID must match the user_id in the OAuth access token.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type HypeTrainBeginCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the broadcaster that you want to get hype train begin notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type HypeTrainProgressCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the broadcaster that you want to get hype train progress notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type HypeTrainEndCondition struct {
	ConditionMarker
	// BroadcasterUserId is an ID of the broadcaster that you want to get hype train end notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type StreamOnlineCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID you want to get stream online notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type StreamOfflineCondition struct {
	ConditionMarker
	// BroadcasterUserId is a broadcaster user ID you want to get stream offline notifications for.
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type UserAuthorizationGrantCondition struct {
	ConditionMarker
	// ClientId is your application’s client id. The provided client_id must match the client id in the application access token.
	ClientId string `json:"client_id"`
}

type UserAuthorizationRevokeCondition struct {
	ConditionMarker
	// ClientId is your application’s client id. The provided client_id must match the client id in the application access token.
	ClientId string `json:"client_id"`
}

type UserUpdateCondition struct {
	ConditionMarker
	// UserId is a user ID for the user you want update notifications for.
	UserId string `json:"user_id"`
}

type WhisperReceivedCondition struct {
	ConditionMarker
	// UserId is a user_id of the person receiving whispers
	UserId string `json:"user_id"`
}
