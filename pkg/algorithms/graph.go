package algorithms

import (
	"fmt"

	"github.com/EAddario/algorithms/pkg/stdin"
	"github.com/EAddario/algorithms/pkg/structures"
)

// Graph ...
type Graph struct {
	v, e int
	adj  []*structures.Bag
}

// NewGraph ...
func NewGraph(in *stdin.In) *Graph {
	v := in.ReadInt()
	adj := make([]*structures.Bag, v)

	for i := 0; i < v; i++ {
		adj[i] = structures.NewBag()
	}

	g := &Graph{v: v, e: 0, adj: adj}
	e := in.ReadInt()

	for i := 0; i < e; i++ {
		v, w := in.ReadInt(), in.ReadInt()
		g.AddEdge(v, w)
	}

	return g
}

// NewGraphV ...
func NewGraphV(v int) *Graph {
	adj := make([]*structures.Bag, v)

	for i := 0; i < v; i++ {
		adj[i] = structures.NewBag()
	}

	return &Graph{v: v, e: 0, adj: adj}
}

// V ...
func (g *Graph) V() int {
	return g.v
}

// E ...
func (g *Graph) E() int {
	return g.e
}

// AddEdge ...
func (g *Graph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.e++
}

// Adj ...
func (g *Graph) Adj(v int) []int {
	return g.adj[v].IntSlice()
}

func (g *Graph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges\n", g.v, g.e)

	for i := 0; i < g.v; i++ {
		var adjs string

		for _, w := range g.Adj(i) {
			adjs = adjs + fmt.Sprintf(" %d", w)
		}

		s += fmt.Sprintf("%d: %s\n", i, adjs)
	}

	return s
}
