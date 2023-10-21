package html

// Node represents an HTML tag or attribute. It is the core type of
// the sanity library. Nodes are immutable and may be safely shared
// across threads.
//
//   - The `tags` package contains constructors for all standard HTML5
//     tags.
//   - The `attr` package contains constructors for all standard HTML5
//     attributes.
//   - The `html` package containing the Node implementation contains
//     generic constructors and utilities.
type Node struct {
	// nodeType is an enum that acts as a tag for Node. Node is basically a
	// tagged union. The initial implementaiton of node was a little cleaner.
	// It used interfaces and dynamic dispatch instead of a switch on nodeType.
	//
	// E.g.
	// type node interface {
	//   renderAsAttribute([]bytes) []bytes
	//   reanderAsContent([]bytes) []bytes
	// }
	//
	// But implementing Node as a tagged union allows allocating a slice of
	// Nodes as a single allocation. The ability to allocate nodes in a slice
	// mean attributes require no allocations. Each tag requires only a single
	// allocation, which is the slice containing all of its children.
	nodeType nodeType
	str1     string
	str2     string
	children []Node
}

func (n Node) String() string {
	return string(n.Render())
}

// Render converts the node and all of its children into a byte array.
// Writing the byte array to an io.Writer is slightly more efficient
// than writing a string.
//
// Example Usage:
// writer.Write(node.Render())
func (n Node) Render() []byte {
	renderer := &renderVisitor{}
	n.Visit(renderer)
	return renderer.bytes
}

// Visit walks the Node tree. It is used by Node.Render to convert the Node
// tree to HTML.
func (n *Node) Visit(visitor Visitor) {
	n.visitAsContent(visitor)
}

// VisitAttributes visits the node's attributes. This should be used by the
// visitor.Tag and visitor.VoidTag implementations to visit the tag's
// attributes.
func (n *Node) VisitAttributes(visitor Visitor) {
	for i := range n.children {
		n.children[i].visitAsAttribute(visitor)
	}
}

// VisitChildren visits the node's children. This should be used by visitor.Tag
// implementation to visit tags and inner HTML contained within the tag.
func (n *Node) VisitChildren(visitor Visitor) {
	for i := range n.children {
		n.children[i].visitAsContent(visitor)
	}
}

// Visitor makes it possible to walk the Node tree. See `render.go` for an
// example visitor implementation. Visitor is used to implement node.Render.
type Visitor interface {
	Tag(name string, node *Node)
	VoidTag(name string, node *Node)
	Content(content string)
	Attribute(name string, value *string)
}

type nodeType uint32

const (
	nodeTypeEmpty nodeType = iota

	nodeTypeAttr
	nodeTypeBoolAttr

	nodeTypeTag
	nodeTypeVoidTag

	nodeTypeRawText

	nodeTypeMany
)

func (n *Node) visitAsAttribute(visitor Visitor) {
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

func (n *Node) visitAsContent(visitor Visitor) {
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
