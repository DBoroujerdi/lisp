package main

import "testing"

func TestParseSymbol(t *testing.T) {

	result, _ := parseSymbol("foo bar bazz")

	if result != "foo" {
		t.Error("Parsed Symbol was", result)
	}
}

func TestElementString(t *testing.T) {
	var elem = Element{SYM, "foo"}

	val := elem.String()

	if val != "foo" {
		t.Error("val was", val)
	}

	typ := elem.typ

	if typ != 1 {
		t.Error("type was", typ)
	}
}

func TestExpression(t *testing.T) {

	var expr = new(Expression)

	var elem = Element{SYM, "foo"}
	expr.Add(elem)

	l := len(expr.elems)
	if l != 1 {
		t.Error("length was", l)
	}

	e := expr.elems[0]

	if e == nil {
		t.Error("element was nil")
	}
}

func TestParseSingleSymbol(t *testing.T) {

	result, err := parse("(foo bar baz)")

	if err != nil {
		t.Error(err)
	}

	len := len(result.elems)
	if len != 3 {
		t.Error("Num Elements expression was", len)
	}

	elem_1 := result.elems[0]
	sym_1 := elem_1.String()

	if sym_1 != "foo" {
		t.Error("Result should be \"foo\" but was", sym_1)
	}

	elem_2 := result.elems[1]
	sym_2 := elem_2.String()

	if sym_2 != "bar" {
		t.Error("Result should be \"bar\" but was", sym_2)
	}

	elem_3 := result.elems[2]
	sym_3 := elem_3.String()

	if sym_3 != "baz" {
		t.Error("Result should be \"baz\" but was", sym_3)
	}
}

func TestParseBoolean(t *testing.T) {

	result, _ := parse("(#t)")

	elem := result.elems[0]
	val := elem.String()

	if val != "true" {
		t.Error("Result should be \"true\" but was")
	}
}

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
