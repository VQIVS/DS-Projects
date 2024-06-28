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
