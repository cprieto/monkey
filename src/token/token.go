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
)

type Token struct {
	Literal   string
	TokenType TokenType
}
