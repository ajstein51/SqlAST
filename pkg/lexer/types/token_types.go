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

	BinaryOperator // going to have to make something to tell a * as all the columns or a * as multiplication
	CloseParen
	OpenParen
	Comma
)