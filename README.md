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

## License

[MIT](LICENSE)
