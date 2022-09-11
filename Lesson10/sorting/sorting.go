package sorting

import "errors"

var (
	ErrorArrayShouldNotBeNil = errors.New("array should not be nil")
)

func Sort(arr []int, useOptimized bool) error {
	if arr == nil {
		return ErrorArrayShouldNotBeNil
	}

	if useOptimized {
		insertionSort(arr)
	} else {
		bubbleSort(arr)
	}

	return nil
}

func insertionSort(arr []int) {
	for i := range arr {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func bubbleSort(arr []int) {
	isDone := false
	arrayLength := len(arr) - 1
	for !isDone {
		isDone = true
		for i := 0; i < arrayLength; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}
		}
	}
}
