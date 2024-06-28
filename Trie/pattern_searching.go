package Trie

import "fmt"

type PatternTrieNode struct {
	children map[rune]*PatternTrieNode
	isEnd    bool
}

type PatternTrie struct {
	root *PatternTrieNode
}

func NewPatternTrieNode() *PatternTrieNode {
	return &PatternTrieNode{children: make(map[rune]*PatternTrieNode)}
}

func NewPatternTrie() *PatternTrie {
	return &PatternTrie{root: NewPatternTrieNode()}
}

func (t *PatternTrie) InsertPattern(pattern string) {
	node := t.root
	for _, ch := range pattern {
		if _, found := node.children[ch]; !found {
			node.children[ch] = NewPatternTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *PatternTrie) SearchPattern(text string) bool {
	node := t.root
	for _, ch := range text {
		if _, found := node.children[ch]; !found {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func RunPatternSearching() {
	fmt.Println("Running Pattern Searching using Trie example")
	trie := NewPatternTrie()
	trie.InsertPattern("pattern")
	trie.InsertPattern("search")

	fmt.Println("Search for 'pattern':", trie.SearchPattern("pattern"))   // Should return true
	fmt.Println("Search for 'notfound':", trie.SearchPattern("notfound")) // Should return false
}
