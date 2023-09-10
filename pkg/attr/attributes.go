package attr

import "github.com/jeffswenson/sanity/pkg/html"

// Abbr constructs an html.Node for the `abbr` attribute.
//
// The `abbr` attribute is used to specify an abbreviation or acronym for an
// HTML element. It helps provide a concise description or clarification of the
// element's content. When the abbreviation or acronym is added to an element
// with the `abbr` attribute, browsers and assistive technologies can display a
// tooltip or provide additional context for the abbreviation. This attribute
// is particularly useful for improving accessibility and enhancing the
// understanding of complex terms or technical jargon.
//
// Example Usage:
// <abbr title="World Wide Web Consortium">W3C</abbr>
// <p>The <abbr title="HyperText Markup Language">HTML</abbr> standard is used for creating web pages.</p>
func Abbr(value string) html.Node {
	return html.NewAttribute("abbr", value)
}

// Accept constructs an html.Node for the `accept` attribute.
//
// The `accept` attribute is used to specify the types of files that can be
// uploaded through an HTML file input element. It allows developers to
// restrict the file types that users can select for uploading, providing a
// more controlled and secure file input experience. The value of the `accept`
// attribute should be a comma-separated list of file extensions or MIME types.
//
// Example usage:
// <input type="file" accept=".jpg, .png">
func Accept(value string) html.Node {
	return html.NewAttribute("accept", value)
}

// AcceptCharSet constructs an html.Node for the `accept-charset` attribute.
//
// `accept-charset` is used to specify the character encodings that are
// accepted by the server when submitting a form. It allows the browser to
// inform the server about the character encoding used in the submitted form
// data, ensuring that the server can correctly interpret and process the data.
// The value of the `accept-charset` attribute is a space-separated list of
// character encoding names.
//
// Example Usage:
// <form action="/submit" method="post" accept-charset="UTF-8">
// <!-- Form fields here -->
// <input type="submit" value="Submit">
// </form>
func AcceptCharSet(value string) html.Node {
	return html.NewAttribute("accept-charset", value)
}

// AccessKey constructs an html.Node for the `accesskey` attribute.
//
// `accesskey` is used to specify a keyboard shortcut for quickly accessing an
// HTML element. It allows users to navigate through the webpage using keyboard
// shortcuts instead of relying on a mouse. When the user presses the specified
// key combination (usually a letter or a number) along with the defined
// accesskey attribute, the focus is moved to the associated element. This
// attribute improves accessibility and makes navigation more efficient for
// users who may have difficulty using a mouse.
//
// Example Usage:
// <input type="text" accesskey="u">This input field can be accessed quickly with the 'Alt+U' key combination.
// <button accesskey="s">This button can be triggered using the 'Alt+S' key combination.
// <textarea accesskey="e">This textarea can be focused with 'Alt+E' key combination.
func AccessKey(value string) html.Node {
	return html.NewAttribute("accesskey", value)
}

// Action constructs an html.Node for the `action` attribute.
//
// `action` is used to specify the URL of the server-side script or program that
// will process the data submitted through an HTML form. It is primarily used in
// the `<form>` element and determines where the form data will be sent for
// processing. When the form is submitted, the browser will navigate to the URL
// specified in the `action` attribute, sending the form data along with it. This
// allows the server to receive and process the data, and provide a response back
// to the user.
//
// Example Usage:
// <form action="/submit-form" method="POST">
// <!-- Form inputs go here -->
// <input type="submit" value="Submit">
// </form>
func Action(value string) html.Node {
	return html.NewAttribute("action", value)
}

// Allow constructs an html.Node for the `allow` attribute.
//
// `allow` is used in HTML5 to specify the types of content that are allowed to
// be displayed or executed within an `<iframe>` element. This attribute
// provides an additional layer of security by allowing the website developer
// to restrict the actions that can be performed within the embedded content.
// The value of the `allow` attribute can include one or more of the following
// keywords: `accelerometer`, `autoplay`, `camera`, `encrypted-media`,
// `fullscreen`, `geolocation`, `gyroscope`, `magnetometer`, `microphone`,
// `midi`, `payment`, and `picture-in-picture`. Each keyword represents a
// specific capability that can be enabled or disabled within the `<iframe>`.
//
// Example Usage:
// <iframe src="https://www.example.com"
// allow="accelerometer; fullscreen; camera;"></iframe>
func Allow(value string) html.Node {
	return html.NewAttribute("allow", value)
}

// Alt constructs an html.Node for the `alt` attribute.
//
// `alt` is used to provide alternative text for an image in an HTML document.
// This text is displayed if the image fails to load or cannot be displayed for
// some reason. The `alt` attribute is important for accessibility, as it allows
// screen readers to describe the image to visually impaired users. It also helps
// with search engine optimization (SEO) by providing relevant information about
// the image to search engines. The value of the `alt` attribute should be brief
// but descriptive, conveying the purpose or content of the image.
//
// Example Usage:
// <img src="example.jpg" alt="A beautiful sunset over the ocean">
func Alt(value string) html.Node {
	return html.NewAttribute("alt", value)
}

// As constructs an html.Node for the `as` attribute.
//
// `as` is used to specify the expected media type or format of a linked resource
// in an HTML document. It is primarily used in the `link` and `script` tags to
// inform the browser how to handle or interpret the resource. This attribute helps
// optimize the loading and rendering of web content by enabling the browser to
// begin processing the resource even before it is fully downloaded. Common values
// for the `as` attribute include `image`, `style`, `font`, `script`, and `fetch`.
//
// Example Usage:
// <link rel="stylesheet" href="styles.css" as="style">
// <script src="script.js" as="script"></script>
// <img src="image.jpg" alt="Example Image" as="image">
func As(value string) html.Node {
	return html.NewAttribute("as", value)
}

// AutoCapitalize constructs an html.Node for the `autocapitalize` attribute.
//
// `autocapitalize` is used to specify whether or not text input in an HTML
// element should be automatically capitalized. This attribute is particularly
// useful for input fields where the user is expected to enter text in a
// specific capitalization style, such as names or addresses. The value of the
// `autocapitalize` attribute can be set to various options, such as "none" (to
// disable automatic capitalization), "sentences" (to capitalize the first letter
// of each sentence), "words" (to capitalize the first letter of each word), or
// "characters" (to capitalize every character).
//
// Example Usage:
// <input type="text" autocapitalize="words" placeholder="Enter your name">
// <input type="text" autocapitalize="none" placeholder="Enter a lowercase email">
func AutoCapitalize(value string) html.Node {
	return html.NewAttribute("autocapitalize", value)
}

// AutoComplete constructs an html.Node for the `autocomplete` attribute.
//
// The `autocomplete` attribute is used to specify whether or not an input
// field should have autocomplete functionality enabled. Autocomplete
// functionality provides suggestions or predictions as the user types, based
// on previously entered values or a predefined list. This attribute can have
// the values "on" or "off" to control the behavior of autocomplete.
//
// Example Usage:
// <input type="text" name="username" autocomplete="off">
// <input type="password" name="password" autocomplete="on">
func AutoComplete(value string) html.Node {
	return html.NewAttribute("autocomplete", value)
}

// Blocking constructs an html.Node for the `blocking` attribute.
//
// The `blocking` attribute is used to indicate whether a specified resource
// should block the rendering of the HTML document until it is fully loaded or
// not. By default, all resources, such as stylesheets and scripts, are
// considered blocking, meaning that the loading of the resource will delay the
// rendering of the page until it is finished loading. However, by setting the
// `blocking` attribute to "none" on a particular resource, it allows the HTML
// document to continue rendering, while the resource is fetched in the
// background.
//
// Example Usage:
// <link href="styles.css" rel="stylesheet" blocking="none">
// <script src="script.js" blocking="none"></script>
func Blocking(value string) html.Node {
	return html.NewAttribute("blocking", value)
}

// CharSet constructs an html.Node for the `charset` attribute.
//
// The `charset` attribute is used to specify the character encoding for an
// HTML document. It tells the browser how to interpret the text within the
// document, ensuring that special characters and symbols are displayed
// correctly. The value of the `charset` attribute is usually set to a specific
// encoding, such as UTF-8 or ISO-8859-1.
//
// Example Usage:
// <meta charset="UTF-8">
func CharSet(value string) html.Node {
	return html.NewAttribute("charset", value)
}

// Cite constructs an html.Node for the `cite` attribute.
//
// `cite` is used to specify the source or reference for a quoted or cited
// content within an HTML document. It is primarily used in blockquote or q tags
// to provide a link or reference to the original author, publication, or source
// of the quoted text. The `cite` attribute helps to attribute and provide
// credibility to the quoted content.
//
// Example Usage:
// <blockquote cite="https://www.example.com/article">Lorem ipsum dolor sit
// amet.</blockquote>
func Cite(value string) html.Node {
	return html.NewAttribute("cite", value)
}

// Class constructs an html.Node for the `class` attribute.
//
// The `class` attribute is used to assign one or more class names to an HTML
// element, allowing developers to apply CSS styles or JavaScript functionality
// to specific elements. This attribute helps in customizing the appearance and
// behavior of the element by associating it with predefined styles and
// behaviors defined in CSS or JavaScript files. The value of the `class`
// attribute is a space-separated list of class names, enabling multiple styles
// or behaviors to be applied simultaneously.
//
// Example Usage:
// <div class="red-text bold">This div has the classes 'red-text' and 'bold' applied to it.</div>
func Class(value string) html.Node {
	return html.NewAttribute("class", value)
}

// Color constructs an html.Node for the `color` attribute.
//
// The `color` attribute is used to specify the text color of an HTML element.
// It allows developers to customize the appearance of text by setting a
// specific color. The value of the `color` attribute can be a named color
// (e.g. "red", "blue"), a hexadecimal color code (e.g. "#FF0000"), or an RGB
// value (e.g. "rgb(255, 0, 0)"). By applying the `color` attribute to an HTML
// element, the text within that element will be displayed in the specified
// color.
//
// Example Usage:
// <p style="color: red;">This text is displayed in red.</p>
func Color(value string) html.Node {
	return html.NewAttribute("color", value)
}

// Cols constructs an html.Node for the `cols` attribute.
//
// `cols` is used to specify the number of columns in an HTML table. It
// determines the layout and organization of data within the table, allowing
// developers to define the width of each column. The value of the `cols`
// attribute is an integer that represents the number of columns. This
// attribute helps in creating structured and organized tables with consistent
// column widths.
//
// Example Usage:
// <table cols="3">
// <tr>
// <td>Column 1</td>
// <td>Column 2</td>
// <td>Column 3</td>
// </tr>
// <tr>
// <td>Data 1</td>
// <td>Data 2</td>
// <td>Data 3</td>
// </tr>
// </table>
func Cols(value string) html.Node {
	return html.NewAttribute("cols", value)
}

// ColSpan constructs an html.Node for the `colspan` attribute.
//
// `colspan` is used to specify the number of columns a cell should span in an
// HTML table. It allows for the merging of multiple columns in a single cell,
// creating a wider cell that spans across multiple columns. This attribute is
// helpful when there is a need to combine multiple cells horizontally to display
// a larger block of content or data within a table.
//
// Example Usage:
// <td colspan="2">This cell spans across two columns.</td>
func ColSpan(value string) html.Node {
	return html.NewAttribute("colspan", value)
}

// Content constructs an html.Node for the `content` attribute.
//
// The `content` attribute is used to specify the content for a specific
// HTML element. It is commonly used with the `meta` element to provide
// additional information about the webpage, such as the author, description,
// or keywords. The value of the `content` attribute can vary depending on
// the context and purpose of the element it is used with.
//
// Example Usage:
// <meta name="description" content="This is a description of the webpage.">
// <meta name="keywords" content="html, attribute, content, example">
func Content(value string) html.Node {
	return html.NewAttribute("content", value)
}

// ContentEditable constructs an html.Node for the `contenteditable` attribute.
//
// The `contenteditable` attribute is used to make an HTML element editable by
// the user. When this attribute is set to "true", the element can be modified
// directly on the webpage, allowing users to input text or make changes to the
// content. This attribute is commonly used in web applications or content
// management systems where users need to edit or update text without using a
// separate text editor.
//
// Example Usage:
// <div contenteditable="true">This is an editable div where users can modify the content.</div>
// <span contenteditable="true">Users can type directly into this span to add or edit text.</span>
func ContentEditable(value string) html.Node {
	return html.NewAttribute("contenteditable", value)
}

// Coords constructs an html.Node for the `coords` attribute.
//
// The `coords` attribute is used to specify the coordinates of
// specified shape in an image map. It is primarily used in the `area` element
// when defining clickable areas within an image. The `coords` attribute
// takes a comma-separated list of coordinates that define the shape and size
// of the clickable area. The exact format of the coordinates depends on the
// shape being used (e.g., rectangular, circular, or polygonal). The values
// represent percentage or pixel values relative to the dimensions of the
// image.
//
// Example Usage:
// <map name="image-map">
// <area shape="circle" coords="50,50,30" alt="Circle" href="circle.html">
// <area shape="rect" coords="10,10,100,100" alt="Rectangle" href="rectangle.html">
// <area shape="polygon" coords="10,10,100,10,100,100,10,100" alt="Polygon" href="polygon.html">
// </map>
// <img src="image.jpg" usemap="#image-map">
func Coords(value string) html.Node {
	return html.NewAttribute("coords", value)
}

// CrossOrigin constructs an html.Node for the `crossorigin` attribute.
//
// `crossorigin` is used to specify how the browser should handle cross-origin
// resource requests when loading an external resource, such as a script or an
// image. This attribute provides a way to control whether the browser should
// send credentials, such as cookies or HTTP authentication, when making the
// request. The `crossorigin` attribute can have one of the following values:
// - "anonymous": The browser will not send any credentials with the request.
// - "use-credentials": The browser will send credentials with the request if
// the origin of the page and the resource being accessed have the same
// origin.
//
// Example Usage:
// <script src="https://example.com/script.js" crossorigin="anonymous"></script>
// <img src="https://example.com/image.jpg" crossorigin="use-credentials">
func CrossOrigin(value string) html.Node {
	return html.NewAttribute("crossorigin", value)
}

// Data constructs an html.Node for the `data` attribute.
//
// `data` is a custom attribute that allows developers to store additional
// information within an HTML element. It is used to attach data to specific
// elements, providing a way to store metadata that is not visible or directly
// relevant to the presentation or behavior of the element. The value of the
// `data` attribute can be any valid string or JSON data, representing various
// types of information such as configuration settings, dynamic data, or
// customized data attributes.
//
// Example Usage:
// <div data-id="1234" data-category="electronics">This div stores data about a product.</div>
// <button data-action="submit-form" data-active="true">This button has data attributes to control its behavior.</button>
func Data(value string) html.Node {
	return html.NewAttribute("data", value)
}

// DateTime constructs an html.Node for the `datetime` attribute.
//
// The `datetime` attribute is used to specify a machine-readable date and time
// value for an HTML element. It is primarily used for semantic markup and
// improving accessibility. The value of the `datetime` attribute should follow
// the ISO 8601 format, providing the date and time in a consistent and
// universal format. This attribute is commonly used in elements such as <time>
// to provide additional context and understanding of a specific date and time.
//
// Example Usage:
// <time datetime="2021-09-30T18:30:00Z">September 30, 2021 at 6:30 PM</time>
func DateTime(value string) html.Node {
	return html.NewAttribute("datetime", value)
}

// Decoding constructs an html.Node for the `decoding` attribute.
//
// The `decoding` attribute is used in HTML to specify how the browser should
// decode and display media files. It allows developers to control how the video,
// audio, or image content is processed and presented to the user. The value of
// the `decoding` attribute can be set to "sync", "async", or "auto". When set to
// "sync", the browser will decode the media file synchronously, meaning it will
// process the file immediately. When set to "async", the browser will decode the
// media file asynchronously, meaning it will process the file in the background
// while the page is loading. Lastly, when set to "auto", the browser will
// determine the optimal decoding method based on the type of media and the user's
// device capabilities.
//
// Example Usage:
// <img src="image.jpg" decoding="auto">
// <video src="video.mp4" decoding="sync"></video>
// <audio src="audio.mp3" decoding="async"></audio>
func Decoding(value string) html.Node {
	return html.NewAttribute("decoding", value)
}

// Dir constructs an html.Node for the `dir` attribute.
//
// `dir` is used to specify the text directionality for the content within an
// HTML element. It allows developers to control the ordering of text and the
// orientation of characters within the element. The `dir` attribute can have two
// possible values: "ltr" (left-to-right) for languages that read from left to
// right, and "rtl" (right-to-left) for languages that read from right to left.
//
// Example Usage:
// <p dir="ltr">This paragraph has left-to-right text direction.</p>
// <p dir="rtl">This paragraph has right-to-left text direction.</p>
func Dir(value string) html.Node {
	return html.NewAttribute("dir", value)
}

// DirName constructs an html.Node for the `dirname` attribute.
//
// `dirname` is used to specify the text directionality of the content within
// an HTML element. It is primarily used for languages that are written from
// right-to-left, such as Arabic or Hebrew, to ensure that the text is rendered
// and displayed correctly. The value of the `dirname` attribute can be either
// `ltr` (left-to-right) or `rtl` (right-to-left), indicating the directionality
// of the text within the element.
//
// Example Usage:
// <p dirname="rtl">This paragraph contains right-to-left text.</p>
// <input type="text" dirname="ltr" placeholder="Enter your name">
func DirName(value string) html.Node {
	return html.NewAttribute("dirname", value)
}

// Draggable constructs an html.Node for the `draggable` attribute.
//
// The `draggable` attribute is used to indicate whether an element can be
// dragged by the user. It can be applied to a wide range of HTML elements,
// including images, text, and divs. When set to `true`, the element can be
// dragged and dropped within the same page or between different applications or
// tabs. By default, all elements are not draggable unless the `draggable`
// attribute is explicitly set to `true`.
//
// Example Usage:
// <img src="image.jpg" draggable="true">
// <p draggable="false">This paragraph cannot be dragged by the user.</p>
func Draggable(value string) html.Node {
	return html.NewAttribute("draggable", value)
}

// Enctype constructs an html.Node for the `enctype` attribute.
//
// `enctype` is used to specify how form data should be encoded and sent to the
// server when an HTML form is submitted. It is primarily used in the `<form>`
// element to control how the data is formatted and transmitted. The `enctype`
// attribute is especially important when the form contains file uploads, as it
// determines how the files will be encoded and sent.
//
// Example Usage:
// <form action="/submit" method="post" enctype="multipart/form-data">
// <input type="file" name="file">
// <input type="submit">
// </form>
func Enctype(value string) html.Node {
	return html.NewAttribute("enctype", value)
}

// EnterKeyHint constructs an html.Node for the `enterkeyhint` attribute.
//
// The `enterkeyhint` attribute is used to provide a hint to the browser about the
// expected user action when the "Enter" key is pressed. It helps improve the
// user experience by suggesting the appropriate action, such as submitting a
// form, searching, or creating a new line. This attribute is particularly useful
// for input fields where the default behavior may not be ideal.
//
// Example Usage:
// <input type="text" enterkeyhint="search" placeholder="Search...">
// <input type="text" enterkeyhint="next" placeholder="Next item...">
// <input type="text" enterkeyhint="done" placeholder="Complete task...">
func EnterKeyHint(value string) html.Node {
	return html.NewAttribute("enterkeyhint", value)
}

// FetchPriority constructs an html.Node for the `fetchpriority` attribute.
//
// The `fetchpriority` attribute is used to indicate the priority of fetching a
// resource in an HTML document. This attribute is typically used in the
// `<link>` or `<img>` tags to specify the importance of retrieving the
// resource. By setting a higher priority value for an element, browsers can
// prioritize fetching and rendering that element before others. This can help
// improve the perceived performance of a web page by ensuring that important
// resources are prioritized and loaded quickly.
//
// Example Usage:
// <link rel="stylesheet" href="styles.css" fetchpriority="high">
// <img src="image.jpg" fetchpriority="low">
func FetchPriority(value string) html.Node {
	return html.NewAttribute("fetchpriority", value)
}

// For constructs an html.Node for the `for` attribute.
//
// The `for` attribute is used to create a relationship between a label element
// and another element on the web page. It is primarily used to associate a label
// with a form input element, allowing users to click on the label to activate
// the corresponding input. This improves accessibility and usability by
// increasing the clickable area of the input. The value of the `for` attribute
// should match the `id` attribute of the related element.
//
// Example Usage:
// <label for="name">Name:</label>
// <input type="text" id="name" name="name" placeholder="Enter your name">
func For(value string) html.Node {
	return html.NewAttribute("for", value)
}

// Form constructs an html.Node for the `form` attribute.
//
// The `form` attribute is used to associate an HTML element with a specific
// form. It allows input elements, such as buttons or text fields, to be
// grouped together and submitted as a single form. When the form is submitted,
// the input values from all associated elements are sent to the server.
//
// Example Usage:
// <input type="text" form="myForm" placeholder="Enter your name">
// <button type="submit" form="myForm">Submit</button>
func Form(value string) html.Node {
	return html.NewAttribute("form", value)
}

// FormAction constructs an html.Node for the `formaction` attribute.
//
// `formaction` is used to override the default submission URL of a form when
// the submit button is clicked. It allows developers to specify a different URL
// that will receive the form data upon submission. This attribute is typically
// used in conjunction with the `form` and `input` elements to control the
// behavior of form submission.
//
// Example Usage:
// <input type="submit" formaction="/submit-form">
func FormAction(value string) html.Node {
	return html.NewAttribute("formaction", value)
}

// FormEncType constructs an html.Node for the `formenctype` attribute.
//
// `formenctype` is used to specify the encoding type to be used when submitting
// data from an HTML form to the server. It allows developers to indicate whether
// the data should be encoded as `application/x-www-form-urlencoded` (the default)
// or as `multipart/form-data`. The `application/x-www-form-urlencoded` encoding
// is used for sending simple form data, while the `multipart/form-data` encoding
// is used for sending file uploads or binary data. The value of the `formenctype`
// attribute should be set to either `application/x-www-form-urlencoded` or
// `multipart/form-data`.
//
// Example Usage:
// <form method="post" action="/submit" formenctype="multipart/form-data">
// <input type="file" name="myfile">
// <input type="submit" value="Submit">
// </form>
func FormEncType(value string) html.Node {
	return html.NewAttribute("formenctype", value)
}

// FormMethod constructs an html.Node for the `formmethod` attribute.
//
// `formmethod` is used to specify the HTTP method to be used when submitting
// a form in an HTML document. This attribute is applied to the `button` or
// `input` elements with a `type` of "submit" or "image". The `formmethod`
// attribute allows developers to override the default method (which is usually
// "GET") and specify a different method such as "POST". This is especially
// useful when submitting sensitive information or when the form data modifies
// server-side resources. The value of the `formmethod` attribute should be a
// valid HTTP method, such as "GET" or "POST".
//
// Example Usage:
// <button formmethod="POST" type="submit">Submit Form</button>
// <input formmethod="DELETE" type="submit" value="Delete Item">
func FormMethod(value string) html.Node {
	return html.NewAttribute("formmethod", value)
}

// FormTarget constructs an html.Node for the `formtarget` attribute.
//
// `formtarget` is used to specify where the form data should be submitted when
// the user submits a form. It overrides the default behavior of submitting the
// form to the same page. The value of the `formtarget` attribute can be a URL
// or one of the following keywords:
//
// - `_blank`: Opens the form response in a new tab or window.
// - `_self`: Loads the form response in the same frame or window.
// - `_parent`: Loads the form response in the parent frame or window.
// - `_top`: Loads the form response in the full body of the window.
//
// Example Usage:
// <form action="/submit" method="post" target="_blank">
// <input type="text" name="name">
// <input type="submit" value="Submit">
// </form>
func FormTarget(value string) html.Node {
	return html.NewAttribute("formtarget", value)
}

// Headers constructs an html.Node for the `headers` attribute.
//
// The `headers` attribute is used to establish a relationship between a data
// cell (`<td>`) and its corresponding header cell (`<th>`) in an HTML table.
// It defines a space-separated list of header cell IDs that the data cell is
// associated with. This allows assistive technologies, such as screen readers,
// to correctly interpret and present tabular data to users with disabilities.
//
// Example Usage:
// <table>
// <tr>
// <th id="name">Name</th>
// <th id="age">Age</th>
// </tr>
// <tr>
// <td headers="name">John Doe</td>
// <td headers="age">25</td>
// </tr>
// </table>
func Headers(value string) html.Node {
	return html.NewAttribute("headers", value)
}

// Height constructs an html.Node for the `height` attribute.
//
// The `height` attribute is used to specify the height of an HTML element. It
// determines the vertical size of the element and can be applied to various
// types of elements such as images, tables, divs, and iframes. The value of
// the `height` attribute can be specified in pixels, percentages, or other CSS
// length units. It allows developers to control the visual layout of elements
// and ensure consistency in the presentation of the web page.
//
// Example Usage:
// <img src="image.jpg" alt="An image" height="200">
// <div style="height: 300px;">This div has a fixed height of 300 pixels.</div>
func Height(value string) html.Node {
	return html.NewAttribute("height", value)
}

// High constructs an html.Node for the `high` attribute.
//
// The `high` attribute is used to specify a numerical value that represents
// the importance or priority of an HTML element. It is primarily used in
// ordered lists (ol) to indicate the level of importance for each list item.
// The value of the `high` attribute should be an integer, with a higher value
// indicating a higher level of importance. This attribute is often used in
// conjunction with CSS to style the list items based on their importance.
//
// Example Usage:
// <ol>
// <li high="3">This is the most important item</li>
// <li high="2">This is a moderately important item</li>
// <li high="1">This is the least important item</li>
// </ol>
func High(value string) html.Node {
	return html.NewAttribute("high", value)
}

// HRef constructs an html.Node for the `href` attribute.
//
// The `href` attribute is used to specify the URL of a linked resource in an
// HTML document. It enables the creation of hyperlinks, allowing users to
// navigate between different pages or sections of a website. The value of the
// `href` attribute acts as the address that the hyperlink points to. When
// clicked, the browser will navigate to the URL specified by the `href`
// attribute value, loading the corresponding webpage. The value can be an
// absolute or relative URL, allowing links to external sites or different
// sections within the same site.
//
// Example Usage:
// <a href="https://www.example.com">This link directs to an external website.</a>
// <a href="/about">This link directs to the 'about' page within the same website.</a>
func HRef(value string) html.Node {
	return html.NewAttribute("href", value)
}

// HRefLang constructs an html.Node for the `hreflang` attribute.
//
// `hreflang` is used to specify the language of the linked resource in an HTML
// document. It is primarily used in anchor (a) tags to provide information
// about the language of the target webpage. This attribute helps search
// engines understand the language of the linked page, allowing them to display
// the correct version of the webpage to users who speak the same language. The
// value of the `hreflang` attribute should be a language code, following the
// ISO 639-1 or ISO 639-2 standards.
//
// Example Usage:
// <a href="https://www.example.com" hreflang="en">This link directs to an English version of the website.</a>
// <a href="https://www.example.es" hreflang="es">Este enlace dirige a la versión en español del sitio web.</a>
func HRefLang(value string) html.Node {
	return html.NewAttribute("hreflang", value)
}

// HttpEquiv constructs an html.Node for the `http-equiv` attribute.
//
// `http-equiv` is used to provide an HTTP header for an HTML document.
// It allows developers to specify information about the document's content type,
// refresh rate, character encoding, and other important metadata that affects
// how the document is interpreted and displayed by the browser. The value of the
// `http-equiv` attribute is a string that corresponds to a specific HTTP header
// field, such as "Content-Type" or "Refresh". This attribute is commonly used in
// legacy HTML documents or in situations where server-side headers cannot be
// modified directly.
//
// Example Usage:
// <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
// <meta http-equiv="Refresh" content="5; URL=https://www.example.com">
func HttpEquiv(value string) html.Node {
	return html.NewAttribute("http-equiv", value)
}

// Id constructs an html.Node for the `id` attribute.
//
// The `id` attribute is used to uniquely identify an HTML element. It allows
// developers to reference and target specific elements in CSS and JavaScript.
// The value of the `id` attribute should be unique within the HTML document.
//
// Example Usage:
// <div id="header">This div represents the header section of the webpage.</div>
// <button id="submitBtn">This button triggers a form submission.</button>
func Id(value string) html.Node {
	return html.NewAttribute("id", value)
}

// ImageSizes constructs an html.Node for the `imagesizes` attribute.
//
// `imagesizes` is used to specify the sizes of the available image sources in
// an HTML document. It helps browsers select the appropriate image source
// based on the device's screen size and resolution, improving performance and
// optimizing the display of images. The value of the `imagesizes` attribute is
// a space-separated list of image widths, allowing developers to provide
// multiple sizes for different screen sizes and resolutions.
//
// Example Usage:
// <img src="image.jpg" alt="Example" imagesizes="320px, 640px, 960px">
// <img srcset="image_small.jpg 320w, image_medium.jpg 640w, image_large.jpg 960w" sizes="(max-width: 600px) 320px, (max-width: 1200px) 640px, 960px" alt="Example">
func ImageSizes(value string) html.Node {
	return html.NewAttribute("imagesizes", value)
}

// ImageSrcSet constructs an html.Node for the `imagesrcset` attribute.
//
// `imagesrcset` is used to specify multiple sources for an image in an HTML
// document. This attribute allows the browser to choose the most appropriate
// image source based on factors like screen resolution or device capabilities.
// The `imagesrcset` attribute is typically used in conjunction with the `src`
// attribute to provide alternative image sources for different scenarios. Each
// source in the `imagesrcset` attribute is defined with a URL and a descriptor
// that specifies the image's width or pixel density. The browser selects the
// source that best matches the conditions and loads that image.
//
// Example Usage:
// <img src="small.jpg" imagesrcset="medium.jpg 800w, large.jpg 1200w">
func ImageSrcSet(value string) html.Node {
	return html.NewAttribute("imagesrcset", value)
}

// InputMode constructs an html.Node for the `inputmode` attribute.
//
// `inputmode` is used to specify the expected input method for an HTML input
// element. It helps optimize the user experience by suggesting the appropriate
// on-screen keyboard layout or input method based on the expected input type.
// By using the `inputmode` attribute, developers can provide better input
// suggestions, autocorrect, and validation for user input. The value of the
// `inputmode` attribute can be set to various values such as "numeric",
// "tel", "email", "url", "search", etc., depending on the type of input expected.
//
// Example Usage:
// <input type="text" inputmode="numeric" placeholder="Enter a number">
// <input type="tel" inputmode="tel" placeholder="Enter a phone number">
// <input type="email" inputmode="email" placeholder="Enter an email address">
func InputMode(value string) html.Node {
	return html.NewAttribute("inputmode", value)
}

// Integrity constructs an html.Node for the `integrity` attribute.
//
// The `integrity` attribute is used to ensure the authenticity and integrity
// of a linked resource in an HTML document. It allows developers to include a
// cryptographic hash value, such as an SHA-256 hash, which is used to verify
// that the resource has not been tampered with or altered. This is
// particularly important for resources like scripts or stylesheets that are
// served from external sources.
//
// Example Usage:
// <script src="https://example.com/script.js" integrity="sha256-ABC123DEF456GHI789"></script>
func Integrity(value string) html.Node {
	return html.NewAttribute("integrity", value)
}

// Is constructs an html.Node for the `is` attribute.
//
// The `is` attribute is used to apply a custom element or behavior to an HTML
// element. It allows developers to define their own custom elements and extend
// the functionality of existing HTML elements. By specifying the value of the
// `is` attribute as the name of a custom element, the HTML element is
// transformed into the custom element with its associated behaviors.
//
// Example Usage:
// <button is="custom-button">This button has custom functionality applied to it.</button>
func Is(value string) html.Node {
	return html.NewAttribute("is", value)
}

// ItemId constructs an html.Node for the `itemid` attribute.
//
// `itemid` is used to specify a unique identifier for an item in an HTML
// document. It is primarily used in conjunction with structured data markup,
// such as the schema.org vocabulary, to provide additional information about the
// item. The `itemid` attribute enables search engines and other applications to
// identify and index specific items on a webpage, improving the visibility and
// understanding of the content.
//
// Example Usage:
// <div itemscope itemtype="http://schema.org/Book">
// <span itemprop="name">The Catcher in the Rye</span>
// <link itemprop="url" href="https://www.example.com/books/catcher-in-the-rye">
// </div>
func ItemId(value string) html.Node {
	return html.NewAttribute("itemid", value)
}

// ItemProp constructs an html.Node for the `itemprop` attribute.
//
// `itemprop` is used to specify a specific property or attribute of an HTML
// element that is part of a structured data set. It is primarily used for
// creating semantic markup and providing context to search engines about the
// content of the element. By using `itemprop`, developers can enhance the
// meaning and relevance of their content for search engine optimization (SEO)
// purposes. The value of the `itemprop` attribute typically corresponds to a
// specific schema.org property, such as "name" or "description".
//
// Example Usage:
// <span itemprop="name">Product Name</span>
// <img itemprop="image" src="product.jpg" alt="Product Image">
func ItemProp(value string) html.Node {
	return html.NewAttribute("itemprop", value)
}

// ItemRef constructs an html.Node for the `itemref` attribute.
//
// The `itemref` attribute is used to create associations between elements in
// an HTML document. It is primarily used in HTML microdata to specify
// additional elements that contain properties related to a main element. By
// referencing the IDs of other elements using the `itemref` attribute, the
// properties from those elements can be included in the main element's
// microdata.
//
// Example Usage:
// <div id="item1" itemscope itemtype="http://schema.org/Book"></div>
// <div id="item2" itemscope itemtype="http://schema.org/Person"></div>
// <div id="mainItem" itemscope itemtype="http://schema.org/Review" itemref="item1 item2"></div>
func ItemRef(value string) html.Node {
	return html.NewAttribute("itemref", value)
}

// ItemType constructs an html.Node for the `itemtype` attribute.
//
// `itemtype` is used to specify the type of an item in an HTML document using a
// URL. It is primarily used in conjunction with the `itemscope` attribute to
// create structured data markup using the Schema.org vocabulary. The value of the
// `itemtype` attribute is a URL that identifies the type of the item, providing
// contextual information about its meaning and properties. This attribute helps
// search engines and other applications understand the content and its
// relationships, contributing to improved search results and enhanced
// presentation of the data.
//
// Example Usage:
// <div itemscope itemtype="http://schema.org/Person">
// <span itemprop="name">John Doe</span>
// <span itemprop="jobTitle">Web Developer</span>
// </div>
func ItemType(value string) html.Node {
	return html.NewAttribute("itemtype", value)
}

// Kind constructs an html.Node for the `kind` attribute.
//
// The `kind` attribute is used to specify the type or category of a media
// resource in an HTML document. It is primarily used in the `<source>` element
// within `<video>` or `<audio>` tags to provide alternative media sources. The
// `kind` attribute allows developers to differentiate between different
// formats or qualities of the media file, such as captions, subtitles, or
// alternative audio tracks. By specifying the `kind` attribute, the browser
// can determine which media source to use based on the user's preferences or
// accessibility requirements.
//
// Example Usage:
// <video>
// <source src="video.mp4" type="video/mp4" kind="main">
// <source src="video.webm" type="video/webm" kind="alternative">
// <track src="video.vtt" kind="captions" srclang="en" label="English">
// </video>
func Kind(value string) html.Node {
	return html.NewAttribute("kind", value)
}

// Label constructs an html.Node for the `label` attribute.
//
// The `label` attribute is used to associate text with form elements in an HTML
// document. It helps improve accessibility by providing a textual description or
// name for the form element, making it easier for users to understand the purpose
// or function of the element. The `label` attribute is usually used in conjunction
// with the `for` attribute, which specifies which form element the label is
// associated with. This allows users to click on the label to activate the
// associated form element, enhancing usability.
//
// Example Usage:
// <label for="username">Username:</label>
func Label(value string) html.Node {
	return html.NewAttribute("label", value)
}

// Lang constructs an html.Node for the `lang` attribute.
//
// `lang` is used to specify the language of the text contained within an
// element. It allows developers to indicate the language of the content to
// support accessibility and SEO purposes. The value of the `lang` attribute
// should be a valid language code defined by the W3C, such as "en" for English
// or "ja" for Japanese.
//
// Example Usage:
// <p lang="fr">Ce paragraphe est écrit en français.</p>
// <span lang="es">Este texto está en español.</span>
func Lang(value string) html.Node {
	return html.NewAttribute("lang", value)
}

// List constructs an html.Node for the `list` attribute.
//
// The `list` attribute is used to associate an input element with a datalist
// element. It allows for autocomplete functionality, where the user can choose
// from a predefined set of options while typing in the input field. The `list`
// attribute's value should be equal to the `id` attribute of the datalist
// element it is associated with. When the user types in the input field, a
// dropdown list of options will appear based on the values specified in the
// associated datalist element. The user can select an option from the
// dropdown, and the selected value will be filled in the input field.
//
// Example Usage:
// <input type="text" list="fruits">
// <datalist id="fruits">
// <option value="Apple">
// <option value="Banana">
// <option value="Orange">
// </datalist>
func List(value string) html.Node {
	return html.NewAttribute("list", value)
}

// Loading constructs an html.Node for the `loading` attribute.
//
// The `loading` attribute is used to control the loading behavior of external
// resources, such as images or scripts, in an HTML document. It determines
// when and how these resources are loaded, allowing developers to optimize
// page loading speed and improve user experience.
//
// Example Usage:
// <img src="image.jpg" loading="lazy">
// <script src="script.js" loading="defer"></script>
func Loading(value string) html.Node {
	return html.NewAttribute("loading", value)
}

// Low constructs an html.Node for the `low` attribute.
//
// The `low` attribute is used to indicate the lower bound value of a range in
// an HTML input element. It is primarily used with the `input` element to
// define the minimum value that can be selected or entered by the user. This
// attribute ensures that the user does not enter a value below the specified
// lower bound.
//
// Example Usage:
// <input type="number" low="0" max="100">
// This input field allows the user to enter a number between 0 and 100,
// inclusive, with 0 being the minimum value.
func Low(value string) html.Node {
	return html.NewAttribute("low", value)
}

// Max constructs an html.Node for the `max` attribute.
//
// The `max` attribute is used to set the maximum value that can be entered or
// selected in an input field or element. It is commonly used with input types
// such as "number" or "date" to define an upper limit for the value that can
// be inputted. The `max` value restricts the user from exceeding the specified
// limit, providing validation and ensuring data integrity.
//
// Example Usage:
// <input type="number" max="100"> // The user can enter a number up to 100.
// <input type="date" max="2022-12-31"> // The user can select a date up to December 31, 2022.
func Max(value string) html.Node {
	return html.NewAttribute("max", value)
}

// MaxLength constructs an html.Node for the `maxlength` attribute.
//
// `maxlength` is used to specify the maximum number of characters allowed in an
// input field in an HTML form. It restricts the user from entering more characters
// than the specified limit. The `maxlength` attribute is commonly used with text
// input types such as `input type="text"` or `input type="password"`.
//
// Example Usage:
// <input type="text" maxlength="10" placeholder="Enter up to 10 characters">
func MaxLength(value string) html.Node {
	return html.NewAttribute("maxlength", value)
}

// Media constructs an html.Node for the `media` attribute.
//
// The `media` attribute is used to specify the media types for which a specific
// CSS style should be applied. It allows developers to control the presentation
// of elements based on the device or medium used to view the web page, such as
// screen, print, or handheld devices. The `media` attribute is commonly used in
// link tags to conditionally load CSS files based on the media type.
//
// Example Usage:
// <link rel="stylesheet" href="styles.css" media="screen">
// <link rel="stylesheet" href="print.css" media="print">
// <link rel="stylesheet" href="mobile.css" media=" handheld">
func Media(value string) html.Node {
	return html.NewAttribute("media", value)
}

// Method constructs an html.Node for the `method` attribute.
//
// The `method` attribute is used in a `<form>` element to specify the HTTP
// request method to be used when submitting the form data to the server. It
// determines how the data from the form will be transmitted. The `method`
// attribute can have two possible values: "GET" or "POST".
//
// The "GET" method appends the form data to the URL in the form of query
// parameters. It is commonly used for retrieving data from the server without
// modifying it.
// Example: `<form method="GET" action="search.php">`
//
// The "POST" method sends the form data in the body of the HTTP request. It is
// commonly used for submitting data to the server that may modify or update
// the server's data.
// Example: `<form method="POST" action="submit.php">`
func Method(value string) html.Node {
	return html.NewAttribute("method", value)
}

// Min constructs an html.Node for the `min` attribute.
//
// `min` is used to specify a minimum value for numerical input fields in an
// HTML form. It restricts the range of acceptable input values to be greater
// than or equal to the specified minimum value. The `min` attribute is
// commonly used with the `input` element, particularly with the `type`
// attribute set to `number` or `date`, but it can also be used with other
// input types such as `range` or `datetime-local`.
//
// Example Usage:
// <input type="number" min="0">    // Input field accepts only positive numbers.
// <input type="date" min="2021-01-01">    // Input field only accepts dates after January 1, 2021.
//
// Note: The `min` attribute alone does not enforce any validation or prevent
// users from manually entering values below the specified minimum. It is
// essential to use JavaScript or HTML5 form validation to ensure the input
// meets the specified criteria.
func Min(value string) html.Node {
	return html.NewAttribute("min", value)
}

// MinLength constructs an html.Node for the `minlength` attribute.
//
// `minlength` is used to specify the minimum number of characters or values
// that should be entered or selected in an HTML input field or textarea. It is
// primarily used to enforce data validation and ensure that a certain level of
// input is provided by the user. The `minlength` attribute works in
// conjunction with other input-related attributes, such as `required`, to
// create more robust and user-friendly forms.
//
// Example Usage:
// <input type="text" minlength="5" required>
// <textarea minlength="10" required></textarea>
func MinLength(value string) html.Node {
	return html.NewAttribute("minlength", value)
}

// Name constructs an html.Node for the `name` attribute.
//
// `name` is used to specify a name for an HTML element, typically used in form
// elements to identify data that will be submitted to a server. The `name`
// attribute provides a way to access and manipulate the value of the element
// through JavaScript or server-side scripts. This attribute is particularly
// important for input elements, such as text fields or checkboxes, as it
// allows the server to identify the data associated with each input.
//
// Example Usage:
// <input type="text" name="username" placeholder="Enter your username">
// <input type="checkbox" name="subscribe" value="yes"> Subscribe to newsletter
func Name(value string) html.Node {
	return html.NewAttribute("name", value)
}

// Nonce constructs an html.Node for the `nonce` attribute.
//
// The `nonce` attribute is used to specify a cryptographic nonce (number used
// once) for inline scripts and styles in an HTML document. It helps prevent
// cross-site scripting (XSS) attacks by ensuring that only trusted scripts and
// styles are executed. The value of the `nonce` attribute should be a randomly
// generated string that is unique for each page load.
//
// Example Usage:
// <script nonce="abc123">This inline script has a unique nonce value.</script>
// <style nonce="def456">This inline style has a unique nonce value.</style>
func Nonce(value string) html.Node {
	return html.NewAttribute("nonce", value)
}

// Optimum constructs an html.Node for the `optimum` attribute.
//
// The `optimum` attribute is used to specify the ideal or optimal value for a
// progress element in an HTML document. It helps define the point at which the
// task or process represented by the progress element is considered complete or
// successful. This attribute is primarily used in conjunction with the `value`
// attribute to provide feedback to users about the progress of a task or
// process.
//
// Example Usage:
// <progress value="50" max="100" optimum="80"></progress>
func Optimum(value string) html.Node {
	return html.NewAttribute("optimum", value)
}

// Pattern constructs an html.Node for the `pattern` attribute.
//
// The `pattern` attribute is used to specify a regular expression pattern that
// an input element's value must match. It is primarily used with text-based input
// fields, such as `<input type="text">` or `<input type="email">`, to enforce a
// specific format or validate user input. The `pattern` attribute helps ensure
// that the data entered by the user follows a specified pattern, such as a
// specific phone number format or a required combination of letters and numbers.
//
// Example Usage:
// <input type="text" pattern="[0-9]{3}-[0-9]{3}-[0-9]{4}">
// <input type="email" pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}">
func Pattern(value string) html.Node {
	return html.NewAttribute("pattern", value)
}

// Ping constructs an html.Node for the `ping` attribute.
//
// The `ping` attribute is used to specify a list of URLs that should receive
// notification when a user interacts with an HTML element. This includes actions
// such as clicking on a link or submitting a form. When the specified event
// occurs, the URLs listed in the `ping` attribute will be alerted, allowing them
// to track and analyze user interactions on the website.
//
// Example Usage:
// <a href="https://www.example.com" ping="https://analytics.example.com">This link directs to an external website and alerts the analytics server.</a>
func Ping(value string) html.Node {
	return html.NewAttribute("ping", value)
}

// Placeholder constructs an html.Node for the `placeholder` attribute.
//
// The `placeholder` attribute is used to provide a hint or example of the
// expected input for an HTML form element. It is commonly used in input fields
// to display a brief description or example of the type of data that should be
// entered. The text specified in the `placeholder` attribute is typically
// displayed in a lighter color and disappears when the user starts typing.
// This attribute improves user experience by providing guidance for filling
// out forms.
//
// Example Usage:
// <input type="text" placeholder="Enter your name">
// <textarea placeholder="Enter your message"></textarea>
func Placeholder(value string) html.Node {
	return html.NewAttribute("placeholder", value)
}

// PopOver constructs an html.Node for the `popover` attribute.
//
// The `popover` attribute is used to create a pop-up dialog or tooltip that
// displays additional information when a user interacts with an element. This
// attribute is typically used in conjunction with JavaScript or CSS to define the
// content and behavior of the popover. When the user hovers over or clicks on an
// element with the `popover` attribute, the popover is triggered and the
// specified content is displayed.
//
// Example Usage:
// <button popover="This is a popover message.">Hover over me</button>
func PopOver(value string) html.Node {
	return html.NewAttribute("popover", value)
}

// PopOverTarget constructs an html.Node for the `popovertarget` attribute.
//
// `popovertarget` is a custom attribute that can be added to HTML elements to
// specify the target element or elements for a popover. A popover is a small,
// dynamically displayed overlay that appears when a user interacts with a
// specific element, usually triggered by a hover or click event. The
// `popovertarget` attribute allows developers to easily associate a popover
// with its corresponding target element, enhancing the user experience by
// providing additional information or functionality.
//
// Example Usage:
// <button popovertarget="popover1">Hover over me to see the popover!</button>
// <div id="popover1" class="popover">This is the content of the popover.</div>
//
// In the example above, the `popovertarget` attribute is added to the button
// element, indicating that the popover with the ID "popover1" should be
// displayed when the button is hovered over. The popover content is contained
// within a div element with the corresponding ID, allowing the CSS styles and
// JavaScript functionality associated with the popover to be easily applied.
func PopOverTarget(value string) html.Node {
	return html.NewAttribute("popovertarget", value)
}

// PopOverTargetAction constructs an html.Node for the `popovertargetaction` attribute.
//
// The `popovertargetaction` attribute is used to specify the action that
// should be performed when a target element is clicked or interacted with to
// open a popover. It allows developers to define custom behavior for popovers,
// such as displaying additional content, triggering animations, or executing
// JavaScript functions.
//
// Example Usage:
// <button popovertargetaction="showPopover()">Click me to open a popover</button>
func PopOverTargetAction(value string) html.Node {
	return html.NewAttribute("popovertargetaction", value)
}

// Poster constructs an html.Node for the `poster` attribute.
//
// The `poster` attribute is used to specify an image that should be displayed
// while a video is loading or before it starts playing. It is primarily used
// in the `<video>` element to provide a visually appealing preview of the
// video content. The value of the `poster` attribute should be the URL of an
// image file.
//
// Example Usage:
// <video poster="video-preview.jpg">
// <source src="video.mp4" type="video/mp4">
// </video>
func Poster(value string) html.Node {
	return html.NewAttribute("poster", value)
}

// PreLoad constructs an html.Node for the `preload` attribute.
//
// The `preload` attribute is used to provide a hint to the browser to load a
// specific resource, such as an audio or video file, before it is actually
// needed. This helps improve performance by reducing the delay in rendering
// media content when it is requested by the user. The value of the `preload`
// attribute can be set to different values to control how and when the
// resource should be preloaded.
//
// Example Usage:
// <video src="video.mp4" preload="auto">
// This video will start preloading as soon as the page loads.
// </video>
//
// <audio src="audio.mp3" preload="metadata">
// Only the metadata of the audio file will be preloaded, not the entire file.
// </audio>
func PreLoad(value string) html.Node {
	return html.NewAttribute("preload", value)
}

// ReferrerPolicy constructs an html.Node for the `referrerpolicy` attribute.
//
// The `referrerpolicy` attribute is used to control the referring information
// that is sent when a user navigates from one webpage to another. It specifies
// the policy that the browser should use when sending the `Referer` header, which
// contains the URL of the webpage that linked to the current webpage. This can be
// used to enhance user privacy and security by controlling the amount of
// information shared with external websites.
//
// Example Usage:
// <a href="https://www.example.com" referrerpolicy="origin">This link will
// only send the origin (domain) of the referring webpage.</a>
// <a href="https://www.example.com" referrerpolicy="no-referrer">This link will
// not send any referring information to the linked webpage.</a>
func ReferrerPolicy(value string) html.Node {
	return html.NewAttribute("referrerpolicy", value)
}

// Rel constructs an html.Node for the `rel` attribute.
//
// `rel` is used to specify the relationship between the current document and
// the linked document in an HTML document. It is primarily used in anchor (a) tags
// to indicate the type of relationship the linked document has with the current
// document, such as "stylesheet" for linking to CSS files, "icon" for specifying
// favicon images, or "canonical" for indicating the preferred version of a page.
// The value of the `rel` attribute can vary depending on the purpose of the link
// and is often combined with other attributes, such as `href` and `type`, to
// provide additional context.
//
// Example Usage:
// <link rel="stylesheet" href="styles.css">
// <link rel="icon" type="image/png" href="favicon.png">
// <link rel="canonical" href="https://www.example.com/main-page.html">
func Rel(value string) html.Node {
	return html.NewAttribute("rel", value)
}

// Rows constructs an html.Node for the `rows` attribute.
//
// The `rows` attribute is used to specify the number of visible rows in a text
// area or a table in HTML. It determines the height of the element, allowing
// users to input or display multiline text or tabular data. The value of the
// `rows` attribute should be a positive integer, indicating the desired number of
// rows to be displayed.
//
// Example Usage:
// <textarea rows="5">This text area has 5 visible rows.</textarea>
// <table>
// <tr>
// <td>Row 1</td>
// </tr>
// <tr>
// <td>Row 2</td>
// </tr>
// <tr>
// <td>Row 3</td>
// </tr>
// </table>
// The table has a default number of visible rows based on the number of table
// rows.
func Rows(value string) html.Node {
	return html.NewAttribute("rows", value)
}

// RowSpan constructs an html.Node for the `rowspan` attribute.
//
// The `rowspan` attribute is used to specify the number of rows that a table
// cell should span vertically. It allows the content of a single cell to
// occupy multiple rows in a table, merging the cells below it. This attribute
// is commonly used in table structures where cells need to span across
// multiple rows or when creating complex table layouts.
//
// Example Usage:
// <td rowspan="2">This cell spans 2 rows.</td>
func RowSpan(value string) html.Node {
	return html.NewAttribute("rowspan", value)
}

// Sandbox constructs an html.Node for the `sandbox` attribute.
//
// The `sandbox` attribute is used to restrict the behavior of an iframe element
// within an HTML document. It creates a secure environment for the embedded
// content, preventing it from accessing or modifying the parent document or
// executing potentially harmful scripts. The `sandbox` attribute can be used with
// four different values:
//
// - If `sandbox` is set to an empty string, it activates all of the available
// restrictions, preventing the iframe from any interaction with the parent
// document.
//
// - If `sandbox` is set to "allow-same-origin", the iframe is allowed to
// navigate within the same origin as the parent document, but still can't
// access or modify it.
//
// - If `sandbox` is set to "allow-scripts", the iframe is allowed to execute
// scripts within its own context, but not in the parent document.
//
// - If `sandbox` is set to "allow-forms", the iframe is allowed to submit forms,
// but not perform other interactions with the parent document.
//
// Example Usage:
// <iframe src="https://www.example.com" sandbox></iframe>
// <iframe src="https://www.example.com" sandbox="allow-same-origin"></iframe>
// <iframe src="https://www.example.com" sandbox="allow-scripts"></iframe>
// <iframe src="https://www.example.com" sandbox="allow-forms"></iframe>
func Sandbox(value string) html.Node {
	return html.NewAttribute("sandbox", value)
}

// Scope constructs an html.Node for the `scope` attribute.
//
// The `scope` attribute is used to specify the scope of data cells in an HTML
// table. It determines whether a header cell applies to a single column, a
// single row, or a group of columns or rows. By defining the scope, assistive
// technologies can properly associate the header cell with its corresponding
// data cells, improving accessibility and usability.
//
// Example Usage:
// <table>
// <thead>
// <tr>
// <th scope="col">Name</th>
// <th scope="col">Age</th>
// </tr>
// </thead>
// <tbody>
// <tr>
// <th scope="row">John</th>
// <td>25</td>
// </tr>
// <tr>
// <th scope="row">Jane</th>
// <td>30</td>
// </tr>
// </tbody>
// </table>
func Scope(value string) html.Node {
	return html.NewAttribute("scope", value)
}

// Shape constructs an html.Node for the `shape` attribute.
//
// The `shape` attribute is used to define the shape of an area in an image map
// in an HTML document. It is primarily used in conjunction with the `coords`
// attribute to create clickable areas within an image. The value of the
// `shape` attribute can be one of the following shapes: "rect" (rectangle),
// "circle" (circle), or "poly" (polygon).
//
// Example Usage:
// <img src="image.jpg" usemap="#myMap" alt="Image">
// <map name="myMap">
// <area shape="rect" coords="0,0,100,100" href="page1.html" alt="Area 1">
// <area shape="circle" coords="150,150,50" href="page2.html" alt="Area 2">
// <area shape="poly" coords="200,200,250,300,300,250,250,200" href="page3.html" alt="Area 3">
// </map>
func Shape(value string) html.Node {
	return html.NewAttribute("shape", value)
}

// Size constructs an html.Node for the `size` attribute.
//
// The `size` attribute is used to specify the visible width, in characters,
// of an input element like text fields. This attribute allows developers to
// control the width of the input field, giving users a visual indication of the
// amount of text that can be entered. The value of the `size` attribute should
// be a positive integer, representing the number of visible characters.
//
// Example Usage:
// <input type="text" size="20">
func Size(value string) html.Node {
	return html.NewAttribute("size", value)
}

// Sizes constructs an html.Node for the `sizes` attribute.
//
// The `sizes` attribute is used to specify the sizes of images or icons in an
// HTML document. It allows the browser to determine the appropriate display
// size for the image based on the device's viewport and screen density. The
// value of the `sizes` attribute is a space-separated list of image sizes,
// each specified as a media condition followed by a corresponding size hint.
// Media conditions can be used to apply different sizes based on factors like
// screen width or resolution. Size hints are specified using the `w`
// descriptor, followed by the width in pixels or a value relative to the
// viewport width. This attribute is commonly used in conjunction with the
// `srcset` attribute to provide responsive images that adapt to different
// screen sizes and resolutions.
//
// Example Usage:
// <img src="image.jpg" sizes="(max-width: 600px) 100vw, (max-width: 1200px) 50vw, 25vw" srcset="image.jpg 1200w, image-m.jpg 600w, image-s.jpg 300w">
// <link href="icon.png" sizes="192x192" rel="icon" type="image/png">
func Sizes(value string) html.Node {
	return html.NewAttribute("sizes", value)
}

// Slot constructs an html.Node for the `slot` attribute.
//
// The `slot` attribute is used to specify where content should be placed
// within a web component. It provides a way to define insertion points within
// the component's template, allowing developers to dynamically inject content
// into specific slots. This attribute is particularly useful in creating
// reusable components with customizable content. By assigning elements to
// different slots, developers can easily customize the layout and structure of
// their components without modifying the component itself.
//
// Example Usage:
// <my-component>
// <div slot="header">This content will be placed in the header slot</div>
// <div slot="content">This content will be placed in the content slot</div>
// <div slot="footer">This content will be placed in the footer slot</div>
// </my-component>
func Slot(value string) html.Node {
	return html.NewAttribute("slot", value)
}

// Span constructs an html.Node for the `span` attribute.
//
// The `span` attribute is used to group inline elements and apply styles or
// functionalities to them as a unit. It does not create any visual or
// structural impact on the HTML document. It is often used to target specific
// portions of text within larger elements, such as paragraphs or headings, for
// styling purposes. The `span` element does not cause line breaks and can be
// nested within other elements.
//
// Example Usage:
// <p>This is a paragraph with a <span style="color: blue;">blue</span> word.</p>
// <p>This is a <span class="highlight">highlighted</span> text within a paragraph.</p>
func Span(value string) html.Node {
	return html.NewAttribute("span", value)
}

// SpellCheck constructs an html.Node for the `spellcheck` attribute.
//
// `spellcheck` is an attribute used to control the automatic spell checking
// behavior of an HTML element. When this attribute is present, it informs the
// browser whether the element's text content should be checked for spelling
// errors or not. The value of the `spellcheck` attribute can be either "true"
// or "false".
//
// Example Usage:
// <input type="text" spellcheck="false" value="I have intentionally misspelled words.">
// <textarea spellcheck="true" placeholder="Type here..."></textarea>
func SpellCheck(value string) html.Node {
	return html.NewAttribute("spellcheck", value)
}

// Src constructs an html.Node for the `src` attribute.
//
// `src` is used to specify the source URL or file path of an external resource
// that needs to be embedded or displayed within an HTML document. It is
// primarily used in tags such as `<img>`, `<audio>`, and `<video>` to specify
// the source of images, audio files, or video files respectively. The `src`
// attribute is essential for rendering these media elements correctly.
//
// Example Usage:
// <img src="image.jpg" alt="A beautiful image">
// <video src="video.mp4" controls>
// Your browser does not support the video tag.
// </video>
func Src(value string) html.Node {
	return html.NewAttribute("src", value)
}

// SrcDoc constructs an html.Node for the `srcdoc` attribute.
//
// The `srcdoc` attribute is used to embed HTML content directly within an HTML
// document. It allows developers to include inline HTML code within an iframe
// element, without the need for a separate external file. This attribute is
// often used when the content to be displayed in the iframe is dynamic or
// generated on the fly. The value of the `srcdoc` attribute is the actual HTML
// markup that will be rendered within the iframe.
//
// Example Usage:
// <iframe srcdoc="<h1>Hello, world!</h1><p>This is some inline HTML content.</p>"></iframe>
func SrcDoc(value string) html.Node {
	return html.NewAttribute("srcdoc", value)
}

// SrcLang constructs an html.Node for the `srclang` attribute.
//
// The `srclang` attribute is used to specify the language of the text within a
// media element in HTML, such as the `track` element in a video or audio
// player. It helps the browser to correctly display and interpret the text in the
// appropriate language, allowing the user to understand the content
// better. The value of the `srclang` attribute should be a valid language code,
// such as "en" for English or "fr" for French.
//
// Example Usage:
// <track src="subtitles_en.vtt" kind="subtitles" srclang="en" label="English subtitles">
// <track src="subtitles_fr.vtt" kind="subtitles" srclang="fr" label="French subtitles">
func SrcLang(value string) html.Node {
	return html.NewAttribute("srclang", value)
}

// SrcSet constructs an html.Node for the `srcset` attribute.
//
// `srcset` is used to specify a list of image sources and their corresponding
// descriptor widths or pixel densities, allowing the browser to choose the most
// appropriate image to display based on the user's device capabilities and
// screen size. This attribute is particularly useful for responsive web design,
// as it ensures that images are optimized for different devices and network
// conditions, improving both the loading speed and the visual quality of the
// website. The value of the `srcset` attribute is a comma-separated list of
// source descriptors, each consisting of a URL followed by a space and the width
// or pixel density descriptor.
//
// Example Usage:
// <img src="image.jpg" srcset="image.jpg 1x, image-2x.jpg 2x, image-3x.jpg 3x">
func SrcSet(value string) html.Node {
	return html.NewAttribute("srcset", value)
}

// Start constructs an html.Node for the `start` attribute.
//
// The `start` attribute is used to specify the starting number of an ordered
// list in HTML. By default, ordered lists start at the number 1, but the
// `start` attribute allows developers to customize the starting number of the
// list. This attribute is especially useful when a list needs to continue from
// a previous list or needs to start at a number other than 1.
//
// Example Usage:
// <ol start="10">
// <li>This is item number 10</li>
// <li>This is item number 11</li>
// <li>This is item number 12</li>
// </ol>
//
// <ol start="50">
// <li>This is item number 50</li>
// <li>This is item number 51</li>
// <li>This is item number 52</li>
// </ol>
func Start(value string) html.Node {
	return html.NewAttribute("start", value)
}

// Step constructs an html.Node for the `step` attribute.
//
// The `step` attribute is used to specify the interval or step size for
// numeric input fields in an HTML form. It defines the amount by which the
// value should increase or decrease when using the arrow controls or keyboard
// input. The `step` value can be a positive or negative number, or even a
// decimal value, allowing for fine-grained control over the input increments.
//
// Example Usage:
// <input type="number" step="1">
// This input field accepts whole numbers only, incrementing or decrementing by 1 each time.
//
// <input type="number" step="0.5">
// This input field accepts decimal numbers, incrementing or decrementing by 0.5 each time.
//
// <input type="number" step="-10">
// This input field accepts negative numbers, decrementing by 10 each time.
func Step(value string) html.Node {
	return html.NewAttribute("step", value)
}

// Style constructs an html.Node for the `style` attribute.
//
// The `style` attribute is used to add inline CSS styles to an HTML element. It allows developers to directly apply specific visual formatting, such as color, font size, or padding, to individual elements within the HTML document. The value of the `style` attribute consists of one or more CSS property-value pairs, separated by semicolons. Each property-value pair defines a specific style rule that will be applied to the element.
//
// Example Usage:
// <p style="color: blue; font-size: 20px;">This paragraph has a blue color and a font size of 20 pixels.</p>
// <p style="background-color: yellow; padding: 10px;">This paragraph has a yellow background color and a padding of 10 pixels.</p>
func Style(value string) html.Node {
	return html.NewAttribute("style", value)
}

// TabIndex constructs an html.Node for the `tabindex` attribute.
//
// `tabindex` is used to specify the order in which elements should be
// navigated when the user interacts with a web page using the keyboard. It
// allows developers to define a custom tab sequence for elements, ensuring
// that keyboard-only users can navigate through the page efficiently. The
// value of the `tabindex` attribute can be a positive integer to specify the
// order in which elements should be focused, or it can be set to "-1" for an
// element that should not be included in the default tab order. If multiple
// elements have the same `tabindex` value, they are navigated in the order they
// appear in the HTML document.
//
// Example Usage:
// <input type="text" tabindex="1">This input field will be focused first when
// tabbing through the page.
// <button tabindex="2">This button will be focused second when tabbing through
// the page.
// <a href="#" tabindex="-1">This link will be skipped when tabbing through the
// page, as it has a `tabindex` value of -1.
func TabIndex(value string) html.Node {
	return html.NewAttribute("tabindex", value)
}

// Target constructs an html.Node for the `target` attribute.
//
// The `target` attribute is used to specify where a linked resource should be
// opened when clicked. It determines the browsing context in which the linked
// resource should be loaded, such as a new window, a new tab, or the same frame
// or window. By default, linked resources are opened in the same browsing
// context, but the `target` attribute can be used to override this behavior. The
// value of the `target` attribute can be set to `_blank` to open the link in a
// new tab or window, `_self` to open the link in the same frame or window, or
// a custom name that can be used as a target for other links.
//
// Example Usage:
// <a href="https://www.example.com" target="_blank">This link opens in a new tab.</a>
// <a href="/about" target="_self">This link opens in the same window.</a>
func Target(value string) html.Node {
	return html.NewAttribute("target", value)
}

// Title constructs an html.Node for the `title` attribute.
//
// The `title` attribute is used to provide a text description or tooltip for
// an HTML element. It is primarily used to offer additional information or
// context about the element when a user hovers over it with their cursor. This
// attribute is commonly used in images, links, and form inputs to provide more
// details about the content or purpose of the element.
//
// Example Usage:
// <img src="image.jpg" alt="An image" title="This is a beautiful landscape photo.">
// <a href="https://www.example.com" title="Visit our website">Click here to visit our website</a>
// <input type="text" placeholder="Enter your name" title="Please enter your full name.">
func Title(value string) html.Node {
	return html.NewAttribute("title", value)
}

// Translate constructs an html.Node for the `translate` attribute.
//
// The `translate` attribute is used to specify whether the content of an HTML
// element should be translated or not. It is primarily used for localization
// purposes, allowing developers to indicate if the text within an element
// should be translated into the user's language. By default, the `translate`
// attribute is set to "yes", meaning the content should be translated.
// However, it can be set to "no" to indicate that the content should not be
// translated.
//
// Example Usage:
// <p translate="yes">This paragraph should be translated.</p>
// <p translate="no">This paragraph should not be translated.</p>
func Translate(value string) html.Node {
	return html.NewAttribute("translate", value)
}

// Type constructs an html.Node for the `type` attribute.
//
// `type` is used to specify the type or format of data entered or displayed in
// an HTML input element. It determines how the browser interprets and handles
// the input, allowing for validation and control over user input. The value of
// the `type` attribute can be a variety of options, including text, number,
// email, password, etc., each indicating the expected input format.
//
// Example Usage:
// <input type="text" placeholder="Enter your name">
// <input type="number" min="1" max="100">
func Type(value string) html.Node {
	return html.NewAttribute("type", value)
}

// UseMap constructs an html.Node for the `usemap` attribute.
//
// `usemap` is used to associate an image with a client-side image map in an
// HTML document. It allows developers to define clickable areas or hotspots on
// an image, each linked to a specific URL or JavaScript function. The `usemap`
// attribute's value should be set to the ID of the corresponding `map`
// element, which specifies the shape and coordinates of the image map.
//
// Example Usage:
// <img src="image.png" usemap="#map">
//
// <map id="map" name="map">
// <area shape="circle" coords="50,50,30" href="https://www.example.com">
// <area shape="rectangle" coords="100,100,200,200" href="https://www.example.com">
// </map>
func UseMap(value string) html.Node {
	return html.NewAttribute("usemap", value)
}

// Value constructs an html.Node for the `value` attribute.
//
// The `value` attribute is used to specify the initial value of an input
// element in an HTML form. It allows users to pre-fill input fields with a
// default value, providing a starting point for the user to edit or submit.
// The value can be text, numbers, or other valid input depending on the type
// of the input element. The `value` attribute can also be dynamically changed
// using JavaScript to update the input field's value based on user
// interactions or other events.
//
// Example Usage:
// <input type="text" value="John Doe">
// <input type="number" value="25">
// <textarea rows="4" cols="50">Default text in the textarea.</textarea>
// <input type="checkbox" value="apple" checked> Apple
// <input type="radio" name="fruit" value="apple" checked> Apple
// <select>
// <option value="volvo">Volvo</option>
// <option value="saab" selected>Saab</option>
// <option value="bmw">BMW</option>
// </select>
func Value(value string) html.Node {
	return html.NewAttribute("value", value)
}

// Width constructs an html.Node for the `width` attribute.
//
// The `width` attribute is used to specify the width of an HTML element. It allows
// developers to control the size of elements, such as images, tables, or
// containers, on a web page. The value of the `width` attribute can be specified
// in pixels, percentage, or other units of measurement. It determines the
// amount of horizontal space that the element occupies within its parent
// container.
//
// Example Usage:
// <img src="image.jpg" alt="Example Image" width="200">
// <table width="100%">
// <tr>
// <td>Content 1</td>
// <td>Content 2</td>
// </tr>
// </table>
// <div style="width: 50%">This div has a width of 50% of its parent container.</div>
func Width(value string) html.Node {
	return html.NewAttribute("width", value)
}

// Wrap constructs an html.Node for the `wrap` attribute.
//
// The `wrap` attribute is used to specify how the text within a text area
// should be wrapped when it exceeds the width of the text area. It determines
// whether the text should wrap automatically or if horizontal scrolling should
// be enabled to view the overflowing text.
//
// Example Usage:
// <textarea wrap="hard">This text area has hard wrapping enabled.</textarea>
// <textarea wrap="soft">This text area has soft wrapping enabled.</textarea>
func Wrap(value string) html.Node {
	return html.NewAttribute("wrap", value)
}
