/******************************************************************************
 *  Execution:    go run cmd/binaryDump/main.go n < file
 *  Data file:    https://introcs.cs.princeton.edu/stdlib/abra.txt
 *  
 *  Reads in a binary file and writes out the bits, n per line.
 *
 *  % more abra.txt 
 *  ABRACADABRA!
 *
 *  % go run cmd/binaryDump/main.go 16 < abra.txt
 *  0100000101000010
 *  0101001001000001
 *  0100001101000001
 *  0100010001000001
 *  0100001001010010
 *  0100000100100001
 *  96 bits
 *
 ******************************************************************************/

package main

import (
	"fmt"
	"os"
	"strconv"

	"../../stdin"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Usage: binaryDump [bits] < file ")
		os.Exit(0)
	}

	bitsPerLine := 16

	if len(os.Args) == 2{
		bitsPerLine, _ = strconv.Atoi(os.Args[1])
	}

	var count int

	for count=0;!stdin.BinaryStdin.IsEmpty();count++{

		if bitsPerLine == 0{
			stdin.BinaryStdin.ReadBool()
			continue
		} else if count != 0 && count  % bitsPerLine == 0 {
			fmt.Println()
		}

		if stdin.BinaryStdin.ReadBool(){
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}

	}

	if bitsPerLine != 0 {
		fmt.Println()
	}

	fmt.Println(count, "bits")
}
