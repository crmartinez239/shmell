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
			token, tknErr := p.lexer.PeekRuneToken()

			if tknErr != nil {
				// Empty AST
				// lexer will only return system errors
				return tknErr
			}

			if token.Type != Bang {
				// Set new parser state
				continue
			}

			parseErr := p.parsePreprocessorStatement()

			if parseErr != nil {
				return parseErr
			}

			break //REMOVE
		}
	}
	return nil
}

func (p *Parser) parsePreprocessorStatement() error {
	// ![tag]
	// ![tag] [attribute] = [value]

	p.lexer.ReadRuneToken() // the Bang token
	tag, tagErr := p.lexer.ReadTagToken()

	if tagErr != nil {
		return tagErr
	}

	if tag.Tag == None {
		return &ParserError{"Expected: tag", tag}
	}

	peek, _ := p.lexer.PeekRuneToken()
	if peek.Type == EOL {
		// AST logic goes here
		p.lexer.ReadRuneToken()
		return nil
	}

	err := p.parseAttributes()
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseAttributes() error {
	// [attribute] ,
	// [attribute] : [value] ,
	for {
		attr, attrErr := p.lexer.ReadAttributeToken()
		if attrErr != nil {
			return attrErr
		}

		if attr.Type != Attribute {
			return &ParserError{"Expected: attribute", attr}
		}

		runeTkn, runeErr := p.lexer.ReadRuneToken()
		if runeErr != nil {
			return runeErr
		}

		// expect EOL, Comma, Colon
		if runeTkn.Type == EOL {
			// AST logic
			return nil
		}

		if runeTkn.Type == Comma {
			// AST logic
			continue
		}

		if runeTkn.Type != Colon {
			return &ParserError{"Expected: end of line", runeTkn}
		}

		p.getAttributeValue()

		break
	}
	return nil
}

func (p *Parser) getAttributeValue() (*Token, error) {
	// '[value]'
	// [value]
	peek, _ := p.lexer.PeekRuneToken()
	if peek.Type == Quote {

		value, err := p.lexer.ReadStringValueToken()

		if err != nil {
			if err.Error() == "mcq" {
				return nil, &ParserError{"Missing closing quote", nil}
			}
			return nil, err
		}

		return value, nil
	}

	// value, err := p.lexer.ReadSingleWordValue()

	return nil, nil
}
