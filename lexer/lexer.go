package lexer

import (
	"slices"

	"sw/json-parser/token"
)

type ParseContext struct {
	Line   int
	Column int
}

func newParseContext() *ParseContext {
	return &ParseContext{Line: 1, Column: 0}
}

type Lexer struct {
	input       string
	position    int
	currentChar byte
	Context     *ParseContext
}

func New(input string) *Lexer {
	l := Lexer{input: input, Context: newParseContext()}
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

	l.Context.Column += 1
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

func (l *Lexer) eatWhitespace() {
	character := l.currentChar
	whitespaceChars := []byte{' ', '\t', '\n', '\r', '\v', '\f'}

	for slices.Contains(whitespaceChars, character) {
		if l.currentChar == '\n' {
			l.Context.Column = 0
			l.Context.Line += 1
		}

		l.readChar()
		character = l.currentChar
	}
}

func (l *Lexer) readKeyword() string {
	startPos := l.position
	for l.isCharLetter() {
		l.readChar()
	}
	endPos := l.position

	return l.input[startPos-1 : endPos-1]
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	// NOTE: this will also read and tokenize faulty 'numbers', like: 1.1.1
	for l.isCharDigit() || l.currentChar == '.' {
		l.readChar()
	}
	endPos := l.position

	return l.input[startPos-1 : endPos-1]
}

func (l *Lexer) isCharLetter() bool {
	ch := l.currentChar

	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) isCharDigit() bool {
	return '0' <= l.currentChar && '9' >= l.currentChar
}

func (l *Lexer) ReadToken() token.Token {
	var newToken token.Token

	l.eatWhitespace()

	switch l.currentChar {
	case ',':
		newToken = *token.New(token.COMMA, string(l.currentChar), l.Context.Line, l.Context.Column)
	case ':':
		newToken = *token.New(token.COLON, string(l.currentChar), l.Context.Line, l.Context.Column)
	case '-':
		newToken = *token.New(token.MINUS, string(l.currentChar), l.Context.Line, l.Context.Column)
	case '{':
		newToken = *token.New(token.LBRACE, string(l.currentChar), l.Context.Line, l.Context.Column)
	case '}':
		newToken = *token.New(token.RBRACE, string(l.currentChar), l.Context.Line, l.Context.Column)
	case '[':
		newToken = *token.New(token.LSQUARE_BRACE, string(l.currentChar), l.Context.Line, l.Context.Column)
	case ']':
		newToken = *token.New(token.RSQUARE_BRACE, string(l.currentChar), l.Context.Line, l.Context.Column)
	case '"':
		jsonString := l.readJsonString()
		beginningColumn := l.Context.Column - len(jsonString)
		newToken = *token.New(token.STRING, jsonString, l.Context.Line, beginningColumn)
	case 0:
		newToken = *token.New(token.EoF, "", l.Context.Line, l.Context.Column)
	default:
		if l.isCharLetter() {
			keyword := l.readKeyword()
			beginningColumn := l.Context.Column - len(keyword)
			newToken = *token.New(token.LookupKeyword(keyword), keyword, l.Context.Line, beginningColumn)

			return newToken
		} else if l.isCharDigit() {
			digit := l.readNumber()
			beginningColumn := l.Context.Column - len(digit)
			newToken = *token.New(token.NUMBER, digit, l.Context.Line, beginningColumn)

			return newToken
		}

		newToken = *token.New(token.INVALID, string(l.currentChar), l.Context.Line, l.Context.Column)
	}

	l.readChar()

	return newToken
}
