# Hashing

## Topics Covered
- Hashing
- Chained Hash Table
- Multiplicative Hashing
- Linear Hash Table
- SHA-256
- Merkle Hash Tree

## Description
This section covers various hashing techniques and structures. It includes implementations and examples of chained hash tables, multiplicative hashing, linear hash tables, SHA-256, and Merkle hash trees.

## Files
- `hash.go`: General information and introduction to hashing.
- `chained_hash_table.go`: Implementation of a chained hash table.
- `multiplicative_hashing.go`: Implementation of multiplicative hashing.
- `linear_hash_table.go`: Implementation of a linear hash table.
- `sha256.go`: Implementation of the SHA-256 hashing algorithm.
- `merkle_hash_tree.go`: Implementation of a Merkle hash tree.

## Hashing

### What is Hashing?
Hashing is the process of converting an input (or 'key') into a fixed-size string of bytes. The output is typically a 'hash code' used for indexing in hash tables.

## Chained Hash Table

### What is a Chained Hash Table?
A chained hash table is a hash table that resolves collisions using linked lists. Each bucket of the hash table contains a linked list of elements that hash to the same index.

### Key Operations
- **Insert**: Add a new key-value pair to the hash table.
- **Search**: Find a value by its key.
- **Delete**: Remove a key-value pair from the hash table.

## Multiplicative Hashing

### What is Multiplicative Hashing?
Multiplicative hashing is a hash function that uses multiplication by a constant to distribute keys uniformly across the hash table.

## Linear Hash Table

### What is a Linear Hash Table?
A linear hash table uses linear probing to resolve collisions. When a collision occurs, it linearly searches the next available slot.

## SHA-256

### What is SHA-256?
SHA-256 (Secure Hash Algorithm 256-bit) is a cryptographic hash function that produces a 256-bit hash value. It is widely used in security applications and protocols.

## Merkle Hash Tree

### What is a Merkle Hash Tree?
A Merkle hash tree is a tree in which every leaf node is labeled with a data block, and every non-leaf node is labeled with the cryptographic hash of the labels of its child nodes. It is used for efficient and secure verification of data integrity.

## Usage
Each file contains a main function to demonstrate the functionality of the respective data structure or algorithm.

## Implementation

### hash.go

```go
package Hash

import (
	"fmt"
)

func RunHash() {
	fmt.Println("Hashing module includes:")
	fmt.Println("- Chained Hash Table")
	fmt.Println("- Multiplicative Hashing")
	fmt.Println("- Linear Hash Table")
	fmt.Println("- SHA-256")
	fmt.Println("- Merkle Hash Tree")
}
