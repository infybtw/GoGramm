# GoGramm

GoGramm is a typed Go client for the Telegram Bot API. It includes a small,
long-polling dispatcher with command routing, callback-query handlers,
composable command groups, keyboards, and helpers for responding to an update.

## Contents

- [Install](#install)
- [Quick Start](#quick-start)
- [Routing](#routing)
- [Reply And Edit Helpers](#reply-and-edit-helpers)
- [Buttons](#buttons)
- [Low-Level API](#low-level-api)
- [License](#license)

## Install

GoGramm requires Go 1.27 or later.

```bash
go get github.com/infybtw/GoGramm@latest
```

Create a bot and obtain its token from [@BotFather](https://t.me/BotFather),
then expose it as an environment variable:

```bash
export TELEGRAM_BOT_TOKEN="123456:token"
```

## Quick Start

This bot handles `/start`, sends an inline keyboard, and updates the original
message after the user presses a button.

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

	bot.Command("start", func(ctx *gogram.Context) error {
		return ctx.Reply("Choose an option", &api.SendMessageParams{
			ReplyMarkup: gogram.InlineKeyboard([]api.InlineKeyboardButton{
				gogram.InlineButton("Say hello", "hello"),
			}),
		})
	})

	bot.CallbackQuery("hello", func(ctx *gogram.Context) error {
		query := ctx.Update.CallbackQuery
		if err := ctx.Api.AnswerCallbackQuery(context.Background(), &api.AnswerCallbackQueryParams{
			CallbackQueryID: query.ID,
		}); err != nil {
			return err
		}
		return ctx.EditMessageText("Hello!")
	})

	if err := bot.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}
```

`Bot.Start` uses long polling and blocks until its context is canceled or a
handler returns an error. Updates and handlers run in order. It fetches the
bot username before polling, so `/start@YourBot` is routed correctly. Remove a
configured Telegram webhook before starting long polling.

## Routing

Register a command without its leading slash. A command-specific handler takes
precedence over the general `OnMessage` handler; a callback-specific handler
takes precedence over `On(gogram.UpdateCallbackQuery, handler)`.

```go
bot.OnMessage(func(ctx *gogram.Context) error {
	message := ctx.Update.Message
	if message.Text == nil {
		return nil
	}
	return ctx.Reply(*message.Text)
})

bot.On(gogram.UpdateChatJoinRequest, func(ctx *gogram.Context) error {
	// Handle any supported update type.
	return nil
})
```

Use a `Composer` to group commands, for example per feature or package. Pass
it to `Bot.Use`; direct `Bot.Command` handlers still take precedence.

```go
admin := gogram.NewComposer()
admin.Command("stats", func(ctx *gogram.Context) error {
	return ctx.Reply("42 active users")
})
admin.Command("reload", func(ctx *gogram.Context) error {
	return ctx.Reply("Configuration reloaded")
})

bot.Use(admin)
```

`Composer` ignores non-command messages, unknown commands, and commands
addressed to another bot. Registering the same command or callback again
replaces its handler; passing `nil` removes it.

## Sessions

Sessions are server-side state held in memory by `Bot`. They are keyed by both
the sender's user ID and chat ID, so different users in a group have separate
sessions. Register a session handler and activate it from a command; active
sessions receive subsequent messages without a registered command handler
before `OnMessage`:

```go
bot.Session("StartMessage", func(update *gogram.Context) error {
	defer update.EndSession()
	return update.Reply("Your next message was: " + *update.Update.Message.Text)
})

bot.Command("start", func(update *gogram.Context) error {
	if err := update.StartSession("StartMessage"); err != nil {
		return err
	}
	return update.Reply("Send one message.")
})
```

`Context.EndSession` ends the active session. Registered commands still take
precedence, allowing commands such as `/start` or `/cancel` to interrupt a
session. Sessions are not persisted and are lost when the process restarts.

## Reply And Edit Helpers

Handlers receive `*gogram.Context`, which exposes the bot, low-level API
client, and update as `Bot`, `Api`, and `Update` respectively.

`ctx.Reply` sends text to the chat of the incoming message. For a callback
query, it uses the chat of the message containing the button. It accepts an
optional `api.SendMessageParams` for send options, but always sets `ChatID` and
`Text` from the update and method argument. Set `ReplyParameters` yourself if
you need Telegram's message-reply behavior.

```go
parseMode := api.ParseModeHTML
return ctx.Reply("Saved", &api.SendMessageParams{
	ParseMode: &parseMode,
})
```

`ctx.ReplyWithMedia` sends an `api.InputMediaPhoto`, `InputMediaVideo`,
`InputMediaAudio`, `InputMediaDocument`, `InputMediaAnimation`,
`InputMediaLivePhoto`, or `InputMediaVoiceNote` to that same chat. Set its
`Caption` before sending.

```go
caption := "A photo"
return ctx.ReplyWithMedia(&api.InputMediaPhoto{
	Media:   api.FileID("AgACAgQAAxkBAAIB..."),
	Caption: &caption,
})
```

For message or callback-query updates, the edit helpers infer the target
message. Their optional parameter struct supplies edit options, while target
fields are ignored.

```go
return ctx.EditMessageText("Updated", &api.EditMessageTextParams{
	ReplyMarkup: gogram.InlineKeyboard([]api.InlineKeyboardButton{
		gogram.InlineButton("Back", "back"),
	}),
})

// Also available: EditMessageCaption, EditMessageMedia, EditMessageImage,
// and EditMessageReplyMarkup.
```

The reply helpers return `gogram.ErrNoReplyChat` when the update has no source
chat. The edit helpers return `gogram.ErrNoEditMessage` when no message can be
targeted.

## Buttons

Use the root-package helpers for common inline and reply keyboards. Each slice
passed to `InlineKeyboard` or `ReplyKeyboard` is a row.

```go
inline := gogram.InlineKeyboard(
	[]api.InlineKeyboardButton{
		gogram.InlineButton("Choose", "choice:one"),
		gogram.InlineURLButton("Website", "https://example.com"),
	},
	[]api.InlineKeyboardButton{
		gogram.InlineButton("Back", "back"),
	},
)

reply := gogram.ReplyKeyboard(gogram.ReplyKeyboardOptions{
	ResizeKeyboard: true,
}, []api.KeyboardButton{
	gogram.ReplyButton("Yes"),
	gogram.ReplyButton("No"),
})
```

Pass either keyboard as `ReplyMarkup` to `ctx.Reply` or a low-level method.
Always answer callback queries so Telegram stops showing its loading indicator.
For less common button actions, such as requesting a contact or location,
create `api.InlineKeyboardButton` or `api.KeyboardButton` directly. Remove a
reply keyboard with `&api.ReplyKeyboardRemove{RemoveKeyboard: true}`.

## Low-Level API

Use `api.Client` directly when you need a Bot API method without dispatcher
logic. Every method accepts a `context.Context` and a typed parameter struct;
all requests are made through `Client.Invoke`.

```go
client := api.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
_, err := client.SendMessage(context.Background(), &api.SendMessageParams{
	ChatID: api.NewChatID(123456789),
	Text:   "Hello from GoGramm",
})
```

See the exported types and methods in [`api`](./api) for the supported Bot API
surface.

## License

[MIT](LICENSE)
