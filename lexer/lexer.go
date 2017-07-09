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
	l.skipWhitespace()

	switch l.char {
	case 0:
		tok = token.Token{TokenType: token.EOF}
	case ';':
		tok = token.Token{string(l.char), token.SEMICOLON}
	case ',':
		tok = token.Token{string(l.char), token.COMMA}
	case '(':
		tok = token.Token{string(l.char), token.LPAREN}
	case ')':
		tok = token.Token{string(l.char), token.RPAREN}
	case '{':
		tok = token.Token{string(l.char), token.LBRACE}
	case '}':
		tok = token.Token{string(l.char), token.RBRACE}
	case '+':
		tok = token.Token{string(l.char), token.PLUS}
	case '-':
		tok = token.Token{string(l.char), token.LESS}
	case '=':
		tok = token.Token{string(l.char), token.ASSIGN}
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

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}
