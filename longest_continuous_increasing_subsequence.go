package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	maxLength := 1
	windowStart := 0

	for i, num := range nums[1:] {
		index := i + 1

		if nums[index-1] >= num {
			windowStart = index
		}

		if length := index - windowStart + 1; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
