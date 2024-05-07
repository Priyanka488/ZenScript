package tokens

const TT_INT = "INT"
const TT_FLOAT = "FLOAT"
const TT_PLUS = "PLUS"
const TT_MINUS = "MINUS"
const TT_MUL = "MUL"
const TT_DIV = "DIV"
const TT_LPAREN = "LPAREN"
const TT_RPAREN = "RPAREN"
const TT_SPACE = "SPACE"

type Token struct {
	Ttype string
	Value string
}

func New(ttype string, value string) *Token {
	token := &Token{
		Ttype: ttype,
		Value: value,
	}
	return token
}
