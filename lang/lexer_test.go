package lang

import (
	"testing"
)

func TestReadTokenType(t *testing.T) {
	lexer, err := NewLexer("test.shm")
	if err != nil {
		t.Errorf("Could not create new lexer\nError: %q", err)
	}

	assert := func(t *testing.T, got, want TokenType) {
		t.Helper()
		if got != want {
			t.Errorf("got: %s - want: %s", got.String(), want.String())
		}
	}

	t.Run("Preprocessor (Bang) token", func(t *testing.T) {
		got := lexer.ReadToken().Type
		want := Bang
		assert(t, got, want)
	})

	t.Run("Identifier token", func(t *testing.T) {
		got := lexer.ReadToken().Type
		want := Ident
		assert(t, got, want)
	})

	t.Run("Semi token", func(t *testing.T) {
		got := lexer.ReadToken().Type
		want := Semi
		assert(t, got, want)
	})

	t.Run("EOL token", func(t *testing.T) {
		got := lexer.ReadToken().Type
		want := EOL
		assert(t, got, want)
	})
}
