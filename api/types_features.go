package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GiftBackground describes the background of a gift.
type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

// Gift represents a gift that can be sent by the bot.
type Gift struct {
	ID                     string          `json:"id"`
	Sticker                Sticker         `json:"sticker"`
	StarCount              int64           `json:"star_count"`
	UpgradeStarCount       *int64          `json:"upgrade_star_count"`
	IsPremium              *bool           `json:"is_premium"`
	HasColors              *bool           `json:"has_colors"`
	TotalCount             *int            `json:"total_count"`
	RemainingCount         *int            `json:"remaining_count"`
	PersonalTotalCount     *int            `json:"personal_total_count"`
	PersonalRemainingCount *int            `json:"personal_remaining_count"`
	Background             *GiftBackground `json:"background"`
	UniqueGiftVariantCount *int            `json:"unique_gift_variant_count"`
	PublisherChat          *Chat           `json:"publisher_chat"`
}

// Gifts represents a list of gifts.
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// UniqueGiftModel describes the model of a unique gift.
type UniqueGiftModel struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
	Rarity         *string `json:"rarity"`
}

// UniqueGiftSymbol describes the symbol shown on the pattern of a unique gift.
type UniqueGiftSymbol struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors describes the colors of the backdrop of a unique gift.
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

// UniqueGiftBackdrop describes the backdrop of a unique gift.
type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

// UniqueGiftColors contains information about the color scheme for a user's
// name, message replies and link previews based on a unique gift.
type UniqueGiftColors struct {
	ModelCustomEmojiID    string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor   int    `json:"light_theme_main_color"`
	LightThemeOtherColors []int  `json:"light_theme_other_colors"`
	DarkThemeMainColor    int    `json:"dark_theme_main_color"`
	DarkThemeOtherColors  []int  `json:"dark_theme_other_colors"`
}

// UniqueGift describes a unique gift that was upgraded from a regular gift.
type UniqueGift struct {
	GiftID           string             `json:"gift_id"`
	BaseName         string             `json:"base_name"`
	Name             string             `json:"name"`
	Number           int                `json:"number"`
	Model            UniqueGiftModel    `json:"model"`
	Symbol           UniqueGiftSymbol   `json:"symbol"`
	Backdrop         UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        *bool              `json:"is_premium"`
	IsBurned         *bool              `json:"is_burned"`
	IsFromBlockchain *bool              `json:"is_from_blockchain"`
	Colors           *UniqueGiftColors  `json:"colors"`
	PublisherChat    *Chat              `json:"publisher_chat"`
}

// GiftInfo describes a service message about a regular gift that was sent or received.
type GiftInfo struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             *string         `json:"owned_gift_id"`
	ConvertStarCount        *int64          `json:"convert_star_count"`
	PrepaidUpgradeStarCount *int64          `json:"prepaid_upgrade_star_count"`
	IsUpgradeSeparate       *bool           `json:"is_upgrade_separate"`
	CanBeUpgraded           *bool           `json:"can_be_upgraded"`
	Text                    *string         `json:"text"`
	Entities                []MessageEntity `json:"entities"`
	IsPrivate               *bool           `json:"is_private"`
	UniqueGiftNumber        *int            `json:"unique_gift_number"`
}

// UniqueGiftInfo describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	Gift               UniqueGift      `json:"gift"`
	Origin             string          `json:"origin"`
	Text               *string         `json:"text"`
	Entities           []MessageEntity `json:"entities"`
	IsPrivate          *bool           `json:"is_private"`
	LastResaleCurrency *string         `json:"last_resale_currency"`
	LastResaleAmount   *int64          `json:"last_resale_amount"`
	OwnedGiftID        *string         `json:"owned_gift_id"`
	TransferStarCount  *int64          `json:"transfer_star_count"`
	NextTransferDate   *int64          `json:"next_transfer_date"`
}

// OwnedGift describes a gift received and owned by a user or a chat.
// Currently, it can be one of OwnedGiftRegular or OwnedGiftUnique.
type OwnedGift interface {
	ownedGift()
}

// OwnedGiftRegular describes a regular gift owned by a user or a chat.
type OwnedGiftRegular struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             *string         `json:"owned_gift_id"`
	SenderUser              *User           `json:"sender_user"`
	SendDate                int64           `json:"send_date"`
	Text                    *string         `json:"text"`
	Entities                []MessageEntity `json:"entities"`
	IsPrivate               *bool           `json:"is_private"`
	IsSaved                 *bool           `json:"is_saved"`
	CanBeUpgraded           *bool           `json:"can_be_upgraded"`
	WasRefunded             *bool           `json:"was_refunded"`
	ConvertStarCount        *int64          `json:"convert_star_count"`
	PrepaidUpgradeStarCount *int64          `json:"prepaid_upgrade_star_count"`
	IsUpgradeSeparate       *bool           `json:"is_upgrade_separate"`
	UniqueGiftNumber        *int            `json:"unique_gift_number"`
}

func (*OwnedGiftRegular) ownedGift() {}

func (v OwnedGiftRegular) MarshalJSON() ([]byte, error) {
	type alias OwnedGiftRegular
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "regular"})
}

// OwnedGiftUnique describes a unique gift received and owned by a user or a chat.
type OwnedGiftUnique struct {
	Gift              UniqueGift `json:"gift"`
	OwnedGiftID       *string    `json:"owned_gift_id"`
	SenderUser        *User      `json:"sender_user"`
	SendDate          int64      `json:"send_date"`
	IsSaved           *bool      `json:"is_saved"`
	CanBeTransferred  *bool      `json:"can_be_transferred"`
	TransferStarCount *int64     `json:"transfer_star_count"`
	NextTransferDate  *int64     `json:"next_transfer_date"`
}

func (*OwnedGiftUnique) ownedGift() {}

func (v OwnedGiftUnique) MarshalJSON() ([]byte, error) {
	type alias OwnedGiftUnique
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "unique"})
}

// OwnedGifts contains the list of gifts received and owned by a user or a chat.
type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset *string     `json:"next_offset"`
}

func (o *OwnedGifts) UnmarshalJSON(data []byte) error {
	var raw struct {
		TotalCount int               `json:"total_count"`
		Gifts      []json.RawMessage `json:"gifts"`
		NextOffset *string           `json:"next_offset"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	gifts := make([]OwnedGift, len(raw.Gifts))
	for i := range raw.Gifts {
		gift, err := decodeOwnedGift(raw.Gifts[i])
		if err != nil {
			return err
		}
		gifts[i] = gift
	}
	o.TotalCount, o.Gifts, o.NextOffset = raw.TotalCount, gifts, raw.NextOffset
	return nil
}

func decodeOwnedGift(data []byte) (OwnedGift, error) {
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	var gift OwnedGift
	switch probe.Type {
	case "regular":
		gift = &OwnedGiftRegular{}
	case "unique":
		gift = &OwnedGiftUnique{}
	default:
		return nil, fmt.Errorf("telegram: unknown OwnedGift type %q", probe.Type)
	}
	if err := json.Unmarshal(data, gift); err != nil {
		return nil, err
	}
	return gift, nil
}

// BotAccessSettings describes the access settings of a bot.
type BotAccessSettings struct {
	IsAccessRestricted bool   `json:"is_access_restricted"`
	AddedUsers         []User `json:"added_users"`
}

// AcceptedGiftTypes describes the types of gifts that can be gifted to a user or a chat.
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	GiftsFromChannels   bool `json:"gifts_from_channels"`
}

// StarAmount describes an amount of Telegram Stars.
type StarAmount struct {
	Amount         int64  `json:"amount"`
	NanostarAmount *int64 `json:"nanostar_amount"`
}

// ChatBoostSource describes the source of a chat boost. It can be one of
// ChatBoostSourcePremium, ChatBoostSourceGiftCode or ChatBoostSourceGiveaway.
type ChatBoostSource interface {
	chatBoostSource()
}

// ChatBoostSourcePremium describes a boost obtained by subscribing to Telegram
// Premium or by gifting a Telegram Premium subscription to another user.
type ChatBoostSourcePremium struct {
	User User `json:"user"`
}

func (*ChatBoostSourcePremium) chatBoostSource() {}

func (v ChatBoostSourcePremium) MarshalJSON() ([]byte, error) {
	type alias ChatBoostSourcePremium
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "premium"})
}

// ChatBoostSourceGiftCode describes a boost obtained by the creation of
// Telegram Premium gift codes to boost a chat.
type ChatBoostSourceGiftCode struct {
	User User `json:"user"`
}

func (*ChatBoostSourceGiftCode) chatBoostSource() {}

func (v ChatBoostSourceGiftCode) MarshalJSON() ([]byte, error) {
	type alias ChatBoostSourceGiftCode
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "gift_code"})
}

// ChatBoostSourceGiveaway describes a boost obtained by the creation of a
// Telegram Premium or a Telegram Star giveaway.
type ChatBoostSourceGiveaway struct {
	GiveawayMessageID int64  `json:"giveaway_message_id"`
	User              *User  `json:"user"`
	PrizeStarCount    *int64 `json:"prize_star_count"`
	IsUnclaimed       *bool  `json:"is_unclaimed"`
}

func (*ChatBoostSourceGiveaway) chatBoostSource() {}

func (v ChatBoostSourceGiveaway) MarshalJSON() ([]byte, error) {
	type alias ChatBoostSourceGiveaway
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "giveaway"})
}

// ChatBoost contains information about a chat boost.
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int64           `json:"add_date"`
	ExpirationDate int64           `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

// ChatBoostUpdated represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved represents a boost removed from a chat.
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int64           `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

// UserChatBoosts represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

// BusinessBotRights represents the rights of a business bot.
type BusinessBotRights struct {
	CanReply                   *bool `json:"can_reply"`
	CanReadMessages            *bool `json:"can_read_messages"`
	CanDeleteSentMessages      *bool `json:"can_delete_sent_messages"`
	CanDeleteAllMessages       *bool `json:"can_delete_all_messages"`
	CanEditName                *bool `json:"can_edit_name"`
	CanEditBio                 *bool `json:"can_edit_bio"`
	CanEditProfilePhoto        *bool `json:"can_edit_profile_photo"`
	CanEditUsername            *bool `json:"can_edit_username"`
	CanChangeGiftSettings      *bool `json:"can_change_gift_settings"`
	CanViewGiftsAndStars       *bool `json:"can_view_gifts_and_stars"`
	CanConvertGiftsToStars     *bool `json:"can_convert_gifts_to_stars"`
	CanTransferAndUpgradeGifts *bool `json:"can_transfer_and_upgrade_gifts"`
	CanTransferStars           *bool `json:"can_transfer_stars"`
	CanManageStories           *bool `json:"can_manage_stories"`
}

// BusinessConnection describes the connection of the bot with a business account.
type BusinessConnection struct {
	ID         string             `json:"id"`
	User       User               `json:"user"`
	UserChatID int64              `json:"user_chat_id"`
	Date       int64              `json:"date"`
	Rights     *BusinessBotRights `json:"rights"`
	IsEnabled  bool               `json:"is_enabled"`
}

// BusinessMessagesDeleted is received when messages are deleted from a
// connected business account.
type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	Chat                 Chat    `json:"chat"`
	MessageIDs           []int64 `json:"message_ids"`
}

// SentWebAppMessage describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageID *string `json:"inline_message_id"`
}

// SentGuestMessage describes an inline message sent by a guest bot.
type SentGuestMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// BotCommand represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	IsEphemeral *bool  `json:"is_ephemeral"`
}

// BotCommandScope represents the scope to which bot commands are applied. It
// is one of BotCommandScopeDefault, BotCommandScopeAllPrivateChats,
// BotCommandScopeAllGroupChats, BotCommandScopeAllChatAdministrators,
// BotCommandScopeChat, BotCommandScopeChatAdministrators or
// BotCommandScopeChatMember.
type BotCommandScope interface {
	botCommandScope()
}

// BotCommandScopeDefault represents the default scope of bot commands. Default
// commands are used if no commands with a narrower scope are specified for the
// user.
type BotCommandScopeDefault struct{}

func (*BotCommandScopeDefault) botCommandScope() {}

func (v BotCommandScopeDefault) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeDefault
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "default"})
}

// BotCommandScopeAllPrivateChats represents the scope of bot commands,
// covering all private chats.
type BotCommandScopeAllPrivateChats struct{}

func (*BotCommandScopeAllPrivateChats) botCommandScope() {}

func (v BotCommandScopeAllPrivateChats) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeAllPrivateChats
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "all_private_chats"})
}

// BotCommandScopeAllGroupChats represents the scope of bot commands, covering
// all group and supergroup chats.
type BotCommandScopeAllGroupChats struct{}

func (*BotCommandScopeAllGroupChats) botCommandScope() {}

func (v BotCommandScopeAllGroupChats) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeAllGroupChats
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "all_group_chats"})
}

// BotCommandScopeAllChatAdministrators represents the scope of bot commands,
// covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct{}

func (*BotCommandScopeAllChatAdministrators) botCommandScope() {}

func (v BotCommandScopeAllChatAdministrators) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeAllChatAdministrators
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "all_chat_administrators"})
}

// BotCommandScopeChat represents the scope of bot commands, covering a
// specific chat.
type BotCommandScopeChat struct {
	ChatID ChatID `json:"chat_id"`
}

func (*BotCommandScopeChat) botCommandScope() {}

func (v BotCommandScopeChat) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeChat
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "chat"})
}

// BotCommandScopeChatAdministrators represents the scope of bot commands,
// covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	ChatID ChatID `json:"chat_id"`
}

func (*BotCommandScopeChatAdministrators) botCommandScope() {}

func (v BotCommandScopeChatAdministrators) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeChatAdministrators
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "chat_administrators"})
}

// BotCommandScopeChatMember represents the scope of bot commands, covering a
// specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	ChatID ChatID `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

func (*BotCommandScopeChatMember) botCommandScope() {}

func (v BotCommandScopeChatMember) MarshalJSON() ([]byte, error) {
	type alias BotCommandScopeChatMember
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "chat_member"})
}

// BotName represents the bot's name.
type BotName struct {
	Name string `json:"name"`
}

// BotDescription represents the bot's description.
type BotDescription struct {
	Description string `json:"description"`
}

// BotShortDescription represents the bot's short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// MenuButton describes the bot's menu button in a private chat. It is one of
// MenuButtonCommands, MenuButtonWebApp or MenuButtonDefault.
type MenuButton interface {
	menuButton()
}

// MenuButtonCommands represents a menu button, which opens the bot's list of
// commands.
type MenuButtonCommands struct{}

func (*MenuButtonCommands) menuButton() {}

func (v MenuButtonCommands) MarshalJSON() ([]byte, error) {
	type alias MenuButtonCommands
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "commands"})
}

// MenuButtonWebApp represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}

func (*MenuButtonWebApp) menuButton() {}

func (v MenuButtonWebApp) MarshalJSON() ([]byte, error) {
	type alias MenuButtonWebApp
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "web_app"})
}

// MenuButtonDefault describes that no specific value for the menu button was
// set.
type MenuButtonDefault struct{}

func (*MenuButtonDefault) menuButton() {}

func (v MenuButtonDefault) MarshalJSON() ([]byte, error) {
	type alias MenuButtonDefault
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "default"})
}

// InlineQuery represents an incoming inline query. When the user sends an empty
// query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType *string   `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

// InlineQueryResultsButton represents a button to be shown above inline query
// results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter *string     `json:"start_parameter,omitempty"`
}

// InlineQueryResult represents one result of an inline query.
type InlineQueryResult interface {
	inlineQueryResult()
}

// InlineQueryResultArticle represents a link to an article or web page.
type InlineQueryResultArticle struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultArticle) inlineQueryResult() {}

func (v InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultArticle
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "article"})
}

// InlineQueryResultPhoto represents a link to a photo. By default, this photo
// will be sent by the user with optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of
// the photo.
type InlineQueryResultPhoto struct {
	ID                    string                `json:"id"`
	PhotoURL              string                `json:"photo_url"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	PhotoWidth            *int                  `json:"photo_width,omitempty"`
	PhotoHeight           *int                  `json:"photo_height,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultPhoto) inlineQueryResult() {}

func (v InlineQueryResultPhoto) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// InlineQueryResultGif represents a link to an animated GIF file. By default,
// this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the
// specified content instead of the animation.
type InlineQueryResultGif struct {
	ID                    string                `json:"id"`
	GifURL                string                `json:"gif_url"`
	GifWidth              *int                  `json:"gif_width,omitempty"`
	GifHeight             *int                  `json:"gif_height,omitempty"`
	GifDuration           *int                  `json:"gif_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     *string               `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultGif) inlineQueryResult() {}

func (v InlineQueryResultGif) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultGif
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "gif"})
}

// InlineQueryResultMpeg4Gif represents a link to a video animation
// (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4
// file will be sent by the user with optional caption. Alternatively, you can
// use input_message_content to send a message with the specified content
// instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	ID                    string                `json:"id"`
	Mpeg4URL              string                `json:"mpeg4_url"`
	Mpeg4Width            *int                  `json:"mpeg4_width,omitempty"`
	Mpeg4Height           *int                  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         *int                  `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     *string               `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultMpeg4Gif) inlineQueryResult() {}

func (v InlineQueryResultMpeg4Gif) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultMpeg4Gif
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "mpeg4_gif"})
}

// InlineQueryResultVideo represents a link to a page containing an embedded
// video player or a video file. By default, this video file will be sent by the
// user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of
// the video.
type InlineQueryResultVideo struct {
	ID                    string                `json:"id"`
	VideoURL              string                `json:"video_url"`
	MimeType              string                `json:"mime_type"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	VideoWidth            *int                  `json:"video_width,omitempty"`
	VideoHeight           *int                  `json:"video_height,omitempty"`
	VideoDuration         *int                  `json:"video_duration,omitempty"`
	Description           *string               `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultVideo) inlineQueryResult() {}

func (v InlineQueryResultVideo) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// InlineQueryResultAudio represents a link to an MP3 audio file. By default,
// this audio file will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of
// the audio.
type InlineQueryResultAudio struct {
	ID                  string                `json:"id"`
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int                  `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultAudio) inlineQueryResult() {}

func (v InlineQueryResultAudio) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultAudio
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "audio"})
}

// InlineQueryResultVoice represents a link to a voice recording in an .OGG
// container encoded with OPUS. By default, this voice recording will be sent by
// the user. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the voice message.
type InlineQueryResultVoice struct {
	ID                  string                `json:"id"`
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	VoiceDuration       *int                  `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultVoice) inlineQueryResult() {}

func (v InlineQueryResultVoice) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultVoice
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "voice"})
}

// InlineQueryResultDocument represents a link to a file. By default, this file
// will be sent by the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with the specified content instead of
// the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultDocument) inlineQueryResult() {}

func (v InlineQueryResultDocument) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultDocument
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "document"})
}

// InlineQueryResultLocation represents a location on a map. By default, the
// location will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of
// the location.
type InlineQueryResultLocation struct {
	ID                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   *float64              `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int                  `json:"live_period,omitempty"`
	Heading              *int                  `json:"heading,omitempty"`
	ProximityAlertRadius *int                  `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL         *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      *int                  `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultLocation) inlineQueryResult() {}

func (v InlineQueryResultLocation) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultLocation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "location"})
}

// InlineQueryResultVenue represents a venue. By default, the venue will be sent
// by the user. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	ID                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        *string               `json:"foursquare_id,omitempty"`
	FoursquareType      *string               `json:"foursquare_type,omitempty"`
	GooglePlaceID       *string               `json:"google_place_id,omitempty"`
	GooglePlaceType     *string               `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultVenue) inlineQueryResult() {}

func (v InlineQueryResultVenue) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultVenue
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "venue"})
}

// InlineQueryResultContact represents a contact with a phone number. By
// default, this contact will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of
// the contact.
type InlineQueryResultContact struct {
	ID                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	Vcard               *string               `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

func (*InlineQueryResultContact) inlineQueryResult() {}

func (v InlineQueryResultContact) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultContact
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "contact"})
}

// InlineQueryResultGame represents a Game.
type InlineQueryResultGame struct {
	ID            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (*InlineQueryResultGame) inlineQueryResult() {}

func (v InlineQueryResultGame) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultGame
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "game"})
}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the
// Telegram servers. By default, this photo will be sent by the user with an
// optional caption. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	ID                    string                `json:"id"`
	PhotoFileID           string                `json:"photo_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedPhoto) inlineQueryResult() {}

func (v InlineQueryResultCachedPhoto) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// InlineQueryResultCachedGif represents a link to an animated GIF file stored
// on the Telegram servers. By default, this animated GIF file will be sent by
// the user with an optional caption. Alternatively, you can use
// input_message_content to send a message with specified content instead of the
// animation.
type InlineQueryResultCachedGif struct {
	ID                    string                `json:"id"`
	GifFileID             string                `json:"gif_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedGif) inlineQueryResult() {}

func (v InlineQueryResultCachedGif) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedGif
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "gif"})
}

// InlineQueryResultCachedMpeg4Gif represents a link to a video animation
// (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By
// default, this animated MPEG-4 file will be sent by the user with an optional
// caption. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	ID                    string                `json:"id"`
	Mpeg4FileID           string                `json:"mpeg4_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedMpeg4Gif) inlineQueryResult() {}

func (v InlineQueryResultCachedMpeg4Gif) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedMpeg4Gif
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "mpeg4_gif"})
}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the
// Telegram servers. By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the
// specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	ID                  string                `json:"id"`
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedSticker) inlineQueryResult() {}

func (v InlineQueryResultCachedSticker) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedSticker
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "sticker"})
}

// InlineQueryResultCachedDocument represents a link to a file stored on the
// Telegram servers. By default, this file will be sent by the user with an
// optional caption. Alternatively, you can use input_message_content to send a
// message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	ID                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedDocument) inlineQueryResult() {}

func (v InlineQueryResultCachedDocument) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedDocument
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "document"})
}

// InlineQueryResultCachedVideo represents a link to a video file stored on the
// Telegram servers. By default, this video file will be sent by the user with
// an optional caption. Alternatively, you can use input_message_content to send
// a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	ID                    string                `json:"id"`
	VideoFileID           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *string               `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedVideo) inlineQueryResult() {}

func (v InlineQueryResultCachedVideo) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// InlineQueryResultCachedVoice represents a link to a voice message stored on
// the Telegram servers. By default, this voice message will be sent by the
// user. Alternatively, you can use input_message_content to send a message with
// the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	ID                  string                `json:"id"`
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedVoice) inlineQueryResult() {}

func (v InlineQueryResultCachedVoice) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedVoice
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "voice"})
}

// InlineQueryResultCachedAudio represents a link to an MP3 audio file stored on
// the Telegram servers. By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the
// specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	ID                  string                `json:"id"`
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *string               `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

func (*InlineQueryResultCachedAudio) inlineQueryResult() {}

func (v InlineQueryResultCachedAudio) MarshalJSON() ([]byte, error) {
	type alias InlineQueryResultCachedAudio
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "audio"})
}

// InputMessageContent represents the content of a message to be sent as a
// result of an inline query.
type InputMessageContent interface {
	inputMessageContent()
}

// InputTextMessageContent represents the content of a text message to be sent
// as the result of an inline query.
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          *string             `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (*InputTextMessageContent) inputMessageContent() {}

// InputRichMessageContent represents the content of a rich message to be sent
// as the result of an inline query.
type InputRichMessageContent struct {
	RichMessage InputRichMessage `json:"rich_message"`
}

func (*InputRichMessageContent) inputMessageContent() {}

// InputLocationMessageContent represents the content of a location message to
// be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

func (*InputLocationMessageContent) inputMessageContent() {}

// InputVenueMessageContent represents the content of a venue message to be sent
// as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

func (*InputVenueMessageContent) inputMessageContent() {}

// InputContactMessageContent represents the content of a contact message to be
// sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	Vcard       *string `json:"vcard,omitempty"`
}

func (*InputContactMessageContent) inputMessageContent() {}

// InputInvoiceMessageContent represents the content of an invoice message to be
// sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             *string        `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              *int64         `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoURL                  *string        `json:"photo_url,omitempty"`
	PhotoSize                 *int           `json:"photo_size,omitempty"`
	PhotoWidth                *int           `json:"photo_width,omitempty"`
	PhotoHeight               *int           `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
}

func (*InputInvoiceMessageContent) inputMessageContent() {}

// PreparedInlineMessage describes an inline message to be sent by a user of a
// Mini App.
type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int64  `json:"expiration_date"`
}

// PreparedKeyboardButton describes a keyboard button to be used by a user of a
// Mini App.
type PreparedKeyboardButton struct {
	ID string `json:"id"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by
// the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// InputMedia describes the content of a media message to be sent.
type InputMedia interface {
	inputMedia()
}

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Media                 InputFile       `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	Width                 *int            `json:"width,omitempty"`
	Height                *int            `json:"height,omitempty"`
	Duration              *int            `json:"duration,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

func (*InputMediaAnimation) inputMedia() {}

func (v InputMediaAnimation) MarshalJSON() ([]byte, error) {
	type alias InputMediaAnimation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "animation"})
}

// InputMediaAudio represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Media           InputFile       `json:"media"`
	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int            `json:"duration,omitempty"`
	Performer       *string         `json:"performer,omitempty"`
	Title           *string         `json:"title,omitempty"`
}

func (*InputMediaAudio) inputMedia() {}

func (v InputMediaAudio) MarshalJSON() ([]byte, error) {
	type alias InputMediaAudio
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "audio"})
}

// InputMediaDocument represents a general file to be sent.
type InputMediaDocument struct {
	Media                       InputFile       `json:"media"`
	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
	Caption                     *string         `json:"caption,omitempty"`
	ParseMode                   *string         `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool           `json:"disable_content_type_detection,omitempty"`
}

func (*InputMediaDocument) inputMedia() {}

func (v InputMediaDocument) MarshalJSON() ([]byte, error) {
	type alias InputMediaDocument
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "document"})
}

// InputMediaLink represents an HTTP link to be sent.
type InputMediaLink struct {
	URL string `json:"url"`
}

func (*InputMediaLink) inputMedia() {}

func (v InputMediaLink) MarshalJSON() ([]byte, error) {
	type alias InputMediaLink
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "link"})
}

// InputMediaLivePhoto represents a live photo to be sent.
type InputMediaLivePhoto struct {
	Media                 InputFile       `json:"media"`
	Photo                 InputFile       `json:"photo"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

func (*InputMediaLivePhoto) inputMedia() {}

func (v InputMediaLivePhoto) MarshalJSON() ([]byte, error) {
	type alias InputMediaLivePhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "live_photo"})
}

// InputMediaLocation represents a location to be sent.
type InputMediaLocation struct {
	Latitude           float64  `json:"latitude"`
	Longitude          float64  `json:"longitude"`
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`
}

func (*InputMediaLocation) inputMedia() {}

func (v InputMediaLocation) MarshalJSON() ([]byte, error) {
	type alias InputMediaLocation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "location"})
}

// InputMediaPhoto represents a photo to be sent.
type InputMediaPhoto struct {
	Media                 InputFile       `json:"media"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

func (*InputMediaPhoto) inputMedia() {}

func (v InputMediaPhoto) MarshalJSON() ([]byte, error) {
	type alias InputMediaPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// InputMediaSticker represents a sticker file to be sent.
type InputMediaSticker struct {
	Media InputFile `json:"media"`
	Emoji *string   `json:"emoji,omitempty"`
}

func (*InputMediaSticker) inputMedia() {}

func (v InputMediaSticker) MarshalJSON() ([]byte, error) {
	type alias InputMediaSticker
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "sticker"})
}

// InputMediaVenue represents a venue to be sent.
type InputMediaVenue struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

func (*InputMediaVenue) inputMedia() {}

func (v InputMediaVenue) MarshalJSON() ([]byte, error) {
	type alias InputMediaVenue
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "venue"})
}

// InputMediaVideo represents a video to be sent.
type InputMediaVideo struct {
	Media                 InputFile       `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Cover                 *InputFile      `json:"cover,omitempty"`
	StartTimestamp        *int            `json:"start_timestamp,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *string         `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	Width                 *int            `json:"width,omitempty"`
	Height                *int            `json:"height,omitempty"`
	Duration              *int            `json:"duration,omitempty"`
	SupportsStreaming     *bool           `json:"supports_streaming,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

func (*InputMediaVideo) inputMedia() {}

func (v InputMediaVideo) MarshalJSON() ([]byte, error) {
	type alias InputMediaVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// InputMediaVoiceNote represents a voice message file to be sent.
type InputMediaVoiceNote struct {
	Media           InputFile       `json:"media"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *string         `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int            `json:"duration,omitempty"`
}

func (*InputMediaVoiceNote) inputMedia() {}

func (v InputMediaVoiceNote) MarshalJSON() ([]byte, error) {
	type alias InputMediaVoiceNote
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "voice_note"})
}

// InputPaidMedia describes the paid media to be sent.
type InputPaidMedia interface {
	inputPaidMedia()
}

// InputPaidMediaLivePhoto means the paid media to send is a live photo.
type InputPaidMediaLivePhoto struct {
	Media InputFile `json:"media"`
	Photo InputFile `json:"photo"`
}

func (*InputPaidMediaLivePhoto) inputPaidMedia() {}

func (v InputPaidMediaLivePhoto) MarshalJSON() ([]byte, error) {
	type alias InputPaidMediaLivePhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "live_photo"})
}

// InputPaidMediaPhoto means the paid media to send is a photo.
type InputPaidMediaPhoto struct {
	Media InputFile `json:"media"`
}

func (*InputPaidMediaPhoto) inputPaidMedia() {}

func (v InputPaidMediaPhoto) MarshalJSON() ([]byte, error) {
	type alias InputPaidMediaPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// InputPaidMediaVideo means the paid media to send is a video.
type InputPaidMediaVideo struct {
	Media             InputFile  `json:"media"`
	Thumbnail         *InputFile `json:"thumbnail,omitempty"`
	Cover             *InputFile `json:"cover,omitempty"`
	StartTimestamp    *int       `json:"start_timestamp,omitempty"`
	Width             *int       `json:"width,omitempty"`
	Height            *int       `json:"height,omitempty"`
	Duration          *int       `json:"duration,omitempty"`
	SupportsStreaming *bool      `json:"supports_streaming,omitempty"`
}

func (*InputPaidMediaVideo) inputPaidMedia() {}

func (v InputPaidMediaVideo) MarshalJSON() ([]byte, error) {
	type alias InputPaidMediaVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// InputProfilePhoto describes a profile photo to set.
type InputProfilePhoto interface {
	inputProfilePhoto()
}

// InputProfilePhotoStatic is a static profile photo in the .JPG format.
type InputProfilePhotoStatic struct {
	Photo InputFile `json:"photo"`
}

func (*InputProfilePhotoStatic) inputProfilePhoto() {}

func (v InputProfilePhotoStatic) MarshalJSON() ([]byte, error) {
	type alias InputProfilePhotoStatic
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "static"})
}

// InputProfilePhotoAnimated is an animated profile photo in the MPEG4 format.
type InputProfilePhotoAnimated struct {
	Animation          InputFile `json:"animation"`
	MainFrameTimestamp *float64  `json:"main_frame_timestamp,omitempty"`
}

func (*InputProfilePhotoAnimated) inputProfilePhoto() {}

func (v InputProfilePhotoAnimated) MarshalJSON() ([]byte, error) {
	type alias InputProfilePhotoAnimated
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "animated"})
}

// InputStoryContent describes the content of a story to post.
type InputStoryContent interface {
	inputStoryContent()
}

// InputStoryContentPhoto describes a photo to post as a story.
type InputStoryContentPhoto struct {
	Photo InputFile `json:"photo"`
}

func (*InputStoryContentPhoto) inputStoryContent() {}

func (v InputStoryContentPhoto) MarshalJSON() ([]byte, error) {
	type alias InputStoryContentPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// InputStoryContentVideo describes a video to post as a story.
type InputStoryContentVideo struct {
	Video               InputFile `json:"video"`
	Duration            *float64  `json:"duration,omitempty"`
	CoverFrameTimestamp *float64  `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         *bool     `json:"is_animation,omitempty"`
}

func (*InputStoryContentVideo) inputStoryContent() {}

func (v InputStoryContentVideo) MarshalJSON() ([]byte, error) {
	type alias InputStoryContentVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// Sticker represents a sticker.
type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail"`
	Emoji            *string       `json:"emoji"`
	SetName          *string       `json:"set_name"`
	PremiumAnimation *File         `json:"premium_animation"`
	MaskPosition     *MaskPosition `json:"mask_position"`
	CustomEmojiID    *string       `json:"custom_emoji_id"`
	NeedsRepainting  *bool         `json:"needs_repainting"`
	FileSize         *int64        `json:"file_size"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers"`
	Thumbnail   *PhotoSize `json:"thumbnail"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// InputSticker describes a sticker to be added to a sticker set.
type InputSticker struct {
	Sticker      InputFile     `json:"sticker"`
	Format       string        `json:"format"`
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     *[]string     `json:"keywords,omitempty"`
}

// StoryAreaPosition describes the position of a clickable area within a story.
type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

// StoryAreaType describes the type of a clickable area on a story.
type StoryAreaType interface {
	storyAreaType()
}

// StoryAreaTypeLocation describes a story area pointing to a location.
type StoryAreaTypeLocation struct {
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Address   *LocationAddress `json:"address,omitempty"`
}

func (*StoryAreaTypeLocation) storyAreaType() {}

func (v StoryAreaTypeLocation) MarshalJSON() ([]byte, error) {
	type alias StoryAreaTypeLocation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "location"})
}

// StoryAreaTypeSuggestedReaction describes a story area pointing to a suggested reaction.
type StoryAreaTypeSuggestedReaction struct {
	ReactionType ReactionType `json:"reaction_type"`
	IsDark       *bool        `json:"is_dark,omitempty"`
	IsFlipped    *bool        `json:"is_flipped,omitempty"`
}

func (*StoryAreaTypeSuggestedReaction) storyAreaType() {}

func (v StoryAreaTypeSuggestedReaction) MarshalJSON() ([]byte, error) {
	type alias StoryAreaTypeSuggestedReaction
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "suggested_reaction"})
}

// StoryAreaTypeLink describes a story area pointing to an HTTP or tg:// link.
type StoryAreaTypeLink struct {
	URL string `json:"url"`
}

func (*StoryAreaTypeLink) storyAreaType() {}

func (v StoryAreaTypeLink) MarshalJSON() ([]byte, error) {
	type alias StoryAreaTypeLink
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "link"})
}

// StoryAreaTypeWeather describes a story area containing weather information.
type StoryAreaTypeWeather struct {
	Temperature     float64 `json:"temperature"`
	Emoji           string  `json:"emoji"`
	BackgroundColor int     `json:"background_color"`
}

func (*StoryAreaTypeWeather) storyAreaType() {}

func (v StoryAreaTypeWeather) MarshalJSON() ([]byte, error) {
	type alias StoryAreaTypeWeather
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "weather"})
}

// StoryAreaTypeUniqueGift describes a story area pointing to a unique gift.
type StoryAreaTypeUniqueGift struct {
	Name string `json:"name"`
}

func (*StoryAreaTypeUniqueGift) storyAreaType() {}

func (v StoryAreaTypeUniqueGift) MarshalJSON() ([]byte, error) {
	type alias StoryAreaTypeUniqueGift
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "unique_gift"})
}

// UnmarshalJSON decodes a StoryArea, probing the area type's "type"
// discriminator to select the concrete StoryAreaType variant. (Go forbids
// methods on pointer-to-interface receivers, so the union is decoded at the
// container level; StoryArea is the only place StoryAreaType occurs.)
func (a *StoryArea) UnmarshalJSON(b []byte) error {
	var raw struct {
		Position StoryAreaPosition `json:"position"`
		Type     json.RawMessage   `json:"type"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	t, err := unmarshalStoryAreaType(raw.Type)
	if err != nil {
		return err
	}
	a.Position = raw.Position
	a.Type = t
	return nil
}

func unmarshalStoryAreaType(b []byte) (StoryAreaType, error) {
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case "location":
		var v StoryAreaTypeLocation
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "suggested_reaction":
		var v StoryAreaTypeSuggestedReaction
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "link":
		var v StoryAreaTypeLink
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "weather":
		var v StoryAreaTypeWeather
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "unique_gift":
		var v StoryAreaTypeUniqueGift
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("telegram: unknown StoryAreaType type %q", probe.Type)
	}
}

// StoryArea describes a clickable area on a story media.
type StoryArea struct {
	Position StoryAreaPosition `json:"position"`
	Type     StoryAreaType     `json:"type"`
}

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int64  `json:"amount"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	Name            *string          `json:"name"`
	PhoneNumber     *string          `json:"phone_number"`
	Email           *string          `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int64      `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate *int64     `json:"subscription_expiration_date"`
	IsRecurring                *bool      `json:"is_recurring"`
	IsFirstRecurring           *bool      `json:"is_first_recurring"`
	ShippingOptionID           *string    `json:"shipping_option_id"`
	OrderInfo                  *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id"`
}

// RefundedPayment contains basic information about a refunded payment.
type RefundedPayment struct {
	Currency                string  `json:"currency"`
	TotalAmount             int64   `json:"total_amount"`
	InvoicePayload          string  `json:"invoice_payload"`
	TelegramPaymentChargeID string  `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID *string `json:"provider_payment_charge_id"`
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID *string    `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

// PaidMediaPurchased contains information about a paid media purchase.
type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

// RevenueWithdrawalState describes the state of a revenue withdrawal operation.
// It can be one of RevenueWithdrawalStatePending, RevenueWithdrawalStateSucceeded
// or RevenueWithdrawalStateFailed.
type RevenueWithdrawalState interface {
	revenueWithdrawalState()
}

// RevenueWithdrawalStatePending means the withdrawal is in progress.
type RevenueWithdrawalStatePending struct{}

func (*RevenueWithdrawalStatePending) revenueWithdrawalState() {}

func (v RevenueWithdrawalStatePending) MarshalJSON() ([]byte, error) {
	type alias RevenueWithdrawalStatePending
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "pending"})
}

// RevenueWithdrawalStateSucceeded means the withdrawal succeeded.
type RevenueWithdrawalStateSucceeded struct {
	Date int64  `json:"date"`
	URL  string `json:"url"`
}

func (*RevenueWithdrawalStateSucceeded) revenueWithdrawalState() {}

func (v RevenueWithdrawalStateSucceeded) MarshalJSON() ([]byte, error) {
	type alias RevenueWithdrawalStateSucceeded
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "succeeded"})
}

// RevenueWithdrawalStateFailed means the withdrawal failed and the transaction was refunded.
type RevenueWithdrawalStateFailed struct{}

func (*RevenueWithdrawalStateFailed) revenueWithdrawalState() {}

func (v RevenueWithdrawalStateFailed) MarshalJSON() ([]byte, error) {
	type alias RevenueWithdrawalStateFailed
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "failed"})
}

// AffiliateInfo contains information about the affiliate that received a
// commission via this transaction.
type AffiliateInfo struct {
	AffiliateUser      *User  `json:"affiliate_user"`
	AffiliateChat      *Chat  `json:"affiliate_chat"`
	CommissionPerMille int    `json:"commission_per_mille"`
	Amount             int64  `json:"amount"`
	NanostarAmount     *int64 `json:"nanostar_amount"`
}

// TransactionPartner describes the source of a transaction, or its recipient
// for outgoing transactions. It can be one of TransactionPartnerUser,
// TransactionPartnerChat, TransactionPartnerAffiliateProgram,
// TransactionPartnerFragment, TransactionPartnerTelegramAds,
// TransactionPartnerTelegramApi or TransactionPartnerOther.
type TransactionPartner interface {
	transactionPartner()
}

// TransactionPartnerUser describes a transaction with a user.
type TransactionPartnerUser struct {
	TransactionType             string         `json:"transaction_type"`
	User                        User           `json:"user"`
	Affiliate                   *AffiliateInfo `json:"affiliate"`
	InvoicePayload              *string        `json:"invoice_payload"`
	SubscriptionPeriod          *int           `json:"subscription_period"`
	PaidMedia                   *[]PaidMedia   `json:"paid_media"`
	PaidMediaPayload            *string        `json:"paid_media_payload"`
	Gift                        *Gift          `json:"gift"`
	PremiumSubscriptionDuration *int           `json:"premium_subscription_duration"`
}

func (p *TransactionPartnerUser) UnmarshalJSON(data []byte) error {
	type alias TransactionPartnerUser
	var raw struct {
		PaidMedia []json.RawMessage `json:"paid_media"`
		*alias
	}
	raw.alias = (*alias)(p)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if raw.PaidMedia == nil {
		return nil
	}
	media := make([]PaidMedia, len(raw.PaidMedia))
	for i := range raw.PaidMedia {
		v, err := decodePaidMedia(raw.PaidMedia[i])
		if err != nil {
			return err
		}
		media[i] = v
	}
	p.PaidMedia = &media
	return nil
}

func (*TransactionPartnerUser) transactionPartner() {}

func (v TransactionPartnerUser) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerUser
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "user"})
}

// TransactionPartnerChat describes a transaction with a chat.
type TransactionPartnerChat struct {
	Chat Chat  `json:"chat"`
	Gift *Gift `json:"gift"`
}

func (*TransactionPartnerChat) transactionPartner() {}

func (v TransactionPartnerChat) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerChat
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "chat"})
}

// TransactionPartnerAffiliateProgram describes the affiliate program that
// issued the affiliate commission received via this transaction.
type TransactionPartnerAffiliateProgram struct {
	SponsorUser        *User `json:"sponsor_user"`
	CommissionPerMille int   `json:"commission_per_mille"`
}

func (*TransactionPartnerAffiliateProgram) transactionPartner() {}

func (v TransactionPartnerAffiliateProgram) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerAffiliateProgram
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "affiliate_program"})
}

// TransactionPartnerFragment describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state"`
}

func (p *TransactionPartnerFragment) UnmarshalJSON(data []byte) error {
	var raw struct {
		WithdrawalState json.RawMessage `json:"withdrawal_state"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	state, err := decodeRevenueWithdrawalState(raw.WithdrawalState)
	if err != nil {
		return err
	}
	p.WithdrawalState = state
	return nil
}

func (*TransactionPartnerFragment) transactionPartner() {}

func (v TransactionPartnerFragment) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerFragment
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "fragment"})
}

// TransactionPartnerTelegramAds describes a withdrawal transaction to the Telegram Ads platform.
type TransactionPartnerTelegramAds struct{}

func (*TransactionPartnerTelegramAds) transactionPartner() {}

func (v TransactionPartnerTelegramAds) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerTelegramAds
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "telegram_ads"})
}

// TransactionPartnerTelegramApi describes a transaction with payment for paid broadcasting.
type TransactionPartnerTelegramApi struct {
	RequestCount int `json:"request_count"`
}

func (*TransactionPartnerTelegramApi) transactionPartner() {}

func (v TransactionPartnerTelegramApi) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerTelegramApi
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "telegram_api"})
}

// TransactionPartnerOther describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct{}

func (*TransactionPartnerOther) transactionPartner() {}

func (v TransactionPartnerOther) MarshalJSON() ([]byte, error) {
	type alias TransactionPartnerOther
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "other"})
}

// StarTransaction describes a Telegram Star transaction.
type StarTransaction struct {
	ID             string             `json:"id"`
	Amount         int64              `json:"amount"`
	NanostarAmount *int64             `json:"nanostar_amount"`
	Date           int64              `json:"date"`
	Source         TransactionPartner `json:"source"`
	Receiver       TransactionPartner `json:"receiver"`
}

func (s *StarTransaction) UnmarshalJSON(data []byte) error {
	type alias StarTransaction
	var raw struct {
		Source   json.RawMessage `json:"source"`
		Receiver json.RawMessage `json:"receiver"`
		*alias
	}
	raw.alias = (*alias)(s)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	source, err := decodeTransactionPartner(raw.Source)
	if err != nil {
		return err
	}
	receiver, err := decodeTransactionPartner(raw.Receiver)
	if err != nil {
		return err
	}
	s.Source, s.Receiver = source, receiver
	return nil
}

func decodeRevenueWithdrawalState(data []byte) (RevenueWithdrawalState, error) {
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	var state RevenueWithdrawalState
	switch probe.Type {
	case "pending":
		state = &RevenueWithdrawalStatePending{}
	case "succeeded":
		state = &RevenueWithdrawalStateSucceeded{}
	case "failed":
		state = &RevenueWithdrawalStateFailed{}
	default:
		return nil, fmt.Errorf("telegram: unknown RevenueWithdrawalState type %q", probe.Type)
	}
	if err := json.Unmarshal(data, state); err != nil {
		return nil, err
	}
	return state, nil
}

func decodeTransactionPartner(data []byte) (TransactionPartner, error) {
	if len(data) == 0 || bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return nil, nil
	}
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	var partner TransactionPartner
	switch probe.Type {
	case "user":
		partner = &TransactionPartnerUser{}
	case "chat":
		partner = &TransactionPartnerChat{}
	case "affiliate_program":
		partner = &TransactionPartnerAffiliateProgram{}
	case "fragment":
		partner = &TransactionPartnerFragment{}
	case "telegram_ads":
		partner = &TransactionPartnerTelegramAds{}
	case "telegram_api":
		partner = &TransactionPartnerTelegramApi{}
	case "other":
		partner = &TransactionPartnerOther{}
	default:
		return nil, fmt.Errorf("telegram: unknown TransactionPartner type %q", probe.Type)
	}
	if err := json.Unmarshal(data, partner); err != nil {
		return nil, err
	}
	return partner, nil
}

// StarTransactions contains a list of Telegram Star transactions.
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// PassportData describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

// PassportFile represents a file uploaded to Telegram Passport. Currently all
// Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

// EncryptedPassportElement describes documents or other Telegram Passport
// elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        *string         `json:"data"`
	PhoneNumber *string         `json:"phone_number"`
	Email       *string         `json:"email"`
	Files       *[]PassportFile `json:"files"`
	FrontSide   *PassportFile   `json:"front_side"`
	ReverseSide *PassportFile   `json:"reverse_side"`
	Selfie      *PassportFile   `json:"selfie"`
	Translation *[]PassportFile `json:"translation"`
	Hash        string          `json:"hash"`
}

// EncryptedCredentials describes data required for decrypting and
// authenticating EncryptedPassportElement.
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// PassportElementError represents an error in a Telegram Passport element
// which was submitted that should be resolved by the user. It can be one of
// PassportElementErrorDataField, PassportElementErrorFrontSide,
// PassportElementErrorReverseSide, PassportElementErrorSelfie,
// PassportElementErrorFile, PassportElementErrorFiles,
// PassportElementErrorTranslationFile, PassportElementErrorTranslationFiles
// or PassportElementErrorUnspecified.
type PassportElementError interface {
	passportElementError()
}

// PassportElementErrorDataField represents an issue in one of the data fields
// that was provided by the user. The error is considered resolved when the
// field's value changes.
type PassportElementErrorDataField struct {
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

func (*PassportElementErrorDataField) passportElementError() {}

func (v PassportElementErrorDataField) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorDataField
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "data"})
}

// PassportElementErrorFrontSide represents an issue with the front side of a
// document. The error is considered resolved when the file with the front side
// of the document changes.
type PassportElementErrorFrontSide struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorFrontSide) passportElementError() {}

func (v PassportElementErrorFrontSide) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorFrontSide
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "front_side"})
}

// PassportElementErrorReverseSide represents an issue with the reverse side of
// a document. The error is considered resolved when the file with the reverse
// side of the document changes.
type PassportElementErrorReverseSide struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorReverseSide) passportElementError() {}

func (v PassportElementErrorReverseSide) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorReverseSide
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "reverse_side"})
}

// PassportElementErrorSelfie represents an issue with the selfie with a
// document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorSelfie) passportElementError() {}

func (v PassportElementErrorSelfie) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorSelfie
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "selfie"})
}

// PassportElementErrorFile represents an issue with a document scan. The error
// is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorFile) passportElementError() {}

func (v PassportElementErrorFile) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorFile
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "file"})
}

// PassportElementErrorFiles represents an issue with a list of scans. The
// error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (*PassportElementErrorFiles) passportElementError() {}

func (v PassportElementErrorFiles) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorFiles
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "files"})
}

// PassportElementErrorTranslationFile represents an issue with one of the
// files that constitute the translation of a document. The error is
// considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

func (*PassportElementErrorTranslationFile) passportElementError() {}

func (v PassportElementErrorTranslationFile) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorTranslationFile
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "translation_file"})
}

// PassportElementErrorTranslationFiles represents an issue with the translated
// version of a document. The error is considered resolved when a file with the
// document translation changes.
type PassportElementErrorTranslationFiles struct {
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

func (*PassportElementErrorTranslationFiles) passportElementError() {}

func (v PassportElementErrorTranslationFiles) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorTranslationFiles
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "translation_files"})
}

// PassportElementErrorUnspecified represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

func (*PassportElementErrorUnspecified) passportElementError() {}

func (v PassportElementErrorUnspecified) MarshalJSON() ([]byte, error) {
	type alias PassportElementErrorUnspecified
	return json.Marshal(struct {
		alias
		Source string `json:"source"`
	}{alias(v), "unspecified"})
}

// Game represents a game. Use BotFather to create and edit games; their short
// names act as unique identifiers.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

// CallbackGame is a placeholder, currently holding no information. Use
// BotFather to set up your game.
type CallbackGame struct{}

// GameHighScore represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}
