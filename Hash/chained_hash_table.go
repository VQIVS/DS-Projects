package Hash

import (
	"fmt"
)

type HashTable struct {
	table map[int]*LinkedList
	size  int
}

type LinkedList struct {
	head *Node
}

type Node struct {
	key   int
	value int
	next  *Node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		table: make(map[int]*LinkedList, size),
		size:  size,
	}
}

func (ht *HashTable) hash(key int) int {
	return key % ht.size
}

func (ht *HashTable) Insert(key, value int) {
	index := ht.hash(key)
	if ht.table[index] == nil {
		ht.table[index] = &LinkedList{}
	}
	ht.table[index].Insert(key, value)
}

func (ll *LinkedList) Insert(key, value int) {
	node := &Node{key: key, value: value}
	if ll.head == nil {
		ll.head = node
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (ht *HashTable) Search(key int) (int, bool) {
	index := ht.hash(key)
	if ht.table[index] == nil {
		return 0, false
	}
	return ht.table[index].Search(key)
}

func (ll *LinkedList) Search(key int) (int, bool) {
	current := ll.head
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

func (ht *HashTable) Delete(key int) {
	index := ht.hash(key)
	if ht.table[index] == nil {
		return
	}
	ht.table[index].Delete(key)
}

func (ll *LinkedList) Delete(key int) {
	if ll.head == nil {
		return
	}
	if ll.head.key == key {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.key != key {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

func RunChainedHashTable() {
	fmt.Println("Running Chained Hash Table example")
	ht := NewHashTable(10)
	ht.Insert(1, 100)
	ht.Insert(11, 200)
	ht.Insert(21, 300)

	value, found := ht.Search(11)
	if found {
		fmt.Println("Found value:", value) // Should print 200
	} else {
		fmt.Println("Value not found")
	}

	ht.Delete(11)
	_, found = ht.Search(11)
	if found {
		fmt.Println("Found value after deletion")
	} else {
		fmt.Println("Value not found after deletion")
	}
}
