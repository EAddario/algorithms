/*
Sorts a sequence of strings from standard input using selection sort.

% more tiny.txt
S O R T E X A M P L E

% go run cmd/selection/main.go < tiny.txt
A E E L M O P R S T X                 [ one string per line ]

% more words3.txt
bed bug dad yes zoo ... all bad yet

% go run cmd/selection/main.go < words3.txt
all bad bed bug dad ... yes yet zoo    [ one string per line ]
*/

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
		fmt.Println("Usage: selection file")
		os.Exit(0)
	}

	items := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllStrings()
	algorithms.SelectionSort(algorithms.StringSlice(items))

	if !algorithms.IsSorted(algorithms.StringSlice(items)) {
		fmt.Println("Not Sorted")
		fmt.Println(items)
	} else {
		fmt.Println("sorted items: ", items)
	}

}
