package main

import (
	"testing"
)

func TestConvertNumberToHexidecimal(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{26, "1a"},
		{113, "71"},
		{257, "101"},
		{60, "3c"},
		{4096, "1000"},
		{4160, "1040"},
		{-1, "ffffffff"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := toHex(test.input)
			if got != test.expected {
				t.Errorf("toHex(%d) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
