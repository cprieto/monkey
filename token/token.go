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
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	BANG      = "!"
	GT        = ">"
	LT        = "<"
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
	IDENT     = "IDENT"
	LET       = "LET"
	FUNC      = "FUNCTION"
	NUMBER    = "NUMBER"
	RETURN    = "RETURN"
	IF        = "IF"
	ELSE      = "ELSE"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
)

type Token struct {
	Literal   string
	TokenType TokenType
}
