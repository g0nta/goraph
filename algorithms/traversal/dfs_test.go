package traversal

import (
	"testing"
)

// Test case structure of NewBFS
type testCaseNewDFS struct {
	IsContainsVertex bool
	IsErrNil         bool
	IsBfsNil         bool
	ExpectedCurrent  interface{}
}

func TestNewDFS(t *testing.T) {
	testCases := []testCaseNewDFS{
		testCaseNewDFS{
			IsContainsVertex: true,
			IsErrNil:         true,
			IsBfsNil:         false,
		},
		testCaseNewDFS{
			IsContainsVertex: false,
			IsErrNil:         false,
			IsBfsNil:         true,
		},
	}

	for _, value := range testCases {
		g := NewMGraph(value.IsContainsVertex, nil)
		bfs, err := NewDFS(g, "isVisit", 1)
		if (err == nil) != value.IsErrNil {
			t.Errorf("TestNewBFS1 has failed.")
		}

		if (bfs == nil) != value.IsBfsNil {
			t.Errorf("TestNewBFS1 has failed.")
		}
	}
}

// Test case structure of Next.
type testCaseDFSNext struct {
	Neighbors           map[interface{}]map[string]interface{}
	ExpectedNextResult  bool
	ExpectedStackLength int
}

func TestDFSNext(t *testing.T) {
	visitFlagKey := "isVisit"
	testCase := []testCaseDFSNext{
		{
			Neighbors: map[interface{}]map[string]interface{}{
				2: {visitFlagKey: false},
				3: {visitFlagKey: false},
			},
			ExpectedNextResult:  true,
			ExpectedStackLength: 2,
		},
		{
			Neighbors: map[interface{}]map[string]interface{}{
				2: {visitFlagKey: true},
				3: {visitFlagKey: true},
			},
			ExpectedNextResult:  false,
			ExpectedStackLength: 0,
		},
	}

	for _, c := range testCase {
		g := NewMGraph(true, c.Neighbors)
		bfs, _ := NewBFS(g, visitFlagKey, 1)
		if bfs.Next() != c.ExpectedNextResult {
			t.Errorf("TestNext1 has failed. %T must be returned.", c.ExpectedNextResult)
		}
		if bfs.queue.Len() != c.ExpectedStackLength {
			t.Errorf("TestNext1 has failed. %T must be returned.", c.ExpectedStackLength)
		}
	}
}
