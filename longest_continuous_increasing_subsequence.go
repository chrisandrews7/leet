package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	maxLength := 1
	length := 1

	for index, num := range nums[1:] {
		if nums[index] < num {
			length++

			if length > maxLength {
				maxLength = length
			}

			continue
		}

		length = 1
	}

	return maxLength
}
