package goraph

import (
	"errors"
)

//Graph : グラフの構造体。privateにすべき？
// 頂点集合と辺集合を持つようにする。
// LinkedListとかはもう少し検討が必要。
type Graph struct {
	VertexSet map[Vertex]Vertex
	EdgeSet   map[Edge]Edge
	Neighbors map[Vertex]map[Vertex]Vertex
}

//NewGraph is constractor of Graph struct. It returns an empty graph.
//スライスの容量も指定できるようにするべき。。。
func NewGraph() *Graph {
	g := new(Graph)
	g.VertexSet = make(map[Vertex]Vertex)
	g.EdgeSet = make(map[Edge]Edge)
	g.Neighbors = make(map[Vertex]map[Vertex]Vertex)
	return g
}

//AddVertex adds a vertex to a graph g.
func (g *Graph) AddVertex(v Vertex) error {
	if _, isExist := g.VertexSet[v]; isExist == true {
		return errors.New("the graph already has a same vertex")
	}

	g.VertexSet[v] = v
	g.Neighbors[v] = make(map[Vertex]Vertex, 0)
	return nil
}

//DeleteVertex deletes a vertex from g.
func (g *Graph) DeleteVertex(v Vertex) {
	delete(g.VertexSet, v)

	neighbors := g.Neighbors[v]
	for _, u := range neighbors {
		delete(g.Neighbors[u], v)
	}
	delete(g.Neighbors, v)
}

//AddEdge adds an edge to a graph g.
func (g *Graph) AddEdge(e Edge) error {
	from := e.From
	to := e.To

	// check the VertexSet contains each vertecies
	if _, isExist := g.VertexSet[from]; isExist == false {
		g.AddVertex(from)
	}
	if _, isExist := g.VertexSet[to]; isExist == false {
		g.AddVertex(to)
	}

	//check the edge set has an same edge.
	if _, isExist := g.EdgeSet[e]; isExist == true {
		return errors.New("the graph already has a same edge")
	}

	g.EdgeSet[e] = e
	g.Neighbors[from][to] = to
	g.Neighbors[to][from] = from

	return nil
}
