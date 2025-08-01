package main

import (
	"testing"
)

func TestRansomNote(t *testing.T) {
	var tests = []struct {
		inputNote     string
		inputMagazine string
		expected      bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := canConstruct(test.inputNote, test.inputMagazine)
			if got != test.expected {
				t.Errorf("canConstruct(%s, %s) got %v, want %v", test.inputNote, test.inputMagazine, got, test.expected)
			}
		})
	}
}
