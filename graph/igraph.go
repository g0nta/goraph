package graph

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
	ContainsVertex(v interface{}) bool
	GetVertexCount() int
	AddVertex(v interface{}, attr map[string]interface{}) bool
	AddVertices(vertices map[interface{}]map[string]interface{}) int
	GetVertexAttributes(v interface{}) map[string]interface{}
	GetNeighbors(v interface{}) map[interface{}]map[string]interface{}
	UpdateVertexAttribute(v interface{}, key string, value interface{}) bool
	//UpdateVertices
	DeleteVertex(v interface{}) bool
	DeleteVertices(vertices []interface{}) int

	GetEdgeAttributes(u interface{}, v interface{}) map[string]interface{}
	//GetEdges
	AddEdge(u interface{}, v interface{}, attr map[string]interface{}) bool
	AddEdges(edges []Edge) int
	UpdateEdgeAttribute(u interface{}, v interface{}, key string, value interface{}) bool
	//UpdateEdges
	DeleteEdge(u interface{}, v interface{}) bool
	DeleteEdges(edges []Edge) int

	//AdjacencyMatrix() map[Vertex]map[Vertex]float32 //重み付きとかを考えてfloat. methodというかプロパティっぽく使いたい
}
