package traversal

import (
	"container/list"
	"errors"

	g "github.com/g0nta/goraph/graph"
)

//BFS do breadth first search.
type BFS struct {
	graph        g.IGraph
	visitFlagKey string
	queue        *list.List
}

//NewBFS returns new BFS struct.
func NewBFS(graph g.IGraph, visitFlagKey string, start interface{}) (*BFS, error) {
	bfs := new(BFS)
	bfs.graph = graph
	bfs.visitFlagKey = visitFlagKey
	if graph.ContainsVertex(start) == false {
		return nil, errors.New("graph does not have a start vertex")
	}
	bfs.queue = list.New()
	bfs.queue.PushBack(start)
	bfs.graph.UpdateVertexAttribute(start, visitFlagKey, true)
	return bfs, nil
}

// Next updates a current vertex to a next vertex
// and returns whether a next vertex exists or not.
func (bfs *BFS) Next() bool {
	current := bfs.queue.Front().Value
	visitFlagKey := bfs.visitFlagKey

	neighbors := bfs.graph.GetNeighbors(current)
	for key, value := range neighbors {
		if flag, isExist := value[visitFlagKey]; isExist == false || flag == false {
			bfs.queue.PushBack(key)
		}
	}

	if current == bfs.queue.Front().Value {
		bfs.queue.Remove(bfs.queue.Front())
	}

	if bfs.queue.Len() == 0 {
		return false
	}
	return true
}

// Value returns current node.
func (bfs *BFS) Value() interface{} {
	return bfs.queue.Front().Value
}
