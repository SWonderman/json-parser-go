package lexer

import (
	"testing"

	"sw/json-parser/token"
)

func TestLexerCanTokenizeSymbols(t *testing.T) {
	input := `{,:[]}`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.COMMA, ","},
		{token.COLON, ":"},
		{token.LSQUARE_BRACE, "["},
		{token.RSQUARE_BRACE, "]"},
		{token.RBRACE, "}"},
	}

	lexer := New(input)

	for i, exp := range expected {
		token := lexer.ReadToken()

		if token.Type != exp.expectedType {
			t.Fatalf("tests[%d] - tokentype is wrong. Expected=%q, but got=%q", i, exp.expectedType, token.Type)
		}

		if token.Literal != exp.expectedLiteral {
			t.Fatalf("tests[%d] - literal is wrong. Expected=%q, but got=%q", i, exp.expectedLiteral, token.Literal)
		}
	}
}

func TestLexerCanTokenizeStrings(t *testing.T) {
	input := `{"name":"Joe"}`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "name"},
		{token.COLON, ":"},
		{token.STRING, "Joe"},
		{token.RBRACE, "}"},
	}

	lexer := New(input)

	for i, exp := range expected {
		token := lexer.ReadToken()

		if token.Type != exp.expectedType {
			t.Fatalf("tests[%d] - tokentype is wrong. Expected=%q, but got=%q", i, exp.expectedType, token.Type)
		}

		if token.Literal != exp.expectedLiteral {
			t.Fatalf("tests[%d] - literal is wrong. Expected=%q, but got=%q", i, exp.expectedLiteral, token.Literal)
		}
	}
}

func TestLexerOmitWhitespace(t *testing.T) {
	input := `  { "name"  :   "Joe"  }   `

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "name"},
		{token.COLON, ":"},
		{token.STRING, "Joe"},
		{token.RBRACE, "}"},
	}

	lexer := New(input)

	for i, exp := range expected {
		token := lexer.ReadToken()

		if token.Type != exp.expectedType {
			t.Fatalf("tests[%d] - tokentype is wrong. Expected=%q, but got=%q", i, exp.expectedType, token.Type)
		}

		if token.Literal != exp.expectedLiteral {
			t.Fatalf("tests[%d] - literal is wrong. Expected=%q, but got=%q", i, exp.expectedLiteral, token.Literal)
		}
	}
}
