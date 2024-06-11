/*
   Execution:    go run cmd/directedDFS/main.go digraph.txt s
   Data files:   https://algs4.cs.princeton.edu/42digraph/tinyDG.txt
                 https://algs4.cs.princeton.edu/42digraph/mediumDG.txt
                 https://algs4.cs.princeton.edu/42digraph/largeDG.txt

   Determine single-source or multiple-source reachability in a digraph
   using depth first search.
   Runs in O(E + V) time.

   % go run cmd/directedDFS/main.go tinyDG.txt 1
   1

   % go run cmd/directedDFS/main.go tinyDG.txt 2
   0 1 2 3 4 5

   % go run cmd/directedDFS/main.go tinyDG.txt 1 2 6
   0 1 2 3 4 5 6 8 9 10 11 12
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/EAddario/algorithms/pkg/algorithms"
	"github.com/EAddario/algorithms/pkg/stdin"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: directedDFS graph-file vertex1 vertex2 vertexN")
		os.Exit(0)
	}

	graph := algorithms.NewDigraph(stdin.NewIn(os.Args[1], bufio.ScanWords))

	var sources []int

	for i := 2; i < len(os.Args); i++ {
		s, _ := strconv.Atoi(os.Args[i])
		sources = append(sources, s)
	}

	reachable := algorithms.NewDirectedDFSSources(graph, sources)

	for v := 0; v < graph.V(); v++ {
		if reachable.Marked(v) {
			fmt.Println(v, " ")
		}
	}

}
