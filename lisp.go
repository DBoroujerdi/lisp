package main

import "fmt"
import "bufio"
import "os"
import "unicode/utf8"
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
			if !isValid(input) {
				fmt.Printf("Lisp is invalid!\n")
			} else {
				parse(input)
			}
		}
	}
}

// TODO remove??
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
