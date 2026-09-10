package api

import (
	"context"
)

// BanChatMember bans a user in a group, supergroup or channel. Returns
// True on success.
func (c *Client) BanChatMember(ctx context.Context, p *BanChatMemberParams) error {
	return c.Invoke(ctx, "banChatMember", p, nil)
}

// BanChatMemberParams holds the parameters for banChatMember.
type BanChatMemberParams struct {
	ChatID         ChatID `json:"chat_id"`
	UserID         int64  `json:"user_id"`
	UntilDate      *int64 `json:"until_date,omitempty"`
	RevokeMessages *bool  `json:"revoke_messages,omitempty"`
}

// UnbanChatMember unbans a previously banned user in a supergroup or
// channel. Returns True on success.
func (c *Client) UnbanChatMember(ctx context.Context, p *UnbanChatMemberParams) error {
	return c.Invoke(ctx, "unbanChatMember", p, nil)
}

// UnbanChatMemberParams holds the parameters for unbanChatMember.
type UnbanChatMemberParams struct {
	ChatID       ChatID `json:"chat_id"`
	UserID       int64  `json:"user_id"`
	OnlyIfBanned *bool  `json:"only_if_banned,omitempty"`
}

// RestrictChatMember restricts a user in a supergroup. Returns True on
// success.
func (c *Client) RestrictChatMember(ctx context.Context, p *RestrictChatMemberParams) error {
	return c.Invoke(ctx, "restrictChatMember", p, nil)
}

// RestrictChatMemberParams holds the parameters for restrictChatMember.
type RestrictChatMemberParams struct {
	ChatID                        ChatID          `json:"chat_id"`
	UserID                        int64           `json:"user_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions *bool           `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     *int64          `json:"until_date,omitempty"`
}

// PromoteChatMember promotes or demotes a user in a supergroup or a
// channel. Returns True on success.
func (c *Client) PromoteChatMember(ctx context.Context, p *PromoteChatMemberParams) error {
	return c.Invoke(ctx, "promoteChatMember", p, nil)
}

// PromoteChatMemberParams holds the parameters for promoteChatMember.
type PromoteChatMemberParams struct {
	ChatID                  ChatID `json:"chat_id"`
	UserID                  int64  `json:"user_id"`
	IsAnonymous             *bool  `json:"is_anonymous,omitempty"`
	CanManageChat           *bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages       *bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     *bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      *bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       *bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo           *bool  `json:"can_change_info,omitempty"`
	CanInviteUsers          *bool  `json:"can_invite_users,omitempty"`
	CanPostStories          *bool  `json:"can_post_stories,omitempty"`
	CanEditStories          *bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories        *bool  `json:"can_delete_stories,omitempty"`
	CanPostMessages         *bool  `json:"can_post_messages,omitempty"`
	CanEditMessages         *bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages          *bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics         *bool  `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages *bool  `json:"can_manage_direct_messages,omitempty"`
	CanManageTags           *bool  `json:"can_manage_tags,omitempty"`
	CanSendWelcomeMessages  *bool  `json:"can_send_welcome_messages,omitempty"`
}

// SetChatAdministratorCustomTitle sets a custom title for an administrator
// in a supergroup promoted by the bot. Returns True on success.
func (c *Client) SetChatAdministratorCustomTitle(ctx context.Context, p *SetChatAdministratorCustomTitleParams) error {
	return c.Invoke(ctx, "setChatAdministratorCustomTitle", p, nil)
}

// SetChatAdministratorCustomTitleParams holds the parameters for
// setChatAdministratorCustomTitle.
type SetChatAdministratorCustomTitleParams struct {
	ChatID      ChatID `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

// SetChatMemberTag sets a tag for a regular member in a group or a
// supergroup. Returns True on success.
func (c *Client) SetChatMemberTag(ctx context.Context, p *SetChatMemberTagParams) error {
	return c.Invoke(ctx, "setChatMemberTag", p, nil)
}

// SetChatMemberTagParams holds the parameters for setChatMemberTag.
type SetChatMemberTagParams struct {
	ChatID ChatID  `json:"chat_id"`
	UserID int64   `json:"user_id"`
	Tag    *string `json:"tag,omitempty"`
}

// BanChatSenderChat bans a channel chat in a supergroup or a channel.
// Returns True on success.
func (c *Client) BanChatSenderChat(ctx context.Context, p *BanChatSenderChatParams) error {
	return c.Invoke(ctx, "banChatSenderChat", p, nil)
}

// BanChatSenderChatParams holds the parameters for banChatSenderChat.
type BanChatSenderChatParams struct {
	ChatID       ChatID `json:"chat_id"`
	SenderChatID int64  `json:"sender_chat_id"`
}

// UnbanChatSenderChat unbans a previously banned channel chat in a
// supergroup or channel. Returns True on success.
func (c *Client) UnbanChatSenderChat(ctx context.Context, p *UnbanChatSenderChatParams) error {
	return c.Invoke(ctx, "unbanChatSenderChat", p, nil)
}

// UnbanChatSenderChatParams holds the parameters for unbanChatSenderChat.
type UnbanChatSenderChatParams struct {
	ChatID       ChatID `json:"chat_id"`
	SenderChatID int64  `json:"sender_chat_id"`
}

// SetChatPermissions sets default chat permissions for all members.
// Returns True on success.
func (c *Client) SetChatPermissions(ctx context.Context, p *SetChatPermissionsParams) error {
	return c.Invoke(ctx, "setChatPermissions", p, nil)
}

// SetChatPermissionsParams holds the parameters for setChatPermissions.
type SetChatPermissionsParams struct {
	ChatID                        ChatID          `json:"chat_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions *bool           `json:"use_independent_chat_permissions,omitempty"`
}

// ExportChatInviteLink generates a new primary invite link for a chat; any
// previously generated primary link is revoked. Returns the new invite
// link as a String on success.
func (c *Client) ExportChatInviteLink(ctx context.Context, p *ExportChatInviteLinkParams) (string, error) {
	var s string
	if err := c.Invoke(ctx, "exportChatInviteLink", p, &s); err != nil {
		return "", err
	}
	return s, nil
}

// ExportChatInviteLinkParams holds the parameters for exportChatInviteLink.
type ExportChatInviteLinkParams struct {
	ChatID ChatID `json:"chat_id"`
}

// CreateChatInviteLink creates an additional invite link for a chat.
// Returns the new invite link as a ChatInviteLink object.
func (c *Client) CreateChatInviteLink(ctx context.Context, p *CreateChatInviteLinkParams) (*ChatInviteLink, error) {
	var m ChatInviteLink
	if err := c.Invoke(ctx, "createChatInviteLink", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// CreateChatInviteLinkParams holds the parameters for createChatInviteLink.
type CreateChatInviteLinkParams struct {
	ChatID             ChatID  `json:"chat_id"`
	Name               *string `json:"name,omitempty"`
	ExpireDate         *int64  `json:"expire_date,omitempty"`
	MemberLimit        *int    `json:"member_limit,omitempty"`
	CreatesJoinRequest *bool   `json:"creates_join_request,omitempty"`
}

// EditChatInviteLink edits a non-primary invite link created by the bot.
// Returns the edited invite link as a ChatInviteLink object.
func (c *Client) EditChatInviteLink(ctx context.Context, p *EditChatInviteLinkParams) (*ChatInviteLink, error) {
	var m ChatInviteLink
	if err := c.Invoke(ctx, "editChatInviteLink", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// EditChatInviteLinkParams holds the parameters for editChatInviteLink.
type EditChatInviteLinkParams struct {
	ChatID             ChatID  `json:"chat_id"`
	InviteLink         string  `json:"invite_link"`
	Name               *string `json:"name,omitempty"`
	ExpireDate         *int64  `json:"expire_date,omitempty"`
	MemberLimit        *int    `json:"member_limit,omitempty"`
	CreatesJoinRequest *bool   `json:"creates_join_request,omitempty"`
}

// CreateChatSubscriptionInviteLink creates a subscription invite link for
// a channel chat. Returns the new invite link as a ChatInviteLink object.
func (c *Client) CreateChatSubscriptionInviteLink(ctx context.Context, p *CreateChatSubscriptionInviteLinkParams) (*ChatInviteLink, error) {
	var m ChatInviteLink
	if err := c.Invoke(ctx, "createChatSubscriptionInviteLink", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// CreateChatSubscriptionInviteLinkParams holds the parameters for
// createChatSubscriptionInviteLink.
type CreateChatSubscriptionInviteLinkParams struct {
	ChatID             ChatID  `json:"chat_id"`
	Name               *string `json:"name,omitempty"`
	SubscriptionPeriod int     `json:"subscription_period"`
	SubscriptionPrice  int64   `json:"subscription_price"`
}

// EditChatSubscriptionInviteLink edits a subscription invite link created
// by the bot. Returns the edited invite link as a ChatInviteLink object.
func (c *Client) EditChatSubscriptionInviteLink(ctx context.Context, p *EditChatSubscriptionInviteLinkParams) (*ChatInviteLink, error) {
	var m ChatInviteLink
	if err := c.Invoke(ctx, "editChatSubscriptionInviteLink", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// EditChatSubscriptionInviteLinkParams holds the parameters for
// editChatSubscriptionInviteLink.
type EditChatSubscriptionInviteLinkParams struct {
	ChatID     ChatID  `json:"chat_id"`
	InviteLink string  `json:"invite_link"`
	Name       *string `json:"name,omitempty"`
}

// RevokeChatInviteLink revokes an invite link created by the bot. Returns
// the revoked invite link as a ChatInviteLink object.
func (c *Client) RevokeChatInviteLink(ctx context.Context, p *RevokeChatInviteLinkParams) (*ChatInviteLink, error) {
	var m ChatInviteLink
	if err := c.Invoke(ctx, "revokeChatInviteLink", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// RevokeChatInviteLinkParams holds the parameters for revokeChatInviteLink.
type RevokeChatInviteLinkParams struct {
	ChatID     ChatID `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

// ApproveChatJoinRequest approves a chat join request. Returns True on
// success.
func (c *Client) ApproveChatJoinRequest(ctx context.Context, p *ApproveChatJoinRequestParams) error {
	return c.Invoke(ctx, "approveChatJoinRequest", p, nil)
}

// ApproveChatJoinRequestParams holds the parameters for
// approveChatJoinRequest.
type ApproveChatJoinRequestParams struct {
	ChatID ChatID `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

// DeclineChatJoinRequest declines a chat join request. Returns True on
// success.
func (c *Client) DeclineChatJoinRequest(ctx context.Context, p *DeclineChatJoinRequestParams) error {
	return c.Invoke(ctx, "declineChatJoinRequest", p, nil)
}

// DeclineChatJoinRequestParams holds the parameters for
// declineChatJoinRequest.
type DeclineChatJoinRequestParams struct {
	ChatID ChatID `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

// AnswerChatJoinRequestQuery processes a received chat join request query.
// Returns True on success.
func (c *Client) AnswerChatJoinRequestQuery(ctx context.Context, p *AnswerChatJoinRequestQueryParams) error {
	return c.Invoke(ctx, "answerChatJoinRequestQuery", p, nil)
}

// AnswerChatJoinRequestQueryParams holds the parameters for
// answerChatJoinRequestQuery.
type AnswerChatJoinRequestQueryParams struct {
	ChatJoinRequestQueryID string `json:"chat_join_request_query_id"`
	Result                 string `json:"result"`
}

// SendChatJoinRequestWebApp processes a received chat join request query
// by showing a Mini App to the user before deciding the outcome. Returns
// True on success.
func (c *Client) SendChatJoinRequestWebApp(ctx context.Context, p *SendChatJoinRequestWebAppParams) error {
	return c.Invoke(ctx, "sendChatJoinRequestWebApp", p, nil)
}

// SendChatJoinRequestWebAppParams holds the parameters for
// sendChatJoinRequestWebApp.
type SendChatJoinRequestWebAppParams struct {
	ChatJoinRequestQueryID string `json:"chat_join_request_query_id"`
	WebAppURL              string `json:"web_app_url"`
}

// SetChatPhoto sets a new profile photo for the chat. Returns True on
// success.
func (c *Client) SetChatPhoto(ctx context.Context, p *SetChatPhotoParams) error {
	return c.Invoke(ctx, "setChatPhoto", p, nil)
}

// SetChatPhotoParams holds the parameters for setChatPhoto.
type SetChatPhotoParams struct {
	ChatID ChatID    `json:"chat_id"`
	Photo  InputFile `json:"photo"`
}

// DeleteChatPhoto deletes a chat photo. Returns True on success.
func (c *Client) DeleteChatPhoto(ctx context.Context, p *DeleteChatPhotoParams) error {
	return c.Invoke(ctx, "deleteChatPhoto", p, nil)
}

// DeleteChatPhotoParams holds the parameters for deleteChatPhoto.
type DeleteChatPhotoParams struct {
	ChatID ChatID `json:"chat_id"`
}

// SetChatTitle changes the title of a chat. Returns True on success.
func (c *Client) SetChatTitle(ctx context.Context, p *SetChatTitleParams) error {
	return c.Invoke(ctx, "setChatTitle", p, nil)
}

// SetChatTitleParams holds the parameters for setChatTitle.
type SetChatTitleParams struct {
	ChatID ChatID `json:"chat_id"`
	Title  string `json:"title"`
}

// SetChatDescription changes the description of a group, a supergroup or a
// channel. Returns True on success.
func (c *Client) SetChatDescription(ctx context.Context, p *SetChatDescriptionParams) error {
	return c.Invoke(ctx, "setChatDescription", p, nil)
}

// SetChatDescriptionParams holds the parameters for setChatDescription.
type SetChatDescriptionParams struct {
	ChatID      ChatID  `json:"chat_id"`
	Description *string `json:"description,omitempty"`
}

// PinChatMessage adds a message to the list of pinned messages in a chat.
// Returns True on success.
func (c *Client) PinChatMessage(ctx context.Context, p *PinChatMessageParams) error {
	return c.Invoke(ctx, "pinChatMessage", p, nil)
}

// PinChatMessageParams holds the parameters for pinChatMessage.
type PinChatMessageParams struct {
	BusinessConnectionID *string `json:"business_connection_id,omitempty"`
	ChatID               ChatID  `json:"chat_id"`
	MessageID            int64   `json:"message_id"`
	DisableNotification  *bool   `json:"disable_notification,omitempty"`
}

// UnpinChatMessage removes a message from the list of pinned messages in a
// chat. Returns True on success.
func (c *Client) UnpinChatMessage(ctx context.Context, p *UnpinChatMessageParams) error {
	return c.Invoke(ctx, "unpinChatMessage", p, nil)
}

// UnpinChatMessageParams holds the parameters for unpinChatMessage.
type UnpinChatMessageParams struct {
	BusinessConnectionID *string `json:"business_connection_id,omitempty"`
	ChatID               ChatID  `json:"chat_id"`
	MessageID            *int64  `json:"message_id,omitempty"`
}

// UnpinAllChatMessages clears the list of pinned messages in a chat.
// Returns True on success.
func (c *Client) UnpinAllChatMessages(ctx context.Context, p *UnpinAllChatMessagesParams) error {
	return c.Invoke(ctx, "unpinAllChatMessages", p, nil)
}

// UnpinAllChatMessagesParams holds the parameters for unpinAllChatMessages.
type UnpinAllChatMessagesParams struct {
	ChatID ChatID `json:"chat_id"`
}

// LeaveChat makes the bot leave a group, supergroup or channel. Returns
// True on success.
func (c *Client) LeaveChat(ctx context.Context, p *LeaveChatParams) error {
	return c.Invoke(ctx, "leaveChat", p, nil)
}

// LeaveChatParams holds the parameters for leaveChat.
type LeaveChatParams struct {
	ChatID ChatID `json:"chat_id"`
}

// GetChat gets up-to-date information about the chat. Returns a
// ChatFullInfo object on success.
func (c *Client) GetChat(ctx context.Context, p *GetChatParams) (*ChatFullInfo, error) {
	var m ChatFullInfo
	if err := c.Invoke(ctx, "getChat", p, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// GetChatParams holds the parameters for getChat.
type GetChatParams struct {
	ChatID ChatID `json:"chat_id"`
}

// GetChatAdministrators gets a list of administrators in a chat. Returns
// an Array of ChatMember objects.
func (c *Client) GetChatAdministrators(ctx context.Context, p *GetChatAdministratorsParams) ([]ChatMember, error) {
	var ms []ChatMember
	if err := c.Invoke(ctx, "getChatAdministrators", p, &ms); err != nil {
		return nil, err
	}
	return ms, nil
}

// GetChatAdministratorsParams holds the parameters for
// getChatAdministrators.
type GetChatAdministratorsParams struct {
	ChatID     ChatID `json:"chat_id"`
	ReturnBots *bool  `json:"return_bots,omitempty"`
}

// GetChatMemberCount gets the number of members in a chat. Returns Int on
// success.
func (c *Client) GetChatMemberCount(ctx context.Context, p *GetChatMemberCountParams) (int, error) {
	var n int
	if err := c.Invoke(ctx, "getChatMemberCount", p, &n); err != nil {
		return 0, err
	}
	return n, nil
}

// GetChatMemberCountParams holds the parameters for getChatMemberCount.
type GetChatMemberCountParams struct {
	ChatID ChatID `json:"chat_id"`
}

// GetChatMember gets information about a member of a chat. Returns a
// ChatMember object on success.
func (c *Client) GetChatMember(ctx context.Context, p *GetChatMemberParams) (ChatMember, error) {
	var m ChatMember
	if err := c.Invoke(ctx, "getChatMember", p, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// GetChatMemberParams holds the parameters for getChatMember.
type GetChatMemberParams struct {
	ChatID ChatID `json:"chat_id"`
	UserID int64  `json:"user_id"`
}

// GetUserPersonalChatMessages gets the last messages from the personal
// chat of a given user. Returns an Array of Message objects on success.
func (c *Client) GetUserPersonalChatMessages(ctx context.Context, p *GetUserPersonalChatMessagesParams) ([]Message, error) {
	var ms []Message
	if err := c.Invoke(ctx, "getUserPersonalChatMessages", p, &ms); err != nil {
		return nil, err
	}
	return ms, nil
}

// GetUserPersonalChatMessagesParams holds the parameters for
// getUserPersonalChatMessages.
type GetUserPersonalChatMessagesParams struct {
	UserID int64 `json:"user_id"`
	Limit  int   `json:"limit"`
}

// SetChatStickerSet sets a new group sticker set for a supergroup. Returns
// True on success.
func (c *Client) SetChatStickerSet(ctx context.Context, p *SetChatStickerSetParams) error {
	return c.Invoke(ctx, "setChatStickerSet", p, nil)
}

// SetChatStickerSetParams holds the parameters for setChatStickerSet.
type SetChatStickerSetParams struct {
	ChatID         ChatID `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

// DeleteChatStickerSet deletes a group sticker set from a supergroup.
// Returns True on success.
func (c *Client) DeleteChatStickerSet(ctx context.Context, p *DeleteChatStickerSetParams) error {
	return c.Invoke(ctx, "deleteChatStickerSet", p, nil)
}

// DeleteChatStickerSetParams holds the parameters for deleteChatStickerSet.
type DeleteChatStickerSetParams struct {
	ChatID ChatID `json:"chat_id"`
}

// GetForumTopicIconStickers returns custom emoji stickers that can be used as forum topic icons.
func (c *Client) GetForumTopicIconStickers(ctx context.Context) ([]Sticker, error) {
	var stickers []Sticker
	if err := c.Invoke(ctx, "getForumTopicIconStickers", nil, &stickers); err != nil {
		return nil, err
	}
	return stickers, nil
}

// CreateForumTopicParams holds the parameters for the createForumTopic method.
type CreateForumTopicParams struct {
	ChatID            ChatID  `json:"chat_id"`
	Name              string  `json:"name"`
	IconColor         *int    `json:"icon_color,omitempty"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

// CreateForumTopic creates a topic in a forum supergroup chat or a private chat with a user.
func (c *Client) CreateForumTopic(ctx context.Context, p *CreateForumTopicParams) (*ForumTopic, error) {
	var t ForumTopic
	if err := c.Invoke(ctx, "createForumTopic", p, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

// EditForumTopicParams holds the parameters for the editForumTopic method.
type EditForumTopicParams struct {
	ChatID            ChatID  `json:"chat_id"`
	MessageThreadID   int64   `json:"message_thread_id"`
	Name              *string `json:"name,omitempty"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

// EditForumTopic edits the name and icon of a topic in a forum supergroup chat.
func (c *Client) EditForumTopic(ctx context.Context, p *EditForumTopicParams) error {
	return c.Invoke(ctx, "editForumTopic", p, nil)
}

// CloseForumTopicParams holds the parameters for the closeForumTopic method.
type CloseForumTopicParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int64  `json:"message_thread_id"`
}

// CloseForumTopic closes an open topic in a forum supergroup chat.
func (c *Client) CloseForumTopic(ctx context.Context, p *CloseForumTopicParams) error {
	return c.Invoke(ctx, "closeForumTopic", p, nil)
}

// ReopenForumTopicParams holds the parameters for the reopenForumTopic method.
type ReopenForumTopicParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int64  `json:"message_thread_id"`
}

// ReopenForumTopic reopens a closed topic in a forum supergroup chat.
func (c *Client) ReopenForumTopic(ctx context.Context, p *ReopenForumTopicParams) error {
	return c.Invoke(ctx, "reopenForumTopic", p, nil)
}

// DeleteForumTopicParams holds the parameters for the deleteForumTopic method.
type DeleteForumTopicParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int64  `json:"message_thread_id"`
}

// DeleteForumTopic deletes a forum topic along with all its messages.
func (c *Client) DeleteForumTopic(ctx context.Context, p *DeleteForumTopicParams) error {
	return c.Invoke(ctx, "deleteForumTopic", p, nil)
}

// UnpinAllForumTopicMessagesParams holds the parameters for the unpinAllForumTopicMessages method.
type UnpinAllForumTopicMessagesParams struct {
	ChatID          ChatID `json:"chat_id"`
	MessageThreadID int64  `json:"message_thread_id"`
}

// UnpinAllForumTopicMessages clears the list of pinned messages in a forum topic.
func (c *Client) UnpinAllForumTopicMessages(ctx context.Context, p *UnpinAllForumTopicMessagesParams) error {
	return c.Invoke(ctx, "unpinAllForumTopicMessages", p, nil)
}

// EditGeneralForumTopicParams holds the parameters for the editGeneralForumTopic method.
type EditGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
	Name   string `json:"name"`
}

// EditGeneralForumTopic edits the name of the 'General' topic in a forum supergroup chat.
func (c *Client) EditGeneralForumTopic(ctx context.Context, p *EditGeneralForumTopicParams) error {
	return c.Invoke(ctx, "editGeneralForumTopic", p, nil)
}

// CloseGeneralForumTopicParams holds the parameters for the closeGeneralForumTopic method.
type CloseGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

// CloseGeneralForumTopic closes the open 'General' topic in a forum supergroup chat.
func (c *Client) CloseGeneralForumTopic(ctx context.Context, p *CloseGeneralForumTopicParams) error {
	return c.Invoke(ctx, "closeGeneralForumTopic", p, nil)
}

// ReopenGeneralForumTopicParams holds the parameters for the reopenGeneralForumTopic method.
type ReopenGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

// ReopenGeneralForumTopic reopens the closed 'General' topic in a forum supergroup chat.
func (c *Client) ReopenGeneralForumTopic(ctx context.Context, p *ReopenGeneralForumTopicParams) error {
	return c.Invoke(ctx, "reopenGeneralForumTopic", p, nil)
}

// HideGeneralForumTopicParams holds the parameters for the hideGeneralForumTopic method.
type HideGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

// HideGeneralForumTopic hides the 'General' topic in a forum supergroup chat.
func (c *Client) HideGeneralForumTopic(ctx context.Context, p *HideGeneralForumTopicParams) error {
	return c.Invoke(ctx, "hideGeneralForumTopic", p, nil)
}

// UnhideGeneralForumTopicParams holds the parameters for the unhideGeneralForumTopic method.
type UnhideGeneralForumTopicParams struct {
	ChatID ChatID `json:"chat_id"`
}

// UnhideGeneralForumTopic unhides the 'General' topic in a forum supergroup chat.
func (c *Client) UnhideGeneralForumTopic(ctx context.Context, p *UnhideGeneralForumTopicParams) error {
	return c.Invoke(ctx, "unhideGeneralForumTopic", p, nil)
}

// UnpinAllGeneralForumTopicMessagesParams holds the parameters for the unpinAllGeneralForumTopicMessages method.
type UnpinAllGeneralForumTopicMessagesParams struct {
	ChatID ChatID `json:"chat_id"`
}

// UnpinAllGeneralForumTopicMessages clears the list of pinned messages in the 'General' forum topic.
func (c *Client) UnpinAllGeneralForumTopicMessages(ctx context.Context, p *UnpinAllGeneralForumTopicMessagesParams) error {
	return c.Invoke(ctx, "unpinAllGeneralForumTopicMessages", p, nil)
}
