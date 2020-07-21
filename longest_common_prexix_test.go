package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		input    []string
		expected string
	}{
		{[]string{"flow", "flower", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"dogs", "dogdog", "dogs"}, "dog"},
		{[]string{}, ""},
		{[]string{"a"}, "a"},
		{[]string{"abab", "aba", ""}, ""},
		{[]string{"ca", "a"}, ""},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := longestCommonPrefix(test.input)
			if got != test.expected {
				t.Errorf("longestCommonPrefix(%s) got %v, want %v", test.input, got, test.expected)
			}
		})
	}
}
