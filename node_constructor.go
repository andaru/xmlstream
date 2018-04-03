package xmlstream

import "encoding/xml"

// ElementNode returns a new schema node for matching XML elements.
func ElementNode(name xml.Name, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeElement, Name: name, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// ProcInstNode returns a new schema node for matching XML processing
// instructions.
func ProcInstNode(target string, opts ...NodeOption) *Node {
	typ := NodeTypeProcInst
	n := &Node{T: typ, Value: xml.ProcInst{Target: target}, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// TextNode returns a new schema node for matching text (character data) nodes.
func TextNode(opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeText, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// CallbackNode returns a new callback schema node.
func CallbackNode(name xml.Name, fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: name, Value: fn, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// TokenEventNode returns a new callback when the token is read, prior
// to any specific node type callback such as StartElementEventNode
// being called.
func TokenEventNode(fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBTokenize, Value: fn, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// StartElementEventNode returns a new callback node which will be executed when its
// parent node is matched during StartElement processing.
func StartElementEventNode(fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBStartElement, Value: fn, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// EndElementEventNode returns a callback node which will be executed when its
// parent node ends (i.e., <foo>'s ending </foo> element is seen) as a token.
func EndElementEventNode(fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBEndElement, Value: fn, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// TextEventNode returns a callback node which will be executed when parent
// TextNode is tokenized.
func TextEventNode(fn NodeTokenCallback, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBText, Value: fn, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

// HandoffEventNode returns a StateFn callback node which will be
// executed when it is found in a parent prior to the next XML token
// being read.
func HandoffEventNode(fn StateFn, opts ...NodeOption) *Node {
	n := &Node{T: NodeTypeCB, Name: CBHandoff, Value: fn, Opt: defaultNodeOptions(), Status: &nodeData{}}
	for _, opt := range opts {
		opt(n)
	}
	return n
}
