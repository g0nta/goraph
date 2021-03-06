package traversal

import g "github.com/g0nta/goraph/graph"

// MGraph is mock graph structure.
type MGraph struct {
	vertexSet               map[interface{}]map[string]interface{}
	adj                     map[interface{}]map[interface{}]map[string]interface{}
	isContainsVertexSuccess bool
	neighbors               map[interface{}]map[string]interface{}
}

// NewMGraph is a constructor of MGraph.
func NewMGraph(isContainsVertexSuccess bool, neighbors map[interface{}]map[string]interface{}) *MGraph {
	g := new(MGraph)
	g.isContainsVertexSuccess = isContainsVertexSuccess
	g.neighbors = neighbors
	return g
}

// ContainsVertex is a mock method.
func (mg *MGraph) ContainsVertex(v interface{}) bool {
	return mg.isContainsVertexSuccess
}

func (mg *MGraph) GetVertexCount() int {
	panic("not implemented")
}

func (mg *MGraph) AddVertex(v interface{}, attr map[string]interface{}) bool {
	panic("not implemented")
}

func (mg *MGraph) AddVertices(vertices map[interface{}]map[string]interface{}) int {
	panic("not implemented")
}

func (mg *MGraph) GetVertexAttributes(v interface{}) map[string]interface{} {
	panic("not implemented")
}

func (mg *MGraph) GetNeighbors(v interface{}) map[interface{}]map[string]interface{} {
	return mg.neighbors
}

// UpdateVertexAttribute is a mock method.
func (mg *MGraph) UpdateVertexAttribute(v interface{}, key string, value interface{}) bool {
	return true
}

func (mg *MGraph) DeleteVertex(v interface{}) bool {
	panic("not implemented")
}

func (mg *MGraph) DeleteVertices(vertices []interface{}) int {
	panic("not implemented")
}

func (mg *MGraph) GetEdgeAttributes(u interface{}, v interface{}) map[string]interface{} {
	panic("not implemented")
}

func (mg *MGraph) AddEdge(u interface{}, v interface{}, attr map[string]interface{}) bool {
	panic("not implemented")
}

func (mg *MGraph) AddEdges(edges []g.Edge) int {
	panic("not implemented")
}

func (mg *MGraph) UpdateEdgeAttribute(u interface{}, v interface{}, key string, value interface{}) bool {
	panic("not implemented")
}

func (mg *MGraph) DeleteEdge(u interface{}, v interface{}) bool {
	panic("not implemented")
}

func (mg *MGraph) DeleteEdges(edges []g.Edge) int {
	panic("not implemented")
}
