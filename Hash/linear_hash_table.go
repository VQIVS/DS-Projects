package Hash

import (
	"fmt"
)

type LinearHashTable struct {
	table []int
	size  int
}

func NewLinearHashTable(size int) *LinearHashTable {
	return &LinearHashTable{
		table: make([]int, size),
		size:  size,
	}
}

func (ht *LinearHashTable) hash(key int) int {
	return key % ht.size
}

func (ht *LinearHashTable) Insert(key int) {
	index := ht.hash(key)
	for ht.table[index] != 0 {
		index = (index + 1) % ht.size
	}
	ht.table[index] = key
}

func (ht *LinearHashTable) Search(key int) bool {
	index := ht.hash(key)
	for ht.table[index] != 0 {
		if ht.table[index] == key {
			return true
		}
		index = (index + 1) % ht.size
	}
	return false
}

func (ht *LinearHashTable) Delete(key int) {
	index := ht.hash(key)
	for ht.table[index] != 0 {
		if ht.table[index] == key {
			ht.table[index] = 0
			return
		}
		index = (index + 1) % ht.size
	}
}

func RunLinearHashTable() {
	fmt.Println("Running Linear Hash Table example")
	ht := NewLinearHashTable(10)
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
