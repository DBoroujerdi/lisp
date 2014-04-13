package main

import "fmt"
import "bufio"
import "os"
import "unicode/utf8"

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

func (o *Add) apply(exprs ...Expression) float64 {
	var r float64
	for _, e := range exprs {
		r += e.eval()
	}
	return r
}

type Mult struct{}

func (o *Mult) apply(exprs ...Expression) float64 {
	var r float64
	for _, e := range exprs {
		r *= e.eval()
	}
	return r
}

type Div struct{}

func (o *Div) apply(exprs ...Expression) float64 {
	if len(exprs) < 2 {
		return 0 //, errors.new("Div operand requires more than 1 expression")
	}
	first := exprs[0]
	r := first.eval()
	tail := exprs[1:len(exprs)]
	for _, e := range tail {
		er := e.eval()
		r = r / er
	}
	return r
}

type Sub struct{}

func (o *Sub) apply(exprs ...Expression) float64 {
	var r float64
	for _, e := range exprs {
		r *= e.eval()
	}
	return r
}

func parse(input string) Expression {

	sub := func(strExpr string) bool {
		return false
	}

	inner := func(str string) int {
		index := 0
		for len(str) > 0 {
			r, size := utf8.DecodeRuneInString(str)

			if r == '(' {
				return index
			}

			str = str[size:]
			index++
		}
		return -1
	}

	outer := func(str string) int {
		index := len(str)
		for len(str) > 0 {
			r, size := utf8.DecodeLastRuneInString(str)

			if r == ')' {
				return index
			}

			str = str[:len(str)-size]
			index--
		}
		return -1
	}

	b := []byte(input)

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c\n", r)

		b = b[size:]
	}

	return nil
}

func evaluate(expr Expression) float64 {
	return expr.eval()
}
