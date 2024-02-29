package main

import "fmt"

type GraphEdge struct {
	to     int
	weight int
}

func NewGraphEdge(to int, weight int) *GraphEdge {
	return &GraphEdge{
		to:     to,
		weight: weight,
	}
}

type AdjacencyList struct {
	list [][]*GraphEdge
}

func NewAdjacencyList(list [][]*GraphEdge) *AdjacencyList {
	return &AdjacencyList{
		list: list,
	}
}

func sampleGraphAdjacencyList() *AdjacencyList {
	list := [][]*GraphEdge{
		{NewGraphEdge(1, 3), NewGraphEdge(2, 1)},
		{NewGraphEdge(4, 1)},
		{NewGraphEdge(3, 7)},
		{},
		{NewGraphEdge(1, 1), NewGraphEdge(3, 5), NewGraphEdge(5, 2)},
		{NewGraphEdge(2, 18), NewGraphEdge(6, 1)},
		{NewGraphEdge(3, 1)},
	}

	return NewAdjacencyList(list)
}

// ---------------------------------------------------------------------

type graphDFS struct {
	adjacencyList *AdjacencyList
	seen          []bool
	path          []int
}

func NewGraphDFS(adjacencyList *AdjacencyList) *graphDFS {
	return &graphDFS{
		adjacencyList: adjacencyList,
		seen:          make([]bool, len(adjacencyList.list)),
		path:          make([]int, 0),
	}
}

func (g *graphDFS) walk(current, needle int) bool {
	fmt.Println()
	fmt.Println("\t start walking...")
	// BASE CASE - needle found , already visited
	// if current == needle {
	// 	return true
	// }
	if g.seen[current] {
		g.showState(current)
		fmt.Println("\t already visited! :", current)
		return false
	}

	// set visited - true
	g.seen[current] = true

	// RECURSION
	// 1 - push & check for needle
	g.path = append(g.path, current)
	if current == needle {
		g.showState(current)
		fmt.Println("\t needle found !!")
		return true
	}

	// 2 - go through each edge in sublist of vertex
	list := g.adjacencyList.list[current]
	g.showState(current)
	for i := 0; i < len(list); i++ {
		// recurse for each edge
		edge := list[i]
		fmt.Println("\t edge: ", edge)
		if g.walk(edge.to, needle) {
			return true
		}
	}

	// 3 - pop
	lastPos := len(g.path) - 1
	fmt.Println("pop out from path: ", g.path[lastPos])
	g.path = g.path[:lastPos]

	g.showState(current)
	return false
}

func (g *graphDFS) dfs(start, needle int) []int {
	g.walk(start, needle)
	return g.path
}

func (g *graphDFS) showState(current int) {
	fmt.Println("\t current ->", current)
	fmt.Println("\t seen ->", g.seen)
	fmt.Println("\t path ->", g.path)
}
