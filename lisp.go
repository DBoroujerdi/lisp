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

			fmt.Printf(err.String())
			break
		}

		if input == "exit\n" {

			break
		}

		if len(input) > 1 {
			fmt.Printf(input)
		}
	}
}
