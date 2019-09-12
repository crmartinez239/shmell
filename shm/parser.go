package shm

// ParserError error interface
type ParserError struct {
	err   string
	token Token
}

// Token that casued the parser error
func (e *ParserError) Token() *Token {
	return &e.token
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

func (p *Parser) Parse() error {
	for {
		if p.currentState == preprocessor {
			// Look for preprocessor statements
			token := p.lexer.ReadRuneToken()

			if token.Type == Bang {
				p.parsePreprocessorStatement()
			}

		}
	}
}

func (p *Parser) parsePreprocessorStatement() error {

}
