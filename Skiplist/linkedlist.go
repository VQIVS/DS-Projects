package Skiplist

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (ll *LinkedList) Insert(value int) {
	newNode := &ListNode{value: value}
	newNode.next = ll.head
	ll.head = newNode
}

func (ll *LinkedList) Search(value int) *ListNode {
	current := ll.head
	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}
	return nil
}

func (ll *LinkedList) Delete(value int) {
	current := ll.head
	var prev *ListNode
	for current != nil && current.value != value {
		prev = current
		current = current.next
	}
	if current == nil {
		return
	}
	if prev == nil {
		ll.head = current.next
	} else {
		prev.next = current.next
	}
}

func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

func RunLinkedList() {
	fmt.Println("Running Linked List example")
	ll := &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)

	fmt.Print("Linked List: ")
	ll.PrintList() // Should print 4 3 2 1

	fmt.Println("Search for 3:", ll.Search(3)) // Should find 3
	ll.Delete(3)
	fmt.Println("Search for 3 after deletion:", ll.Search(3)) // Should not find 3

	fmt.Print("Linked List after deletion: ")
	ll.PrintList() // Should print 4 2 1
}
