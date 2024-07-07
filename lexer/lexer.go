package lexer

import (
	"sw/json-parser/token"
)

type Lexer struct {
	input       string
	position    int
	currentChar byte
}

func New(input string) *Lexer {
	l := Lexer{input: input}
	l.readChar()

	return &l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		// NOTE: 0 corresponds to a space character and will be used later to catch an EOF
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.position]
	}
	l.position += 1
}

func (l *Lexer) readJsonString() string {
	l.readChar()
	startPosition := l.position
	for l.currentChar != '"' {
		l.readChar()
	}
	endPosition := l.position

	return l.input[startPosition-1 : endPosition-1]
}

func (l *Lexer) ReadToken() token.Token {
	var newToken token.Token

	switch l.currentChar {
	case ',':
		newToken = *token.New(token.COMMA, string(l.currentChar))
	case ':':
		newToken = *token.New(token.COLON, string(l.currentChar))
	case '{':
		newToken = *token.New(token.LBRACE, string(l.currentChar))
	case '}':
		newToken = *token.New(token.RBRACE, string(l.currentChar))
	case '[':
		newToken = *token.New(token.LSQUARE_BRACE, string(l.currentChar))
	case ']':
		newToken = *token.New(token.RSQUARE_BRACE, string(l.currentChar))
	case '"':
		newToken = *token.New(token.STRING, l.readJsonString())
	case 0:
		newToken = *token.New(token.EoF, "")
	}

	l.readChar()

	return newToken
}
