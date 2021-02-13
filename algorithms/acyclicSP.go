//Computes a shortest paths tree from {@code s} to every other vertex in
//the directed acyclic graph
//G the acyclic digraph
//s the source vertex

package algorithms

import (
	"fmt"
	"os"
	"strconv"
)

// PositiveInfinity is a fake positive infinity value
const PositiveInfinity = 999999.0

// AcyclicSP ...
type AcyclicSP struct {
	distTo []float32
	edgeTo []*DirectedEdge
}

// NewAcyclicSP ...
func NewAcyclicSP(g *EdgeWeightedDigraph, s int) *AcyclicSP {
	edgeTo := make([]*DirectedEdge, g.V())
	distTo := make([]float32, g.V())

	validateVertex(len(distTo), s)

	for v := 0; v < g.V(); v++ {
		distTo[v] = PositiveInfinity
	}

	distTo[s] = 0.0
	a := &AcyclicSP{distTo, edgeTo}
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

func (a *AcyclicSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()

	if a.distTo[w] > a.distTo[v] + e.Weight() {
		a.distTo[w] = a.distTo[v] + e.Weight()
		a.edgeTo[w] = e
	}
}

// DistTo ...
func (a *AcyclicSP) DistTo(v int) float32 {
	validateVertex(len(a.distTo), v)
	return a.distTo[v]
}

// HasPathTo ...
func (a *AcyclicSP) HasPathTo(v int) bool {
	validateVertex(len(a.distTo), v)
	return a.distTo[v] < PositiveInfinity
}

// PathTo ...
func (a *AcyclicSP) PathTo(v int) (edges []*DirectedEdge) {

	if !a.HasPathTo(v) {
		return nil
	}

	for e := a.edgeTo[v]; e != nil; e = a.edgeTo[e.From()] {
		// stack
		edges = append([]*DirectedEdge{e}, edges...)
	}

	return
}

func validateVertex(v, w int) {

	if w < 0 || w >= v {
		panic("vertex " + strconv.Itoa(w) + " is not between 0 and " + strconv.Itoa(v - 1))
	}

}
