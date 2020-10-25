package main

type vertex struct {
	value int
	prev  *vertex
}

func allPathsSourceTarget(graph [][]int) (paths [][]int) {
	target := len(graph) - 1
	stack := []vertex{}

	for _, value := range graph[0] {
		stack = append(stack, vertex{
			value: value,
			prev: &vertex{
				value: 0,
			},
		})
	}

	for len(stack) > 0 {
		var currentVertex vertex
		currentVertex, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if currentVertex.value == target {
			path := []int{}
			current := &currentVertex
			for current != nil {
				path = append([]int{current.value}, path...)
				current = current.prev
			}

			paths = append(paths, path)
		}

		for _, value := range graph[currentVertex.value] {
			stack = append(stack, vertex{
				value: value,
				prev:  &currentVertex,
			})
		}
	}

	return
}
