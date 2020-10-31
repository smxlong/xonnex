package graph

// edgeMetadata is metadata for an Edge.
type edgeMetadata struct {
	name string
}

func (e *edgeMetadata) Name() string {
	return e.name
}

// Edge is a edge in a graph.
type Edge struct {
	from, to      *Node
	isDirectional bool
	metadata      edgeMetadata
}

// NewEdge returns a new Edge between the given nodes. The Edge is not directional.
func NewEdge(from, to *Node) *Edge {
	return &Edge{
		from:          from,
		to:            to,
		isDirectional: false,
	}
}

// NewEdgeDirectional returns a new Edge between the given nodes. The Edge is directional.
func NewEdgeDirectional(from, to *Node) *Edge {
	return &Edge{
		from:          from,
		to:            to,
		isDirectional: true,
	}
}

// Metadata returns the Metadata of an Edge.
func (e *Edge) Metadata() Metadata {
	return &e.metadata
}

// From returns the Node where the Edge begins/originates.
func (e *Edge) From() *Node {
	return e.from
}

// To returns the Node where the Edge ends/terminates.
func (e *Edge) To() *Node {
	return e.to
}

// IsDirectional returns whether the Edge is directional.
func (e *Edge) IsDirectional() bool {
	return e.isDirectional
}
