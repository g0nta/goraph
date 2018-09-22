package goraph

//Vertex : 頂点構造体
type Vertex struct {
	ID     uint
	Weight float32
}

//Edge : struct of edge. Maybe this struct don't need...
type Edge struct {
	From   Vertex
	To     Vertex
	Weight float32
}

//IGraph is interface of graphs.
//AddVertex adds vertex to its graph. If the graph already has same vertex, it returns error.
//AddEdge add edge to its graph. If the graph already has same vertex, it returns error.
type IGraph interface {
	AddVertex(v Vertex) error
	AddVertices(vertices []Vertex)
	DeleteVertex(v Vertex)
	DeleteVertices(vertices []Vertex)
	AddEdge(e Edge) error
	AddEdges(edges []Edge)
	DeleteEdge(e Edge)
	DeleteEdges(edges []Edge)

	BuildAdjacencyMatrix() [][]float32 //重み付きとかを考えてfloat
}
