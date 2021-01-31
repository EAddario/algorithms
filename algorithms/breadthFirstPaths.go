package algorithms

import "../structures"

// BreadFirstPaths ...
type BreadFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

// NewBreadFirstPaths ...
func NewBreadFirstPaths(g *Graph, s int) *BreadFirstPaths {
	b := &BreadFirstPaths{marked: make([]bool, g.V()), edgeTo: make([]int, g.V()), s: s}
	b.Bfs(g, s)
	return b
}

// Bfs ...
func (s *BreadFirstPaths) Bfs(g *Graph, v int) {
	s.marked[v] = true
	queue := structures.NewQueue()
	queue.Enqueue(v)

	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)
		for _, w := range g.Adj(v) {
			if !s.marked[w] {
				s.edgeTo[w] = v
				s.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}

}

// HasPathTo ...
func (s *BreadFirstPaths) HasPathTo(v int) bool {
	return s.marked[v]
}

// PathTo ...
func (s *BreadFirstPaths) PathTo(v int) []int {

	if !s.HasPathTo(v) {
		return nil
	}

	path := structures.NewStack()

	for x := v; x != s.s; x = s.edgeTo[x] {
		path.Push(x)
	}

	path.Push(s.s)
	return path.IntSlice()
}
