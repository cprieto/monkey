package lexer

import "github.com/cprieto/monkey/token"

type Lexer struct {
	input    string
	char     byte
	position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.readChar()
	switch l.char {
	case 0:
		tok = token.Token{TokenType: token.EOF}
	case ';':
		tok = token.Token{Literal: string(l.char), TokenType: token.SEMICOLON}
	}

	return tok
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.position]
	}
	l.position += 1
}
