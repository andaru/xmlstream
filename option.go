package xmlstream

import (
	"context"
	"encoding/xml"
	"errors"
)

// NodeOption is a Node constructor option for setting schema options
type NodeOption func(*Node)

// NodeParameters are the schema parameters for a node
type NodeParameters interface {
	MinOccurs() int
	MaxOccurs() int
}

// WithMaxOccurs sets the maximum number of times the node may be expressed in
// input matching the schema. Use 0 for unlimited.
func WithMaxOccurs(n int) NodeOption {
	return func(o *Node) {
		o.Opt.maxOccurs = n
		if !hasCallbackChild(o, CBSEOccurs) {
			o.Append(CallbackNode(CBSEOccurs, func(ctx context.Context, _ *Node, t xml.Token) {
				o.Status.Occurs++
				if se, ok := xml.CopyToken(t).(xml.StartElement); ok {
					o.Status.SE = se
				}
			}))
		}
	}
}

func hasCallbackChild(n *Node, name xml.Name) (ok bool) {
	n.Iter(func(it *Node) error {
		if it.T != NodeTypeCB || name != it.Name {
			return nil
		}
		ok = true
		return errors.New("stop iteration")
	})
	return
}

// WithMinOccurs sets the minimum number of times the node must be expressed in
// input matching the schema. Use 0 to indicate the node is optional.
func WithMinOccurs(n int) NodeOption {
	return func(o *Node) {
		o.Opt.minOccurs = n
		if !hasCallbackChild(o, CBSEOccurs) {
			o.Append(CallbackNode(CBSEOccurs, func(ctx context.Context, _ *Node, t xml.Token) {
				o.Status.Occurs++
				if se, ok := xml.CopyToken(t).(xml.StartElement); ok {
					o.Status.SE = se
				}
			}))
		}
	}
}

func WithValidator(validator NodeTokenCallback) NodeOption {
	return func(o *Node) {
		o.Validator = validator
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

type nodeData struct {
	Occurs int
	SE     xml.StartElement
}

func defaultNodeOptions() *nodeOptions {
	return &nodeOptions{minOccurs: -1, maxOccurs: -1}
}

func (o nodeOptions) MinOccurs() int { return o.minOccurs }
func (o nodeOptions) MaxOccurs() int { return o.maxOccurs }

// compile time interface validation
var _ NodeParameters = nodeOptions{}
var _ NodeParameters = &nodeOptions{}
