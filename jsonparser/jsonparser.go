package jsonparser

import (
	"sw/json-parser/lexer"
	"sw/json-parser/parser"
)

func Parse(input string) (map[string]any, error) {
    lexer := lexer.New(input)
    parser := parser.New(lexer)

    return parser.Parse()
}

