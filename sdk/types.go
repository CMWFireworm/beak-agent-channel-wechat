package sdk

import (
	"strings"

	adaptor "beak/pkg/channeladaptor"
)

const (
	ChatTypeGroup  = "group"
	ChatTypeDirect = "direct"

	LoginModeQRCode     = adaptor.LoginModeQRCode
	LoginModeCredential = adaptor.LoginModeCredential
)

type Connector = adaptor.RuntimeConnector
type ConnectorMetadata = adaptor.Metadata
type Metadata = adaptor.Metadata
type Capabilities = adaptor.Capabilities
type CredentialSchema = adaptor.CredentialSchema
type CredentialField = adaptor.CredentialField
type Runtime = adaptor.Runtime
type Gateway = adaptor.Gateway
type AccountStore = adaptor.AccountStore
type Channel = adaptor.Channel
type ChannelAccount = adaptor.RuntimeAccount
type RuntimeAccount = adaptor.RuntimeAccount
type Account = adaptor.Account
type LoginStartRequest = adaptor.LoginStartRequest
type LoginPollRequest = adaptor.LoginPollRequest
type LoginChallenge = adaptor.LoginChallenge
type LoginStatus = adaptor.LoginStatus
type OutboundMessage = adaptor.OutboundMessage
type SendResult = adaptor.SendResult
type EnsureChannelRequest = adaptor.EnsureChannelRequest
type EnsureChannelLinkSessionRequest = adaptor.EnsureChannelLinkSessionRequest
type EnsureChatSessionRequest = adaptor.EnsureChatSessionRequest
type CreateMessageRequest = adaptor.CreateMessageRequest
type StreamSessionRequest = adaptor.StreamSessionRequest
type StreamEvent = adaptor.StreamEvent

type InboundMessage struct {
	WorkspaceUUID string         `json:"workspace_uuid"`
	Platform      string         `json:"platform"`
	AccountUUID   string         `json:"account_uuid"`
	ChannelUUID   string         `json:"channel_uuid"`
	ChatType      string         `json:"chat_type"`
	ChatID        string         `json:"chat_id"`
	SenderID      string         `json:"sender_id"`
	MessageID     string         `json:"message_id,omitempty"`
	Text          string         `json:"text"`
	DedupeKey     string         `json:"dedupe_key,omitempty"`
	Raw           map[string]any `json:"raw,omitempty"`
}

func IMUserParticipantID(platform, chatType, chatID, senderID string) string {
	parts := []string{"im", platform, chatType, chatID, "user", senderID}
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return strings.Join(parts, ":")
}
