/*
  Execution:    go run cmd/directedCycle/main.go input.txt
  Data files:   https://algs4.cs.princeton.edu/42digraph/tinyDG.txt
                https://algs4.cs.princeton.edu/42digraph/tinyDAG.txt

  Finds a directed cycle in a digraph.
  Runs in O(E + V) time.

  % go run cmd/directedCycle/main.go tinyDG.txt
  Directed cycle: 3 5 4 3

  %  go run cmd/directedCycle/main.go tinyDAG.txt
  No directed cycle

*/

package main

import (
	"bufio"
	"fmt"
	"os"

	"../../algorithms"
	"../../stdin"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: directedCycle graph-file")
		os.Exit(0)
	}

	graph := algorithms.NewDigraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	finder := algorithms.NewDirectedCycle(graph)

	if finder.HasCycle() {

		for _, v := range finder.Cycle().Slice() {
			fmt.Print(v, " ")
		}

		fmt.Println()
	} else {
		fmt.Println("Graph is acyclic")
	}

}
