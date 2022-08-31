package main

import (
	"fmt"
	"os"
)

func fibonacci(number int) int {
	if number < 2 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to read number")
		os.Exit(1)
	}

	result := fibonacci(i)
	fmt.Printf("Result: %d\n", result)
}
