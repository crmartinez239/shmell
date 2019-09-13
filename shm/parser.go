package shm

// ParserError error interface
type ParserError struct {
	err   string
	token *Token
}

// Token that casued the parser error
func (e *ParserError) Token() *Token {
	return e.token
}

func (e *ParserError) Error() string {
	return e.err
}

type parserState int

const (
	preprocessor parserState = iota
)

// Parser undefined
type Parser struct {
	lexer        *Lexer
	currentState parserState
}

// NewParser undefined
func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer, preprocessor}
}

// Parse parses the current Lexer stream
func (p *Parser) Parse() error {
	for {

		if p.currentState == preprocessor {
			// Look for preprocessor statements
			token, tknErr := p.lexer.ReadRuneToken()

			if tknErr != nil {
				// lexer will only return system errors
				return tknErr
			}

			if token.Type != Bang {
				// Set new parser state
				p.lexer.UnreadRune()
				continue
			}

			parseErr := p.parsePreprocessorStatement()
		}

	}
}

func (p *Parser) parsePreprocessorStatement() error {
	tag, tagErr := p.lexer.ReadTagToken()

	if tag.Tag == None {
		return &ParserError{"Expected: tag", tag}
	}

	return nil
}
