package html

// NewTag creates an element that has a closing tag. Like <div> or </button>.
func NewTag(name string, options ...Node) Node {
	return Node{
		nodeType: nodeTypeTag,
		str1:     name,
		children: options,
	}
}

// NewVoidTag creates an element that has no closing tag. Like <img> or
// <input>.
func NewVoidTag(name string, options ...Node) Node {
	return Node{
		nodeType: nodeTypeVoidTag,
		str1:     name,
		children: options,
	}
}
