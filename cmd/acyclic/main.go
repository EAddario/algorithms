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
		fmt.Println("Usage: acyclic graph-file vertex")
		os.Exit(0)
	}

	graph := algorithms.NewEdgeWeightedDigraph(stdin.NewIn(os.Args[1], bufio.ScanWords))
	fmt.Println("NewEdgeWeightedDigraph:", graph)

	s, _ := strconv.Atoi(os.Args[2])
	sp := algorithms.NewAcyclicSP(graph, s)
	fmt.Println("Shortest path from s to each other vertex in DAG", sp)

	for v := 0; v < graph.V(); v++ {

		// find shortest path from s to each other vertex in DAG
		if sp.HasPathTo(v) {
			fmt.Printf("%d to %d (%.2f)  ", s, v, sp.DistTo(v))

			for _, e := range sp.PathTo(v) {
				fmt.Print(e, "   ")
			}

			fmt.Println()
		} else {
			fmt.Printf("%d to %d         no path\n", s, v)
		}
	}

	lp := algorithms.NewAcyclicLP(graph, s)
	fmt.Println("Longest path from s to each other vertex in DAG", lp)

	for v := 0; v < graph.V(); v++ {

		// find longest path from s to each other vertex in DAG
		if lp.HasPathTo(v) {
			fmt.Printf("%d to %d (%.2f)  ", s, v, lp.DistTo(v))

			for _, e := range lp.PathTo(v) {
				fmt.Print(e, "   ")
			}

			fmt.Println()
		} else {
			fmt.Printf("%d to %d         no path\n", s, v)
		}

	}

}
