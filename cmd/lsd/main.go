/******************************************************************************
 *  Execution:    go run cmd/lsd/main.go w < input.txt
 *  Data files:   https://algs4.cs.princeton.edu/51radix/words3.txt
 *
 *  LSD radix sort
 *
 *    - Sort a String[] array of n extended ASCII strings (R = 256), each of length w.
 *
 *    - Sort an int[] array of n 32-bit integers, treating each integer as
 *      a sequence of w = 4 bytes (R = 256).
 *
 *  Uses extra space proportional to n + R.
 *
 *
 *  % go run cmd/lsd/main.go 3 < words3.txt
 *  all
 *  bad
 *  bed
 *  bug
 *  dad
 *  ...
 *  yes
 *  yet
 *  zoo
 *
 ******************************************************************************/

package main

import (
	"github.com/EAddario/algorithms/pkg/algorithms"
	"github.com/EAddario/algorithms/pkg/stdin"
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: lsd file")
		os.Exit(0)
	}

	//words := stdin.ReadAllStrings()
	words := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	ln := len(words[0])

	for i, w := range words {

		if l := len(w); l != ln {
			fmt.Println("Strings must have fixed length:", words[i-1], w, words[i+1])
			os.Exit(0)
		}

	}

	algorithms.LSDSort(words, ln)
	fmt.Println(words)
}
