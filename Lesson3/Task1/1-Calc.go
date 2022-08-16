package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

type inputArguments struct {
	FirstNumber  float64
	SecondNumber float64
}

type operationFunction func(*inputArguments) (float64, error)

func main() {
	args, operation := getInputArguments()
	result, err := calculate(args, operation)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func getInputArguments() (*inputArguments, string) {
	var a, b float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	fmt.Print("Enter operation (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)

	return &inputArguments{FirstNumber: a, SecondNumber: b}, op
}

func calculate(args *inputArguments, operationName string) (float64, error) {
	supportedOperations := map[string]operationFunction{
		"+": plusOperation,
		"-": minusOperation,
		"*": multiplyOperation,
		"/": divideOperation,
		"!": factOperation,
		"^": powerOperation,
	}

	// Validation
	operationFunction, isOperationSupported := supportedOperations[operationName]
	if !isOperationSupported {
		return 0, fmt.Errorf("operation '%s' is not supported", operationName)
	}

	// Calculation
	return operationFunction(args)
}

func plusOperation(args *inputArguments) (float64, error) {
	return args.FirstNumber + args.SecondNumber, nil
}

func minusOperation(args *inputArguments) (float64, error) {
	return args.FirstNumber - args.SecondNumber, nil
}

func multiplyOperation(args *inputArguments) (float64, error) {
	return args.FirstNumber * args.SecondNumber, nil
}

func divideOperation(args *inputArguments) (float64, error) {
	if args.SecondNumber == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return args.FirstNumber / args.SecondNumber, nil
}

func powerOperation(args *inputArguments) (float64, error) {
	return math.Pow(args.FirstNumber, float64(args.SecondNumber)), nil
}

func factOperation(args *inputArguments) (float64, error) {
	if args.FirstNumber < 0 {
		return 0, errors.New("factorial of negative number doesn't exist")
	}

	var result float64 = 1
	for i := 1; i <= int(args.FirstNumber); i++ {
		result *= float64(i) // mismatched types int64 and int
	}

	return result, nil
}
