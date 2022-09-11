package fibonacci

import "errors"

var ErrorNegativeNumbers = errors.New("there is no fibonacci for negative number")

func Fibonacci(number int, useOptimized bool) (uint64, error) {
	if number < 0 {
		return 0, ErrorNegativeNumbers
	}

	if !useOptimized {
		return oldFibonacci(number), nil
	}

	cache := map[int]uint64{
		0: 0,
		1: 1,
		2: 1,
	}

	return cachedFibonacci(cache, number), nil
}

func cachedFibonacci(cache map[int]uint64, number int) uint64 {
	if number < 2 {
		return uint64(number)
	}
	fibonacciResult, isResultExist := cache[number]
	if !isResultExist {
		fibonacciResult = cachedFibonacci(cache, number-1) + cachedFibonacci(cache, number-2)
		cache[number] = fibonacciResult
	}

	return fibonacciResult
}

func oldFibonacci(number int) uint64 {
	if number < 2 {
		return uint64(number)
	}

	return oldFibonacci(number-1) + oldFibonacci(number-2)
}
