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
	const input = ";(,}-{=)+!<>*"
	expect := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{";", token.SEMICOLON},
		{"(", token.LPAREN},
		{",", token.COMMA},
		{"}", token.RBRACE},
		{"-", token.MINUS},
		{"{", token.LBRACE},
		{"=", token.ASSIGN},
		{")", token.RPAREN},
		{"+", token.PLUS},
		{"!", token.BANG},
		{"<", token.LT},
		{">", token.GT},
		{"*", token.ASTERISK},
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

func TestLexerIgnoresWhitespacesAndEnter(t *testing.T) {
	const input = `+ =
	{
	`
	expect := []token.TokenType{
		token.PLUS,
		token.ASSIGN,
		token.LBRACE,
		token.EOF,
	}

	l := NewLexer(input)
	for _, e := range expect {
		tok := l.NextToken()

		if tok.TokenType != e {
			t.Fatalf("Expected type %v but got %v", e, tok.TokenType)
		}
	}
}

func TestIllegalCharacters(t *testing.T) {
	const input = "~@"
	l := NewLexer(input)

	for i := 0; i < len(input); i++ {
		tok := l.NextToken()
		if tok.TokenType != token.ILLEGAL {
			t.Fatalf("Expected an ILLEGAL token, but got %v", tok.TokenType)
		}
	}

	tok := l.NextToken()
	if tok.TokenType != token.EOF {
		t.Fatalf("Expected EOF token but got %v", tok.TokenType)
	}
}

func TestLexerRecognizeIdentifierToken(t *testing.T) {
	const input = "foo bar _hello"
	expect := []struct {
		TokenType token.TokenType
		Literal   string
	}{
		{token.IDENT, "foo"},
		{token.IDENT, "bar"},
		{token.IDENT, "_hello"},
		{token.EOF, ""},
	}

	l := NewLexer(input)
	for _, r := range expect {
		tok := l.NextToken()
		if r.TokenType != tok.TokenType {
			t.Fatalf("Expected token %v but got %v", r.TokenType, tok.TokenType)
		}

		if r.Literal != tok.Literal {
			t.Fatalf("Expected literal %v but got %v", r.Literal, tok.Literal)
		}
	}
}

func TestLexerReturnsFunctionAndIdent(t *testing.T) {
	const input = "let fn else if true false return"
	expect := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{"let", token.LET},
		{"fn", token.FUNC},
		{"else", token.ELSE},
		{"if", token.IF},
		{"true", token.TRUE},
		{"false", token.FALSE},
		{"return", token.RETURN},
		{"", token.EOF},
	}

	l := NewLexer(input)
	for _, r := range expect {
		tok := l.NextToken()
		if r.TokenType != tok.TokenType {
			t.Fatalf("Expected token %v but got %v", r.TokenType, tok.TokenType)
		}

		if r.Literal != tok.Literal {
			t.Fatalf("Expected literal %v but got %v", r.Literal, tok.Literal)
		}
	}
}

func TestLexerReturnsFunctionWithTokenInMiddle(t *testing.T) {
	const input = "let=fn"
	expect := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{"let", token.LET},
		{"=", token.ASSIGN},
		{"fn", token.FUNC},
		{"", token.EOF},
	}

	l := NewLexer(input)
	for _, r := range expect {
		tok := l.NextToken()
		if r.TokenType != tok.TokenType {
			t.Fatalf("Expected token %v but got %v", r.TokenType, tok.TokenType)
		}

		if r.Literal != tok.Literal {
			t.Fatalf("Expected literal %v but got %v", r.Literal, tok.Literal)
		}
	}
}

func TestLexerUppercaseKeywordsAreRecognized(t *testing.T) {
	const input = "LET FN"
	expected := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{"LET", token.LET},
		{"FN", token.FUNC},
	}

	l := NewLexer(input)
	for _, r := range expected {
		tok := l.NextToken()
		if r.TokenType != tok.TokenType {
			t.Fatalf("Expected token %v but got %v", r.TokenType, tok.TokenType)
		}

		if r.Literal != tok.Literal {
			t.Fatalf("Expected literal %v but got %v", r.Literal, tok.Literal)
		}
	}
}

func TestLexerRecognizeFunctionAndLetAssignation(t *testing.T) {
	const input = "let add = fn(x, y){ x + y};"
	tokens := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{"let", token.LET},
		{"add", token.IDENT},
		{"=", token.ASSIGN},
		{"fn", token.FUNC},
		{"(", token.LPAREN},
		{"x", token.IDENT},
		{",", token.COMMA},
		{"y", token.IDENT},
		{")", token.RPAREN},
		{"{", token.LBRACE},
		{"x", token.IDENT},
		{"+", token.PLUS},
		{"y", token.IDENT},
		{"}", token.RBRACE},
		{";", token.SEMICOLON},
		{"", token.EOF},
	}

	l := NewLexer(input)
	for _, r := range tokens {
		tok := l.NextToken()
		if r.Literal != tok.Literal {
			t.Errorf("Expected literal %v but got %v", r.Literal, tok.Literal)
		}
		if r.TokenType != tok.TokenType {
			t.Errorf("Expected token %v but got %v", r.TokenType, tok.TokenType)
		}
	}
}

func TestLexerCanRecognizeNumbers(t *testing.T) {
	const input = "let a = 123;"
	tokens := []struct {
		Literal   string
		TokenType token.TokenType
	}{
		{"let", token.LET},
		{"a", token.IDENT},
		{"=", token.ASSIGN},
		{"123", token.NUMBER},
	}

	l := NewLexer(input)
	for _, r := range tokens {
		tok := l.NextToken()
		if r.Literal != tok.Literal {
			t.Errorf("Expected literal %v but got %v", r.Literal, tok.Literal)
		}
		if r.TokenType != tok.TokenType {
			t.Errorf("Expected token %v but got %v", r.TokenType, tok.TokenType)
		}
	}
}
