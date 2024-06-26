/******************************************************************************
 *  Execution:    go run cmd/huffman/main.go - < input.txt   (compress)
 *  Execution:    go run cmd/huffman/main.go + < input.txt   (expand)
 *  Data files:   https://algs4.cs.princeton.edu/55compression/abra.txt
 *                https://algs4.cs.princeton.edu/55compression/tinytinyTale.txt
 *                https://algs4.cs.princeton.edu/55compression/medTale.txt
 *                https://algs4.cs.princeton.edu/55compression/tale.txt
 *
 *  Compress or expand a binary input stream using the Huffman algorithm.
 *
 *  % go run cmd/huffman/main.go - < abra.txt | go run cmd/binaryDump/main.go 60
 *  010100000100101000100010010000110100001101010100101010000100
 *  000000000000000000000000000110001111100101101000111110010100
 *  120 bits
 *
 *  % go run cmd/huffman/main.go - < abra.txt | go run cmd/huffman/main.go +
 *  ABRACADABRA!
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/EAddario/algorithms/pkg/algorithms"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: huffman mode input")
		os.Exit(0)
	}

	huffman := algorithms.NewHuffman()

	if os.Args[1] == "-" {
		huffman.Compress()
	} else {
		huffman.Expand()
	}

}
