package SegmentTree

import (
	"fmt"
)

func RunRangeQueries() {
	fmt.Println("Running Range Queries example")
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)
	fmt.Println("Sum of values in range (1, 3):", st.query(1, 3, 0, len(arr)-1, 0)) // Output: 15
	fmt.Println("Sum of values in range (2, 5):", st.query(2, 5, 0, len(arr)-1, 0)) // Output: 32
}
