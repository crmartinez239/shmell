package shm

import (
	"testing"
)

func TestReadRuneToken(t *testing.T) {
	lexer, err := NewLexer("test/runetoken.shm")
	defer lexer.Close()

	if err != nil {
		t.Errorf("Could not create new lexer\nError: %q", err)
	}

	assert := func(t *testing.T, got, want TokenType) {
		t.Helper()
		if got != want {
			t.Errorf("got: %s - want: %s", got.String(), want.String())
		}
	}

	t.Run("Lparen token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := Lparen
		assert(t, got, want)
	})

	t.Run("Rparen token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := Rparen
		assert(t, got, want)
	})

	t.Run("Colon token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := Colon
		assert(t, got, want)
	})

	t.Run("Semi token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := Semi
		assert(t, got, want)
	})

	t.Run("Bslash token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := Bslash
		assert(t, got, want)
	})

	t.Run("EOL token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := EOL
		assert(t, got, want)
	})

	t.Run("Undefined token", func(t *testing.T) {
		got := lexer.ReadRuneToken().Type
		want := Undefined
		assert(t, got, want)
	})
}
