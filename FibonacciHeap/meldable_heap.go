package FibonacciHeap

import (
	"fmt"
	"math/rand"
	"time"
)

// MeldableNode represents a node in the meldable heap
type MeldableNode struct {
	value int
	left  *MeldableNode
	right *MeldableNode
}

// MeldableHeap represents a meldable heap
type MeldableHeap struct {
	root *MeldableNode
}

// NewMeldableHeap creates a new meldable heap
func NewMeldableHeap() *MeldableHeap {
	return &MeldableHeap{}
}

// Meld merges two meldable heaps and returns the new root
func (h *MeldableHeap) Meld(h1, h2 *MeldableNode) *MeldableNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.value > h2.value {
		h1, h2 = h2, h1
	}
	// Randomly decide to meld with left or right child
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		h1.left = h.Meld(h1.left, h2)
	} else {
		h1.right = h.Meld(h1.right, h2)
	}
	return h1
}

// Insert inserts a new value into the meldable heap
func (h *MeldableHeap) Insert(value int) {
	newNode := &MeldableNode{value: value}
	h.root = h.Meld(h.root, newNode)
}

// ExtractMin extracts the minimum value from the meldable heap
func (h *MeldableHeap) ExtractMin() int {
	if h.root == nil {
		panic("Heap is empty")
	}
	minValue := h.root.value
	h.root = h.Meld(h.root.left, h.root.right)
	return minValue
}

// RunMeldableHeap demonstrates the functionality of the meldable heap
func RunMeldableHeap() {
	fmt.Println("Running Meldable Heap example")
	h := NewMeldableHeap()
	h.Insert(3)
	h.Insert(2)
	h.Insert(7)
	h.Insert(1)

	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 1
	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 2
	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 3
	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 7
}
