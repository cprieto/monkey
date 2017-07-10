package lexer

import "github.com/cprieto/monkey/token"

var keywords = map[string]token.TokenType{
	"let": token.LET,
	"fn":  token.FUNC,
}

type Lexer struct {
	input      string
	char       byte
	currentPos int
	nextPos    int
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
	default:
		if isLetter(l.char) {
			tok.Literal = l.getIdent()
			tok.TokenType = lookupIdent(tok.Literal)
		} else {
			tok.TokenType = token.ILLEGAL
		}
	}

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

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func lookupIdent(input string) token.TokenType {
	if val, ok := keywords[input]; ok {
		return val
	}
	return token.IDENT
}
