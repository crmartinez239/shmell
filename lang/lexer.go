package lang

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
	// EOL end-of-line
	EOL
	// Undefined is undefined
	Undefined
)

// TypeString returns the string value for TokenType
func (t TokenType) String() string {
	switch t {
	case Ident:
		return "Ident"
	case Rparen:
		return "Rparen"
	case Lparen:
		return "Lparen"
	case Semi:
		return "Semi"
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

// Token token data structure
type Token struct {
	// Type the TokenType of token
	Type TokenType
	// Value is the raw token text
	Value []rune
}

func isNewLine(r rune) bool {
	if r == 10 {
		return true
	}

	return false
}

func isPreprocessor(r rune) bool {
	if string(r) == "!" {
		return true
	}
	return false
}

func isControlOperator(r rune) bool {
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

// ReadToken reads the next token from file stream
func (l *Lexer) ReadToken() *Token {
	currentRune, _, err := l.reader.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	if isPreprocessor(currentRune) {
		return &Token{Bang, []rune{currentRune}}
	}

	if unicode.IsLetter(currentRune) {
		l.reader.UnreadRune()
		word, _ := l.getLetterNumeric()
		return &Token{Ident, word}
	}

	if isControlOperator(currentRune) {
		return tokenFromRune(currentRune)
	}

	if isNewLine(currentRune) {
		return &Token{EOL, []rune{currentRune}}
	}

	return &Token{Undefined, []rune{}}
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
