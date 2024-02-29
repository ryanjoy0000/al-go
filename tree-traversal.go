package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

type BinaryTree struct {
	head   *TreeNode
	path   []int
	height int
}

func NewBinaryTree(head *TreeNode) *BinaryTree {
	return &BinaryTree{
		head:   head,
		path:   make([]int, 0),
		height: 0,
	}
}

func NewTreeNode(value int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		left:  left,
		right: right,
		value: value,
	}
}

// ------------------------------------------------------------
func (t *BinaryTree) traverse_tree_PRE_ORDER() {
	t.walk_PRE_ORDER(t.head)
}

func (t *BinaryTree) walk_PRE_ORDER(current *TreeNode) {
	// BASE CASE
	if current == nil {
		return
	}

	// 1. visit node
	t.path = append(t.path, current.value)

	// 2. recurse - go left
	t.walk_PRE_ORDER(current.left)

	// 3. recurse - go right
	t.walk_PRE_ORDER(current.right)
}

// ------------------------------------------------------------
func (t *BinaryTree) traverse_tree_IN_ORDER() {
	t.walk_IN_ORDER(t.head)
}

func (t *BinaryTree) walk_IN_ORDER(current *TreeNode) {
	// BASE CASE
	if current == nil {
		return
	}

	// 2. recurse - go left
	t.walk_IN_ORDER(current.left)

	// 1. visit node
	t.path = append(t.path, current.value)

	// 3. recurse - go right
	t.walk_IN_ORDER(current.right)
}

// ------------------------------------------------------------

func (t *BinaryTree) traverse_tree_POST_ORDER() {
	t.walk_POST_ORDER(t.head)
}

func (t *BinaryTree) walk_POST_ORDER(current *TreeNode) {
	// BASE CASE
	if current == nil {
		return
	}

	// 2. recurse - go left
	t.walk_POST_ORDER(current.left)

	// 3. recurse - go right
	t.walk_POST_ORDER(current.right)

	// 1. visit node
	t.path = append(t.path, current.value)
}

// ---------------------------------------------------------
// Breadth first search
func (t *BinaryTree) BreadthFirstSearch(needle int) bool {
	fmt.Println("-----------------")
	fmt.Println("searching by breadth first...")
	q1 := NewQueue()
	if t.head.value == needle {
		return true
	}
	q1.Enqueue(NewNode(t.head))

	for q1.length > 0 {
		next := q1.Dequeue()
		fmt.Println("TreeNode: ", next.(*TreeNode).value)
		if next.(*TreeNode).value == needle {
			fmt.Println("needle found!")
			return true
		}
		if next.(*TreeNode).left != nil {
			q1.Enqueue(NewNode(next.(*TreeNode).left))
		}
		if next.(*TreeNode).right != nil {
			q1.Enqueue(NewNode(next.(*TreeNode).right))
		}
	}

	return false
}

func (t *BinaryTree) compareNodes(node1, node2 *TreeNode) bool {
	// BASE conditions
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.value == node2.value {
		return true
	}

	// Recursion
	return t.compareNodes(node1.left, node2.left) && t.compareNodes(node1.right, node2.right)
}

func (t *BinaryTree) DepthFirstSearch(needle int) bool {
	fmt.Println("-----------------")
	fmt.Println("searching by depth first...")
	return find(t.head, needle)
}

func find(node *TreeNode, needle int) bool {
	// BASE conditions
	// 1 - node does not exist
	if node == nil {
		return false
	}
	fmt.Println("TreeNode: ", node.value)
	// 2 - value found
	if node.value == needle {
		fmt.Println("needle found!")
		return true
	}
	// Recursion
	// GO right
	if needle > node.value {
		return find(node.right, needle)
	}
	// GO left
	return find(node.left, needle)
}
