package types

type TokenType int

const (
	ILLEGAL TokenType = iota

	Identifier	
	Number
	Select
	From
	Where
	And
	Or
	Join
	InnerJoin
	OuterJoin
	LeftJoin
	RightJoin
	As
	On
	OrderBy
	GroupBy

	BinaryOperator
	CloseParen
	OpenParen
	Comma
)