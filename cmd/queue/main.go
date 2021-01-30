/******************************************************************************
 *  Execution:    go run cmd/queue/main.go < input.txt
 *
 *  A generic queue or multiset, implemented using a singly linked list.
 *
 *  % more tobe.txt
 *  to be or not to - be - - that - - - is
 *
 *  % go run cmd/queue/main.go < tobe.txt
 *  size of queue = 14
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

	"../../algorithms"
	"../../stdin"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: queue file")
		os.Exit(0)
	}

	//words := stdin.ReadAllStrings()
	words := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	queue := algorithms.NewQueue()

	for _, word := range words {
		queue.Enqueue(word)
	}

	fmt.Println("size of queue = ", queue.Size())

	for _, item := range queue.Slice() {
		fmt.Println(item)
	}
}
