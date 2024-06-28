package Trie

import (
	"fmt"
	"math"
)

type XFastTrieNode struct {
	children map[int]*XFastTrieNode
	isEnd    bool
}

type XFastTrie struct {
	root *XFastTrieNode
}

func NewXFastTrieNode() *XFastTrieNode {
	return &XFastTrieNode{children: make(map[int]*XFastTrieNode)}
}

func NewXFastTrie() *XFastTrie {
	return &XFastTrie{root: NewXFastTrieNode()}
}

func (t *XFastTrie) Insert(value int) {
	node := t.root
	for i := int(math.Floor(math.Log2(float64(value)))); i >= 0; i-- {
		bit := (value >> i) & 1
		if _, found := node.children[bit]; !found {
			node.children[bit] = NewXFastTrieNode()
		}
		node = node.children[bit]
	}
	node.isEnd = true
}

func (t *XFastTrie) Predecessor(value int) int {
	// Implement predecessor logic
	return -1
}

func (t *XFastTrie) Successor(value int) int {
	// Implement successor logic
	return -1
}

func RunXFastTrie() {
	fmt.Println("Running X-fast Trie example")
	xft := NewXFastTrie()
	xft.Insert(10)
	xft.Insert(15)
	xft.Insert(20)

	fmt.Println("Predecessor of 15:", xft.Predecessor(15)) // Implement the logic to return the predecessor
	fmt.Println("Successor of 15:", xft.Successor(15))     // Implement the logic to return the successor
}
