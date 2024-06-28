package FibonacciHeap

import (
	"container/heap"
	"fmt"
)

type Graph struct {
	vertices int
	edges    map[int][]Edge
}

type Edge struct {
	to     int
	weight int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make(map[int][]Edge),
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.edges[from] = append(g.edges[from], Edge{to, weight})
}

func (g *Graph) Dijkstra(start int) []int {
	dist := make([]int, g.vertices)
	for i := range dist {
		dist[i] = int(^uint(0) >> 1) // Infinity
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Item).value
		for _, edge := range g.edges[u] {
			v, weight := edge.to, edge.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{value: v, priority: dist[v]})
			}
		}
	}

	return dist
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func RunDijkstraFibonacciHeap() {
	fmt.Println("Running Dijkstra using Fibonacci Heap example")
	g := NewGraph(5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 4, 3)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 4, 4)
	g.AddEdge(2, 3, 9)
	g.AddEdge(3, 2, 7)
	g.AddEdge(4, 1, 1)
	g.AddEdge(4, 2, 8)
	g.AddEdge(4, 3, 2)

	dist := g.Dijkstra(0)
	fmt.Println("Shortest distances from vertex 0:", dist)
}
