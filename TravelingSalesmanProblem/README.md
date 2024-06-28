# Traveling Salesman Problem

## Topics Covered
- Traveling Salesman Problem (TSP)
- Minimum Spanning Tree (MST)
- Special Cases of TSP

## Description
This section covers the implementation of the Traveling Salesman Problem (TSP) and related concepts such as Minimum Spanning Tree (MST) and special cases of TSP.

## Files
- `tsp.go`: Implementation of TSP.
- `mst.go`: Implementation of MST.
- `special_cases.go`: Implementation of special cases of TSP.

## Traveling Salesman Problem

### What is the Traveling Salesman Problem?
The Traveling Salesman Problem (TSP) is an optimization problem where a salesman is required to visit a list of cities, each exactly once, and return to the origin city, with the goal of minimizing the total travel distance.

### Key Operations
- **Exact Algorithms**: Find the exact solution to the TSP.
- **Approximation Algorithms**: Find a solution close to the optimal for the TSP.

### Advantages
Solving the TSP has numerous applications in logistics, planning, and microchip manufacturing.

## Minimum Spanning Tree (MST)

### What is a Minimum Spanning Tree?
A Minimum Spanning Tree (MST) of a graph is a subset of the edges that connects all the vertices together, without any cycles, and with the minimum possible total edge weight.

### Key Operations
- **Kruskal's Algorithm**: Finds the MST using a greedy approach.
- **Prim's Algorithm**: Another greedy algorithm to find the MST.

## Special Cases of TSP

### What are Special Cases of TSP?
Special cases of TSP include scenarios where certain properties of the graph make the problem easier to solve or provide approximation algorithms.

## Usage
Each file contains a main function to demonstrate the functionality of the respective algorithm.

## Implementation

### tsp.go

```go
package TravelingSalesmanProblem

import (
	"fmt"
	"math"
)

const maxInt = int(^uint(0) >> 1)

func tsp(graph [][]int, visited []bool, currPos, n, count, cost int) int {
	if count == n && graph[currPos][0] > 0 {
		return cost + graph[currPos][0]
	}
	minCost := maxInt
	for i := 0; i < n; i++ {
		if !visited[i] && graph[currPos][i] > 0 {
			visited[i] = true
			tempCost := tsp(graph, visited, i, n, count+1, cost+graph[currPos][i])
			if tempCost < minCost {
				minCost = tempCost
			}
			visited[i] = false
		}
	}
	return minCost
}

func RunTSP() {
	fmt.Println("Running Traveling Salesman Problem example")
	graph := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}
	n := len(graph)
	visited := make([]bool, n)
	visited[0] = true
	result := tsp(graph, visited, 0, n, 1, 0)
	fmt.Println("The minimum cost is", result)
}
