package main

func houseRobber(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	cache := make(map[int]int)
	cache[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		nextDoor := cache[i-1]
		twoDown := cache[i-2] + nums[i]

		if nextDoor > twoDown {
			cache[i] = nextDoor
			continue
		}

		cache[i] = twoDown
	}

	return cache[len(nums)-1]
}
