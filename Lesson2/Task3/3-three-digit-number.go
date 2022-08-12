package main

import (
	"fmt"
	"os"
)

func splitNumber(number int) (int, int, int) {
	var hundreds = number / 100
	var tens = (number - (hundreds * 100)) / 10
	var units = number - (hundreds * 100) - (tens * 10)
	return units, tens, hundreds
}

func main() {
	var number int

	fmt.Println("This app calculate diameter and circumference by given area of a circle")
	fmt.Print("Enter 3 digits number: ")
	fmt.Scanln(&number)

	if number > 999 || number <= 99 {
		fmt.Fprintln(os.Stderr, "This app works with only three digits number")
		os.Exit(1)
	}

	units, tens, hundreds := splitNumber(number)
	fmt.Printf("Hundreds are equal to: %d\n", hundreds)
	fmt.Printf("Tens are equal to: %d\n", tens)
	fmt.Printf("Units are equal to: %d\n", units)
}
