package parser

import (
	"cprieto.com/monkey/ast"
	"cprieto.com/monkey/lexer"
	"cprieto.com/monkey/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lxr     *lexer.Lexer
	current token.Token
	peek    token.Token
	errors  []string

	prefixFn map[token.TokenType]prefixParseFn
	infixFn  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lxr: l, errors: []string{}}
	p.prefixFn = make(map[token.TokenType]prefixParseFn)
	p.nextToken()
	p.nextToken()

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lxr.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.current.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.current.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.current}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.current, Value: p.current.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Not doing expressions (yet)
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peek.Type == t
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.current.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token `%s` but got `%s`", t, p.peek.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.current}
	p.nextToken()

	// TODO: We don't support expressions
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixFn[t] = fn
}

func (p *Parser) registerInfix(t token.TokenType, fn infixParseFn) {
	p.infixFn[t] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.current}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(lowest int) ast.Expression {
	prefix := p.prefixFn[p.current.Type]
	if prefix == nil {
		p.noPrefixParseError(p.current.Type)
		return nil
	}
	leftExpr := prefix()
	return leftExpr
}

// / ** Prefix functions
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.current, Value: p.current.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{Token: p.current}
	value, err := strconv.ParseInt(p.current.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Couldn't parse %q as integer", p.current.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{Token: p.current, Operator: p.current.Literal}

	p.nextToken()
	exp.Right = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) noPrefixParseError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s", t)
	p.errors = append(p.errors, msg)
}
