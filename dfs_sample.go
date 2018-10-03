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
	dfs, err := trv.NewDFS(g, "isVisit", "A")

	if err != nil {
		fmt.Println("g doesn't have a start vertex.")
	}

	for dfs.Next() {
		fmt.Println(dfs.Value())
	}
}
