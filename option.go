package xmlstream

// NodeOption is a Node constructor option for setting schema options
type NodeOption func(*nodeOptions)

// NodeOptions is the inteface to options for a schema node
type NodeOptions interface {
	MinOccurs() int
	MaxOccurs() int
}

// WithMaxOccurs sets the maximum number of times the node may be expressed in
// input matching the schema. Use -1 for unlimited.
func WithMaxOccurs(n int) NodeOption {
	return func(o *nodeOptions) {
		o.maxOccurs = n
	}
}

// WithMinOccurs sets the minimum number of times the node must be expressed in
// input matching the schema. Use 0 to indicate the node is optional.
func WithMinOccurs(n int) NodeOption {
	return func(o *nodeOptions) {
		o.minOccurs = n
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
var _ NodeOptions = nodeOptions{}
var _ NodeOptions = &nodeOptions{}
