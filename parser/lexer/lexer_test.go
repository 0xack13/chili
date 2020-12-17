package lexer

import (
	"reflect"
	"testing"

	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

func TestLexerNormalExpression(t *testing.T) {
	expression := "34534 + 345.34 - 222 / 43435 * 745.234"
	lex := FromString(expression)
	number1, _ := decimal.NewFromString("34534")
	number2, _ := decimal.NewFromString("345.34")
	number3, _ := decimal.NewFromString("222")
	number4, _ := decimal.NewFromString("43435")
	number5, _ := decimal.NewFromString("745.234")

	tokens := []token.Token{
		{Type: token.Number{}, Literal: number1, Lexeme: "34534", Column: 5},
		{Type: token.Plus{}, Literal: nil, Lexeme: "+", Column: 7},
		{Type: token.Number{}, Literal: number2, Lexeme: "345.34", Column: 14},
		{Type: token.Minus{}, Literal: nil, Lexeme: "-", Column: 16},
		{Type: token.Number{}, Literal: number3, Lexeme: "222", Column: 20},
		{Type: token.CommonSlash{}, Literal: nil, Lexeme: "/", Column: 22},
		{Type: token.Number{}, Literal: number4, Lexeme: "43435", Column: 28},
		{Type: token.Star{}, Literal: nil, Lexeme: "*", Column: 30},
		{Type: token.Number{}, Literal: number5, Lexeme: "745.234", Column: 38},
	}
	for _, tt := range tokens {
		t.Run(tt.Lexeme, func(t *testing.T) {
			got, err := lex.Next()
			if err != nil {
				t.Errorf("Lexer.Next() error = %v, wantErr %v", err, false)
				return
			}
			if !reflect.DeepEqual(got, tt) {
				t.Errorf("Lexer.Next() = %v, want %v", got, tt)
			}
		})
	}
}
