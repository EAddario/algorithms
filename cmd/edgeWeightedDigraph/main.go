/******************************************************************************
 *  Execution:    go run cmd/edgeWeightedDigraph/main.go filename.txt
 *  Data files:   https://algs4.cs.princeton.edu/43mst/tinyEWG.txt
 *                https://algs4.cs.princeton.edu/43mst/mediumEWG.txt
 *                https://algs4.cs.princeton.edu/43mst/largeEWG.txt
 *
 *  An edge-weighted undirected graph, implemented using adjacency lists.
 *  Parallel edges and self-loops are permitted.
 *
 *  % go run cmd/edgeWeightedDigraph/main.go tinyEWD.txt
 * 8 vertices, 15 edges
 * 0:  0->2 0.26000  0->4 0.38000
 * 1:  1->3 0.29000
 * 2:  2->7 0.34000
 * 3:  3->6 0.52000
 * 4:  4->7 0.37000  4->5 0.35000
 * 5:  5->1 0.32000  5->7 0.28000  5->4 0.35000
 * 6:  6->4 0.93000  6->0 0.58000  6->2 0.40000
 * 7:  7->3 0.39000  7->5 0.28000
 *
 ******************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/EAddario/algorithms/pkg/algorithms"
	"github.com/EAddario/algorithms/pkg/stdin"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: edgeWeightedDigraph graph-file")
		os.Exit(0)
	}

	graph := algorithms.NewEdgeWeightedDigraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	fmt.Println(graph)
}
