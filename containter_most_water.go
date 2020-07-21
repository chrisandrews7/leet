package main

import (
	"math"
)

func containerMostWater(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		lowestValue := int(math.Min(float64(height[left]), float64(height[right])))

		if max := lowestValue * (right - left); max > maxArea {
			maxArea = max
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
