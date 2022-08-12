package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber float32

	fmt.Println("This app calculate rectangle area based on two numbers")
	fmt.Print("Enter first number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&secondNumber)

	fmt.Printf("Rectangle area is equal to: %f\n", firstNumber*secondNumber)
}
