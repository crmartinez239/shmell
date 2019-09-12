package shm

// TokenType undefined
type TokenType int

const (
	// Ident identifier token type
	Ident TokenType = iota
	// Rparen right parenthesies character '('
	Rparen
	// Lparen left parenthesies character ')'
	Lparen
	// Colon colon character ':'
	Colon
	// Semi semi-colon character ';'
	Semi
	// Comma comma character ','
	Comma
	// Bslash back slash character
	Bslash
	// Dot period character
	Dot
	//Hash hash character
	Hash
	//Bang exclamtion character
	Bang
	// Equal equal sign character
	Equal
	// EOL end-of-line
	EOL
	// Undefined is undefined
	Undefined
)

// TypeString returns the string value for TokenType
func (t TokenType) String() string {
	switch t {
	case Ident:
		return "Identifier"
	case Rparen:
		return ")"
	case Lparen:
		return "("
	case Semi:
		return ";"
	case Colon:
		return ":"
	case Comma:
		return ","
	case Bslash:
		return "\\"
	case Dot:
		return "."
	case Hash:
		return "#"
	case Bang:
		return "!"
	case Equal:
		return "="
	}
	return "Undefined"
}

// Token token data structure
type Token struct {
	// Type the TokenType of token
	Type TokenType
	// Value is the raw token text
	Value []rune
}

// func tokenFromTag([]rune) *Token {

// }

func tokenFromRune(r rune) *Token {
	var t TokenType
	switch string(r) {
	case "(":
		t = Lparen
	case ")":
		t = Rparen
	case ";":
		t = Semi
	case ":":
		t = Colon
	case ",":
		t = Comma
	case "\\":
		t = Bslash
	case ".":
		t = Dot
	case "#":
		t = Hash
	case "!":
		t = Bang
	case "\n":
		t = EOL
	default:
		t = Undefined
	}
	return &Token{t, []rune{r}}
}
