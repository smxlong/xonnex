package graph

import "errors"

// ErrorNodeAlreadyInGraph means an attempt was made to add a node to a graph
// that the node was already part of.
var ErrorNodeAlreadyInGraph = errors.New("node is already in graph")

// ErrorNodeNotInGraph means the references node was not in the graph.
var ErrorNodeNotInGraph = errors.New("node is not in graph")

// ErrorEdgeAlreadyInGraph means an attempt was made to add an edge to a graph
// that already existed.
var ErrorEdgeAlreadyInGraph = errors.New("edge is already in graph")

// graphMetadata is metadata for a Graph.
type graphMetadata struct {
	name string
}

// Name returns the name of the graph.
func (g *graphMetadata) Name() string {
	return g.name
}

// Graph is a graph of Nodes and Edges.
type Graph struct {
	nodes    []*Node
	edges    []*Edge
	metadata graphMetadata
}

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		nodes: []*Node{},
		edges: []*Edge{},
	}
}

// NewGraphWithName returns a new Graph with a name.
func NewGraphWithName(name string) *Graph {
	return &Graph{
		nodes: []*Node{},
		edges: []*Edge{},
		metadata: graphMetadata{
			name: name,
		},
	}
}

// Metadata returns the Metadata of a Graph.
func (g *Graph) Metadata() Metadata {
	return &g.metadata
}

// AddNode adds a Node to a Graph. If the Node is already part of the Graph,
// ErrorNodeAlreadyInGraph is returned.
func (g *Graph) AddNode(node *Node) error {
	for _, existingNode := range g.nodes {
		if existingNode == node {
			return ErrorNodeAlreadyInGraph
		}
	}
	g.nodes = append(g.nodes, node)
	return nil
}

// RemoveNode removes a Node from a Graph. If the Node is not part of the Graph,
// ErrorNodeNotInGraph is returned.
func (g *Graph) RemoveNode(node *Node) error {
	for i, existingNode := range g.nodes {
		if existingNode == node {
			g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
			return nil
		}
	}
	return ErrorNodeNotInGraph
}

// AddEdge adds an Edge to a graph. If the Edge is already part of the Graph,
// ErrorEdgeAlreadyInGraph is returned.
func (g *Graph) AddEdge(edge *Edge) error {
	for _, existingEdge := range g.edges {
		if existingEdge == edge {
			return ErrorEdgeAlreadyInGraph
		}
		if existingEdge.isDirectional && edge.isDirectional &&
			existingEdge.from == edge.from &&
			existingEdge.to == edge.to {
			return ErrorEdgeAlreadyInGraph
		}
		if !existingEdge.isDirectional && !edge.isDirectional &&
			((existingEdge.from == edge.from &&
				existingEdge.to == edge.to) ||
				(existingEdge.from == edge.to &&
					existingEdge.to == edge.from)) {
			return ErrorEdgeAlreadyInGraph
		}
	}
	g.edges = append(g.edges, edge)
	return nil
}
