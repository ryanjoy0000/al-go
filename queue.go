package main

import "fmt"

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	next  *Node
	value any
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewNode(val any) *Node {
	return &Node{
		value: val,
	}
}

func (q *Queue) Enqueue(n *Node) {
	if q.length > 0 {
		q.tail.next = n
		q.tail = n
	} else {
		q.head = n
		q.tail = n
	}
	// fmt.Println("Enqueued...")
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	var result interface{}
	if q.length > 0 {
		result = q.head.value
		tmpPtr := q.head
		q.head = q.head.next
		tmpPtr.next = nil
		// fmt.Println("Dequeued...")
		q.length--
	} else {
		fmt.Println("queue is empty")
	}
	return result
}

func (q *Queue) Show() {
	if q.length > 0 {
		fmt.Print(q.head.value, " -> ")
		if q.head.next != nil {
			n := q.head.next
			for i := 0; i < q.length; i++ {
				fmt.Print(n.value, " -> ")
				if n.next != nil {
					n = n.next
				} else {
					break
				}
			}
		}
	}
	fmt.Println("empty queue...")
}
