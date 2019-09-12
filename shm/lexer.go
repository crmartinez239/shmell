package shm

import (
	"bufio"
	"log"
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

// func (l *Lexer) ReadTagToken() *Token {

// }

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
