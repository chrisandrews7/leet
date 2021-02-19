package main

func shuffle(nums []int, n int) []int {
	result := make([]int, 0, len(nums))

	left := 0
	right := n

	for right < len(nums) {
		result = append(result, nums[left], nums[right])
		left++
		right++
	}

	return result
}
