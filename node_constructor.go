package xmlstream

import "encoding/xml"

// ElementNode returns a new schema node for matching XML elements.
func ElementNode(name xml.Name, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeElement, Name: name, Opt: defaultNodeOptions()}
	for _, opt := range opts {
		opt(n.Opt)
	}
	return n
}

// ProcInstNode returns a new schema node for matching XML processing
// instructions.
func ProcInstNode(target string, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeProcInst, Value: xml.ProcInst{Target: target}, Opt: &nodeOptions{}}
	for _, opt := range opts {
		opt(n.Opt)
	}
	return n
}

// TextNode returns a new schema node for matching text (character data) nodes.
func TextNode(opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeText, Opt: defaultNodeOptions()}
	for _, opt := range opts {
		opt(n.Opt)
	}
	return n
}

// CallbackNode returns a new callback schema node.
func CallbackNode(name xml.Name, fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: name, Value: fn, Opt: defaultNodeOptions()}
	for _, opt := range opts {
		opt(n.Opt)
	}
	return n
}

// TokenEventNode returns a new callback node which will be executed when its
// parent node is matched during tokenization.
func TokenEventNode(fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBTokenize, Value: fn, Opt: defaultNodeOptions()}
	for _, opt := range opts {
		opt(n.Opt)
	}
	return n
}

// EndElementEventNode returns a callback node which will be executed when its
// parent node ends (i.e., <foo>'s ending </foo> element is seen) as a token.
func EndElementEventNode(fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBEndElement, Value: fn, Opt: defaultNodeOptions()}
	for _, opt := range opts {
		opt(n.Opt)
	}
	return n
}
