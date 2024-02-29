package main

import "fmt"

type Stack struct {
	top    *Node
	length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(n *Node) {
	if s.length == 0 {
		s.top = n
	} else {
		n.next = s.top
		s.top = n
	}
	fmt.Println("PUSH...")
	s.length++
}

func (s *Stack) show() {
	if s.length > 0 {
		fmt.Println("--STACK--")
		fmt.Println(s.top.value)
		if s.top.next != nil {
			n := s.top.next
			for n.next != nil {
				fmt.Println(n.value)
				n = n.next
			}
		}
		fmt.Println("---")
	} else {
		fmt.Println("empty stack...")
	}
}

func (s *Stack) pop() {
	if s.length > 0 {
		s.top = s.top.next
		fmt.Println("POP...")
		s.length--
	} else {
		fmt.Println("empty stack...")
	}
}
