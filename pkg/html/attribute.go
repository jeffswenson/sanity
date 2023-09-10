package html

import "html"

// NewAttribute creates an attribute with a value. Like id="some-id" or
// class="class-a class-b".
func NewAttribute(name string, value string) Node {
	return Node{
		nodeType: nodeTypeAttr,
		str1:    html.EscapeString(name),
		str2:    html.EscapeString(value),
	}
}

// NewBoolAttribute creates a bool attribute. A bool attribute is an attribute
// with no value. An example bool attribute is the `disabled` attribute in
// <button disabled>Submit</button>.
func NewBoolAttribute(name string) Node {
	return Node{
		nodeType: nodeTypeBoolAttr,
		str1:    html.EscapeString(name),
	}
}
