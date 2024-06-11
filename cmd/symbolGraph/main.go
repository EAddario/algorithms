/*
   Execution:    go run cmd/symbolGraph/main.go filename.txt delimiter
   Data files:   https://algs4.cs.princeton.edu/41graph/routes.txt
                 https://algs4.cs.princeton.edu/41graph/movies.txt
                 https://algs4.cs.princeton.edu/41graph/moviestiny.txt
                 https://algs4.cs.princeton.edu/41graph/moviesG.txt
                 https://algs4.cs.princeton.edu/41graph/moviestopGrossing.txt

   %  go run cmd/symbolGraph/main.go routes.txt " "
   JFK
      MCO
      ATL
      ORD
   LAX
      PHX
      LAS

   % go run cmd/symbolGraph/main.go movies.txt "/"
   Tin Men (1987)
      Hershey, Barbara
      Geppi, Cindy
      Jones, Kathy (II)
      Herr, Marcia
      ...
      Blumenfeld, Alan
      DeBoy, David
   Bacon, Kevin
      Woodsman, The (2004)
      Wild Things (1998)
      Where the Truth Lies (2005)
      Tremors (1990)
      ...
      Apollo 13 (1995)
      Animal House (1978)


   Assumes that input file is encoded using UTF-8.
   % iconv -f ISO-8859-1 -t UTF-8 movies-iso8859.txt > movies.txt
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EAddario/algorithms/pkg/algorithms"
	"github.com/EAddario/algorithms/pkg/stdin"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: symbolGraph file delimiter")
		os.Exit(0)
	}

	sg := algorithms.NewSymbolGraph(os.Args[1], os.Args[2])
	g := sg.G()
	si := stdin.NewStdIn(bufio.ScanLines)

	for !si.IsEmpty() {
		source := strings.Trim(si.ReadString(), " ")

		if sg.Contains(source) {
			s := sg.Index(source)

			for _, v := range g.Adj(s) {
				fmt.Println("  ", sg.Name(v))
			}
			fmt.Println()
		} else {
			fmt.Println("input does not contains source: ", source)
		}

	}

}
