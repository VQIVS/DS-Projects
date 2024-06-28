package SegmentTree

import (
	"fmt"
	"math"
)

type LazySegmentTree struct {
	tree []int
	lazy []int
	arr  []int
	size int
}

func NewLazySegmentTree(arr []int) *LazySegmentTree {
	size := len(arr)
	treeSize := int(2*math.Pow(2, math.Ceil(math.Log2(float64(size)))) - 1)
	tree := make([]int, treeSize)
	lazy := make([]int, treeSize)
	st := &LazySegmentTree{tree: tree, lazy: lazy, arr: arr, size: size}
	st.build(0, size-1, 0)
	return st
}

func (st *LazySegmentTree) build(start, end, node int) {
	if start == end {
		st.tree[node] = st.arr[start]
	} else {
		mid := (start + end) / 2
		st.build(start, mid, 2*node+1)
		st.build(mid+1, end, 2*node+2)
		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
	}
}

func (st *LazySegmentTree) updateRange(l, r, diff, start, end, node int) {
	if st.lazy[node] != 0 {
		st.tree[node] += (end - start + 1) * st.lazy[node]
		if start != end {
			st.lazy[2*node+1] += st.lazy[node]
			st.lazy[2*node+2] += st.lazy[node]
		}
		st.lazy[node] = 0
	}

	if start > end || start > r || end < l {
		return
	}

	if start >= l && end <= r {
		st.tree[node] += (end - start + 1) * diff
		if start != end {
			st.lazy[2*node+1] += diff
			st.lazy[2*node+2] += diff
		}
		return
	}

	mid := (start + end) / 2
	st.updateRange(l, r, diff, start, mid, 2*node+1)
	st.updateRange(l, r, diff, mid+1, end, 2*node+2)
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *LazySegmentTree) queryRange(l, r, start, end, node int) int {
	if st.lazy[node] != 0 {
		st.tree[node] += (end - start + 1) * st.lazy[node]
		if start != end {
			st.lazy[2*node+1] += st.lazy[node]
			st.lazy[2*node+2] += st.lazy[node]
		}
		st.lazy[node] = 0
	}

	if start > end || start > r || end < l {
		return 0
	}

	if start >= l && end <= r {
		return st.tree[node]
	}

	mid := (start + end) / 2
	leftSum := st.queryRange(l, r, start, mid, 2*node+1)
	rightSum := st.queryRange(l, r, mid+1, end, 2*node+2)
	return leftSum + rightSum
}

func RunLazyPropagation() {
	fmt.Println("Running Lazy Propagation example")
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewLazySegmentTree(arr)
	fmt.Println("Sum of values in given range (1, 3):", st.queryRange(1, 3, 0, len(arr)-1, 0)) // Output: 15
	st.updateRange(1, 3, 10, 0, len(arr)-1, 0)
	fmt.Println("Updated sum of values in given range (1, 3):", st.queryRange(1, 3, 0, len(arr)-1, 0)) // Output: 45
}
