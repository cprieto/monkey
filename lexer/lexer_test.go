package lexer

import (
	"github.com/cprieto/monkey/token"
	"testing"
)

func TestLexerCanRecognizeEOF(t *testing.T) {
	const input = ""

	l := NewLexer(input)
	tok := l.NextToken()
	if tok.TokenType != token.EOF {
		t.Fatalf("Expected EOF but got %q", tok.TokenType)
	}
}

func TestLexerCanRecognizeSymbolTokens(t *testing.T) {
	const input = ";(,}-{=)+"
	expect := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{";", token.SEMICOLON},
		{"(", token.LPAREN},
		{",", token.COMMA},
		{"}", token.RBRACE},
		{"-", token.LESS},
		{"{", token.LBRACE},
		{"=", token.ASSIGN},
		{")", token.RPAREN},
		{"+", token.PLUS},
		{"", token.EOF},
	}

	l := NewLexer(input)
	for _, tt := range expect {
		tok := l.NextToken()
		if tok.Literal != tt.Literal {
			t.Fatalf("Expected literal %q but got %q", tt.Literal, tok.Literal)
		}

		if tok.TokenType != tt.TokenType {
			t.Fatalf("Expected type %v but got %v", tt.TokenType, tok.TokenType)
		}
	}
}
