package parser

import (
	"errors"
	"strconv"

	"sw/json-parser/lexer"
	"sw/json-parser/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := Parser{lexer: lexer}

	parser.nextToken()

	return &parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.lexer.ReadToken()
}

func (parser *Parser) Parse() (map[string]any, error) {
	// TODO: check if the first token is '{'
	return parser.parseObject()
}

func (parser *Parser) parseJson() (any, error) {
	switch parser.currentToken.Type {
	case token.LBRACE:
		return parser.parseObject()
	case token.LSQUARE_BRACE:
		return parser.parseArray()
	case token.STRING:
		return parser.parseString()
	case token.NUMBER:
		return parser.parseNumber()
	default:
		return nil, errors.New("Unknown current token with type: " + string(parser.currentToken.Type) + " and value: " + string(parser.currentToken.Literal))
	}
}

func (parser *Parser) parseArray() ([]any, error) {
	jsonArr := []any{}

	// consume '['
	parser.nextToken()

	for parser.currentToken.Type != token.RSQUARE_BRACE {
		parsedJson, err := parser.parseJson()
		if err != nil {
			return nil, err
		}

		jsonArr = append(jsonArr, parsedJson)

		// consume array value
		parser.nextToken()

		if parser.currentToken.Type != token.RSQUARE_BRACE {
			// consume ',' if we are not at the end of the object
			parser.nextToken()
		}
	}

	return jsonArr, nil
}

func (parser *Parser) parseObject() (map[string]any, error) {
	jsonObj := make(map[string]any)

	// consume '{'
	parser.nextToken()

	for parser.currentToken.Type != token.RBRACE {
		if parser.currentToken.Type != token.STRING {
			return nil, errors.New("Key has to be of type string, got: " + string(parser.currentToken.Type))
		}

		key := parser.currentToken.Literal

		// move past the key string
		parser.nextToken()

		if parser.currentToken.Type != token.COLON {
			return nil, errors.New("Key has to be followed by a colon, got: " + string(parser.currentToken.Type))
		}

		// consume ':'
		parser.nextToken()

		value, err := parser.parseJson()
		if err != nil {
			return nil, err
		}

		jsonObj[key] = value

		// consume value
		parser.nextToken()

		if parser.currentToken.Type != token.RBRACE {
			// consume ',' if we are not at the end of the object
			parser.nextToken()
		}
	}

	return jsonObj, nil
}

func (parser *Parser) parseString() (string, error) {
	return parser.currentToken.Literal, nil
}

func (parser *Parser) parseNumber() (int, error) {
	// TODO: make it work with floats
	return strconv.Atoi(parser.currentToken.Literal)
}
