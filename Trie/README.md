# Trie

## Topics Covered
- Trie
- X-fast Trie
- Y-fast Trie
- Pattern Searching

## Description
This section covers the implementation of Trie data structures and their variants. It includes examples and explanations for X-fast Trie, Y-fast Trie, and pattern searching using Trie.

## Files
- `trie.go`: Implementation of Trie.
- `x_fast_trie.go`: Implementation of X-fast Trie.
- `y_fast_trie.go`: Implementation of Y-fast Trie.
- `pattern_searching.go`: Implementation of pattern searching using Trie.

## Trie

### What is a Trie?
A Trie (pronounced "try") is a tree-like data structure that stores a dynamic set of strings. It is commonly used for searching words in a dictionary, autocomplete features, and spell checking.

### Key Operations
- **Insert**: Add a new string to the Trie.
- **Search**: Find a string in the Trie.
- **Delete**: Remove a string from the Trie.

### Advantages
Tries provide efficient search, insert, and delete operations, typically in O(L) time complexity, where L is the length of the string.

## X-fast Trie

### What is an X-fast Trie?
An X-fast Trie is a data structure that supports fast predecessor and successor queries. It is a variant of the Trie that uses hash tables to achieve efficient queries.

### Key Operations
- **Insert**: Add a new element to the Trie.
- **Predecessor**: Find the largest element smaller than a given value.
- **Successor**: Find the smallest element larger than a given value.

## Y-fast Trie

### What is a Y-fast Trie?
A Y-fast Trie is a further optimization of the X-fast Trie. It combines the properties of X-fast Tries and balanced binary search trees to achieve better space and time complexity.

### Key Operations
- **Insert**: Add a new element to the Trie.
- **Predecessor**: Find the largest element smaller than a given value.
- **Successor**: Find the smallest element larger than a given value.

## Pattern Searching

### What is Pattern Searching?
Pattern searching involves finding occurrences of a pattern string within a text. Tries can be used for efficient pattern matching, especially in multiple-pattern scenarios.

### Key Operations
- **Insert Pattern**: Add a new pattern to the Trie.
- **Search Pattern**: Search for a pattern in the text using the Trie.

## Usage
Each file contains a main function to demonstrate the functionality of the respective data structure.

## Implementation

### trie.go

```go
package Trie

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, found := node.children[ch]; !found {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, found := node.children[ch]; !found {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (t *Trie) Delete(word string) {
	t.deleteHelper(t.root, word, 0)
}

func (t *Trie) deleteHelper(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}
	if depth == len(word) {
		if node.isEnd {
			node.isEnd = false
		}
		return len(node.children) == 0
	}
	ch := rune(word[depth])
	if t.deleteHelper(node.children[ch], word, depth+1) {
		delete(node.children, ch)
		return len(node.children) == 0 && !node.isEnd
	}
	return false
}

func RunTrie() {
	fmt.Println("Running Trie example")
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println("Search for 'hello':", trie.Search("hello")) // Should return true
	trie.Delete("hello")
	fmt.Println("Search for 'hello' after deletion:", trie.Search("hello")) // Should return false
}
