package goraph

//Vertex : 頂点構造体
type Vertex struct {
	Vertex     interface{}
	Attributes map[string]interface{}
}

//Edge : struct of edge. Maybe this struct don't need...
type Edge struct {
	From       interface{}
	To         interface{}
	Attributes map[string]interface{}
}

//IGraph is interface of graphs.
//AddVertex adds vertex to its graph. If the graph already has same vertex, it returns error.
//AddEdge add edge to its graph. If the graph already has same vertex, it returns error.
type IGraph interface {
	AddVertex(v interface{}, attr map[string]interface{})
	AddVertices(vertices []Vertex)
	//UpdateVertex
	//UpdateVertices
	DeleteVertex(v interface{})
	DeleteVertices(vertices []interface{})
	AddEdge(u interface{}, v interface{}, attr map[string]interface{})
	AddEdges(edges []Edge)
	DeleteEdge(u interface{}, v interface{})
	DeleteEdges(edges []Edge)

	//AdjacencyMatrix() map[Vertex]map[Vertex]float32 //重み付きとかを考えてfloat. methodというかプロパティっぽく使いたい
}
