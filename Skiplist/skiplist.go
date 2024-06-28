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
