package graph

import (
	"testing"
)

// C#で言うとこのTestCaseみたいなのを使いたい。。。

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g == nil {
		t.Error("TestNewGraph has faild.")
	}
	if g.vertexSet == nil {
		t.Error("TestNewGraph has faild. The vertexSet is nil.")
	}
	if g.adj == nil {
		t.Error("TestNewGraph has faild. The adj is nil.")
	}
	t.Log("TestNewGraph has successed.")
}

//TestAddVertex1 tests AddVertex.
//A vertex that will be added doesn't exist in g.
func TestAddVertex1(t *testing.T) {
	g := NewGraph()
	attr := map[string]interface{}{"weight": 10}
	if g.AddVertex(1, attr) == false {
		t.Error("TestAddVertex1 has faild. True must be returned.")
	}
	if _, isExists := g.vertexSet[1]; isExists == false {
		t.Error("TestAddVertex1 has faild. A vertex wasn't added")
	}
	if g.vertexSet[1]["weight"] != 10 {
		t.Errorf("TestAddVertex1 has faild. An added attribute is wrong. Expected: %d, Acutually: %d", 10, g.vertexSet[1]["weight"])
	}
	if _, isExists := g.adj[1]; isExists == false {
		t.Error("TestAddVertex1 has faild. Neighbors doesn't exist.")
	}
	t.Log("TestAddVertex1 has successed.")
}

//TestAddVertex2 tests AddVertex.
//A vertex that will be added already exists in g.
func TestAddVertex2(t *testing.T) {
	g := NewGraph()
	attr := map[string]interface{}{"weight": 10}
	g.AddVertex(1, attr)
	if g.AddVertex(1, nil) == true {
		t.Error("TestAddVertex2 has faild. False must be returned.")
	}
	if g.vertexSet[1]["weight"] != 10 {
		t.Error("TestAddVertex2 has faild. Invalid update has done.")
	}
	t.Log("TestAddVertex2 has successed.")
}

//TestAddVertices1 tests AddVertices
//Vertices that will be added don't exist in g.
func TestAddVertices1(t *testing.T) {
	g := NewGraph()
	vertices := map[interface{}]map[string]interface{}{
		1: {"weight": 10},
		2: {"weight": 20},
		3: {"weight": 30},
	}
	actual := g.AddVertices(vertices)
	expected := 3
	if actual != 3 {
		t.Errorf("TestAddVertices1 has faild. It must return %d, but actually %d.", expected, actual)
	}
	if len(g.vertexSet) != 3 {
		t.Error("TestAddVertices1 has faild.")
	}
	t.Log("TestAddVertices1 has successed.")
}

//TestAddvertices2 tests AddVertices
//Some vertices that will be added already exist in g.
func TestAddVertices2(t *testing.T) {
	g := NewGraph()
	g.vertexSet[1] = map[string]interface{}{"weight": 10}
	g.adj[1] = make(map[interface{}]map[string]interface{}, 0)
	vertices := map[interface{}]map[string]interface{}{
		1: {"weight": 10},
		2: {"weight": 20},
		3: {"weight": 30},
	}
	actual := g.AddVertices(vertices)
	expected := 2
	if actual != 2 {
		t.Errorf("TestAddVertices1 has faild. It must return %d, but actually %d.", expected, actual)
	}
	if len(g.vertexSet) != 3 {
		t.Error("TestAddVertices2 has faild.")
	}
	t.Log("TestAddVertices2 has successed.")
}

func TestGetVertexAttributes1(t *testing.T) {
	g := NewGraph()
	g.vertexSet[1] = map[string]interface{}{"weight": 10}
	g.adj[1] = make(map[interface{}]map[string]interface{}, 0)
	attr := g.GetVertexAttributes(1)
	if attr["weight"] != 10 {
		t.Errorf("TestGetVertex1 has faild. An Expected value is %d, but %d is returned.", 10, attr["weight"])
	}
	t.Log("TestGetVertex1 has successed.")
}

func TestGetVertexAttributes2(t *testing.T) {
	g := NewGraph()
	attr := g.GetVertexAttributes(1)

	if attr != nil {
		t.Errorf("TestGetVertex2 has faild. An Expected value is nil.")
	}
	t.Log("TestGetVertex2 has successed.")
}

func TestGetNeighbors1(t *testing.T) {
	g := NewGraph()
	u := 1
	attru := map[string]interface{}{"weight": 10}
	v := 2
	attrv := map[string]interface{}{"weight": 20}
	w := 3
	attrw := map[string]interface{}{"weight": 30}

	g.vertexSet[u] = attru
	g.vertexSet[v] = attrv
	g.vertexSet[w] = attrw
	g.adj[u] = map[interface{}]map[string]interface{}{
		v: attrv,
		w: attrw,
	}
	g.adj[v] = map[interface{}]map[string]interface{}{
		u: attru,
	}
	g.adj[w] = map[interface{}]map[string]interface{}{
		w: attrw,
	}

	neighbors := g.GetNeighbors(u)

	if len(neighbors) != 2 {
		t.Errorf("TestGetNeighbors has faild. The length of neighbors must be 2 but it returns %d", len(neighbors))
	}
	if neighbors[v]["weight"] != attrv["weight"] {
		t.Errorf("TestGetNeighbors has faild. An Expected value is %d, but it returns %d", attrv["weight"], neighbors[v]["weight"])
	}
	if neighbors[w]["weight"] != attrw["weight"] {
		t.Errorf("TestGetNeighbors has faild. An Expected value is %d, but it returns %d", attrw["weight"], neighbors[w]["weight"])
	}
}

//TestAddEdge1 tests AddEdge.
//An edge e that will be added doesn't exist in g.
//Each endpoint of e also don't exist in g.
func TestAddEdge1(t *testing.T) {
	g := NewGraph()
	u := 1
	v := 2

	if g.AddEdge(u, v, nil) == false {
		t.Error("TestAddEdge1 has faild. True must be returned.")
	}
	if _, isExists := g.vertexSet[u]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A New Vertex %d wasn't added.", u)
	}
	if _, isExists := g.vertexSet[v]; isExists == false {
		t.Errorf("TestAddEdge1 has faild. A New Vertex %d wasn't added.", v)
	}
	if _, isExist := g.adj[u][v]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	}
	if _, isExist := g.adj[v][u]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	}
}

//TestAddEdge2 tests AddEdge.
//An edge e that will be added doesn't exist in g.
//Each endpoint of e already exst in g.
func TestAddEdge2(t *testing.T) {
	g := NewGraph()
	u := 1
	attru := map[string]interface{}{"weight": 10}
	v := 2
	attrv := map[string]interface{}{"weight": 20}
	g.AddVertex(u, attru)
	g.AddVertex(v, attrv)

	if g.AddEdge(u, v, nil) == false {
		t.Error("TestAddEdge1 has faild. True must be returned.")
	}
	if _, isExist := g.adj[u][v]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
	}
	if _, isExist := g.adj[v][u]; isExist == false {
		t.Error("TestAddEdge1 has faild. Neigbhors wasn't updated.")
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
	if g.AddEdge(u, v, nil) == true {
		t.Error("TestAddEdge3 has faild. True must be returned.")
	}
}
