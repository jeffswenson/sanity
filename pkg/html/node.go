package html

// Node represents an HTML tag or attribute. It is the core type of
// the sanity library. Nodes are immutable and may be safely shared
// across threads.
//
// - The `tags` package contains constructors for all standard HTML5
//   tags.
// - The `attr` package contains constructors for all standard HTML5
//   attributes.
// - The `html` package containing the Node implementation contains
//   generic constructors and utilities.
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
	str1    string
	str2    string
	data    any
}

func (n Node) String() string {
	return string(n.Render())
}

// Render converts the node and all of its children into a byte array. Writing
// the byte array to an io.Writer is slightly more efficient than writing a
// string.
//
// Example Usage:
// writer.Write(node.Render())
func (n Node) Render() []byte {
	return n.renderAsContent(nil)
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

func (n *Node) renderAsAttribute(bytes []byte) []byte {
	switch n.nodeType {
	case nodeTypeAttr:
		bytes = append(bytes, " "...)
		return n.renderAttribute(bytes)
	case nodeTypeBoolAttr:
		bytes = append(bytes, " "...)
		return n.renderBoolAttribute(bytes)
	case nodeTypeMany:
		return n.renderManyAsAttribute(bytes)
	}
	return bytes
}

func (n *Node) renderAsContent(bytes []byte) []byte {
	switch n.nodeType {
	case nodeTypeTag:
		return n.renderTag(bytes)
	case nodeTypeVoidTag:
		return n.renderVoidTag(bytes)
	case nodeTypeRawText:
		return n.renderRawText(bytes)
	case nodeTypeMany:
		return n.renderManyAsContent(bytes)
	}
	return bytes
}

func (n *Node) renderAttribute(bytes []byte) []byte {
	bytes = append(bytes, n.str1...)
	bytes = append(bytes, "=\""...)
	bytes = append(bytes, n.str2...)
	bytes = append(bytes, "\""...)
	return bytes
}

func (n *Node) renderBoolAttribute(bytes []byte) []byte {
	bytes = append(bytes, n.str1...)
	return bytes
}

func (n *Node) renderTag(bytes []byte) []byte {
	children := n.data.([]Node)

	bytes = append(bytes, "<"...)
	bytes = append(bytes, n.str1...)

	for i := range children {
		bytes = children[i].renderAsAttribute(bytes)
	}

	bytes = append(bytes, ">"...)

	for i := range children {
		bytes = children[i].renderAsContent(bytes)
	}

	bytes = append(bytes, "</"...)
	bytes = append(bytes, n.str1...)
	bytes = append(bytes, ">"...)

	return bytes
}

func (n *Node) renderVoidTag(bytes []byte) []byte {
	children := n.data.([]Node)

	bytes = append(bytes, "<"...)
	bytes = append(bytes, n.str1...)

	for i := range children {
		bytes = children[i].renderAsAttribute(bytes)
	}

	bytes = append(bytes, ">"...)

	return bytes
}

func (n *Node) renderRawText(bytes []byte) []byte {
	return append(bytes, n.str1...)
}

func (n *Node) renderManyAsContent(bytes []byte) []byte {
	children := n.data.([]Node)
	for i := range children {
		bytes = children[i].renderAsContent(bytes)
	}
	return bytes
}

func (n *Node) renderManyAsAttribute(bytes []byte) []byte {
	children := n.data.([]Node)
	for i := range children {
		bytes = children[i].renderAsAttribute(bytes)
	}
	return bytes
}
