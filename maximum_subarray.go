package main

func MaximumSubarray(nums []int) int {
	cache := make(map[int]int)
	largestSubArray := nums[0]

	cache[0] = nums[0]

	for i, num := range nums[1:] {
		index := i + 1
		max := num

		if total := cache[index-1] + num; total > num {
			max = total
		}

		cache[index] = max

		if max > largestSubArray {
			largestSubArray = max
		}
	}

	return largestSubArray
}
