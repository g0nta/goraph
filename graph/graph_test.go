package goraph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g.VertexSet == nil {
		t.Error("VertexSet is nil.")
	}
	if g.EdgeSet == nil {
		t.Error("EdgeSet is nil.")
	}
	if g.Neighbors == nil {
		t.Error("Neighbors is nil")
	}
	t.Log("TestNewGraph has passed.")
}
