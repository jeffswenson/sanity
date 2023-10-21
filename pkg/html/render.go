package html

// renderVisitor is the core implementation of Node.Render.
type renderVisitor struct {
	bytes []byte
}

func (rv *renderVisitor) Tag(name string, node *Node) {
	rv.write("<")
	rv.write(name)

	node.VisitAttributes(rv)

	rv.write(">")

	node.VisitChildren(rv)

	rv.write("</")
	rv.write(name)
	rv.write(">")
}

func (rv *renderVisitor) VoidTag(name string, node *Node) {
	rv.write("<")
	rv.write(name)

	node.VisitAttributes(rv)

	rv.write(">")
}

func (rv *renderVisitor) Content(content string) {
	rv.write(content)
}

func (rv *renderVisitor) Attribute(name string, value *string) {
	rv.write(" ")
	rv.write(name)
	if value != nil {
		rv.write("=\"")
		rv.write(*value)
		rv.write("\"")
	}
}

func (rv *renderVisitor) write(str string) {
	rv.bytes = append(rv.bytes, str...)
}
