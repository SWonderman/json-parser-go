package parser

import (
	"errors"
	"fmt"
	"strconv"

	"sw/json-parser/lexer"
	"sw/json-parser/token"
)

type ParserResult struct {
	SingleMap map[string]any
	MapArray  []map[string]any
}

func (parserResult *ParserResult) IsSingleMap() bool {
	return parserResult.SingleMap != nil
}

func (parserResult *ParserResult) IsMapArray() bool {
	return parserResult.MapArray != nil
}

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := Parser{lexer: lexer}

	parser.nextToken()
	parser.nextToken()

	return &parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.ReadToken()
}

func (parser *Parser) Parse() (*ParserResult, error) {
	if parser.currentToken.Literal == token.LBRACE {
		result, error := parser.parseObject()

		return &ParserResult{SingleMap: result}, error
	} else if parser.currentToken.Type == token.LSQUARE_BRACE {
		result, error := parser.parseArray()

        if error != nil {
            return nil, error
        }
        
        // NOTE: is there a better way of convering []any to []map[string]any?
        var mapResult []map[string]any
        for _, res := range result {
            conv, ok := res.(map[string]any)
            if ok == false {
                return nil, errors.New("Error while converting array of objects into a map. Conversion from 'any' type was not possible.")
            }
            mapResult = append(mapResult, conv) 
        }


		return &ParserResult{MapArray: mapResult}, error
	}

	return nil, errors.New(fmt.Sprintf("The input has to begin either with '{' or with '['. Found '%s' instead at line: %d and column: %d.", parser.currentToken.Literal, parser.currentToken.Line, parser.currentToken.Column))
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
	case token.FALSE:
		return false, nil
	case token.TRUE:
		return true, nil
	case token.NULL:
		return nil, nil
	default:
		// Handle negative numbers
		if parser.currentToken.Type == token.MINUS && parser.peekExpected(token.NUMBER) {
			return parser.parseNegativeNumber()
		}

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

func (parser *Parser) parseNegativeNumber() (interface{}, error) {
	// consume '-'
	parser.nextToken()

	parsedValue, error := parser.parseNumber()
	if error != nil {
		return nil, error
	}

	switch val := parsedValue.(type) {
	case int:
		return -val, nil
	case float64:
		return -val, nil
	}

	return nil, errors.New("It was not possible to parse a negative number")
}

func (parser *Parser) parseNumber() (interface{}, error) {
	// Try to parse the current token literal as an int first, if that fails,
	// try to parse it as a float.

	literal := parser.currentToken.Literal

	parsedInt, error := strconv.Atoi(literal)
	if error == nil {
		return parsedInt, nil
	}

	parsedFloat, error := strconv.ParseFloat(literal, 64)
	if error == nil {
		return parsedFloat, nil
	}

	return nil, errors.New("It was not possible to parse the current token literal as either int or float")
}

func (parser *Parser) peekExpected(expectedToken token.TokenType) bool {
	if parser.peekToken.Type == expectedToken {
		return true
	}

	return false
}
