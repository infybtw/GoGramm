package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Update represents an incoming update. At most one of the optional fields can
// be present in any given update.
type Update struct {
	UpdateID                 int64                        `json:"update_id"`
	Message                  *Message                     `json:"message"`
	EditedMessage            *Message                     `json:"edited_message"`
	ChannelPost              *Message                     `json:"channel_post"`
	EditedChannelPost        *Message                     `json:"edited_channel_post"`
	BusinessConnection       *BusinessConnection          `json:"business_connection"`
	BusinessMessage          *Message                     `json:"business_message"`
	EditedBusinessMessage    *Message                     `json:"edited_business_message"`
	DeletedBusinessMessages  *BusinessMessagesDeleted     `json:"deleted_business_messages"`
	GuestMessage             *Message                     `json:"guest_message"`
	MessageReaction          *MessageReactionUpdated      `json:"message_reaction"`
	MessageReactionCount     *MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQuery              *InlineQuery                 `json:"inline_query"`
	ChosenInlineResult       *ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQuery            *CallbackQuery               `json:"callback_query"`
	ShippingQuery            *ShippingQuery               `json:"shipping_query"`
	PreCheckoutQuery         *PreCheckoutQuery            `json:"pre_checkout_query"`
	PurchasedPaidMedia       *PaidMediaPurchased          `json:"purchased_paid_media"`
	Poll                     *Poll                        `json:"poll"`
	PollAnswer               *PollAnswer                  `json:"poll_answer"`
	MyChatMember             *ChatMemberUpdated           `json:"my_chat_member"`
	ChatMember               *ChatMemberUpdated           `json:"chat_member"`
	ChatJoinRequest          *ChatJoinRequest             `json:"chat_join_request"`
	ChatBoost                *ChatBoostUpdated            `json:"chat_boost"`
	RemovedChatBoost         *ChatBoostRemoved            `json:"removed_chat_boost"`
	ManagedBot               *ManagedBotUpdated           `json:"managed_bot"`
	Subscription             *BotSubscriptionUpdated      `json:"subscription"`
	StoppedMessageGeneration *MessageGenerationStopped    `json:"stopped_message_generation"`
}

// WebhookInfo describes the current status of a webhook.
type WebhookInfo struct {
	URL                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int      `json:"pending_update_count"`
	IPAddress                    *string  `json:"ip_address"`
	LastErrorDate                *int64   `json:"last_error_date"`
	LastErrorMessage             *string  `json:"last_error_message"`
	LastSynchronizationErrorDate *int64   `json:"last_synchronization_error_date"`
	MaxConnections               *int     `json:"max_connections"`
	AllowedUpdates               []string `json:"allowed_updates"`
}

// User represents a Telegram user or bot.
type User struct {
	ID                         int64   `json:"id"`
	IsBot                      bool    `json:"is_bot"`
	FirstName                  string  `json:"first_name"`
	LastName                   *string `json:"last_name"`
	Username                   *string `json:"username"`
	LanguageCode               *string `json:"language_code"`
	IsPremium                  *bool   `json:"is_premium"`
	AddedToAttachmentMenu      *bool   `json:"added_to_attachment_menu"`
	CanJoinGroups              *bool   `json:"can_join_groups"`
	CanReadAllGroupMessages    *bool   `json:"can_read_all_group_messages"`
	SupportsGuestQueries       *bool   `json:"supports_guest_queries"`
	SupportsInlineQueries      *bool   `json:"supports_inline_queries"`
	CanConnectToBusiness       *bool   `json:"can_connect_to_business"`
	HasMainWebApp              *bool   `json:"has_main_web_app"`
	HasTopicsEnabled           *bool   `json:"has_topics_enabled"`
	AllowsUsersToCreateTopics  *bool   `json:"allows_users_to_create_topics"`
	CanManageBots              *bool   `json:"can_manage_bots"`
	SupportsJoinRequestQueries *bool   `json:"supports_join_request_queries"`
}

// Chat represents a chat.
type Chat struct {
	ID               int64   `json:"id"`
	Type             string  `json:"type"`
	Title            *string `json:"title"`
	Username         *string `json:"username"`
	FirstName        *string `json:"first_name"`
	LastName         *string `json:"last_name"`
	IsForum          *bool   `json:"is_forum"`
	IsDirectMessages *bool   `json:"is_direct_messages"`
}

// ChatFullInfo contains full information about a chat.
type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              *string               `json:"title"`
	Username                           *string               `json:"username"`
	FirstName                          *string               `json:"first_name"`
	LastName                           *string               `json:"last_name"`
	IsForum                            *bool                 `json:"is_forum"`
	IsDirectMessages                   *bool                 `json:"is_direct_messages"`
	AccentColorID                      int64                 `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo"`
	ActiveUsernames                    *[]string             `json:"active_usernames"`
	Birthdate                          *Birthdate            `json:"birthdate"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro"`
	BusinessLocation                   *BusinessLocation     `json:"business_location"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours"`
	PersonalChat                       *Chat                 `json:"personal_chat"`
	ParentChat                         *Chat                 `json:"parent_chat"`
	AvailableReactions                 *[]ReactionType       `json:"available_reactions"`
	BackgroundCustomEmojiID            *string               `json:"background_custom_emoji_id"`
	ProfileAccentColorID               *int64                `json:"profile_accent_color_id"`
	ProfileBackgroundCustomEmojiID     *string               `json:"profile_background_custom_emoji_id"`
	EmojiStatusCustomEmojiID           *string               `json:"emoji_status_custom_emoji_id"`
	EmojiStatusExpirationDate          *int64                `json:"emoji_status_expiration_date"`
	Bio                                *string               `json:"bio"`
	HasPrivateForwards                 *bool                 `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages"`
	JoinByRequest                      *bool                 `json:"join_by_request"`
	Description                        *string               `json:"description"`
	InviteLink                         *string               `json:"invite_link"`
	PinnedMessage                      *Message              `json:"pinned_message"`
	Permissions                        *ChatPermissions      `json:"permissions"`
	AcceptedGiftTypes                  AcceptedGiftTypes     `json:"accepted_gift_types"`
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media"`
	SlowModeDelay                      *int                  `json:"slow_mode_delay"`
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count"`
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time"`
	HasAggressiveAntiSpamEnabled       *bool                 `json:"has_aggressive_anti_spam_enabled"`
	HasHiddenMembers                   *bool                 `json:"has_hidden_members"`
	HasProtectedContent                *bool                 `json:"has_protected_content"`
	HasVisibleHistory                  *bool                 `json:"has_visible_history"`
	StickerSetName                     *string               `json:"sticker_set_name"`
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set"`
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name"`
	LinkedChatID                       *int64                `json:"linked_chat_id"`
	Location                           *ChatLocation         `json:"location"`
	Rating                             *UserRating           `json:"rating"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors"`
	PaidMessageStarCount               *int64                `json:"paid_message_star_count"`
	GuardBot                           *User                 `json:"guard_bot"`
	Community                          *Community            `json:"community"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// Community represents a community (a group of chats).
type Community struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Birthdate describes the birthdate of a user.
type Birthdate struct {
	Day   int  `json:"day"`
	Month int  `json:"month"`
	Year  *int `json:"year"`
}

// BusinessIntro contains information about the start page settings of a Telegram Business account.
type BusinessIntro struct {
	Title   *string  `json:"title"`
	Message *string  `json:"message"`
	Sticker *Sticker `json:"sticker"`
}

// BusinessLocation contains information about the location of a Telegram Business account.
type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location"`
}

// BusinessOpeningHoursInterval describes an interval of time during which a business is open.
type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

// BusinessOpeningHours describes the opening hours of a business.
type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// UserRating describes the rating of a user based on their Telegram Star spendings.
type UserRating struct {
	Level              int  `json:"level"`
	Rating             int  `json:"rating"`
	CurrentLevelRating int  `json:"current_level_rating"`
	NextLevelRating    *int `json:"next_level_rating"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// Location represents a point on the map.
type Location struct {
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy"`
	LivePeriod           *int     `json:"live_period"`
	Heading              *int     `json:"heading"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius"`
}

// LocationAddress describes the physical address of a location.
type LocationAddress struct {
	CountryCode string  `json:"country_code"`
	State       *string `json:"state"`
	City        *string `json:"city"`
	Street      *string `json:"street"`
}

// Venue represents a venue.
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    *string  `json:"foursquare_id"`
	FoursquareType  *string  `json:"foursquare_type"`
	GooglePlaceID   *string  `json:"google_place_id"`
	GooglePlaceType *string  `json:"google_place_type"`
}

// DirectMessagesTopic describes a topic of a direct messages chat.
type DirectMessagesTopic struct {
	TopicID int64 `json:"topic_id"`
	User    *User `json:"user"`
}

// SharedUser contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
type SharedUser struct {
	UserID    int64        `json:"user_id"`
	FirstName *string      `json:"first_name"`
	LastName  *string      `json:"last_name"`
	Username  *string      `json:"username"`
	Photo     *[]PhotoSize `json:"photo"`
}

// UsersShared contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	RequestID int64        `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

// ChatShared contains information about a chat that was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	RequestID int64        `json:"request_id"`
	ChatID    int64        `json:"chat_id"`
	Title     *string      `json:"title"`
	Username  *string      `json:"username"`
	Photo     *[]PhotoSize `json:"photo"`
}

// WriteAccessAllowed represents a service message about a user allowing a bot to write messages after adding it
// to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent
// by the method requestWriteAccess.
type WriteAccessAllowed struct {
	FromRequest        *bool   `json:"from_request"`
	WebAppName         *string `json:"web_app_name"`
	FromAttachmentMenu *bool   `json:"from_attachment_menu"`
}

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat triggers
// a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// ManagedBotCreated contains information about the bot that was created to be managed by the current bot.
type ManagedBotCreated struct {
	Bot User `json:"bot"`
}

// ManagedBotUpdated contains information about the creation, token update, or owner update of a bot that is
// managed by the current bot.
type ManagedBotUpdated struct {
	User User `json:"user"`
	Bot  User `json:"bot"`
}

// BotSubscriptionUpdated contains information about changes to a user payment subscription toward the current bot.
type BotSubscriptionUpdated struct {
	User           User   `json:"user"`
	InvoicePayload string `json:"invoice_payload"`
	State          string `json:"state"`
}

// MessageGenerationStopped describes an update about a user stopping message generation.
type MessageGenerationStopped struct {
	Chat            Chat   `json:"chat"`
	MessageThreadID *int64 `json:"message_thread_id"`
	DraftID         int64  `json:"draft_id"`
}

// ChatOwnerLeft describes a service message about the chat owner leaving the chat.
type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner"`
}

// ChatOwnerChanged describes a service message about an ownership change in the chat.
type ChatOwnerChanged struct {
	NewOwner User `json:"new_owner"`
}

// Message represents a message.
type Message struct {
	MessageID                     int64                          `json:"message_id"`
	MessageThreadID               *int64                         `json:"message_thread_id"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic"`
	From                          *User                          `json:"from"`
	SenderChat                    *Chat                          `json:"sender_chat"`
	SenderBoostCount              *int                           `json:"sender_boost_count"`
	SenderBusinessBot             *User                          `json:"sender_business_bot"`
	SenderTag                     *string                        `json:"sender_tag"`
	ReceiverUser                  *User                          `json:"receiver_user"`
	EphemeralMessageID            *int64                         `json:"ephemeral_message_id"`
	Date                          int64                          `json:"date"`
	GuestQueryID                  *string                        `json:"guest_query_id"`
	BusinessConnectionID          *string                        `json:"business_connection_id"`
	Chat                          Chat                           `json:"chat"`
	ForwardOrigin                 MessageOrigin                  `json:"forward_origin"`
	IsTopicMessage                *bool                          `json:"is_topic_message"`
	IsAutomaticForward            *bool                          `json:"is_automatic_forward"`
	ReplyToMessage                *Message                       `json:"reply_to_message"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply"`
	Quote                         *TextQuote                     `json:"quote"`
	ReplyToStory                  *Story                         `json:"reply_to_story"`
	ReplyToChecklistTaskID        *int64                         `json:"reply_to_checklist_task_id"`
	ReplyToPollOptionID           *string                        `json:"reply_to_poll_option_id"`
	ViaBot                        *User                          `json:"via_bot"`
	GuestBotCallerUser            *User                          `json:"guest_bot_caller_user"`
	GuestBotCallerChat            *Chat                          `json:"guest_bot_caller_chat"`
	EditDate                      *int64                         `json:"edit_date"`
	HasProtectedContent           *bool                          `json:"has_protected_content"`
	IsFromOffline                 *bool                          `json:"is_from_offline"`
	IsPaidPost                    *bool                          `json:"is_paid_post"`
	MediaGroupID                  *string                        `json:"media_group_id"`
	AuthorSignature               *string                        `json:"author_signature"`
	PaidStarCount                 *int64                         `json:"paid_star_count"`
	Text                          *string                        `json:"text"`
	Entities                      *[]MessageEntity               `json:"entities"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info"`
	EffectID                      *string                        `json:"effect_id"`
	RichMessage                   *RichMessage                   `json:"rich_message"`
	Animation                     *Animation                     `json:"animation"`
	Audio                         *Audio                         `json:"audio"`
	Document                      *Document                      `json:"document"`
	LivePhoto                     *LivePhoto                     `json:"live_photo"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media"`
	Photo                         *[]PhotoSize                   `json:"photo"`
	Sticker                       *Sticker                       `json:"sticker"`
	Story                         *Story                         `json:"story"`
	Video                         *Video                         `json:"video"`
	VideoNote                     *VideoNote                     `json:"video_note"`
	Voice                         *Voice                         `json:"voice"`
	Caption                       *string                        `json:"caption"`
	CaptionEntities               *[]MessageEntity               `json:"caption_entities"`
	ShowCaptionAboveMedia         *bool                          `json:"show_caption_above_media"`
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler"`
	Checklist                     *Checklist                     `json:"checklist"`
	Contact                       *Contact                       `json:"contact"`
	Dice                          *Dice                          `json:"dice"`
	Game                          *Game                          `json:"game"`
	Poll                          *Poll                          `json:"poll"`
	Venue                         *Venue                         `json:"venue"`
	Location                      *Location                      `json:"location"`
	NewChatMembers                *[]User                        `json:"new_chat_members"`
	LeftChatMember                *User                          `json:"left_chat_member"`
	ChatOwnerLeft                 *ChatOwnerLeft                 `json:"chat_owner_left"`
	ChatOwnerChanged              *ChatOwnerChanged              `json:"chat_owner_changed"`
	NewChatTitle                  *string                        `json:"new_chat_title"`
	NewChatPhoto                  *[]PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo"`
	GroupChatCreated              *bool                          `json:"group_chat_created"`
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               *int64                         `json:"migrate_to_chat_id"`
	MigrateFromChatID             *int64                         `json:"migrate_from_chat_id"`
	PinnedMessage                 MaybeInaccessibleMessage       `json:"pinned_message"`
	Invoice                       *Invoice                       `json:"invoice"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment"`
	UsersShared                   *UsersShared                   `json:"users_shared"`
	ChatShared                    *ChatShared                    `json:"chat_shared"`
	Gift                          *GiftInfo                      `json:"gift"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift"`
	GiftUpgradeSent               *GiftInfo                      `json:"gift_upgrade_sent"`
	ConnectedWebsite              *string                        `json:"connected_website"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  *PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set"`
	ChecklistTasksDone            *ChecklistTasksDone            `json:"checklist_tasks_done"`
	ChecklistTasksAdded           *ChecklistTasksAdded           `json:"checklist_tasks_added"`
	CommunityChatAdded            *CommunityChatAdded            `json:"community_chat_added"`
	CommunityChatJoined           *CommunityChatJoined           `json:"community_chat_joined"`
	CommunityChatRemoved          *CommunityChatRemoved          `json:"community_chat_removed"`
	DirectMessagePriceChanged     *DirectMessagePriceChanged     `json:"direct_message_price_changed"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created"`
	Giveaway                      *Giveaway                      `json:"giveaway"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed"`
	ManagedBotCreated             *ManagedBotCreated             `json:"managed_bot_created"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed"`
	PollOptionAdded               *PollOptionAdded               `json:"poll_option_added"`
	PollOptionDeleted             *PollOptionDeleted             `json:"poll_option_deleted"`
	SuggestedPostApproved         *SuggestedPostApproved         `json:"suggested_post_approved"`
	SuggestedPostApprovalFailed   *SuggestedPostApprovalFailed   `json:"suggested_post_approval_failed"`
	SuggestedPostDeclined         *SuggestedPostDeclined         `json:"suggested_post_declined"`
	SuggestedPostPaid             *SuggestedPostPaid             `json:"suggested_post_paid"`
	SuggestedPostRefunded         *SuggestedPostRefunded         `json:"suggested_post_refunded"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    *WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
}

// MessageId represents a unique message identifier.
type MessageId struct {
	MessageID int64 `json:"message_id"`
}

// InaccessibleMessage describes a message that was deleted or is otherwise inaccessible.
type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`
	MessageID int64 `json:"message_id"`
	Date      int64 `json:"date"`
}

// MaybeInaccessibleMessage is either a Message or an InaccessibleMessage.
type MaybeInaccessibleMessage interface {
	maybeInaccessibleMessage()
}

func (*Message) maybeInaccessibleMessage()             {}
func (*InaccessibleMessage) maybeInaccessibleMessage() {}

// MessageEntity represents one special entity in a text message.
type MessageEntity struct {
	Type           string  `json:"type"`
	Offset         int     `json:"offset"`
	Length         int     `json:"length"`
	URL            *string `json:"url"`
	User           *User   `json:"user"`
	Language       *string `json:"language"`
	CustomEmojiID  *string `json:"custom_emoji_id"`
	UnixTime       *int64  `json:"unix_time"`
	DateTimeFormat *string `json:"date_time_format"`
}

// TextQuote contains information about the quoted part of a message.
type TextQuote struct {
	Text     string           `json:"text"`
	Entities *[]MessageEntity `json:"entities"`
	Position int              `json:"position"`
	IsManual *bool            `json:"is_manual"`
}

// ExternalReplyInfo contains information about a message being replied to.
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`
	Chat               *Chat               `json:"chat"`
	MessageID          *int64              `json:"message_id"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`
	Animation          *Animation          `json:"animation"`
	Audio              *Audio              `json:"audio"`
	Document           *Document           `json:"document"`
	LivePhoto          *LivePhoto          `json:"live_photo"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media"`
	Photo              *[]PhotoSize        `json:"photo"`
	Sticker            *Sticker            `json:"sticker"`
	Story              *Story              `json:"story"`
	Video              *Video              `json:"video"`
	VideoNote          *VideoNote          `json:"video_note"`
	Voice              *Voice              `json:"voice"`
	HasMediaSpoiler    *bool               `json:"has_media_spoiler"`
	Checklist          *Checklist          `json:"checklist"`
	Contact            *Contact            `json:"contact"`
	Dice               *Dice               `json:"dice"`
	Game               *Game               `json:"game"`
	Giveaway           *Giveaway           `json:"giveaway"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners"`
	Invoice            *Invoice            `json:"invoice"`
	Location           *Location           `json:"location"`
	Poll               *Poll               `json:"poll"`
	Venue              *Venue              `json:"venue"`
}

// ReplyParameters describes reply parameters for a sent message.
type ReplyParameters struct {
	MessageID                *int64           `json:"message_id"`
	ChatID                   *ChatID          `json:"chat_id"`
	EphemeralMessageID       *int64           `json:"ephemeral_message_id"`
	AllowSendingWithoutReply *bool            `json:"allow_sending_without_reply"`
	Quote                    *string          `json:"quote"`
	QuoteParseMode           *string          `json:"quote_parse_mode"`
	QuoteEntities            *[]MessageEntity `json:"quote_entities"`
	QuotePosition            *int             `json:"quote_position"`
	ChecklistTaskID          *int64           `json:"checklist_task_id"`
	PollOptionID             *string          `json:"poll_option_id"`
}

// EphemeralMessageParameters describes parameters for an ephemeral message.
type EphemeralMessageParameters struct {
	ReceiverUserID              int64   `json:"receiver_user_id"`
	CallbackQueryID             *string `json:"callback_query_id"`
	ReplaceCallbackQueryMessage *bool   `json:"replace_callback_query_message"`
}

// MessageOrigin describes the origin of a forwarded message.
type MessageOrigin interface {
	messageOrigin()
}

// MessageOriginUser indicates a message originally sent by a known user.
type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       int64  `json:"date"`
	SenderUser User   `json:"sender_user"`
}

func (*MessageOriginUser) messageOrigin() {}

func (m MessageOriginUser) MarshalJSON() ([]byte, error) {
	type alias MessageOriginUser
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "user", alias: alias(m)})
}

// MessageOriginHiddenUser indicates a message originally sent by an unknown user.
type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           int64  `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

func (*MessageOriginHiddenUser) messageOrigin() {}

func (m MessageOriginHiddenUser) MarshalJSON() ([]byte, error) {
	type alias MessageOriginHiddenUser
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "hidden_user", alias: alias(m)})
}

// MessageOriginChat indicates a message originally sent on behalf of a chat.
type MessageOriginChat struct {
	Type            string  `json:"type"`
	Date            int64   `json:"date"`
	SenderChat      Chat    `json:"sender_chat"`
	AuthorSignature *string `json:"author_signature"`
}

func (*MessageOriginChat) messageOrigin() {}

func (m MessageOriginChat) MarshalJSON() ([]byte, error) {
	type alias MessageOriginChat
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "chat", alias: alias(m)})
}

// MessageOriginChannel indicates a message originally sent to a channel.
type MessageOriginChannel struct {
	Type            string  `json:"type"`
	Date            int64   `json:"date"`
	Chat            Chat    `json:"chat"`
	MessageID       int64   `json:"message_id"`
	AuthorSignature *string `json:"author_signature"`
}

func (*MessageOriginChannel) messageOrigin() {}

func (m MessageOriginChannel) MarshalJSON() ([]byte, error) {
	type alias MessageOriginChannel
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "channel", alias: alias(m)})
}

// LinkPreviewOptions describes options used for link preview generation.
type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled"`
	URL              *string `json:"url"`
	PreferSmallMedia *bool   `json:"prefer_small_media"`
	PreferLargeMedia *bool   `json:"prefer_large_media"`
	ShowAboveText    *bool   `json:"show_above_text"`
}

// WebAppData describes data sent from a Web App.
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// ChecklistTask describes a task in a checklist.
type ChecklistTask struct {
	ID              int64            `json:"id"`
	Text            string           `json:"text"`
	TextEntities    *[]MessageEntity `json:"text_entities"`
	CompletedByUser *User            `json:"completed_by_user"`
	CompletedByChat *Chat            `json:"completed_by_chat"`
	CompletionDate  *int64           `json:"completion_date"`
}

// Checklist describes a checklist.
type Checklist struct {
	Title                    string           `json:"title"`
	TitleEntities            *[]MessageEntity `json:"title_entities"`
	Tasks                    []ChecklistTask  `json:"tasks"`
	OthersCanAddTasks        *bool            `json:"others_can_add_tasks"`
	OthersCanMarkTasksAsDone *bool            `json:"others_can_mark_tasks_as_done"`
}

// InputChecklistTask describes a task to add to a checklist.
type InputChecklistTask struct {
	ID           int64            `json:"id"`
	Text         string           `json:"text"`
	ParseMode    *string          `json:"parse_mode"`
	TextEntities *[]MessageEntity `json:"text_entities"`
}

// InputChecklist describes a checklist to create.
type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                *string              `json:"parse_mode"`
	TitleEntities            *[]MessageEntity     `json:"title_entities"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        *bool                `json:"others_can_add_tasks"`
	OthersCanMarkTasksAsDone *bool                `json:"others_can_mark_tasks_as_done"`
}

// ChecklistTasksDone describes checklist tasks marked as done or not done.
type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message"`
	MarkedAsDoneTaskIDs    *[]int64 `json:"marked_as_done_task_ids"`
	MarkedAsNotDoneTaskIDs *[]int64 `json:"marked_as_not_done_task_ids"`
}

// ChecklistTasksAdded describes tasks added to a checklist.
type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message"`
	Tasks            []ChecklistTask `json:"tasks"`
}

// ForumTopicCreated describes a newly created forum topic.
type ForumTopicCreated struct {
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id"`
	IsNameImplicit    *bool   `json:"is_name_implicit"`
}

type ForumTopicClosed struct{}

type ForumTopicEdited struct {
	Name              *string `json:"name"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id"`
}

type ForumTopicReopened struct{}

type GeneralForumTopicHidden struct{}

type GeneralForumTopicUnhidden struct{}

// VideoChatScheduled describes a scheduled video chat.
type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

type VideoChatStarted struct{}

// VideoChatEnded describes an ended video chat.
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited describes participants invited to a video chat.
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// PaidMessagePriceChanged describes a change in the price of paid messages.
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int64 `json:"paid_message_star_count"`
}

// DirectMessagePriceChanged describes a change in the price of direct messages.
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool   `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   *int64 `json:"direct_message_star_count"`
}

type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message"`
	Price                *SuggestedPostPrice `json:"price"`
	SendDate             int64               `json:"send_date"`
}

type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message           `json:"suggested_post_message"`
	Price                SuggestedPostPrice `json:"price"`
}

type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message"`
	Comment              *string  `json:"comment"`
}

type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message"`
	Currency             string      `json:"currency"`
	Amount               *int64      `json:"amount"`
	StarAmount           *StarAmount `json:"star_amount"`
}

type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message"`
	Reason               string   `json:"reason"`
}

// SuggestedPostPrice describes the price of a suggested post.
type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

// SuggestedPostInfo contains information about a suggested post.
type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price"`
	SendDate *int64              `json:"send_date"`
}

// SuggestedPostParameters contains parameters for a suggested post.
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price"`
	SendDate *int64              `json:"send_date"`
}

// GiveawayCreated describes creation of a scheduled giveaway.
type GiveawayCreated struct {
	PrizeStarCount *int64 `json:"prize_star_count"`
}

// Giveaway describes a scheduled giveaway.
type Giveaway struct {
	Chats                         []Chat    `json:"chats"`
	WinnersSelectionDate          int64     `json:"winners_selection_date"`
	WinnerCount                   int       `json:"winner_count"`
	OnlyNewMembers                *bool     `json:"only_new_members"`
	HasPublicWinners              *bool     `json:"has_public_winners"`
	PrizeDescription              *string   `json:"prize_description"`
	CountryCodes                  *[]string `json:"country_codes"`
	PrizeStarCount                *int64    `json:"prize_star_count"`
	PremiumSubscriptionMonthCount *int      `json:"premium_subscription_month_count"`
}

// GiveawayWinners describes winners of a completed giveaway.
type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`
	GiveawayMessageID             int64   `json:"giveaway_message_id"`
	WinnersSelectionDate          int64   `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []User  `json:"winners"`
	AdditionalChatCount           *int    `json:"additional_chat_count"`
	PrizeStarCount                *int64  `json:"prize_star_count"`
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count"`
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count"`
	OnlyNewMembers                *bool   `json:"only_new_members"`
	WasRefunded                   *bool   `json:"was_refunded"`
	PrizeDescription              *string `json:"prize_description"`
}

// GiveawayCompleted describes completion of a giveaway without public winners.
type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count"`
	GiveawayMessage     *Message `json:"giveaway_message"`
	IsStarGiveaway      *bool    `json:"is_star_giveaway"`
}

// CommunityChatAdded describes a chat or bot added to a community.
type CommunityChatAdded struct {
	Community Community `json:"community"`
}

// CommunityChatJoined describes a chat joined by a user from a community.
type CommunityChatJoined struct {
	Community Community `json:"community"`
}

type CommunityChatRemoved struct{}

// PollOptionAdded describes an option added to a poll.
type PollOptionAdded struct {
	PollMessage        MaybeInaccessibleMessage `json:"poll_message"`
	OptionPersistentID string                   `json:"option_persistent_id"`
	OptionText         string                   `json:"option_text"`
	OptionTextEntities *[]MessageEntity         `json:"option_text_entities"`
}

// PollOptionDeleted describes an option deleted from a poll.
type PollOptionDeleted struct {
	PollMessage        MaybeInaccessibleMessage `json:"poll_message"`
	OptionPersistentID string                   `json:"option_persistent_id"`
	OptionText         string                   `json:"option_text"`
	OptionTextEntities *[]MessageEntity         `json:"option_text_entities"`
}

// ChatBoostAdded describes a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

// ChatBackground represents a chat background.
type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

// BackgroundFill describes how a background is filled.
type BackgroundFill interface {
	backgroundFill()
}

// BackgroundFillSolid describes a solid background fill.
type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

func (*BackgroundFillSolid) backgroundFill() {}

func (m BackgroundFillSolid) MarshalJSON() ([]byte, error) {
	type alias BackgroundFillSolid
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "solid", alias: alias(m)})
}

// BackgroundFillGradient describes a gradient background fill.
type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

func (*BackgroundFillGradient) backgroundFill() {}

func (m BackgroundFillGradient) MarshalJSON() ([]byte, error) {
	type alias BackgroundFillGradient
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "gradient", alias: alias(m)})
}

// BackgroundFillFreeformGradient describes a freeform gradient background fill.
type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`
	Colors []int  `json:"colors"`
}

func (*BackgroundFillFreeformGradient) backgroundFill() {}

func (m BackgroundFillFreeformGradient) MarshalJSON() ([]byte, error) {
	type alias BackgroundFillFreeformGradient
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "freeform_gradient", alias: alias(m)})
}

// BackgroundType describes the type of a background.
type BackgroundType interface {
	backgroundType()
}

// BackgroundTypeFill describes an automatically colored background.
type BackgroundTypeFill struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

func (*BackgroundTypeFill) backgroundType() {}

func (m BackgroundTypeFill) MarshalJSON() ([]byte, error) {
	type alias BackgroundTypeFill
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "fill", alias: alias(m)})
}

// BackgroundTypeWallpaper describes a JPEG wallpaper background.
type BackgroundTypeWallpaper struct {
	Type             string   `json:"type"`
	Document         Document `json:"document"`
	DarkThemeDimming int      `json:"dark_theme_dimming"`
	IsBlurred        *bool    `json:"is_blurred"`
	IsMoving         *bool    `json:"is_moving"`
}

func (*BackgroundTypeWallpaper) backgroundType() {}

func (m BackgroundTypeWallpaper) MarshalJSON() ([]byte, error) {
	type alias BackgroundTypeWallpaper
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "wallpaper", alias: alias(m)})
}

// BackgroundTypePattern describes a pattern background.
type BackgroundTypePattern struct {
	Type       string         `json:"type"`
	Document   Document       `json:"document"`
	Fill       BackgroundFill `json:"fill"`
	Intensity  int            `json:"intensity"`
	IsInverted *bool          `json:"is_inverted"`
	IsMoving   *bool          `json:"is_moving"`
}

func (*BackgroundTypePattern) backgroundType() {}

func (m BackgroundTypePattern) MarshalJSON() ([]byte, error) {
	type alias BackgroundTypePattern
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "pattern", alias: alias(m)})
}

// BackgroundTypeChatTheme describes a built-in chat theme background.
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

func (*BackgroundTypeChatTheme) backgroundType() {}

func (m BackgroundTypeChatTheme) MarshalJSON() ([]byte, error) {
	type alias BackgroundTypeChatTheme
	return json.Marshal(struct {
		Type string `json:"type"`
		alias
	}{Type: "chat_theme", alias: alias(m)})
}

func decodeMaybeInaccessibleMessage(data []byte) (MaybeInaccessibleMessage, error) {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return nil, nil
	}
	var probe struct {
		Date int64 `json:"date"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	if probe.Date == 0 {
		var v InaccessibleMessage
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	}
	var v Message
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func decodeMessageOrigin(data []byte) (MessageOrigin, error) {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return nil, nil
	}
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case "user":
		var v MessageOriginUser
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "hidden_user":
		var v MessageOriginHiddenUser
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "chat":
		var v MessageOriginChat
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "channel":
		var v MessageOriginChannel
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("telegram: unknown message origin type %q", probe.Type)
	}
}

func decodeBackgroundFill(data []byte) (BackgroundFill, error) {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return nil, nil
	}
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case "solid":
		var v BackgroundFillSolid
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "gradient":
		var v BackgroundFillGradient
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "freeform_gradient":
		var v BackgroundFillFreeformGradient
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("telegram: unknown background fill type %q", probe.Type)
	}
}

func decodeBackgroundType(data []byte) (BackgroundType, error) {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		return nil, nil
	}
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case "fill":
		var v BackgroundTypeFill
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "wallpaper":
		var v BackgroundTypeWallpaper
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "pattern":
		var v BackgroundTypePattern
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "chat_theme":
		var v BackgroundTypeChatTheme
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("telegram: unknown background type %q", probe.Type)
	}
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type alias Message
	var aux struct {
		ForwardOrigin json.RawMessage `json:"forward_origin"`
		PinnedMessage json.RawMessage `json:"pinned_message"`
		*alias
	}
	aux.alias = (*alias)(m)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.ForwardOrigin) != 0 {
		origin, err := decodeMessageOrigin(aux.ForwardOrigin)
		if err != nil {
			return err
		}
		m.ForwardOrigin = origin
	}
	if len(aux.PinnedMessage) != 0 {
		pinned, err := decodeMaybeInaccessibleMessage(aux.PinnedMessage)
		if err != nil {
			return err
		}
		m.PinnedMessage = pinned
	}
	return nil
}

func (e *ExternalReplyInfo) UnmarshalJSON(data []byte) error {
	type alias ExternalReplyInfo
	var aux struct {
		Origin json.RawMessage `json:"origin"`
		*alias
	}
	aux.alias = (*alias)(e)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Origin) != 0 {
		origin, err := decodeMessageOrigin(aux.Origin)
		if err != nil {
			return err
		}
		e.Origin = origin
	}
	return nil
}

func (p *PollOptionAdded) UnmarshalJSON(data []byte) error {
	type alias PollOptionAdded
	var aux struct {
		PollMessage json.RawMessage `json:"poll_message"`
		*alias
	}
	aux.alias = (*alias)(p)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.PollMessage) != 0 {
		pollMessage, err := decodeMaybeInaccessibleMessage(aux.PollMessage)
		if err != nil {
			return err
		}
		p.PollMessage = pollMessage
	}
	return nil
}

func (p *PollOptionDeleted) UnmarshalJSON(data []byte) error {
	type alias PollOptionDeleted
	var aux struct {
		PollMessage json.RawMessage `json:"poll_message"`
		*alias
	}
	aux.alias = (*alias)(p)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.PollMessage) != 0 {
		pollMessage, err := decodeMaybeInaccessibleMessage(aux.PollMessage)
		if err != nil {
			return err
		}
		p.PollMessage = pollMessage
	}
	return nil
}

func (b *ChatBackground) UnmarshalJSON(data []byte) error {
	type alias ChatBackground
	var aux struct {
		Type json.RawMessage `json:"type"`
		*alias
	}
	aux.alias = (*alias)(b)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Type) != 0 {
		backgroundType, err := decodeBackgroundType(aux.Type)
		if err != nil {
			return err
		}
		b.Type = backgroundType
	}
	return nil
}

func (b *BackgroundTypeFill) UnmarshalJSON(data []byte) error {
	type alias BackgroundTypeFill
	var aux struct {
		Fill json.RawMessage `json:"fill"`
		*alias
	}
	aux.alias = (*alias)(b)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Fill) != 0 {
		fill, err := decodeBackgroundFill(aux.Fill)
		if err != nil {
			return err
		}
		b.Fill = fill
	}
	return nil
}

func (b *BackgroundTypePattern) UnmarshalJSON(data []byte) error {
	type alias BackgroundTypePattern
	var aux struct {
		Fill json.RawMessage `json:"fill"`
		*alias
	}
	aux.alias = (*alias)(b)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Fill) != 0 {
		fill, err := decodeBackgroundFill(aux.Fill)
		if err != nil {
			return err
		}
		b.Fill = fill
	}
	return nil
}

// PhotoSize represents one size of a photo or a file/sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     *int64 `json:"file_size"`
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
	FileName     *string    `json:"file_name"`
	MimeType     *string    `json:"mime_type"`
	FileSize     *int64     `json:"file_size"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    *string    `json:"performer"`
	Title        *string    `json:"title"`
	FileName     *string    `json:"file_name"`
	MimeType     *string    `json:"mime_type"`
	FileSize     *int64     `json:"file_size"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
	FileName     *string    `json:"file_name"`
	MimeType     *string    `json:"mime_type"`
	FileSize     *int64     `json:"file_size"`
}

// LivePhoto represents a live photo.
type LivePhoto struct {
	Photo        *[]PhotoSize `json:"photo"`
	FileID       string       `json:"file_id"`
	FileUniqueID string       `json:"file_unique_id"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	Duration     int          `json:"duration"`
	MimeType     *string      `json:"mime_type"`
	FileSize     *int64       `json:"file_size"`
}

// Story represents a story.
type Story struct {
	Chat Chat  `json:"chat"`
	ID   int64 `json:"id"`
}

// VideoQuality represents a video file of a specific quality.
type VideoQuality struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Codec        string `json:"codec"`
	FileSize     *int64 `json:"file_size"`
}

// Video represents a video file.
type Video struct {
	FileID         string          `json:"file_id"`
	FileUniqueID   string          `json:"file_unique_id"`
	Width          int             `json:"width"`
	Height         int             `json:"height"`
	Duration       int             `json:"duration"`
	Thumbnail      *PhotoSize      `json:"thumbnail"`
	Cover          *[]PhotoSize    `json:"cover"`
	StartTimestamp *int            `json:"start_timestamp"`
	Qualities      *[]VideoQuality `json:"qualities"`
	FileName       *string         `json:"file_name"`
	MimeType       *string         `json:"mime_type"`
	FileSize       *int64          `json:"file_size"`
}

// VideoNote represents a video message.
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail"`
	FileSize     *int64     `json:"file_size"`
}

// Voice represents a voice note.
type Voice struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type"`
	FileSize     *int64  `json:"file_size"`
}

// PaidMediaInfo describes the paid media added to a message.
type PaidMediaInfo struct {
	StarCount int64       `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

// PaidMedia describes paid media. It can be one of PaidMediaPreview,
// PaidMediaPhoto, PaidMediaVideo or PaidMediaLivePhoto.
type PaidMedia interface {
	paidMedia()
}

// PaidMediaPreview means the paid media isn't available before the payment.
type PaidMediaPreview struct {
	Width    *int `json:"width"`
	Height   *int `json:"height"`
	Duration *int `json:"duration"`
}

// MarshalJSON implements the json.Marshaler interface, injecting the "preview" discriminator.
func (v PaidMediaPreview) MarshalJSON() ([]byte, error) {
	type alias PaidMediaPreview
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "preview"})
}

// PaidMediaPhoto means the paid media is a photo.
type PaidMediaPhoto struct {
	Photo []PhotoSize `json:"photo"`
}

// MarshalJSON implements the json.Marshaler interface, injecting the "photo" discriminator.
func (v PaidMediaPhoto) MarshalJSON() ([]byte, error) {
	type alias PaidMediaPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// PaidMediaVideo means the paid media is a video.
type PaidMediaVideo struct {
	Video Video `json:"video"`
}

// MarshalJSON implements the json.Marshaler interface, injecting the "video" discriminator.
func (v PaidMediaVideo) MarshalJSON() ([]byte, error) {
	type alias PaidMediaVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// PaidMediaLivePhoto means the paid media is a live photo.
type PaidMediaLivePhoto struct {
	LivePhoto LivePhoto `json:"live_photo"`
}

// MarshalJSON implements the json.Marshaler interface, injecting the "live_photo" discriminator.
func (v PaidMediaLivePhoto) MarshalJSON() ([]byte, error) {
	type alias PaidMediaLivePhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "live_photo"})
}

func (*PaidMediaPreview) paidMedia()   {}
func (*PaidMediaPhoto) paidMedia()     {}
func (*PaidMediaVideo) paidMedia()     {}
func (*PaidMediaLivePhoto) paidMedia() {}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	UserID      *int64  `json:"user_id"`
	Vcard       *string `json:"vcard"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// Link represents an HTTP link.
type Link struct {
	URL string `json:"url"`
}

// PollMedia describes the media attached to a poll description, explanation
// or option. At most one of the optional fields can be present.
type PollMedia struct {
	Animation *Animation   `json:"animation"`
	Audio     *Audio       `json:"audio"`
	Document  *Document    `json:"document"`
	Link      *Link        `json:"link"`
	LivePhoto *LivePhoto   `json:"live_photo"`
	Location  *Location    `json:"location"`
	Photo     *[]PhotoSize `json:"photo"`
	Sticker   *Sticker     `json:"sticker"`
	Venue     *Venue       `json:"venue"`
	Video     *Video       `json:"video"`
}

// InputPollMedia represents the content of a poll description or a quiz
// explanation to be sent: one of the InputMedia variants.
type InputPollMedia = InputMedia

// InputPollOptionMedia represents the content of a poll option to be sent:
// one of the InputMedia variants.
type InputPollOptionMedia = InputMedia

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	PersistentID string           `json:"persistent_id"`
	Text         string           `json:"text"`
	TextEntities *[]MessageEntity `json:"text_entities"`
	Media        *PollMedia       `json:"media"`
	VoterCount   int              `json:"voter_count"`
	AddedByUser  *User            `json:"added_by_user"`
	AddedByChat  *Chat            `json:"added_by_chat"`
	AdditionDate *int64           `json:"addition_date"`
}

// InputPollOption contains information about one answer option in a poll to be sent.
type InputPollOption struct {
	Text          string               `json:"text"`
	TextParseMode *string              `json:"text_parse_mode,omitempty"`
	TextEntities  *[]MessageEntity     `json:"text_entities,omitempty"`
	Media         InputPollOptionMedia `json:"media,omitempty"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID              string   `json:"poll_id"`
	VoterChat           *Chat    `json:"voter_chat"`
	User                *User    `json:"user"`
	OptionIDs           []int    `json:"option_ids"`
	OptionPersistentIDs []string `json:"option_persistent_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	QuestionEntities      *[]MessageEntity `json:"question_entities"`
	Options               []PollOption     `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	AllowsRevoting        bool             `json:"allows_revoting"`
	MembersOnly           bool             `json:"members_only"`
	CountryCodes          *[]string        `json:"country_codes"`
	CorrectOptionIDs      *[]int           `json:"correct_option_ids"`
	Explanation           *string          `json:"explanation"`
	ExplanationEntities   *[]MessageEntity `json:"explanation_entities"`
	ExplanationMedia      *PollMedia       `json:"explanation_media"`
	OpenPeriod            *int             `json:"open_period"`
	CloseDate             *int64           `json:"close_date"`
	Description           *string          `json:"description"`
	DescriptionEntities   *[]MessageEntity `json:"description_entities"`
	Media                 *PollMedia       `json:"media"`
}

// File represents a file ready to be downloaded.
type File struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int64  `json:"file_size"`
	FilePath     *string `json:"file_path"`
}

// UserProfilePhotos represents a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// UserProfileAudios represents the audios displayed on a user's profile.
type UserProfileAudios struct {
	TotalCount int     `json:"total_count"`
	Audios     []Audio `json:"audios"`
}

// ReplyMarkup is the interface implemented by all reply markup types:
// InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove and ForceReply.
type ReplyMarkup interface {
	replyMarkup()
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next
// to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
	ForceReply     *bool                    `json:"force_reply"`
}

func (*InlineKeyboardMarkup) replyMarkup() {}

// InlineKeyboardButton represents one button of an inline keyboard. Exactly
// one of the fields other than text, icon_custom_emoji_id and style must be
// used to specify the type of the button.
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	IconCustomEmojiID            *string                      `json:"icon_custom_emoji_id"`
	Style                        *string                      `json:"style"`
	URL                          *string                      `json:"url"`
	CallbackData                 *string                      `json:"callback_data"`
	WebApp                       *WebAppInfo                  `json:"web_app"`
	LoginURL                     *LoginUrl                    `json:"login_url"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`
	CopyText                     *CopyTextButton              `json:"copy_text"`
	CallbackGame                 *CallbackGame                `json:"callback_game"`
	Pay                          *bool                        `json:"pay"`
	Disabled                     *DisabledButton              `json:"disabled"`
}

// LoginUrl represents a parameter of the inline keyboard button used to
// automatically authorize a user.
type LoginUrl struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text"`
	BotUsername        *string `json:"bot_username"`
	RequestWriteAccess *bool   `json:"request_write_access"`
}

// SwitchInlineQueryChosenChat represents an inline button that switches the
// current user to inline mode in a chosen chat, with an optional default
// inline query.
type SwitchInlineQueryChosenChat struct {
	Query             *string `json:"query"`
	AllowUserChats    *bool   `json:"allow_user_chats"`
	AllowBotChats     *bool   `json:"allow_bot_chats"`
	AllowGroupChats   *bool   `json:"allow_group_chats"`
	AllowChannelChats *bool   `json:"allow_channel_chats"`
}

// CopyTextButton represents an inline keyboard button that copies the
// specified text to the clipboard.
type CopyTextButton struct {
	Text string `json:"text"`
}

// DisabledButton represents a disabled button which does nothing.
type DisabledButton struct{}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          *bool              `json:"is_persistent"`
	ResizeKeyboard        *bool              `json:"resize_keyboard"`
	OneTimeKeyboard       *bool              `json:"one_time_keyboard"`
	InputFieldPlaceholder *string            `json:"input_field_placeholder"`
	Selective             *bool              `json:"selective"`
	ForceReply            *bool              `json:"force_reply"`
}

func (*ReplyKeyboardMarkup) replyMarkup() {}

// KeyboardButton represents one button of the reply keyboard.
type KeyboardButton struct {
	Text              string                           `json:"text"`
	IconCustomEmojiID *string                          `json:"icon_custom_emoji_id"`
	Style             *string                          `json:"style"`
	RequestUsers      *KeyboardButtonRequestUsers      `json:"request_users"`
	RequestChat       *KeyboardButtonRequestChat       `json:"request_chat"`
	RequestManagedBot *KeyboardButtonRequestManagedBot `json:"request_managed_bot"`
	RequestContact    *bool                            `json:"request_contact"`
	RequestLocation   *bool                            `json:"request_location"`
	RequestPoll       *KeyboardButtonPollType          `json:"request_poll"`
	WebApp            *WebAppInfo                      `json:"web_app"`
}

// KeyboardButtonRequestUsers defines the criteria used to request suitable
// users.
type KeyboardButtonRequestUsers struct {
	RequestID       int64 `json:"request_id"`
	UserIsBot       *bool `json:"user_is_bot"`
	UserIsPremium   *bool `json:"user_is_premium"`
	MaxQuantity     *int  `json:"max_quantity"`
	RequestName     *bool `json:"request_name"`
	RequestUsername *bool `json:"request_username"`
	RequestPhoto    *bool `json:"request_photo"`
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable
// chat.
type KeyboardButtonRequestChat struct {
	RequestID               int64                    `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             *bool                    `json:"chat_is_forum"`
	ChatHasUsername         *bool                    `json:"chat_has_username"`
	ChatIsCreated           *bool                    `json:"chat_is_created"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights"`
	BotIsMember             *bool                    `json:"bot_is_member"`
	RequestTitle            *bool                    `json:"request_title"`
	RequestUsername         *bool                    `json:"request_username"`
	RequestPhoto            *bool                    `json:"request_photo"`
}

// KeyboardButtonRequestManagedBot defines the parameters for the creation of
// a managed bot.
type KeyboardButtonRequestManagedBot struct {
	RequestID         int64   `json:"request_id"`
	SuggestedName     *string `json:"suggested_name"`
	SuggestedUsername *string `json:"suggested_username"`
}

// KeyboardButtonPollType represents the type of a poll which is allowed to be
// created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type *string `json:"type"`
}

// ReplyKeyboardRemove requests clients to remove the current custom keyboard
// and display the default letter-keyboard.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective"`
}

func (*ReplyKeyboardRemove) replyMarkup() {}

// ForceReply requests clients to display a reply interface to the user.
type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	InputFieldPlaceholder *string `json:"input_field_placeholder"`
	Selective             *bool   `json:"selective"`
}

func (*ForceReply) replyMarkup() {}

// CallbackQuery represents an incoming callback query from a callback button
// in an inline keyboard.
type CallbackQuery struct {
	ID              string                   `json:"id"`
	From            User                     `json:"from"`
	Message         MaybeInaccessibleMessage `json:"message"`
	InlineMessageID *string                  `json:"inline_message_id"`
	ChatInstance    string                   `json:"chat_instance"`
	Data            *string                  `json:"data"`
	GameShortName   *string                  `json:"game_short_name"`
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name"`
	ExpireDate              *int64  `json:"expire_date"`
	MemberLimit             *int    `json:"member_limit"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count"`
	SubscriptionPeriod      *int    `json:"subscription_period"`
	SubscriptionPrice       *int64  `json:"subscription_price"`
}

// ChatAdministratorRights represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	IsAnonymous             bool `json:"is_anonymous"`
	CanManageChat           bool `json:"can_manage_chat"`
	CanDeleteMessages       bool `json:"can_delete_messages"`
	CanManageVideoChats     bool `json:"can_manage_video_chats"`
	CanRestrictMembers      bool `json:"can_restrict_members"`
	CanPromoteMembers       bool `json:"can_promote_members"`
	CanChangeInfo           bool `json:"can_change_info"`
	CanInviteUsers          bool `json:"can_invite_users"`
	CanPostStories          bool `json:"can_post_stories"`
	CanEditStories          bool `json:"can_edit_stories"`
	CanDeleteStories        bool `json:"can_delete_stories"`
	CanPostMessages         bool `json:"can_post_messages"`
	CanEditMessages         bool `json:"can_edit_messages"`
	CanPinMessages          bool `json:"can_pin_messages"`
	CanManageTopics         bool `json:"can_manage_topics"`
	CanManageDirectMessages bool `json:"can_manage_direct_messages"`
	CanManageTags           bool `json:"can_manage_tags"`
	CanSendWelcomeMessages  bool `json:"can_send_welcome_messages"`
}

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int64           `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link"`
	ViaJoinRequest          *bool           `json:"via_join_request"`
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link"`
}

// ChatMember contains information about one member of a chat.
type ChatMember interface {
	chatMember()
}

// ChatMemberOwner represents a chat member that owns the chat and has all
// administrator privileges.
type ChatMemberOwner struct {
	User        User    `json:"user"`
	IsAnonymous bool    `json:"is_anonymous"`
	CustomTitle *string `json:"custom_title"`
}

func (*ChatMemberOwner) chatMember() {}

func (v ChatMemberOwner) MarshalJSON() ([]byte, error) {
	type alias ChatMemberOwner
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{alias(v), "creator"})
}

// ChatMemberAdministrator represents a chat member that has some additional
// privileges.
type ChatMemberAdministrator struct {
	User                    User    `json:"user"`
	CanBeEdited             bool    `json:"can_be_edited"`
	IsAnonymous             bool    `json:"is_anonymous"`
	CanManageChat           bool    `json:"can_manage_chat"`
	CanDeleteMessages       bool    `json:"can_delete_messages"`
	CanManageVideoChats     bool    `json:"can_manage_video_chats"`
	CanRestrictMembers      bool    `json:"can_restrict_members"`
	CanPromoteMembers       bool    `json:"can_promote_members"`
	CanChangeInfo           bool    `json:"can_change_info"`
	CanInviteUsers          bool    `json:"can_invite_users"`
	CanPostStories          bool    `json:"can_post_stories"`
	CanEditStories          bool    `json:"can_edit_stories"`
	CanDeleteStories        bool    `json:"can_delete_stories"`
	CanPostMessages         *bool   `json:"can_post_messages"`
	CanEditMessages         *bool   `json:"can_edit_messages"`
	CanPinMessages          *bool   `json:"can_pin_messages"`
	CanManageTopics         *bool   `json:"can_manage_topics"`
	CanManageDirectMessages *bool   `json:"can_manage_direct_messages"`
	CanManageTags           *bool   `json:"can_manage_tags"`
	CanSendWelcomeMessages  bool    `json:"can_send_welcome_messages"`
	CustomTitle             *string `json:"custom_title"`
}

func (*ChatMemberAdministrator) chatMember() {}

func (v ChatMemberAdministrator) MarshalJSON() ([]byte, error) {
	type alias ChatMemberAdministrator
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{alias(v), "administrator"})
}

// ChatMemberMember represents a chat member that has no additional privileges
// or restrictions.
type ChatMemberMember struct {
	Tag       *string `json:"tag"`
	User      User    `json:"user"`
	UntilDate *int64  `json:"until_date"`
}

func (*ChatMemberMember) chatMember() {}

func (v ChatMemberMember) MarshalJSON() ([]byte, error) {
	type alias ChatMemberMember
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{alias(v), "member"})
}

// ChatMemberRestricted represents a chat member that is under certain
// restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
	Tag                   *string `json:"tag"`
	User                  User    `json:"user"`
	IsMember              bool    `json:"is_member"`
	CanSendMessages       bool    `json:"can_send_messages"`
	CanSendAudios         bool    `json:"can_send_audios"`
	CanSendDocuments      bool    `json:"can_send_documents"`
	CanSendPhotos         bool    `json:"can_send_photos"`
	CanSendVideos         bool    `json:"can_send_videos"`
	CanSendVideoNotes     bool    `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool    `json:"can_send_voice_notes"`
	CanSendPolls          bool    `json:"can_send_polls"`
	CanSendOtherMessages  bool    `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool    `json:"can_add_web_page_previews"`
	CanReactToMessages    bool    `json:"can_react_to_messages"`
	CanEditTag            bool    `json:"can_edit_tag"`
	CanChangeInfo         bool    `json:"can_change_info"`
	CanInviteUsers        bool    `json:"can_invite_users"`
	CanPinMessages        bool    `json:"can_pin_messages"`
	CanManageTopics       bool    `json:"can_manage_topics"`
	UntilDate             int64   `json:"until_date"`
}

func (*ChatMemberRestricted) chatMember() {}

func (v ChatMemberRestricted) MarshalJSON() ([]byte, error) {
	type alias ChatMemberRestricted
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{alias(v), "restricted"})
}

// ChatMemberLeft represents a chat member that isn't currently a member of
// the chat, but may join it themselves.
type ChatMemberLeft struct {
	User User `json:"user"`
}

func (*ChatMemberLeft) chatMember() {}

func (v ChatMemberLeft) MarshalJSON() ([]byte, error) {
	type alias ChatMemberLeft
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{alias(v), "left"})
}

// ChatMemberBanned represents a chat member that was banned in the chat and
// can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	User      User  `json:"user"`
	UntilDate int64 `json:"until_date"`
}

func (*ChatMemberBanned) chatMember() {}

func (v ChatMemberBanned) MarshalJSON() ([]byte, error) {
	type alias ChatMemberBanned
	return json.Marshal(struct {
		alias
		Status string `json:"status"`
	}{alias(v), "kicked"})
}

// ChatJoinRequest represents a join request sent to a chat.
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int64           `json:"date"`
	Bio        *string         `json:"bio"`
	InviteLink *ChatInviteLink `json:"invite_link"`
	QueryID    *string         `json:"query_id"`
}

// ChatPermissions describes actions that a non-administrator user is allowed
// to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanReactToMessages    bool `json:"can_react_to_messages"`
	CanEditTag            bool `json:"can_edit_tag"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics"`
}

// ReactionType describes the type of a reaction.
type ReactionType interface {
	reactionType()
}

// ReactionTypeEmoji is a reaction based on an emoji.
type ReactionTypeEmoji struct {
	Emoji string `json:"emoji"`
}

func (*ReactionTypeEmoji) reactionType() {}

func (v ReactionTypeEmoji) MarshalJSON() ([]byte, error) {
	type alias ReactionTypeEmoji
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "emoji"})
}

// ReactionTypeCustomEmoji is a reaction based on a custom emoji.
type ReactionTypeCustomEmoji struct {
	CustomEmojiID string `json:"custom_emoji_id"`
}

func (*ReactionTypeCustomEmoji) reactionType() {}

func (v ReactionTypeCustomEmoji) MarshalJSON() ([]byte, error) {
	type alias ReactionTypeCustomEmoji
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "custom_emoji"})
}

// ReactionTypePaid is a paid reaction.
type ReactionTypePaid struct{}

func (*ReactionTypePaid) reactionType() {}

func (v ReactionTypePaid) MarshalJSON() ([]byte, error) {
	type alias ReactionTypePaid
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "paid"})
}

// ReactionCount represents a reaction added to a message along with the
// number of times it was added.
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

// MessageReactionUpdated represents a change of a reaction on a message
// performed by a user.
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int64          `json:"message_id"`
	User        *User          `json:"user"`
	ActorChat   *Chat          `json:"actor_chat"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

// MessageReactionCountUpdated represents reaction changes on a message with
// anonymous reactions.
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int64           `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

// ForumTopic represents a forum topic.
type ForumTopic struct {
	MessageThreadID   int64   `json:"message_thread_id"`
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id"`
	IsNameImplicit    *bool   `json:"is_name_implicit"`
}
