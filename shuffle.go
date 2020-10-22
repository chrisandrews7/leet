package main

func shuffle(nums []int, n int) (result []int) {
	left := 0
	right := n

	for right < len(nums) {
		result = append(result, nums[left], nums[right])
		left++
		right++
	}

	return
}
