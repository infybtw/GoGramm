// Package gogram is a high-level Telegram bot layer on top of the low-level
// api package: command and update-type handler registration plus a
// long-polling dispatcher.
package gogram

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/infybtw/GoGramm/api"
)

// Api is the low-level Bot API client the high-level layer runs on. It is an
// alias, so a client built with api.New plugs in directly.
type Api = api.Client

// Update is the complete update received from Telegram.
type Update = api.Update

// Handler handles a single update. A returned error propagates to Bot.Start,
// which stops polling and returns it unchanged.
type Handler func(*Context) error

// UpdateType identifies which optional api.Update field an incoming update
// carries.
type UpdateType string

// Update types, one per optional api.Update field.
const (
	UpdateMessage                  UpdateType = "message"
	UpdateEditedMessage            UpdateType = "edited_message"
	UpdateChannelPost              UpdateType = "channel_post"
	UpdateEditedChannelPost        UpdateType = "edited_channel_post"
	UpdateBusinessConnection       UpdateType = "business_connection"
	UpdateBusinessMessage          UpdateType = "business_message"
	UpdateEditedBusinessMessage    UpdateType = "edited_business_message"
	UpdateDeletedBusinessMessages  UpdateType = "deleted_business_messages"
	UpdateGuestMessage             UpdateType = "guest_message"
	UpdateMessageReaction          UpdateType = "message_reaction"
	UpdateMessageReactionCount     UpdateType = "message_reaction_count"
	UpdateInlineQuery              UpdateType = "inline_query"
	UpdateChosenInlineResult       UpdateType = "chosen_inline_result"
	UpdateCallbackQuery            UpdateType = "callback_query"
	UpdateShippingQuery            UpdateType = "shipping_query"
	UpdatePreCheckoutQuery         UpdateType = "pre_checkout_query"
	UpdatePurchasedPaidMedia       UpdateType = "purchased_paid_media"
	UpdatePoll                     UpdateType = "poll"
	UpdatePollAnswer               UpdateType = "poll_answer"
	UpdateMyChatMember             UpdateType = "my_chat_member"
	UpdateChatMember               UpdateType = "chat_member"
	UpdateChatJoinRequest          UpdateType = "chat_join_request"
	UpdateChatBoost                UpdateType = "chat_boost"
	UpdateRemovedChatBoost         UpdateType = "removed_chat_boost"
	UpdateManagedBot               UpdateType = "managed_bot"
	UpdateSubscription             UpdateType = "subscription"
	UpdateStoppedMessageGeneration UpdateType = "stopped_message_generation"
)

// Context carries one update to its handler.
type Context struct {
	Bot    *Bot
	Api    *Api
	Update *Update
}

// Bot is a Telegram bot: handler registries plus a low-level API client.
// Registration and polling are safe for concurrent use.
type Bot struct {
	// Api is the low-level client used for polling and available to
	// handlers via Context.Api. It is exported so callers can substitute
	// a client built with api.New and api.Options.
	Api *Api

	mu        sync.RWMutex
	commands  map[string]Handler
	callbacks map[string]Handler
	handlers  map[UpdateType]Handler
	username  string
}

// NewBot returns a Bot authenticating with token. Replace the exported Api
// field to point the bot at a custom client (for example one built with
// api.WithServerURL).
func NewBot(token string) *Bot {
	return &Bot{
		Api:       api.New(token),
		commands:  make(map[string]Handler),
		callbacks: make(map[string]Handler),
		handlers:  make(map[UpdateType]Handler),
	}
}

// Command registers handler for name, the command without its leading
// slash. Registration replaces any previous handler for the command; a nil
// handler removes it.
func (b *Bot) Command(name string, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if handler == nil {
		delete(b.commands, name)
		return
	}
	b.commands[name] = handler
}

// CallbackQuery registers handler for an inline button whose callback_data
// exactly matches data. Registration replaces any previous handler for data;
// a nil handler removes it. A matching handler takes precedence over the
// general UpdateCallbackQuery handler registered with On.
func (b *Bot) CallbackQuery(data string, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if handler == nil {
		delete(b.callbacks, data)
		return
	}
	b.callbacks[data] = handler
}

// On registers handler for updateType. Registration replaces any previous
// handler for the type; a nil handler removes it.
func (b *Bot) On(updateType UpdateType, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if handler == nil {
		delete(b.handlers, updateType)
		return
	}
	b.handlers[updateType] = handler
}

// OnMessage registers handler for message updates.
func (b *Bot) OnMessage(handler Handler) {
	b.On(UpdateMessage, handler)
}

// classify reports which optional api.Update field u carries, inspecting
// members in declaration order. It returns "" when none is set.
func classify(u *api.Update) UpdateType {
	switch {
	case u.Message != nil:
		return UpdateMessage
	case u.EditedMessage != nil:
		return UpdateEditedMessage
	case u.ChannelPost != nil:
		return UpdateChannelPost
	case u.EditedChannelPost != nil:
		return UpdateEditedChannelPost
	case u.BusinessConnection != nil:
		return UpdateBusinessConnection
	case u.BusinessMessage != nil:
		return UpdateBusinessMessage
	case u.EditedBusinessMessage != nil:
		return UpdateEditedBusinessMessage
	case u.DeletedBusinessMessages != nil:
		return UpdateDeletedBusinessMessages
	case u.GuestMessage != nil:
		return UpdateGuestMessage
	case u.MessageReaction != nil:
		return UpdateMessageReaction
	case u.MessageReactionCount != nil:
		return UpdateMessageReactionCount
	case u.InlineQuery != nil:
		return UpdateInlineQuery
	case u.ChosenInlineResult != nil:
		return UpdateChosenInlineResult
	case u.CallbackQuery != nil:
		return UpdateCallbackQuery
	case u.ShippingQuery != nil:
		return UpdateShippingQuery
	case u.PreCheckoutQuery != nil:
		return UpdatePreCheckoutQuery
	case u.PurchasedPaidMedia != nil:
		return UpdatePurchasedPaidMedia
	case u.Poll != nil:
		return UpdatePoll
	case u.PollAnswer != nil:
		return UpdatePollAnswer
	case u.MyChatMember != nil:
		return UpdateMyChatMember
	case u.ChatMember != nil:
		return UpdateChatMember
	case u.ChatJoinRequest != nil:
		return UpdateChatJoinRequest
	case u.ChatBoost != nil:
		return UpdateChatBoost
	case u.RemovedChatBoost != nil:
		return UpdateRemovedChatBoost
	case u.ManagedBot != nil:
		return UpdateManagedBot
	case u.Subscription != nil:
		return UpdateSubscription
	case u.StoppedMessageGeneration != nil:
		return UpdateStoppedMessageGeneration
	default:
		return ""
	}
}

// commandHandler parses the first whitespace-delimited token of a text
// message. A command addressed to this bot with @botname uses the same key as
// an unqualified command; commands addressed to another bot fall through to
// the message handler. Returns nil when the message is not a command with a
// registered handler.
func (b *Bot) commandHandler(msg *api.Message) Handler {
	if msg == nil || msg.Text == nil {
		return nil
	}
	text := *msg.Text
	if !strings.HasPrefix(text, "/") {
		return nil
	}
	token := text
	if i := strings.IndexFunc(text, unicode.IsSpace); i >= 0 {
		token = text[:i]
	}
	command := strings.TrimPrefix(token, "/")
	b.mu.RLock()
	defer b.mu.RUnlock()
	if name, target, ok := strings.Cut(command, "@"); ok {
		if b.username == "" || !strings.EqualFold(target, b.username) {
			return nil
		}
		command = name
	}
	return b.commands[command]
}

// callbackHandler returns the handler registered for q's callback data, if
// q carries data and the data has a registered handler.
func (b *Bot) callbackHandler(q *api.CallbackQuery) Handler {
	if q == nil || q.Data == nil {
		return nil
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.callbacks[*q.Data]
}

// handlerFor selects the handler for u: for a message update, a registered
// command handler wins over the message handler; a non-command message, a
// command without a registered handler, and all other update types use the
// update-type handler.
func (b *Bot) handlerFor(u *api.Update, typ UpdateType) Handler {
	if typ == UpdateMessage {
		if h := b.commandHandler(u.Message); h != nil {
			return h
		}
	}
	if typ == UpdateCallbackQuery {
		if h := b.callbackHandler(u.CallbackQuery); h != nil {
			return h
		}
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.handlers[typ]
}

// dispatch runs the handler selected for u, if any. Updates without a
// recognized member and updates with no registered handler are no-ops.
func (b *Bot) dispatch(u *api.Update) error {
	typ := classify(u)
	if typ == "" {
		return nil
	}
	h := b.handlerFor(u, typ)
	if h == nil {
		return nil
	}
	return h(&Context{Bot: b, Api: b.Api, Update: u})
}

const (
	// pollTimeoutSeconds is the getUpdates long-poll timeout.
	pollTimeoutSeconds = 30
	// retryDelay is the wait before re-polling after a transient failure
	// without a Retry-After hint.
	retryDelay = time.Second
)

// transient reports whether a polling failure is worth retrying: network
// and decode failures, 429, and 5xx API errors. Other API errors are
// permanent (bad request, auth, webhook conflicts, ...).
func transient(err error) bool {
	var apiErr *api.Error
	if !errors.As(err, &apiErr) {
		return true
	}
	return apiErr.Code == 429 || apiErr.Code >= 500
}

// retryDelayFor returns how long to wait after a failed poll before
// retrying: a 429's RetryAfter seconds when positive, otherwise one second.
func retryDelayFor(err error) time.Duration {
	var apiErr *api.Error
	if errors.As(err, &apiErr) && apiErr.Code == 429 &&
		apiErr.Parameters != nil && apiErr.Parameters.RetryAfter > 0 {
		return time.Duration(apiErr.Parameters.RetryAfter) * time.Second
	}
	return retryDelay
}

// Start polls for updates until ctx is canceled or a handler returns an
// error, dispatching each update in order. It obtains and caches the bot
// username before polling so it can route @botname commands safely. Handlers
// run inline. Transient polling failures (network errors, 429, 5xx) are
// retried after a context-aware wait; other API errors stop Start and are
// returned. The poll offset is not persisted: a fresh Start begins at offset
// zero.
func (b *Bot) Start(ctx context.Context) error {
	me, err := b.Api.GetMe(ctx)
	if err != nil {
		return err
	}
	b.mu.Lock()
	b.username = ""
	if me.Username != nil {
		b.username = *me.Username
	}
	b.mu.Unlock()

	var (
		offset      int64
		timeoutSecs = pollTimeoutSeconds
	)
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		updates, err := b.Api.GetUpdates(ctx, &api.GetUpdatesParams{
			Offset:  &offset,
			Timeout: &timeoutSecs,
		})
		if err != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				return ctxErr
			}
			if !transient(err) {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryDelayFor(err)):
			}
			continue
		}
		for i := range updates {
			offset = updates[i].UpdateID + 1
			if err := b.dispatch(&updates[i]); err != nil {
				return err
			}
		}
	}
}
