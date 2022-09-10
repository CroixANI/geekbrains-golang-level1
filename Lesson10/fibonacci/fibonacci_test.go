package fibonacci

import (
	"errors"
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name          string
		input         int
		expected      uint64
		expectedError error
	}{
		{
			name:          "simple case for 10",
			input:         10,
			expected:      55,
			expectedError: nil,
		}, {
			name:          "simple case for 11",
			input:         11,
			expected:      89,
			expectedError: nil,
		}, {
			name:          "edge case for 0",
			input:         0,
			expected:      0,
			expectedError: nil,
		}, {
			name:          "edge case for 1",
			input:         1,
			expected:      1,
			expectedError: nil,
		}, {
			name:          "edge case for 2",
			input:         2,
			expected:      1,
			expectedError: nil,
		}, {
			name:          "edge case for negative number",
			input:         -1,
			expected:      0,
			expectedError: ErrorNegativeNumbers,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase // for async support
		t.Run(testCase.name, func(t *testing.T) {
			result, err := Fibonacci(testCase.input, true)
			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("function does not return expected error: %v", testCase.expectedError)
			}
			if testCase.expectedError == nil && result != testCase.expected {
				t.Errorf("function returns wrong answer: expected: %d, got: %d", testCase.expected, result)
			}
		})
	}
}

func BenchmarkFibonacci_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20, true)
	}
}

func BenchmarkFibonacci_NotOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20, false)
	}
}

func ExampleFibonacci() {
	in := 10
	useOptimized := true

	result, err := Fibonacci(in, useOptimized)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	// output:
	// 55
}
