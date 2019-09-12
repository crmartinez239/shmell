package shm

type ParserState int

const (
	Initial ParserState = iota
	Preprocessor
)

// Parser undefined
type Parser struct {
	lexer        *Lexer
	currentState ParserState
}

// NewParser undefined
func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer, Initial}
}

func (p *Parser) Parse() error {
	for {
		if p.currentState == Initial {
			// Look for preprocessor statements
			token := p.lexer.ReadToken()

			if token.Type == Bang {
				p.parsePreprocessorStatement()
			}

		}
	}
}

func (p *Parser) parsePreprocessorStatement() error {

}
