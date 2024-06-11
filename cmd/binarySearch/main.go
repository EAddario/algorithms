/******************************************************************************
 *  Execution:    go run cmd/binarySearch/main.go tinyAllowlist.txt < tinyText.txt
 *  50
 *  99
 *
 *  % binarySearch largeAllowlist.txt < largeText.txt | more
 *  499569
 *  984875
 *  295754
 *  207807
 *  140925
 *  161828
 *  [367,966 total values]
 *
 ******************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/EAddario/algorithms/pkg/algorithms"
	"github.com/EAddario/algorithms/pkg/stdin"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: binarySearch file < file")
		os.Exit(0)
	}

	arr := stdin.NewIn(os.Args[1], bufio.ScanWords).ReadAllInts()
	sort.Ints(arr)
	si := stdin.NewStdIn(bufio.ScanLines)

	for !si.IsEmpty() {
		key := si.ReadInt()

		if algorithms.IndexOf(&arr, key) == -1 {
			fmt.Println(key)
		}

	}

}
