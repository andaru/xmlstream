package xmlstream

import (
	"encoding/xml"
	"fmt"
)

const (
	// XMLNS is the XML namespace of this package
	XMLNS = "https://github.com/andaru/xmlstream"
)

var (
	// CBTokenize is the NodeTokenCallback called during tokenization.
	CBTokenize = xml.Name{Space: XMLNS, Local: "callback-tokenize"}
	// CBEndElement is the NodeTokenCallback called at element end.
	CBEndElement = xml.Name{Space: XMLNS, Local: "callback-end-element"}
	// CBHandoff denotes the node option which calls StateFn for handoff.
	CBHandoff = xml.Name{Space: XMLNS, Local: "callback-handoff"}
)

// NodeTokenCallback is the prototype for schema attached callbacks
type NodeTokenCallback func(*Node, xml.Token)

// NodeIterFn is the prototype for node iterators
type NodeIterFn func(*Node) error

// NodeType is the schema node type
type NodeType int

const (
	NodeTypeRoot NodeType = iota
	NodeTypeElement
	NodeTypeText
	NodeTypeProcInst
	NodeTypeCB
)

// Node is a node forming a schema tree
type Node struct {
	Parent, Child    *Node
	NextSib, PrevSib *Node
	Opt              *nodeOptions
	T                NodeType
	Name             xml.Name
	Value            interface{}
}

func (t NodeType) Format(f fmt.State, _ rune) {
	switch t {
	case NodeTypeCB:
		f.Write([]byte("callback"))
	case NodeTypeElement:
		f.Write([]byte("element"))
	case NodeTypeRoot:
		f.Write([]byte("root"))
	case NodeTypeText:
		f.Write([]byte("text"))
	case NodeTypeProcInst:
		f.Write([]byte("processing-instruction"))
	default:
		f.Write([]byte(fmt.Sprintf("NodeType(%d)", int(t))))
	}
}

// NewSchema returns a new schema for parsing XML streams configured
// with supplied options.
func NewSchema(options ...NodeOption) *Node {
	// opts will be heap allocated only if NodeOption are provided
	var opts nodeOptions
	root := &Node{T: NodeTypeRoot}
	root.PrevSib = root
	for _, option := range options {
		option(&opts)
		root.Opt = &opts
	}
	root.Append(ProcInstNode("xml", WithMinOccurs(0), WithMaxOccurs(1)))
	return root
}

func (n *Node) ParentElement() (it *Node) {
	for it = n.Parent; it != nil; it = it.Parent {
		if it.T == NodeTypeRoot || it.T == NodeTypeElement && it.Name.Local != "" {
			return
		}
	}
	return
}

func (n *Node) Append(child *Node) *Node {
	appendNode(child, n)
	return child
}

func (n *Node) Prepend(child *Node) *Node {
	prependNode(child, n)
	return child
}

func (n *Node) InsertBefore(ref, child *Node) *Node {
	if ref == nil {
		return n.Prepend(child)
	}
	insertNodeBefore(child, ref)
	return child
}

func (n *Node) InsertAfter(ref, child *Node) *Node {
	if ref == nil {
		return n.Append(child)
	}
	insertNodeAfter(child, ref)
	return child
}

func (n *Node) Iter(fn NodeIterFn) (err error) {
	for it := n.Child; it != nil; it = it.NextSib {
		if err = fn(it); err != nil {
			break
		}
	}
	return
}

func (n *Node) IterReverse(fn NodeIterFn) (err error) {
	if n.Child == nil {
		return
	}
	for it := n.Child.PrevSib; it.NextSib != nil; it = it.PrevSib {
		if err = fn(it); err != nil {
			break
		}
	}
	return
}

func appendNode(child, parent *Node) {
	child.Parent = parent
	child.NextSib = nil
	if head := parent.Child; head != nil {
		tail := head.PrevSib
		tail.NextSib = child
		child.PrevSib = tail
		head.PrevSib = child
	} else {
		parent.Child = child
		child.PrevSib = child
	}
}

func prependNode(child, parent *Node) {
	child.Parent = parent
	head := parent.Child
	if head != nil {
		child.PrevSib = head.PrevSib
		head.PrevSib = child
	} else {
		child.PrevSib = child
	}
	child.NextSib = head
	parent.Child = child
}

func insertNodeBefore(child, before *Node) {
	parent := before.Parent
	child.Parent = parent
	if before.PrevSib.NextSib != nil {
		before.PrevSib.NextSib = child
	} else {
		// before is head, so child will now be head
		parent.Child.PrevSib = before.PrevSib
		parent.Child = child
	}
	child.PrevSib = before.PrevSib
	child.NextSib = before
	before.PrevSib = child
}

func insertNodeAfter(child, after *Node) {
	parent := after.Parent
	child.Parent = parent
	if next := after.NextSib; next != nil {
		next.PrevSib = child
	} else {
		parent.Child.PrevSib = child
	}
	child.NextSib = after.NextSib
	child.PrevSib = after
	after.NextSib = child
}
