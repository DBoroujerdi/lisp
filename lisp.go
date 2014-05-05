package main

import "fmt"
import "bufio"
import "os"

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
