# Minimum Spanning Tree (MST)

## Topics Covered
- Cut Property
- Cycle Property
- Disjoint Set Union
- Fredman-Tarjan Algorithm
- MST Algorithms

## Description
This section covers the implementation of Minimum Spanning Tree (MST) and related concepts such as Cut Property, Cycle Property, Disjoint Set Union, and Fredman-Tarjan Algorithm.

## Files
- `mst.go`: Implementation of MST algorithms.
- `cut_property.go`: Explanation and example of Cut Property.
- `cycle_property.go`: Explanation and example of Cycle Property.
- `disjoint_set_union.go`: Implementation of Disjoint Set Union (Union-Find).
- `fredman_tarjan.go`: Explanation and example of Fredman-Tarjan Algorithm.

## Minimum Spanning Tree

### What is a Minimum Spanning Tree?
A Minimum Spanning Tree (MST) of a graph is a subset of the edges that connects all the vertices together, without any cycles, and with the minimum possible total edge weight.

### Key Operations
- **Kruskal's Algorithm**: Finds the MST using a greedy approach.
- **Prim's Algorithm**: Another greedy algorithm to find the MST.

## Cut Property

### What is the Cut Property?
The Cut Property states that for any cut in a graph, the minimum weight edge that crosses the cut is part of the MST.

## Cycle Property

### What is the Cycle Property?
The Cycle Property states that for any cycle in the graph, the maximum weight edge in the cycle is not part of the MST.

## Disjoint Set Union (Union-Find)

### What is Disjoint Set Union?
Disjoint Set Union (Union-Find) is a data structure that keeps track of a partition of a set into disjoint subsets and supports union and find operations.

## Fredman-Tarjan Algorithm

### What is the Fredman-Tarjan Algorithm?
The Fredman-Tarjan Algorithm is a more advanced algorithm for finding MSTs, using a combination of Borůvka's algorithm and a heap-based approach for efficiency.

## Usage
Each file contains a main function to demonstrate the functionality of the respective algorithm or concept.

## Implementation

### mst.go

```go
package MST

import (
	"fmt"
	"sort"
)

type Edge struct {
	u, v, weight int
}

type Graph struct {
	edges []Edge
	n     int
}

func NewGraph(n int) *Graph {
	return &Graph{n: n}
}

func (g *Graph) AddEdge(u, v, weight int) {
	g.edges = append(g.edges, Edge{u, v, weight})
}

func (g *Graph) FindMST() int {
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	parent := make([]int, g.n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootY] = rootX
		}
	}

	mstWeight := 0
	for _, edge := range g.edges {
		u, v := find(edge.u), find(edge.v)
		if u != v {
			mstWeight += edge.weight
			union(u, v)
		}
	}
	return mstWeight
}

func RunMST() {
	fmt.Println("Running Minimum Spanning Tree example")
	g := NewGraph(4)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)

	mstWeight := g.FindMST()
	fmt.Println("The weight of the MST is", mstWeight)
}
