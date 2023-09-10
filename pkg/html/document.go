package html

// Document renders the preamble required by the HTML standard. 
//
// html.Document(attr.Lang("en"), html.InnerText("Hello World!")) would render
// as:
// <!DOCTYPE html>
// <html lang="en">
//     Hello world! 
// </html>
func Document(options ... Node) Node {
	return Combine(
		NewVoidTag("!DOCTYPE", NewBoolAttribute("html")),
		NewTag("html", options...),
	)
}
