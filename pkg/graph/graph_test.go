package graph

import "testing"

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if len(g.edges) != 0 {
		t.Fatal()
	}
	if len(g.nodes) != 0 {
		t.Fatal()
	}
	if g.Metadata().Name() != "" {
		t.Fatal()
	}
}

func TestNewGraphWithName(t *testing.T) {
	g := NewGraphWithName("testing")
	if len(g.edges) != 0 {
		t.Fatal()
	}
	if len(g.nodes) != 0 {
		t.Fatal()
	}
	if g.Metadata().Name() != "testing" {
		t.Fatal()
	}
}

func TestAddNode(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.AddNode(n)
	if err != nil {
		t.Fatal()
	}
	if len(g.nodes) != 1 {
		t.Fatal()
	}
	if g.nodes[0] != n {
		t.Fatal()
	}
}

func TestAddNodeAlreadyExists(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.AddNode(n)
	if err != nil {
		t.Fatal()
	}
	err = g.AddNode(n)
	if err == nil {
		t.Fatal()
	}
	if len(g.nodes) != 1 {
		t.Fatal()
	}
	if g.nodes[0] != n {
		t.Fatal()
	}
}

func TestAddNodeAlreadyExistsExtended(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.AddNode(n)
	if err != nil {
		t.Fatal()
	}
	n2 := NewNode()
	err = g.AddNode(n2)
	if err != nil {
		t.Fatal()
	}
	err = g.AddNode(n)
	if err == nil {
		t.Fatal()
	}
	err = g.AddNode(n2)
	if err == nil {
		t.Fatal()
	}
	if len(g.nodes) != 2 {
		t.Fatal()
	}
}

func TestRemoveNode(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.AddNode(n)
	if err != nil {
		t.Fatal()
	}
	err = g.RemoveNode(n)
	if err != nil {
		t.Fatal()
	}
	if len(g.nodes) != 0 {
		t.Fatal()
	}
}

func TestRemoveNodeDoesNotExist(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.RemoveNode(n)
	if err != ErrorNodeNotInGraph {
		t.Fatal()
	}
	if len(g.nodes) != 0 {
		t.Fatal()
	}
}

func TestRemoveNodeExtended1(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.AddNode(n)
	if err != nil {
		t.Fatal()
	}
	n2 := NewNode()
	err = g.AddNode(n2)
	if err != nil {
		t.Fatal()
	}
	err = g.RemoveNode(n)
	if err != nil {
		t.Fatal()
	}
	if len(g.nodes) != 1 {
		t.Fatal()
	}
	if g.nodes[0] != n2 {
		t.Fatal()
	}
}

func TestRemoveNodeExtended2(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	err := g.AddNode(n)
	if err != nil {
		t.Fatal()
	}
	n2 := NewNode()
	err = g.AddNode(n2)
	if err != nil {
		t.Fatal()
	}
	err = g.RemoveNode(n2)
	if err != nil {
		t.Fatal()
	}
	if len(g.nodes) != 1 {
		t.Fatal()
	}
	if g.nodes[0] != n {
		t.Fatal()
	}
}

func TestAddEdgeSelf(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	g.AddNode(n)
	e := NewEdge(n, n)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	if len(g.edges) != 1 {
		t.Fatal()
	}
	if g.edges[0] != e {
		t.Fatal()
	}
}

func TestAddEdgeSelfTwin(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	g.AddNode(n)
	e := NewEdge(n, n)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	err = g.AddEdge(e)
	if err != ErrorEdgeAlreadyInGraph {
		t.Fatal()
	}
}

func TestAddEdgeSelfTwice(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	g.AddNode(n)
	e := NewEdge(n, n)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	e2 := NewEdge(n, n)
	err = g.AddEdge(e2)
	if err != ErrorEdgeAlreadyInGraph {
		t.Fatal()
	}
}

func TestAddEdgeSelfDirectional(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	g.AddNode(n)
	e := NewEdgeDirectional(n, n)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	if len(g.edges) != 1 {
		t.Fatal()
	}
	if g.edges[0] != e {
		t.Fatal()
	}
	err = g.AddEdge(e)
	if err != ErrorEdgeAlreadyInGraph {
		t.Fatal()
	}
}

func TestAddEdgeSelfDirectionalTwice(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	g.AddNode(n)
	e := NewEdgeDirectional(n, n)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	if len(g.edges) != 1 {
		t.Fatal()
	}
	if g.edges[0] != e {
		t.Fatal()
	}
	e2 := NewEdgeDirectional(n, n)
	err = g.AddEdge(e2)
	if err != ErrorEdgeAlreadyInGraph {
		t.Fatal()
	}
	if len(g.edges) != 1 {
		t.Fatal()
	}
	if g.edges[0] != e {
		t.Fatal()
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	n2 := NewNode()
	g.AddNode(n)
	g.AddNode(n2)
	e := NewEdge(n, n2)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
}

func TestAddEdgeBackAndForth(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	n2 := NewNode()
	g.AddNode(n)
	g.AddNode(n2)
	e := NewEdge(n, n2)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	e2 := NewEdge(n2, n)
	err = g.AddEdge(e2)
	if err != ErrorEdgeAlreadyInGraph {
		t.Fatal()
	}
}

func TestAddEdgeBackAndForthDirectional(t *testing.T) {
	g := NewGraph()
	n := NewNode()
	n2 := NewNode()
	g.AddNode(n)
	g.AddNode(n2)
	e := NewEdgeDirectional(n, n2)
	err := g.AddEdge(e)
	if err != nil {
		t.Fatal()
	}
	e2 := NewEdgeDirectional(n2, n)
	err = g.AddEdge(e2)
	if err != nil {
		t.Fatal()
	}
}
