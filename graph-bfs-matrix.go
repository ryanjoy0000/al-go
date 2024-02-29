package main

import (
	"slices"
)

// type graphBFS struct {
// 	q         []int
// 	seen      []int
// 	prev      []int
// 	path      []int
// 	adjMatrix [][]int
// 	startV    int
// }

func graphBFS(graph [][]int, start, needle int) []int {
	// initial state - queue, seen, parent
	q := make([]int, 0)
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		seen[i] = false
		prev[i] = -1
	}

	// adding start vertex which changes state
	q = append(q, start)
	seen[start] = true

	// ----------A: graph traversal -----------------------------
	for len(q) != 0 {
		// pop
		current := q[0]
		q = q[1:]

		// check for needle
		if current == needle {
			break
		}

		// range through the weights of respective vertex in matrix
		subList := graph[current]
		for i := 0; i < len(subList); i++ {
			// skip if empty or already visited
			if subList[i] == 0 || seen[i] {
				continue
			} else {
				// if weight present, set state - seen, parent, queue
				seen[i] = true
				prev[i] = current
				q = append(q, i)
			}
		}

	}

	// ----------B: reconstruct path  -----------------------------

	// start from end point
	curr := needle
	// path: end to start
	out := make([]int, 0)

	// until parent is empty
	for prev[curr] != -1 {
		// store the vertex
		out = append(out, curr)
		// set the parent as current
		curr = prev[curr]
	}

	if len(out) != 0 {
		// reverse path: start -> end
		slices.Reverse(out)
		// prepend the start vertex to out
		out = append([]int{start}, out...)
	}

	return out
}

func sampleGraphData() [][]int {
	return [][]int{
		{0, 3, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}
}
