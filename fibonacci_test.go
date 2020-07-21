package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			got := fibonacci(test.input)
			if got != test.expected {
				t.Errorf("fibonacci(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
