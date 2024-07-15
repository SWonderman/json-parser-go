package jsonparser

import (
	"sw/json-parser/lexer"
	"sw/json-parser/parser"
)

func Parse(input string) (*parser.ParserResult, parser.ParserErrors) {
	lexer := lexer.New(input)
	parser := parser.New(lexer)

	return parser.Parse()
}
