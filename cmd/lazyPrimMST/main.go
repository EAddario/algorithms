/******************************************************************************
 *  Execution:    go run cmd/lazyPrimMST/main.go filename.txt
 *  Data files:   https://algs4.cs.princeton.edu/43mst/tinyEWG.txt
 *                https://algs4.cs.princeton.edu/43mst/mediumEWG.txt
 *                https://algs4.cs.princeton.edu/43mst/largeEWG.txt
 *
 *  Compute a minimum spanning forest using a lazy version of Prim's
 *  algorithm.
 *
 *  %  go run cmd/lazyPrimMST/main.go tinyEWG.txt
 *  0-7 0.16000
 *  1-7 0.19000
 *  0-2 0.26000
 *  2-3 0.17000
 *  5-7 0.28000
 *  4-5 0.35000
 *  6-2 0.40000
 *  1.81000
 *
 *  % go run cmd/lazyPrimMST/main.go mediumEWG.txt
 *  0-225   0.02383
 *  49-225  0.03314
 *  44-49   0.02107
 *  44-204  0.01774
 *  49-97   0.03121
 *  202-204 0.04207
 *  176-202 0.04299
 *  176-191 0.02089
 *  68-176  0.04396
 *  58-68   0.04795
 *  10.46351
 *
 *  % go run cmd/lazyPrimMST/main.go largeEWG.txt
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
		fmt.Println("Usage: lazyPrimMST file")
		os.Exit(0)
	}

	g := algorithms.NewEdgeWeightedGraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	mst := algorithms.NewLazyPrimMST(g)

	for _, e := range mst.Edges() {
		fmt.Println(e)
	}

	fmt.Printf("%.5f\n", mst.Weight())
}
