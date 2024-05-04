package lexer

import "unicode"

// example to lex the following SQL query: SELECT * FROM table WHERE column = 10

type TokenType int

const (
	ILLEGAL TokenType = iota

	Identifier // TODO: identifier is a SELECT, FROM, WHERE, AND, OR, JOIN, INNER JOIN, OUTER JOIN, LEFT JOIN, RIGHT JOIN, AS, ETC
	Number
	String

	BinaryOperator
	CloseParen
	OpenParen
	Comma
)

type BinaryOperatorType rune

const (
	BinaryOperatorTypeUnknown BinaryOperatorType = 0

	BinaryOperatorTypePlus BinaryOperatorType = '+'
	BinaryOperatorTypeMinus BinaryOperatorType = '-'
	BinaryOperatorTypeMultiply BinaryOperatorType = '*'
	BinaryOperatorTypeDivide BinaryOperatorType = '/'
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(source string) []Token {
	var tokens []Token

	for _, c := range source {
		// go through each character in the source and make tokens
		switch {
		case unicode.IsLetter(c):
			tokens = append(tokens, Token{String, string(c)})
		case unicode.IsDigit(c):
			tokens = append(tokens, Token{Number, string(c)})


		// operators
		case c == '(':
			tokens = append(tokens, Token{OpenParen, string(c)})
		case c == ')':
			tokens = append(tokens, Token{CloseParen, string(c)})
		case c == ',':
			tokens = append(tokens, Token{Comma, string(c)})
		case BinaryOperatorType(c) != BinaryOperatorTypeUnknown:
			tokens = append(tokens, Token{BinaryOperator, string(c)})

		default:
			tokens = append(tokens, Token{ILLEGAL, string(c)})
		}
	}

	return tokens

}
