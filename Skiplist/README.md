# Skiplist

## Topics Covered
- Skiplist
- XOR-Linkedlist
- Linked List

## Description
This section covers the implementation of Skiplists and related data structures. It includes examples and explanations for XOR-Linkedlist and basic Linked List.

## Files
- `skiplist.go`: Implementation of Skiplist.
- `xor_linkedlist.go`: Implementation of XOR-Linkedlist.
- `linkedlist.go`: Implementation of basic Linked List.

## Skiplist

### What is a Skiplist?
A Skiplist is a probabilistic data structure that allows fast search, insertion, and deletion operations. It consists of multiple levels of linked lists, with each level skipping over a certain number of elements, allowing for efficient traversal.

### Key Operations
- **Insert**: Add a new element to the skiplist.
- **Search**: Find an element in the skiplist.
- **Delete**: Remove an element from the skiplist.

### Advantages
Skiplists provide a balance between the simplicity of linked lists and the efficiency of balanced trees. They are easy to implement and provide good average-case performance.

## XOR-Linkedlist

### What is an XOR-Linkedlist?
An XOR-Linkedlist is a memory-efficient doubly linked list where each node stores the XOR of the addresses of the previous and next nodes. This reduces the memory overhead of storing two pointers per node.

### Key Operations
- **Insert**: Add a new element to the list.
- **Search**: Find an element in the list.
- **Delete**: Remove an element from the list.

### Advantages
XOR-Linkedlists reduce memory usage compared to traditional doubly linked lists, making them useful in memory-constrained environments.

## Linked List

### What is a Linked List?
A Linked List is a linear data structure where each element is a separate object, called a node, which contains a reference (or link) to the next node in the sequence.

### Key Operations
- **Insert**: Add a new element to the list.
- **Search**: Find an element in the list.
- **Delete**: Remove an element from the list.

### Advantages
Linked Lists provide dynamic memory allocation and efficient insertions and deletions compared to arrays.

## Usage
Each file contains a main function to demonstrate the functionality of the respective data structure.

## Implementation

### skiplist.go

```go
package Skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLevel = 16

type Node struct {
	value int
	next  []*Node
}

type Skiplist struct {
	header *Node
	level  int
}

func NewNode(value, level int) *Node {
	return &Node{
		value: value,
		next:  make([]*Node, level),
	}
}

func NewSkiplist() *Skiplist {
	rand.Seed(time.Now().UnixNano())
	return &Skiplist{
		header: NewNode(0, maxLevel),
		level:  1,
	}
}

func (sl *Skiplist) randomLevel() int {
	level := 1
	for rand.Float32() < 0.5 && level < maxLevel {
		level++
	}
	return level
}

func (sl *Skiplist) Insert(value int) {
	update := make([]*Node, maxLevel)
	current := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}
	level := sl.randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}
	newNode := NewNode(value, level)
	for i := 0; i < level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func (sl *Skiplist) Search(value int) *Node {
	current := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
	}
	current = current.next[0]
	if current != nil && current.value == value {
		return current
	}
	return nil
}

func (sl *Skiplist) Delete(value int) {
	update := make([]*Node, maxLevel)
	current := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}
	current = current.next[0]
	if current != nil && current.value == value {
		for i := 0; i < sl.level; i++ {
			if update[i].next[i] != current {
				break
			}
			update[i].next[i] = current.next[i]
		}
		for sl.level > 1 && sl.header.next[sl.level-1] == nil {
			sl.level--
		}
	}
}

func RunSkiplist() {
	fmt.Println("Running Skiplist example")
	sl := NewSkiplist()
	sl.Insert(3)
	sl.Insert(6)
	sl.Insert(7)
	sl.Insert(9)
	sl.Insert(12)
	sl.Insert(19)
	sl.Insert(17)
	sl.Insert(26)
	sl.Insert(21)
	sl.Insert(25)

	fmt.Println("Search for 19:", sl.Search(19)) // Should find 19
	sl.Delete(19)
	fmt.Println("Search for 19 after deletion:", sl.Search(19)) // Should not find 19
}
