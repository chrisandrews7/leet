package main

func Fibonacci(n int) int {
	cache := make(map[int]int)

	cache[1] = 1
	cache[2] = 1

	for num := 3; num <= n; num++ {
		cache[num] = cache[num-1] + cache[num-2]
	}

	return cache[n]
}
