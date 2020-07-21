package main

func removeArrayDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0

	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}

		right++
	}

	return left + 1
}
