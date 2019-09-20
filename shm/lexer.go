package shm

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

// Lexer undefined
type Lexer struct {
	file        *os.File
	reader      *bufio.Reader
	currentLine uint
	currentChar uint
}

// NewLexer returns a new Lexer stream from fileName.
// If there is an error, it will be of type *PathError
func NewLexer(fileName string) (*Lexer, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	l := &Lexer{file, reader, 1, 1}
	return l, nil
}

func (l *Lexer) newLine() {
	l.currentChar = 1
	l.currentLine++
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
		l.currentChar++

		if err != nil {
			return l.checkEOF(err)
		}

		s := string(currentRune)
		if s == " " || s == "\t" {
			continue
		}

		if s == "\n" {
			return l.eolToken(), nil
		}

		return tokenFromRune(currentRune, l.currentLine, l.currentChar), nil
	}
}

// ReadTagToken attempts to read a Tag type token.
// Will return error if there was a problem reading from file.
func (l *Lexer) ReadTagToken() (*Token, error) {
	tkn, err := l.getLetterNumeric()

	if err != nil {
		return l.checkEOF(err)
	}

	return tokenFromTag(tkn, l.currentLine, l.currentChar), nil
}

// ReadAttributeToken reads the next Letter Numeric string and
// stores it in an Attribute token
func (l *Lexer) ReadAttributeToken() (*Token, error) {
	tkn, err := l.getLetterNumeric()

	if err != nil {
		return l.checkEOF(err)
	}

	return &Token{Attribute, None, tkn, l.currentLine, l.currentChar}, nil
}

// ReadStringValueToken reads all text withing quotes and returns  	 	the value in a token
func (l *Lexer) ReadStringValueToken() (*Token, error) {
	value := []rune{}
	l.ReadRuneToken() //fist quote
	for {
		nextRune, _, err := l.reader.ReadRune()
		l.currentChar++
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

	return &Token{Value, None, value, l.currentLine, l.currentChar}, nil
}

// ReadWordValueToken reads an attribute value not within quotes
func (l *Lexer) ReadWordValueToken() (*Token, error) {
	value := []rune{}
	for {
		nextRune, _, err := l.reader.ReadRune()
		l.currentChar++
		if err != nil {
			if len(value) != 0 {
				return tokenFromValue(value, l.currentLine, l.currentChar), nil
			}
			if err.Error() == "EOF" {
				return l.eofToken(), nil
			}
			return nil, err
		}

		if l.isBreakRune(nextRune) {
			l.reader.UnreadRune()
			l.currentChar--
			if len(value) == 0 {
				return nil, nil
			}
			return tokenFromValue(value, l.currentLine, l.currentChar), nil
		}

		value = append(value, nextRune)
	}
}

// PeekRuneToken attempts to read next signle rune token
// but does not remove it from lexer file stream
func (l *Lexer) PeekRuneToken() (*Token, error) {
	cLine := l.currentLine
	cPos := l.currentChar

	tkn, err := l.ReadRuneToken()
	l.reader.UnreadRune()

	l.currentLine = cLine
	l.currentChar = cPos

	return tkn, err
}

// rune slice will start with a unicode character from category L
// the rest of the slice will be a combination of category L and N
func (l *Lexer) getLetterNumeric() ([]rune, error) {
	str := []rune{}
	l.eatSpace()

	for {
		currentRune, _, err := l.reader.ReadRune()
		l.currentChar++
		if err != nil {
			if len(str) == 0 {
				return nil, err
			}

			return str, nil
		}

		if unicode.IsLetter(currentRune) || unicode.IsNumber(currentRune) {
			str = append(str, currentRune)
		} else {
			l.reader.UnreadRune()
			l.currentChar--
			return str, nil
		}
	}

}

func (l *Lexer) eatSpace() {
	for {
		r, _, _ := l.reader.ReadRune()
		l.currentChar++
		space := string(r)

		if space != " " && space != "\t" {
			l.reader.UnreadRune()
			l.currentChar--
			return
		}
	}
}

func (l *Lexer) eolToken() *Token {
	nl := &Token{EOL, None, nil, l.currentLine, l.currentChar}
	l.newLine()
	return nl
}

func (l *Lexer) eofToken() *Token {
	return &Token{EOF, None, nil, l.currentLine, l.currentChar}
}

func (l *Lexer) checkEOF(err error) (*Token, error) {
	if err.Error() == "EOF" {
		return l.eofToken(), nil
	}
	return nil, err
}

func (l *Lexer) isBreakRune(r rune) bool {
	if unicode.IsSpace(r) {
		return true
	}

	check := tokenFromRune(r, l.currentLine, l.currentChar)

	if check.Type == Comma {
		return true
	}

	if check.Type == EOL {
		return true
	}

	return false
}
