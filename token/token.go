package token

type TokenType string

const (
	LBRACE    = "{"
	RBRACE    = "}"
	LPAREN    = "("
	RPAREN    = ")"
	COMMA     = ","
	SEMICOLON = ";"
	ASSIGN    = "="
	PLUS      = "+"
	LESS      = "-"
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
	IDENT     = "IDENT"
	LET       = "LET"
	FUNC      = "FUNCTION"
)

type Token struct {
	Literal   string
	TokenType TokenType
}
