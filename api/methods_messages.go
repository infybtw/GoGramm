package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// GetUpdatesParams holds the parameters for the getUpdates method.
type GetUpdatesParams struct {
	Offset         *int     `json:"offset,omitempty"`
	Limit          *int     `json:"limit,omitempty"`
	Timeout        *int     `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// GetUpdates receives incoming updates using long polling.
func (c *Client) GetUpdates(ctx context.Context, p *GetUpdatesParams) ([]Update, error) {
	var updates []Update
	if err := c.Invoke(ctx, "getUpdates", p, &updates); err != nil {
		return nil, err
	}
	return updates, nil
}

// SetWebhookParams holds the parameters for the setWebhook method.
type SetWebhookParams struct {
	URL                string     `json:"url"`
	Certificate        *InputFile `json:"certificate,omitempty"`
	IPAddress          *string    `json:"ip_address,omitempty"`
	MaxConnections     *int       `json:"max_connections,omitempty"`
	AllowedUpdates     []string   `json:"allowed_updates,omitempty"`
	DropPendingUpdates *bool      `json:"drop_pending_updates,omitempty"`
	SecretToken        *string    `json:"secret_token,omitempty"`
}

// SetWebhook specifies a URL and receives incoming updates via an outgoing webhook.
func (c *Client) SetWebhook(ctx context.Context, p *SetWebhookParams) error {
	return c.Invoke(ctx, "setWebhook", p, nil)
}

// DeleteWebhookParams holds the parameters for the deleteWebhook method.
type DeleteWebhookParams struct {
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook removes webhook integration if you decide to switch back to getUpdates.
func (c *Client) DeleteWebhook(ctx context.Context, p *DeleteWebhookParams) error {
	return c.Invoke(ctx, "deleteWebhook", p, nil)
}

// GetWebhookInfo returns the current webhook status.
func (c *Client) GetWebhookInfo(ctx context.Context) (*WebhookInfo, error) {
	var w WebhookInfo
	if err := c.Invoke(ctx, "getWebhookInfo", nil, &w); err != nil {
		return nil, err
	}
	return &w, nil
}

// GetMe returns basic information about the bot in the form of a User object.
func (c *Client) GetMe(ctx context.Context) (*User, error) {
	var u User
	if err := c.Invoke(ctx, "getMe", nil, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

// LogOut logs out from the cloud Bot API server before launching the bot locally.
func (c *Client) LogOut(ctx context.Context) error {
	return c.Invoke(ctx, "logOut", nil, nil)
}

// Close closes the bot instance before moving it from one local server to another.
func (c *Client) Close(ctx context.Context) error {
	return c.Invoke(ctx, "close", nil, nil)
}

// SendMessageParams holds the parameters for sendMessage.
type SendMessageParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Text                       string                      `json:"text"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	Entities                   *[]MessageEntity            `json:"entities,omitempty"`
	LinkPreviewOptions         *LinkPreviewOptions         `json:"link_preview_options,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendMessage sends text messages. On success, the sent Message is returned.
func (c *Client) SendMessage(ctx context.Context, p *SendMessageParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendMessage", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// ForwardMessageParams holds the parameters for forwardMessage.
type ForwardMessageParams struct {
	ChatID                  ChatID                   `json:"chat_id"`
	MessageThreadID         *int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   *int64                   `json:"direct_messages_topic_id,omitempty"`
	FromChatID              ChatID                   `json:"from_chat_id"`
	VideoStartTimestamp     *int64                   `json:"video_start_timestamp,omitempty"`
	DisableNotification     *bool                    `json:"disable_notification,omitempty"`
	ProtectContent          *bool                    `json:"protect_content,omitempty"`
	MessageEffectID         *string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	MessageID               int64                    `json:"message_id"`
}

// ForwardMessage forwards messages of any kind; service messages and messages
// with protected content can't be forwarded. On success, the sent Message is returned.
func (c *Client) ForwardMessage(ctx context.Context, p *ForwardMessageParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "forwardMessage", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// ForwardMessagesParams holds the parameters for forwardMessages.
type ForwardMessagesParams struct {
	ChatID                ChatID  `json:"chat_id"`
	MessageThreadID       *int64  `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID *int64  `json:"direct_messages_topic_id,omitempty"`
	FromChatID            ChatID  `json:"from_chat_id"`
	MessageIDs            []int64 `json:"message_ids"`
	DisableNotification   *bool   `json:"disable_notification,omitempty"`
	ProtectContent        *bool   `json:"protect_content,omitempty"`
}

// ForwardMessages forwards multiple messages of any kind. On success, an
// Array of MessageId of the sent messages is returned.
func (c *Client) ForwardMessages(ctx context.Context, p *ForwardMessagesParams) ([]MessageId, error) {
	var ids []MessageId
	if err := c.Invoke(ctx, "forwardMessages", p, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// CopyMessageParams holds the parameters for copyMessage.
type CopyMessageParams struct {
	ChatID                  ChatID                   `json:"chat_id"`
	MessageThreadID         *int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   *int64                   `json:"direct_messages_topic_id,omitempty"`
	FromChatID              ChatID                   `json:"from_chat_id"`
	MessageID               int64                    `json:"message_id"`
	VideoStartTimestamp     *int64                   `json:"video_start_timestamp,omitempty"`
	Caption                 *string                  `json:"caption,omitempty"`
	ParseMode               *string                  `json:"parse_mode,omitempty"`
	CaptionEntities         *[]MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   *bool                    `json:"show_caption_above_media,omitempty"`
	DisableNotification     *bool                    `json:"disable_notification,omitempty"`
	ProtectContent          *bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      *bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         *string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             ReplyMarkup              `json:"reply_markup,omitempty"`
}

// CopyMessage copies messages of any kind. Returns the MessageId of the sent
// message on success.
func (c *Client) CopyMessage(ctx context.Context, p *CopyMessageParams) (*MessageId, error) {
	var id MessageId
	if err := c.Invoke(ctx, "copyMessage", p, &id); err != nil {
		return nil, err
	}
	return &id, nil
}

// CopyMessagesParams holds the parameters for copyMessages.
type CopyMessagesParams struct {
	ChatID                ChatID  `json:"chat_id"`
	MessageThreadID       *int64  `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID *int64  `json:"direct_messages_topic_id,omitempty"`
	FromChatID            ChatID  `json:"from_chat_id"`
	MessageIDs            []int64 `json:"message_ids"`
	DisableNotification   *bool   `json:"disable_notification,omitempty"`
	ProtectContent        *bool   `json:"protect_content,omitempty"`
	RemoveCaption         *bool   `json:"remove_caption,omitempty"`
}

// CopyMessages copies multiple messages of any kind. On success, an Array of
// MessageId of the sent messages is returned.
func (c *Client) CopyMessages(ctx context.Context, p *CopyMessagesParams) ([]MessageId, error) {
	var ids []MessageId
	if err := c.Invoke(ctx, "copyMessages", p, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// SendPhotoParams holds the parameters for sendPhoto.
type SendPhotoParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Photo                      InputFile                   `json:"photo"`
	Caption                    *string                     `json:"caption,omitempty"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	CaptionEntities            *[]MessageEntity            `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia      *bool                       `json:"show_caption_above_media,omitempty"`
	HasSpoiler                 *bool                       `json:"has_spoiler,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendPhoto sends photos. On success, the sent Message is returned.
func (c *Client) SendPhoto(ctx context.Context, p *SendPhotoParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendPhoto", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendLivePhotoParams holds the parameters for sendLivePhoto.
type SendLivePhotoParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	LivePhoto                  InputFile                   `json:"live_photo"`
	Photo                      InputFile                   `json:"photo"`
	Caption                    *string                     `json:"caption,omitempty"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	CaptionEntities            *[]MessageEntity            `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia      *bool                       `json:"show_caption_above_media,omitempty"`
	HasSpoiler                 *bool                       `json:"has_spoiler,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendLivePhoto sends live photos. On success, the sent Message is returned.
func (c *Client) SendLivePhoto(ctx context.Context, p *SendLivePhotoParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendLivePhoto", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendAudioParams holds the parameters for sendAudio.
type SendAudioParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Audio                      InputFile                   `json:"audio"`
	Caption                    *string                     `json:"caption,omitempty"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	CaptionEntities            *[]MessageEntity            `json:"caption_entities,omitempty"`
	Duration                   *int                        `json:"duration,omitempty"`
	Performer                  *string                     `json:"performer,omitempty"`
	Title                      *string                     `json:"title,omitempty"`
	Thumbnail                  *InputFile                  `json:"thumbnail,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendAudio sends audio files to be displayed by Telegram clients in the music
// player. On success, the sent Message is returned.
func (c *Client) SendAudio(ctx context.Context, p *SendAudioParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendAudio", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendDocumentParams holds the parameters for sendDocument.
type SendDocumentParams struct {
	BusinessConnectionID        *string                     `json:"business_connection_id,omitempty"`
	ChatID                      ChatID                      `json:"chat_id"`
	MessageThreadID             *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID       *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters  *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Document                    InputFile                   `json:"document"`
	Thumbnail                   *InputFile                  `json:"thumbnail,omitempty"`
	Caption                     *string                     `json:"caption,omitempty"`
	ParseMode                   *string                     `json:"parse_mode,omitempty"`
	CaptionEntities             *[]MessageEntity            `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool                       `json:"disable_content_type_detection,omitempty"`
	DisableNotification         *bool                       `json:"disable_notification,omitempty"`
	ProtectContent              *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast          *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID             *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters     *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters             *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                 ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendDocument sends general files. On success, the sent Message is returned.
func (c *Client) SendDocument(ctx context.Context, p *SendDocumentParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendDocument", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendVideoParams holds the parameters for sendVideo.
type SendVideoParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Video                      InputFile                   `json:"video"`
	Duration                   *int                        `json:"duration,omitempty"`
	Width                      *int                        `json:"width,omitempty"`
	Height                     *int                        `json:"height,omitempty"`
	Thumbnail                  *InputFile                  `json:"thumbnail,omitempty"`
	Cover                      *InputFile                  `json:"cover,omitempty"`
	StartTimestamp             *int64                      `json:"start_timestamp,omitempty"`
	Caption                    *string                     `json:"caption,omitempty"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	CaptionEntities            *[]MessageEntity            `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia      *bool                       `json:"show_caption_above_media,omitempty"`
	HasSpoiler                 *bool                       `json:"has_spoiler,omitempty"`
	SupportsStreaming          *bool                       `json:"supports_streaming,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendVideo sends video files. On success, the sent Message is returned.
func (c *Client) SendVideo(ctx context.Context, p *SendVideoParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendVideo", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendAnimationParams holds the parameters for sendAnimation.
type SendAnimationParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Animation                  InputFile                   `json:"animation"`
	Duration                   *int                        `json:"duration,omitempty"`
	Width                      *int                        `json:"width,omitempty"`
	Height                     *int                        `json:"height,omitempty"`
	Thumbnail                  *InputFile                  `json:"thumbnail,omitempty"`
	Caption                    *string                     `json:"caption,omitempty"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	CaptionEntities            *[]MessageEntity            `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia      *bool                       `json:"show_caption_above_media,omitempty"`
	HasSpoiler                 *bool                       `json:"has_spoiler,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendAnimation sends animation files (GIF or H.264/MPEG-4 AVC video without
// sound). On success, the sent Message is returned.
func (c *Client) SendAnimation(ctx context.Context, p *SendAnimationParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendAnimation", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendVoiceParams holds the parameters for sendVoice.
type SendVoiceParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Voice                      InputFile                   `json:"voice"`
	Caption                    *string                     `json:"caption,omitempty"`
	ParseMode                  *string                     `json:"parse_mode,omitempty"`
	CaptionEntities            *[]MessageEntity            `json:"caption_entities,omitempty"`
	Duration                   *int                        `json:"duration,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendVoice sends audio files displayed as a playable voice message. On
// success, the sent Message is returned.
func (c *Client) SendVoice(ctx context.Context, p *SendVoiceParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendVoice", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendVideoNoteParams holds the parameters for sendVideoNote.
type SendVideoNoteParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	VideoNote                  InputFile                   `json:"video_note"`
	Duration                   *int                        `json:"duration,omitempty"`
	Length                     *int                        `json:"length,omitempty"`
	Thumbnail                  *InputFile                  `json:"thumbnail,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendVideoNote sends a rounded square MPEG4 video of up to 1 minute long. On
// success, the sent Message is returned.
func (c *Client) SendVideoNote(ctx context.Context, p *SendVideoNoteParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendVideoNote", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendPaidMediaParams holds the parameters for sendPaidMedia.
type SendPaidMediaParams struct {
	BusinessConnectionID    *string                  `json:"business_connection_id,omitempty"`
	ChatID                  ChatID                   `json:"chat_id"`
	MessageThreadID         *int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   *int64                   `json:"direct_messages_topic_id,omitempty"`
	StarCount               int64                    `json:"star_count"`
	Media                   []InputPaidMedia         `json:"media"`
	Payload                 *string                  `json:"payload,omitempty"`
	Caption                 *string                  `json:"caption,omitempty"`
	ParseMode               *string                  `json:"parse_mode,omitempty"`
	CaptionEntities         *[]MessageEntity         `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   *bool                    `json:"show_caption_above_media,omitempty"`
	DisableNotification     *bool                    `json:"disable_notification,omitempty"`
	ProtectContent          *bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      *bool                    `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendPaidMedia sends paid media. On success, the sent Message is returned.
func (c *Client) SendPaidMedia(ctx context.Context, p *SendPaidMediaParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendPaidMedia", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendMediaGroupParams holds the parameters for sendMediaGroup.
type SendMediaGroupParams struct {
	BusinessConnectionID  *string          `json:"business_connection_id,omitempty"`
	ChatID                ChatID           `json:"chat_id"`
	MessageThreadID       *int64           `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID *int64           `json:"direct_messages_topic_id,omitempty"`
	Media                 []InputMedia     `json:"media"`
	DisableNotification   *bool            `json:"disable_notification,omitempty"`
	ProtectContent        *bool            `json:"protect_content,omitempty"`
	AllowPaidBroadcast    *bool            `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID       *string          `json:"message_effect_id,omitempty"`
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`
}

// SendMediaGroup sends a group of photos, live photos, videos, documents or
// audios as an album. On success, an Array of Message objects that were sent
// is returned.
func (c *Client) SendMediaGroup(ctx context.Context, p *SendMediaGroupParams) ([]Message, error) {
	var ms []Message
	if err := c.Invoke(ctx, "sendMediaGroup", p, &ms); err != nil {
		return nil, err
	}
	return ms, nil
}

// SendLocationParams holds the parameters for sendLocation.
type SendLocationParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Latitude                   float64                     `json:"latitude"`
	Longitude                  float64                     `json:"longitude"`
	HorizontalAccuracy         *float64                    `json:"horizontal_accuracy,omitempty"`
	LivePeriod                 *int                        `json:"live_period,omitempty"`
	Heading                    *int                        `json:"heading,omitempty"`
	ProximityAlertRadius       *int                        `json:"proximity_alert_radius,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendLocation sends a point on the map. On success, the sent Message is returned.
func (c *Client) SendLocation(ctx context.Context, p *SendLocationParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendLocation", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendVenueParams holds the parameters for sendVenue.
type SendVenueParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Latitude                   float64                     `json:"latitude"`
	Longitude                  float64                     `json:"longitude"`
	Title                      string                      `json:"title"`
	Address                    string                      `json:"address"`
	FoursquareID               *string                     `json:"foursquare_id,omitempty"`
	FoursquareType             *string                     `json:"foursquare_type,omitempty"`
	GooglePlaceID              *string                     `json:"google_place_id,omitempty"`
	GooglePlaceType            *string                     `json:"google_place_type,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendVenue sends information about a venue. On success, the sent Message is
// returned.
func (c *Client) SendVenue(ctx context.Context, p *SendVenueParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendVenue", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendContactParams holds the parameters for sendContact.
type SendContactParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	PhoneNumber                string                      `json:"phone_number"`
	FirstName                  string                      `json:"first_name"`
	LastName                   *string                     `json:"last_name,omitempty"`
	VCard                      *string                     `json:"vcard,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendContact sends phone contacts. On success, the sent Message is returned.
func (c *Client) SendContact(ctx context.Context, p *SendContactParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendContact", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendPollParams holds the parameters for sendPoll.
type SendPollParams struct {
	BusinessConnectionID   *string           `json:"business_connection_id,omitempty"`
	ChatID                 ChatID            `json:"chat_id"`
	MessageThreadID        *int64            `json:"message_thread_id,omitempty"`
	Question               string            `json:"question"`
	QuestionParseMode      *string           `json:"question_parse_mode,omitempty"`
	QuestionEntities       *[]MessageEntity  `json:"question_entities,omitempty"`
	Options                []InputPollOption `json:"options"`
	IsAnonymous            *bool             `json:"is_anonymous,omitempty"`
	Type                   *string           `json:"type,omitempty"`
	AllowsMultipleAnswers  *bool             `json:"allows_multiple_answers,omitempty"`
	AllowsRevoting         *bool             `json:"allows_revoting,omitempty"`
	ShuffleOptions         *bool             `json:"shuffle_options,omitempty"`
	AllowAddingOptions     *bool             `json:"allow_adding_options,omitempty"`
	HideResultsUntilCloses *bool             `json:"hide_results_until_closes,omitempty"`
	MembersOnly            *bool             `json:"members_only,omitempty"`
	CountryCodes           *[]string         `json:"country_codes,omitempty"`
	CorrectOptionIDs       *[]int            `json:"correct_option_ids,omitempty"`
	Explanation            *string           `json:"explanation,omitempty"`
	ExplanationParseMode   *string           `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities    *[]MessageEntity  `json:"explanation_entities,omitempty"`
	ExplanationMedia       InputPollMedia    `json:"explanation_media,omitempty"`
	OpenPeriod             *int              `json:"open_period,omitempty"`
	CloseDate              *int64            `json:"close_date,omitempty"`
	IsClosed               *bool             `json:"is_closed,omitempty"`
	Description            *string           `json:"description,omitempty"`
	DescriptionParseMode   *string           `json:"description_parse_mode,omitempty"`
	DescriptionEntities    *[]MessageEntity  `json:"description_entities,omitempty"`
	Media                  InputPollMedia    `json:"media,omitempty"`
	DisableNotification    *bool             `json:"disable_notification,omitempty"`
	ProtectContent         *bool             `json:"protect_content,omitempty"`
	AllowPaidBroadcast     *bool             `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID        *string           `json:"message_effect_id,omitempty"`
	ReplyParameters        *ReplyParameters  `json:"reply_parameters,omitempty"`
	ReplyMarkup            ReplyMarkup       `json:"reply_markup,omitempty"`
}

// SendPoll sends a native poll. On success, the sent Message is returned.
func (c *Client) SendPoll(ctx context.Context, p *SendPollParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendPoll", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendChecklistParams holds the parameters for sendChecklist.
type SendChecklistParams struct {
	BusinessConnectionID string           `json:"business_connection_id"`
	ChatID               ChatID           `json:"chat_id"`
	Checklist            InputChecklist   `json:"checklist"`
	DisableNotification  *bool            `json:"disable_notification,omitempty"`
	ProtectContent       *bool            `json:"protect_content,omitempty"`
	MessageEffectID      *string          `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`
}

// SendChecklist sends a checklist on behalf of a connected business account.
// On success, the sent Message is returned.
func (c *Client) SendChecklist(ctx context.Context, p *SendChecklistParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendChecklist", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendDiceParams holds the parameters for sendDice.
type SendDiceParams struct {
	BusinessConnectionID    *string                  `json:"business_connection_id,omitempty"`
	ChatID                  ChatID                   `json:"chat_id"`
	MessageThreadID         *int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID   *int64                   `json:"direct_messages_topic_id,omitempty"`
	Emoji                   *string                  `json:"emoji,omitempty"`
	DisableNotification     *bool                    `json:"disable_notification,omitempty"`
	ProtectContent          *bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast      *bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID         *string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             ReplyMarkup              `json:"reply_markup,omitempty"`
}

// SendDice sends an animated emoji that will display a random value. On
// success, the sent Message is returned.
func (c *Client) SendDice(ctx context.Context, p *SendDiceParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendDice", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendMessageDraftParams holds the parameters for sendMessageDraft.
type SendMessageDraftParams struct {
	ChatID          ChatID           `json:"chat_id"`
	MessageThreadID *int64           `json:"message_thread_id,omitempty"`
	DraftID         int64            `json:"draft_id"`
	Text            *string          `json:"text,omitempty"`
	ParseMode       *string          `json:"parse_mode,omitempty"`
	Entities        *[]MessageEntity `json:"entities,omitempty"`
	CanStop         *bool            `json:"can_stop,omitempty"`
	KeepOnStop      *bool            `json:"keep_on_stop,omitempty"`
}

// SendMessageDraft streams a partial message to a user while the message is
// being generated. Returns True on success.
func (c *Client) SendMessageDraft(ctx context.Context, p *SendMessageDraftParams) error {
	return c.Invoke(ctx, "sendMessageDraft", p, nil)
}

// SendChatActionParams holds the parameters for sendChatAction.
type SendChatActionParams struct {
	BusinessConnectionID *string `json:"business_connection_id,omitempty"`
	ChatID               ChatID  `json:"chat_id"`
	MessageThreadID      *int64  `json:"message_thread_id,omitempty"`
	Action               string  `json:"action"`
}

// SendChatAction tells the user that something is happening on the bot's side.
// Returns True on success.
func (c *Client) SendChatAction(ctx context.Context, p *SendChatActionParams) error {
	return c.Invoke(ctx, "sendChatAction", p, nil)
}

// SetMessageReactionParams holds the parameters for setMessageReaction.
type SetMessageReactionParams struct {
	ChatID    ChatID          `json:"chat_id"`
	MessageID int64           `json:"message_id"`
	Reaction  *[]ReactionType `json:"reaction,omitempty"`
	IsBig     *bool           `json:"is_big,omitempty"`
}

// SetMessageReaction changes the chosen reactions on a message. Returns True
// on success.
func (c *Client) SetMessageReaction(ctx context.Context, p *SetMessageReactionParams) error {
	return c.Invoke(ctx, "setMessageReaction", p, nil)
}

// DeleteMessageReactionParams holds the parameters for deleteMessageReaction.
type DeleteMessageReactionParams struct {
	ChatID      ChatID `json:"chat_id"`
	MessageID   int64  `json:"message_id"`
	UserID      *int64 `json:"user_id,omitempty"`
	ActorChatID *int64 `json:"actor_chat_id,omitempty"`
}

// DeleteMessageReaction removes a reaction from a message in a group or a
// supergroup chat. Returns True on success.
func (c *Client) DeleteMessageReaction(ctx context.Context, p *DeleteMessageReactionParams) error {
	return c.Invoke(ctx, "deleteMessageReaction", p, nil)
}

// DeleteAllMessageReactionsParams holds the parameters for deleteAllMessageReactions.
type DeleteAllMessageReactionsParams struct {
	ChatID      ChatID `json:"chat_id"`
	UserID      *int64 `json:"user_id,omitempty"`
	ActorChatID *int64 `json:"actor_chat_id,omitempty"`
}

// DeleteAllMessageReactions removes up to 10000 recent reactions in a group
// or a supergroup chat added by a given user or chat. Returns True on success.
func (c *Client) DeleteAllMessageReactions(ctx context.Context, p *DeleteAllMessageReactionsParams) error {
	return c.Invoke(ctx, "deleteAllMessageReactions", p, nil)
}

// ApproveSuggestedPostParams holds the parameters for approveSuggestedPost.
type ApproveSuggestedPostParams struct {
	ChatID    ChatID `json:"chat_id"`
	MessageID int64  `json:"message_id"`
	SendDate  *int64 `json:"send_date,omitempty"`
}

// ApproveSuggestedPost approves a suggested post in a direct messages chat.
// Returns True on success.
func (c *Client) ApproveSuggestedPost(ctx context.Context, p *ApproveSuggestedPostParams) error {
	return c.Invoke(ctx, "approveSuggestedPost", p, nil)
}

// DeclineSuggestedPostParams holds the parameters for declineSuggestedPost.
type DeclineSuggestedPostParams struct {
	ChatID    ChatID  `json:"chat_id"`
	MessageID int64   `json:"message_id"`
	Comment   *string `json:"comment,omitempty"`
}

// DeclineSuggestedPost declines a suggested post in a direct messages chat.
// Returns True on success.
func (c *Client) DeclineSuggestedPost(ctx context.Context, p *DeclineSuggestedPostParams) error {
	return c.Invoke(ctx, "declineSuggestedPost", p, nil)
}

// DeleteMessageParams holds the parameters for deleteMessage.
type DeleteMessageParams struct {
	ChatID    ChatID `json:"chat_id"`
	MessageID int64  `json:"message_id"`
}

// DeleteMessage deletes a message, including service messages. Returns True
// on success.
func (c *Client) DeleteMessage(ctx context.Context, p *DeleteMessageParams) error {
	return c.Invoke(ctx, "deleteMessage", p, nil)
}

// DeleteMessagesParams holds the parameters for deleteMessages.
type DeleteMessagesParams struct {
	ChatID     ChatID  `json:"chat_id"`
	MessageIDs []int64 `json:"message_ids"`
}

// DeleteMessages deletes multiple messages simultaneously. Returns True on
// success.
func (c *Client) DeleteMessages(ctx context.Context, p *DeleteMessagesParams) error {
	return c.Invoke(ctx, "deleteMessages", p, nil)
}

// DeleteEphemeralMessageParams holds the parameters for deleteEphemeralMessage.
type DeleteEphemeralMessageParams struct {
	ChatID             ChatID `json:"chat_id"`
	ReceiverUserID     int64  `json:"receiver_user_id"`
	EphemeralMessageID int64  `json:"ephemeral_message_id"`
}

// DeleteEphemeralMessage deletes an ephemeral message. Returns True on success.
func (c *Client) DeleteEphemeralMessage(ctx context.Context, p *DeleteEphemeralMessageParams) error {
	return c.Invoke(ctx, "deleteEphemeralMessage", p, nil)
}

// EditMessageTextParams holds parameters for editMessageText.
type EditMessageTextParams struct {
	BusinessConnectionID *string             `json:"business_connection_id,omitempty"`
	ChatID               *ChatID             `json:"chat_id,omitempty"`
	MessageID            *int64              `json:"message_id,omitempty"`
	InlineMessageID      *string             `json:"inline_message_id,omitempty"`
	Text                 *string             `json:"text,omitempty"`
	ParseMode            *string             `json:"parse_mode,omitempty"`
	Entities             []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	RichMessage          *InputRichMessage   `json:"rich_message,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
}

// EditMessageText edits text, rich and game messages. Returns the edited
// Message, or nil if the edited message is an inline message (Telegram
// returns True).
func (c *Client) EditMessageText(ctx context.Context, p *EditMessageTextParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "editMessageText", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode editMessageText result: %w", err)
	}
	return &m, nil
}

// EditMessageCaptionParams holds parameters for editMessageCaption.
type EditMessageCaptionParams struct {
	BusinessConnectionID  *string         `json:"business_connection_id,omitempty"`
	ChatID                *ChatID         `json:"chat_id,omitempty"`
	MessageID             *int64          `json:"message_id,omitempty"`
	InlineMessageID       *string         `json:"inline_message_id,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup     `json:"reply_markup,omitempty"`
}

// EditMessageCaption edits captions of messages. Returns the edited Message,
// or nil if the edited message is an inline message (Telegram returns True).
func (c *Client) EditMessageCaption(ctx context.Context, p *EditMessageCaptionParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "editMessageCaption", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode editMessageCaption result: %w", err)
	}
	return &m, nil
}

// EditMessageMediaParams holds parameters for editMessageMedia.
type EditMessageMediaParams struct {
	BusinessConnectionID *string     `json:"business_connection_id,omitempty"`
	ChatID               *ChatID     `json:"chat_id,omitempty"`
	MessageID            *int64      `json:"message_id,omitempty"`
	InlineMessageID      *string     `json:"inline_message_id,omitempty"`
	Media                InputMedia  `json:"media"`
	ReplyMarkup          ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMedia edits animation, audio, document, live photo, photo, or
// video messages, or replaces a text or rich message with a media. Returns
// the edited Message, or nil if the edited message is an inline message
// (Telegram returns True).
func (c *Client) EditMessageMedia(ctx context.Context, p *EditMessageMediaParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "editMessageMedia", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode editMessageMedia result: %w", err)
	}
	return &m, nil
}

// EditMessageLiveLocationParams holds parameters for editMessageLiveLocation.
type EditMessageLiveLocationParams struct {
	BusinessConnectionID *string     `json:"business_connection_id,omitempty"`
	ChatID               *ChatID     `json:"chat_id,omitempty"`
	MessageID            *int64      `json:"message_id,omitempty"`
	InlineMessageID      *string     `json:"inline_message_id,omitempty"`
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	LivePeriod           *int        `json:"live_period,omitempty"`
	HorizontalAccuracy   *float64    `json:"horizontal_accuracy,omitempty"`
	Heading              *int        `json:"heading,omitempty"`
	ProximityAlertRadius *int        `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation edits live location messages. Returns the edited
// Message, or nil if the edited message is an inline message (Telegram
// returns True).
func (c *Client) EditMessageLiveLocation(ctx context.Context, p *EditMessageLiveLocationParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "editMessageLiveLocation", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode editMessageLiveLocation result: %w", err)
	}
	return &m, nil
}

// EditMessageReplyMarkupParams holds parameters for editMessageReplyMarkup.
type EditMessageReplyMarkupParams struct {
	BusinessConnectionID *string     `json:"business_connection_id,omitempty"`
	ChatID               *ChatID     `json:"chat_id,omitempty"`
	MessageID            *int64      `json:"message_id,omitempty"`
	InlineMessageID      *string     `json:"inline_message_id,omitempty"`
	ReplyMarkup          ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup edits only the reply markup of messages. Returns
// the edited Message, or nil if the edited message is an inline message
// (Telegram returns True).
func (c *Client) EditMessageReplyMarkup(ctx context.Context, p *EditMessageReplyMarkupParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "editMessageReplyMarkup", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode editMessageReplyMarkup result: %w", err)
	}
	return &m, nil
}

// StopMessageLiveLocationParams holds parameters for stopMessageLiveLocation.
type StopMessageLiveLocationParams struct {
	BusinessConnectionID *string     `json:"business_connection_id,omitempty"`
	ChatID               *ChatID     `json:"chat_id,omitempty"`
	MessageID            *int64      `json:"message_id,omitempty"`
	InlineMessageID      *string     `json:"inline_message_id,omitempty"`
	ReplyMarkup          ReplyMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation stops updating a live location message before
// live_period expires. Returns the edited Message, or nil if the message is
// an inline message (Telegram returns True).
func (c *Client) StopMessageLiveLocation(ctx context.Context, p *StopMessageLiveLocationParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "stopMessageLiveLocation", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode stopMessageLiveLocation result: %w", err)
	}
	return &m, nil
}

// EditMessageChecklistParams holds parameters for editMessageChecklist.
type EditMessageChecklistParams struct {
	BusinessConnectionID string         `json:"business_connection_id"`
	ChatID               ChatID         `json:"chat_id"`
	MessageID            int64          `json:"message_id"`
	Checklist            InputChecklist `json:"checklist"`
	ReplyMarkup          ReplyMarkup    `json:"reply_markup,omitempty"`
}

// EditMessageChecklist edits a checklist on behalf of a connected business
// account.
func (c *Client) EditMessageChecklist(ctx context.Context, p *EditMessageChecklistParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "editMessageChecklist", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// StopPollParams holds parameters for stopPoll.
type StopPollParams struct {
	BusinessConnectionID *string     `json:"business_connection_id,omitempty"`
	ChatID               ChatID      `json:"chat_id"`
	MessageID            int64       `json:"message_id"`
	ReplyMarkup          ReplyMarkup `json:"reply_markup,omitempty"`
}

// StopPoll stops a poll which was sent by the bot.
func (c *Client) StopPoll(ctx context.Context, p *StopPollParams) (*Poll, error) {
	var m Poll
	if err := c.Invoke(ctx, "stopPoll", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// EditEphemeralMessageTextParams holds parameters for
// editEphemeralMessageText.
type EditEphemeralMessageTextParams struct {
	ChatID             ChatID              `json:"chat_id"`
	ReceiverUserID     int64               `json:"receiver_user_id"`
	EphemeralMessageID int64               `json:"ephemeral_message_id"`
	Text               *string             `json:"text,omitempty"`
	ParseMode          *string             `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	RichMessage        *InputRichMessage   `json:"rich_message,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	ReplyMarkup        ReplyMarkup         `json:"reply_markup,omitempty"`
}

// EditEphemeralMessageText edits an ephemeral text or rich message. It is
// not guaranteed that the user will receive the message edit event,
// especially if they are offline.
func (c *Client) EditEphemeralMessageText(ctx context.Context, p *EditEphemeralMessageTextParams) error {
	return c.Invoke(ctx, "editEphemeralMessageText", p, nil)
}

// EditEphemeralMessageMediaParams holds parameters for
// editEphemeralMessageMedia.
type EditEphemeralMessageMediaParams struct {
	ChatID             ChatID      `json:"chat_id"`
	ReceiverUserID     int64       `json:"receiver_user_id"`
	EphemeralMessageID int64       `json:"ephemeral_message_id"`
	Media              InputMedia  `json:"media"`
	ReplyMarkup        ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditEphemeralMessageMedia edits the media of an ephemeral message. It is
// not guaranteed that the user will receive the message edit event,
// especially if they are offline.
func (c *Client) EditEphemeralMessageMedia(ctx context.Context, p *EditEphemeralMessageMediaParams) error {
	return c.Invoke(ctx, "editEphemeralMessageMedia", p, nil)
}

// EditEphemeralMessageCaptionParams holds parameters for
// editEphemeralMessageCaption.
type EditEphemeralMessageCaptionParams struct {
	ChatID                ChatID          `json:"chat_id"`
	ReceiverUserID        int64           `json:"receiver_user_id"`
	EphemeralMessageID    int64           `json:"ephemeral_message_id"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           ReplyMarkup     `json:"reply_markup,omitempty"`
}

// EditEphemeralMessageCaption edits the caption of an ephemeral message. It
// is not guaranteed that the user will receive the message edit event,
// especially if they are offline.
func (c *Client) EditEphemeralMessageCaption(ctx context.Context, p *EditEphemeralMessageCaptionParams) error {
	return c.Invoke(ctx, "editEphemeralMessageCaption", p, nil)
}

// EditEphemeralMessageReplyMarkupParams holds parameters for
// editEphemeralMessageReplyMarkup.
type EditEphemeralMessageReplyMarkupParams struct {
	ChatID             ChatID      `json:"chat_id"`
	ReceiverUserID     int64       `json:"receiver_user_id"`
	EphemeralMessageID int64       `json:"ephemeral_message_id"`
	ReplyMarkup        ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditEphemeralMessageReplyMarkup edits only the reply markup of an
// ephemeral message. It is not guaranteed that the user will receive the
// message edit event, especially if they are offline.
func (c *Client) EditEphemeralMessageReplyMarkup(ctx context.Context, p *EditEphemeralMessageReplyMarkupParams) error {
	return c.Invoke(ctx, "editEphemeralMessageReplyMarkup", p, nil)
}

// GetUserProfilePhotosParams holds the parameters for the getUserProfilePhotos method.
type GetUserProfilePhotosParams struct {
	UserID int64 `json:"user_id"`
	Offset *int  `json:"offset,omitempty"`
	Limit  *int  `json:"limit,omitempty"`
}

// GetUserProfilePhotos returns a list of profile pictures for a user.
func (c *Client) GetUserProfilePhotos(ctx context.Context, p *GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	var photos UserProfilePhotos
	if err := c.Invoke(ctx, "getUserProfilePhotos", p, &photos); err != nil {
		return nil, err
	}
	return &photos, nil
}

// GetUserProfileAudiosParams holds the parameters for the getUserProfileAudios method.
type GetUserProfileAudiosParams struct {
	UserID int64 `json:"user_id"`
	Offset *int  `json:"offset,omitempty"`
	Limit  *int  `json:"limit,omitempty"`
}

// GetUserProfileAudios returns a list of profile audios for a user.
func (c *Client) GetUserProfileAudios(ctx context.Context, p *GetUserProfileAudiosParams) (*UserProfileAudios, error) {
	var audios UserProfileAudios
	if err := c.Invoke(ctx, "getUserProfileAudios", p, &audios); err != nil {
		return nil, err
	}
	return &audios, nil
}

// SetUserEmojiStatusParams holds the parameters for the setUserEmojiStatus method.
type SetUserEmojiStatusParams struct {
	UserID                    int64   `json:"user_id"`
	EmojiStatusCustomEmojiID  *string `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate *int64  `json:"emoji_status_expiration_date,omitempty"`
}

// SetUserEmojiStatus changes the emoji status for a given user that previously allowed the bot to manage it.
func (c *Client) SetUserEmojiStatus(ctx context.Context, p *SetUserEmojiStatusParams) error {
	return c.Invoke(ctx, "setUserEmojiStatus", p, nil)
}

// GetFileParams holds the parameters for the getFile method.
type GetFileParams struct {
	FileID string `json:"file_id"`
}

// GetFile returns basic information about a file and prepares it for downloading.
func (c *Client) GetFile(ctx context.Context, p *GetFileParams) (*File, error) {
	var f File
	if err := c.Invoke(ctx, "getFile", p, &f); err != nil {
		return nil, err
	}
	return &f, nil
}

// AnswerCallbackQueryParams holds parameters for answerCallbackQuery.
type AnswerCallbackQueryParams struct {
	CallbackQueryID string  `json:"callback_query_id"`
	Text            *string `json:"text,omitempty"`
	ShowAlert       *bool   `json:"show_alert,omitempty"`
	URL             *string `json:"url,omitempty"`
	CacheTime       *int    `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery sends answers to callback queries sent from inline
// keyboards. The answer is displayed to the user as a notification at the
// top of the chat screen or as an alert.
func (c *Client) AnswerCallbackQuery(ctx context.Context, p *AnswerCallbackQueryParams) error {
	return c.Invoke(ctx, "answerCallbackQuery", p, nil)
}

// AnswerGuestQueryParams holds parameters for answerGuestQuery.
type AnswerGuestQueryParams struct {
	GuestQueryID string            `json:"guest_query_id"`
	Result       InlineQueryResult `json:"result"`
}

// AnswerGuestQuery replies to a received guest message.
func (c *Client) AnswerGuestQuery(ctx context.Context, p *AnswerGuestQueryParams) (*SentGuestMessage, error) {
	var m SentGuestMessage
	if err := c.Invoke(ctx, "answerGuestQuery", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// AnswerWebAppQueryParams holds parameters for answerWebAppQuery.
type AnswerWebAppQueryParams struct {
	WebAppQueryID string            `json:"web_app_query_id"`
	Result        InlineQueryResult `json:"result"`
}

// AnswerWebAppQuery sets the result of an interaction with a Web App and
// sends a corresponding message on behalf of the user to the chat from
// which the query originated.
func (c *Client) AnswerWebAppQuery(ctx context.Context, p *AnswerWebAppQueryParams) (*SentWebAppMessage, error) {
	var m SentWebAppMessage
	if err := c.Invoke(ctx, "answerWebAppQuery", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SavePreparedInlineMessageParams holds parameters for
// savePreparedInlineMessage.
type SavePreparedInlineMessageParams struct {
	UserID            int64             `json:"user_id"`
	Result            InlineQueryResult `json:"result"`
	AllowUserChats    *bool             `json:"allow_user_chats,omitempty"`
	AllowBotChats     *bool             `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   *bool             `json:"allow_group_chats,omitempty"`
	AllowChannelChats *bool             `json:"allow_channel_chats,omitempty"`
}

// SavePreparedInlineMessage stores a message that can be sent by a user of
// a Mini App.
func (c *Client) SavePreparedInlineMessage(ctx context.Context, p *SavePreparedInlineMessageParams) (*PreparedInlineMessage, error) {
	var m PreparedInlineMessage
	if err := c.Invoke(ctx, "savePreparedInlineMessage", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SavePreparedKeyboardButtonParams holds parameters for
// savePreparedKeyboardButton.
type SavePreparedKeyboardButtonParams struct {
	UserID int64          `json:"user_id"`
	Button KeyboardButton `json:"button"`
}

// SavePreparedKeyboardButton stores a keyboard button that can be used by a
// user within a Mini App.
func (c *Client) SavePreparedKeyboardButton(ctx context.Context, p *SavePreparedKeyboardButtonParams) (*PreparedKeyboardButton, error) {
	var m PreparedKeyboardButton
	if err := c.Invoke(ctx, "savePreparedKeyboardButton", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
