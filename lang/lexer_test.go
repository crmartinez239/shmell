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
			t.Errorf("got: %q - want: %q", got, want)
		}
	}

	t.Run("Preprocessor token", func(t *testing.T) {
		got := lexer.ReadToken().Type
		want := Bang
		assert(t, got, want)
	})

}
