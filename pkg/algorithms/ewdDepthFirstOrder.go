package algorithms

import "github.com/EAddario/algorithms/pkg/structures"

// DepthFirstOrder ...
type EWDDepthFirstOrder struct {
	marked      []bool
	pre, post   *structures.Queue
	reversePost *structures.Stack
}

// NewDepthFirstOrder ...
func NewEWDDepthFirstOrder(g *EdgeWeightedDigraph) *EWDDepthFirstOrder {
	d := &EWDDepthFirstOrder{
		marked:      make([]bool, g.V()),
		pre:         structures.NewQueue(),
		post:        structures.NewQueue(),
		reversePost: structures.NewStack(),
	}

	for v := 0; v < g.V(); v++ {

		if !d.marked[v] {
			d.Dfs(g, v)
		}

	}

	return d
}

// Dfs ...
func (d *EWDDepthFirstOrder) Dfs(g *EdgeWeightedDigraph, v int) {
	d.pre.Enqueue(v)
	d.marked[v] = true

	for _, w := range g.Adj(v) {

		if !d.marked[w.w] {
			d.Dfs(g, w.w)
		}

	}

	d.post.Enqueue(v)
	d.reversePost.Push(v)
}

// Pre ...
func (d *EWDDepthFirstOrder) Pre() *structures.Queue {
	return d.pre
}

// Post ...
func (d *EWDDepthFirstOrder) Post() *structures.Queue {
	return d.post
}

// ReversePost ...
func (d *EWDDepthFirstOrder) ReversePost() *structures.Stack {
	rp := structures.NewStack()

	for !d.reversePost.IsEmpty() {
		rp.Push(d.reversePost.Pop())
	}

	return rp
}
