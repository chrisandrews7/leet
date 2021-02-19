package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		input    int
		expected []string
	}{
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := fizzBuzz(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("fizzBuzz(%d) got %s, want %s", test.input, got, test.expected)
			}
		})
	}
}
