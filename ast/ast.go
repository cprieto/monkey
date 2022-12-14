package ast

import "cprieto.com/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	// TODO: use identifier
	Name  string
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (l *ReturnStatement) TokenLiteral() string {
	return l.Token.Literal
}
