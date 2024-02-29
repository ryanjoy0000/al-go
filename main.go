package main

import "fmt"

func main() {
	fmt.Println("Popular Algorithms...")

	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// val := 9

	// result := linearSearch(x, val)
	// result := binarySearchSortedArr(x, val)
	// fmt.Println("Search for", val, " in ", x ,": ", result)

	// crystal ball problem
	// resultChan := make(chan int)
	// go oldWay(a, resultChan)
	// go solveCrystalBalls(a, resultChan)
	// fmt.Println("Found true at position: ",<-resultChan)

	// bubble sort
	// a := []int{4, 1, 7, 6, 2}
	// a := getLargeSlice(1_000_000_000)
	// fmt.Println("input: ", a)
	// fmt.Println(bubbleSort(a))

	// a := []int{5, 3, 1, 7, 8, 9}
	// fmt.Println(sortThisQuick(a))
	//
	//
	// --------- LINKED LIST --------------
	// l1 := NewLinkedList()
	// fmt.Println("l1", *l1)
	// l1.insert(NewNode(5), 0)
	// l1.insert(NewNode(6), 5)
	// l1.insert(NewNode(7), 6)
	// l1.insert(NewNode(2), 7)
	// l1.insert(NewNode(8), 2)
	// l1.show()
	//
	// l2 := NewLinkedList()
	// n1 := NewNode(8)
	// n2 := NewNode(9)
	// l2.insert(n1, 0)
	// l2.insert(n2, 8)
	// l2.show()
	// l2.delete(n2)
	// l2.show()
	// l2.delete(n1)
	// l2.show()

	// --------- QUEUE --------------
	//
	// q1 := NewQueue()
	// fmt.Println("q1", *q1)
	// q1.enqueue(NewNode(5))
	// q1.show()
	// q1.enqueue(NewNode(6))
	// q1.show()
	// q1.enqueue(NewNode(7))
	// q1.show()
	// q1.enqueue(NewNode(2))
	// q1.show()
	// q1.enqueue(NewNode(8))
	// q1.show()
	//
	// for q1.length > 0 {
	// 	q1.dequeue()
	// 	q1.show()
	// }
	//
	// --------- STACK ---------------
	// s1.push(NewNode(546))
	// s1.push(NewNode(32))
	// s1.push(NewNode(87))
	// s1.push(NewNode(78))
	// s1.push(NewNode(98))
	// s1.push(NewNode(778))
	// s1.push(NewNode(34))
	// s1.show()
	// s1.pop()
	// s1.pop()
	// s1.pop()
	// s1.pop()
	// s1.show()

	// --------Tree Traversal-------------------
	// b4 := NewTreeNode(1, nil, nil)
	// b5 := NewTreeNode(3, nil, nil)
	// b6 := NewTreeNode(15, nil, nil)
	// b7 := NewTreeNode(30, nil, nil)
	// b2 := NewTreeNode(2, b4, b5)
	// b3 := NewTreeNode(20, b6, b7)
	// b1 := NewTreeNode(10, b2, b3)
	//
	// t1 := NewBinaryTree(b1)
	// t1.traverse_tree_PRE_ORDER()
	// fmt.Println(t1.path)

	// t1.path = nil
	// t1.traverse_tree_IN_ORDER()
	// fmt.Println(t1.path)

	// t1.path = nil
	// t1.traverse_tree_POST_ORDER()
	// fmt.Println(t1.path)
	// needle := 30
	// fmt.Println(t1.BreadthFirstSearch(needle))
	// fmt.Println(t1.DepthFirstSearch(needle))
	//

	// -------- HEAP ---------------
	// m1 := NewMinHeap()
	// m1.insert(3)
	// m1.insert(5)
	// m1.insert(7)
	// m1.insert(10)
	// m1.show()
	// m1.insert(1)
	// m1.show()
	// m1.pop()
	// m1.show()
	// m1.pop()
	// m1.show()
	// m1.pop()
	// m1.show()

	// --- TRIE tree ----------
	// root := NewTrieNode()
	// t1 := NewTrie(root)
	// list := []string{"CAR", "DONE", "TRY", "CAT", "TRIE", "DO"}
	// for _, word := range list {
	// 	t1.insert(word)
	// }
	// t1.show()
	// fmt.Println(t1.hasWord("DO"))

	// ------- GRAPH: BFS on adjacency matrix
	// sample := sampleGraphData()
	// out := graphBFS(sample, 0, 6)
	// fmt.Println("Graph: BfS on matrix: ->", out)
	//
	// ------ GRAPH: DFS on adjacency list
	adjList1 := sampleGraphAdjacencyList()
	fmt.Println(adjList1.list)
	for _, subList := range adjList1.list {
		fmt.Println()
		for _, item := range subList {
			fmt.Println(item)
		}
	}
	fmt.Println("--------------------")

	g1 := NewGraphDFS(adjList1)

	fmt.Println("\n\nConstructed path to needle: ", g1.dfs(0, 6))
	fmt.Println("Expected Solution: [0 1 4 5 6]")
}
