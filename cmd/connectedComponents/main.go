/*
  Execution:    go run cmd/connectedComponents/main.go filename.txt
  Data files:   https://algs4.cs.princeton.edu/41graph/tinyG.txt
                https://algs4.cs.princeton.edu/41graph/mediumG.txt
                https://algs4.cs.princeton.edu/41graph/largeG.txt

  Compute connected components using depth first search.
  Runs in O(E + V) time.

  % go run cmd/connectedComponents/main.go tinyG.txt
  3 components
  0 1 2 3 4 5 6
  7 8
  9 10 11 12

  % go run cmd/connectedComponents/main.go largeG.txt
  1 components
  0 1 2 3 4 5 6 7 8 9 10 ...

  Note: This implementation uses a recursive DFS. To avoid needing
        a potentially very large stack size, replace with a non-recurisve
        DFS ala NonrecursiveDFS.
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
		fmt.Println("Usage: connectedComponents graph-file")
		os.Exit(0)
	}

	graph := algorithms.NewGraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	cc := algorithms.NewCC(graph)
	fmt.Println(cc.Count(), " components")
	var components []*algorithms.Bag

	for i := 0; i < cc.Count(); i++ {
		components = append(components, algorithms.NewBag())
	}

	for v := 0; v < graph.V(); v++ {
		components[cc.ID(v)].Add(v)
	}

	for i := 0; i < cc.Count(); i++ {

		for _, v := range components[i].Slice() {
			fmt.Print(v, " ")
		}

		fmt.Println()
	}

}
