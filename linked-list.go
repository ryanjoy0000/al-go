package main

import "fmt"

type LinkedStruct interface {
	show()
}

type LL struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LL {
	return &LL{}
}

func (lList *LL) show() {
	if lList.length > 0 {
		fmt.Print("HEAD->")
		fmt.Print(lList.head.value, "->")
		if lList.head.next != nil {
			n := lList.head.next
			for {
				fmt.Print(n.value, "->")
				if n.next != nil {
					n = n.next
				} else {
					break
				}
			}
		}
		fmt.Println("NULL")
	} else {
		fmt.Println("empty linked list...")
	}
}

func (lList *LL) insert(n *Node, afterWhichVal int) {
	if lList.length == 0 {
		lList.head = n
		lList.length++
		fmt.Println("Node inserted with val:", n.value)
	} else {
		ptr1 := lList.head
		for i := 0; i < lList.length; i++ {
			if ptr1.value == afterWhichVal {
				break
			} else {
				ptr1 = ptr1.next
			}
		}

		if ptr1 != nil {
			ptr2 := ptr1.next
			n.next = ptr2
			ptr1.next = n
			lList.length++
			fmt.Println("Node inserted with val:", n.value)
		}
	}
}

func (lList *LL) delete(n *Node) {
	// empty
	if lList.length == 0 {
		fmt.Println("empty linked list")
	} else if lList.length == 1 {
		lList.head = nil
		fmt.Println("node removed, linked list is now empty")
		lList.length--
	} else if lList.length == 2 {
		if lList.head.value == n.value {
			ptr1 := lList.head
			lList.head = ptr1.next
		} else {
			lList.head.next = nil
		}
		fmt.Println("node removed")
		lList.length--
	} else {
		// TODO - deletion to be implemented
	}
}
