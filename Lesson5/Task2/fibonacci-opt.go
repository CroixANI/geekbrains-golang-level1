package main

import (
	"fmt"
	"os"
)

func fibonacci(cache map[int]uint64, number int) uint64 {
	fibonacciResult, isResultExist := cache[number]
	if !isResultExist {
		fibonacciResult = fibonacci(cache, number-1) + fibonacci(cache, number-2)
		cache[number] = fibonacciResult
	}

	return fibonacciResult
}

func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to read number")
		os.Exit(1)
	}
	if i < 0 {
		fmt.Fprintln(os.Stderr, "There is no fibonacci for negative number")
		os.Exit(2)
	}

	cache := map[int]uint64{
		0: 0,
		1: 1,
		2: 1,
	}
	result := fibonacci(cache, i)
	fmt.Printf("Result: %d\n", result)
}
