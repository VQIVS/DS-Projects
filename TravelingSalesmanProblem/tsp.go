package TravelingSalesmanProblem

import (
	"fmt"
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
