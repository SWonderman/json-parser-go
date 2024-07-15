package parser

import (
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
    errorHandler *ErrorHandler
	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := Parser{lexer: lexer, errorHandler: &ErrorHandler{}}

	parser.nextToken()
	parser.nextToken()

	return &parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.ReadToken()
}

func (parser *Parser) Parse() (*ParserResult, ParserErrors) {
	if parser.currentToken.Literal == token.LBRACE {
		result := parser.parseObject()

		return &ParserResult{SingleMap: result}, parser.errorHandler.GetErrors()
	} else if parser.currentToken.Type == token.LSQUARE_BRACE {
		result := parser.parseArray()

		// NOTE: is there a better way of converting []any to []map[string]any?
		var mapResult []map[string]any
		for _, res := range result {
			conv, ok := res.(map[string]any)
			if ok == false {
                parser.errorHandler.AddPlainError("Error while converting array of objects into a map. Conversion from 'any' type was not possible.")

				return nil, parser.errorHandler.GetErrors()
			}
			mapResult = append(mapResult, conv)
		}

		return &ParserResult{MapArray: mapResult}, parser.errorHandler.GetErrors()
	}

    parser.errorHandler.AddTokenError(fmt.Sprintf("The input has to begin either with '{' or with '[', but got '%s' instead.", parser.currentToken.Literal), &parser.currentToken)
    
    return nil, parser.errorHandler.GetErrors()
}

func (parser *Parser) parseJson() any {
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
		return false
	case token.TRUE:
		return true
	case token.NULL:
		return nil
	default:
		// Handle negative numbers
		if parser.currentToken.Type == token.MINUS && parser.peekExpected(token.NUMBER) {
			return parser.parseNegativeNumber()
		}

        parser.errorHandler.AddTokenError("Unknown token", &parser.currentToken)

        return nil
	}
}

func (parser *Parser) parseArray() []any {
	jsonArr := []any{}

	// consume '['
	parser.nextToken()

	for parser.currentToken.Type != token.RSQUARE_BRACE {
		parsedJson := parser.parseJson()

		jsonArr = append(jsonArr, parsedJson)

		// consume array value
		parser.nextToken()

		if parser.currentToken.Type != token.RSQUARE_BRACE {
			// consume ',' if we are not at the end of the object
			parser.nextToken()
		}
	}

	return jsonArr
}

func (parser *Parser) parseObject() map[string]any {
	jsonObj := make(map[string]any)

	// consume '{'
	parser.nextToken()

	for parser.currentToken.Type != token.RBRACE {
		if parser.currentToken.Type != token.STRING {

            parser.errorHandler.AddTokenError("Key value has to be of type string. Did you add quotation marks around the key value?", &parser.currentToken)

            return nil
		}

		key := parser.currentToken.Literal

		// move past the key string
		parser.nextToken()

		if parser.currentToken.Type != token.COLON {
            parser.errorHandler.AddTokenError("Key value has to be followed by a colon, but got " + string(parser.currentToken.Type), &parser.currentToken)
            
            return nil
		}

		// consume ':'
		parser.nextToken()

		value := parser.parseJson()

		jsonObj[key] = value

		// consume value
		parser.nextToken()

		if parser.currentToken.Type != token.RBRACE {
			// consume ',' if we are not at the end of the object
			parser.nextToken()
		}
	}

	return jsonObj
}

func (parser *Parser) parseString() string {
	return parser.currentToken.Literal
}

func (parser *Parser) parseNegativeNumber() interface{} {
	// consume '-'
	parser.nextToken()

	parsedValue := parser.parseNumber()

	switch val := parsedValue.(type) {
	case int:
		return -val
	case float64:
		return -val
	}

    return nil
}

func (parser *Parser) parseNumber() interface{} {
	// Try to parse the current token literal as an int first, if that fails,
	// try to parse it as a float.

	literal := parser.currentToken.Literal

	parsedInt, error := strconv.Atoi(literal)
	if error == nil {
		return parsedInt
	}

	parsedFloat, error := strconv.ParseFloat(literal, 64)
	if error == nil {
		return parsedFloat
	}

    parser.errorHandler.AddTokenError("It was not possible to parse the token literal as either int or float.", &parser.currentToken)

    return nil
}

func (parser *Parser) peekExpected(expectedToken token.TokenType) bool {
	if parser.peekToken.Type == expectedToken {
		return true
	}

	return false
}
