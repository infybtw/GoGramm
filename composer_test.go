package gogram

import (
	"errors"
	"testing"

	"github.com/infybtw/GoGramm/api"
)

func TestComposerCommandDispatch(t *testing.T) {
	composer := NewComposer()
	called := 0
	composer.Command("start", func(ctx *Context) error {
		called++
		return nil
	})

	b := NewBot("TESTTOKEN")
	b.username = "TestBot"
	b.Use(composer)

	for _, text := range []string{"/start", "/start@TestBot arguments"} {
		if err := b.dispatch(textUpdate(text)); err != nil {
			t.Fatalf("dispatch %q: %v", text, err)
		}
	}
	if called != 2 {
		t.Errorf("command handler called %d times, want 2", called)
	}
}

func TestComposerIgnoresUnmatchedCommands(t *testing.T) {
	composer := NewComposer()
	called := 0
	composer.Command("start", func(*Context) error {
		called++
		return nil
	})

	b := NewBot("TESTTOKEN")
	b.username = "TestBot"
	b.Use(composer)
	for _, text := range []string{"hello", "/unknown", "/start@OtherBot"} {
		if err := b.dispatch(textUpdate(text)); err != nil {
			t.Fatalf("dispatch %q: %v", text, err)
		}
	}
	if called != 0 {
		t.Errorf("command handler called %d times, want 0", called)
	}
}

func TestComposerCommandReplaceDeleteAndError(t *testing.T) {
	composer := NewComposer()
	first, second := 0, 0
	composer.Command("start", func(*Context) error {
		first++
		return nil
	})
	composer.Command("start", func(*Context) error {
		second++
		return nil
	})

	b := NewBot("TESTTOKEN")
	b.Use(composer)
	if err := b.dispatch(textUpdate("/start")); err != nil {
		t.Fatalf("dispatch: %v", err)
	}
	if first != 0 || second != 1 {
		t.Errorf("handlers called %d and %d times, want 0 and 1", first, second)
	}

	composer.Command("start", nil)
	if err := b.dispatch(textUpdate("/start")); err != nil {
		t.Fatalf("dispatch after delete: %v", err)
	}
	if second != 1 {
		t.Errorf("handler called %d times after delete, want 1", second)
	}

	sentinel := errors.New("composer boom")
	composer.Command("fail", func(*Context) error { return sentinel })
	if err := b.dispatch(&api.Update{Message: &api.Message{Text: strptr("/fail")}}); !errors.Is(err, sentinel) {
		t.Errorf("dispatch error = %v, want %v", err, sentinel)
	}
}

func strptr(s string) *string { return &s }
