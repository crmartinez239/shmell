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
			t.Errorf("got: %s - want: %s", typeToString(got), typeToString(want))
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
}

func typeToString(t TokenType) string {
	switch t {
	case Ident:
		return "Ident"
	case Rparen:
		return "Rparen"
	case Lparen:
		return "Lparen"
	case Semi:
		return "Semi-Colon"
	case Colon:
		return "Colon"
	case Comma:
		return "Comma"
	case Bslash:
		return "\\"
	case Dot:
		return "."
	case Hash:
		return "#"
	case Bang:
		return "!"
	}
	return "Undefined"
}
