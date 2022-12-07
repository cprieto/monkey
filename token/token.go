package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	THEN     = "THEN"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	EQ = "=="
	NE = "!="
)

type Token struct {
	Type    TokenType
	Literal string
}

var keyword = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"if":     IF,
	"then":   THEN,
	"else":   ELSE,
	"return": RETURN,
	"false":  FALSE,
	"true":   TRUE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return IDENT
}
