package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2, 4, 1}, 2},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := maxProfit(test.input)
			if got != test.expected {
				t.Errorf("maxProfit(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
