package MST

import "fmt"

type DisjointSetUnion struct {
	parent, rank []int
}

func NewDisjointSetUnion(size int) *DisjointSetUnion {
	dsu := &DisjointSetUnion{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range dsu.parent {
		dsu.parent[i] = i
	}
	return dsu
}

func (dsu *DisjointSetUnion) Find(u int) int {
	if dsu.parent[u] != u {
		dsu.parent[u] = dsu.Find(dsu.parent[u])
	}
	return dsu.parent[u]
}

func (dsu *DisjointSetUnion) Union(u, v int) {
	rootU := dsu.Find(u)
	rootV := dsu.Find(v)
	if rootU != rootV {
		if dsu.rank[rootU] > dsu.rank[rootV] {
			dsu.parent[rootV] = rootU
		} else if dsu.rank[rootU] < dsu.rank[rootV] {
			dsu.parent[rootU] = rootV
		} else {
			dsu.parent[rootV] = rootU
			dsu.rank[rootU]++
		}
	}
}

func RunDisjointSetUnion() {
	fmt.Println("Running Disjoint Set Union (Union-Find) example")
	dsu := NewDisjointSetUnion(5)
	dsu.Union(0, 1)
	dsu.Union(1, 2)
	dsu.Union(3, 4)
	fmt.Println("Find(0):", dsu.Find(0)) // Should print 0
	fmt.Println("Find(1):", dsu.Find(1)) // Should print 0
	fmt.Println("Find(2):", dsu.Find(2)) // Should print 0
	fmt.Println("Find(3):", dsu.Find(3)) // Should print 3
	fmt.Println("Find(4):", dsu.Find(4)) // Should print 3
}
