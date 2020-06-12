package main

func TwoSum(nums []int, target int) []int {
	previousSubtractions := make(map[int]int)

	for index, num := range nums {
		if partner, exists := previousSubtractions[num]; exists {
			return []int{partner, index}
		}

		previousSubtractions[target-num] = index
	}

	return []int{0, 0}
}
