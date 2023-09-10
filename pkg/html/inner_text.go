package html

import "html"

// InnerText is rendered as text inside a tag. HTML in the content is escaped.
//
// Example Usage:
// node := tag.Div(InnerText("hello world!"))
// node.String() == "<div>hello world!</div>"
func InnerText(content string) Node {
	return Node{
		nodeType: nodeTypeRawText,
		str1:     html.EscapeString(content),
	}
}

// RawInnerText is rendered as text inside a tag. Danger: HTML in the content
// is not escaped.
//
// Example Usage:
// node := tag.Div(RawInnerText("<script>alert("hello world")</script>"))
// node.String() == "<div>"<script>alert("hello world")</script>"</div>"
func RawInnerText(content string) Node {
	return Node{
		nodeType: nodeTypeRawText,
		str1:     content,
	}
}
