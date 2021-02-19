package main

func runningSum(nums []int) []int {
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])

	for index, num := range nums[1:] {
		result = append(result, num+result[index])
	}

	return result
}
