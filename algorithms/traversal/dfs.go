package traversal

import (
	"container/list"
	"errors"

	g "github.com/g0nta/goraph/graph"
)

//DFS do depth first search.
type DFS struct {
	graph        g.IGraph
	visitFlagKey string
	stack        *list.List
}

//NewDFS returns a new dfs struct.
//A graph must have a start vertex.
//You must specify vertex attribute key for using visit flag.
func NewDFS(graph g.IGraph, visitFlagKey string, start interface{}) (*DFS, error) {
	dfs := new(DFS)
	dfs.graph = graph
	dfs.visitFlagKey = visitFlagKey
	if graph.ContainsVertex(start) == false {
		return nil, errors.New("graph does not have a start vertex")
	}
	dfs.stack = list.New()
	dfs.stack.PushFront(start)
	// start は 訪問すみにする
	dfs.graph.UpdateVertexAttribute(start, visitFlagKey, true)
	return dfs, nil
}

// Next updates a current vertex to a next vertex
// and returns whether a next vertex exists or not.
func (dfs *DFS) Next() bool {
	if dfs.stack.Len() == 0 {
		return false
	}

	current := dfs.stack.Front().Value
	visitFlagKey := dfs.visitFlagKey
	// get current's neighbors.
	neighbors := dfs.graph.GetNeighbors(current)
	for key, value := range neighbors {
		if flag, isExist := value[visitFlagKey]; isExist == false || flag == false { //未訪問の頂点
			dfs.stack.PushFront(key)
			dfs.graph.UpdateVertexAttribute(key, visitFlagKey, true)
			break
		}
	}
	if current == dfs.stack.Front().Value {
		dfs.stack.Remove(dfs.stack.Front())
	}

	if dfs.stack.Len() == 0 {
		return false
	}
	return true
}

// Value returns current node.
func (dfs *DFS) Value() interface{} {
	return dfs.stack.Front().Value
}
