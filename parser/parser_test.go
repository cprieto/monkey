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

		if letStmt.Name.Value != tt.literal {
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

func TestIdentifierExpression(t *testing.T) {
	input := `foobar`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if len(program.Statements) == 0 {
		t.Fatalf("Expected some statement value, but got nothing")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected an expression statement and got something else: `%T`", program.Statements[0])
	}

	id, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Expecting an identifier as expression but got `%T`", stmt.Expression)
	}

	if id.Value != "foobar" {
		t.Fatalf("Identifier value is not `foobar` but `%s`", id.Value)
	}

	if id.TokenLiteral() != "foobar" {
		t.Fatalf("Identifier token is not `foobar` but `%s`", id.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := `10`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if len(program.Statements) == 0 {
		t.Fatalf("Expected some statement value, but got nothing")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected an expression statement and got something else: `%T`", program.Statements[0])
	}

	id, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Expecting an identifier as expression but got `%T`", stmt.Expression)
	}

	if id.Value != 10 {
		t.Fatalf("Identifier value is not `foobar` but `%d`", id.Value)
	}

	if id.TokenLiteral() != "10" {
		t.Fatalf("Identifier token is not `foobar` but `%s`", id.TokenLiteral())
	}
}
