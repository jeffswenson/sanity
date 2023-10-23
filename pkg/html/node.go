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
