package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbersString := readStringFromStdin()
	arr := convertToNumbers(numbersString)
	fmt.Printf("Parsed array is: %+v\n", arr)
	insertionSort(arr)
	fmt.Printf("Sorted array is: %+v\n", arr)
}

func readStringFromStdin() *string {
	var line string

	fmt.Println("Enter array where numbers separated by whitespace e.g. ' '")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}

	fmt.Printf("Provided string is: %q\n", line)

	return &line
}

func convertToNumbers(numbersString *string) []int {
	stringsArray := strings.Split(*numbersString, " ")
	result := make([]int, len(stringsArray))
	for i := range result {
		result[i], _ = strconv.Atoi(stringsArray[i])
	}
	return result
}

func insertionSort(arr []int) {
	for i := range arr {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
