package html

// Visit calls the `TagVisitor.Tag`, `TagVisitor.VoidTag`, or
// `TagVisitor.Content` method depending on the type of thhe Node. The Visit*
// methods are used by Node.Render to convert the Node tree to HTML.
func (n *Node) Visit(visitor TagVisitor) {
	n.visitAsContent(visitor)
}

// VisitAttributes visits the Node's attributes. This should be used by the
// visitor.Tag and visitor.VoidTag implementations to iterate over the tag's
// attributes.
func (n *Node) VisitAttributes(visitor AttributeVisitor) {
	for i := range n.children {
		n.children[i].visitAsAttribute(visitor)
	}
}

// VisitChildren visits the Node's children. This should be used by visitor.Tag
// implementation to visit the tag's child tags and inner HTML.
func (n *Node) VisitChildren(visitor TagVisitor) {
	for i := range n.children {
		n.children[i].visitAsContent(visitor)
	}
}

// TagVisitor makes it possible to walk the Node tree. See `pkg/html/render.go`
// for an example visitor implementation. TagVisitor is used to implement
// node.Render.
type TagVisitor interface {
	Tag(name string, node *Node)
	VoidTag(name string, node *Node)
	Content(content string)
}

// AttributeVisitor is used to iterate over a tag's attributes. See
// `pkg/html/render.go` for an example of how to use the visitor interfaces.
type AttributeVisitor interface {
	Attribute(name string, value *string)
}

func (n *Node) visitAsAttribute(visitor AttributeVisitor) {
	switch n.nodeType {
	case nodeTypeAttr:
		visitor.Attribute(n.str1, &n.str2)
	case nodeTypeBoolAttr:
		visitor.Attribute(n.str1, nil)
	case nodeTypeMany:
		for i := range n.children {
			n.children[i].visitAsAttribute(visitor)
		}
	}
}

func (n *Node) visitAsContent(visitor TagVisitor) {
	switch n.nodeType {
	case nodeTypeTag:
		visitor.Tag(n.str1, n)
	case nodeTypeVoidTag:
		visitor.VoidTag(n.str1, n)
	case nodeTypeRawText:
		visitor.Content(n.str1)
	case nodeTypeMany:
		for i := range n.children {
			n.children[i].visitAsContent(visitor)
		}
	}
}
