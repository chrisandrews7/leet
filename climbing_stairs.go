package main

func climbStairs(n int) int {
	cache := make(map[int]int)

	cache[0] = 1
	cache[1] = 1

	for i := 1; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}
