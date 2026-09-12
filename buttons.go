package gogram

import "github.com/infybtw/GoGramm/api"

// InlineButton creates an inline button that sends callbackData to the bot.
// An inline button must have exactly one action.
func InlineButton(text, callbackData string) api.InlineKeyboardButton {
	return api.InlineKeyboardButton{
		Text:         text,
		CallbackData: &callbackData,
	}
}

// InlineURLButton creates an inline button that opens url.
// An inline button must have exactly one action.
func InlineURLButton(text, url string) api.InlineKeyboardButton {
	return api.InlineKeyboardButton{
		Text: text,
		URL:  &url,
	}
}

// InlineKeyboard creates an inline keyboard from button rows.
func InlineKeyboard(rows ...[]api.InlineKeyboardButton) *api.InlineKeyboardMarkup {
	return &api.InlineKeyboardMarkup{InlineKeyboard: rows}
}

// ReplyButton creates a text button for a reply keyboard.
func ReplyButton(text string) api.KeyboardButton {
	return api.KeyboardButton{Text: text}
}

// ReplyKeyboardOptions configures a reply keyboard.
type ReplyKeyboardOptions struct {
	ResizeKeyboard  bool
	OneTimeKeyboard bool
	IsPersistent    bool
}

// ReplyKeyboard creates a reply keyboard from button rows and options.
func ReplyKeyboard(options ReplyKeyboardOptions, rows ...[]api.KeyboardButton) *api.ReplyKeyboardMarkup {
	return &api.ReplyKeyboardMarkup{
		Keyboard:        rows,
		ResizeKeyboard:  &options.ResizeKeyboard,
		OneTimeKeyboard: &options.OneTimeKeyboard,
		IsPersistent:    &options.IsPersistent,
	}
}
