# GoGramm

GoGramm is a typed Go client for the Telegram Bot API with a small
long-polling dispatcher for commands and updates.

## Install

```bash
go get github.com/infybtw/GoGramm@latest
```

Create a bot and obtain its token from [@BotFather](https://t.me/BotFather),
then expose it as an environment variable:

```bash
export TELEGRAM_BOT_TOKEN="123456:token"
```

## Quick Start

This bot replies to `/start` and echoes every other text message.

```go
package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	gogram "github.com/infybtw/GoGramm"
	"github.com/infybtw/GoGramm/api"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	bot := gogram.NewBot(os.Getenv("TELEGRAM_BOT_TOKEN"))

	bot.Command("start", func(update *gogram.Context) error {
		_, err := update.Api.SendMessage(context.Background(), &api.SendMessageParams{
			ChatID: api.NewChatID(update.Update.Message.Chat.ID),
			Text:   "Hello! Send me a message.",
		})
		return err
	})

	bot.OnMessage(func(update *gogram.Context) error {
		message := update.Update.Message
		if message.Text == nil {
			return nil
		}
		_, err := update.Api.SendMessage(context.Background(), &api.SendMessageParams{
			ChatID: api.NewChatID(message.Chat.ID),
			Text:   *message.Text,
		})
		return err
	})

	if err := bot.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}
```

`Bot.Start` uses long polling and blocks until its context is canceled or a
handler returns an error. It dispatches updates in order. Register a command
without its leading slash; a command handler takes precedence over `OnMessage`.
Remove a configured Telegram webhook before using long polling.

## Low-Level API

Use `api.Client` directly when you need a Bot API method without dispatcher
logic. All methods accept a `context.Context` and a typed parameter struct:

```go
client := api.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
_, err := client.SendMessage(context.Background(), &api.SendMessageParams{
	ChatID: api.NewChatID(123456789),
	Text:   "Hello from GoGramm",
})
```

See the exported types and methods in [`api`](./api) for the supported Bot API
surface.

## Buttons

Use the root package helpers for common inline and reply keyboards. The helpers
return `api.ReplyMarkup` implementations, so they can be passed directly to
`SendMessageParams.ReplyMarkup`.

```go
keyboard := gogram.InlineKeyboard(
	[]api.InlineKeyboardButton{
		gogram.InlineButton("Choose", "choice:one"),
		gogram.InlineURLButton("Website", "https://example.com"),
	},
)
_, err := bot.Api.SendMessage(ctx, &api.SendMessageParams{
	ChatID:      api.NewChatID(123456789),
	Text:        "Choose an option",
	ReplyMarkup: keyboard,
})

bot.CallbackQuery("choice:one", func(update *gogram.Context) error {
	query := update.Update.CallbackQuery
	log.Printf("choice: %s", *query.Data)
	return update.Api.AnswerCallbackQuery(ctx, &api.AnswerCallbackQueryParams{
		CallbackQueryID: query.ID,
	})
})
```

Always answer callback queries so Telegram stops showing the loading indicator.
`Bot.CallbackQuery` matches `callback_data` exactly; use
`Bot.On(gogram.UpdateCallbackQuery, handler)` to handle every callback query.

```go
keyboard := gogram.ReplyKeyboard(gogram.ReplyKeyboardOptions{
	ResizeKeyboard: true,
}, []api.KeyboardButton{
	gogram.ReplyButton("Yes"),
	gogram.ReplyButton("No"),
})
_, err := bot.Api.SendMessage(ctx, &api.SendMessageParams{
	ChatID:      api.NewChatID(123456789),
	Text:        "Continue?",
	ReplyMarkup: keyboard,
})

// Remove a reply keyboard with the existing API type.
removeKeyboard := &api.ReplyKeyboardRemove{RemoveKeyboard: true}
```

For uncommon button actions, such as requesting a contact or location, create
`api.InlineKeyboardButton` and `api.KeyboardButton` values directly.

## License

[MIT](LICENSE)
