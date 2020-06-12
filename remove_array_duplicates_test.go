package main

import (
	"reflect"
	"testing"
)

func TestRemoveArrayDuplicates(t *testing.T) {
	var tests = []struct {
		input       []int
		expectedArr []int
		expectedLen int
	}{
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := RemoveArrayDuplicates(test.input)
			if got != test.expectedLen {
				t.Errorf("RemoveArrayDuplicates(%d) de-duplicated length was %v, wanted %v", test.input, got, test.expectedLen)
			}
			if !reflect.DeepEqual(test.input[:test.expectedLen], test.expectedArr) {
				t.Errorf("RemoveArrayDuplicates(%d) de-duplicated array was %v, wanted %v", test.input, test.input[:test.expectedLen], test.expectedArr)
			}
		})
	}
}
