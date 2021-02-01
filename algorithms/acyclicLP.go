//Computes the longest path tree from {@code s} to every other vertex in
//the directed acyclic graph
//G the acyclic digraph
//s the source vertex

package algorithms

import (
	"fmt"
	"os"
)

// NegativeInfinity is a fake positive infinity value
const NegativeInfinity = -999999.0

// AcyclicLP ...
type AcyclicLP struct {
	distTo []float32
	edgeTo []*DirectedEdge
}

// NewAcyclicLP ...
func NewAcyclicLP(g *EdgeWeightedDigraph, s int) *AcyclicLP {
	edgeTo := make([]*DirectedEdge, g.V())
	distTo := make([]float32, g.V())

	for v := 0; v < g.V(); v++ {
		distTo[v] = NegativeInfinity
	}

	distTo[s] = 0.0
	a := &AcyclicLP{distTo, edgeTo}
	top := NewEWDTopological(g)

	if !top.IsDAG() {
		fmt.Println("Digraph is not acyclic.")
		os.Exit(0)
	}

	for _, v := range top.Order() {

		for _, e := range g.Adj(v) {
			a.relax(e)
		}

	}

	return a
}

func (a *AcyclicLP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()

	if a.distTo[w] < a.distTo[v] + e.Weight() {
		a.distTo[w] = a.distTo[v] + e.Weight()
		a.edgeTo[w] = e
	}
}

// DistTo ...
func (a *AcyclicLP) DistTo(v int) float32 {
	return a.distTo[v]
}

// HasPathTo ...
func (a *AcyclicLP) HasPathTo(v int) bool {
	return a.distTo[v] > NegativeInfinity
}

// PathTo ...
func (a *AcyclicLP) PathTo(v int) (edges []*DirectedEdge) {

	if !a.HasPathTo(v) {
		return nil
	}

	for e := a.edgeTo[v]; e != nil; e = a.edgeTo[e.From()] {
		// stack
		edges = append([]*DirectedEdge{e}, edges...)
	}

	return
}
