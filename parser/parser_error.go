package parser

import (
	"fmt"
	"sw/json-parser/token"
)

type ParserErrors []string

type ErrorHandler struct {
    errors []string
}

const (
    ANSI_RESET = "\033[0m"
    ANSI_RED = "\033[31m"
)

func makeStringRed(message string) string {
    return fmt.Sprintf("%s%s%s", ANSI_RED, message, ANSI_RESET)
}

func (errorHandler *ErrorHandler) AddTokenError(errorMessage string, token *token.Token) {
    errorPosition := fmt.Sprintf("%s line %d and column %d near token literal '%s'.", makeStringRed("PARSER ERROR:"), token.Line, token.Column, token.Literal)
    error := fmt.Sprintf("%s\n%s", errorPosition, errorMessage)

    errorHandler.errors = append(errorHandler.errors, error)
}

func (errorHandler *ErrorHandler) AddPlainError(errorMessage string) {
    errorHandler.errors = append(errorHandler.errors, errorMessage)
}

func (errorHandler *ErrorHandler) GetErrors() []string {
    return errorHandler.errors
}

