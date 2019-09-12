package shm

import (
	"bufio"
	"os"
	"unicode"
)

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
func (l *Lexer) ReadRuneToken() (*Token, error) {
	for {
		currentRune, _, err := l.reader.ReadRune()

		if err != nil {
			return nil, err
		}

		s := string(currentRune)
		if s == " " || s == "\t" {
			continue
		}

		return tokenFromRune(currentRune), nil
	}
}

func (l *Lexer) ReadTagToken() (*Token, error) {
	token, err := l.getLetterNumeric()

	if err != nil {
		if token == nil {
			return nil, err
		}
	}

	return tokenFromTag(token), nil
}

// rune slice will start with a unicode character from category L
// the rest of the slice will be a combination of category L and N
func (l *Lexer) getLetterNumeric() ([]rune, error) {
	str := []rune{}

	l.eatSpace()
	for {
		currentRune, _, err := l.reader.ReadRune()

		if err != nil {
			return str, err
		}

		if unicode.IsLetter(currentRune) || unicode.IsNumber(currentRune) {
			str = append(str, currentRune)
		} else {
			break
		}
	}

	l.reader.UnreadRune()
	return str, nil
}

func (l *Lexer) eatSpace() {
	for {
		r, _, _ := l.reader.ReadRune()
		space := string(r)
		if space != " " && space != "\t" {
			l.reader.UnreadRune()
			return
		}
	}
}
