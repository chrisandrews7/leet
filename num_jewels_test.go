package main

import (
	"reflect"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	var tests = []struct {
		inputJewels string
		inputStones string
		expected    int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := numJewelsInStones(test.inputJewels, test.inputStones)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("numJewelsInStones(%s, %s) got %d, want %d", test.inputStones, test.inputJewels, got, test.expected)
			}
		})
	}
}
