package main

import "unicode/utf8"
import "github.com/dboroujerdi/stack"
import "errors"
import "bytes"
import "container/list"
import "fmt"

// =====================================================

type Expression struct {
	elems *list.List
}

func (expr *Expression) Len() int {
	return expr.elems.Len()
}

func (expr *Expression) Add(e Element) {
	expr.elems.PushBack(&e)
}

type Element struct {
	elem *Element
	typ  ElementType
	val  interface{}
}

func (e *Element) String() string {
	switch e.typ {
	case EXP:
		return "Not implemented yet"
	case SYM:
		return fmt.Sprintf("%v", e.val)
	default:
		return "nil"
	}
}

type ElementType int

const (
	EXP ElementType = iota
	SYM
)

type Symbol struct {
	val string
}

// =====================================================

func parse(input string) (*Expression, error) {
	var elems list.List
	var expr = Expression{&elems}

	if !isValid(input[0:]) {
		return nil, errors.New("Invalid Parenthesis!")
	}

	_, err := parseR(&expr, input[1:])

	if err != nil {
		return nil, err
	}

	return &expr, nil
}

func isLetter(r rune) bool {
	isCaps := r >= 'A' && r <= 'Z'
	isLower := r >= 'a' && r <= 'z'
	return isCaps || isLower
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func parseSymbol(s string) (string, int) {
	var str = s
	var res bytes.Buffer

	for len(str) > 0 {

		r, size := utf8.DecodeRuneInString(str)

		if r == ' ' || r == ')' {
			break
		}

		res.WriteRune(r)
		str = str[size:]
	}

	rs := res.String()

	return rs, len(rs)
}

func isValid(input string) bool {
	var str = input
	s := new(stack.Stack)

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)

		if r == '(' {
			s.Push(r)

		} else if r == ')' {
			_, err := s.Pop()

			if err != nil {
				return false
			}
		}

		str = str[size:]
	}

	if !s.IsEmpty() {
		return false
	}

	return true
}

func parseR(expr *Expression, input string) (int, error) {
	var str = input
	index := 0

	for len(str) > 0 {

		r, size := utf8.DecodeRuneInString(str)

		if r == ')' {

			index++
			break
		} else if r == '(' {

			var elems list.List
			var subExpr = Expression{&elems}
			s, err := parseR(&subExpr, str[0:])

			if err != nil {
				return -1, err
			}

			elem := Element{nil, EXP, subExpr}
			expr.Add(elem)
			str = str[s:]
		} else if isLetter(r) {

			sym, s := parseSymbol(str)
			elem := Element{nil, SYM, sym}
			expr.elems.PushBack(&elem)
			str = str[s:]
		} else {
			str = str[size:]
		}

		index++
	}

	return index, nil
}
