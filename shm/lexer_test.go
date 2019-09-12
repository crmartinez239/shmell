package shm

import (
	"testing"
)

func TestSHM(t *testing.T) {
	lexer, err := NewLexer("test/test.shm")
	defer lexer.Close()

	if err != nil {
		t.Errorf("Could not create new lexer\nError: %q", err)
	}

	assertTokenType := func(t *testing.T, got, want TokenType) {
		t.Helper()
		if got != want {
			t.Errorf("got: %s - want: %s", got.String(), want.String())
		}
	}

	assertTagType := func(t *testing.T, got, want TagType) {
		t.Helper()
		if got != want {
			t.Errorf("got: %d - want: %d", got, want)
		}
	}

	t.Run("Bang token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := Bang
		assertTokenType(t, got.Type, want)
	})

	t.Run("Std tag", func(t *testing.T) {
		got, _ := lexer.ReadTagToken()
		want := Std
		assertTagType(t, got.Tag, want)
	})

	t.Run("EOL token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := EOL
		assertTokenType(t, got.Type, want)
	})
}

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
		got, _ := lexer.ReadRuneToken()
		want := Lparen
		assert(t, got.Type, want)
	})

	t.Run("Rparen token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := Rparen
		assert(t, got.Type, want)
	})

	t.Run("Colon token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := Colon
		assert(t, got.Type, want)
	})

	t.Run("Semi token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := Semi
		assert(t, got.Type, want)
	})

	t.Run("Bslash token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := Bslash
		assert(t, got.Type, want)
	})

	t.Run("EOL token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := EOL
		assert(t, got.Type, want)
	})

	t.Run("Undefined token", func(t *testing.T) {
		got, _ := lexer.ReadRuneToken()
		want := Undefined
		assert(t, got.Type, want)
	})
}
