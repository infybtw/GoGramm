package gogram

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infybtw/GoGramm/api"
)

func TestButtonConstructors(t *testing.T) {
	callback := InlineButton("Choose", "choice:1")
	if callback.Text != "Choose" || callback.CallbackData == nil || *callback.CallbackData != "choice:1" || callback.URL != nil {
		t.Errorf("callback button = %+v, want only text and callback data", callback)
	}

	url := InlineURLButton("Website", "https://example.com")
	if url.Text != "Website" || url.URL == nil || *url.URL != "https://example.com" || url.CallbackData != nil {
		t.Errorf("URL button = %+v, want only text and URL", url)
	}

	inline := InlineKeyboard([]api.InlineKeyboardButton{callback, url}, []api.InlineKeyboardButton{InlineButton("Back", "back")})
	if got := inline.InlineKeyboard; len(got) != 2 || len(got[0]) != 2 || got[0][1].Text != "Website" || got[1][0].Text != "Back" {
		t.Errorf("inline keyboard = %+v, want rows in supplied order", got)
	}

	reply := ReplyKeyboard(ReplyKeyboardOptions{ResizeKeyboard: true, OneTimeKeyboard: true, IsPersistent: true}, []api.KeyboardButton{ReplyButton("Yes"), ReplyButton("No")})
	if reply.Keyboard[0][0].Text != "Yes" || reply.Keyboard[0][1].Text != "No" {
		t.Errorf("reply keyboard = %+v, want text buttons in supplied order", reply.Keyboard)
	}
	if reply.ResizeKeyboard == nil || !*reply.ResizeKeyboard || reply.OneTimeKeyboard == nil || !*reply.OneTimeKeyboard || reply.IsPersistent == nil || !*reply.IsPersistent {
		t.Errorf("reply options = %+v, want all options enabled", reply)
	}

	var _ api.ReplyMarkup = inline
	var _ api.ReplyMarkup = reply
}

func TestSendMessageKeyboardJSON(t *testing.T) {
	tests := []struct {
		name   string
		markup api.ReplyMarkup
		check  func(*testing.T, map[string]any)
	}{
		{
			name: "inline",
			markup: InlineKeyboard([]api.InlineKeyboardButton{
				InlineButton("Choose", "choice:1"),
				InlineURLButton("Website", "https://example.com"),
			}),
			check: func(t *testing.T, markup map[string]any) {
				t.Helper()
				row := markup["inline_keyboard"].([]any)[0].([]any)
				callback := row[0].(map[string]any)
				if callback["text"] != "Choose" || callback["callback_data"] != "choice:1" {
					t.Errorf("callback button = %v", callback)
				}
				if _, ok := callback["url"]; ok {
					t.Errorf("callback button contains URL action: %v", callback)
				}
				url := row[1].(map[string]any)
				if url["text"] != "Website" || url["url"] != "https://example.com" {
					t.Errorf("URL button = %v", url)
				}
				if _, ok := url["callback_data"]; ok {
					t.Errorf("URL button contains callback action: %v", url)
				}
			},
		},
		{
			name: "reply",
			markup: ReplyKeyboard(ReplyKeyboardOptions{ResizeKeyboard: true}, []api.KeyboardButton{
				ReplyButton("Yes"), ReplyButton("No"),
			}),
			check: func(t *testing.T, markup map[string]any) {
				t.Helper()
				row := markup["keyboard"].([]any)[0].([]any)
				if row[0].(map[string]any)["text"] != "Yes" || row[1].(map[string]any)["text"] != "No" {
					t.Errorf("reply row = %v", row)
				}
				if markup["resize_keyboard"] != true || markup["one_time_keyboard"] != false || markup["is_persistent"] != false {
					t.Errorf("reply options = %v", markup)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var err error
				body, err = io.ReadAll(r.Body)
				if err != nil {
					t.Errorf("read request body: %v", err)
				}
				_, _ = w.Write([]byte(`{"ok":true,"result":{"message_id":1,"date":1,"chat":{"id":1,"type":"private"}}}`))
			}))
			t.Cleanup(srv.Close)

			client := api.New("TESTTOKEN", api.WithServerURL(srv.URL), api.WithHTTPClient(srv.Client()))
			_, err := client.SendMessage(context.Background(), &api.SendMessageParams{ChatID: api.NewChatID(1), Text: "Choose", ReplyMarkup: tt.markup})
			if err != nil {
				t.Fatalf("SendMessage: %v", err)
			}

			var request struct {
				ReplyMarkup map[string]any `json:"reply_markup"`
			}
			if err := json.Unmarshal(body, &request); err != nil {
				t.Fatalf("decode request: %v", err)
			}
			tt.check(t, request.ReplyMarkup)
		})
	}
}
