package main

import (
	"fmt"
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
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			got := climbStairs(test.input)
			if got != test.expected {
				t.Errorf("climbStairs(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
