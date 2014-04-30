package main

import "testing"

func TestParse_1(t *testing.T) {

	input := "(+ 5 5)"
	var _, _ = parse(input)
}

func TestValidParens(t *testing.T) {

	input := "(+ 5 5)"

	valid := isValid(input)

	if !valid {
		t.Error("Valid lisp " + input + " should be valid!")
	}
}

func TestInvalidParens(t *testing.T) {
	input := "(+ 5 ((5)"

	valid := isValid(input)

	if valid {
		t.Error("Invalid lisp " + input + " should be Invalid!")
	}
}
