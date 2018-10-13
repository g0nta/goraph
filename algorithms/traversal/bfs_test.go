package traversal

import (
	"testing"

	g "github.com/g0nta/goraph/graph"
)

// Test case structure of NewBFS
type testCaseNewBFS struct {
	IsContainsVertex bool
	IsErrNil         bool
	IsBfsNil         bool
	ExpectedCurrent  interface{}
}

func TestNewBFS(t *testing.T) {
	testCases := []testCaseNewBFS{
		testCaseNewBFS{
			IsContainsVertex: true,
			IsErrNil:         true,
			IsBfsNil:         false,
		},
		testCaseNewBFS{
			IsContainsVertex: false,
			IsErrNil:         false,
			IsBfsNil:         true,
		},
	}

	for _, value := range testCases {
		g := g.NewMGraph(value.IsContainsVertex, nil)
		bfs, err := NewBFS(g, "isVisit", 1)
		if (err == nil) != value.IsErrNil {
			t.Errorf("TestNewBFS1 has failed.")
		}

		if (bfs == nil) != value.IsBfsNil {
			t.Errorf("TestNewBFS1 has failed.")
		}
	}
}

// Test case structure of Next.
type testCaseNext struct {
	Neighbors          map[interface{}]map[string]interface{}
	ExpectedNextResult bool
	ExpectedQLength    int
}

func TestNext(t *testing.T) {
	testCase := []testCaseNext{
		{
			Neighbors: map[interface{}]map[string]interface{}{
				2: {"isVisit": false},
				3: {"isVisit": false},
			},
			ExpectedNextResult: true,
			ExpectedQLength:    2,
		},
		{
			Neighbors: map[interface{}]map[string]interface{}{
				2: {"isVisit": true},
				3: {"isVisit": true},
			},
			ExpectedNextResult: false,
			ExpectedQLength:    0,
		},
	}

	for _, c := range testCase {
		g := g.NewMGraph(true, c.Neighbors)
		bfs, _ := NewBFS(g, "isVisit", 1)
		if bfs.Next() != c.ExpectedNextResult {
			t.Errorf("TestNext1 has failed. %T must be returned.", c.ExpectedNextResult)
		}
		if bfs.queue.Len() != c.ExpectedQLength {
			t.Errorf("TestNext1 has failed. %T must be returned.", c.ExpectedQLength)
		}
	}
}
