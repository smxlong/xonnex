package graph

import "testing"

func TestNewNode(t *testing.T) {
	n := NewNode()
	if n == nil {
		t.Fatal()
	}
	if n.Metadata().Name() != "" {
		t.Fatal()
	}
}

func TestNewNodeWithName(t *testing.T) {
	n := NewNodeWithName("testing")
	if n == nil {
		t.Fatal()
	}
	if n.Metadata().Name() != "testing" {
		t.Fatal()
	}
}
