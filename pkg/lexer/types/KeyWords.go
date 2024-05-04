package types

var KeyWords = map[string]TokenType {
	"SELECT" : Select,
	"FROM" : From,
	"WHERE" : Where,
	"AND" : And,
	"OR" : Or,
	"JOIN" : Join,
	"INNER JOIN" : InnerJoin,
	"OUTER JOIN" : OuterJoin,
	"LEFT JOIN" : LeftJoin,
	"RIGHT JOIN" : RightJoin,
	"AS" : As,
	"ON" : On,
	"ORDER BY" : OrderBy,
	"GROUP BY" : GroupBy,
}