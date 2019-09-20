package shm

import "strings"

type TagType int

const (
	None TagType = iota
	Base
	Body
	HTML
	Link
	Meta
	Script
	Std
	Title
)

// TokenType undefined
type TokenType int

const (
	// Ident identifier token type
	Ident TokenType = iota
	// Rparen right parenthesies character '('
	Rparen
	// Lparen left parenthesies character ')'
	Lparen
	// Quote single quote character
	Quote
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
	// Bang exclamtion character
	Bang
	// Equal equal sign character
	Equal
	// Tag html tag
	Tag
	// Attribute tag attribute
	Attribute
	// Value tag or attribute value
	Value
	// EOL end-of-line
	EOL
	//EOF end-of-file
	EOF
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
	case Quote:
		return "'"
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
	case Tag:
		return "Tag"
	case Attribute:
		return "Attribute"
	case EOL:
		return "EOL"
	case EOF:
		return "EOF"
	default:
		return "Undefined"
	}
}

// Token token data structure
type Token struct {
	// Type the TokenType of token
	Type TokenType
	Tag  TagType
	// Value is the raw token text
	Value    []rune
	Line     uint
	Position uint
}

func tokenFromTag(r []rune, line uint, position uint) *Token {
	var t TagType
	str := strings.ToLower(string(r))
	switch str {
	case "base":
		t = Base
	case "body":
		t = Body
	case "html":
		t = HTML
	case "link":
		t = Link
	case "meta":
		t = Meta
	case "script":
		t = Script
	case "std":
		t = Std
	case "title":
		t = Title
	default:
		t = None
	}
	return &Token{Tag, t, r, line, position - uint(len(str))}
}

func tokenFromRune(r rune, line uint, position uint) *Token {
	var t TokenType
	switch string(r) {
	case "(":
		t = Lparen
	case ")":
		t = Rparen
	case "'":
		t = Quote
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
	return &Token{t, None, []rune{r}, line, position - 1}
}

func tokenFromValue(value []rune, line uint, position uint) *Token {
	return &Token{Value, None, value, line, position - uint(len(value))}
}
