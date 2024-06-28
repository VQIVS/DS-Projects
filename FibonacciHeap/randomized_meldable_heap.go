package FibonacciHeap

import (
	"fmt"
)

// RunRandomizedMeldableHeap demonstrates the functionality of a randomized meldable heap
func RunRandomizedMeldableHeap() {
	fmt.Println("Running Randomized Meldable Heap example")
	h := NewMeldableHeap()
	h.Insert(10)
	h.Insert(20)
	h.Insert(5)
	h.Insert(15)

	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 5
	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 10
	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 15
	fmt.Println("Minimum value:", h.ExtractMin()) // Should print 20
}
