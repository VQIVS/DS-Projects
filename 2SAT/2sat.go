package TwoSAT

import (
	"fmt"
)

type TwoSAT struct {
	n          int
	adj        [][]int
	adjInv     [][]int
	visited    []bool
	order      []int
	comp       []int
	assignment []bool
}

func NewTwoSAT(n int) *TwoSAT {
	return &TwoSAT{
		n:          n,
		adj:        make([][]int, 2*n),
		adjInv:     make([][]int, 2*n),
		visited:    make([]bool, 2*n),
		comp:       make([]int, 2*n),
		assignment: make([]bool, n),
	}
}

func (ts *TwoSAT) addImplication(u, v int) {
	ts.adj[u] = append(ts.adj[u], v)
	ts.adjInv[v] = append(ts.adjInv[v], u)
}

func (ts *TwoSAT) addClause(u, v int) {
	ts.addImplication(u^1, v)
	ts.addImplication(v^1, u)
}

func (ts *TwoSAT) dfs1(v int) {
	ts.visited[v] = true
	for _, u := range ts.adj[v] {
		if !ts.visited[u] {
			ts.dfs1(u)
		}
	}
	ts.order = append(ts.order, v)
}

func (ts *TwoSAT) dfs2(v, cl int) {
	ts.comp[v] = cl
	for _, u := range ts.adjInv[v] {
		if ts.comp[u] == -1 {
			ts.dfs2(u, cl)
		}
	}
}

func (ts *TwoSAT) solve() bool {
	for i := 0; i < 2*ts.n; i++ {
		if !ts.visited[i] {
			ts.dfs1(i)
		}
	}
	for i := 0; i < 2*ts.n; i++ {
		ts.comp[i] = -1
	}
	cl := 0
	for i := 0; i < 2*ts.n; i++ {
		v := ts.order[2*ts.n-1-i]
		if ts.comp[v] == -1 {
			ts.dfs2(v, cl)
			cl++
		}
	}
	for i := 0; i < ts.n; i++ {
		if ts.comp[2*i] == ts.comp[2*i+1] {
			return false
		}
		ts.assignment[i] = ts.comp[2*i] > ts.comp[2*i+1]
	}
	return true
}

func RunTwoSAT() {
	fmt.Println("Running 2-SAT example")
	n := 3
	twoSAT := NewTwoSAT(n)

	// Example 2-SAT problem: (x1 OR x2) AND (NOT x1 OR x3) AND (NOT x2 OR NOT x3)
	twoSAT.addClause(0*2+1, 1*2)   // x1 OR x2
	twoSAT.addClause(0*2+1, 2*2+1) // NOT x1 OR x3
	twoSAT.addClause(1*2+1, 2*2+1) // NOT x2 OR NOT x3

	if twoSAT.solve() {
		fmt.Println("Satisfiable")
		for i := 0; i < n; i++ {
			fmt.Printf("x%d: %t\n", i+1, twoSAT.assignment[i])
		}
	} else {
		fmt.Println("Not Satisfiable")
	}
}
