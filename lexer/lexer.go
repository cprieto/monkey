package lexer

import (
	"github.com/cprieto/monkey/token"
	"strings"
)

var keywords = map[string]token.TokenType{
	"let":    token.LET,
	"fn":     token.FUNC,
	"return": token.RETURN,
	"if":     token.IF,
	"else":   token.ELSE,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"!=":     token.NE,
	"==":     token.EQ,
}

type Lexer struct {
	input      string
	char       byte
	currentPos int
	nextPos    int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case 0:
		tok = token.Token{Type: token.EOF}
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
		tok = token.Token{string(l.char), token.MINUS}
	case '/':
		tok = token.Token{string(l.char), token.SLASH}
	case '*':
		tok = token.Token{string(l.char), token.ASTERISK}
	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{string(char) + string(l.char), token.NE}
		} else {
			tok = token.Token{string(l.char), token.BANG}
		}
	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{string(char) + string(l.char), token.EQ}
		} else {
			tok = token.Token{string(l.char), token.ASSIGN}
		}
	case '>':
		tok = token.Token{string(l.char), token.GT}
	case '<':
		tok = token.Token{string(l.char), token.LT}
	default:
		if isLetter(l.char) {
			tok.Literal = l.getIdent()
			tok.Type = lookupIdent(tok.Literal)

			return tok // I really don't like this
		} else if isNumber(l.char) {
			tok.Literal = l.getNumber()
			tok.Type = token.NUMBER
			return tok
		}

		tok.Type = token.ILLEGAL
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.nextPos]
	}
	l.currentPos = l.nextPos
	l.nextPos += 1
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) getIdent() string {
	position := l.currentPos
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.currentPos]
}

func (l *Lexer) getNumber() string {
	position := l.currentPos
	for isNumber(l.char) {
		l.readChar()
	}
	return l.input[position:l.currentPos]
}

func (l *Lexer) peekChar() byte {
	if l.currentPos >= len(l.input) {
		return 0
	}
	return l.input[l.nextPos]
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func lookupIdent(input string) token.TokenType {
	key := strings.ToLower(input)
	if val, ok := keywords[key]; ok {
		return val
	}
	return token.IDENT
}
