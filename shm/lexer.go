package shm

import (
	"bufio"
	"log"
	"os"
	"unicode"
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

func isOperator(r rune) bool {
	str := string(r)
	switch str {
	case "(":
		return true
	case ")":
		return true
	case ";":
		return true
	case ":":
		return true
	case ",":
		return true
	case "\\":
		return true
	case ".":
		return true
	case "#":
		return true
	}
	return false
}

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

// Lexer undefined
type Lexer struct {
	file   *os.File
	reader *bufio.Reader
}

// NewLexer returns a new Lexer stream from fileName.
// If there is an error, it will be of type *PathError
func NewLexer(fileName string) (*Lexer, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	l := &Lexer{file, reader}
	return l, nil
}

// Close closes the lexer stream and frees memory.
// Should always be called when done with lexer
func (l *Lexer) Close() {
	l.file.Close()
}

// ReadRuneToken reads the next single rune token from file stream
func (l *Lexer) ReadRuneToken() *Token {
	for {
		currentRune, _, err := l.reader.ReadRune()

		if err != nil {
			log.Fatalln(err)
		}

		s := string(currentRune)
		if s == " " || s == "\t" {
			continue
		}

		return tokenFromRune(currentRune)
	}
}

// rune slice will start with a unicode character from category L
// the rest of the slice will be a combination of category L and N
func (l *Lexer) getLetterNumeric() ([]rune, error) {
	str := []rune{}

	// first rune has already been verified
	firstRune, _, _ := l.reader.ReadRune()
	str = append(str, firstRune)

	for {
		r, _, err := l.reader.ReadRune()

		if err != nil {
			return str, err
		}

		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			str = append(str, r)
		} else {
			break
		}
	}

	l.reader.UnreadRune()
	return str, nil
}
