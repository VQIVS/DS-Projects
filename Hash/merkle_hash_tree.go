package Hash

import (
	"crypto/sha256"
	"fmt"
)

type MerkleTree struct {
	root *MerkleNode
}

type MerkleNode struct {
	left  *MerkleNode
	right *MerkleNode
	hash  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := &MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.hash = hash[:]
	} else {
		prevHashes := append(left.hash, right.hash...)
		hash := sha256.Sum256(prevHashes)
		node.hash = hash[:]
	}
	node.left = left
	node.right = right
	return node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for len(nodes) > 1 {
		var newLevel []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
				newLevel = append(newLevel, *node)
			} else {
				newLevel = append(newLevel, nodes[i])
			}
		}
		nodes = newLevel
	}

	tree := MerkleTree{&nodes[0]}
	return &tree
}

func RunMerkleHashTree() {
	fmt.Println("Running Merkle Hash Tree example")
	data := [][]byte{
		[]byte("Hello, World!"),
		[]byte("Merkle Tree Example"),
		[]byte("Hash Function"),
		[]byte("Cryptographic Hash"),
	}

	tree := NewMerkleTree(data)
	fmt.Printf("Root Hash: %x\n", tree.root.hash)
}
