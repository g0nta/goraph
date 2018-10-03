package traversal

import (
	"container/list"

	g "github.com/g0nta/goraph/graph"
)

//BFS do breadth first search.
type BFS struct {
	graph        g.IGraph
	visitFlagKey string
	queue        *list.List
}

//NewBFS returns new BFS struct.
func NewBFS(raph g.IGraph, visitFlagKey string, start interface{}) (*DFS, error) {
	return nil, nil
}

// Next updates a current vertex to a next vertex
// and returns whether a next vertex exists or not.
func (bfs *BFS) Next() bool {
	return false
}

// Value returns current node.
func (bfs *BFS) Value() interface{} {
	return nil
}
