package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testToken = "TESTTOKEN"

// newTestClient spins an httptest server answering with reply and returns a
// Client pointed at it plus any recorded request.
func newTestClient(t *testing.T, reply string) (*Client, *capturedRequest) {
	t.Helper()
	cap := &capturedRequest{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}
		cap.Method = r.Method
		cap.Path = r.URL.Path
		cap.ContentType = r.Header.Get("Content-Type")
		cap.Body = body
		_, _ = w.Write([]byte(reply))
	}))
	t.Cleanup(srv.Close)
	return New(testToken, WithServerURL(srv.URL), WithHTTPClient(srv.Client())), cap
}

type capturedRequest struct {
	Method      string
	Path        string
	ContentType string
	Body        []byte
}

func TestGetMe(t *testing.T) {
	c, cap := newTestClient(t, `{"ok":true,"result":{"id":42,"is_bot":true,"first_name":"GoGram"}}`)

	u, err := c.GetMe(context.Background())
	if err != nil {
		t.Fatalf("GetMe: %v", err)
	}
	if cap.Method != http.MethodPost {
		t.Errorf("method = %q, want POST", cap.Method)
	}
	if cap.Path != "/bot"+testToken+"/getMe" {
		t.Errorf("path = %q, want /bot%s/getMe", cap.Path, testToken)
	}
	if cap.ContentType != "application/json" {
		t.Errorf("content type = %q, want application/json", cap.ContentType)
	}
	if len(cap.Body) != 0 {
		t.Errorf("body = %q, want empty", cap.Body)
	}
	if u.ID != 42 || u.FirstName != "GoGram" || !u.IsBot {
		t.Errorf("user = %+v, want ID 42, FirstName GoGram, IsBot true", u)
	}
}

func TestSendMessage(t *testing.T) {
	c, cap := newTestClient(t, `{"ok":true,"result":{"message_id":777,"date":1700000000,"chat":{"id":123456,"type":"private"}}}`)

	p := &SendMessageParams{
		ChatID: NewChatID(123456),
		Text:   "hi",
		ReplyMarkup: &InlineKeyboardMarkup{
			InlineKeyboard: [][]InlineKeyboardButton{{{Text: "btn"}}},
		},
	}
	m, err := c.SendMessage(context.Background(), p)
	if err != nil {
		t.Fatalf("SendMessage: %v", err)
	}
	if m == nil || m.MessageID != 777 {
		t.Fatalf("message = %+v, want MessageID 777", m)
	}
	if !strings.HasPrefix(cap.ContentType, "application/json") {
		t.Errorf("content type = %q, want application/json", cap.ContentType)
	}
	var body struct {
		ChatID      json.Number    `json:"chat_id"`
		Text        string         `json:"text"`
		ReplyMarkup map[string]any `json:"reply_markup"`
	}
	if err := json.Unmarshal(cap.Body, &body); err != nil {
		t.Fatalf("decode body %s: %v", cap.Body, err)
	}
	if body.ChatID.String() != "123456" {
		t.Errorf("chat_id = %s, want number 123456", body.ChatID.String())
	}
	if body.Text != "hi" {
		t.Errorf("text = %q, want hi", body.Text)
	}
	kb, _ := body.ReplyMarkup["inline_keyboard"].([]any)
	if len(kb) != 1 {
		t.Fatalf("inline_keyboard = %v, want 1 row", body.ReplyMarkup)
	}
	row, _ := kb[0].([]any)
	btn, _ := row[0].(map[string]any)
	if btn["text"] != "btn" {
		t.Errorf("button = %v, want text btn", btn)
	}
}

func TestSendPhotoUpload(t *testing.T) {
	c, cap := newTestClient(t, `{"ok":true,"result":{"message_id":778,"date":1700000000,"chat":{"id":1,"type":"private"}}}`)

	p := &SendPhotoParams{
		ChatID:  NewChatID(1),
		Photo:   FileUpload("cat.png", strings.NewReader("fakepng")),
		Caption: ptr("a cat"),
	}
	m, err := c.SendPhoto(context.Background(), p)
	if err != nil {
		t.Fatalf("SendPhoto: %v", err)
	}
	if m == nil || m.MessageID != 778 {
		t.Fatalf("message = %+v, want MessageID 778", m)
	}
	if !strings.HasPrefix(cap.ContentType, "multipart/form-data") {
		t.Fatalf("content type = %q, want multipart/form-data", cap.ContentType)
	}
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(cap.Body))
	r.Header.Set("Content-Type", cap.ContentType)
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		t.Fatalf("parse multipart: %v", err)
	}
	if got := r.FormValue("photo"); got != "attach://0" {
		t.Errorf("photo field = %q, want attach://0", got)
	}
	if got := r.FormValue("caption"); got != "a cat" {
		t.Errorf("caption field = %q, want a cat", got)
	}
	file, header, err := r.FormFile("0")
	if err != nil {
		t.Fatalf("file part 0: %v", err)
	}
	defer file.Close()
	if header.Filename != "cat.png" {
		t.Errorf("filename = %q, want cat.png", header.Filename)
	}
}

func ptr[T any](v T) *T { return &v }

func TestSendMediaGroupNestedUpload(t *testing.T) {
	c, cap := newTestClient(t, `{"ok":true,"result":[{"message_id":1,"date":1700000000,"chat":{"id":1,"type":"private"}}]}`)

	p := &SendMediaGroupParams{
		ChatID: NewChatID(1),
		Media: []InputMedia{
			&InputMediaPhoto{Media: FileUpload("a.png", strings.NewReader("fake"))},
		},
	}
	msgs, err := c.SendMediaGroup(context.Background(), p)
	if err != nil {
		t.Fatalf("SendMediaGroup: %v", err)
	}
	if len(msgs) != 1 || msgs[0].MessageID != 1 {
		t.Fatalf("messages = %+v, want one MessageID 1", msgs)
	}
	if !strings.HasPrefix(cap.ContentType, "multipart/form-data") {
		t.Fatalf("content type = %q, want multipart/form-data", cap.ContentType)
	}
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(cap.Body))
	r.Header.Set("Content-Type", cap.ContentType)
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		t.Fatalf("parse multipart: %v", err)
	}
	if media := r.FormValue("media"); !strings.Contains(media, `"attach://0"`) {
		t.Errorf("media field = %q, want attach://0 reference", media)
	}
	if _, _, err := r.FormFile("0"); err != nil {
		t.Errorf("file part 0: %v", err)
	}
}

func TestErrorPath(t *testing.T) {
	c, _ := newTestClient(t, `{"ok":false,"error_code":429,"description":"Too Many Requests: retry after 3","parameters":{"retry_after":3}}`)

	_, err := c.SendMessage(context.Background(), &SendMessageParams{ChatID: NewChatID(1), Text: "x"})
	if err == nil {
		t.Fatal("SendMessage: want error, got nil")
	}
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %T (%v), want *api.Error", err, err)
	}
	if apiErr.Code != 429 {
		t.Errorf("code = %d, want 429", apiErr.Code)
	}
	if apiErr.Parameters == nil || apiErr.Parameters.RetryAfter != 3 {
		t.Errorf("parameters = %+v, want RetryAfter 3", apiErr.Parameters)
	}
	if !strings.HasPrefix(apiErr.Error(), "telegram: 429:") {
		t.Errorf("Error() = %q, want prefix \"telegram: 429:\"", apiErr.Error())
	}
}

func TestEditMessageTextDualMode(t *testing.T) {
	t.Run("true result yields nil message", func(t *testing.T) {
		c, _ := newTestClient(t, `{"ok":true,"result":true}`)
		m, err := c.EditMessageText(context.Background(), &EditMessageTextParams{
			ChatID:    new(NewChatID(1)),
			MessageID: new(int64(2)),
			Text:      new("hi"),
		})
		if err != nil {
			t.Fatalf("EditMessageText: %v", err)
		}
		if m != nil {
			t.Fatalf("message = %+v, want nil", m)
		}
	})
	t.Run("message result decodes", func(t *testing.T) {
		c, _ := newTestClient(t, `{"ok":true,"result":{"message_id":9,"date":1700000000,"chat":{"id":1,"type":"private"}}}`)
		m, err := c.EditMessageText(context.Background(), &EditMessageTextParams{
			ChatID:    new(NewChatID(1)),
			MessageID: new(int64(2)),
			Text:      new("hi"),
		})
		if err != nil {
			t.Fatalf("EditMessageText: %v", err)
		}
		if m == nil || m.MessageID != 9 {
			t.Fatalf("message = %+v, want MessageID 9", m)
		}
	})
}

func TestExportChatInviteLink(t *testing.T) {
	c, _ := newTestClient(t, `{"ok":true,"result":"https://t.me/+abc"}`)

	link, err := c.ExportChatInviteLink(context.Background(), &ExportChatInviteLinkParams{ChatID: NewChatID(1)})
	if err != nil {
		t.Fatalf("ExportChatInviteLink: %v", err)
	}
	if link != "https://t.me/+abc" {
		t.Errorf("link = %q", link)
	}
}

func TestUnionResults(t *testing.T) {
	t.Run("chat member", func(t *testing.T) {
		c, _ := newTestClient(t, `{"ok":true,"result":{"status":"member","user":{"id":7,"is_bot":false,"first_name":"A"},"can_send_messages":true}}`)
		member, err := c.GetChatMember(context.Background(), &GetChatMemberParams{ChatID: NewChatID(1), UserID: 7})
		if err != nil {
			t.Fatalf("GetChatMember: %v", err)
		}
		if _, ok := member.(*ChatMemberMember); !ok {
			t.Fatalf("member = %T, want *ChatMemberMember", member)
		}
	})
	t.Run("menu button", func(t *testing.T) {
		c, _ := newTestClient(t, `{"ok":true,"result":{"type":"web_app","text":"Open","web_app":{"url":"https://example.test"}}}`)
		button, err := c.GetChatMenuButton(context.Background(), &GetChatMenuButtonParams{})
		if err != nil {
			t.Fatalf("GetChatMenuButton: %v", err)
		}
		webApp, ok := button.(*MenuButtonWebApp)
		if !ok || webApp.Text != "Open" {
			t.Fatalf("button = %#v, want web app button", button)
		}
	})
	t.Run("administrator list", func(t *testing.T) {
		c, _ := newTestClient(t, `{"ok":true,"result":[{"status":"creator","user":{"id":8,"is_bot":false,"first_name":"B"},"is_anonymous":false}]}`)
		members, err := c.GetChatAdministrators(context.Background(), &GetChatAdministratorsParams{ChatID: NewChatID(1)})
		if err != nil {
			t.Fatalf("GetChatAdministrators: %v", err)
		}
		if len(members) != 1 {
			t.Fatalf("members = %#v, want one member", members)
		}
		if _, ok := members[0].(*ChatMemberOwner); !ok {
			t.Fatalf("member = %T, want *ChatMemberOwner", members[0])
		}
	})
}
