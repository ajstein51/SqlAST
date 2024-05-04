package main

import (
	"fmt"

	// "github.com/ajstein51/sql-parser/pkg/lexer"
	// "github.com/ajstein51/sql-parser/pkg/lexer/types"
)

func main() {
	sql := "SELECT TOP(100) Col1, * FROM table"

	fmt.Print(sql)
	// tokens := lexer.Tokenize(sql)

	// fmt.Println("Test 2 passed %t", tokens[0].Type == types.Select)
	

}