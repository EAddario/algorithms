/******************************************************************************
 *  Execution:   go run cmd/msd/main.go < input.txt
 *  Data files:   https://algs4.cs.princeton.edu/51radix/words3.txt
 *                https://algs4.cs.princeton.edu/51radix/shells.txt
 *
 *  Sort an array of strings or integers using MSD radix sort.
 *
 *  % go run cmd/msd/main.go < shells.txt
 *  are
 *  by
 *  sea
 *  seashells
 *  seashells
 *  sells
 *  sells
 *  she
 *  she
 *  shells
 *  shore
 *  surely
 *  the
 *  the
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
		fmt.Println("Usage: msd file")
		os.Exit(0)
	}

	words := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	algorithms.NewMSD(words)
	fmt.Println("sorted items:", words)
}
