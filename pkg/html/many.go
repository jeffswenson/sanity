package html

// ForEach applies the view function to each item, then returns a node
// containing the combined results.
//
// Example Usage:
//
//	fruits := []string{"apple", "bannana", "orange"}
//	list := tag.Ul(ForEach(fruits, func(fruit string) Node {
//		return tag.Li(InnerText(fruit))
//	})
//	list.String() == "<ul><li>apple</li><li>bannana</li><li>orange</li></ul>"
func ForEach[T any](items []T, view func(T) Node) Node {
	results := make([]Node, len(items))
	for i := range items {
		results[i] = view(items[i])
	}
	return Node{
		nodeType: nodeTypeMany,
		children: results,
	}
}

// Combine multiple Nodes into a single Node. This can be useful for grouping
// nodes without adding a wrapping element.
//
// Example Usage:
//
//	c := Combine(tag.Div(InnerText("hello")), tag.Div(InnerText("world")))
//	c.String() == "<div>hello</div><div>world</div>"
func Combine(options ...Node) Node {
	return Node{
		nodeType: nodeTypeMany,
		children: options,
	}
}
