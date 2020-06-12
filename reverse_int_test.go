package main

import (
	"testing"
)

func TestReverseInt(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{657, 756},
		{1534236469, 0},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			got := ReverseInt(test.input)
			if got != test.expected {
				t.Errorf("ReverseInt(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
