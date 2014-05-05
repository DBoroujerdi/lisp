package main

import "fmt"
import "unicode/utf8"
import "github.com/dboroujerdi/stack"

// =====================================================

type SyntaxTree struct {
	exp *Expression
}

type Expression struct {
	elem *[]Element
}

type Element struct {
	typ ElementType
	val interface{}
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

func parse(input string) (string, error) {
	tree := new(SyntaxTree)
}

// recursive
func parseR(input string) (string, error) {

	var str = input
	s := new(stack.Stack)
	index := 0
	l := new(Expression)

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)

		if r == '(' {
			s.Push(r)

		} else if r == ')' {
			p, err := s.Pop()

			if err != nil {
				break
			}

			if p != '(' {
				break
			}
		} else if r >= '0' && r <= '9' {
			fmt.Printf("rune digit found! %c\n", r)
		}
		str = str[size:]

		index++
	}

	return input, nil
}
