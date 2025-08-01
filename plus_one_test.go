package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			got := plusOne(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("plusOne(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
