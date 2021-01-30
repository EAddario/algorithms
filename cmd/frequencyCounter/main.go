/******************************************************************************
 *  Execution:    go run cmd/frequencyCounter/main.go L < input.txt
*  Data files:   https://algs4.cs.princeton.edu/31elementary/tinyTale.txt
*                https://algs4.cs.princeton.edu/31elementary/tale.txt
*                https://algs4.cs.princeton.edu/31elementary/leipzig100K.txt
*                https://algs4.cs.princeton.edu/31elementary/leipzig300K.txt
*                https://algs4.cs.princeton.edu/31elementary/leipzig1M.txt
*
*  Read in a list of words from standard input and print out
*  the most frequently occurring word that has length greater than
*  a given threshold.
*
*  % go run cmd/frequencyCounter/main.go 1 < tinyTale.txt
*  it 10
*
*  % go run cmd/frequencyCounter/main.go 8 < tale.txt
*  business 122
*
*  % go run cmd/frequencyCounter/main.go 10 < leipzig1M.txt
*  government 24763
*
*
******************************************************************************/

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
		fmt.Println("Usage: frequencyCounter file threshold")
		os.Exit(0)
	}

	minLen, _ := strconv.Atoi(os.Args[2])
	//st := algorithms.NewSequentialSearchST()
	//st := algorithms.NewSeparateChainHashST(0)
	st := algorithms.NewLinearProbingHashST(0)
	si := stdin.NewIn(os.Args[1], bufio.ScanWords)

	for !si.IsEmpty() {
		word := si.ReadString()

		if len(word) < minLen {
			continue
		}

		wordKey := algorithms.StringHashKey(word)

		if !st.Contains(wordKey) {
			st.Put(wordKey, 1)
		} else {
			st.Put(wordKey, st.GetInt(wordKey)+1)
		}

	}

	max := algorithms.StringHashKey("")
	st.Put(max, 0)

	for _, word := range st.Keys() {

		if st.GetInt(word.(algorithms.StringHashKey)) > st.GetInt(max) {
			max = word.(algorithms.StringHashKey)
		}

	}

	fmt.Printf("%s %d\n", max, st.GetInt(max))
}
