package ast

import (
	"bytes"
	"cprieto.com/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

/// ** Program statement

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

/// ** Identifier

type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i Identifier) String() string {
	return i.Value
}

/// ** Integer literals

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i IntegerLiteral) String() string {
	return i.Token.Literal
}

/// ** LET statements

type LetStatement struct {
	Token token.Token
	// TODO: use identifier
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.Value)
	out.WriteString(" = ")

	// TODO: Put value

	out.WriteString(";")

	return out.String()
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

/// ** RETURN statement

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (l *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " ")

	// TODO: Handle expression

	out.WriteString(";")

	return out.String()
}

func (l *ReturnStatement) TokenLiteral() string {
	return l.Token.Literal
}

/// ** Expression statement

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e ExpressionStatement) String() string {
	var out bytes.Buffer
	// TODO: Implement expression
	return out.String()
}
