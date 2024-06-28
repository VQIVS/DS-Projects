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
