package sorting

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	randomArray []int
)

func TestMain(m *testing.M) {
	setupTestData()

	st := m.Run()

	os.Exit(st)
}

func TestSort(t *testing.T) {
	testCases := []struct {
		name          string
		input         []int
		expected      []int
		expectedError error
	}{
		{
			name:          "simple case for reverse array",
			input:         []int{3, 2, 1, 0},
			expected:      []int{0, 1, 2, 3},
			expectedError: nil,
		}, {
			name:          "simple case for random array",
			input:         []int{49, 67, 406, 214, 426, 18},
			expected:      []int{18, 49, 67, 214, 406, 426},
			expectedError: nil,
		}, {
			name:          "edge case for empty array",
			input:         []int{},
			expected:      []int{},
			expectedError: nil,
		}, {
			name:          "edge case for nil array",
			input:         nil,
			expected:      nil,
			expectedError: ErrorArrayShouldNotBeNil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase // for async support
		t.Run(testCase.name, func(t *testing.T) {
			var toSort []int = nil
			if testCase.input != nil {
				toSort = make([]int, len(testCase.input))
				copy(toSort, testCase.input)
			}

			err := Sort(toSort, true)
			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("function does not return expected error: %v", testCase.expectedError)
			}
			if testCase.expectedError == nil && !reflect.DeepEqual(testCase.expected, toSort) {
				t.Errorf("function returns wrong answer: expected: %+v, got: %+v", testCase.expected, toSort)
			}
		})
	}
}

func BenchmarkSort_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(randomArray, true)
	}
}

func BenchmarkSort_NotOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(randomArray, false)
	}
}

func ExampleSort() {
	in := []int{9, 6, 4, 0, 7, 1}
	useOptimized := true

	err := Sort(in, useOptimized)
	if err != nil {
		panic(err)
	}

	fmt.Println(in)
	// output:
	// [0 1 4 6 7 9]
}

func setupTestData() {
	size := 500
	rand.Seed(time.Now().Unix())
	randomArray = rand.Perm(size)
}
