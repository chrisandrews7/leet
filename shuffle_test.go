package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	var tests = []struct {
		inputArr []int
		inputSum int
		expected []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := shuffle(test.inputArr, test.inputSum)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("shuffle(%d, %d) got %v, want %v", test.inputSum, test.inputArr, got, test.expected)
			}
		})
	}
}
