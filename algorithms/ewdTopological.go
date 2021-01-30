package algorithms

// EWDTopological implements topological order for a edge weighted digraph
type EWDTopological struct {
	order []*DirectedEdge
}

// NewEWDTopological ...
func NewEWDTopological(g *EdgeWeightedDigraph) *EWDTopological {
	t := new(EWDTopological)
	cycleFinder := NewEdgeWeightedDirectedCycle(g)

	if !cycleFinder.HasCycle() {
		dfs := NewEWDDepthFirstOrder(g)
		var order []*DirectedEdge

		for _, v := range dfs.ReversePost().Slice() {
			order = append(order, v.(*DirectedEdge))
		}

		t.order = order
	}
	//return &EWDTopological{}
	return t
}

// Order ...
func (t *EWDTopological) Order() []*DirectedEdge {
	return t.order
}

// IsDAG ...
func (t *EWDTopological) IsDAG() bool {
	return t.order != nil
}
