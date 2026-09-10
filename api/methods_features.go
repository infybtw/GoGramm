package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// GetBusinessConnectionParams holds parameters for getBusinessConnection.
type GetBusinessConnectionParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// GetBusinessConnection gets information about the connection of the bot with a business account.
func (c *Client) GetBusinessConnection(ctx context.Context, p *GetBusinessConnectionParams) (*BusinessConnection, error) {
	var b BusinessConnection
	if err := c.Invoke(ctx, "getBusinessConnection", p, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

// GetManagedBotTokenParams holds parameters for getManagedBotToken.
type GetManagedBotTokenParams struct {
	UserID int64 `json:"user_id"`
}

// GetManagedBotToken gets the token of a managed bot.
func (c *Client) GetManagedBotToken(ctx context.Context, p *GetManagedBotTokenParams) (string, error) {
	var s string
	if err := c.Invoke(ctx, "getManagedBotToken", p, &s); err != nil {
		return "", err
	}
	return s, nil
}

// ReplaceManagedBotTokenParams holds parameters for replaceManagedBotToken.
type ReplaceManagedBotTokenParams struct {
	UserID int64 `json:"user_id"`
}

// ReplaceManagedBotToken revokes the current token of a managed bot and generates a new one.
func (c *Client) ReplaceManagedBotToken(ctx context.Context, p *ReplaceManagedBotTokenParams) (string, error) {
	var s string
	if err := c.Invoke(ctx, "replaceManagedBotToken", p, &s); err != nil {
		return "", err
	}
	return s, nil
}

// GetManagedBotAccessSettingsParams holds parameters for getManagedBotAccessSettings.
type GetManagedBotAccessSettingsParams struct {
	UserID int64 `json:"user_id"`
}

// GetManagedBotAccessSettings gets the access settings of a managed bot.
func (c *Client) GetManagedBotAccessSettings(ctx context.Context, p *GetManagedBotAccessSettingsParams) (*BotAccessSettings, error) {
	var s BotAccessSettings
	if err := c.Invoke(ctx, "getManagedBotAccessSettings", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// SetManagedBotAccessSettingsParams holds parameters for setManagedBotAccessSettings.
type SetManagedBotAccessSettingsParams struct {
	UserID             int64    `json:"user_id"`
	IsAccessRestricted bool     `json:"is_access_restricted"`
	AddedUserIDs       *[]int64 `json:"added_user_ids,omitempty"`
}

// SetManagedBotAccessSettings changes the access settings of a managed bot.
func (c *Client) SetManagedBotAccessSettings(ctx context.Context, p *SetManagedBotAccessSettingsParams) error {
	return c.Invoke(ctx, "setManagedBotAccessSettings", p, nil)
}

// SetMyCommandsParams holds parameters for setMyCommands.
type SetMyCommandsParams struct {
	Commands     []BotCommand    `json:"commands"`
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode *string         `json:"language_code,omitempty"`
}

// SetMyCommands changes the list of the bot's commands.
func (c *Client) SetMyCommands(ctx context.Context, p *SetMyCommandsParams) error {
	return c.Invoke(ctx, "setMyCommands", p, nil)
}

// DeleteMyCommandsParams holds parameters for deleteMyCommands.
type DeleteMyCommandsParams struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode *string         `json:"language_code,omitempty"`
}

// DeleteMyCommands deletes the list of the bot's commands for the given scope and user language.
func (c *Client) DeleteMyCommands(ctx context.Context, p *DeleteMyCommandsParams) error {
	return c.Invoke(ctx, "deleteMyCommands", p, nil)
}

// GetMyCommandsParams holds parameters for getMyCommands.
type GetMyCommandsParams struct {
	Scope        BotCommandScope `json:"scope,omitempty"`
	LanguageCode *string         `json:"language_code,omitempty"`
}

// GetMyCommands gets the current list of the bot's commands for the given scope and user language.
func (c *Client) GetMyCommands(ctx context.Context, p *GetMyCommandsParams) ([]BotCommand, error) {
	var cs []BotCommand
	if err := c.Invoke(ctx, "getMyCommands", p, &cs); err != nil {
		return nil, err
	}
	return cs, nil
}

// SetMyNameParams holds parameters for setMyName.
type SetMyNameParams struct {
	Name         *string `json:"name,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"`
}

// SetMyName changes the bot's name.
func (c *Client) SetMyName(ctx context.Context, p *SetMyNameParams) error {
	return c.Invoke(ctx, "setMyName", p, nil)
}

// GetMyNameParams holds parameters for getMyName.
type GetMyNameParams struct {
	LanguageCode *string `json:"language_code,omitempty"`
}

// GetMyName gets the current bot name for the given user language.
func (c *Client) GetMyName(ctx context.Context, p *GetMyNameParams) (*BotName, error) {
	var n BotName
	if err := c.Invoke(ctx, "getMyName", p, &n); err != nil {
		return nil, err
	}
	return &n, nil
}

// SetMyDescriptionParams holds parameters for setMyDescription.
type SetMyDescriptionParams struct {
	Description  *string `json:"description,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"`
}

// SetMyDescription changes the bot's description, which is shown in the chat with the bot if the chat is empty.
func (c *Client) SetMyDescription(ctx context.Context, p *SetMyDescriptionParams) error {
	return c.Invoke(ctx, "setMyDescription", p, nil)
}

// GetMyDescriptionParams holds parameters for getMyDescription.
type GetMyDescriptionParams struct {
	LanguageCode *string `json:"language_code,omitempty"`
}

// GetMyDescription gets the current bot description for the given user language.
func (c *Client) GetMyDescription(ctx context.Context, p *GetMyDescriptionParams) (*BotDescription, error) {
	var d BotDescription
	if err := c.Invoke(ctx, "getMyDescription", p, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// SetMyShortDescriptionParams holds parameters for setMyShortDescription.
type SetMyShortDescriptionParams struct {
	ShortDescription *string `json:"short_description,omitempty"`
	LanguageCode     *string `json:"language_code,omitempty"`
}

// SetMyShortDescription changes the bot's short description, which is shown on the bot's profile page.
func (c *Client) SetMyShortDescription(ctx context.Context, p *SetMyShortDescriptionParams) error {
	return c.Invoke(ctx, "setMyShortDescription", p, nil)
}

// GetMyShortDescriptionParams holds parameters for getMyShortDescription.
type GetMyShortDescriptionParams struct {
	LanguageCode *string `json:"language_code,omitempty"`
}

// GetMyShortDescription gets the current bot short description for the given user language.
func (c *Client) GetMyShortDescription(ctx context.Context, p *GetMyShortDescriptionParams) (*BotShortDescription, error) {
	var d BotShortDescription
	if err := c.Invoke(ctx, "getMyShortDescription", p, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// SetMyProfilePhotoParams holds parameters for setMyProfilePhoto.
type SetMyProfilePhotoParams struct {
	Photo InputProfilePhoto `json:"photo"`
}

// SetMyProfilePhoto changes the profile photo of the bot.
func (c *Client) SetMyProfilePhoto(ctx context.Context, p *SetMyProfilePhotoParams) error {
	return c.Invoke(ctx, "setMyProfilePhoto", p, nil)
}

// RemoveMyProfilePhoto removes the profile photo of the bot. Requires no parameters.
func (c *Client) RemoveMyProfilePhoto(ctx context.Context) error {
	return c.Invoke(ctx, "removeMyProfilePhoto", nil, nil)
}

// SetChatMenuButtonParams holds parameters for setChatMenuButton.
type SetChatMenuButtonParams struct {
	ChatID     *int64     `json:"chat_id,omitempty"`
	MenuButton MenuButton `json:"menu_button,omitempty"`
}

// SetChatMenuButton changes the bot's menu button in a private chat, or the default menu button.
func (c *Client) SetChatMenuButton(ctx context.Context, p *SetChatMenuButtonParams) error {
	return c.Invoke(ctx, "setChatMenuButton", p, nil)
}

// GetChatMenuButtonParams holds parameters for getChatMenuButton.
type GetChatMenuButtonParams struct {
	ChatID *int64 `json:"chat_id,omitempty"`
}

// GetChatMenuButton gets the current value of the bot's menu button in a private chat, or the default menu button.
func (c *Client) GetChatMenuButton(ctx context.Context, p *GetChatMenuButtonParams) (MenuButton, error) {
	var m MenuButton
	if err := c.Invoke(ctx, "getChatMenuButton", p, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// SetMyDefaultAdministratorRightsParams holds parameters for setMyDefaultAdministratorRights.
type SetMyDefaultAdministratorRightsParams struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels *bool                    `json:"for_channels,omitempty"`
}

// SetMyDefaultAdministratorRights changes the default administrator rights requested by the bot when it's added as an administrator to groups or channels.
func (c *Client) SetMyDefaultAdministratorRights(ctx context.Context, p *SetMyDefaultAdministratorRightsParams) error {
	return c.Invoke(ctx, "setMyDefaultAdministratorRights", p, nil)
}

// GetMyDefaultAdministratorRightsParams holds parameters for getMyDefaultAdministratorRights.
type GetMyDefaultAdministratorRightsParams struct {
	ForChannels *bool `json:"for_channels,omitempty"`
}

// GetMyDefaultAdministratorRights gets the current default administrator rights of the bot.
func (c *Client) GetMyDefaultAdministratorRights(ctx context.Context, p *GetMyDefaultAdministratorRightsParams) (*ChatAdministratorRights, error) {
	var r ChatAdministratorRights
	if err := c.Invoke(ctx, "getMyDefaultAdministratorRights", p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// GetAvailableGifts returns the list of gifts that can be sent by the bot
// to users and channel chats.
func (c *Client) GetAvailableGifts(ctx context.Context) (*Gifts, error) {
	var g Gifts
	if err := c.Invoke(ctx, "getAvailableGifts", nil, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

// SendGift sends a gift to the given user or channel chat. The gift can't
// be converted to Telegram Stars by the receiver.
func (c *Client) SendGift(ctx context.Context, p *SendGiftParams) error {
	return c.Invoke(ctx, "sendGift", p, nil)
}

// SendGiftParams holds the parameters for sendGift.
type SendGiftParams struct {
	UserID        *int64          `json:"user_id,omitempty"`
	ChatID        *ChatID         `json:"chat_id,omitempty"`
	GiftID        string          `json:"gift_id"`
	PayForUpgrade *bool           `json:"pay_for_upgrade,omitempty"`
	Text          *string         `json:"text,omitempty"`
	TextParseMode *string         `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// GiftPremiumSubscription gifts a Telegram Premium subscription to the
// given user.
func (c *Client) GiftPremiumSubscription(ctx context.Context, p *GiftPremiumSubscriptionParams) error {
	return c.Invoke(ctx, "giftPremiumSubscription", p, nil)
}

// GiftPremiumSubscriptionParams holds the parameters for
// giftPremiumSubscription.
type GiftPremiumSubscriptionParams struct {
	UserID        int64           `json:"user_id"`
	MonthCount    int             `json:"month_count"`
	StarCount     int64           `json:"star_count"`
	Text          *string         `json:"text,omitempty"`
	TextParseMode *string         `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// VerifyUser verifies a user on behalf of the organization which is
// represented by the bot.
func (c *Client) VerifyUser(ctx context.Context, p *VerifyUserParams) error {
	return c.Invoke(ctx, "verifyUser", p, nil)
}

// VerifyUserParams holds the parameters for verifyUser.
type VerifyUserParams struct {
	UserID            int64   `json:"user_id"`
	CustomDescription *string `json:"custom_description,omitempty"`
}

// VerifyChat verifies a chat on behalf of the organization which is
// represented by the bot.
func (c *Client) VerifyChat(ctx context.Context, p *VerifyChatParams) error {
	return c.Invoke(ctx, "verifyChat", p, nil)
}

// VerifyChatParams holds the parameters for verifyChat.
type VerifyChatParams struct {
	ChatID            ChatID  `json:"chat_id"`
	CustomDescription *string `json:"custom_description,omitempty"`
}

// RemoveUserVerification removes verification from a user who is currently
// verified on behalf of the organization represented by the bot.
func (c *Client) RemoveUserVerification(ctx context.Context, p *RemoveUserVerificationParams) error {
	return c.Invoke(ctx, "removeUserVerification", p, nil)
}

// RemoveUserVerificationParams holds the parameters for
// removeUserVerification.
type RemoveUserVerificationParams struct {
	UserID int64 `json:"user_id"`
}

// RemoveChatVerification removes verification from a chat that is currently
// verified on behalf of the organization represented by the bot.
func (c *Client) RemoveChatVerification(ctx context.Context, p *RemoveChatVerificationParams) error {
	return c.Invoke(ctx, "removeChatVerification", p, nil)
}

// RemoveChatVerificationParams holds the parameters for
// removeChatVerification.
type RemoveChatVerificationParams struct {
	ChatID ChatID `json:"chat_id"`
}

// GetUserGifts returns the gifts owned and hosted by a user.
func (c *Client) GetUserGifts(ctx context.Context, p *GetUserGiftsParams) (*OwnedGifts, error) {
	var g OwnedGifts
	if err := c.Invoke(ctx, "getUserGifts", p, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

// GetUserGiftsParams holds the parameters for getUserGifts.
type GetUserGiftsParams struct {
	UserID                      int64   `json:"user_id"`
	ExcludeUnlimited            *bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    *bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable *bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeFromBlockchain       *bool   `json:"exclude_from_blockchain,omitempty"`
	ExcludeUnique               *bool   `json:"exclude_unique,omitempty"`
	SortByPrice                 *bool   `json:"sort_by_price,omitempty"`
	Offset                      *string `json:"offset,omitempty"`
	Limit                       *int    `json:"limit,omitempty"`
}

// GetChatGifts returns the gifts owned by a chat.
func (c *Client) GetChatGifts(ctx context.Context, p *GetChatGiftsParams) (*OwnedGifts, error) {
	var g OwnedGifts
	if err := c.Invoke(ctx, "getChatGifts", p, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

// GetChatGiftsParams holds the parameters for getChatGifts.
type GetChatGiftsParams struct {
	ChatID                      ChatID  `json:"chat_id"`
	ExcludeUnsaved              *bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                *bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            *bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    *bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable *bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeFromBlockchain       *bool   `json:"exclude_from_blockchain,omitempty"`
	ExcludeUnique               *bool   `json:"exclude_unique,omitempty"`
	SortByPrice                 *bool   `json:"sort_by_price,omitempty"`
	Offset                      *string `json:"offset,omitempty"`
	Limit                       *int    `json:"limit,omitempty"`
}

// ConvertGiftToStars converts a given regular gift to Telegram Stars.
// Requires the can_convert_gifts_to_stars business bot right.
func (c *Client) ConvertGiftToStars(ctx context.Context, p *ConvertGiftToStarsParams) error {
	return c.Invoke(ctx, "convertGiftToStars", p, nil)
}

// ConvertGiftToStarsParams holds the parameters for convertGiftToStars.
type ConvertGiftToStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
}

// UpgradeGift upgrades a given regular gift to a unique gift. Requires the
// can_transfer_and_upgrade_gifts business bot right; additionally requires
// the can_transfer_stars business bot right if the upgrade is paid.
func (c *Client) UpgradeGift(ctx context.Context, p *UpgradeGiftParams) error {
	return c.Invoke(ctx, "upgradeGift", p, nil)
}

// UpgradeGiftParams holds the parameters for upgradeGift.
type UpgradeGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	KeepOriginalDetails  *bool  `json:"keep_original_details,omitempty"`
	StarCount            *int64 `json:"star_count,omitempty"`
}

// TransferGift transfers an owned unique gift to another user. Requires the
// can_transfer_and_upgrade_gifts business bot right; requires the
// can_transfer_stars business bot right if the transfer is paid.
func (c *Client) TransferGift(ctx context.Context, p *TransferGiftParams) error {
	return c.Invoke(ctx, "transferGift", p, nil)
}

// TransferGiftParams holds the parameters for transferGift.
type TransferGiftParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	OwnedGiftID          string `json:"owned_gift_id"`
	NewOwnerChatID       int64  `json:"new_owner_chat_id"`
	StarCount            *int64 `json:"star_count,omitempty"`
}

// PostStory posts a story on behalf of a managed business account. Requires
// the can_manage_stories business bot right.
func (c *Client) PostStory(ctx context.Context, p *PostStoryParams) (*Story, error) {
	var s Story
	if err := c.Invoke(ctx, "postStory", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// PostStoryParams holds the parameters for postStory.
type PostStoryParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	Content              InputStoryContent `json:"content"`
	ActivePeriod         int               `json:"active_period"`
	Caption              *string           `json:"caption,omitempty"`
	ParseMode            *string           `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []StoryArea       `json:"areas,omitempty"`
	PostToChatPage       *bool             `json:"post_to_chat_page,omitempty"`
	ProtectContent       *bool             `json:"protect_content,omitempty"`
}

// RepostStory reposts a story on behalf of a business account from another
// business account. Both business accounts must be managed by the same bot.
// Requires the can_manage_stories business bot right for both business
// accounts.
func (c *Client) RepostStory(ctx context.Context, p *RepostStoryParams) (*Story, error) {
	var s Story
	if err := c.Invoke(ctx, "repostStory", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// RepostStoryParams holds the parameters for repostStory.
type RepostStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	FromChatID           int64  `json:"from_chat_id"`
	FromStoryID          int64  `json:"from_story_id"`
	ActivePeriod         int    `json:"active_period"`
	PostToChatPage       *bool  `json:"post_to_chat_page,omitempty"`
	ProtectContent       *bool  `json:"protect_content,omitempty"`
}

// EditStory edits a story previously posted by the bot on behalf of a
// managed business account. Requires the can_manage_stories business bot
// right.
func (c *Client) EditStory(ctx context.Context, p *EditStoryParams) (*Story, error) {
	var s Story
	if err := c.Invoke(ctx, "editStory", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// EditStoryParams holds the parameters for editStory.
type EditStoryParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	StoryID              int64             `json:"story_id"`
	Content              InputStoryContent `json:"content,omitempty"`
	Caption              *string           `json:"caption,omitempty"`
	ParseMode            *string           `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity   `json:"caption_entities,omitempty"`
	Areas                []StoryArea       `json:"areas,omitempty"`
}

// DeleteStory deletes a story previously posted by the bot on behalf of a
// managed business account. Requires the can_manage_stories business bot
// right.
func (c *Client) DeleteStory(ctx context.Context, p *DeleteStoryParams) error {
	return c.Invoke(ctx, "deleteStory", p, nil)
}

// DeleteStoryParams holds the parameters for deleteStory.
type DeleteStoryParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StoryID              int64  `json:"story_id"`
}

// ReadBusinessMessageParams holds parameters for readBusinessMessage.
type ReadBusinessMessageParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ChatID               int64  `json:"chat_id"`
	MessageID            int64  `json:"message_id"`
}

// ReadBusinessMessage marks an incoming message as read on behalf of a business account.
func (c *Client) ReadBusinessMessage(ctx context.Context, p *ReadBusinessMessageParams) error {
	return c.Invoke(ctx, "readBusinessMessage", p, nil)
}

// DeleteBusinessMessagesParams holds parameters for deleteBusinessMessages.
type DeleteBusinessMessagesParams struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	MessageIDs           []int64 `json:"message_ids"`
}

// DeleteBusinessMessages deletes messages on behalf of a business account.
func (c *Client) DeleteBusinessMessages(ctx context.Context, p *DeleteBusinessMessagesParams) error {
	return c.Invoke(ctx, "deleteBusinessMessages", p, nil)
}

// SetBusinessAccountNameParams holds parameters for setBusinessAccountName.
type SetBusinessAccountNameParams struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	FirstName            string  `json:"first_name"`
	LastName             *string `json:"last_name,omitempty"`
}

// SetBusinessAccountName changes the first and last name of a managed business account.
func (c *Client) SetBusinessAccountName(ctx context.Context, p *SetBusinessAccountNameParams) error {
	return c.Invoke(ctx, "setBusinessAccountName", p, nil)
}

// SetBusinessAccountUsernameParams holds parameters for setBusinessAccountUsername.
type SetBusinessAccountUsernameParams struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	Username             *string `json:"username,omitempty"`
}

// SetBusinessAccountUsername changes the username of a managed business account.
func (c *Client) SetBusinessAccountUsername(ctx context.Context, p *SetBusinessAccountUsernameParams) error {
	return c.Invoke(ctx, "setBusinessAccountUsername", p, nil)
}

// SetBusinessAccountBioParams holds parameters for setBusinessAccountBio.
type SetBusinessAccountBioParams struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	Bio                  *string `json:"bio,omitempty"`
}

// SetBusinessAccountBio changes the bio of a managed business account.
func (c *Client) SetBusinessAccountBio(ctx context.Context, p *SetBusinessAccountBioParams) error {
	return c.Invoke(ctx, "setBusinessAccountBio", p, nil)
}

// SetBusinessAccountProfilePhotoParams holds parameters for setBusinessAccountProfilePhoto.
type SetBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	Photo                InputProfilePhoto `json:"photo"`
	IsPublic             *bool             `json:"is_public,omitempty"`
}

// SetBusinessAccountProfilePhoto changes the profile photo of a managed business account.
func (c *Client) SetBusinessAccountProfilePhoto(ctx context.Context, p *SetBusinessAccountProfilePhotoParams) error {
	return c.Invoke(ctx, "setBusinessAccountProfilePhoto", p, nil)
}

// RemoveBusinessAccountProfilePhotoParams holds parameters for removeBusinessAccountProfilePhoto.
type RemoveBusinessAccountProfilePhotoParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	IsPublic             *bool  `json:"is_public,omitempty"`
}

// RemoveBusinessAccountProfilePhoto removes the current profile photo of a managed business account.
func (c *Client) RemoveBusinessAccountProfilePhoto(ctx context.Context, p *RemoveBusinessAccountProfilePhotoParams) error {
	return c.Invoke(ctx, "removeBusinessAccountProfilePhoto", p, nil)
}

// SetBusinessAccountGiftSettingsParams holds parameters for setBusinessAccountGiftSettings.
type SetBusinessAccountGiftSettingsParams struct {
	BusinessConnectionID string            `json:"business_connection_id"`
	ShowGiftButton       bool              `json:"show_gift_button"`
	AcceptedGiftTypes    AcceptedGiftTypes `json:"accepted_gift_types"`
}

// SetBusinessAccountGiftSettings changes the privacy settings pertaining to incoming gifts in a managed business account.
func (c *Client) SetBusinessAccountGiftSettings(ctx context.Context, p *SetBusinessAccountGiftSettingsParams) error {
	return c.Invoke(ctx, "setBusinessAccountGiftSettings", p, nil)
}

// GetBusinessAccountStarBalanceParams holds parameters for getBusinessAccountStarBalance.
type GetBusinessAccountStarBalanceParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
}

// GetBusinessAccountStarBalance returns the amount of Telegram Stars owned by a managed business account.
func (c *Client) GetBusinessAccountStarBalance(ctx context.Context, p *GetBusinessAccountStarBalanceParams) (*StarAmount, error) {
	var a StarAmount
	if err := c.Invoke(ctx, "getBusinessAccountStarBalance", p, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

// TransferBusinessAccountStarsParams holds parameters for transferBusinessAccountStars.
type TransferBusinessAccountStarsParams struct {
	BusinessConnectionID string `json:"business_connection_id"`
	StarCount            int64  `json:"star_count"`
}

// TransferBusinessAccountStars transfers Telegram Stars from the business account balance to the bot's balance.
func (c *Client) TransferBusinessAccountStars(ctx context.Context, p *TransferBusinessAccountStarsParams) error {
	return c.Invoke(ctx, "transferBusinessAccountStars", p, nil)
}

// GetBusinessAccountGiftsParams holds parameters for getBusinessAccountGifts.
type GetBusinessAccountGiftsParams struct {
	BusinessConnectionID        string  `json:"business_connection_id"`
	ExcludeUnsaved              *bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved                *bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited            *bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimitedUpgradable    *bool   `json:"exclude_limited_upgradable,omitempty"`
	ExcludeLimitedNonUpgradable *bool   `json:"exclude_limited_non_upgradable,omitempty"`
	ExcludeUnique               *bool   `json:"exclude_unique,omitempty"`
	ExcludeFromBlockchain       *bool   `json:"exclude_from_blockchain,omitempty"`
	SortByPrice                 *bool   `json:"sort_by_price,omitempty"`
	Offset                      *string `json:"offset,omitempty"`
	Limit                       *int    `json:"limit,omitempty"`
}

// GetBusinessAccountGifts returns the gifts received and owned by a managed business account.
func (c *Client) GetBusinessAccountGifts(ctx context.Context, p *GetBusinessAccountGiftsParams) (*OwnedGifts, error) {
	var g OwnedGifts
	if err := c.Invoke(ctx, "getBusinessAccountGifts", p, &g); err != nil {
		return nil, err
	}
	return &g, nil
}

// GetUserChatBoostsParams holds parameters for getUserChatBoosts.
type GetUserChatBoostsParams struct {
	ChatID ChatID `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

// GetUserChatBoosts gets the list of boosts added to a chat by a user. Requires administrator rights in the chat.
func (c *Client) GetUserChatBoosts(ctx context.Context, p *GetUserChatBoostsParams) (*UserChatBoosts, error) {
	var b UserChatBoosts
	if err := c.Invoke(ctx, "getUserChatBoosts", p, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

// SendSticker sends static .WEBP, animated .TGS, or video .WEBM stickers.
// On success, the sent Message is returned.
func (c *Client) SendSticker(ctx context.Context, p *SendStickerParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendSticker", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendStickerParams holds the parameters for sendSticker.
type SendStickerParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	Sticker                    InputFile                   `json:"sticker"`
	Emoji                      *string                     `json:"emoji,omitempty"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// GetStickerSet gets a sticker set. On success, a StickerSet object is
// returned.
func (c *Client) GetStickerSet(ctx context.Context, p *GetStickerSetParams) (*StickerSet, error) {
	var s StickerSet
	if err := c.Invoke(ctx, "getStickerSet", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// GetStickerSetParams holds the parameters for getStickerSet.
type GetStickerSetParams struct {
	Name string `json:"name"`
}

// GetCustomEmojiStickers gets information about custom emoji stickers by
// their identifiers. Returns an Array of Sticker objects.
func (c *Client) GetCustomEmojiStickers(ctx context.Context, p *GetCustomEmojiStickersParams) ([]Sticker, error) {
	var s []Sticker
	if err := c.Invoke(ctx, "getCustomEmojiStickers", p, &s); err != nil {
		return nil, err
	}
	return s, nil
}

// GetCustomEmojiStickersParams holds the parameters for
// getCustomEmojiStickers.
type GetCustomEmojiStickersParams struct {
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

// UploadStickerFile uploads a file with a sticker for later use in
// createNewStickerSet, addStickerToSet, or replaceStickerInSet. Returns the
// uploaded File on success.
func (c *Client) UploadStickerFile(ctx context.Context, p *UploadStickerFileParams) (*File, error) {
	var f File
	if err := c.Invoke(ctx, "uploadStickerFile", p, &f); err != nil {
		return nil, err
	}
	return &f, nil
}

// UploadStickerFileParams holds the parameters for uploadStickerFile.
type UploadStickerFileParams struct {
	UserID        int64     `json:"user_id"`
	Sticker       InputFile `json:"sticker"`
	StickerFormat string    `json:"sticker_format"`
}

// CreateNewStickerSet creates a new sticker set owned by a user. The bot
// will be able to edit the sticker set thus created.
func (c *Client) CreateNewStickerSet(ctx context.Context, p *CreateNewStickerSetParams) error {
	return c.Invoke(ctx, "createNewStickerSet", p, nil)
}

// CreateNewStickerSetParams holds the parameters for createNewStickerSet.
type CreateNewStickerSetParams struct {
	UserID          int64          `json:"user_id"`
	Name            string         `json:"name"`
	Title           string         `json:"title"`
	Stickers        []InputSticker `json:"stickers"`
	StickerType     *string        `json:"sticker_type,omitempty"`
	NeedsRepainting *bool          `json:"needs_repainting,omitempty"`
}

// AddStickerToSet adds a new sticker to a set created by the bot.
func (c *Client) AddStickerToSet(ctx context.Context, p *AddStickerToSetParams) error {
	return c.Invoke(ctx, "addStickerToSet", p, nil)
}

// AddStickerToSetParams holds the parameters for addStickerToSet.
type AddStickerToSetParams struct {
	UserID  int64        `json:"user_id"`
	Name    string       `json:"name"`
	Sticker InputSticker `json:"sticker"`
}

// SetStickerPositionInSet moves a sticker in a set created by the bot to a
// specific position.
func (c *Client) SetStickerPositionInSet(ctx context.Context, p *SetStickerPositionInSetParams) error {
	return c.Invoke(ctx, "setStickerPositionInSet", p, nil)
}

// SetStickerPositionInSetParams holds the parameters for
// setStickerPositionInSet.
type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

// DeleteStickerFromSet deletes a sticker from a set created by the bot.
func (c *Client) DeleteStickerFromSet(ctx context.Context, p *DeleteStickerFromSetParams) error {
	return c.Invoke(ctx, "deleteStickerFromSet", p, nil)
}

// DeleteStickerFromSetParams holds the parameters for deleteStickerFromSet.
type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"`
}

// ReplaceStickerInSet replaces an existing sticker in a sticker set with a
// new one.
func (c *Client) ReplaceStickerInSet(ctx context.Context, p *ReplaceStickerInSetParams) error {
	return c.Invoke(ctx, "replaceStickerInSet", p, nil)
}

// ReplaceStickerInSetParams holds the parameters for replaceStickerInSet.
type ReplaceStickerInSetParams struct {
	UserID     int64        `json:"user_id"`
	Name       string       `json:"name"`
	OldSticker string       `json:"old_sticker"`
	Sticker    InputSticker `json:"sticker"`
}

// SetStickerEmojiList changes the list of emoji assigned to a regular or
// custom emoji sticker.
func (c *Client) SetStickerEmojiList(ctx context.Context, p *SetStickerEmojiListParams) error {
	return c.Invoke(ctx, "setStickerEmojiList", p, nil)
}

// SetStickerEmojiListParams holds the parameters for setStickerEmojiList.
type SetStickerEmojiListParams struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

// SetStickerKeywords changes search keywords assigned to a regular or
// custom emoji sticker.
func (c *Client) SetStickerKeywords(ctx context.Context, p *SetStickerKeywordsParams) error {
	return c.Invoke(ctx, "setStickerKeywords", p, nil)
}

// SetStickerKeywordsParams holds the parameters for setStickerKeywords.
type SetStickerKeywordsParams struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

// SetStickerMaskPosition changes the mask position of a mask sticker.
func (c *Client) SetStickerMaskPosition(ctx context.Context, p *SetStickerMaskPositionParams) error {
	return c.Invoke(ctx, "setStickerMaskPosition", p, nil)
}

// SetStickerMaskPositionParams holds the parameters for
// setStickerMaskPosition.
type SetStickerMaskPositionParams struct {
	Sticker      string        `json:"sticker"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// SetStickerSetTitle sets the title of a created sticker set.
func (c *Client) SetStickerSetTitle(ctx context.Context, p *SetStickerSetTitleParams) error {
	return c.Invoke(ctx, "setStickerSetTitle", p, nil)
}

// SetStickerSetTitleParams holds the parameters for setStickerSetTitle.
type SetStickerSetTitleParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

// SetStickerSetThumbnail sets the thumbnail of a regular or mask sticker
// set. The format of the thumbnail file must match the format of the
// stickers in the set.
func (c *Client) SetStickerSetThumbnail(ctx context.Context, p *SetStickerSetThumbnailParams) error {
	return c.Invoke(ctx, "setStickerSetThumbnail", p, nil)
}

// SetStickerSetThumbnailParams holds the parameters for
// setStickerSetThumbnail.
type SetStickerSetThumbnailParams struct {
	Name      string     `json:"name"`
	UserID    int64      `json:"user_id"`
	Thumbnail *InputFile `json:"thumbnail,omitempty"`
	Format    string     `json:"format"`
}

// SetCustomEmojiStickerSetThumbnail sets the thumbnail of a custom emoji
// sticker set.
func (c *Client) SetCustomEmojiStickerSetThumbnail(ctx context.Context, p *SetCustomEmojiStickerSetThumbnailParams) error {
	return c.Invoke(ctx, "setCustomEmojiStickerSetThumbnail", p, nil)
}

// SetCustomEmojiStickerSetThumbnailParams holds the parameters for
// setCustomEmojiStickerSetThumbnail.
type SetCustomEmojiStickerSetThumbnailParams struct {
	Name          string  `json:"name"`
	CustomEmojiID *string `json:"custom_emoji_id,omitempty"`
}

// DeleteStickerSet deletes a sticker set that was created by the bot.
func (c *Client) DeleteStickerSet(ctx context.Context, p *DeleteStickerSetParams) error {
	return c.Invoke(ctx, "deleteStickerSet", p, nil)
}

// DeleteStickerSetParams holds the parameters for deleteStickerSet.
type DeleteStickerSetParams struct {
	Name string `json:"name"`
}

// SendRichMessage sends a rich message. If the message contains a block with
// a media element, the bot must have the right to send the media to the chat.
// On success, the sent Message is returned.
func (c *Client) SendRichMessage(ctx context.Context, p *SendRichMessageParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendRichMessage", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendRichMessageParams holds the parameters for sendRichMessage.
type SendRichMessageParams struct {
	BusinessConnectionID       *string                     `json:"business_connection_id,omitempty"`
	ChatID                     ChatID                      `json:"chat_id"`
	MessageThreadID            *int64                      `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID      *int64                      `json:"direct_messages_topic_id,omitempty"`
	EphemeralMessageParameters *EphemeralMessageParameters `json:"ephemeral_message_parameters,omitempty"`
	RichMessage                InputRichMessage            `json:"rich_message"`
	DisableNotification        *bool                       `json:"disable_notification,omitempty"`
	ProtectContent             *bool                       `json:"protect_content,omitempty"`
	AllowPaidBroadcast         *bool                       `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID            *string                     `json:"message_effect_id,omitempty"`
	SuggestedPostParameters    *SuggestedPostParameters    `json:"suggested_post_parameters,omitempty"`
	ReplyParameters            *ReplyParameters            `json:"reply_parameters,omitempty"`
	ReplyMarkup                ReplyMarkup                 `json:"reply_markup,omitempty"`
}

// SendRichMessageDraft streams a partial rich message to a user while the
// message is being generated; once the output is finalized, sendRichMessage
// must be called with the complete message to persist it. Returns True on
// success.
func (c *Client) SendRichMessageDraft(ctx context.Context, p *SendRichMessageDraftParams) error {
	return c.Invoke(ctx, "sendRichMessageDraft", p, nil)
}

// SendRichMessageDraftParams holds the parameters for sendRichMessageDraft.
type SendRichMessageDraftParams struct {
	ChatID          int64            `json:"chat_id"`
	MessageThreadID *int64           `json:"message_thread_id,omitempty"`
	DraftID         int64            `json:"draft_id"`
	RichMessage     InputRichMessage `json:"rich_message"`
	CanStop         *bool            `json:"can_stop,omitempty"`
	KeepOnStop      *bool            `json:"keep_on_stop,omitempty"`
}

// AnswerInlineQuery sends answers to an inline query. On success, True is
// returned. No more than 50 results per query are allowed.
func (c *Client) AnswerInlineQuery(ctx context.Context, p *AnswerInlineQueryParams) error {
	return c.Invoke(ctx, "answerInlineQuery", p, nil)
}

// AnswerInlineQueryParams holds the parameters for answerInlineQuery.
type AnswerInlineQueryParams struct {
	InlineQueryID string                    `json:"inline_query_id"`
	Results       []InlineQueryResult       `json:"results"`
	CacheTime     *int                      `json:"cache_time,omitempty"`
	IsPersonal    *bool                     `json:"is_personal,omitempty"`
	NextOffset    *string                   `json:"next_offset,omitempty"`
	Button        *InlineQueryResultsButton `json:"button,omitempty"`
}

// SendInvoice sends invoices. On success, the sent Message is returned.
func (c *Client) SendInvoice(ctx context.Context, p *SendInvoiceParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendInvoice", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendInvoiceParams holds the parameters for sendInvoice.
type SendInvoiceParams struct {
	ChatID                    ChatID                   `json:"chat_id"`
	MessageThreadID           *int64                   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicID     *int64                   `json:"direct_messages_topic_id,omitempty"`
	Title                     string                   `json:"title"`
	Description               string                   `json:"description"`
	Payload                   string                   `json:"payload"`
	ProviderToken             *string                  `json:"provider_token,omitempty"`
	Currency                  string                   `json:"currency"`
	Prices                    []LabeledPrice           `json:"prices"`
	MaxTipAmount              *int64                   `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64                  `json:"suggested_tip_amounts,omitempty"`
	StartParameter            *string                  `json:"start_parameter,omitempty"`
	ProviderData              *string                  `json:"provider_data,omitempty"`
	PhotoURL                  *string                  `json:"photo_url,omitempty"`
	PhotoSize                 *int                     `json:"photo_size,omitempty"`
	PhotoWidth                *int                     `json:"photo_width,omitempty"`
	PhotoHeight               *int                     `json:"photo_height,omitempty"`
	NeedName                  *bool                    `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool                    `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool                    `json:"need_email,omitempty"`
	NeedShippingAddress       *bool                    `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool                    `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool                    `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool                    `json:"is_flexible,omitempty"`
	DisableNotification       *bool                    `json:"disable_notification,omitempty"`
	ProtectContent            *bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast        *bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID           *string                  `json:"message_effect_id,omitempty"`
	SuggestedPostParameters   *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters           *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup    `json:"reply_markup,omitempty"`
}

// CreateInvoiceLink creates a link for an invoice. Returns the created
// invoice link as String on success.
func (c *Client) CreateInvoiceLink(ctx context.Context, p *CreateInvoiceLinkParams) (string, error) {
	var s string
	if err := c.Invoke(ctx, "createInvoiceLink", p, &s); err != nil {
		return "", err
	}
	return s, nil
}

// CreateInvoiceLinkParams holds the parameters for createInvoiceLink.
type CreateInvoiceLinkParams struct {
	BusinessConnectionID      *string        `json:"business_connection_id,omitempty"`
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             *string        `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	SubscriptionPeriod        *int           `json:"subscription_period,omitempty"`
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

// AnswerShippingQuery replies to shipping queries. On success, True is
// returned.
func (c *Client) AnswerShippingQuery(ctx context.Context, p *AnswerShippingQueryParams) error {
	return c.Invoke(ctx, "answerShippingQuery", p, nil)
}

// AnswerShippingQueryParams holds the parameters for answerShippingQuery.
type AnswerShippingQueryParams struct {
	ShippingQueryID string           `json:"shipping_query_id"`
	OK              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    *string          `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQuery responds to pre-checkout queries. On success, True
// is returned.
func (c *Client) AnswerPreCheckoutQuery(ctx context.Context, p *AnswerPreCheckoutQueryParams) error {
	return c.Invoke(ctx, "answerPreCheckoutQuery", p, nil)
}

// AnswerPreCheckoutQueryParams holds the parameters for answerPreCheckoutQuery.
type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryID string  `json:"pre_checkout_query_id"`
	OK                 bool    `json:"ok"`
	ErrorMessage       *string `json:"error_message,omitempty"`
}

// GetMyStarBalance gets the current Telegram Stars balance of the bot. On
// success, returns a StarAmount object.
func (c *Client) GetMyStarBalance(ctx context.Context, p *GetMyStarBalanceParams) (*StarAmount, error) {
	var s StarAmount
	if err := c.Invoke(ctx, "getMyStarBalance", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// GetMyStarBalanceParams holds the parameters for getMyStarBalance.
type GetMyStarBalanceParams struct{}

// GetStarTransactions returns the bot's Telegram Star transactions in
// chronological order. On success, returns a StarTransactions object.
func (c *Client) GetStarTransactions(ctx context.Context, p *GetStarTransactionsParams) (*StarTransactions, error) {
	var s StarTransactions
	if err := c.Invoke(ctx, "getStarTransactions", p, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// GetStarTransactionsParams holds the parameters for getStarTransactions.
type GetStarTransactionsParams struct {
	Offset *int `json:"offset,omitempty"`
	Limit  *int `json:"limit,omitempty"`
}

// RefundStarPayment refunds a successful payment in Telegram Stars. Returns
// True on success.
func (c *Client) RefundStarPayment(ctx context.Context, p *RefundStarPaymentParams) error {
	return c.Invoke(ctx, "refundStarPayment", p, nil)
}

// RefundStarPaymentParams holds the parameters for refundStarPayment.
type RefundStarPaymentParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

// EditUserStarSubscription allows the bot to cancel or re-enable extension
// of a subscription paid in Telegram Stars. Returns True on success.
func (c *Client) EditUserStarSubscription(ctx context.Context, p *EditUserStarSubscriptionParams) error {
	return c.Invoke(ctx, "editUserStarSubscription", p, nil)
}

// EditUserStarSubscriptionParams holds the parameters for
// editUserStarSubscription.
type EditUserStarSubscriptionParams struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

// SetPassportDataErrors informs a user that some of the Telegram Passport
// elements they provided contains errors. Returns True on success.
func (c *Client) SetPassportDataErrors(ctx context.Context, p *SetPassportDataErrorsParams) error {
	return c.Invoke(ctx, "setPassportDataErrors", p, nil)
}

// SetPassportDataErrorsParams holds the parameters for setPassportDataErrors.
type SetPassportDataErrorsParams struct {
	UserID int64                  `json:"user_id"`
	Errors []PassportElementError `json:"errors"`
}

// SendGame sends a game. On success, the sent Message is returned.
func (c *Client) SendGame(ctx context.Context, p *SendGameParams) (*Message, error) {
	var m Message
	if err := c.Invoke(ctx, "sendGame", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// SendGameParams holds the parameters for sendGame.
type SendGameParams struct {
	BusinessConnectionID *string               `json:"business_connection_id,omitempty"`
	ChatID               ChatID                `json:"chat_id"`
	MessageThreadID      *int64                `json:"message_thread_id,omitempty"`
	GameShortName        string                `json:"game_short_name"`
	DisableNotification  *bool                 `json:"disable_notification,omitempty"`
	ProtectContent       *bool                 `json:"protect_content,omitempty"`
	AllowPaidBroadcast   *bool                 `json:"allow_paid_broadcast,omitempty"`
	MessageEffectID      *string               `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SetGameScore sets the score of the specified user in a game message. On
// success, if the message is not an inline message, the Message is returned,
// otherwise True is returned.
func (c *Client) SetGameScore(ctx context.Context, p *SetGameScoreParams) (*Message, error) {
	var raw json.RawMessage
	if err := c.Invoke(ctx, "setGameScore", p, &raw); err != nil {
		return nil, err
	}
	if string(bytes.TrimSpace(raw)) == "true" {
		return nil, nil
	}
	var m Message
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("telegram: decode setGameScore result: %w", err)
	}
	return &m, nil
}

// SetGameScoreParams holds the parameters for setGameScore.
type SetGameScoreParams struct {
	UserID             int64   `json:"user_id"`
	Score              int     `json:"score"`
	Force              *bool   `json:"force,omitempty"`
	DisableEditMessage *bool   `json:"disable_edit_message,omitempty"`
	ChatID             *int64  `json:"chat_id,omitempty"`
	MessageID          *int64  `json:"message_id,omitempty"`
	InlineMessageID    *string `json:"inline_message_id,omitempty"`
}

// GetGameHighScores gets data for high score tables. Returns an Array of
// GameHighScore objects.
func (c *Client) GetGameHighScores(ctx context.Context, p *GetGameHighScoresParams) ([]GameHighScore, error) {
	var hs []GameHighScore
	if err := c.Invoke(ctx, "getGameHighScores", p, &hs); err != nil {
		return nil, err
	}
	return hs, nil
}

// GetGameHighScoresParams holds the parameters for getGameHighScores.
type GetGameHighScoresParams struct {
	UserID          int64   `json:"user_id"`
	ChatID          *int64  `json:"chat_id,omitempty"`
	MessageID       *int64  `json:"message_id,omitempty"`
	InlineMessageID *string `json:"inline_message_id,omitempty"`
}
