package types

type BinaryOperatorType rune

const (
	BinaryOperatorTypeUnknown BinaryOperatorType = 0

	BinaryOperatorTypePlus BinaryOperatorType = '+'
	BinaryOperatorTypeMinus BinaryOperatorType = '-'
	BinaryOperatorTypeMultiply BinaryOperatorType = '*'
	BinaryOperatorTypeDivide BinaryOperatorType = '/'
)