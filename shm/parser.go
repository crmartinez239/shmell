package shm

type ParserState int

const (
	Initial ParserState = iota
	Preprocessor
)

// Parser undefined
type Parser struct {
	lexer *Lexer,
	currentState ParserState
}

// NewParser undefined
func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer, Initial}
}

func (p *Parser) ParsePreprocessor() (error) {
	for {
		if p.currentState == Initial {
			
		}
	}
}
