package parser

import (
	"cprieto.com/monkey/ast"
	"cprieto.com/monkey/lexer"
	"testing"
)

func TestParseLetIntegers(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 34;
`
	tests := []struct {
		literal string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatal("Expected a program and got nothing")
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if stmt.TokenLiteral() != "let" {
			t.Fatalf("Expected `let` literal but got `%s`", stmt.TokenLiteral())
		}

		letStmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("Expected a let statement but it is not")
		}

		if letStmt.Name != tt.literal {
			t.Fatalf("Expected literal with name `%s` but got `%s`", tt.literal, letStmt.Name)
		}
	}
}

func TestPeekErrors(t *testing.T) {
	input := `let x 5;`
	l := lexer.New(input)
	p := New(l)

	p.ParseProgram()
	if len(p.Errors()) == 0 {
		t.Fatalf("Expected a parsing error but got nothing")
	}

}

func TestReturnStatement(t *testing.T) {
	input := `return 5;`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("Expected a program and got nothing")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("Expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatal("Statement is not a return statement")
	}

	if stmt.TokenLiteral() != "return" {
		t.Fatalf("Expected a token literal `return` but got `%s`", stmt.TokenLiteral())
	}
}
