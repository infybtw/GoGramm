package gogram

import (
	"strings"
	"sync"
	"unicode"

	"github.com/infybtw/GoGramm/api"
)

// Composer groups related command handlers so they can be registered on a bot
// as one message handler.
type Composer struct {
	mu       sync.RWMutex
	commands map[string]Handler
}

// NewComposer returns an empty command composer.
func NewComposer() *Composer {
	return &Composer{commands: make(map[string]Handler)}
}

// Command registers handler for name, the command without its leading slash.
// Registration replaces any previous handler for the command; a nil handler
// removes it.
func (c *Composer) Command(name string, handler Handler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if handler == nil {
		delete(c.commands, name)
		return
	}
	c.commands[name] = handler
}

// Handle dispatches a matching command. Commands addressed to another bot and
// updates without a matching command are ignored.
func (c *Composer) Handle(ctx *Context) error {
	if c == nil || ctx == nil || ctx.Update == nil {
		return nil
	}
	name, ok := commandName(ctx.Update.Message, botUsername(ctx.Bot))
	if !ok {
		return nil
	}
	c.mu.RLock()
	handler := c.commands[name]
	c.mu.RUnlock()
	if handler == nil {
		return nil
	}
	return handler(ctx)
}

func botUsername(b *Bot) string {
	if b == nil {
		return ""
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.username
}

// commandName returns the command in msg when it is addressed to username.
// An empty username accepts only unqualified commands.
func commandName(msg *api.Message, username string) (string, bool) {
	if msg == nil || msg.Text == nil || !strings.HasPrefix(*msg.Text, "/") {
		return "", false
	}
	token := *msg.Text
	if i := strings.IndexFunc(token, unicode.IsSpace); i >= 0 {
		token = token[:i]
	}
	command := strings.TrimPrefix(token, "/")
	if name, target, addressed := strings.Cut(command, "@"); addressed {
		if username == "" || !strings.EqualFold(target, username) {
			return "", false
		}
		command = name
	}
	return command, true
}
