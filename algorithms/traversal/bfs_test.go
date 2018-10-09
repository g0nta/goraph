package traversal

import (
	"testing"

	g "github.com/g0nta/goraph/graph"
)

func TestNewBFS1(t *testing.T) {
	g := g.NewMGraph(true)
	bfs, err := NewBFS(g, "isVisit", 1)

	if err != nil {
		t.Errorf("TestNewBFS1 has failed. Err must not occured.")
	}

	if bfs.queue.Front().Value != 1 {
		t.Errorf("TestNewBFS1 has failed. 1 must be returned.")
	}
}

func TestNewBFS2(t *testing.T) {
	g := g.NewMGraph(false)
	bfs, err := NewBFS(g, "isVisit", 1)

	if err == nil {
		t.Errorf("TestNewBFS1 has failed. Err must be nil.")
	}

	if bfs != nil {
		t.Errorf("TestNewBFS1 has failed. bfs must be nil.")
	}
}
