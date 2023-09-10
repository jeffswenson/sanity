package tag

import "github.com/jeffswenson/sanity/pkg/html"

// Area constructs an html.Node for the `<area>` tag.
//
// The <area> tag is used to define clickable areas within an image map in an
// HTML document. It is primarily used in conjunction with the <map> tag to
// create interactive image elements. Each <area> tag specifies a shape
// (rectangular, circular, or polygonal) and coordinates that define the area
// where a user can click. When clicked, the <area> tag triggers a specified
// action or navigation, such as opening a new webpage or running a JavaScript
// function.
//
// Example Usage:
// <img src="image.jpg" alt="Image Map">
// <map name="example">
// <area shape="rect" coords="0,0,100,100" href="page1.html" alt="Area 1">
// <area shape="circle" coords="150,100,50" href="page2.html" alt="Area 2">
// <area shape="poly" coords="200,200,250,300,300,250" href="page3.html" alt="Area 3">
// </map>
func Area(children ...html.Node) html.Node {
	return html.NewVoidTag("area", children...)
}

// Base constructs an html.Node for the `<base>` tag.
//
// The <base> tag is used to specify the base URL for all relative URLs within
// an HTML document. It allows developers to define a default URL that is used as
// the base reference point for all relative links, including hyperlinks, images,
// scripts, and stylesheets. This helps streamline the process of linking to
// resources within the document and ensures consistent behavior across multiple
// pages.
//
// Example Usage:
// <base href="https://www.example.com/">
func Base(children ...html.Node) html.Node {
	return html.NewVoidTag("base", children...)
}

// BaseFont constructs an html.Node for the `<basefont>` tag.
//
// The <basefont> tag is used to specify the base font size, color, and face
// for an HTML document. It affects the default styling of the text within the
// document, allowing developers to set a consistent and uniform look across
// multiple elements. The <basefont> tag is deprecated and should be avoided in
// favor of CSS styling, which offers more flexibility and control over the
// appearance of text.
//
// Example Usage:
// <html>
// <head>
// <basefont size="4" color="blue" face="Arial">
// </head>
// <body>
// <p>This paragraph has a larger font size, blue color, and Arial font face.</p>
// <h1>This heading has the same font style as the paragraph above.</h1>
// </body>
// </html>
func BaseFont(children ...html.Node) html.Node {
	return html.NewVoidTag("basefont", children...)
}

// Br constructs an html.Node for the `<br>` tag.
//
// The <br> tag is used to insert a single line break in an HTML document. It
// does not require a closing tag. When the <br> tag is encountered, it creates
// a new line and moves the following content to the next line. This tag is
// commonly used to add line breaks within paragraphs or to create vertical
// spacing between elements.
//
// Example Usage:
// <p>This is the first line.<br>
// This is the second line.</p>
func Br(children ...html.Node) html.Node {
	return html.NewVoidTag("br", children...)
}

// Col constructs an html.Node for the `<col>` tag.
//
// The <col> tag is used to define a column within an HTML table. It is primarily
// used to specify formatting and styling properties for a group of table cells
// or columns. The <col> tag does not have any visible content, but it allows
// developers to define properties such as width, background color, and
// borders for the associated table cells or columns. This tag is commonly used in
// conjunction with the <colgroup> tag to apply formatting to multiple columns.
//
// Example Usage:
// <table>
// <colgroup>
// <col style="width: 50%">
// <col style="width: 50%">
// </colgroup>
// <tr>
// <td>Column 1</td>
// <td>Column 2</td>
// </tr>
// </table>
func Col(children ...html.Node) html.Node {
	return html.NewVoidTag("col", children...)
}

// Embed constructs an html.Node for the `<embed>` tag.
//
// The <embed> tag is used to embed external content, such as multimedia files,
// into an HTML document. It allows you to integrate content like videos, audio
// files, or interactive applications directly into your webpage. The <embed> tag
// requires the `src` attribute to specify the URL of the embedded content. The
// type of content that can be embedded depends on the browser's capabilities and
// supported file formats.
//
// Example Usage:
// <embed src="video.mp4" type="video/mp4">
// <embed src="audio.wav" type="audio/wav">
// <embed src="game.swf" type="application/x-shockwave-flash">
func Embed(children ...html.Node) html.Node {
	return html.NewVoidTag("embed", children...)
}

// Hr constructs an html.Node for the `<hr>` tag.
//
// The <hr> tag is used to create a horizontal rule or a line in an HTML
// document. It is primarily used to visually separate content within a page,
// providing a clear visual break between sections. The <hr> tag does not
// require any closing tag and is self-closing. It can also be customized with
// CSS to modify its appearance, such as changing its color, size, or style.
//
// Example Usage:
// <hr>
// <p>This is the content above the horizontal rule.</p>
// <hr>
// <p>This is the content below the horizontal rule.</p>
func Hr(children ...html.Node) html.Node {
	return html.NewVoidTag("hr", children...)
}

// Img constructs an html.Node for the `<img>` tag.
//
// The <img> tag is used to embed an image in an HTML document. It allows
// developers to display visual content on a webpage. The "src" attribute is
// required and specifies the path or URL to the image file. The image can be in
// different formats such as JPEG, PNG, or GIF. The <img> tag also supports
// attributes like "alt" to provide alternative text for screen readers and
// "height" and "width" to define the dimensions of the image.
//
// Example Usage:
// <img src="image.jpg" alt="Description of the image" width="300" height="200">
func Img(children ...html.Node) html.Node {
	return html.NewVoidTag("img", children...)
}

// Input constructs an html.Node for the `<input>` tag.
//
// The <input> tag is used to create interactive form controls in an HTML
// document. It allows users to input data such as text, numbers, checkboxes,
// radio buttons, and more. The specific behavior of the <input> tag depends on
// the type attribute, which determines the type of input control to be
// rendered. The entered data can be processed and submitted to a server for
// further processing, typically through a form submission.
//
// Example usages:
// <input type="text" placeholder="Enter your name">
// <input type="number" min="0" max="100" value="50">
// <input type="checkbox" name="fruit" value="apple"> Apple
// <input type="radio" name="gender" value="male"> Male
// <input type="submit" value="Submit">
func Input(children ...html.Node) html.Node {
	return html.NewVoidTag("input", children...)
}

// Link constructs an html.Node for the `<link>` tag.
//
// The <link> tag is used to define a relationship between an HTML document and
// an external resource. It is primarily used to link CSS stylesheets to HTML
// documents, allowing developers to apply styles to the content. The <link>
// tag requires the `rel` attribute to specify the type of relationship, such
// as "stylesheet" for CSS files. The `href` attribute is used to specify the
// URL of the external resource. The <link> tag is placed inside the <head>
// section of an HTML document.
//
// Example Usage:
// <head>
// <link rel="stylesheet" href="styles.css">
// </head>
func Link(children ...html.Node) html.Node {
	return html.NewVoidTag("link", children...)
}

// Meta constructs an html.Node for the `<meta>` tag.
//
// The <meta> tag is used to provide metadata about an HTML document. It
// does not produce any visible content, but instead provides information about
// the document's character encoding, viewport settings, author, description,
// and other metadata. This information is used by browsers, search engines, and
// other web services to understand and categorize the webpage. The <meta> tag is
// typically placed within the <head> section of an HTML document.
//
// Example Usage:
// <meta charset="UTF-8">
// <meta name="viewport" content="width=device-width, initial-scale=1.0">
// <meta name="description" content="This is a website about cooking recipes.">
func Meta(children ...html.Node) html.Node {
	return html.NewVoidTag("meta", children...)
}

// Param constructs an html.Node for the `<param>` tag.
//
// The <param> tag is used to pass parameters or arguments to an embedded
// object in an HTML document, such as a multimedia file or a plugin. It allows
// the author to configure various settings or provide additional information to
// the embedded object. The <param> tag is typically used within the <object> or
// <embed> tags, and requires the `name` and `value` attributes to specify the
// parameter name and its corresponding value.
//
// Example Usage:
// <object>
// <param name="autoplay" value="true">
// <param name="loop" value="false">
// <param name="src" value="video.mp4">
// </object>
func Param(children ...html.Node) html.Node {
	return html.NewVoidTag("param", children...)
}

// Source constructs an html.Node for the `<source>` tag.
//
// The `<source>` tag is used to specify multiple sources for media elements
// (such as `<video>` or `<audio>`) in an HTML document. It allows the browser to
// choose the most suitable media source based on the capabilities of the user's
// device and network conditions. The `<source>` tag requires the `src` attribute
// to specify the URL of the media file and the `type` attribute to indicate the
// mime type of the media. The browser will automatically select and load the
// first compatible source. If none are compatible, the browser will fallback to
// the content provided in the opening and closing tags of the media element.
//
// Example Usage:
// <video controls>
// <source src="video.mp4" type="video/mp4">
// <source src="video.webm" type="video/webm">
// Your browser does not support the video tag.
// </video>
func Source(children ...html.Node) html.Node {
	return html.NewVoidTag("source", children...)
}

// Track constructs an html.Node for the `<track>` tag.
//
// The <track> tag is used to specify captions, subtitles, descriptions, or other
// textual data for multimedia elements, especially video and audio. It is primarily
// used in conjunction with the <video> and <audio> tags to provide additional
// information or accessibility features. The <track> tag requires the `src`
// attribute to specify the URL of the associated text track file, and the `kind`
// attribute to specify the type of data being provided, such as captions or
// descriptions. The text provided by the <track> tag will be displayed based on
// user preferences or device settings.
//
// Example Usage:
// <video controls>
// <source src="video.mp4" type="video/mp4">
// <track src="captions.vtt" kind="captions" label="English" srclang="en">
// <track src="descriptions.vtt" kind="descriptions" label="English" srclang="en">
// </video>
func Track(children ...html.Node) html.Node {
	return html.NewVoidTag("track", children...)
}

// Wbr constructs an html.Node for the `<wbr>` tag.
//
// The <wbr> tag is used to suggest a word break opportunity in a long or
// unbreakable string of text. It allows the browser to break the text at that
// point if necessary for proper formatting. This can be useful for preventing
// overflow or improving readability in situations where long or continuous text
// is present, such as URLs or file paths.
//
// Example Usage:
// <p>Visit our website at www.example<wbr>.com for more information.</p>
// <p>This file is located at C:\Program<wbr>Files\Example\filename.txt.</p>
func Wbr(children ...html.Node) html.Node {
	return html.NewVoidTag("wbr", children...)
}
