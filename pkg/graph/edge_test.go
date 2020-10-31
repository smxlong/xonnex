package graph

import "testing"

func TestNewEdge(t *testing.T) {
	n := &Node{}
	n2 := &Node{}
	e := NewEdge(n, n2)
	if e == nil {
		t.Fatal()
	}
	if e.Metadata().Name() != "" {
		t.Fatal()
	}
	if e.From() != n {
		t.Fatal()
	}
	if e.To() != n2 {
		t.Fatal()
	}
	if e.IsDirectional() {
		t.Fatal()
	}
}

func TestNewEdgeDirectional(t *testing.T) {
	n := &Node{}
	n2 := &Node{}
	e := NewEdgeDirectional(n, n2)
	if e == nil {
		t.Fatal()
	}
	if e.Metadata().Name() != "" {
		t.Fatal()
	}
	if e.From() != n {
		t.Fatal()
	}
	if e.To() != n2 {
		t.Fatal()
	}
	if !e.IsDirectional() {
		t.Fatal()
	}
}
