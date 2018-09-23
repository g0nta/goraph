package goraph

//Digraph is struct of directed graph
type Digraph struct {
	VertexSet map[interface{}]Vertex
	EdgeSet   map[interface{}]map[interface{}]Edge
	Neighbors map[interface{}]map[interface{}]Vertex
}
