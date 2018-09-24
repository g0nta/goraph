package goraph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g == nil {
		t.Error("TestNewGraph has faild.")
	}
	if g.vertexSet == nil {
		t.Error("TestNewGraph has faild. The vertexSet is nil.")
	}
	if g.edgeSet == nil {
		t.Error("TestNewGraph has faild. The edgeSet is nil.")
	}
	if g.neighbors == nil {
		t.Error("TestNewGraph has faild. The neighbors is nil.")
	}
	t.Log("TestNewGraph has successed.")
}

func TestAddVertexWhenNewVertexAdded(t *testing.T) {
	g := NewGraph()
	g.AddVertex(1, nil)
	if _, isExists := g.vertexSet[1]; isExists == false {
		t.Error("TestAddVertexWhenNewVertexAdded has faild. A vertex wasn't added")
	}
	if _, isExists := g.neighbors[1]; isExists == false {
		t.Error("TestAddVertexWhenNewVertexAdded has faild. Neighbors doesn't exist.")
	}
	if len(g.neighbors[1]) != 0 {
		t.Error("TestAddVertexWhenNewVertexAdded has faild. Number of neighbors must be zero.")
	}
	t.Log("TestAddVertexWhenNewVertexAdded has successed.")
}

func TestAddVertexWhenExistingVertexAdded(t *testing.T) {
	g := NewGraph()
	g.AddVertex(1, nil)
	g.AddVertex(1, nil)
	if len(g.vertexSet) != 1 {
		t.Error("TestAddVertexWhenExistingVertexAdded has faild.")
	}
	t.Log("TestAddVertexWhenExistingVertexAdded has successed.")
}

func TestAddVerticesWhenAllVerticesAreNew(t *testing.T) {
	g := NewGraph()
	vertices := []Vertex{Vertex{1, nil}, Vertex{2, nil}, Vertex{3, nil}}

	g.AddVertices(vertices)
	if len(g.vertexSet) != 3 {
		t.Error("TestAddVerticesWhenAllVerticesAreNew has faild.")
	}
	t.Log("TestAddVerticesWhenAllVerticesAreNew has successed.")
}
