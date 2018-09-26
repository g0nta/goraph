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

//TestAddVertex1 tests AddVertex.
//A vertex that will be added doesn't exist in g.
func TestAddVertex1(t *testing.T) {
	g := NewGraph()
	err := g.AddVertex(1, nil)

	if err != nil {
		t.Error("TestAddVertex1 has faild. Non nil error struct returned.")
	}
	if _, isExists := g.vertexSet[1]; isExists == false {
		t.Error("TestAddVertex1 has faild. A vertex wasn't added")
	}
	if g.vertexSet[1].Vertex != 1 {
		t.Errorf("TestAddVertex1 has faild. An added vertex is wrong. Expected: %d, Acutually: %d", 1, g.vertexSet[1].Vertex)
	}
	if _, isExists := g.neighbors[1]; isExists == false {
		t.Error("TestAddVertex1 has faild. Neighbors doesn't exist.")
	}
	if len(g.neighbors[1]) != 0 {
		t.Error("TestAddVertex1 has faild. Number of neighbors must be zero.")
	}
	t.Log("TestAddVertex1 has successed.")
}

//TestAddVertex2 tests AddVertex.
//A vertex that will be added already exists in g.
func TestAddVertex2(t *testing.T) {
	g := NewGraph()
	g.AddVertex(1, nil)
	err := g.AddVertex(1, nil)
	if err == nil {
		t.Error("TestAddVertex2 has faild.")
	}
	t.Log("TestAddVertex2 has successed.")
}

//TestAddVertices1 tests AddVertices
//Vertices that will be added don't exist in g.
func TestAddVertices1(t *testing.T) {
	g := NewGraph()
	vertices := []Vertex{Vertex{1, nil}, Vertex{2, nil}, Vertex{3, nil}}

	g.AddVertices(vertices)
	if len(g.vertexSet) != 3 {
		t.Error("TestAddVertices1 has faild.")
	}
	t.Log("TestAddVertices1 has successed.")
}

//TestAddvertices2 tests AddVertices
//Some vertices that will be added already exist in g.
func TestAddVertices2(t *testing.T) {
	g := NewGraph()
	g.vertexSet[1] = Vertex{1, nil}
	g.neighbors[1] = make(map[interface{}]Vertex, 0)
	vertices := []Vertex{Vertex{1, nil}, Vertex{2, nil}, Vertex{3, nil}}
	g.AddVertices(vertices)

	if len(g.vertexSet) != 3 {
		t.Error("TestAddVertices2 has faild.")
	}
	t.Log("TestAddVertices2 has successed.")
}

func TestGetVertex1(t *testing.T) {
	g := NewGraph()
	g.vertexSet[1] = Vertex{1, nil}
	g.neighbors[1] = make(map[interface{}]Vertex, 0)
	v := g.GetVertex(1)
	if v.Vertex != 1 {
		t.Errorf("TestGetVertex1 has faild. An Expected value is %d, but %d is returned.", 1, v.Vertex)
	}
	t.Log("TestGetVertex1 has successed.")
}

func TestGetVertex2(t *testing.T) {
	g := NewGraph()
	v := g.GetVertex(1)

	if v.Vertex != nil {
		t.Errorf("TestGetVertex2 has faild. An Expected value is nil.")
	}
	t.Log("TestGetVertex2 has successed.")
}

func TestGetNeighbors1(t *testing.T) {
	g := NewGraph()
	u := 1
	v := 2
	w := 3
	g.vertexSet[u] = Vertex{u, nil}
	g.neighbors[u] = map[interface{}]Vertex{v: Vertex{v, nil}, w: Vertex{w, nil}}
	g.vertexSet[v] = Vertex{v, nil}
	g.neighbors[v] = map[interface{}]Vertex{u: Vertex{u, nil}}
	g.vertexSet[w] = Vertex{w, nil}
	g.neighbors[w] = map[interface{}]Vertex{w: Vertex{u, nil}}

	g.edgeSet[u] = map[interface{}]Edge{v: Edge{u, v, nil}, w: Edge{u, w, nil}}
	g.edgeSet[v] = map[interface{}]Edge{u: Edge{u, v, nil}}
	g.edgeSet[w] = map[interface{}]Edge{u: Edge{u, w, nil}}

	neighbors := g.GetNeighbors(u)

	if len(neighbors) != 2 {
		t.Errorf("TestGetNeighbors has faild. The length of neighbors must be 2 but it returns %d", len(neighbors))
	}
	if neighbors[v].Vertex != v {
		t.Errorf("TestGetNeighbors has faild. An Expected value is %d, but it returns %d", v, neighbors[v].Vertex)
	}
	if neighbors[w].Vertex != w {
		t.Errorf("TestGetNeighbors has faild. An Expected value is %d, but it returns %d", w, neighbors[w].Vertex)
	}
}

func TestGetNeighbors2(t *testing.T) {
	g := NewGraph()
	u := 1
	g.vertexSet[u] = Vertex{u, nil}
	g.neighbors[u] = make(map[interface{}]Vertex, 0)

	neighbors := g.GetNeighbors(u)

	if len(neighbors) != 0 {
		t.Errorf("TestGetNeighbors has faild. The length of neighbors must be 0 but it returns %d", len(neighbors))
	}
}

//TestAddEdge1 tests AddEdge.
//An edge e that will be added doesn't exist in g.
//Each endpoint of e also don't exist in g.
func TestAddEdge1(t *testing.T) {
	g := NewGraph()
	u := 1
	v := 2
	err := g.AddEdge(u, v, nil)

	if err != nil {
		t.Error("TestAddEdge1 has faild. Non nil error struct has returned.")
	}
	if _, isExists := g.vertexSet[u]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A New Vertex %d wasn't added.", u)
	}
	if _, isExists := g.vertexSet[v]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A New Vertex %d wasn't added.", v)
	}
	if x, isExist := g.neighbors[u][v]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	} else if x.Vertex != v {
		t.Errorf("TestAddEdge1 has faild. Neibhors of u is wrong. Expedted: %d, Acutually: %d", v, x.Vertex)
	}
	if x, isExist := g.neighbors[v][u]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	} else if x.Vertex != u {
		t.Errorf("TestAddEdge1 has faild. Neibhors of u is wrong. Expedted: %d, Acutually: %d", u, x.Vertex)
	}
	if x, isExists := g.edgeSet[u][v]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A new edge %d-%d wasn't added.", u, v)
	} else if x.From != u || x.To != v || x.Attributes != nil {
		t.Errorf("TestAddEdge1 has faild. A new edge %d-%d is wrong.", u, v)
	}
	if x, isExists := g.edgeSet[v][u]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A New Vertex %d-%d wasn't added.", u, v)
	} else if x.From != u || x.To != v || x.Attributes != nil {
		t.Errorf("TestAddEdge1 has faild. A new edge %d-%d is wrong.", u, v)
	}
}

//TestAddEdge2 tests AddEdge.
//An edge e that will be added doesn't exist in g.
//Each endpoint of e already exst in g.
func TestAddEdge2(t *testing.T) {
	g := NewGraph()
	u := 1
	v := 2
	g.vertexSet[u] = Vertex{u, nil}
	g.neighbors[u] = make(map[interface{}]Vertex, 0)
	g.vertexSet[v] = Vertex{v, nil}
	g.neighbors[v] = make(map[interface{}]Vertex, 0)
	g.AddEdge(u, v, nil)
	if x, isExist := g.neighbors[u][v]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	} else if x.Vertex != v {
		t.Errorf("TestAddEdge1 has faild. Neibhors of u is wrong. Expedted: %d, Acutually: %d", v, x.Vertex)
	}
	if x, isExist := g.neighbors[v][u]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	} else if x.Vertex != u {
		t.Errorf("TestAddEdge1 has faild. Neibhors of u is wrong. Expedted: %d, Acutually: %d", u, x.Vertex)
	}
	if x, isExists := g.edgeSet[u][v]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A new edge %d-%d wasn't added.", u, v)
	} else if x.From != u || x.To != v || x.Attributes != nil {
		t.Errorf("TestAddEdge1 has faild. A new edge %d-%d is wrong.", u, v)
	}
	if x, isExists := g.edgeSet[v][u]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A New Vertex %d-%d wasn't added.", u, v)
	} else if x.From != u || x.To != v || x.Attributes != nil {
		t.Errorf("TestAddEdge1 has faild. A new edge %d-%d is wrong.", u, v)
	}
}

//TestAddEdge2 tests AddEdge.
//An edge e that will be added already exists in g.
//Each endpoint of e also exst in g.
func TestAddEdge3(t *testing.T) {
	g := NewGraph()
	u := 1
	v := 2
	g.AddEdge(u, v, nil)
	err := g.AddEdge(u, v, nil)
	if err == nil {
		t.Error("TestAddEdge3 has faild. Non nil error struct must be returned.")
	}
}
