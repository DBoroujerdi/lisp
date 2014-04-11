package main

import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Printf("GoLisp Version 0.1\n")
	fmt.Printf("Press Ctrl+c to Exit\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			// You may check here if err == io.EOF
			break
		}

		fmt.Printf(input)
	}
}
