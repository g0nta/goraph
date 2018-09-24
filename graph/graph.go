package goraph

//Graph : グラフの構造体。privateにすべき？
// 頂点集合と辺集合を持つようにする。
// LinkedListとかはもう少し検討が必要。
type Graph struct {
	vertexSet map[interface{}]Vertex
	edgeSet   map[interface{}]map[interface{}]Edge
	neighbors map[interface{}]map[interface{}]Vertex
}

//NewGraph is constractor of Graph struct. It returns an empty graph.
//スライスの容量も指定できるようにするべき。。。
func NewGraph() *Graph {
	g := new(Graph)
	g.vertexSet = make(map[interface{}]Vertex)
	g.edgeSet = make(map[interface{}]map[interface{}]Edge)
	g.neighbors = make(map[interface{}]map[interface{}]Vertex)
	return g
}

//AddVertex adds a vertex to g.
//errorを返すべきなのかbool(success or fail)を返すべきなのか。。。
func (g *Graph) AddVertex(v interface{}, attr map[string]interface{}) {
	if _, isExist := g.vertexSet[v]; isExist == true {
		return
	}

	g.vertexSet[v] = Vertex{v, attr}
	g.neighbors[v] = make(map[interface{}]Vertex, 0)
}

//AddVertices adds some vertices to g
func (g *Graph) AddVertices(vertices []Vertex) {
	for _, v := range vertices {
		g.AddVertex(v.Vertex, v.Attributes)
	}
}

//DeleteVertex deletes a vertex from g.
//複数のオブジェクトをいじるからトランザクション処理っぽいことしたほうがいいか。。。
func (g *Graph) DeleteVertex(v interface{}) {
	delete(g.vertexSet, v)

	neighbors := g.neighbors[v]
	for _, u := range neighbors {
		delete(g.neighbors[u], v)
	}
	delete(g.neighbors, v)
}

//DeleteVertices deletes some vertices from g.
func (g *Graph) DeleteVertices(vertices []Vertex) {
	for _, v := range vertices {
		g.DeleteVertex(v)
	}
}

//AddEdge adds an edge to g.
//If g doesn't have vertex u or v, they are added to g without attributes.
func (g *Graph) AddEdge(u interface{}, v interface{}, attr map[string]interface{}) {
	// check the VertexSet contains each vertecies
	if _, isExist := g.vertexSet[v]; isExist == false {
		g.AddVertex(v, nil)
	}
	if _, isExist := g.vertexSet[u]; isExist == false {
		g.AddVertex(u, nil)
	}

	//check the edge set has an same edge.
	if _, isExist := g.neighbors[u][v]; isExist == true {
		return
	}

	g.neighbors[u][v] = Vertex{v, nil}
	g.neighbors[v][u] = Vertex{u, nil}

	e := Edge{u, v, attr}
	g.edgeSet[u][v] = e
	g.edgeSet[v][u] = e
}

//AddEdges adds some edges to g.
func (g *Graph) AddEdges(edges []Edge) {
	for _, e := range edges {
		if _, isExists := g.vertexSet[e.From]; isExists == true {
			continue
		}
		if _, isExists := g.vertexSet[e.To]; isExists == true {
			continue
		}
		g.AddEdge(e.From, e.To, e.Attributes)
	}
}

//DeleteEdge deletes an edge from g.
func (g *Graph) DeleteEdge(e Edge) {
	if _, isExists := g.edgeSet[e]; isExists == false {
		return
	}

	//EdgeSetから削除
	delete(g.edgeSet[e.From], e.To)
	delete(g.edgeSet[e.To], e.From)

	//Neighborsから削除
	delete(g.neighbors[e.From], e.To)
	delete(g.neighbors[e.To], e.From)
}

//DeleteEdges deletes some edges form g.
func (g *Graph) DeleteEdges(edges []Edge) {
	for _, e := range edges {
		g.DeleteEdge(e)
	}
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
