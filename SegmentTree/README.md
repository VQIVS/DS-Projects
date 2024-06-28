# Segment Tree

## Topics Covered
- Segment Tree
- Lazy Propagation
- Range Queries

## Description
This section covers the implementation of Segment Trees and related concepts such as Lazy Propagation and Range Queries.

## Files
- `segment_tree.go`: Implementation of Segment Tree.
- `lazy_propagation.go`: Implementation of Lazy Propagation.
- `range_queries.go`: Implementation of Range Queries using Segment Tree.

## Segment Tree

### What is a Segment Tree?
A Segment Tree is a tree data structure used for storing intervals or segments. It allows querying which of the stored segments contain a given point efficiently.

### Key Operations
- **Build**: Construct the segment tree.
- **Update**: Update elements in the array and reflect those changes in the tree.
- **Query**: Retrieve information about a segment of the array.

### Advantages
Segment Trees provide efficient solutions for range query problems and can handle updates and queries in logarithmic time.

## Lazy Propagation

### What is Lazy Propagation?
Lazy Propagation is an optimization technique for segment trees that helps defer updates to segments until necessary, improving efficiency.

## Range Queries

### What are Range Queries?
Range Queries involve querying information (such as sum, minimum, maximum) over a range of array indices. Segment Trees are well-suited for such tasks.

## Usage
Each file contains a main function to demonstrate the functionality of the respective data structure or algorithm.

## Implementation

### segment_tree.go

```go
package SegmentTree

import (
	"fmt"
	"math"
)

type SegmentTree struct {
	tree []int
	arr  []int
	size int
}

func NewSegmentTree(arr []int) *SegmentTree {
	size := len(arr)
	treeSize := int(2*math.Pow(2, math.Ceil(math.Log2(float64(size)))) - 1)
	tree := make([]int, treeSize)
	st := &SegmentTree{tree: tree, arr: arr, size: size}
	st.build(0, size-1, 0)
	return st
}

func (st *SegmentTree) build(start, end, node int) {
	if start == end {
		st.tree[node] = st.arr[start]
	} else {
		mid := (start + end) / 2
		st.build(start, mid, 2*node+1)
		st.build(mid+1, end, 2*node+2)
		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
	}
}

func (st *SegmentTree) update(index, value, start, end, node int) {
	if start == end {
		st.arr[index] = value
		st.tree[node] = value
	} else {
		mid := (start + end) / 2
		if index <= mid {
			st.update(index, value, start, mid, 2*node+1)
		} else {
			st.update(index, value, mid+1, end, 2*node+2)
		}
		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
	}
}

func (st *SegmentTree) query(l, r, start, end, node int) int {
	if r < start || l > end {
		return 0
	}
	if l <= start && r >= end {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftSum := st.query(l, r, start, mid, 2*node+1)
	rightSum := st.query(l, r, mid+1, end, 2*node+2)
	return leftSum + rightSum
}

func RunSegmentTree() {
	fmt.Println("Running Segment Tree example")
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)
	fmt.Println("Sum of values in given range (1, 3):", st.query(1, 3, 0, len(arr)-1, 0)) // Output: 15
	st.update(1, 10, 0, len(arr)-1, 0)
	fmt.Println("Updated sum of values in given range (1, 3):", st.query(1, 3, 0, len(arr)-1, 0)) // Output: 22
}
