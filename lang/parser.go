package lisp

import "unicode/utf8"
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

const (
	EXP TokenType = iota
	SYM
)

type Symbol struct {
	val string
}

// =====================================================

func Parse(input string) (*Expression, error) {
	var expr = new(Expression)

	if !IsValid(input[0:]) {
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

func IsValid(input string) bool {
	var str = input
	s := new(Stack)

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

			var elem = Token{SYM, sym}

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
