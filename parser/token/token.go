package token

import "fmt"

// Token character
const (
	PlusChar        = '+'
	MinusChar       = '-'
	StarChar        = '*'
	CommonSlashChar = '/'
)

// Type of token
const (
	Plus = 1 + iota
	Minus
	Star
	CommonSlash
	Number
)

var TokenVsTokenLiteral map[uint]string = map[uint]string{
	Plus:        "Plus",
	Minus:       "Minus",
	Star:        "Star",
	CommonSlash: "Common Slash",
	Number:      "Number",
}

//Token character stream of expression is token into token
type Token struct {
	Type    uint
	Literal interface{}
	Lexeme  string
	Column  uint
}

func (t Token) String() string {
	return fmt.Sprintf("< %s %s %v %d>", TokenVsTokenLiteral[t.Type], t.Lexeme, t.Literal, t.Column)
}