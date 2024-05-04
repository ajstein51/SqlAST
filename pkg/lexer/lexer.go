package lexer

import (
	"unicode"

	"github.com/ajstein51/sql-parser/pkg/lexer/types"
)

type Token struct {
	Type  types.TokenType
	Value string
}

func Tokenize(source string) []Token {
	var tokens []Token

	for len(source) > 0 {
		c := source[0]

		// go version of popping the first element off the string
		source = source[1:]

		switch {
		case c == ' ' || c == '\n' || c == '\t':
			continue
		case unicode.IsLetter(rune(c)):
			tokens = append(tokens, HandleLetters(&source, c))

		case unicode.IsDigit(rune(c)):
			str := string(c)

			for len(source) > 0 && unicode.IsDigit(rune(source[0])) {
				str += string(source[0])
				source = source[1:]
			}

			tokens = append(tokens, Token{types.Number, str})
		
		// operators
		case c == '(':
			tokens = append(tokens, Token{types.OpenParen, string(c)})
		case c == ')':
			tokens = append(tokens, Token{types.CloseParen, string(c)})
		case c == ',':
			tokens = append(tokens, Token{types.Comma, string(c)})
		case c == '*':
			tokens = append(tokens, HandleAsterisk(tokens, &source, c))
		case types.BinaryOperatorType(c) != types.BinaryOperatorTypeUnknown:
			tokens = append(tokens, Token{types.BinaryOperator, string(c)})

		default:
			tokens = append(tokens, Token{types.ILLEGAL, string(c)})
		}
	}

	return tokens
}

func HandleLetters(source *string, c byte) Token {
	str := string(c)

	for len(*source) > 0 && unicode.IsLetter(rune((*source)[0])) {
		str += string((*source)[0])
		*source = (*source)[1:]
	}

	if keyWordTokenType, ok := types.KeyWords[str]; ok {
		return Token{keyWordTokenType, str}
	}

	return Token{types.Identifier, str}
}

func HandleAsterisk(tokens []Token, source *string, c byte) Token {
	// we need to either perform a look ahead or look behind
	// to determine if the asterisk is a multiplication operator or a wildcard

	// lets first look behind if we have a column name or variable name then its a binary operator
	// if we have a SELECT, TOP()/Function, or a Comma then its a wildcard

	// get the last token
	// lastToken := tokens[len(tokens)-1]


	// Lets leave this as a TODO since we are gonna need to be further in the imp	
	return Token{types.BinaryOperator, string(c)}
}