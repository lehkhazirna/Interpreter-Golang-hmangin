package token

type TokenType string

/*
golang struct ah chuan a inziak hnuhung hi a type ani.
chuvangin struct chhungah chuan
int a;  TokenType Type
		tih ang nilo in
a int;  Type TokenType
a inti ang.
*/
type Token struct {
	Type    TokenType
	Literal string
}

/*
Tunah chuan a TokenType hrang hrang, monkey language in a hman ho kha
const angin kan declare ang
*/
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
