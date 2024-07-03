package main

import (
	"fmt"
	"os"
	"time"

	TwoSAT "DS-Projects/2sat"
	"DS-Projects/FibonacciHeap"
	"DS-Projects/Hash"
	"DS-Projects/MST"
	"DS-Projects/SegmentTree"
	"DS-Projects/Skiplist"
	"DS-Projects/TravelingSalesmanProblem"
	"DS-Projects/Trie"

	"github.com/fatih/color"
)

func main() {
	for {
		clearScreen()

		// Displaying options for running different topics
		color.Cyan("Welcome to the Data Structures and Algorithms CLI!")
		time.Sleep(1 * time.Second)
		color.Yellow("Select a topic to run the example:")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("1. Fibonacci Heap")
		fmt.Println("2. Skiplist")
		fmt.Println("3. Trie")
		fmt.Println("4. Traveling Salesman Problem")
		fmt.Println("5. Hash")
		fmt.Println("6. 2-SAT")
		fmt.Println("7. Segment Tree")
		fmt.Println("8. MST")
		fmt.Println("9. Exit")

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		if choice == 9 {
			color.Green("Exiting the CLI. Goodbye!")
			return
		}

		runExample(choice)
	}
}

func runExample(choice int) {
	for {
		switch choice {
		case 1:
			runFibonacciHeapExamples()
		case 2:
			runSkiplistExamples()
		case 3:
			runTrieExamples()
		case 4:
			runTSPExamples()
		case 5:
			runHashExamples()
		case 6:
			runTwoSATExamples()
		case 7:
			runSegmentTreeExamples()
		case 8:
			runMSTExamples()
		default:
			color.Red("Invalid choice")
			time.Sleep(1 * time.Second)
			return
		}

		if !showBackOption() {
			return
		}
	}
}

func runFibonacciHeapExamples() {
	clearScreen()
	color.Cyan("\nRunning Fibonacci Heap examples...")
	time.Sleep(1 * time.Second)
	FibonacciHeap.RunFibonacciHeap()
	time.Sleep(1 * time.Second)
	FibonacciHeap.RunMeldableHeap()
	time.Sleep(1 * time.Second)
	FibonacciHeap.RunRandomizedMeldableHeap()
	time.Sleep(1 * time.Second)
	FibonacciHeap.RunDijkstraFibonacciHeap()
}

func runSkiplistExamples() {
	clearScreen()
	color.Cyan("\nRunning Skiplist examples...")
	time.Sleep(1 * time.Second)
	Skiplist.RunSkiplist()
	time.Sleep(1 * time.Second)
	Skiplist.RunXORLinkedList()
	time.Sleep(1 * time.Second)
	Skiplist.RunLinkedList()
}

func runTrieExamples() {
	clearScreen()
	color.Cyan("\nRunning Trie examples...")
	time.Sleep(1 * time.Second)
	Trie.RunTrie()
	time.Sleep(1 * time.Second)
	Trie.RunXFastTrie()
	time.Sleep(1 * time.Second)
	Trie.RunYFastTrie()
	time.Sleep(1 * time.Second)
	Trie.RunPatternSearching()
}

func runTSPExamples() {
	clearScreen()
	color.Cyan("\nRunning Traveling Salesman Problem examples...")
	time.Sleep(1 * time.Second)
	TravelingSalesmanProblem.RunTSP()
	time.Sleep(1 * time.Second)
	TravelingSalesmanProblem.RunMST()
	time.Sleep(1 * time.Second)
	TravelingSalesmanProblem.RunSpecialCases()
}

func runHashExamples() {
	clearScreen()
	color.Cyan("\nRunning Hash examples...")
	time.Sleep(1 * time.Second)
	Hash.RunHash()
	time.Sleep(1 * time.Second)
	Hash.RunChainedHashTable()
	time.Sleep(1 * time.Second)
	Hash.RunMultiplicativeHashing()
	time.Sleep(1 * time.Second)
	Hash.RunLinearHashTable()
	time.Sleep(1 * time.Second)
	Hash.RunSHA256()
	time.Sleep(1 * time.Second)
	Hash.RunMerkleHashTree()
}

func runTwoSATExamples() {
	clearScreen()
	color.Cyan("\nRunning 2-SAT examples...")
	time.Sleep(1 * time.Second)
	TwoSAT.RunBooleanSatisfiability()
	time.Sleep(1 * time.Second)
	TwoSAT.RunTwoSAT()
}

func runSegmentTreeExamples() {
	clearScreen()
	color.Cyan("\nRunning Segment Tree examples...")
	time.Sleep(1 * time.Second)
	SegmentTree.RunSegmentTree()
	time.Sleep(1 * time.Second)
	SegmentTree.RunLazyPropagation()
	time.Sleep(1 * time.Second)
	SegmentTree.RunRangeQueries()
}

func runMSTExamples() {
	clearScreen()
	color.Cyan("\nRunning Minimum Spanning Tree examples...")
	time.Sleep(1 * time.Second)
	MST.RunMST()
	time.Sleep(1 * time.Second)
	MST.RunCutProperty()
	time.Sleep(1 * time.Second)
	MST.RunCycleProperty()
	time.Sleep(1 * time.Second)
	MST.RunDisjointSetUnion()
	time.Sleep(1 * time.Second)
	MST.RunFredmanTarjan()
}

func showBackOption() bool {
	fmt.Println("\n1. Back to main menu")
	fmt.Println("2. Exit")
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	if choice == 1 {
		return false
	} else if choice == 2 {
		color.Green("Exiting the CLI. Goodbye!")
		os.Exit(0)
	}
	return true
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
