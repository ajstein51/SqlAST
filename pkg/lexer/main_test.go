package lexer

import (
	"testing"
)

func TestNewLexer(t *testing.T) {
	sql := "SELECT * FROM table"
	
	test := Tokenize(sql)

	if len(test) != 4 {
		t.Errorf("Expected 4 tokens, got %d", len(test))
	}

}