package main

type Cache map[int]int

func climb(n int, cache Cache) int {
	if n < 2 {
		return 1
	}

	if result, exists := cache[n]; exists {
		return result
	}

	result := climb(n-1, cache) + climb(n-2, cache)
	cache[n] = result

	return result
}

func ClimbStairs(n int) int {
	return climb(n, make(Cache))
}
