package Skiplist

import (
	"fmt"
	"unsafe"
)

type XORNode struct {
	value int
	npx   *XORNode
}

type XORLinkedList struct {
	head *XORNode
	tail *XORNode
}

func NewXORNode(value int) *XORNode {
	return &XORNode{value: value}
}

func (ll *XORLinkedList) XOR(a, b *XORNode) *XORNode {
	return (*XORNode)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) ^ uintptr(unsafe.Pointer(b))))
}

func (ll *XORLinkedList) Insert(value int) {
	newNode := NewXORNode(value)
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.npx = ll.XOR(ll.tail.npx, newNode)
		newNode.npx = ll.tail
		ll.tail = newNode
	}
}

func (ll *XORLinkedList) PrintList() {
	curr := ll.head
	var prev *XORNode
	var next *XORNode

	for curr != nil {
		fmt.Printf("%d ", curr.value)
		next = ll.XOR(prev, curr.npx)
		prev = curr
		curr = next
	}
	fmt.Println()
}

func RunXORLinkedList() {
	fmt.Println("Running XOR Linked List example")
	ll := &XORLinkedList{}
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	fmt.Print("XOR Linked List: ")
	ll.PrintList() // Should print 10 20 30 40
}
