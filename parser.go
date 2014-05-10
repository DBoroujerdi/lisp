package main

import "unicode/utf8"
import "github.com/dboroujerdi/stack"
import "errors"
import "bytes"
import "fmt"

// =====================================================

type Expression struct {
	tkns []*Token
}

func (expr *Expression) Pretty() string {
	p := "("

	for i := range expr.tkns {
		t := expr.tkns[i]

		switch t.typ {
		case EXP:
			p = p + t.val.(*Expression).Pretty()
			break
		case SYM:
			p = p + fmt.Sprintf("%v", t.val) + " "
			break
		default:
			p = p + "nil"
		}
	}

	p = p + ")"
	return p
}

func (expr *Expression) Len() int {
	return len(expr.tkns)
}

func (expr *Expression) Add(e Token) {
	expr.tkns = append(expr.tkns, &e)
}

type Token struct {
	typ TokenType
	val interface{}
}

func (e *Token) String() string {
	switch e.typ {
	case EXP:
		return "Not implemented yet"
	case SYM:
		return fmt.Sprintf("%v", e.val)
	default:
		return "nil"
	}
}

type TokenType int

const (
	EXP TokenType = iota
	SYM
)

type Symbol struct {
	val string
}

// =====================================================

func parse(input string) (*Expression, error) {
	var expr = new(Expression)

	if !isValid(input[0:]) {
		return nil, errors.New("Invalid Parenthesis!")
	}

	_, err := parseR(expr, input[1:])

	if err != nil {
		return nil, err
	}

	return expr, nil
}

func isLetter(r rune) bool {
	isCaps := r >= 'A' && r <= 'Z'
	isLower := r >= 'a' && r <= 'z'
	return isCaps || isLower
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSpecial(r rune) bool {
	return r == '#'
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
	fmt.Printf("Parsing expression [%s]\n", input)
	var str = input
	index := 0

	for len(str) > 0 {

		r, size := utf8.DecodeRuneInString(str)

		if r == ')' {

			index++
			break
		} else if r == '(' {

			b := index + 1
			fmt.Printf("'(' found at index %d in string %s \n", b, str)
			var subExpr = new(Expression)
			s, err := parseR(subExpr, str[1:])

			if err != nil {
				return -1, err
			}

			elem := Token{EXP, subExpr}
			expr.Add(elem)
			str = str[s:]
			index = index + s
		} else if isLetter(r) || isSpecial(r) {

			sym, s := parseSymbol(str)

			var elem Token
			if sym == "#t" {
				elem = Token{SYM, "true"}
			} else if sym == "#f" {
				elem = Token{SYM, "false"}
			} else {
				elem = Token{SYM, sym}
			}

			expr.Add(elem)
			str = str[s:]
			index = index + s
		} else {
			str = str[size:]
		}

		index++
	}

	return index, nil
}
