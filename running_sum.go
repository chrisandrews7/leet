package main

func runningSum(nums []int) (result []int) {
	result = append(result, nums[0])

	for index, num := range nums[1:] {
		result = append(result, num+result[index])
	}

	return
}
