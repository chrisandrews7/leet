package main

import (
	"math"
)

func ReverseInt(x int) int {
	var reversed int

	for x != 0 {
		popped := x % 10
		x /= 10
		reversed = reversed*10 + popped

		if reversed > math.MaxInt32 || reversed < math.MinInt32 {
			return 0
		}
	}

	return reversed
}
