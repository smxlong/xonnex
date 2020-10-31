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
