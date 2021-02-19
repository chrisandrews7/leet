package main

import "fmt"

func getValue(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprint(n)
	}
}

func fizzBuzz(n int) []string {
	result := make([]string, 0, n)

	for i := 1; i < n+1; i++ {
		result = append(result, getValue(i))
	}

	return result
}
