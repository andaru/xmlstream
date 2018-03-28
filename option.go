package xmlstream

// NodeOption is a Node constructor option for setting schema options
type NodeOption func(*nodeOptions)

type nodeOptions struct {
	MinOccurs int
	MaxOccurs int
}

func defaultNodeOptions() *nodeOptions {
	return &nodeOptions{
		MinOccurs: 1,
		MaxOccurs: 1,
	}
}

// WithMaxOccurs sets the maximum number of times the node may be expressed in
// input matching the schema. Use -1 for unlimited.
func WithMaxOccurs(n int) NodeOption {
	return func(o *nodeOptions) {
		o.MaxOccurs = n
	}
}

// WithMinOccurs sets the minimum number of times the node must be expressed in
// input matching the schema. Use 0 to indicate the node is optional.
func WithMinOccurs(n int) NodeOption {
	return func(o *nodeOptions) {
		o.MinOccurs = n
	}
}
