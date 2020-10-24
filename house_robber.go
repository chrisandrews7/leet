package main

func rob(nums []int, current int, cache map[int]int) int {
	switch current {
	case len(nums) - 1:
		return nums[current]
	case len(nums):
		return 0
	}

	if value, ok := cache[current]; ok {
		return value
	}

	twoDown := nums[current]
	if len(nums)+2 > current {
		twoDown += rob(nums, current+2, cache)
	}

	if nextDoor := rob(nums, current+1, cache); nextDoor > twoDown {
		cache[current] = nextDoor
		return nextDoor
	}

	cache[current] = twoDown
	return twoDown
}

func houseRobber(nums []int) int {
	cache := make(map[int]int)
	return rob(nums, 0, cache)
}
