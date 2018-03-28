package xmlstream

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testIterCounter struct{ called int }

func (c *testIterCounter) Iterator(it *Node) error {
	c.called++
	return nil
}

func TestRootSchema(t *testing.T) {
	check := assert.New(t)
	root := NewSchema()
	check.Equal(NodeTypeRoot, root.T)
	check.Empty(root.Name.Local)
	check.Empty(root.Name.Space)
	check.NotNil(root.PrevSib, "PrevSib must be non-nil for all valid *Node")
	check.NotNil(root.Child, "Child must exist on the *Node returned by NewSchema()")
	check.Equal(NodeTypeProcInst, root.Child.T)
	counter := &testIterCounter{}
	check.NoError(root.Iter(counter.Iterator))
	check.Equal(1, counter.called)
}

func TestElementNode(t *testing.T) {
	check := assert.New(t)
	elem := ElementNode(xml.Name{Space: "urn:ns", Local: "foo"})
	check.Equal("urn:ns", elem.Name.Space)
	check.Equal("foo", elem.Name.Local)
	check.Equal(NodeTypeElement, elem.T)
	counter := &testIterCounter{}
	check.NoError(elem.Iter(counter.Iterator))
	check.Equal(0, counter.called)
}

func TestNodeAppend(t *testing.T) {
	check := assert.New(t)
	parent := ElementNode(xml.Name{Local: "parent"})
	var children []*Node
	children = append(children, parent.Append(ElementNode(xml.Name{Local: "child1"})))
	children = append(children, parent.Append(ElementNode(xml.Name{Local: "child2"})))
	children = append(children, parent.Append(ElementNode(xml.Name{Local: "child3"})))
	i := 0
	check.NoError(parent.Iter(func(n *Node) error {
		want := children[i]
		t.Run(fmt.Sprintf("%s:%#v", want.T, want.Name), func(t *testing.T) {
			check = assert.New(t)
			check.Equal(want, n)
			check.Equal(want.Name, n.Name)
			check.Equal(want.Child, n.Child)
		})
		i++
		return nil
	}))
	check.Equal(len(children), i)
}

func TestNodePrepend(t *testing.T) {
	check := assert.New(t)
	parent := ElementNode(xml.Name{Local: "parent"})
	var children []*Node
	children = append(children, parent.Prepend(ElementNode(xml.Name{Local: "child1"})))
	children = append(children, parent.Prepend(ElementNode(xml.Name{Local: "child2"})))
	children = append(children, parent.Prepend(ElementNode(xml.Name{Local: "child3"})))
	i := 0
	check.NoError(parent.Iter(func(n *Node) error {
		want := children[(len(children)-1)-i]
		t.Run(fmt.Sprintf("%s:%#v", want.T, want.Name), func(t *testing.T) {
			check = assert.New(t)
			check.Equal(want, n)
			check.Equal(want.Name, n.Name)
			check.Equal(want.Child, n.Child)
		})
		i++
		return nil
	}))
	check.Equal(len(children), i)
}

func TestNodeChildOperations(t *testing.T) {
	for _, tc := range []struct {
		name      string
		fn        func() *Node
		wantNames []string
	}{
		{
			name: "append*3",
			fn: func() *Node {
				root := ElementNode(xml.Name{})
				root.Append(elemTest("1"))
				root.Append(elemTest("2"))
				root.Append(elemTest("3"))
				return root
			},
			wantNames: []string{"1", "2", "3"},
		},

		{
			name: "append prepend append",
			fn: func() *Node {
				root := ElementNode(xml.Name{})
				root.Append(elemTest("1"))
				root.Prepend(elemTest("2"))
				root.Append(elemTest("3"))
				return root
			},
			wantNames: []string{"2", "1", "3"},
		},

		{
			name: "insert after and before",
			fn: func() *Node {
				root := ElementNode(xml.Name{})
				root.InsertAfter(nil, elemTest("3"))
				e1 := root.InsertBefore(nil, elemTest("1"))
				e2 := root.InsertBefore(e1, elemTest("2"))
				root.InsertAfter(e2, elemTest("4"))
				return root
			},
			wantNames: []string{"2", "4", "1", "3"},
		},

		{
			name: "insert after",
			fn: func() *Node {
				root := ElementNode(xml.Name{})
				e1 := root.InsertAfter(nil, elemTest("1"))
				e2 := root.InsertAfter(e1, elemTest("2"))
				root.InsertAfter(e2, elemTest("3"))
				return root
			},
			wantNames: []string{"1", "2", "3"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			i := 0
			tc.fn().Iter(func(it *Node) error {
				if want := tc.wantNames[i]; want != it.Name.Local {
					t.Errorf("want name %d %q, got %q", i, want, it.Name.Local)
				}
				i++
				return nil
			})
		})
	}
}

func elemTest(local string) *Node {
	return ElementNode(xml.Name{Local: local})
}
