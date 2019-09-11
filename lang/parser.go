package lang

type Parser struct {
	lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer}
}
