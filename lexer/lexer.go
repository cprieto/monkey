package lexer

import "cprieto.com/monkey/token"

type Lexer struct {
	input    string
	position int
	current  int
	char     byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // feed the first reading character

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '!':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			tok = token.Token{Type: token.NE, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.current >= len(l.input) {
		l.char = 0 // set char to NUL
	} else {
		l.char = l.input[l.current]
	}
	l.position = l.current
	l.current += 1
}

func (l *Lexer) peekChar() byte {
	if l.current >= len(l.input) {
		return '0'
	} else {
		return l.input[l.current]
	}
}

// read a given identifier
func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// create a token from a given type and byte
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// check if a byte is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
