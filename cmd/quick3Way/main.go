/*
Sorts a sequence of strings from standard input using quick sort 3 way.

% more tiny.txt
S O R T E X A M P L E

% go run cmd/quick3Way/main.go < tiny.txt
A E E L M O P R S T X                 [ one string per line ]

% more words3.txt
bed bug dad yes zoo ... all bad yet

% go run cmd/quick3Way/main.go < words3.txt
all bad bed bug dad ... yes yet zoo    [ one string per line ]
*/

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
		fmt.Println("Usage: quick3Way file")
		os.Exit(0)
	}

	items := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	fmt.Println("items: ", items)
	algorithms.QuickSort3way(algorithms.StringSlice(items))

	if !algorithms.IsSorted(algorithms.StringSlice(items)) {
		fmt.Println("Not Sorted")
		fmt.Println(items)
	} else {
		fmt.Println("sorted items: ", items)
	}

}
