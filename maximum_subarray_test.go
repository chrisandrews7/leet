package main

import (
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, -3, 2, 1, -1}, 3},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := MaximumSubarray(test.input)
			if got != test.expected {
				t.Errorf("MaximumSubarray(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
