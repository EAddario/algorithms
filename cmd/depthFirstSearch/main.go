/*
  Execution:    go run cmd/depthFirstSearch/main.go filename.txt s
  Data files:   https: // algs4.cs.princeton.edu / 41graph / tinyG.txt
                https: // algs4.cs.princeton.edu / 41graph / mediumG.txt

  Run depth first search on an undirected graph.
  Runs in O(E + V) time.

 % go run cmd/depthFirstSearch/main.go tinyG.txt 0
  0 1 2 3 4 5 6
  NOT connected

 % go run cmd/depthFirstSearch/main.go tinyG.txt 9
  9 10 11 12
  NOT connected
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../../algorithms"
	"../../stdin"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: depthFirstSearch graph-file vertex")
		os.Exit(0)
	}

	graph := algorithms.NewGraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	s, _ := strconv.Atoi(os.Args[2])
	search := algorithms.NewDepthFirstSearch(graph, s)

	for v := 0; v < graph.V(); v++ {
		if search.Marked(v) {
			fmt.Println(v, " ")
		}
	}

	if search.Count() == graph.V() {
		fmt.Println("connected")
	} else {
		fmt.Println("not connected")
	}

}
