package xmlstream

// NodeOption is a Node constructor option for setting schema options
type NodeOption func(*Node)

// NodeParameters are the schema parameters for a node
type NodeParameters interface {
	MinOccurs() int
	MaxOccurs() int
}

// WithMaxOccurs sets the maximum number of times the node may be expressed in
// input matching the schema. Use -1 for unlimited.
func WithMaxOccurs(n int) NodeOption {
	return func(o *Node) {
		o.Opt.maxOccurs = n
	}
}

// WithMinOccurs sets the minimum number of times the node must be expressed in
// input matching the schema. Use 0 to indicate the node is optional.
func WithMinOccurs(n int) NodeOption {
	return func(o *Node) {
		o.Opt.minOccurs = n
	}
}

// WithParent sets the node's parent at construction time. Note that
// this may be over-written by any further mutations to the node.
func WithParent(parent *Node) NodeOption {
	return func(o *Node) {
		o.Parent = parent
	}
}

type nodeOptions struct {
	minOccurs int
	maxOccurs int
}

func defaultNodeOptions(t NodeType) *nodeOptions {
	commonNodeOptions := &nodeOptions{minOccurs: 0, maxOccurs: 1}
	// customization is available per NodeType here
	nodeTypeOptions := map[NodeType]*nodeOptions{
		NodeTypeElement: &nodeOptions{minOccurs: 1, maxOccurs: 1},
	}
	if options, ok := nodeTypeOptions[t]; ok {
		return options
	}
	return commonNodeOptions
}

func (o nodeOptions) MinOccurs() int { return o.minOccurs }
func (o nodeOptions) MaxOccurs() int { return o.maxOccurs }

// compile time interface validation
var _ NodeParameters = nodeOptions{}
var _ NodeParameters = &nodeOptions{}
