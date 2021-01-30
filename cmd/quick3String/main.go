/******************************************************************************
 *  Execution:   go run cmd/quick3String/main.go < input.txt
 *  Data files:   https://algs4.cs.princeton.edu/51radix/words3.txt
 *                https://algs4.cs.princeton.edu/51radix/shells.txt
 *
 *  Reads string from standard input and 3-way string quicksort them.
 *
 *  % go run cmd/quick3String/main.go < shells.txt
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

	"../../algorithms"
	"../../stdin"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: quick3String file")
		os.Exit(0)
	}

	words := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	algorithms.NewQuick3String(words)
	fmt.Println(words)
}
