package tag

import "github.com/jeffswenson/sanity/pkg/html"

// A constructs an html.Node for the `<a>` tag.
//
// The <a> tag is used to create hyperlinks in HTML documents, allowing users
// to navigate between web pages or jump to specific sections within the same
// document. It requires the `href` attribute to specify the URL of the linked
// resource. When clicked, the browser will load the corresponding webpage or
// scroll to the anchor point within the document.
//
// Example Usage:
// <a href="https://www.example.com">Visit Example.com</a>
// <a href="#section2">Jump to Section 2</a>
func A(children ...html.Node) html.Node {
    return html.NewTag("a", children...)
}

// Abbr constructs an html.Node for the `<abbr>` tag.
// 
// The <abbr> tag is used to define an abbreviation or acronym in an HTML
// document. It is primarily used to provide a full expansion or explanation of
// the abbreviation when the user hovers over or clicks on the abbreviated
// text. The <abbr> tag helps improve accessibility and provides additional
// information to readers.
// 
// Example Usage:
// <p>The <abbr title="World Health Organization">WHO</abbr> provides global health guidance.</p>
// <p>Please refer to the <abbr title="HyperText Markup Language">HTML</abbr> specification for more information.</p>
func Abbr(children ...html.Node) html.Node {
    return html.NewTag("abbr", children...)
}

// Acronym constructs an html.Node for the `<acronym>` tag.
// 
// The <acronym> tag is used to define an acronym or abbreviation in an HTML
// document. It provides a means of indicating a shortened form of a word or
// phrase, with the intention of providing additional information or context. The
// <acronym> tag is typically used in conjunction with the `title` attribute, which
// provides an expanded explanation of the acronym or abbreviation when the user
// hovers over the element. This tag can help improve accessibility and aid users
// in understanding potentially unfamiliar terms.
// 
// Example Usage:
// <p>I work at the <acronym title="National Aeronautics and Space
// Administration">NASA</acronym>.<p>
// <p>Please bring your <acronym title="Identification">ID</acronym> to the
// meeting.<p>
func Acronym(children ...html.Node) html.Node {
    return html.NewTag("acronym", children...)
}

// Address constructs an html.Node for the `<address>` tag.
// 
// The <address> tag is used to display contact information or the author
// information for the document. It typically includes the name, address, phone
// number, and email of the author or the organization associated with the
// document. The <address> tag is often used in the footer section of a webpage
// or in the author information section of a blog post.
// 
// Example Usage:
// <address>
// John Doe<br>
// 123 Main Street<br>
// New York, NY 10001<br>
// Phone: (123) 456-7890<br>
// Email: example@example.com
// </address>
func Address(children ...html.Node) html.Node {
    return html.NewTag("address", children...)
}

// Applet constructs an html.Node for the `<applet>` tag.
// 
// The <applet> tag is used to embed Java applets in HTML documents. It
// provides a way to run Java programs within a browser. The <applet> tag
// requires the `code` attribute to specify the location of the Java applet
// class file, and the `width` and `height` attributes to set the dimensions of
// the applet display area. The implemented behavior of the applet is defined
// by the Java code contained in the specified class file.
// 
// Example Usage:
// <applet code="MyApplet.class" width="300" height="200">
// Your browser does not support Java applets.
// </applet>
func Applet(children ...html.Node) html.Node {
    return html.NewTag("applet", children...)
}

// Article constructs an html.Node for the `<article>` tag.
// 
// The `<article>` tag is used to represent a standalone piece of content
// within an HTML document. It is typically used for blog posts, news articles,
// or other similar types of content that can be independently distributed or
// syndicated. The `<article>` tag helps search engines and assistive
// technologies to identify and navigate the main content of a document. It is
// often used in conjunction with other semantic tags, such as `<header>` and
// `<footer>`, to define the structure of the page.
// 
// Example Usage:
// <article>
// <h1>Article Title</h1>
// <p>This is the main content of the article.</p>
// <footer>
// <p>Published on <time datetime="YYYY-MM-DD">Date</time> by <a href="author.html">Author</a></p>
// </footer>
// </article>
func Article(children ...html.Node) html.Node {
    return html.NewTag("article", children...)
}

// Aside constructs an html.Node for the `<aside>` tag.
// 
// The <aside> tag is used to mark content that is tangentially related to the
// main content of an HTML document. It is typically used for sidebars, pull
// quotes, and other additional information that is not essential to the
// overall structure of the page. The <aside> tag is often used in conjunction
// with the <article> and <section> tags to provide additional context or
// related content. It can be styled separately from the main content using
// CSS.
// 
// Example Usage:
// <aside>
// <h3>Related Articles</h3>
// <ul>
// <li><a href="article1.html">Article 1</a></li>
// <li><a href="article2.html">Article 2</a></li>
// </ul>
// </aside>
func Aside(children ...html.Node) html.Node {
    return html.NewTag("aside", children...)
}

// Audio constructs an html.Node for the `<audio>` tag.
// 
// The <audio> tag is used to embed audio content in an HTML document. It allows
// you to play sound files directly in the browser, without the need for external
// plugins. The <audio> tag supports various audio formats such as MP3, WAV, and
// Ogg. It provides controls for play, pause, and volume, and can be styled using
// CSS. Additionally, you can specify alternative audio sources using the <source>
// tag within the <audio> tag, allowing the browser to choose the best format
// based on compatibility.
// 
// Example Usage:
// <audio src="audiofile.mp3" controls></audio>
// <audio>
// <source src="audiofile.mp3" type="audio/mpeg">
// <source src="audiofile.ogg" type="audio/ogg">
// Your browser does not support the audio element.
// </audio>
func Audio(children ...html.Node) html.Node {
    return html.NewTag("audio", children...)
}

// BDI constructs an html.Node for the `<bdi>` tag.
// 
// The <bdi> tag is used to isolate a section of text that has a different
// text direction than the surrounding content. It is primarily used in
// internationalization and localization to correctly display text in
// bidirectional languages, such as Arabic or Hebrew. The <bdi> tag ensures
// that text within it is rendered correctly, maintaining its original
// directionality regardless of the overall document's text direction.
// 
// Example Usage:
// <p>My name is <bdi>محمد</bdi>.</p>
// <p>My username is <bdi>@user123</bdi>.</p>
func BDI(children ...html.Node) html.Node {
    return html.NewTag("bdi", children...)
}

// BDO constructs an html.Node for the `<bdo>` tag.
// 
// The <bdo> tag is used to override the default directionality of text in an
// HTML document. It is primarily used to ensure the correct rendering of text
// in languages that are written from right to left, such as Arabic or Hebrew.
// The <bdo> tag can be used to reverse the direction of text within a specific
// section or element, allowing it to be displayed correctly. This tag is
// useful for ensuring proper text alignment and readability in multilingual
// websites.
// 
// Example Usage:
// <bdo dir="rtl">This text will be displayed from right to left.</bdo>
// <bdo dir="ltr">This text will be displayed from left to right.</bdo>
func BDO(children ...html.Node) html.Node {
    return html.NewTag("bdo", children...)
}

// Big constructs an html.Node for the `<big>` tag.
// 
// The <big> tag is used to increase the size of the text it surrounds in an
// HTML document. It is primarily used for emphasizing certain content or drawing
// attention to specific text by making it larger than the surrounding text. The
// size increase applied by the <big> tag is relative to the default font size of
// the document. Although this tag has been deprecated in HTML5, it can still be
// used for backward compatibility purposes.
// 
// Example Usage:
// <p>This is a <big>big</big> heading.</p>
// <span>Click <big>here</big> to read more.</span>
func Big(children ...html.Node) html.Node {
    return html.NewTag("big", children...)
}

// BlockQuote constructs an html.Node for the `<blockquote>` tag.
// 
// The <blockquote> tag is used to represent a section of quoted text in an
// HTML document. It is primarily used to visually distinguish quoted content
// from the rest of the text, typically by indenting it or presenting it in a
// different style. The <blockquote> tag is commonly used for citing sources,
// displaying quotations, or highlighting important text. It helps to provide
// visual cues to readers that the content is a quote or extract from another
// source.
// 
// Example Usage:
// <blockquote>
// Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
// </blockquote>
// 
// <blockquote>
// "Success is not the key to happiness. Happiness is the key to success. If you love what you are doing, you will be successful." - Albert Schweitzer
// </blockquote>
func BlockQuote(children ...html.Node) html.Node {
    return html.NewTag("blockquote", children...)
}

// Body constructs an html.Node for the `<body>` tag.
// 
// The <body> tag is used to define the main content of an HTML document. It
// represents the content that will be displayed in the browser window. All
// visible elements of a webpage, such as text, images, and interactive
// elements, are enclosed within the <body> tag. The <body> tag is required in
// every HTML document and should only be used once.
// 
// Example Usage:
// <!DOCTYPE html>
// <html>
// <head>
// <title>My Webpage</title>
// </head>
// <body>
// <h1>Welcome to My Webpage</h1>
// <p>This is the main content of my webpage.</p>
// </body>
// </html>
func Body(children ...html.Node) html.Node {
    return html.NewTag("body", children...)
}

// Button constructs an html.Node for the `<button>` tag.
// 
// The <button> tag is used to create a clickable button in an HTML document.
// It allows users to trigger an action or event when clicked. The behavior of
// the button can be customized by adding JavaScript code. The button can be
// used for various purposes like submitting a form, displaying a modal, or
// triggering a function. It can display text or an image as its content.
// 
// Example Usage:
// <button onclick="myFunction()">Click me</button>
// <button type="submit">Submit</button>
func Button(children ...html.Node) html.Node {
    return html.NewTag("button", children...)
}

// Canvas constructs an html.Node for the `<canvas>` tag.
// 
// The <canvas> tag is used to draw graphics, animations, or interactive
// elements in an HTML document. It provides a rectangular drawing area where
// developers can use JavaScript to create and manipulate visual content. The
// <canvas> element itself does not display any content by default, but it can be
// programmatically controlled to render shapes, images, and other graphical
// elements. It is commonly used for creating charts, games, or visualizations
// that require custom graphical rendering.
// 
// Example Usage:
// <canvas id="myCanvas" width="500" height="300"></canvas>
func Canvas(children ...html.Node) html.Node {
    return html.NewTag("canvas", children...)
}

// Caption constructs an html.Node for the `<caption>` tag.
// 
// The <caption> tag is used to add a title or caption to a table in an HTML document.
// It is placed immediately after the opening <table> tag and before the <thead>, <tfoot>,
// or <tbody> tags. The content within the <caption> tag is displayed centered above
// the table, providing a brief description or summary of the table's contents.
// The <captain> tag is commonly used for accessibility purposes or to improve
// the readability and understanding of complex tables.
// 
// Example Usage:
// <table>
// <caption>This table shows the sales data for each month.</caption>
// <thead>
// <tr>
// <th>Month</th>
// <th>Sales</th>
// </tr>
// </thead>
// <tbody>
// <tr>
// <td>January</td>
// <td>$1000</td>
// </tr>
// <tr>
// <td>February</td>
// <td>$1500</td>
// </tr>
// </tbody>
// </table>
func Caption(children ...html.Node) html.Node {
    return html.NewTag("caption", children...)
}

// Center constructs an html.Node for the `<center>` tag.
// 
// The <center> tag is used to horizontally center-align content within an HTML
// document. It affects the positioning of the elements within its container,
// allowing them to appear in the center of the page or parent element. This
// tag is primarily used for visual styling purposes. While it was popular in
// older versions of HTML, it is now considered deprecated and its usage is
// discouraged. Instead, CSS should be used to center-align content.
// 
// Example Usage:
// <center>
// <h1>This heading is center-aligned</h1>
// <p>This paragraph is also center-aligned</p>
// </center>
func Center(children ...html.Node) html.Node {
    return html.NewTag("center", children...)
}

// Cite constructs an html.Node for the `<cite>` tag.
// 
// The <cite> tag is used to indicate a citation or reference to a piece of
// work within an HTML document. It is primarily used to emphasize the title or
// name of a book, article, or other creative or scholarly work. The <cite> tag
// does not affect the style or appearance of the text, but it signifies that
// the content is a reference and should be treated as such. It is often used
// in conjunction with the <blockquote> or <q> tags to denote quoted or
// attributed text.
// 
// Example Usage:
// <p>The novel <cite>To Kill a Mockingbird</cite> by Harper Lee is a classic.</p>
// <blockquote>
// <p><cite>"The only thing we have to fear is fear itself."</cite> - Franklin D. Roosevelt</p>
// </blockquote>
func Cite(children ...html.Node) html.Node {
    return html.NewTag("cite", children...)
}

// Code constructs an html.Node for the `<code>` tag.
// 
// The <code> tag is used to display inline code within an HTML document.
// It is primarily used to show code snippets or examples of programming
// code. The text within the <code> tag is typically displayed in a monospace
// font, making it stand out from the surrounding text. It is often used in
// developer documentation or tutorials to demonstrate specific code usage.
// 
// Example Usage:
// <p>Use the <code>print()</code> function to display a message.</p>
// <p>The following code defines a variable: <code>var x = 5;</code></p>
func Code(children ...html.Node) html.Node {
    return html.NewTag("code", children...)
}

// ColGroup constructs an html.Node for the `<colgroup>` tag.
// 
// The <colgroup> tag is used to group and style columns in an HTML table. It
// allows for the application of common formatting or attributes to multiple
// columns at once. Typically, the <colgroup> tag is placed inside the <table>
// tag and can contain one or more <col> tags, which define the properties of
// individual columns.
// 
// Example Usage:
// <table>
// <colgroup>
// <col style="background-color: yellow;">
// <col span="2" style="background-color: lightblue;">
// </colgroup>
// <tr>
// <th>Header 1</th>
// <th>Header 2</th>
// <th>Header 3</th>
// </tr>
// <tr>
// <td>Data 1</td>
// <td>Data 2</td>
// <td>Data 3</td>
// </tr>
// </table>
func ColGroup(children ...html.Node) html.Node {
    return html.NewTag("colgroup", children...)
}

// Data constructs an html.Node for the `<data>` tag.
// 
// The <data> tag is used to embed machine-readable data in an HTML document.
// It provides a way for developers to include data that can be easily accessed
// by JavaScript or other scripting languages. The content within the <data>
// tag is not displayed to the user, but it can be utilized for various
// purposes, such as storing values, tracking analytics, or dynamic data
// manipulation.
// 
// Example Usage:
// <p>The current temperature is <data value="25">25°C</data>.</p>
// <p>The total sales for the month are <data value="5000">$5000</data>.</p>
func Data(children ...html.Node) html.Node {
    return html.NewTag("data", children...)
}

// DataList constructs an html.Node for the `<datalist>` tag.
// 
// The <datalist> tag is used to provide a predefined list of options for user
// input in an HTML form. It works in conjunction with the <input> tag,
// specifically the type attribute set to "text" or "search". The <datalist>
// tag contains a series of <option> tags that define the available choices.
// When the user interacts with the input field, they are presented with a
// dropdown list of options to choose from. This tag enhances user experience
// by providing a convenient way to select from a predetermined set of options.
// 
// Example Usage:
// <input list="fruits">
// <datalist id="fruits">
// <option value="Apple">
// <option value="Banana">
// <option value="Orange">
// </datalist>
func DataList(children ...html.Node) html.Node {
    return html.NewTag("datalist", children...)
}

// DD constructs an html.Node for the `<dd>` tag.
// 
// The <dd> tag is used to define a description or sub-item in an HTML definition
// list (dl). It is used as a counterpart to the <dt> tag, which represents the
// term being defined. The <dd> tag is typically placed after the <dt> tag within
// a <dl> element. It provides additional information or details about the term
// defined by the <dt> tag. The <dd> tag can contain various elements, including
// text, images, or other HTML tags.
// 
// Example Usage:
// <dl>
// <dt>HTML</dt>
// <dd>Stands for HyperText Markup Language, the standard markup language for
// creating web pages.</dd>
// 
// <dt>CSS</dt>
// <dd>Stands for Cascading Style Sheets, a language used to describe the
// presentation of a document written in HTML.</dd>
// </dl>
func DD(children ...html.Node) html.Node {
    return html.NewTag("dd", children...)
}

// Del constructs an html.Node for the `<del>` tag.
// 
// The <del> tag is used to indicate deleted or removed content in an HTML
// document. It displays the enclosed text with a strikethrough effect to
// visually indicate that it should be considered as deleted. The <del> tag is
// often used in revisions, document histories, or to show changes in content
// over time.
// 
// Example Usage:
// <p>My <del>old</del> new phone model.</p>
// <p>Please ignore the <del>previous</del> new instructions.</p>
func Del(children ...html.Node) html.Node {
    return html.NewTag("del", children...)
}

// Details constructs an html.Node for the `<details>` tag.
// 
// The <details> tag is used to create a collapsible section of content in an
// HTML document. It provides a way to hide or reveal additional information,
// allowing users to toggle the visibility of content within the document. When
// the <details> tag is initially loaded, the content within it is hidden, and
// users can click on the summary element to expand or collapse the content.
// This tag is commonly used to create FAQs, hide/show additional details, or
// create expandable sections of content.
// 
// Example Usage:
// <details>
// <summary>Click here to view more information</summary>
// <p>This is the hidden content that will be revealed when the summary element is clicked.</p>
// </details>
func Details(children ...html.Node) html.Node {
    return html.NewTag("details", children...)
}

// Dfn constructs an html.Node for the `<dfn>` tag.
// 
// The <dfn> tag is used to define a term within an HTML document. It marks the
// term as a definition, allowing it to stand out from the surrounding text.
// The <dfn> tag can be useful for providing explanations or clarifications for
// specific terms in a document. It does not affect the style or appearance of
// the text, but it provides semantic meaning to the term.
// 
// Example Usage:
// <p>The <dfn>HTTP</dfn> protocol is used for communication between web browsers and web servers.</p>
// <p><dfn>HTML</dfn> stands for Hypertext Markup Language.</p>
func Dfn(children ...html.Node) html.Node {
    return html.NewTag("dfn", children...)
}

// Dialog constructs an html.Node for the `<dialog>` tag.
// 
// The <dialog> tag is used to create a modal or pop-up dialog box in an HTML
// document. It is primarily used to display important messages, prompts, or
// user interactions within the context of the current page. The <dialog> tag
// can be triggered to open or close using JavaScript, providing a dynamic and
// interactive experience for users. When the dialog box is open, it typically
// overlays the rest of the page, preventing interaction with the underlying
// content until it is closed.
// 
// Example usage:
// <dialog open>
// <p>This is a dialog box.</p>
// <button>Close</button>
// </dialog>
func Dialog(children ...html.Node) html.Node {
    return html.NewTag("dialog", children...)
}

// Dir constructs an html.Node for the `<dir>` tag.
// 
// The <dir> tag is used to create a directory or list of items in an HTML document.
// It organizes content into a hierarchical structure, typically representing a
// file system or folder structure. The <dir> tag is commonly used in documentation
// or file organization, providing a visual representation of a directory structure.
// 
// Example Usage:
// <dir>
// <ul>
// <li>Folder 1</li>
// <li>Folder 2</li>
// <li>Folder 3</li>
// </ul>
// </dir>
func Dir(children ...html.Node) html.Node {
    return html.NewTag("dir", children...)
}

// Div constructs an html.Node for the `<div>` tag.
// 
// The <div> tag is used to create a generic container or division in an HTML
// document. It is primarily used to group and organize other HTML elements,
// allowing for easier styling and layout control. The <div> tag does not have
// any inherent meaning or styling, but it serves as a way to logically
// organize and structure the content within a web page.
// 
// Example Usage:
// <div>
// <h1>This is a heading</h1>
// <p>This is a paragraph of text.</p>
// </div>
func Div(children ...html.Node) html.Node {
    return html.NewTag("div", children...)
}

// DL constructs an html.Node for the `<dl>` tag.
// 
// The <dl> tag is used to define a description list in an HTML document. It
// consists of a series of term-definition pairs that are contained within the
// <dl> element. The term is represented by the <dt> tag, while the definition is
// represented by the <dd> tag. This tag is commonly used to display glossaries,
// product feature lists, or metadata definitions. The <dl> tag provides a
// structured and semantically meaningful way to present information.
// 
// Example Usage:
// <dl>
// <dt>HTML</dt>
// <dd>HyperText Markup Language</dd>
// <dt>CSS</dt>
// <dd>Cascading Style Sheets</dd>
// </dl>
func DL(children ...html.Node) html.Node {
    return html.NewTag("dl", children...)
}

// DT constructs an html.Node for the `<dt>` tag.
// 
// The <dt> tag is used to define a term in a description list in an HTML document.
// It is primarily used to label a description or definition for a corresponding term.
// The <dt> tag is typically placed within a <dl> element and is followed by a <dd> tag
// that contains the description or definition. This helps organize and structure
// information in a semantic way. When rendered, the <dt> tag usually appears bold
// or with a different font weight to distinguish it from the content.
// 
// Example Usage:
// <dl>
// <dt>HTML</dt>
// <dd>HyperText Markup Language is the standard language for creating web pages.</dd>
// <dt>CSS</dt>
// <dd>Cascading Style Sheets is used to style the appearance of HTML elements.</dd>
// <dt>JavaScript</dt>
// <dd>A programming language used for adding interactivity to web pages.</dd>
// </dl>
func DT(children ...html.Node) html.Node {
    return html.NewTag("dt", children...)
}

// Em constructs an html.Node for the `<em>` tag.
// 
// The <em> tag is used to emphasize or highlight words or phrases in an HTML
// document. It alters the visual presentation of the text, typically by
// italicizing it. The <em> tag indicates that the enclosed text has semantic
// importance and should be given emphasis when rendered. It is often used to
// add emphasis to key terms, titles, or important sections of content.
// 
// Example Usage:
// <p>This is a <em>critical</em> issue that needs immediate attention.</p>
// <p>Please <em>do not</em> share your password with anyone.</p>
func Em(children ...html.Node) html.Node {
    return html.NewTag("em", children...)
}

// FieldSet constructs an html.Node for the `<fieldset>` tag.
// 
// The <fieldset> tag is used to group related form elements together in an
// HTML document. It is primarily used to visually organize and label a set of
// form controls, such as input fields, checkboxes, and radio buttons. The
// <fieldset> tag helps improve the accessibility and usability of forms by
// providing a clear and structured layout.
// 
// Example Usage:
// <fieldset>
// <legend>Personal Information</legend>
// <label for="name">Name:</label>
// <input type="text" id="name" name="name"><br>
// <label for="email">Email:</label>
// <input type="email" id="email" name="email"><br>
// </fieldset>
func FieldSet(children ...html.Node) html.Node {
    return html.NewTag("fieldset", children...)
}

// FigCaption constructs an html.Node for the `<figcaption>` tag.
// 
// The <figcaption> tag is used to provide a caption or description for an
// HTML figure element. It is primarily used to add textual context or explanation
// to an accompanying image, diagram, or illustration. The <figcaption> tag
// appears as a block-level element directly following the <figure> element.
// It helps improve accessibility and allows screen readers to provide
// appropriate descriptions for visually impaired users.
// 
// Example Usage:
// <figure>
// <img src="image.jpg" alt="Description of the image">
// <figcaption>A beautiful sunset over the ocean.</figcaption>
// </figure>
func FigCaption(children ...html.Node) html.Node {
    return html.NewTag("figcaption", children...)
}

// Figure constructs an html.Node for the `<figure>` tag.
// 
// The <figure> tag is used to embed media content, such as images or videos,
// within an HTML document. It provides a semantic way to associate a caption
// or description with the media element. The <figure> tag is often used in
// conjunction with the <figcaption> tag, which is used to provide a caption
// for the media content.
// 
// Example Usage:
// <figure>
// <img src="image.jpg" alt="An example image">
// <figcaption>This is a caption for the image.</figcaption>
// </figure>
func Figure(children ...html.Node) html.Node {
    return html.NewTag("figure", children...)
}

// Font constructs an html.Node for the `<font>` tag.
// 
// The <font> tag is used to define the font style for text in an HTML
// document. It allows developers to specify attributes such as font size,
// color, and face. However, the <font> tag is now deprecated in HTML5 and
// should be avoided in favor of CSS to style text. The use of CSS provides a
// more flexible and efficient way to customize the appearance of text in HTML.
// 
// Example Usage:
// <p style="font-size: 20px;">This text has a font size of 20 pixels.</p>
// <p style="color: red;">This text is displayed in red color.</p>
// <p style="font-family: Arial;">This text is displayed in the Arial font.</p>
func Font(children ...html.Node) html.Node {
    return html.NewTag("font", children...)
}

// Footer constructs an html.Node for the `<footer>` tag.
// 
// The `<footer>` tag is used to define the footer section of an HTML document.
// It is used to present information or content that is typically situated at
// the bottom of a webpage, such as copyright information, contact details, or
// navigation links. The `<footer>` tag is a semantic element that helps define
// the structure and layout of the webpage, and it often appears inside the
// `<body>` element. It can contain various HTML elements, such as headings,
// paragraphs, links, or even other nested tags.
// 
// Example Usage:
// ```html
// <footer>
// <p>© 2022 My Website. All rights reserved.</p>
// <nav>
// <ul>
// <li><a href="#">Home</a></li>
// <li><a href="#">About</a></li>
// <li><a href="#">Contact</a></li>
// </ul>
// </nav>
// </footer>
// ```
func Footer(children ...html.Node) html.Node {
    return html.NewTag("footer", children...)
}

// Form constructs an html.Node for the `<form>` tag.
// 
// The <form> tag is used to create a form in an HTML document. It is primarily
// used to collect user input, such as text, selection options, checkboxes, or
// radio buttons. The <form> tag acts as a container for form elements and
// provides a way to send data to a server for processing. It requires the
// `action` attribute to specify the URL where the form data should be
// submitted, and the `method` attribute to define the HTTP method to be used
// (usually GET or POST). When the form is submitted, the browser will send the
// data to the specified URL and display the response, which can be a new
// webpage or a message.
// 
// Examples:
// <form action="/submit" method="post">
// <label for="name">Name:</label>
// <input type="text" id="name" name="name" required>
// <label for="email">Email:</label>
// <input type="email" id="email" name="email" required>
// <button type="submit">Submit</button>
// </form>
// 
// <form action="/search" method="get">
// <input type="text" name="query" placeholder="Search...">
// <button type="submit">Search</button>
// </form>
func Form(children ...html.Node) html.Node {
    return html.NewTag("form", children...)
}

// Frame constructs an html.Node for the `<frame>` tag.
// 
// The <frame> tag is used to define an individual frame within a frameset in an
// HTML document. It is primarily used to create a layout with multiple sections
// or windows, each displaying separate web pages or content. The <frame> tag
// requires the `src` attribute to specify the URL of the content to be displayed
// in that frame. The frameset element contains multiple <frame> tags to define
// the layout and structure of the frames. Each frame can be resized, scrolled,
// and targeted individually for specific actions.
// 
// Example Usage:
// <frameset rows="50%, 50%">
// <frame src="top.html">
// <frame src="bottom.html">
// </frameset>
func Frame(children ...html.Node) html.Node {
    return html.NewTag("frame", children...)
}

// FrameSet constructs an html.Node for the `<frameset>` tag.
// 
// The <frameset> tag is used to define a set of frames within an HTML
// document. It allows for the creation of a multi-window layout, with each
// frame displaying separate web pages or content. The <frameset> tag is
// primarily used in conjunction with the <frame> tag to specify the individual
// frames and their attributes. It provides a way to divide the browser window
// into multiple sections, each displaying its own HTML document or resource.
// 
// Example Usage:
// <frameset cols="25%,75%">
// <frame src="menu.html" />
// <frame src="content.html" />
// </frameset>
func FrameSet(children ...html.Node) html.Node {
    return html.NewTag("frameset", children...)
}

// Head constructs an html.Node for the `<head>` tag.
// 
// The <head> tag is used to define the head section of an HTML document. It
// contains metadata and other non-visible information about the document, such
// as the document title, character encoding, and linked style sheets. This tag
// does not display any visible content on a webpage, but it plays a crucial
// role in providing instructions and information to web browsers and search
// engines.
// 
// Example Usage:
// <!DOCTYPE html>
// <html>
// <head>
// <title>My Webpage</title>
// <meta charset="UTF-8">
// <link rel="stylesheet" href="styles.css">
// </head>
// <body>
// ...
// </body>
// </html>
func Head(children ...html.Node) html.Node {
    return html.NewTag("head", children...)
}

// Header constructs an html.Node for the `<header>` tag.
// 
// The <header> tag is used to define the introductory or navigational section
// of a document or section. It typically contains the logo, title, and
// navigation elements of a webpage. The <header> tag helps structure the page
// and improves accessibility for users. It is often placed at the top of the
// document or within specific sections.
// 
// Example Usage:
// <header>
// <h1>Welcome to Example Website</h1>
// <nav>
// <ul>
// <li><a href="home.html">Home</a></li>
// <li><a href="about.html">About</a></li>
// <li><a href="products.html">Products</a></li>
// <li><a href="contact.html">Contact</a></li>
// </ul>
// </nav>
// </header>
func Header(children ...html.Node) html.Node {
    return html.NewTag("header", children...)
}

// HGroup constructs an html.Node for the `<hgroup>` tag.
// 
// The <hgroup> tag is used to group heading elements together in an HTML
// document. It is primarily used to create a hierarchical structure for
// headings, allowing for better organization and readability. The <hgroup> tag
// can be used to group multiple <h1> to <h6> tags within a section or article,
// indicating that they are related and part of the same heading group.
// 
// Example Usage:
// <hgroup>
// <h1>Main Heading</h1>
// <h2>Subheading 1</h2>
// <h3>Subheading 2</h3>
// </hgroup>
// 
// In this example, the <hgroup> tag is used to group the main heading and its
// subheadings together, indicating their hierarchical relationship.
func HGroup(children ...html.Node) html.Node {
    return html.NewTag("hgroup", children...)
}

// H1 constructs an html.Node for the `<h1>` tag.
//
// The <h1> tag is used in HTML to define the highest level of headings on a
// webpage. It is primarily used to denote the main title or heading of a page,
// making it important for SEO as search engines pay special attention to <h1>
// content when indexing a webpage. It should ideally be used only once in a
// document, however, it's not an error to use it more than once. Other heading
// tags (<h2> to <h6>) are used to denote subheadings of decreasing importance.
//
// Example Usage:
// <h1>Main Title of the Page</h1>
// <p>This is the introductory paragraph.</p>
// <h2>Subheading of the Page</h2>
// <p>This is the content for the subheading.</p>
func H1(children ...html.Node) html.Node {
    return html.NewTag("h1", children...)
}

// H2 constructs an html.Node for the `<h2>` tag.
//
// The <h2> tag is used in HTML to define the second level heading of a webpage
// or section. It is typically used to break up content into easily readable
// sections, creating a structured outline for the content. Headings are
// important for accessibility, as they enable users with screen readers to
// navigate a webpage with speech output. The <h2> tags are styled in a larger
// and bolder font than normal text, but smaller than the primary <h1> tag.
//
// Example Usage:
// <h2>Section Title</h2>
// <p>This is the content of the section.</p>
func H2(children ...html.Node) html.Node {
    return html.NewTag("h2", children...)
}

// H3 constructs an html.Node for the `<h3>` tag.
//
// The <h3> tag is used in HTML to denote a third level heading, with <h1>
// being the highest and most important and <h6> being the least. It provides a
// hierarchy and structure to the web content, making it easier for users and
// search engines to understand the document structure. The content within the
// <h3> tag automatically appears in bold and slightly larger font compared to
// normal text, but its appearance can be customized using CSS.
//
// Example Usage:
// <h3>This is a Third Level Heading</h3>
// <p>And this is some paragraph text below the heading.</p>
func H3(children ...html.Node) html.Node {
    return html.NewTag("h3", children...)
}

// H4 constructs an html.Node for the `<h4>` tag.
//
// The <h4> tag is used in HTML to define a level four heading, which is
// typically smaller than <h1>, <h2>, and <h3> headings. It is used to group
// content under a common heading to improve the structure and readability of
// the web page. The <h4> tag inherently includes a line break after the
// heading and browsers, by default, make the text enclosed in <h4> tags bold
// and smaller than <h3>.
//
// Example Usage:
// <h4>This is a fourth level heading</h4>
// <p>And this is some paragraph text beneath the heading.</p>
func H4(children ...html.Node) html.Node {
    return html.NewTag("h4", children...)
}

// H5 constructs an html.Node for the `<h5>` tag.
//
// The <h5> tag is used to define a level five heading in an HTML document. It
// is primarily used to create sub-section headings, being the fifth in
// hierarchy after <h1>, <h2>, <h3>, and <h4> tags. The content inside <h5>
// tags is displayed in a smaller font than preceding heading tags, attracting
// less attention.
//
// Example Usage:
// <h5>This is a Level Five Heading</h5>
// <h5>Another Level Five Heading</h5>
func H5(children ...html.Node) html.Node {
    return html.NewTag("h5", children...)
}

// HTML constructs an html.Node for the `<html>` tag.
// 
// The <html> tag is the root element of an HTML document. It defines the
// entire content of the document and acts as a container for all other HTML
// elements. It provides the structure and metadata of the webpage, including
// the title, language, scripts, and stylesheets. The <html> tag must enclose
// the entire HTML content and is located at the beginning and end of the
// document.
// 
// Example Usage:
// <!DOCTYPE html>
// <html lang="en">
// <head>
// <title>My Webpage</title>
// </head>
// <body>
// <h1>Welcome to My Webpage!</h1>
// <p>This is the main content of the webpage.</p>
// </body>
// </html>
func HTML(children ...html.Node) html.Node {
    return html.NewTag("html", children...)
}

// IFrame constructs an html.Node for the `<iframe>` tag.
// 
// The <iframe> tag is used to embed another HTML document within the current
// document. It creates a rectangular area on the page that displays content
// from a different source, such as a separate webpage or a video. This allows
// the integration of external content into a single page, making it versatile
// for displaying maps, videos, or other websites. The content within the
// <iframe> tag is typically defined by the `src` attribute, which specifies
// the URL of the content to be embedded.
// 
// Example Usage:
// <iframe src="https://www.example.com"></iframe>
func IFrame(children ...html.Node) html.Node {
    return html.NewTag("iframe", children...)
}

// Ins constructs an html.Node for the `<ins>` tag.
// 
// The <ins> tag is used to mark inserted text in an HTML document. It is
// primarily used to indicate that content has been added or inserted into the
// document after it was originally written. The text within the <ins> tag is
// typically displayed with an underline, although the appearance can be
// modified using CSS. This tag is often used in collaborative editing, version
// control, or to highlight changes or additions made to a document.
// 
// Example Usage:
// <p>This is some <ins>new</ins> text that was added.</p>
// <p>The <ins>latest</ins> version of the document includes additional information.</p>
func Ins(children ...html.Node) html.Node {
    return html.NewTag("ins", children...)
}

// Kbd constructs an html.Node for the `<kbd>` tag.
// 
// The <kbd> tag is used to define keyboard input in an HTML document. It is
// primarily used to display text or code that represents keyboard input, such
// as keys or key combinations. The <kbd> tag does not affect the style or
// appearance of the text, but it indicates to users that the content should be
// entered using a keyboard. It is often used in tutorials, documentation, or
// forms to instruct users on what to type.
// 
// Example Usage:
// <p>To save the document, press <kbd>Ctrl</kbd> + <kbd>S</kbd>.</p>
// <p>To navigate to the next page, press the <kbd>→</kbd> key.</p>
func Kbd(children ...html.Node) html.Node {
    return html.NewTag("kbd", children...)
}

// Label constructs an html.Node for the `<label>` tag.
// 
// The <label> tag is used to associate text with an input or form element in
// an HTML document. Its purpose is to provide a textual description or caption
// for the associated element, enhancing the accessibility and user experience
// of the form. The text enclosed within the <label> tags is typically
// displayed adjacent to or above the associated element, providing context or
// prompting for user input.
// 
// Example Usage:
// <label for="name">Name:</label>
// <input type="text" id="name" name="name">
// 
// <label for="email">Email:</label>
// <input type="email" id="email" name="email">
// 
// In the examples above, the <label> tags are used to label the input fields
// for name and email. When a user clicks on the label text, focus is
// automatically set to the associated input field, making it easier for users
// to interact with the form.
func Label(children ...html.Node) html.Node {
    return html.NewTag("label", children...)
}

// Legend constructs an html.Node for the `<legend>` tag.
// 
// The <legend> tag is used to provide a caption or title for a fieldset
// element in an HTML form. It helps to provide a clear and concise description
// of the purpose or content of the group of form fields. The <legend> tag is
// typically placed immediately after the opening <fieldset> tag and before the
// form fields it represents.
// 
// Example Usage:
// <fieldset>
// <legend>Personal Information</legend>
// <label for="firstName">First Name:</label>
// <input type="text" id="firstName" name="firstName"><br>
// <label for="lastName">Last Name:</label>
// <input type="text" id="lastName" name="lastName"><br>
// </fieldset>
func Legend(children ...html.Node) html.Node {
    return html.NewTag("legend", children...)
}

// LI constructs an html.Node for the `<li>` tag.
// 
// The <li> tag is used to create a list item in an HTML document. It is
// primarily used within the <ol> (ordered list) or <ul> (unordered list) tags
// to define individual items in a list. The <li> tag is automatically indented
// and displayed with a bullet point or numbering, depending on the parent list
// type. It helps to structure and organize content, making it easier for users
// to read and understand.
// 
// Example Usage:
// <ol>
// <li>First item</li>
// <li>Second item</li>
// <li>Third item</li>
// </ol>
// 
// <ul>
// <li>Apple</li>
// <li>Orange</li>
// <li>Banana</li>
// </ul>
func LI(children ...html.Node) html.Node {
    return html.NewTag("li", children...)
}

// Main constructs an html.Node for the `<main>` tag.
// 
// The <main> tag is used to define the main content of an HTML document. It is
// often used to encapsulate the central content of a webpage, such as articles,
// blogs, or news sections. The <main> tag provides semantic meaning to the content
// it contains, allowing browsers and assistive technologies to better understand
// the structure of the page. It is recommended to only use one <main> tag per
// document for better accessibility and SEO.
// 
// Example Usage:
// <main>
// <h1>Welcome to Our Website</h1>
// <p>Explore our products and services.</p>
// </main>
func Main(children ...html.Node) html.Node {
    return html.NewTag("main", children...)
}

// Map constructs an html.Node for the `<map>` tag.
// 
// The <map> tag is used to create an image map in an HTML document. It is
// primarily used to define clickable areas on an image, associating specific
// coordinates with different links or actions. The <map> tag requires the use
// of the <area> tag, which specifies the shape and coordinates of each
// clickable area. When a user clicks on a defined area of the image, the
// browser will either navigate to the specified URL or trigger a specified
// action. This tag is particularly useful for creating interactive images or
// diagrams.
// 
// Example Usage:
// <img src="diagram.jpg" alt="Diagram" usemap="#diagram-map">
// <map name="diagram-map">
// <area shape="rect" coords="0,0,100,100" href="page1.html" alt="Link 1">
// <area shape="circle" coords="150,150,50" href="page2.html" alt="Link 2">
// </map>
func Map(children ...html.Node) html.Node {
    return html.NewTag("map", children...)
}

// Mark constructs an html.Node for the `<mark>` tag.
// 
// The <mark> tag is used to highlight or mark specific text within an HTML
// document. It applies a yellow background color to the enclosed text, making
// it stand out visually. The <mark> tag does not affect the semantics or
// structure of the document, but it provides a visual cue to users, indicating
// important or relevant information.
// 
// Example Usage:
// <p>This sentence contains <mark>important</mark> information.</p>
// <span>Please <mark>highlight</mark> the keyword in the paragraph.</span>
// <blockquote><mark>Quote</mark> of the day: "Stay positive and keep moving forward."</blockquote>
func Mark(children ...html.Node) html.Node {
    return html.NewTag("mark", children...)
}

// Menu constructs an html.Node for the `<menu>` tag.
// 
// The <menu> tag is used to define a list of commands or choices in an HTML
// document. It is primarily used to create a menu or navigation bar for the
// user to interact with. The <menu> tag can contain various types of menu
// items such as links, buttons, or options. The behavior of the menu items can
// be defined using CSS or JavaScript to provide interactive functionality such
// as dropdown menus or toggling visibility.
// 
// Example Usage:
// <menu>
// <li><a href="#">Home</a></li>
// <li><a href="#">Products</a></li>
// <li><a href="#">Services</a></li>
// </menu>
func Menu(children ...html.Node) html.Node {
    return html.NewTag("menu", children...)
}

// MenuItem constructs an html.Node for the `<menuitem>` tag.
// 
// The <menuitem> tag is used to define a command or action within a context
// menu or toolbar menu in an HTML document. It provides a clickable option
// that triggers a specific action when selected by the user. The <menuitem>
// tag can be used together with the <menu> element to create custom menus with
// various options.
// 
// Example Usage:
// <menu>
// <menuitem label="Save" onclick="save()">Save</menuitem>
// <menuitem label="Copy" onclick="copy()">Copy</menuitem>
// <menuitem label="Paste" onclick="paste()">Paste</menuitem>
// </menu>
func MenuItem(children ...html.Node) html.Node {
    return html.NewTag("menuitem", children...)
}

// Meter constructs an html.Node for the `<meter>` tag.
// 
// The `<meter>` tag is used to represent a scalar measurement or a value
// within a known range in an HTML document. It is primarily used to display
// the level or progress of a specific metric, such as the completion
// percentage of a task or the volume level of an audio player. The `<meter>`
// tag requires the `value` attribute to specify the current value of the
// measurement, and the `min` and `max` attributes to define the minimum and
// maximum values of the range. The browser renders the `<meter>` tag as a
// graphical representation, typically as a horizontal bar with a filled
// portion indicating the current value in relation to the range.
// 
// Example Usage:
// <meter value="75" min="0" max="100">75%</meter>
// <meter value="4" min="0" max="10">40%</meter>
func Meter(children ...html.Node) html.Node {
    return html.NewTag("meter", children...)
}

// Nav constructs an html.Node for the `<nav>` tag.
// 
// The <nav> tag is used to define a section of an HTML document that contains
// navigation links. It is primarily used to create a navigation menu or toolbar,
// providing a convenient way for users to navigate between different pages or
// sections within the website. The <nav> tag helps organize and structure the
// navigation elements, making it easier for users to find and access content.
// 
// Example Usage:
// <nav>
// <ul>
// <li><a href="/home">Home</a></li>
// <li><a href="/about">About</a></li>
// <li><a href="/services">Services</a></li>
// <li><a href="/contact">Contact</a></li>
// </ul>
// </nav>
func Nav(children ...html.Node) html.Node {
    return html.NewTag("nav", children...)
}

// NoFrames constructs an html.Node for the `<noframes>` tag.
// 
// The <noframes> tag is used to provide alternative content for browsers that
// do not support frames. It is primarily used in older HTML documents to
// display a message or instructions for users who are unable to view the
// content within frames. The content within the <noframes> tag is displayed
// when the browser does not support frames or when the user has disabled frame
// rendering.
// 
// Example Usage:
// <noframes>
// <p>Your browser does not support frames. Please upgrade to a modern browser to view this website.</p>
// </noframes>
func NoFrames(children ...html.Node) html.Node {
    return html.NewTag("noframes", children...)
}

// NoScript constructs an html.Node for the `<noscript>` tag.
// 
// The <noscript> tag is used to define content that should be displayed if
// the browser does not support JavaScript, or if JavaScript is disabled.
// It allows developers to provide alternative non-scripted content for users
// who cannot or choose not to use JavaScript. The content within the
// <noscript> tag is only shown when JavaScript is not available or enabled.
// This tag can be useful for displaying important information or instructions
// for users who would otherwise be unable to access or interact with certain
// parts of a website.
// 
// Example Usage:
// <noscript>
// <p>Please enable JavaScript to use the interactive features of this website.</p>
// </noscript>
func NoScript(children ...html.Node) html.Node {
    return html.NewTag("noscript", children...)
}

// Object constructs an html.Node for the `<object>` tag.
// 
// The <object> tag is used to embed external content, such as images, videos,
// or interactive media, into an HTML document. It provides compatibility with
// various file formats and allows for seamless integration of multimedia
// elements. The <object> tag requires the `data` attribute to specify the
// location of the external content, and can also include alternative content,
// displayed if the browser does not support the embedded content.
// 
// Example Usage:
// <object data="image.jpg" type="image/jpeg">
// <img src="fallback.jpg" alt="Fallback Image">
// </object>
// 
// <object data="video.mp4" type="video/mp4">
// <img src="fallback.png" alt="Fallback Image">
// </object>
func Object(children ...html.Node) html.Node {
    return html.NewTag("object", children...)
}

// OL constructs an html.Node for the `<ol>` tag.
// 
// The <ol> tag is used to create an ordered list in an HTML document. It
// defines a numbered list of items, where each item is represented by an
// <li> tag. The list items are automatically numbered sequentially by default,
// but the start attribute can be used to specify a custom starting number.
// The <ol> tag provides structure and organization to a list of items,
// making it easier for users to comprehend and navigate through the content.
// 
// Example Usage:
// <ol>
// <li>First item</li>
// <li>Second item</li>
// <li>Third item</li>
// </ol>
func OL(children ...html.Node) html.Node {
    return html.NewTag("ol", children...)
}

// OptGroup constructs an html.Node for the `<optgroup>` tag.
// 
// The <optgroup> tag is used to group related options within a select dropdown
// menu. It provides a way to organize and categorize options, making it easier
// for users to navigate and select the desired option. The <optgroup> tag
// requires the use of the <select> tag and can contain one or multiple
// <option> tags nested within it.
// 
// Example Usage:
// <select>
// <optgroup label="Fruits">
// <option value="apple">Apple</option>
// <option value="banana">Banana</option>
// </optgroup>
// <optgroup label="Vegetables">
// <option value="carrot">Carrot</option>
// <option value="lettuce">Lettuce</option>
// </optgroup>
// </select>
func OptGroup(children ...html.Node) html.Node {
    return html.NewTag("optgroup", children...)
}

// Option constructs an html.Node for the `<option>` tag.
// 
// The <option> tag is used to define an individual option within a <select> or
// <datalist> element in an HTML document. It is primarily used to provide
// predefined choices in dropdown menus or autocompletion lists. Each <option>
// tag represents a selectable option, and the content within the tag is
// displayed as the option's label. The <option> tag can be used in conjunction
// with the <optgroup> tag to group related options together.
// 
// Example Usage:
// <select>
// <option value="apple">Apple</option>
// <option value="banana">Banana</option>
// <option value="orange">Orange</option>
// </select>
// 
// <datalist>
// <option value="red">Red</option>
// <option value="green">Green</option>
// <option value="blue">Blue</option>
// </datalist>
func Option(children ...html.Node) html.Node {
    return html.NewTag("option", children...)
}

// Output constructs an html.Node for the `<output>` tag.
// 
// The <output> tag is used to display the result of a computation or
// calculation in an HTML document. It is primarily used in forms or
// interactive applications where user input is processed and an output is
// generated. The <output> tag can be used to dynamically update the displayed
// output based on user interactions or changes to the input values. It is
// often used in conjunction with JavaScript to perform calculations and update
// the displayed result in real-time.
// 
// Example Usage:
// <form oninput="result.value=parseInt(a.value)+parseInt(b.value)">
// <input type="number" id="a" name="a">
// +
// <input type="number" id="b" name="b">
// =
// <output name="result" for="a b"></output>
// </form>
func Output(children ...html.Node) html.Node {
    return html.NewTag("output", children...)
}

// P constructs an html.Node for the `<p>` tag.
//
// The <p> tag is used to define a paragraph of text in an HTML document. It is
// primarily used to separate and format blocks of text, creating visually
// distinct sections. The <p> tag automatically adds additional space before and
// after the paragraph, giving it a clear separation from other elements on the page.
//
// Example Usage:
// <p>This is a paragraph of text.</p>
// <p>Here is another paragraph.</p>
func P(children ...html.Node) html.Node {
    return html.NewTag("p", children...)
}

// Picture constructs an html.Node for the `<picture>` tag.
// 
// The <picture> tag is used to define multiple sources of an image and specify
// which one should be displayed based on the device and screen size. It is
// primarily used for responsive images, allowing web developers to provide
// different image files for different devices and resolutions. The <picture> tag
// contains one or more <source> tags, each with a different image source and
// media conditions. The browser will choose the appropriate image source based on
// the available media conditions and display that image.
// 
// Example Usage:
// <picture>
// <source srcset="image-small.jpg" media="(max-width: 600px)">
// <source srcset="image-medium.jpg" media="(max-width: 1200px)">
// <img src="image-large.jpg" alt="Description of the image">
// </picture>
func Picture(children ...html.Node) html.Node {
    return html.NewTag("picture", children...)
}

// Pre constructs an html.Node for the `<pre>` tag.
// 
// The <pre> tag is used to preserve and display the formatting of the text
// within it. It is primarily used for displaying code snippets, poetry, or any
// text that requires a fixed-width font and line breaks to maintain its intended
// formatting. The <pre> tag preserves whitespace and line breaks, allowing for
// the accurate representation of preformatted text.
// 
// Example Usage:
// <pre>
// function helloWorld() {
// console.log("Hello, world!");
// }
// </pre>
// 
// <pre>
// This is a
// multi-line
// text
// </pre>
func Pre(children ...html.Node) html.Node {
    return html.NewTag("pre", children...)
}

// Progress constructs an html.Node for the `<progress>` tag.
// 
// The <progress> tag is used to represent the progress of a specific task or
// completion of a process in an HTML document. It provides a visual indicator,
// typically displayed as a bar, that visually represents the progress of the
// task. The <progress> tag requires the `value` attribute to specify the current
// progress value, and the `max` attribute to define the maximum value of the
// progress. The browser will automatically update the appearance of the
// <progress> element based on the provided values.
// 
// Example Usage:
// <progress value="50" max="100"></progress>
// <progress value="75" max="100"></progress>
func Progress(children ...html.Node) html.Node {
    return html.NewTag("progress", children...)
}

// RP constructs an html.Node for the `<rp>` tag.
// 
// The <rp> tag is used to provide fallback content for browsers that do not
// support the <ruby> tag. It is primarily used in ruby annotations to display
// parentheses or symbols that surround the ruby text, such as the phonetic
// pronunciation of a character in East Asian languages. The content within the
// <rp> tag is typically hidden by default and only displayed if the browser
// cannot render the ruby annotations.
// 
// Example Usage:
// <ruby>
// 漢 <rp>(</rp><rt>Kan</rt><rp>)</rp>
// 字 <rp>(</rp><rt>ji</rt><rp>)</rp>
// </ruby> (Kanji)
// <ruby>
// 動 <rp>(</rp><rt>dō</rt><rp>)</rp>
// 物 <rp>(</rp><rt>mono</rt><rp>)</rp>
// </ruby> (Dōbutsu)
func RP(children ...html.Node) html.Node {
    return html.NewTag("rp", children...)
}

// RT constructs an html.Node for the `<rt>` tag.
// 
// The <rt> tag is used to define the pronunciation of characters in ruby text
// annotations in East Asian typography. It is primarily used for presenting
// supplemental pronunciation information alongside the main text. The <rt> tag is
// typically used in combination with the <ruby> tag, where the main text is
// enclosed within the <ruby> tags and the pronunciation guide is enclosed within
// the <rt> tags. The <rt> tag does not affect the style or appearance of the text,
// but it provides a way to indicate the correct way to pronounce certain
// characters or words.
// 
// Example Usage:
// <ruby>
// 漢<rt>かん</rt>字<rt>じ</rt>
// </ruby>
// <p>The above example shows the pronunciation of the characters "漢字" as "かんじ".</p>
func RT(children ...html.Node) html.Node {
    return html.NewTag("rt", children...)
}

// Ruby constructs an html.Node for the `<ruby>` tag.
// 
// The <ruby> tag is used to add ruby annotations, also known as furigana, to
// text in an HTML document. Ruby annotations are small phonetic characters that
// are typically used in East Asian languages to provide pronunciation guides for
// difficult or uncommon kanji characters. The <ruby> tag consists of a base text
// element and a <rt> element that contains the ruby text. The ruby text is
// usually displayed above or beside the base text, providing additional reading
// assistance. This tag is primarily used for displaying furigana in Japanese web
// pages.
// 
// Example Usage:
// <ruby>日本語<rt>にほんご</rt></ruby>
// <ruby>漢字<rt>かんじ</rt></ruby>
func Ruby(children ...html.Node) html.Node {
    return html.NewTag("ruby", children...)
}

// Samp constructs an html.Node for the `<samp>` tag.
// 
// The `<samp>` tag is used to indicate sample output or example code in an
// HTML document. It is primarily used to display text that represents the
// output of a program or the expected result of a code snippet. The `<samp>`
// tag can be used to distinguish sample output from regular text, making it
// easier for developers to identify and understand the expected output. It
// does not affect the style or appearance of the text, but it conveys its
// purpose to the reader.
// 
// Example Usage:
// ```
// <p>The <samp>print()</samp> function displays text on the console.</p>
// <p>The output of the program is <samp>Hello, World!</samp>.</p>
// ```
func Samp(children ...html.Node) html.Node {
    return html.NewTag("samp", children...)
}

// Script constructs an html.Node for the `<script>` tag.
// 
// The <script> tag is used to embed or reference JavaScript code in an HTML
// document. It allows for dynamic and interactive functionality on webpages.
// JavaScript code placed within the <script> tags can manipulate the content
// and structure of the HTML document, interact with user input, make AJAX
// requests to servers, and perform various other actions. The <script> tag can
// be placed in the head or body section of an HTML document, but it is often
// best practice to place it at the end of the body to ensure that the
// JavaScript executes after the HTML has been loaded.
// 
// Example Usage:
// <script>
// function myFunction() {
// document.getElementById("demo").innerHTML = "Hello, world!";
// }
// </script>
// 
// <script src="script.js"></script>
func Script(children ...html.Node) html.Node {
    return html.NewTag("script", children...)
}

// Section constructs an html.Node for the `<section>` tag.
// 
// The <section> tag is used to define a section within an HTML document. It
// helps in creating a logical grouping of related content, making it easier to
// organize and style the document. The <section> tag is often used to divide
// the webpage into distinct parts or chapters, such as introduction, main
// content, and conclusion. It is a semantic element that adds structure and
// meaning to the content.
// 
// Example Usage:
// <section>
// <h2>About Me</h2>
// <p>Here is some information about myself.</p>
// <p>...</p>
// </section>
// 
// <section>
// <h2>Contact Information</h2>
// <p>Phone: 123-456-7890</p>
// <p>Email: example@example.com</p>
// </section>
func Section(children ...html.Node) html.Node {
    return html.NewTag("section", children...)
}

// Select constructs an html.Node for the `<select>` tag.
// 
// The `<select>` tag is used to create a dropdown menu in an HTML document. It
// allows users to select one option from a list of available choices. The
// options are defined using the `<option>` tag within the `<select>` tag. When
// a user selects an option, the value of the selected option is sent to the
// server when the form is submitted, or it can be accessed using JavaScript.
// The selected option is visually displayed in the dropdown menu.
// 
// Example Usage:
// <select>
// <option value="option1">Option 1</option>
// <option value="option2">Option 2</option>
// <option value="option3">Option 3</option>
// </select>
func Select(children ...html.Node) html.Node {
    return html.NewTag("select", children...)
}

// Slot constructs an html.Node for the `<slot>` tag.
//
// The `<slot>` tag is used in the context of Web Components and the Shadow DOM
// in HTML. It is a placeholder inside a web component where you can insert
// other markup. It allows you to compose complex UIs. When creating reusable
// components, the content inside `<slot>` can be replaced during instantiation
// of the component. The replaced content is referred to as "slotable" and is
// composed by the rendering engine to render inside the shadow DOM tree.
//
// Example Usage:
// <my-custom-element>
// <div slot="mySlot">This is some slotted content.</div>
// </my-custom-element>
//
// <!-- In the shadow DOM of my-custom-element -->
// <div>
// <slot name="mySlot"></slot>
// </div>
func Slot(children ...html.Node) html.Node {
    return html.NewTag("slot", children...)
}

// Small constructs an html.Node for the `<small>` tag.
// 
// The <small> tag is used to indicate that the enclosed text should be
// displayed in a smaller font size compared to the surrounding text. It is
// commonly used for legal disclaimers, copyright information, or fine print.
// 
// Example Usage:
// <p>This website is protected by <small>copyright</small>.</p>
// <p><small>Please read the terms and conditions before proceeding.</small></p>
func Small(children ...html.Node) html.Node {
    return html.NewTag("small", children...)
}

// Span constructs an html.Node for the `<span>` tag.
// 
// The <span> tag is used to group inline elements and apply styles or
// manipulate their content. It does not have any inherent styling or semantic
// meaning. Instead, it serves as a container for styling purposes or to target
// specific sections of text within larger elements. It is commonly used with
// CSS to target and apply styles to specific portions of text within a
// paragraph or heading.
// 
// Example Usage:
// <p>This is a <span style="color: red;">red</span> word.</p>
// <p><span class="highlight">Highlighted</span> text within a paragraph.</p>
func Span(children ...html.Node) html.Node {
    return html.NewTag("span", children...)
}

// Strike constructs an html.Node for the `<strike>` tag.
// 
// The <strike> tag is used to draw a line through the text in an HTML
// document, indicating that it is no longer valid or has been deleted. This
// tag is primarily used to show that certain content is deprecated or should
// be disregarded. It can also be used for crossed-out text in annotations or
// strike-through styling. The <strike> tag does not affect the semantic
// meaning of the text, but it provides visual emphasis.
// 
// Example Usage:
// <p>This is <strike>incorrect</strike> deprecated information.</p>
// <span>This item is <strike>sold out</strike>.</span>
func Strike(children ...html.Node) html.Node {
    return html.NewTag("strike", children...)
}

// Strong constructs an html.Node for the `<strong>` tag.
// 
// The <strong> tag is used to indicate that the text within it should be
// displayed as strong emphasis. It is primarily used to highlight important
// information or emphasize certain words or phrases in an HTML document. The
// <strong> tag does not affect the structure or layout of the text, but it
// visually makes the text appear bold. It is commonly used for headings, key
// points, or emphasizing important messages.
// 
// Example Usage:
// <p>This is a <strong>strongly emphasized</strong> text.</p>
// <h1><strong>This is a heading with strong emphasis</strong></h1>
func Strong(children ...html.Node) html.Node {
    return html.NewTag("strong", children...)
}

// Style constructs an html.Node for the `<style>` tag.
// 
// The <style> tag is used to define the style rules for HTML elements in an
// HTML document. It allows developers to specify the appearance of elements
// using CSS (Cascading Style Sheets). The <style> tag must be placed within
// the <head> section of the HTML document, and the CSS rules enclosed within
// the <style> tags will apply to all elements within the document that match
// the specified selectors.
// 
// Example Usage:
// ```
// <head>
// <style>
// body {
// background-color: lightblue;
// font-family: Arial, sans-serif;
// }
// 
// h1 {
// color: red;
// text-align: center;
// }
// </style>
// </head>
// 
// <body>
// <h1>Welcome to my website!</h1>
// <p>This is the content of my website.</p>
// </body>
// ```
func Style(children ...html.Node) html.Node {
    return html.NewTag("style", children...)
}

// Sub constructs an html.Node for the `<sub>` tag.
// 
// The <sub> tag is used to render subscript text in an HTML document. It is
// primarily used to display characters or symbols that should appear below the
// baseline, indicating a smaller or lower position. The <sub> tag is commonly
// used in mathematical equations, chemical formulas, footnotes, and copyright
// or trademark symbols. When used, the text enclosed by the <sub> tag will be
// displayed with a smaller font size and in a slightly lower position.
// 
// Example Usage:
// <p>The chemical formula for water is H<sub>2</sub>O.</p>
// <p>The equation is a<sub>n</sub> = a<sub>1</sub>(r<sup>n-1</sup>)</p>
func Sub(children ...html.Node) html.Node {
    return html.NewTag("sub", children...)
}

// Summary constructs an html.Node for the `<summary>` tag.
// 
// The <summary> tag is used to provide a summary or caption for a details
// element in an HTML document. It is primarily used to give a brief overview
// or title to expandable content, such as an accordion or dropdown menu. When
// the details element is collapsed, the content inside the <summary> tag is
// displayed as the clickable title. When clicked, the content inside the
// details element is revealed or hidden, depending on its previous state. The
// <summary> tag does not affect the style or appearance of the text by default
// but can be customized using CSS.
// 
// Example Usage:
// <details>
// <summary>Click here for more information</summary>
// <p>This is the hidden content that will be revealed when the summary is clicked.</p>
// </details>
func Summary(children ...html.Node) html.Node {
    return html.NewTag("summary", children...)
}

// Sup constructs an html.Node for the `<sup>` tag.
// 
// The `<sup>` tag is used to create superscript text in an HTML document. It
// is primarily used to display smaller, raised text that appears above the
// regular line of text. This is commonly used to represent footnotes, mathematical
// exponents, and citations. The `<sup>` tag does not alter the appearance of the
// text, but it visually indicates that the content is of smaller size and
// positioned above.
// 
// Example Usage:
// The equation is 2<sup>3</sup> = 8.
// H<sup>2</sup>O is the chemical formula for water.
func Sup(children ...html.Node) html.Node {
    return html.NewTag("sup", children...)
}

// SVG constructs an html.Node for the `<svg>` tag.
// 
// The <svg> tag is used to embed scalable vector graphics (SVG) in an HTML
// document. It allows for the creation of high-quality, resolution-independent
// graphics that can be scaled and manipulated without losing image quality. The
// <svg> tag provides a container for drawing shapes, lines, text, and other
// graphic elements using XML-based syntax. It supports various attributes and
// style properties to control the appearance and behavior of the graphics.
// 
// Example usage:
// <svg width="200" height="200">
// <rect x="50" y="50" width="100" height="100" fill="red" />
// <circle cx="150" cy="150" r="50" fill="blue" />
// <text x="100" y="180" fill="white">SVG Example</text>
// </svg>
func SVG(children ...html.Node) html.Node {
    return html.NewTag("svg", children...)
}

// Table constructs an html.Node for the `<table>` tag.
// 
// The <table> tag is used to create tabular data in an HTML document. It
// organizes data into rows and columns, allowing for structured presentation.
// The <table> tag contains one or more <tr> (table row) tags, which in turn
// contain one or more <td> (table data) or <th> (table header) tags to define
// individual cells. The <th> tag is used for header cells that provide labels
// for the corresponding columns or rows. The <caption> tag can be used to
// provide a title or description for the table. The <table> tag also supports
// various attributes like `border`, `bgcolor`, and `width` to customize the
// appearance and formatting of the table.
// 
// Example Usage:
// <table>
// <tr>
// <th>Name</th>
// <th>Age</th>
// </tr>
// <tr>
// <td>John</td>
// <td>25</td>
// </tr>
// <tr>
// <td>Jane</td>
// <td>30</td>
// </tr>
// </table>
func Table(children ...html.Node) html.Node {
    return html.NewTag("table", children...)
}

// TBody constructs an html.Node for the `<tbody>` tag.
// 
// The <tbody> tag is used to group the body content of an HTML table. It is
// essential for organizing and structuring tabular data. The <tbody> tag
// should be used inside the <table> tag and typically contains multiple <tr>
// (table row) tags. By default, the content within <tbody> will be displayed
// as regular table rows. However, it allows for specifying additional styles
// or behavior if needed, like applying different background colors or hiding
// rows with CSS or JavaScript.
// 
// Example Usage:
// <table>
// <thead>
// <tr>
// <th>Header 1</th>
// <th>Header 2</th>
// </tr>
// </thead>
// <tbody>
// <tr>
// <td>Row 1, Cell 1</td>
// <td>Row 1, Cell 2</td>
// </tr>
// <tr>
// <td>Row 2, Cell 1</td>
// <td>Row 2, Cell 2</td>
// </tr>
// </tbody>
// </table>
func TBody(children ...html.Node) html.Node {
    return html.NewTag("tbody", children...)
}

// TD constructs an html.Node for the `<td>` tag.
// 
// The <td> tag is used to define a cell in an HTML table. It represents a
// single data entry or piece of content within a row. The <td> tag is
// typically used within a <tr> (table row) element to create a grid-like
// structure where each cell contains information. It allows for the
// organization and alignment of data in columns and rows.
// 
// Example Usage:
// <table>
// <tr>
// <td>This cell contains data</td>
// <td>Another cell with data</td>
// </tr>
// <tr>
// <td>A third cell</td>
// <td>And yet another cell</td>
// </tr>
// </table>
func TD(children ...html.Node) html.Node {
    return html.NewTag("td", children...)
}

// Template constructs an html.Node for the `<template>` tag.
// 
// The <template> tag is used to define reusable content that can be
// cloned and inserted into an HTML document. It allows developers to define
// a section of markup that is not rendered immediately, but can be used
// dynamically. The content within a <template> tag is hidden by default and
// can be accessed and rendered using JavaScript. This tag is often used to
// create reusable components or dynamic content that needs to be generated
// programmatically.
// 
// Example Usage:
// <template id="myTemplate">
// <div>
// <h1>This is a template</h1>
// <p>Content inside the template.</p>
// </div>
// </template>
// 
// <script>
// const template = document.getElementById("myTemplate");
// const clone = template.content.cloneNode(true);
// document.body.appendChild(clone);
// </script>
func Template(children ...html.Node) html.Node {
    return html.NewTag("template", children...)
}

// TextArea constructs an html.Node for the `<textarea>` tag.
// 
// The <textarea> tag is used to create a multi-line text input field in an
// HTML document. It allows users to enter and edit large amounts of text. The
// <textarea> tag includes attributes such as rows and cols to control the size
// of the input field. Any text placed between the opening and closing
// <textarea> tags will be displayed as the initial value of the input field.
// 
// Example Usage:
// <textarea rows="4" cols="50">
// This is a sample text that can be edited.
// </textarea>
func TextArea(children ...html.Node) html.Node {
    return html.NewTag("textarea", children...)
}

// TFoot constructs an html.Node for the `<tfoot>` tag.
// 
// The <tfoot> tag is used to define a footer for a table in an HTML document.
// It is used to group and describe the footer content, such as summary
// information, totals, or any other relevant details related to the table. The
// <tfoot> tag is placed inside the <table> element, after the <tbody> or
// <thead> tags. It helps in structuring the table structure by clearly
// separating the header, body, and footer sections.
// 
// Example Usage:
// <table>
// <thead>
// <tr>
// <th>Product</th>
// <th>Price</th>
// </tr>
// </thead>
// <tbody>
// <tr>
// <td>Product 1</td>
// <td>$10.00</td>
// </tr>
// <tr>
// <td>Product 2</td>
// <td>$15.00</td>
// </tr>
// </tbody>
// <tfoot>
// <tr>
// <td>Total:</td>
// <td>$25.00</td>
// </tr>
// </tfoot>
// </table>
func TFoot(children ...html.Node) html.Node {
    return html.NewTag("tfoot", children...)
}

// TH constructs an html.Node for the `<th>` tag.
// 
// The <th> tag is used to define a header cell in an HTML table. It identifies
// a cell as a header, which is typically used to label or describe the content
// in a specific table column or row. The <th> tag can be used within the
// <thead> or <tbody> section of a table.
// 
// Example Usage:
// <tr>
// <th>Name</th>
// <th>Age</th>
// <th>Gender</th>
// </tr>
func TH(children ...html.Node) html.Node {
    return html.NewTag("th", children...)
}

// THead constructs an html.Node for the `<thead>` tag.
// 
// The <thead> tag is used to define the header section of a table in an HTML
// document. It is responsible for grouping and identifying the header row or
// rows in a table. The content placed within the <thead> tag typically
// consists of one or more <tr> (table row) tags containing <th> (table header
// cell) tags. The <thead> tag helps organize and structure the table by
// separating the header information from the main body of the table, improving
// readability and accessibility.
// 
// Example Usage:
// <table>
// <thead>
// <tr>
// <th>Name</th>
// <th>Age</th>
// <th>Gender</th>
// </tr>
// </thead>
// <tbody>
// <tr>
// <td>John Doe</td>
// <td>25</td>
// <td>Male</td>
// </tr>
// <tr>
// <td>Jane Smith</td>
// <td>30</td>
// <td>Female</td>
// </tr>
// </tbody>
// </table>
func THead(children ...html.Node) html.Node {
    return html.NewTag("thead", children...)
}

// Time constructs an html.Node for the `<time>` tag.
// 
// The <time> tag is used to mark up a specific point in time or a duration in
// an HTML document. It helps to semantically represent dates, times, and
// durations on a webpage. This tag can be beneficial for search engines and
// assistive technologies in understanding and presenting time-related
// information accurately. The datetime attribute is used to specify the date,
// time, or duration in a machine-readable format.
// 
// Example Usage:
// <p>Today is <time datetime="2022-11-13">November 13, 2022</time>.</p>
// <p>The event will start at <time datetime="18:30">6:30 PM</time>.</p>
// <p>The video duration is <time datetime="PT5M30S">5 minutes and 30 seconds</time>.</p>
func Time(children ...html.Node) html.Node {
    return html.NewTag("time", children...)
}

// Title constructs an html.Node for the `<title>` tag.
// 
// The <title> tag is used to define the title of an HTML document. It is placed
// within the <head> section of the document and is displayed as the title of the
// browser window or tab. The text within the <title> tag is also used by search
// engines to display the title of the webpage in search results. There should be
// only one <title> tag in each HTML document.
// 
// Example Usage:
// <!DOCTYPE html>
// <html>
// <head>
// <title>My Website</title>
// </head>
// <body>
// ...
// </body>
// </html>
func Title(children ...html.Node) html.Node {
    return html.NewTag("title", children...)
}

// TR constructs an html.Node for the `<tr>` tag.
// 
// The <tr> tag is used to define a row in an HTML table. It is primarily used
// to structure tabular data by separating it into rows. Each <tr> tag contains
// one or more <td> (table data) or <th> (table header) tags, which represent
// the cells within the row. The <td> tags contain the actual data, while the
// <th> tags are used to display header cells. The <tr> tag ensures that the
// table data is displayed in a structured manner with consistent row
// alignment.
// 
// Example Usage:
// <table>
// <tr>
// <th>Header 1</th>
// <th>Header 2</th>
// </tr>
// <tr>
// <td>Data 1</td>
// <td>Data 2</td>
// </tr>
// </table>
func TR(children ...html.Node) html.Node {
    return html.NewTag("tr", children...)
}

// TT constructs an html.Node for the `<tt>` tag.
// 
// The `<tt>` tag is used to display text in a monospaced or "typewriter" font
// style in an HTML document. It is primarily used to indicate code snippets or to
// highlight text that needs to stand out from the rest of the content. The `<tt>`
// tag does not influence the semantic meaning of the text, but it conveys to
// users and developers that the content represents code or a computer terminal
// output. It is often used in programming tutorials, documentation, or to
// emphasize specific instructions.
// 
// Example Usage:
// <tt>printf("Hello, World!");</tt>
// <tt><input type="text" value="Username"></tt>
func TT(children ...html.Node) html.Node {
    return html.NewTag("tt", children...)
}

// UL constructs an html.Node for the `<ul>` tag.
// 
// The <ul> tag is used to create an unordered list in an HTML document.
// It is used to group related items together and display them as a list.
// The <ul> tag does not specify the order or sequence of the items, and
// each item is typically represented by a <li> (list item) tag. The
// unordered list is displayed with bullet points by default, but the
// appearance can be customized using CSS.
// 
// Example usage:
// <ul>
// <li>Item 1</li>
// <li>Item 2</li>
// <li>Item 3</li>
// </ul>
func UL(children ...html.Node) html.Node {
    return html.NewTag("ul", children...)
}

// Var constructs an html.Node for the `<var>` tag.
// 
// The <var> tag is used to display variables or placeholders in code or
// mathematical expressions. It indicates to developers that the content has a
// special meaning, but it does not affect the style or appearance. It is
// commonly used in computer programming documentation or mathematical
// formulas.
// 
// Example usage:
// <p>The value of <var>x</var> is 5.</p>
// <p>The result of <var>a</var> + <var>b</var> is <var>c</var>.</p>
func Var(children ...html.Node) html.Node {
    return html.NewTag("var", children...)
}

// Video constructs an html.Node for the `<video>` tag.
// 
// The <video> tag is used to embed a video or audio file into an HTML
// document. It allows for the playback of media content directly within the
// web page. The <video> tag requires the 'src' attribute to specify the URL or
// file path of the video or audio file. It supports various file formats such
// as MP4, WebM, and Ogg. Developers can also customize the playback controls,
// autoplay behavior, and other settings through additional attributes and
// JavaScript. This tag has greatly enhanced the multimedia capabilities of
// HTML, enabling seamless integration of video and audio content on websites.
// 
// Example Usage:
// <video src="video.mp4" controls>
// Your browser does not support the video tag.
// </video>
func Video(children ...html.Node) html.Node {
    return html.NewTag("video", children...)
}

// WBr constructs an html.Node for the `<wbr>` tag.
//
// The <wbr> tag in HTML is used to specify a possible line-break. It primarily
// allows the browser to break a line at the specified point if the line is too
// long for its container. The <wbr> tag does not have any visual
// representation unless it enforces a line break. It is mainly useful when
// dealing with long words or strings that could potentially break the layout
// if not broken into multiple lines.
//
// Example Usage:
// <p>This is a verylongwordthatmightneeda<wbr>break.</p>
func WBr(children ...html.Node) html.Node {
    return html.NewTag("wbr", children...)
}
