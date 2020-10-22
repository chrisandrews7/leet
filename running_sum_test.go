package main

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := runningSum(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("runningSum(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
