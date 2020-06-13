package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{9, 55},
		{13, 377},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			got := ClimbStairs(test.input)
			if got != test.expected {
				t.Errorf("ClimbStairs(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
