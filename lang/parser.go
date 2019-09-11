package lang

// Parser undefined
type Parser struct {
	lexer *Lexer
}

// NewParser undefined
func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer}
}
