package algorithms

import "github.com/EAddario/algorithms/pkg/structures"

// DepthFirstOrder ...
type DepthFirstOrder struct {
	marked      []bool
	pre, post   *structures.Queue
	reversePost *structures.Stack
}

// NewDepthFirstOrder ...
func NewDepthFirstOrder(g *Digraph) *DepthFirstOrder {
	d := &DepthFirstOrder{
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
func (d *DepthFirstOrder) Dfs(g *Digraph, v int) {
	d.pre.Enqueue(v)
	d.marked[v] = true

	for _, w := range g.Adj(v) {

		if !d.marked[w] {
			d.Dfs(g, w)
		}

	}

	d.post.Enqueue(v)
	d.reversePost.Push(v)
}

// Pre ...
func (d *DepthFirstOrder) Pre() *structures.Queue {
	return d.pre
}

// Post ...
func (d *DepthFirstOrder) Post() *structures.Queue {
	return d.post
}

// ReversePost ...
func (d *DepthFirstOrder) ReversePost() *structures.Stack {
	return d.reversePost
}
