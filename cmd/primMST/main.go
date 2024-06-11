/******************************************************************************
 *  Execution:    go run cmd/primMST/main.go filename.txt
 *  Data files:   https://algs4.cs.princeton.edu/43mst/tinyEWG.txt
 *                https://algs4.cs.princeton.edu/43mst/mediumEWG.txt
 *                https://algs4.cs.princeton.edu/43mst/largeEWG.txt
 *
 *  Compute a minimum spanning forest using Prim's algorithm.
 *
 *  %  go run cmd/primMST/main.go tinyEWG.txt
 *  1-7 0.19000
 *  0-2 0.26000
 *  2-3 0.17000
 *  4-5 0.35000
 *  5-7 0.28000
 *  6-2 0.40000
 *  0-7 0.16000
 *  1.81000
 *
 *  % go run cmd/primMST/main.go mediumEWG.txt
 *  1-72   0.06506
 *  2-86   0.05980
 *  3-67   0.09725
 *  4-55   0.06425
 *  5-102  0.03834
 *  6-129  0.05363
 *  7-157  0.00516
 *  ...
 *  10.46351
 *
 *  % go run cmd/primMST/main.go largeEWG.txt
 *  ...
 *  647.66307
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
		fmt.Println("Usage: primMST file")
		os.Exit(0)
	}

	g := algorithms.NewEdgeWeightedGraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	mst := algorithms.NewPrimMST(g)

	for _, e := range mst.Edges() {
		fmt.Println(e)
	}

	fmt.Printf("%.5f\n", mst.Weight())
}
