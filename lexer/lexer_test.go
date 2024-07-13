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
        expectedLine int
        expectedColumn int
	}{
		{token.LBRACE, "{", 1, 1},
		{token.COMMA, ",", 1, 2},
		{token.COLON, ":", 1, 3},
		{token.LSQUARE_BRACE, "[", 1, 4},
		{token.RSQUARE_BRACE, "]", 1, 5},
		{token.RBRACE, "}", 1, 6},
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

        if token.Line != exp.expectedLine {
            t.Fatalf("tests[%d] - line is wrong. Expected=%d, but got=%d", i, exp.expectedLine, token.Line)
        }

        if token.Column != exp.expectedColumn {
            t.Fatalf("tests[%d] - column is wrong. Expected=%d, but got=%d", i, exp.expectedColumn, token.Column)
        }
	}
}

func TestLexerCanTokenizeStrings(t *testing.T) {
	input := `{"name":"Joe"}`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
        expectedLine int
        expectedColumn int
	}{
		{token.LBRACE, "{", 1, 1},
		{token.STRING, "name", 1, 3},
		{token.COLON, ":", 1, 8},
		{token.STRING, "Joe", 1, 10},
		{token.RBRACE, "}", 1, 14},
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

        if token.Line != exp.expectedLine {
            t.Fatalf("tests[%d] - line is wrong. Expected=%d, but got=%d", i, exp.expectedLine, token.Line)
        }

        if token.Column != exp.expectedColumn {
            t.Fatalf("tests[%d] - column is wrong. Expected=%d, but got=%d", i, exp.expectedColumn, token.Column)
        }
	}
}

func TestLexerOmitWhitespace(t *testing.T) {
	input := `  { "name"  :   "Joe"  }   `

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
        expectedLine int
        expectedColumn int
	}{
		{token.LBRACE, "{", 1, 3},
		{token.STRING, "name",1 , 6},
		{token.COLON, ":", 1, 13},
		{token.STRING, "Joe", 1, 18},
		{token.RBRACE, "}", 1, 24},
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

        if token.Line != exp.expectedLine {
            t.Fatalf("tests[%d] - line is wrong. Expected=%d, but got=%d", i, exp.expectedLine, token.Line)
        }

        if token.Column != exp.expectedColumn {
            t.Fatalf("tests[%d] - column is wrong. Expected=%d, but got=%d", i, exp.expectedColumn, token.Column)
        }
	}
}

func TestLexerTokenizesKeywords(t *testing.T) {
	input := `{"isBold": false, "likesProgramming": true, "hasKids": null}`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
        expectedLine int
        expectedColumn int
	}{
		{token.LBRACE, "{", 1, 1},
		{token.STRING, "isBold", 1, 3},
		{token.COLON, ":", 1, 10},
		{token.FALSE, "false", 1, 12},
		{token.COMMA, ",", 1, 17},
		{token.STRING, "likesProgramming", 1, 20},
		{token.COLON, ":", 1, 37},
		{token.TRUE, "true", 1, 39},
		{token.COMMA, ",", 1, 43},
		{token.STRING, "hasKids", 1, 46},
		{token.COLON, ":", 1, 54},
		{token.NULL, "null", 1, 56},
		{token.RBRACE, "}", 1, 60},
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

        if token.Line != exp.expectedLine {
            t.Fatalf("tests[%d] - line is wrong. Expected=%d, but got=%d", i, exp.expectedLine, token.Line)
        }

        if token.Column != exp.expectedColumn {
            t.Fatalf("tests[%d] - column is wrong. Expected=%d, but got=%d", i, exp.expectedColumn, token.Column)
        }
	}
}

func TestLexerTokenizesNumbers(t *testing.T) {
	input := `{"name": "Joe", "age": 77, "salary": 123.123, "cars": -1}`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
        expectedLine int
        expectedColumn int
	}{
		{token.LBRACE, "{", 1, 1},
		{token.STRING, "name", 1, 3},
		{token.COLON, ":", 1, 8},
		{token.STRING, "Joe", 1, 11},
		{token.COMMA, ",", 1 , 15},
		{token.STRING, "age", 1, 18},
		{token.COLON, ":", 1, 22},
		{token.NUMBER, "77", 1, 24},
		{token.COMMA, ",", 1, 26},
		{token.STRING, "salary", 1, 29},
		{token.COLON, ":", 1, 36},
		{token.NUMBER, "123.123", 1, 38},
		{token.COMMA, ",", 1, 45},
		{token.STRING, "cars", 1, 48},
		{token.COLON, ":", 1, 53},
		{token.MINUS, "-", 1, 55},
		{token.NUMBER, "1", 1, 56},
		{token.RBRACE, "}", 1, 57},
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

        if token.Line != exp.expectedLine {
            t.Fatalf("tests[%d] - line is wrong. Expected=%d, but got=%d", i, exp.expectedLine, token.Line)
        }

        if token.Column != exp.expectedColumn {
            t.Fatalf("tests[%d] - column is wrong. Expected=%d, but got=%d", i, exp.expectedColumn, token.Column)
        }
	}
}

func TestLexerContextTracksTokenPosition(t *testing.T) {
    input := `{
    "first_name": "Joe",
    "last_name": "Doe"
}`

	expected := []struct {
		expectedType    token.TokenType
		expectedLiteral string
        expectedLine int
        expectedColumn int
	}{
		{token.LBRACE, "{", 1, 1},
		{token.STRING, "first_name", 2, 6},
		{token.COLON, ":", 2, 17},
		{token.STRING, "Joe", 2, 20},
		{token.COMMA, ",", 2 , 24},
		{token.STRING, "last_name", 3, 6},
		{token.COLON, ":", 3, 16},
		{token.STRING, "Doe", 3, 19},
		{token.RBRACE, "}", 4, 1},
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

        if token.Line != exp.expectedLine {
            t.Fatalf("tests[%d] - line is wrong. Expected=%d, but got=%d", i, exp.expectedLine, token.Line)
        }

        if token.Column != exp.expectedColumn {
            t.Fatalf("tests[%d] - column is wrong. Expected=%d, but got=%d", i, exp.expectedColumn, token.Column)
        }
	}
}
