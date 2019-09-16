package shm

import (
	"bufio"
	"errors"
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
			if err.Error() == "EOF" {
				return &Token{EOF, None, []rune{currentRune}}, nil
			}
			return nil, err
		}

		s := string(currentRune)
		if s == " " || s == "\t" {
			continue
		}

		return tokenFromRune(currentRune), nil
	}
}

// ReadTagToken attempts to read a Tag type token.
// Will return error if there was a problem reading from file.
func (l *Lexer) ReadTagToken() (*Token, error) {
	token, err := l.getLetterNumeric()

	if err != nil {
		if err.Error() == "EOF" {
			return &Token{EOF, None, token}, nil
		}
		return nil, err
	}

	return tokenFromTag(token), nil
}

// ReadAttributeToken reads the next Letter Numeric string and
// stores it in an Attribute token
func (l *Lexer) ReadAttributeToken() (*Token, error) {
	token, err := l.getLetterNumeric()

	if err != nil {
		if err.Error() == "EOF" {
			return &Token{EOF, None, token}, nil
		}
		return nil, err
	}

	return &Token{Attribute, None, token}, nil
}

// ReadStringValueToken reads all text withing quotes and returns  	 	the value in a token
func (l *Lexer) ReadStringValueToken() (*Token, error) {
	value := []rune{}
	l.ReadRuneToken() //fist quote
	for {
		nextRune, _, err := l.reader.ReadRune()

		if err != nil {
			if err.Error() == "EOF" {
				//Missing closing quote
				return nil, errors.New("mcq")
			}
			return nil, err
		}

		if string(nextRune) != "'" {
			value = append(value, nextRune)
		} else {
			break
		}
	}

	return &Token{Value, None, value}, nil
}

// ReadWordValueToken reads an attribute value not within quotes
func (l *Lexer) ReadWordValueToken() (*Token, error) {
	value := []rune{}
	for {
		nextRune, _, err := l.reader.ReadRune()

		if err != nil {
			if len(value) != 0 {
				return &Token{Value, None, value}, nil
			}
			if err.Error() == "EOF" {
				return &Token{EOF, None, nil}, nil
			}
			return nil, err
		}

		if isBreakRune(nextRune) {
			l.reader.UnreadRune()
			if len(value) == 0 {
				return nil, nil
			}
			return &Token{Value, None, value}, nil
		}

		value = append(value, nextRune)
	}
}

// PeekRuneToken attempts to read next signle rune token
// but does not remove it from lexer file stream
func (l *Lexer) PeekRuneToken() (*Token, error) {
	tkn, err := l.ReadRuneToken()
	l.reader.UnreadRune()
	return tkn, err
}

// rune slice will start with a unicode character from category L
// the rest of the slice will be a combination of category L and N
func (l *Lexer) getLetterNumeric() ([]rune, error) {
	str := []rune{}
	l.eatSpace()

	for {
		currentRune, _, err := l.reader.ReadRune()

		if err != nil {
			if len(str) == 0 {
				return nil, err
			}

			return str, nil
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

func isBreakRune(r rune) bool {
	if unicode.IsSpace(r) {
		return true
	}

	check := tokenFromRune(r)

	if check.Type == Comma {
		return true
	}

	if check.Type == EOL {
		return true
	}

	return false
}
