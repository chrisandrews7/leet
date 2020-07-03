package main

import (
	"testing"
)

func TestContainerMostWater(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 3, 2, 5, 25, 24, 5}, 24},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := ContainerMostWater(test.input)
			if got != test.expected {
				t.Errorf("TestContainerMostWater(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
