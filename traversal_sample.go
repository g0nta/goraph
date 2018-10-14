package main

import (
	"fmt"

	trv "github.com/g0nta/goraph/algorithms/traversal"
	graph "github.com/g0nta/goraph/graph"
)

func main() {
	g := graph.NewGraph()
	edges := []graph.Edge{
		graph.Edge{From: "A", To: "B", Attributes: nil},
		graph.Edge{From: "A", To: "C", Attributes: nil},
		graph.Edge{From: "B", To: "C", Attributes: nil},
		graph.Edge{From: "B", To: "D", Attributes: nil},
	}
	g.AddEdges(edges)
	dfs, errDfs := trv.NewDFS(g, "isVisitDFS", "A")

	if errDfs != nil {
		fmt.Println("g doesn't have a start vertex.")
	}

	fmt.Println("DFS sample.")
	traversalSample(dfs)

	bfs, errBfs := trv.NewBFS(g, "isVisitBFS", "A")

	if errBfs != nil {
		fmt.Println("g doesn't have a start vertex.")
	}

	fmt.Println("BFS sample.")
	traversalSample(bfs)
}

func traversalSample(traveler trv.Traveler) {
	for traveler.Next() {
		fmt.Println(traveler.Value())
	}
}
