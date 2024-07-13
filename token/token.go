package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column int
}

func New(tokenType TokenType, literal string, line int, column int) *Token {
	return &Token{Type: tokenType, Literal: literal, Line: line, Column: column}
}

const (
	INVALID = "INVALID"
	EoF     = "EOF"

	STRING = "STRING"
	NUMBER = "NUMBER"

	COMMA = ","
	COLON = ":"
	MINUS = "-"

	LBRACE        = "{"
	RBRACE        = "}"
	LSQUARE_BRACE = "["
	RSQUARE_BRACE = "]"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func LookupKeyword(keyword string) TokenType {
	value, wasFound := keywords[keyword]

	if wasFound == false {
		return INVALID
	}

	return value
}
