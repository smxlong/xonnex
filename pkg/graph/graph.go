package graph

// graphMetadata is metadata for a Graph.
type graphMetadata struct {
	name string
}

func (g *graphMetadata) Name() string {
	return g.name
}

// Graph is a graph of Nodes and Edges.
type Graph struct {
	nodes    []*Node
	edges    []*Edge
	metadata graphMetadata
}

// Metadata returns the Metadata of a Graph.
func (g *Graph) Metadata() Metadata {
	return &g.metadata
}
