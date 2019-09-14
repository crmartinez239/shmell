package shm

import (
	"testing"
)

func TestParser(t *testing.T) {
	lexer, err := NewLexer("shm/test/test.shm")
	parser := NewParser(lexer)

	if err != nil {
		t.Errorf("Could not create new lexer.\nError: %q", err)
	}

	assertError := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %s - want: %s", got, want)
		}
	}

	t.Run("Parser error", func(t *testing.T) {
		got := parser.Parse().Error()
		want := ""
		assertError(t, got, want)
	})

}
