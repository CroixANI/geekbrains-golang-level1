package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lesson10/numbers"
	"lesson10/sorting"
)

func main() {
	numbersString := readStringFromStdin()
	arr := convertToNumbers(numbersString)
	fmt.Printf("Parsed array is: %+v\n", arr)
	sorting.Sort(arr, false)
	fmt.Printf("Sorted array is: %+v\n", arr)

	_, _, hundreds, err := numbers.SplitNumber(923)
	if err != nil {
		panic(err)
	}

	fmt.Println(hundreds)
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
