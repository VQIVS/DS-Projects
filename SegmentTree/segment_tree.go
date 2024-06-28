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
