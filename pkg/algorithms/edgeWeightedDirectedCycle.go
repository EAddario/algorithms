package algorithms

// EdgeWeightedDirectedCycle for EdgeWeightedDigrapy
// Note: this is file is just placeholder, see detail algorithm in directedCycle.go
type EdgeWeightedDirectedCycle struct {
	marked  []bool
	edgeTo  []int
	cycle   []*DirectedEdge
	onStack []bool
}

// NewEdgeWeightedDirectedCycle ...
func NewEdgeWeightedDirectedCycle(g *EdgeWeightedDigraph) *EdgeWeightedDirectedCycle {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	onStack := make([]bool, g.V())
	c := &EdgeWeightedDirectedCycle{marked: marked, edgeTo: edgeTo, onStack: onStack}

	if c.hasSelfLoop(g) {
		return c
	}

	if c.hasParallelEdges(g) {
		return c
	}

	for v := 0; v < g.V(); v++ {

		if !c.marked[v] {
			c.Dfs(g, v)
		}

	}

	return c
}

// hasSelfLoop ...
func (c *EdgeWeightedDirectedCycle) hasSelfLoop(g *EdgeWeightedDigraph) bool {
	return false
}

// hasParallelEdges ...
func (c *EdgeWeightedDirectedCycle) hasParallelEdges(g *EdgeWeightedDigraph) bool {
	return false
}

// Dfs ...
func (c *EdgeWeightedDirectedCycle) Dfs(g *EdgeWeightedDigraph, v int) {
	c.onStack[v] = true
	c.marked[v] = true

	for _, e := range g.Adj(v) {
		w := e.To()

		if c.cycle != nil {
			return
		} else if !c.marked[w] {
			c.edgeTo[w] = e.From()
			c.Dfs(g, w)
		} else if c.onStack[w] {
			var cy []*DirectedEdge
			f := e

			for f.From() != w {
				cy = append(cy, f)
				f.w = f.From()
			}

			cy = append(cy, f)
			c.cycle = cy
			return
		}

	}

	c.onStack[v] = false
}

// HasCycle ...
func (c *EdgeWeightedDirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

// Cycle ...
func (c *EdgeWeightedDirectedCycle) Cycle() []*DirectedEdge {
	return c.cycle
}
