package main

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceToTarget(t *testing.T) {
	var tests = []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 2, 3}, {0, 1, 3}}},
		{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, [][]int{{0, 1, 4}, {0, 1, 2, 3, 4}, {0, 1, 3, 4}, {0, 3, 4}, {0, 4}}},
		{[][]int{{1}, {}}, [][]int{{0, 1}}},
		{[][]int{{1, 2, 3}, {2}, {3}, {}}, [][]int{{0, 3}, {0, 2, 3}, {0, 1, 2, 3}}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := allPathsSourceTarget(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("allPathsSourceTarget(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
