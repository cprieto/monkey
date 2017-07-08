package lexer

import "token"

type Lexer struct {
}

func NewLexer(input string) *Lexer {
	return &Lexer{}
}

func (l *Lexer) NextToken() *token.Token {
	return &token.Token{}
}
