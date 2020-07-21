package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		inputArr []int
		inputSum int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{45, 32, 84, 55, 35}, 100, []int{0, 3}},
		{[]int{1, 2}, 3, []int{0, 1}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := twoSum(test.inputArr, test.inputSum)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("twoSum(%d, %d) got %v, want %v", test.inputSum, test.inputArr, got, test.expected)
			}
		})
	}
}
