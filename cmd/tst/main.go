/******************************************************************************
 *  Compilation:  go run cmd/tst/main.go
 *  Execution:   go run cmd/tst/main.go < words.txt
 *  Dependencies: StdIn.java
 *  Data files:   https://algs4.cs.princeton.edviu/52trie/shellsST.txt
 *
 *  A string symbol table for extended ASCII strings, implemented
 *  using a 256-way trie.
 *
 *  %go run cmd/tst/main.go < shellsST.txt
 *  by 4
 *  sea 6
 *  sells 1
 *  she 0
 *  shells 3
 *  shore 7
 *  the 5
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
		fmt.Println("Usage: tst file")
		os.Exit(0)
	}

	st := algorithms.NewTST()
	in := stdin.NewIn(os.Args[1], bufio.ScanWords)

	for i := 0; !in.IsEmpty(); i++ {
		key := in.ReadString()
		st.Put(key, i)
	}

	// print results
	if st.Size() < 100 {
		fmt.Println("keys(\"\"):")

		for _, key := range st.Keys() {
			fmt.Println(key, " ", st.Get(key))
		}

		fmt.Println()
	}

	fmt.Println("longestPrefixOf(\"shellsort\"):")
	fmt.Println(st.LongPrefixOf("shellsort"))
	fmt.Println()
	fmt.Println("longestPrefixOf(\"quicksort\"):")
	fmt.Println(st.LongPrefixOf("quicksort"))
	fmt.Println()
	fmt.Println("keysWithPrefix(\"shor\"):")

	for _, s := range st.KeysWithPrefix("shor") {
		fmt.Println(s)
	}

	fmt.Println()
	fmt.Println("keysThatMatch(\".he.l.\"):")

	for _, s := range st.KeysThatMatch(".he.l.") {
		fmt.Println(s)
	}
}
