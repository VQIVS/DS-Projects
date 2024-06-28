package FibonacciHeap

import (
	"fmt"
)

// FibonacciNode represents a node in the Fibonacci heap
type FibonacciNode struct {
	value  int
	parent *FibonacciNode
	child  *FibonacciNode
	left   *FibonacciNode
	right  *FibonacciNode
	degree int
	marked bool
}

// FibonacciHeap represents a Fibonacci heap
type FibonacciHeap struct {
	minNode *FibonacciNode
	size    int
}

// NewFibonacciHeap creates a new Fibonacci heap
func NewFibonacciHeap() *FibonacciHeap {
	h := &FibonacciHeap{}
	h.minNode = nil
	h.size = 0
	return h
}

// Insert inserts a new value into the Fibonacci heap
func (h *FibonacciHeap) Insert(value int) {
	newNode := &FibonacciNode{value: value}
	newNode.left = newNode
	newNode.right = newNode
	h.minNode = h.mergeLists(h.minNode, newNode)
	h.size++
}

// ExtractMin extracts the minimum value from the Fibonacci heap
func (h *FibonacciHeap) ExtractMin() (int, error) {
	if h.minNode == nil {
		return 0, fmt.Errorf("heap is empty")
	}

	minNode := h.minNode
	if minNode.child != nil {
		children := minNode.child
		for {
			children.parent = nil
			children = children.right
			if children == minNode.child {
				break
			}
		}
		h.minNode = h.mergeLists(h.minNode, minNode.child)
	}

	h.removeNode(minNode)
	if minNode == minNode.right {
		h.minNode = nil
	} else {
		h.minNode = minNode.right
		h.consolidate()
	}
	h.size--
	return minNode.value, nil
}

// mergeLists merges two doubly linked lists and returns the new minimum node
func (h *FibonacciHeap) mergeLists(a, b *FibonacciNode) *FibonacciNode {
	if a == nil && b == nil {
		return nil
	} else if a == nil {
		return b
	} else if b == nil {
		return a
	} else {
		aRight := a.right
		a.right = b.right
		a.right.left = a
		b.right = aRight
		b.right.left = b
		if a.value < b.value {
			return a
		}
		return b
	}
}

// removeNode removes a node from the doubly linked list
func (h *FibonacciHeap) removeNode(node *FibonacciNode) {
	node.left.right = node.right
	node.right.left = node.left
	node.left = node
	node.right = node
}

// consolidate consolidates the heap by merging trees of the same degree
func (h *FibonacciHeap) consolidate() {
	degreeTable := make(map[int]*FibonacciNode)
	nodes := []*FibonacciNode{}

	for node := h.minNode; len(nodes) == 0 || node != nodes[0]; node = node.right {
		nodes = append(nodes, node)
	}

	for _, node := range nodes {
		degree := node.degree
		for degreeTable[degree] != nil {
			other := degreeTable[degree]
			if node.value > other.value {
				node, other = other, node
			}
			h.linkNodes(other, node)
			degreeTable[degree] = nil
			degree++
		}
		degreeTable[degree] = node
	}

	h.minNode = nil
	for _, node := range degreeTable {
		if node != nil {
			h.minNode = h.mergeLists(h.minNode, node)
		}
	}
}

// linkNodes links two nodes together in the Fibonacci heap
func (h *FibonacciHeap) linkNodes(child, parent *FibonacciNode) {
	h.removeNode(child)
	child.left = child
	child.right = child
	child.parent = parent
	parent.child = h.mergeLists(parent.child, child)
	parent.degree++
	child.marked = false
}

// RunFibonacciHeap demonstrates the functionality of the Fibonacci heap
func RunFibonacciHeap() {
	fmt.Println("Running Fibonacci Heap example")
	h := NewFibonacciHeap()
	h.Insert(3)
	h.Insert(2)
	h.Insert(7)
	h.Insert(1)

	for h.size > 0 {
		minValue, err := h.ExtractMin()
		if err != nil {
			fmt.Println("Error extracting min:", err)
			return
		}
		fmt.Println("Minimum value:", minValue)
	}
}
