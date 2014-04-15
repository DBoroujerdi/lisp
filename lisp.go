package main

import "fmt"
import "bufio"
import "os"
import "unicode/utf8"
import "errors"
import "github.com/dboroujerdi/stack"

func main() {
	fmt.Printf("GoLisp Version 0.1\n")
	fmt.Printf("Press Ctrl+C to Exit\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("GoLisp>")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf(err.Error())
			break
		}

		if input == "exit\n" {
			fmt.Printf("\nExiting ... \n")
			break
		}

		if len(input) > 1 {
			parse(input)
		}
	}
}

type Expression interface {
	eval() float64
}

type Complex struct {
	op    Operand
	exprs []Expression
}

func (c *Complex) eval() float64 {
	return c.op.apply(c.exprs)
}

type Number struct {
	value float64
}

func (n *Number) eval() float64 {
	return n.value
}

type Operand interface {
	apply(exprs []Expression) float64
}

type Add struct{}

func (o *Add) apply(exprs ...Expression) (float64, error) {
	var r float64
	for _, e := range exprs {
		r += e.eval()
	}
	return r, nil
}

type Mult struct{}

func (o *Mult) apply(exprs ...Expression) (float64, error) {
	var r float64
	for _, e := range exprs {
		r *= e.eval()
	}
	return r, nil
}

type Div struct{}

func (o *Div) apply(exprs ...Expression) (float64, error) {
	if len(exprs) < 2 {
		return 0, errors.New("Div operand requires more than 1 expression")
	}
	first := exprs[0]
	r := first.eval()
	tail := exprs[1:len(exprs)]
	for _, e := range tail {
		er := e.eval()
		r = r / er
	}
	return r, nil
}

type Sub struct{}

func (o *Sub) apply(exprs ...Expression) (float64, error) {
	var r float64
	for _, e := range exprs {
		r *= e.eval()
	}
	return r, nil
}

func parse(input string) (Expression, error) {

	var str = input
	s := new(stack.Stack)

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)

		if r != '\n' {
			fmt.Printf("%c\n", r)
		}

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
	}

	return nil, nil
}

func evaluate(expr Expression) float64 {
	return expr.eval()
}
