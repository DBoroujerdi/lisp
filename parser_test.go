package main

import "testing"
import "fmt"

func TestParseSymbol(t *testing.T) {

	result, _ := parseSymbol("foo bar bazz")

	if result != "foo" {
		t.Error("Parsed Symbol was", result)
	}
}

func TestTokenString(t *testing.T) {
	var elem = Token{SYM, "foo"}

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

	var elem = Token{SYM, "foo"}
	expr.Add(elem)

	l := len(expr.tkns)
	if l != 1 {
		t.Error("length was", l)
	}

	e := expr.tkns[0]

	if e == nil {
		t.Error("element was nil")
	}
}

func TestParseSingleSymbol(t *testing.T) {

	result, err := parse("(foo bar baz)")

	if err != nil {
		t.Error(err)
	}

	len := len(result.tkns)
	if len != 3 {
		t.Error("Num Tokens expression was", len)
	}

	elem_1 := result.tkns[0]
	sym_1 := elem_1.String()

	if sym_1 != "foo" {
		t.Error("Result should be \"foo\" but was", sym_1)
	}

	elem_2 := result.tkns[1]
	sym_2 := elem_2.String()

	if sym_2 != "bar" {
		t.Error("Result should be \"bar\" but was", sym_2)
	}

	elem_3 := result.tkns[2]
	sym_3 := elem_3.String()

	if sym_3 != "baz" {
		t.Error("Result should be \"baz\" but was", sym_3)
	}
}

func TestParseExpressionWithSubExpression(t *testing.T) {
	fmt.Printf("\nRunning test TestParseExpressionWithSubExpression\n")

	result, err := parse("(foo (bar baz))")

	fmt.Printf("Pretty printed: " + result.Pretty() + "\n")

	if err != nil {
		t.Error(err)
	}

	len := result.Len()
	ex := 2
	if len != ex {
		t.Error(fmt.Sprintf("Num Tokens expected [%d] expression was [%d]", ex, len))
	}

	foo := result.tkns[0].val
	if foo != "foo" {
		t.Error(fmt.Sprintf("Token expected to be [foo] but was %s", foo))
	}

	subExprToken := result.tkns[1]
	if subExprToken.typ != 0 {
		t.Error(fmt.Sprintf("Token expected to be of TokenType EXP but was %s", subExprToken.typ))
	}

	subExpr, ok := subExprToken.val.(*Expression)
	if !ok {
		t.Error("Unable to convert type")
	}

	fmt.Printf(subExpr.tkns[0].val.(string) + "\n")
	fmt.Printf(subExpr.tkns[1].val.(string) + "\n")
	l := subExpr.Len()
	fmt.Printf("%d\n", l)

	fmt.Printf("\n")
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
