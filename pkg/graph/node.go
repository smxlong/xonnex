package graph

// nodeMetadata is metadata for a Node.
type nodeMetadata struct {
	name string
}

// Name returns the name of the node.
func (n *nodeMetadata) Name() string {
	return n.name
}

// Node is a node in a graph.
type Node struct {
	Data     map[string]interface{}
	metadata nodeMetadata
}

// NewNode returns a new Node.
func NewNode() *Node {
	return &Node{}
}

// NewNodeWithName returns a new Node with a name.
func NewNodeWithName(name string) *Node {
	return &Node{
		metadata: nodeMetadata{
			name: name,
		},
	}
}

// Metadata returns the Metadata of a Node.
func (n *Node) Metadata() Metadata {
	return &n.metadata
}
