package main

import "testing"
import "container/list"

func TestParseSymbol(t *testing.T) {

	result, _ := parseSymbol("foo bar bazz")

	if result != "foo" {
		t.Error("Parsed Symbol was", result)
	}
}

func TestElementString(t *testing.T) {
	var elem = Element{nil, SYM, "foo"}

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
	var elems list.List
	var expr = Expression{&elems}

	var elem = Element{nil, SYM, "foo"}
	expr.Add(elem)

	l := expr.Len()

	if l != 1 {
		t.Error("length was", l)
	}

	e := expr.elems.Front().Value.(*Element)

	if e == nil {
		t.Error("element was nil")
	}
}

func TestParseSingleSymbol(t *testing.T) {

	result, err := parse("(foo bar baz)")

	if err != nil {
		t.Error(err)
	}

	len := result.elems.Len()
	if len != 3 {
		t.Error("Num Elements expression was", len)
	}

	elem_1 := result.elems.Front()
	sym_1 := elem_1.Value.(*Element).String()

	if sym_1 != "foo" {
		t.Error("Result should be \"foo\" but was", sym_1)
	}

	result.elems.Remove(elem_1)
	elem_2 := result.elems.Front()
	sym_2 := elem_2.Value.(*Element).String()

	if sym_2 != "bar" {
		t.Error("Result should be \"bar\" but was", sym_2)
	}

	result.elems.Remove(elem_2)
	elem_3 := result.elems.Front()
	sym_3 := elem_3.Value.(*Element).String()

	if sym_3 != "baz" {
		t.Error("Result should be \"baz\" but was", sym_3)
	}
}

// TODO
// func TestParseBoolean(t *testing.T) {

// 	tru, _ := parse("#t")

// 	if tru != true {
// 		t.Error("Result should be \"true\" but was")
// 	}
// }

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
