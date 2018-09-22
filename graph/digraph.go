package goraph

//Digraph is struct of directed graph
type Digraph struct {
	VertexSet map[Vertex]Vertex
	EdgeSet   map[Edge]Edge
	Neighbors map[Vertex]map[Vertex]Vertex
}
