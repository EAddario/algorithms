/******************************************************************************
 *  Execution:    go run cmd/bag/main.go < input.txt
 *
 *  A generic bag or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run cmd/bag/main.go < tobe.txt
 *  size of bag = 14
 *  is
 *  -
 *  -
 *  -
 *  that
 *  -
 *  -
 *  be
 *  -
 *  to
 *  not
 *  or
 *  be
 *  to
 *
 ******************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"

	"../../stdin"
	"../../structures"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: bag file ")
		os.Exit(0)
	}

	words := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	bag := structures.NewBag()

	for _, word := range words {
		bag.Add(word)
	}

	fmt.Println("size of bag = ", bag.Size())

	for _, item := range bag.StringSlice() {
		fmt.Println(item)
	}

}
