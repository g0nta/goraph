package goraph

//Graph : グラフの構造体。privateにすべき？
// 頂点集合と辺集合を持つようにする。
// LinkedListとかはもう少し検討が必要。
type Graph struct {
	VertexSet map[interface{}]Vertex
	EdgeSet   map[interface{}]map[interface{}]Edge
	Neighbors map[interface{}]map[interface{}]Vertex
}

//NewGraph is constractor of Graph struct. It returns an empty graph.
//スライスの容量も指定できるようにするべき。。。
func NewGraph() *Graph {
	g := new(Graph)
	g.VertexSet = make(map[interface{}]Vertex)
	g.EdgeSet = make(map[interface{}]map[interface{}]Edge)
	g.Neighbors = make(map[interface{}]map[interface{}]Vertex)
	return g
}

//AddVertex adds a vertex to g.
//errorを返すべきなのかbool(success or fail)を返すべきなのか。。。
func (g *Graph) AddVertex(v interface{}, attr map[string]interface{}) {
	if _, isExist := g.VertexSet[v]; isExist == true {
		return
	}
	// attrが指定されていない場合はweightを設定する
	if attr == nil {
		attr = map[string]interface{}{"weight": 1}
	}
	g.VertexSet[v] = Vertex{v, attr}
	g.Neighbors[v] = make(map[interface{}]Vertex, 0)
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
	delete(g.VertexSet, v)

	neighbors := g.Neighbors[v]
	for _, u := range neighbors {
		delete(g.Neighbors[u], v)
	}
	delete(g.Neighbors, v)
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
	if _, isExist := g.VertexSet[v]; isExist == false {
		g.AddVertex(v, nil)
	}
	if _, isExist := g.VertexSet[u]; isExist == false {
		g.AddVertex(u, nil)
	}

	//check the edge set has an same edge.
	if _, isExist := g.Neighbors[u][v]; isExist == true {
		return
	}

	g.Neighbors[u][v] = Vertex{v, nil}
	g.Neighbors[v][u] = Vertex{u, nil}

	e := Edge{u, v, attr}
	g.EdgeSet[u][v] = e
	g.EdgeSet[v][u] = e
}

//AddEdges adds some edges to g.
func (g *Graph) AddEdges(edges []Edge) {
	for _, e := range edges {
		if _, isExists := g.VertexSet[e.From]; isExists == true {
			continue
		}
		if _, isExists := g.VertexSet[e.To]; isExists == true {
			continue
		}
		g.AddEdge(e.From, e.To, e.Attributes)
	}
}

//DeleteEdge deletes an edge from g.
func (g *Graph) DeleteEdge(e Edge) {
	if _, isExists := g.EdgeSet[e]; isExists == false {
		return
	}

	//EdgeSetから削除
	delete(g.EdgeSet[e.From], e.To)
	delete(g.EdgeSet[e.To], e.From)

	//Neighborsから削除
	delete(g.Neighbors[e.From], e.To)
	delete(g.Neighbors[e.To], e.From)
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
