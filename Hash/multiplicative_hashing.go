package Hash

import (
	"fmt"
)

type MultiplicativeHashTable struct {
	table []int
	size  int
	A     float64
}

func NewMultiplicativeHashTable(size int, A float64) *MultiplicativeHashTable {
	return &MultiplicativeHashTable{
		table: make([]int, size),
		size:  size,
		A:     A,
	}
}

func (ht *MultiplicativeHashTable) hash(key int) int {
	return int(float64(ht.size) * (float64(key)*ht.A - float64(int(float64(key)*ht.A))))
}

func (ht *MultiplicativeHashTable) Insert(key int) {
	index := ht.hash(key)
	ht.table[index] = key
}

func (ht *MultiplicativeHashTable) Search(key int) bool {
	index := ht.hash(key)
	return ht.table[index] == key
}

func (ht *MultiplicativeHashTable) Delete(key int) {
	index := ht.hash(key)
	if ht.table[index] == key {
		ht.table[index] = 0
	}
}

func RunMultiplicativeHashing() {
	fmt.Println("Running Multiplicative Hashing example")
	ht := NewMultiplicativeHashTable(10, 0.618033)
	ht.Insert(1)
	ht.Insert(11)
	ht.Insert(21)

	if ht.Search(11) {
		fmt.Println("Found value:", 11) // Should print 11
	} else {
		fmt.Println("Value not found")
	}

	ht.Delete(11)
	if ht.Search(11) {
		fmt.Println("Found value after deletion")
	} else {
		fmt.Println("Value not found after deletion")
	}
}
