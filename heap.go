package main

import "fmt"

type MinHeap struct {
	data   []int
	length int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data:   make([]int, 0),
		length: 0,
	}
}

func (m *MinHeap) getParentPos(itemPos int) int {
	return (itemPos - 1) / 2
}

func (m *MinHeap) getLeftChildPos(itemPos int) int {
	return 2*itemPos + 1
}

func (m *MinHeap) getRightChildPos(itemPos int) int {
	return 2*itemPos + 2
}

func (m *MinHeap) insert(val int) {
	// 1 - insert at last pos
	m.data = append(m.data, val)

	// 2 - heapify up, from last pos
	m.heapifyUp(m.length)

	// 3 - increase length
	m.length++
}

func (m *MinHeap) heapifyDown(itemPos int) {
	leftChildPos := m.getLeftChildPos(itemPos)
	rightChildPos := m.getRightChildPos(itemPos)

	// BASE CASE: exit when exceeding length or no more children left
	if itemPos >= m.length || leftChildPos >= m.length || rightChildPos >= m.length {
		return
	}
	currentVal := m.data[itemPos]
	leftChildVal := m.data[leftChildPos]
	rightChildVal := m.data[rightChildPos]
	// compare with left and right child values
	if currentVal > leftChildVal && leftChildVal < rightChildVal {
		// swap with left child
		m.data[itemPos], m.data[leftChildPos] = m.data[leftChildPos], m.data[itemPos]
		m.heapifyDown(itemPos)
	} else if currentVal > rightChildVal && rightChildVal < leftChildVal {
		// swap with right child
		m.data[itemPos], m.data[rightChildPos] = m.data[rightChildPos], m.data[itemPos]
		m.heapifyDown(itemPos)
	}
}

func (m *MinHeap) heapifyUp(itemPos int) {
	// BASE case: exit when we reach root / top element
	if itemPos == 0 {
		return
	}

	parentPos := m.getParentPos(itemPos)
	// check if val < parent val
	if m.data[itemPos] < m.data[parentPos] {
		// swap
		m.data[itemPos], m.data[parentPos] = m.data[parentPos], m.data[itemPos]
		// RECURSION
		m.heapifyUp(parentPos)
	}
}

func (m *MinHeap) pop() int {
	// BASE CASE: exit when no elements
	if m.length == 0 {
		fmt.Println("nothing to pop")
		return -1
	}

	out := m.data[0]

	// handle for single element
	if m.length == 1 {
		m.data = nil
		fmt.Println("popping out: ", out)
		return out
	}

	// replace root by last element
	m.data[m.length-1], m.data[0] = m.data[0], m.data[m.length-1]

	// remove last element
	m.data = m.data[0 : m.length-1]

	// decrease length
	m.length--

	// RECURSION: heapifyDown
	m.heapifyDown(0)

	fmt.Println("popping out: ", out)
	return out
}

func (m *MinHeap) show() {
	fmt.Println("Heap with length: ", m.length)
	fmt.Println(m.data)
}
