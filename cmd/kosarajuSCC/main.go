/******************************************************************************
 *  Execution:    go run cmd/kosarajuSCC/main.go filename.txt
 *  Data files:   https://algs4.cs.princeton.edu/42digraph/tinyDG.txt
 *                https://algs4.cs.princeton.edu/42digraph/mediumDG.txt
 *                https://algs4.cs.princeton.edu/42digraph/largeDG.txt
 *
 *  Compute the strongly-connected components of a digraph using the
 *  Kosaraju-Sharir algorithm.
 *
 *  Runs in O(E + V) time.
 *
 *  % go run cmd/kosarajuSCC/main.go tinyDG.txt
 *  5 strong components
 *  1
 *  0 2 3 4 5
 *  9 10 11 12
 *  6 8
 *  7
 *
 *  % go run cmd/kosarajuSCC/main.go mediumDG.txt
 *  10 strong components
 *  21
 *  2 5 6 8 9 11 12 13 15 16 18 19 22 23 25 26 28 29 30 31 32 33 34 35 37 38 39 40 42 43 44 46 47 48 49
 *  14
 *  3 4 17 20 24 27 36
 *  41
 *  7
 *  45
 *  1
 *  0
 *  10
 *
 ******************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"

	"../../algorithms"
	"../../stdin"
	"../../structures"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: kosarajuSCC file")
		os.Exit(0)
	}

	g := algorithms.NewDigraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	scc := algorithms.NewKosarajuSCC(g)
	m := scc.Count()
	fmt.Println(m, "strong components")
	components := make([]*structures.Queue, m)

	for i := 0; i < m; i++ {
		components[i] = structures.NewQueue()
	}

	for v := 0; v < g.V(); v++ {
		components[scc.ID(v)].Enqueue(v)
	}

	for i := 0; i < m; i++ {

		for _, v := range components[i].Slice() {
			fmt.Print(v, " ")
		}

		fmt.Println()
	}
}
