package graph

//Graph : グラフの構造体。privateにすべき？
// 頂点集合と辺集合を持つようにする。
// LinkedListとかはもう少し検討が必要。
type Graph struct {
	// vertex は任意の型をとれるようにする。それにattributeがつく。
	// attributeはkeyがstring、valueが任意の型であるようなmap
	vertexSet map[interface{}]map[string]interface{}

	// edge は vertexの2つ組（hyper graphは考えない。。。）
	// attributeはvertexと同じ。
	adj map[interface{}]map[interface{}]map[string]interface{}
}

//NewGraph is constractor of Graph struct. It returns an empty graph.
func NewGraph() *Graph {
	g := new(Graph)
	g.vertexSet = make(map[interface{}]map[string]interface{})
	g.adj = make(map[interface{}]map[interface{}]map[string]interface{})
	return g
}

//ContainsVertex checks that g has a vertex v.
func (g *Graph) ContainsVertex(v interface{}) bool {
	_, isExist := g.vertexSet[v]
	return isExist
}

//GetVertexCount returns a number of g's vertices.
func (g *Graph) GetVertexCount() int {
	return len(g.vertexSet)
}

//AddVertex adds a vertex to g.
//errorを返すべきなのかbool(success or fail)を返すべきなのか。。。
func (g *Graph) AddVertex(v interface{}, attr map[string]interface{}) bool {
	if _, isExist := g.vertexSet[v]; isExist == true {
		return false
	}

	g.vertexSet[v] = attr
	g.adj[v] = make(map[interface{}]map[string]interface{}, 0)
	return true
}

//AddVertices adds some vertices to g
func (g *Graph) AddVertices(vertices map[interface{}]map[string]interface{}) int {
	successCount := 0
	for v, attribute := range vertices {
		if g.AddVertex(v, attribute) == true {
			successCount++
		}
	}
	return successCount
}

//GetVertexAttributes get v's attribute. If v is not in g, it returns nil.
func (g *Graph) GetVertexAttributes(v interface{}) map[string]interface{} {
	return g.vertexSet[v]
}

//GetNeighbors gets a v's neighbors. If v doesn't have neighbors, it returns nil.
func (g *Graph) GetNeighbors(v interface{}) map[interface{}]map[string]interface{} {
	neighbors := make(map[interface{}]map[string]interface{}, 0)
	for key := range g.adj[v] {
		neighbors[key] = g.vertexSet[key]
	}
	return neighbors
}

// UpdateVertexAttribute update v's attribute.
func (g *Graph) UpdateVertexAttribute(v interface{}, key string, value interface{}) bool {
	if _, isExist := g.vertexSet[v]; isExist == true {
		g.vertexSet[v][key] = value
	}
	return true
}

//DeleteVertex deletes a vertex from g.
//複数のオブジェクトをいじるからトランザクション処理っぽいことしたほうがいいか。。。
func (g *Graph) DeleteVertex(v interface{}) {
	delete(g.vertexSet, v)

	neighbors := g.adj[v]
	for key := range neighbors {
		delete(g.adj[key], v)
	}
	delete(g.adj, v)
}

//DeleteVertices deletes some vertices from g.
func (g *Graph) DeleteVertices(vertices []interface{}) {
	for _, v := range vertices {
		g.DeleteVertex(v)
	}
}

//GetEdgeAttributes gets edge's attributes.
func (g *Graph) GetEdgeAttributes(u interface{}, v interface{}) map[string]interface{} {
	return g.adj[u][v]
}

//AddEdge adds an edge to g.
//If g doesn't have vertex u or v, they are added to g without attributes.
func (g *Graph) AddEdge(u interface{}, v interface{}, attr map[string]interface{}) bool {
	// check the VertexSet contains each vertecies
	if _, isExist := g.vertexSet[v]; isExist == false {
		g.AddVertex(v, make(map[string]interface{}, 0))
	}
	if _, isExist := g.vertexSet[u]; isExist == false {
		g.AddVertex(u, make(map[string]interface{}, 0))
	}

	//check the edge set has an same edge.
	if _, isExist := g.adj[u][v]; isExist == true {
		return false
	}

	g.adj[u][v] = attr
	g.adj[v][u] = attr

	return true
}

//AddEdges adds some edges to g.
//attributeまで含めた辺のリストをそのまま加える仕様にすると使いづらいのでラップしたEdge構造体のリストを引数にする。
func (g *Graph) AddEdges(edges []Edge) int {
	successCount := 0
	for _, e := range edges {
		if g.AddEdge(e.From, e.To, e.Attributes) == true {
			successCount++
		}
	}
	return successCount
}

//UpdateEdgeAttribute updates edge's attribute
func (g *Graph) UpdateEdgeAttribute(u interface{}, v interface{}, key string, value interface{}) bool {
	if _, isExist := g.adj[u][v]; isExist == false {
		return false
	}
	g.adj[u][v][key] = value
	g.adj[v][u][key] = value
	return true
}

//DeleteEdge deletes an edge from g.
func (g *Graph) DeleteEdge(u interface{}, v interface{}) bool {
	if _, isExists := g.adj[u][v]; isExists == false {
		return false
	}

	delete(g.adj[u], v)
	delete(g.adj[v], u)
	return true
}

//DeleteEdges deletes some edges form g.
func (g *Graph) DeleteEdges(edges []Edge) int {
	successCount := 0
	for _, e := range edges {
		if isExist := g.DeleteEdge(e.To, e.From); isExist == true {
			successCount++
		}
	}
	return successCount
}

//AdjancencyMatrix returns a matrix of vertex adjacency.
//func (g *Graph) AdjacencyMatrix() map[Vertex]map[Vertex]float32 {
//	if g.AdjacencyMatrix != nil {
//		return g.AdjacencyMatrix
//	}
//	g.AdjacencyMatrix = make(map[Vertex]map[Vertex]float32, len(g.VertexSet))
//	for _, v := range g.VertexSet {
//		for _, u := range g.Neighbors[v] {
//			g.AdjacencyMatrix[v][u] =
//
//		}
//	}
//}
