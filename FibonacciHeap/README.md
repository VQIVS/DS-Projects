# Fibonacci Heap

## Topics Covered
- Fibonacci Heap
- Meldable Heap
- Randomized Meldable Heap
- Dijkstra using Fibonacci Heap

## Description
This section covers the implementation of Fibonacci Heaps and related data structures. It includes examples and explanations for meldable heaps and their randomized versions, as well as how to use a Fibonacci Heap in Dijkstra's algorithm.

## Files
- `fibonacci_heap.go`: Implementation of Fibonacci Heap.
- `meldable_heap.go`: Implementation of Meldable Heap.
- `randomized_meldable_heap.go`: Implementation of Randomized Meldable Heap.
- `dijkstra_fibonacci_heap.go`: Example of Dijkstra's algorithm using Fibonacci Heap.

## Fibonacci Heap

### What is a Fibonacci Heap?
A Fibonacci Heap is a data structure consisting of a collection of trees satisfying the minimum-heap property. It is particularly useful for implementing priority queues with efficient operations.

### Key Operations
- **Insert**: Add a new element to the heap.
- **Extract Min**: Remove and return the minimum element from the heap.
- **Decrease Key**: Decrease the value of a key in the heap.
- **Delete**: Remove an element from the heap.

### Advantages
Fibonacci Heaps have better amortized running times for some operations compared to other heap structures. For example, `Decrease Key` and `Delete` operations are more efficient.

## Meldable Heap

### What is a Meldable Heap?
A Meldable Heap is a type of heap that supports a `Meld` operation, which combines two heaps into one. The primary advantage is the ability to efficiently merge two priority queues.

### Key Operations
- **Meld**: Combine two heaps into one.
- **Insert**: Add a new element to the heap.
- **Extract Min**: Remove and return the minimum element from the heap.

## Randomized Meldable Heap

### What is a Randomized Meldable Heap?
A Randomized Meldable Heap is a variant of the Meldable Heap where the merging of heaps is done randomly. This randomness helps in maintaining a balanced structure over a series of operations.

### Key Operations
- **Meld**: Combine two heaps into one with random choices.
- **Insert**: Add a new element to the heap.
- **Extract Min**: Remove and return the minimum element from the heap.

## Dijkstra using Fibonacci Heap

### What is Dijkstra's Algorithm?
Dijkstra's Algorithm is a famous algorithm for finding the shortest paths between nodes in a graph. It is widely used in network routing and geographical mapping.

### How does Fibonacci Heap help?
Using a Fibonacci Heap can significantly improve the performance of Dijkstra's Algorithm, especially for graphs with many nodes and edges. The `Decrease Key` operation is particularly efficient with Fibonacci Heaps.

## Usage
Each file contains a main function to demonstrate the functionality of the respective data structure. Below are the implementations for each topic.

## Implementation

### fibonacci_heap.go

```go
package main

import (
    "fmt"
)

type Node struct {
    key      int
    degree   int
    marked   bool
    parent   *Node
    child    *Node
    left     *Node
    right    *Node
}

type FibonacciHeap struct {
    min   *Node
    count int
}

func NewNode(key int) *Node {
    node := &Node{
        key: key,
    }
    node.left = node
    node.right = node
    return node
}

func NewFibonacciHeap() *FibonacciHeap {
    return &FibonacciHeap{}
}

func (heap *FibonacciHeap) Insert(key int) {
    node := NewNode(key)
    if heap.min == nil {
        heap.min = node
    } else {
        heap.min.left.right = node
        node.left = heap.min.left
        node.right = heap.min
        heap.min.left = node
        if key < heap.min.key {
            heap.min = node
        }
    }
    heap.count++
}

func (heap *FibonacciHeap) ExtractMin() *Node {
    min := heap.min
    if min != nil {
        if min.child != nil {
            children := min.child
            for ok := true; ok; ok = (children != min.child) {
                next := children.right
                children.left.right = children.right
                children.right.left = children.left
                children.left = heap.min
                children.right = heap.min.right
                heap.min.right = children
                children.parent = nil
                children = next
            }
        }
        min.left.right = min.right
        min.right.left = min.left
        if min == min.right {
            heap.min = nil
        } else {
            heap.min = min.right
            heap.Consolidate()
        }
        heap.count--
    }
    return min
}

func (heap *FibonacciHeap) Consolidate() {
    // Consolidation logic to maintain heap structure
}

func main() {
    heap := NewFibonacciHeap()
    heap.Insert(3)
    heap.Insert(2)
    heap.Insert(1)

    fmt.Println("Minimum value: ", heap.ExtractMin().key) // Should print 1
}
