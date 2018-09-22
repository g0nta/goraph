package goraph

//Vertex : 重みなしVertex. 重みありVertexは別の構造体で。。。？
type Vertex uint

//Edge : struct of edge. Maybe this struct don't need...
type Edge struct {
	From Vertex
	To   Vertex
}

//Graph : グラフの構造体。privateにすべき？
//グラフは辺集合を持つべき？隣接行列かLinkedList持ってれば持たなくてもいいが、隣接行列とLinkeListどっちを持てばいいかはグラフの疎密具合に依存する。。。
//辺集合を持っておいて後でユーザー側でBuildAdjMatrix or BuildLinkedListみたいなのをしたほうがいい気もする
//隣接行列なら辺の問い合わせは定数。隣接行列持った方がいいかな。
//隣接行列だけもつ。頂点の追加は隣接行列に行と列を足す。辺の追加は該当要素に1を立てる。とすれば良いのか？？？？
//いやーMST求める時とかはEdge Set持ってた方が楽だったりするんだよな。。。
// 頂点集合と辺集合と有向無向のフラグを持つようにする。
type Graph struct {
	VertexSet       map[Vertex]Vertex //Key: ID of vertex, Value: vertex.
	EdgeSet         map[Edge]Edge
	IsDirected      bool
	LinkedList      map[Vertex][]Vertex
	AdjacencyMatrix [][]Vertex
}

//New returns an empty graph.
func New(isDirected bool) *Graph {
	g := new(Graph)
	g.VertexSet = make(map[Vertex]Vertex)
	g.EdgeSet = make(map[Edge]Edge)
	g.IsDirected = isDirected
	g.LinkedList = make(map[Vertex][]Vertex)
	g.AdjacencyMatrix = make([][]Vertex, 0)
	return g
}

func (g *Graph) AddNode(v Vertex) error {
	if _, isExist := g.VertexSet[v]; isExist==true {
		return error.
	}
}
