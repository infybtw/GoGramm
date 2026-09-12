package gogram

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/infybtw/GoGramm/api"
)

// recorder records the contexts a handler was invoked with.
type recorder struct {
	mu    sync.Mutex
	calls []*Context
}

func (r *recorder) handle(c *Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, c)
	return nil
}

func (r *recorder) count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.calls)
}

func (r *recorder) last() *Context {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.calls) == 0 {
		return nil
	}
	return r.calls[len(r.calls)-1]
}

// newTestBot returns a bot whose Api client is pointed at srv.
func newTestBot(t *testing.T, srv *httptest.Server) *Bot {
	t.Helper()
	b := NewBot("TESTTOKEN")
	b.Api = api.New("TESTTOKEN", api.WithServerURL(srv.URL), api.WithHTTPClient(srv.Client()))
	return b
}

// updateTypes lists every UpdateType constant.
func updateTypes() []UpdateType {
	return []UpdateType{
		UpdateMessage, UpdateEditedMessage, UpdateChannelPost, UpdateEditedChannelPost,
		UpdateBusinessConnection, UpdateBusinessMessage, UpdateEditedBusinessMessage,
		UpdateDeletedBusinessMessages, UpdateGuestMessage, UpdateMessageReaction,
		UpdateMessageReactionCount, UpdateInlineQuery, UpdateChosenInlineResult,
		UpdateCallbackQuery, UpdateShippingQuery, UpdatePreCheckoutQuery,
		UpdatePurchasedPaidMedia, UpdatePoll, UpdatePollAnswer, UpdateMyChatMember,
		UpdateChatMember, UpdateChatJoinRequest, UpdateChatBoost, UpdateRemovedChatBoost,
		UpdateManagedBot, UpdateSubscription, UpdateStoppedMessageGeneration,
	}
}

func TestDispatchClassifiesEveryUpdateType(t *testing.T) {
	tests := []struct {
		name string
		typ  UpdateType
		fill func(*api.Update)
	}{
		{"message", UpdateMessage, func(u *api.Update) { u.Message = &api.Message{} }},
		{"edited message", UpdateEditedMessage, func(u *api.Update) { u.EditedMessage = &api.Message{} }},
		{"channel post", UpdateChannelPost, func(u *api.Update) { u.ChannelPost = &api.Message{} }},
		{"edited channel post", UpdateEditedChannelPost, func(u *api.Update) { u.EditedChannelPost = &api.Message{} }},
		{"business connection", UpdateBusinessConnection, func(u *api.Update) { u.BusinessConnection = &api.BusinessConnection{} }},
		{"business message", UpdateBusinessMessage, func(u *api.Update) { u.BusinessMessage = &api.Message{} }},
		{"edited business message", UpdateEditedBusinessMessage, func(u *api.Update) { u.EditedBusinessMessage = &api.Message{} }},
		{"deleted business messages", UpdateDeletedBusinessMessages, func(u *api.Update) { u.DeletedBusinessMessages = &api.BusinessMessagesDeleted{} }},
		{"guest message", UpdateGuestMessage, func(u *api.Update) { u.GuestMessage = &api.Message{} }},
		{"message reaction", UpdateMessageReaction, func(u *api.Update) { u.MessageReaction = &api.MessageReactionUpdated{} }},
		{"message reaction count", UpdateMessageReactionCount, func(u *api.Update) { u.MessageReactionCount = &api.MessageReactionCountUpdated{} }},
		{"inline query", UpdateInlineQuery, func(u *api.Update) { u.InlineQuery = &api.InlineQuery{} }},
		{"chosen inline result", UpdateChosenInlineResult, func(u *api.Update) { u.ChosenInlineResult = &api.ChosenInlineResult{} }},
		{"callback query", UpdateCallbackQuery, func(u *api.Update) { u.CallbackQuery = &api.CallbackQuery{} }},
		{"shipping query", UpdateShippingQuery, func(u *api.Update) { u.ShippingQuery = &api.ShippingQuery{} }},
		{"pre checkout query", UpdatePreCheckoutQuery, func(u *api.Update) { u.PreCheckoutQuery = &api.PreCheckoutQuery{} }},
		{"purchased paid media", UpdatePurchasedPaidMedia, func(u *api.Update) { u.PurchasedPaidMedia = &api.PaidMediaPurchased{} }},
		{"poll", UpdatePoll, func(u *api.Update) { u.Poll = &api.Poll{} }},
		{"poll answer", UpdatePollAnswer, func(u *api.Update) { u.PollAnswer = &api.PollAnswer{} }},
		{"my chat member", UpdateMyChatMember, func(u *api.Update) { u.MyChatMember = &api.ChatMemberUpdated{} }},
		{"chat member", UpdateChatMember, func(u *api.Update) { u.ChatMember = &api.ChatMemberUpdated{} }},
		{"chat join request", UpdateChatJoinRequest, func(u *api.Update) { u.ChatJoinRequest = &api.ChatJoinRequest{} }},
		{"chat boost", UpdateChatBoost, func(u *api.Update) { u.ChatBoost = &api.ChatBoostUpdated{} }},
		{"removed chat boost", UpdateRemovedChatBoost, func(u *api.Update) { u.RemovedChatBoost = &api.ChatBoostRemoved{} }},
		{"managed bot", UpdateManagedBot, func(u *api.Update) { u.ManagedBot = &api.ManagedBotUpdated{} }},
		{"subscription", UpdateSubscription, func(u *api.Update) { u.Subscription = &api.BotSubscriptionUpdated{} }},
		{"stopped message generation", UpdateStoppedMessageGeneration, func(u *api.Update) { u.StoppedMessageGeneration = &api.MessageGenerationStopped{} }},
	}

	// One Update field per UpdateType plus UpdateID itself; a mismatch
	// means a new api.Update member lacks a classifier branch.
	if got := reflect.TypeOf(api.Update{}).NumField(); got != len(updateTypes())+1 {
		t.Fatalf("api.Update has %d fields, want %d (update types + update_id)", got, len(updateTypes())+1)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBot("TESTTOKEN")
			rec := &recorder{}
			b.On(tt.typ, rec.handle)

			u := api.Update{UpdateID: 7}
			tt.fill(&u)
			if err := b.dispatch(&u); err != nil {
				t.Fatalf("dispatch: %v", err)
			}
			if rec.count() != 1 {
				t.Fatalf("handler called %d times, want 1", rec.count())
			}
			c := rec.last()
			if c.Bot != b {
				t.Errorf("ctx.Bot = %p, want %p", c.Bot, b)
			}
			if c.Api != b.Api {
				t.Errorf("ctx.Api = %p, want %p", c.Api, b.Api)
			}
			if c.Update != &u {
				t.Errorf("ctx.Update = %p, want %p", c.Update, &u)
			}
		})
	}
}

func TestDispatchIgnoresUpdateWithoutMember(t *testing.T) {
	b := NewBot("TESTTOKEN")
	recs := make(map[UpdateType]*recorder)
	for _, typ := range updateTypes() {
		rec := &recorder{}
		recs[typ] = rec
		b.On(typ, rec.handle)
	}
	if err := b.dispatch(&api.Update{UpdateID: 1}); err != nil {
		t.Fatalf("dispatch: %v", err)
	}
	for typ, rec := range recs {
		if rec.count() != 0 {
			t.Errorf("%s handler called %d times, want 0", typ, rec.count())
		}
	}
}

func TestCommandRouting(t *testing.T) {
	textUpdate := func(s string) *api.Update {
		return &api.Update{UpdateID: 1, Message: &api.Message{Text: &s}}
	}
	photoUpdate := &api.Update{UpdateID: 1, Message: &api.Message{Photo: &[]api.PhotoSize{{}}}}

	tests := []struct {
		name            string
		update          *api.Update
		registerCommand bool
		wantCommand     int
		wantMessage     int
	}{
		{"plain command", textUpdate("/start"), true, 1, 0},
		{"command with arguments", textUpdate("/start arguments here"), true, 1, 0},
		{"command addressed to this bot", textUpdate("/start@TestBot"), true, 1, 0},
		{"non-command text", textUpdate("hello there"), true, 0, 1},
		{"unregistered command", textUpdate("/unknown"), true, 0, 1},
		{"command addressed to other bot", textUpdate("/start@OtherBot"), true, 0, 1},
		{"no command registered", textUpdate("/start"), false, 0, 1},
		{"non-text message", photoUpdate, true, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBot("TESTTOKEN")
			b.username = "TestBot"
			cmd := &recorder{}
			msg := &recorder{}
			if tt.registerCommand {
				b.Command("start", cmd.handle)
			}
			b.OnMessage(msg.handle)

			if err := b.dispatch(tt.update); err != nil {
				t.Fatalf("dispatch: %v", err)
			}
			if cmd.count() != tt.wantCommand {
				t.Errorf("command handler called %d times, want %d", cmd.count(), tt.wantCommand)
			}
			if msg.count() != tt.wantMessage {
				t.Errorf("message handler called %d times, want %d", msg.count(), tt.wantMessage)
			}
		})
	}
}

func TestRegistrationReplacesAndDeletes(t *testing.T) {
	t.Run("command replace", func(t *testing.T) {
		b := NewBot("TESTTOKEN")
		first, second := &recorder{}, &recorder{}
		b.Command("start", first.handle)
		b.Command("start", second.handle)
		b.OnMessage((&recorder{}).handle)
		if err := b.dispatch(textUpdate("/start")); err != nil {
			t.Fatalf("dispatch: %v", err)
		}
		if first.count() != 0 || second.count() != 1 {
			t.Errorf("first = %d, second = %d calls, want 0 and 1", first.count(), second.count())
		}
	})
	t.Run("command delete falls back to message handler", func(t *testing.T) {
		b := NewBot("TESTTOKEN")
		cmd, msg := &recorder{}, &recorder{}
		b.Command("start", cmd.handle)
		b.OnMessage(msg.handle)
		b.Command("start", nil)
		if err := b.dispatch(textUpdate("/start")); err != nil {
			t.Fatalf("dispatch: %v", err)
		}
		if cmd.count() != 0 || msg.count() != 1 {
			t.Errorf("command = %d, message = %d calls, want 0 and 1", cmd.count(), msg.count())
		}
	})
	t.Run("update type replace and delete", func(t *testing.T) {
		b := NewBot("TESTTOKEN")
		first, second := &recorder{}, &recorder{}
		b.On(UpdateCallbackQuery, first.handle)
		b.On(UpdateCallbackQuery, second.handle)
		u := &api.Update{UpdateID: 1, CallbackQuery: &api.CallbackQuery{}}
		if err := b.dispatch(u); err != nil {
			t.Fatalf("dispatch: %v", err)
		}
		if first.count() != 0 || second.count() != 1 {
			t.Fatalf("first = %d, second = %d calls, want 0 and 1", first.count(), second.count())
		}
		b.On(UpdateCallbackQuery, nil)
		if err := b.dispatch(u); err != nil {
			t.Fatalf("dispatch: %v", err)
		}
		if second.count() != 1 {
			t.Errorf("second called %d times after delete, want 1", second.count())
		}
	})
	t.Run("OnMessage delegates to On", func(t *testing.T) {
		b := NewBot("TESTTOKEN")
		msg := &recorder{}
		b.OnMessage(msg.handle)
		b.On(UpdateMessage, nil)
		if err := b.dispatch(textUpdate("hi")); err != nil {
			t.Fatalf("dispatch: %v", err)
		}
		if msg.count() != 0 {
			t.Errorf("message handler called %d times after delete, want 0", msg.count())
		}
	})
}

func textUpdate(s string) *api.Update {
	return &api.Update{UpdateID: 1, Message: &api.Message{Text: &s}}
}

// testServer answers getUpdates with successive replies; bodies of every
// request are appended to the returned slice.
func testServer(t *testing.T, replies ...string) (*httptest.Server, func() []string) {
	t.Helper()
	var mu sync.Mutex
	var bodies []string
	i := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/getMe") {
			_, _ = w.Write([]byte(`{"ok":true,"result":{"id":1,"is_bot":true,"first_name":"Test","username":"TestBot"}}`))
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		mu.Lock()
		bodies = append(bodies, string(body))
		reply := replies[min(i, len(replies)-1)]
		i++
		mu.Unlock()
		_, _ = w.Write([]byte(reply))
	}))
	t.Cleanup(srv.Close)
	return srv, func() []string {
		mu.Lock()
		defer mu.Unlock()
		return append([]string(nil), bodies...)
	}
}

func pollParams(t *testing.T, body string) (offset *int64, timeout *int) {
	t.Helper()
	var p struct {
		Offset  *int64 `json:"offset"`
		Timeout *int   `json:"timeout"`
	}
	if err := json.Unmarshal([]byte(body), &p); err != nil {
		t.Fatalf("decode request body %q: %v", body, err)
	}
	return p.Offset, p.Timeout
}

func TestStartPollsAdvancesOffsetAndCancels(t *testing.T) {
	const (
		withUpdate = `{"ok":true,"result":[{"update_id":10,"message":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"text":"hi"}}]}`
		empty      = `{"ok":true,"result":[]}`
	)
	srv, bodies := testServer(t, withUpdate, empty, empty, empty)
	b := newTestBot(t, srv)
	rec := &recorder{}
	b.OnMessage(rec.handle)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error, 1)
	go func() { errCh <- b.Start(ctx) }()

	// Wait until the second poll has been issued, then inspect both.
	deadline := time.Now().Add(5 * time.Second)
	for {
		if n := len(bodies()); n >= 2 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for second request; got %d requests", len(bodies()))
		}
		time.Sleep(5 * time.Millisecond)
	}

	logged := bodies()
	for i, body := range logged[:2] {
		_, timeout := pollParams(t, body)
		if timeout == nil || *timeout != pollTimeoutSeconds {
			t.Errorf("request %d timeout = %v, want %d", i+1, timeout, pollTimeoutSeconds)
		}
	}
	offset, _ := pollParams(t, logged[1])
	if offset == nil || *offset != 11 {
		t.Errorf("second request offset = %v, want 11", offset)
	}

	if rec.count() != 1 {
		t.Fatalf("message handler called %d times, want 1", rec.count())
	}
	if id := rec.last().Update.UpdateID; id != 10 {
		t.Errorf("handled update id = %d, want 10", id)
	}

	cancel()
	select {
	case err := <-errCh:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Start err = %v, want context.Canceled", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Start did not return after cancel")
	}
}

func TestStartRetriesTransient500(t *testing.T) {
	const (
		fiveHundred = `{"ok":false,"error_code":500,"description":"internal server error"}`
		withUpdate  = `{"ok":true,"result":[{"update_id":1,"message":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"text":"after retry"}}]}`
		empty       = `{"ok":true,"result":[]}`
	)
	srv, bodies := testServer(t, fiveHundred, withUpdate, empty, empty)
	b := newTestBot(t, srv)

	handled := make(chan error, 1)
	b.OnMessage(func(c *Context) error {
		handled <- nil
		return nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error, 1)
	go func() { errCh <- b.Start(ctx) }()

	select {
	case <-handled:
	case <-time.After(5 * time.Second):
		t.Fatal("update was not processed after transient failure")
	}
	cancel()
	select {
	case err := <-errCh:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("Start err = %v, want context.Canceled", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Start did not return after cancel")
	}
	if n := len(bodies()); n < 2 {
		t.Errorf("issued %d requests, want at least 2 (failure then retry)", n)
	}
}

func TestStartReturnsPermanent4xxWithoutRetry(t *testing.T) {
	const badRequest = `{"ok":false,"error_code":400,"description":"Bad Request: malformed query"}`
	srv, bodies := testServer(t, badRequest)
	b := newTestBot(t, srv)

	err := b.Start(context.Background())
	var apiErr *api.Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("Start err = %v, want *api.Error", err)
	}
	if apiErr.Code != 400 || apiErr.Description != "Bad Request: malformed query" {
		t.Errorf("api error = %d/%q, want 400/%q", apiErr.Code, apiErr.Description, "Bad Request: malformed query")
	}
	if n := len(bodies()); n != 1 {
		t.Errorf("issued %d requests, want 1 (no retry on permanent error)", n)
	}
}

func TestStartReturnsHandlerErrorUnchanged(t *testing.T) {
	const withCommand = `{"ok":true,"result":[{"update_id":1,"message":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"},"text":"/start@TestBot"}}]}`
	srv, bodies := testServer(t, withCommand, withCommand)
	b := newTestBot(t, srv)

	sentinel := errors.New("handler boom")
	b.Command("start", func(c *Context) error { return sentinel })

	if err := b.Start(context.Background()); !errors.Is(err, sentinel) {
		t.Fatalf("Start err = %v, want sentinel %v", err, sentinel)
	}
	if n := len(bodies()); n != 1 {
		t.Errorf("issued %d requests, want 1 (handler error stops polling)", n)
	}
}

func TestTransient(t *testing.T) {
	tests := []struct {
		err  error
		want bool
	}{
		{errors.New("connection refused"), true},
		{fmt.Errorf("telegram: decode response: %w", errors.New("unexpected EOF")), true},
		{&api.Error{Code: 429}, true},
		{&api.Error{Code: 500}, true},
		{&api.Error{Code: 502}, true},
		{&api.Error{Code: 400}, false},
		{&api.Error{Code: 401}, false},
		{&api.Error{Code: 403}, false},
		{&api.Error{Code: 404}, false},
		{&api.Error{Code: 409}, false},
	}
	for _, tt := range tests {
		if got := transient(tt.err); got != tt.want {
			t.Errorf("transient(%v) = %v, want %v", tt.err, got, tt.want)
		}
	}
}

func TestRetryDelayFor(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want time.Duration
	}{
		{"429 with retry after", &api.Error{Code: 429, Parameters: &api.ResponseParameters{RetryAfter: 9}}, 9 * time.Second},
		{"429 without parameters", &api.Error{Code: 429}, retryDelay},
		{"429 with zero retry after", &api.Error{Code: 429, Parameters: &api.ResponseParameters{}}, retryDelay},
		{"429 with negative retry after", &api.Error{Code: 429, Parameters: &api.ResponseParameters{RetryAfter: -3}}, retryDelay},
		{"5xx ignores retry after", &api.Error{Code: 500, Parameters: &api.ResponseParameters{RetryAfter: 9}}, retryDelay},
		{"network error", errors.New("dial tcp: refused"), retryDelay},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := retryDelayFor(tt.err); got != tt.want {
				t.Errorf("retryDelayFor(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}
