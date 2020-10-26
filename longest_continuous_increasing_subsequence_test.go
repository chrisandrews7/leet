package main

import (
	"reflect"
	"testing"
)

func TestLongestContinuousIncreasingSubsequence(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
		{[]int{}, 0},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := findLengthOfLCIS(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("findLengthOfLCIS(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
