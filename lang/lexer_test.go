package lisp

import "testing"

func TestLexerConstructor(t *testing.T) {
	lexer := NewLexer("(+ 2 2)")

	if lexer.input != "(+ 2 2)" {
		t.Error("lexer.input was not initialized properly", lexer)
	}
}

func TestGetTokens(t *testing.T) {
	lexer := NewLexer("(+ 2 2)")

	tokens, err := lexer.GetTokens()

	if err != nil {
		t.Error(err)
	}

	if len(tokens) != 5 {
		t.Error("Incorrect number of tokens", tokens)
	}
}

func TestGetTokensWithEnexpectedRune(t *testing.T) {
	lexer := NewLexer("(& 2 2)")

	_, err := lexer.GetTokens()

	if err == nil {
		t.Error("Error expected")
	}
}
