package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, literal string) *Token {
	newToken := Token{Type: tokenType, Literal: literal}

	return &newToken
}

const (
	EoF = "EOF"

	STRING = "STRING"
	NUMBER = "NUMBER"

	COMMA = ","
	COLON = ":"

	LBRACE        = "{"
	RBRACE        = "}"
	LSQUARE_BRACE = "["
	RSQUARE_BRACE = "]"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)
