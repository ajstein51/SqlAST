package lexer

import (
	"unicode"
	"github.com/ajstein51/sql-parser/pkg/lexer/types"
)

// example to lex the following SQL query: SELECT * FROM table WHERE column = 10

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(source string) []Token {
	var tokens []Token

	for len(source) > 0 {
		c := source[0]

		// go version of popping the first element off the string
		source = source[1:]

		// go through each character in the source and make tokens
		switch {
		case unicode.IsLetter(rune(c)):
			tokens = append(tokens, HandleLetters(source, c))

		case unicode.IsDigit(rune(c)):
			str := string(c)

			for len(source) > 0 && unicode.IsDigit(rune(source[0])) {
				str += string(source[0])
			}

			tokens = append(tokens, Token{Number, str})
		
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

func HandleLetters(source string, c byte) Token {
	str := string(c)

	for len(source) > 0 && unicode.IsLetter(rune(source[0])) {
		str += string(source[0])
	}

	if tokenType, ok := types.KeyWords[str]; ok {
		return Token{tokenType, str}
	}

	return Token{types.Identifier, str}
}