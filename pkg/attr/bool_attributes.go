package attr

import "github.com/jeffswenson/sanity/pkg/html"

// AllowFullScreen constructs an html.Node for the `allowfullscreen` attribute.
//
// `allowfullscreen` is used to enable or disable fullscreen mode for media
// elements like videos or iframes. When the `allowfullscreen` attribute is
// present and set to "true" or "1", the user is able to enter fullscreen mode
// by clicking on the media element or using the browser's fullscreen controls.
// If the attribute is not present or set to "false" or "0", fullscreen mode is
// disabled and the element will not be able to occupy the entire screen.
//
// Example Usage:
// <iframe src="https://www.youtube.com/embed/abcdef123" allowfullscreen></iframe>
func AllowFullScreen() html.Node {
	return html.NewBoolAttribute("allowfullscreen")
}

// Async constructs an html.Node for the `async` attribute.
//
// The `async` attribute is used to specify that an external script should be
// downloaded and executed asynchronously. This means that the fetching of the
// script can happen in parallel with the rendering of the HTML document, without
// blocking the rendering process. Once the script has finished downloading, it
// will be executed immediately, regardless of whether the HTML document has
// finished parsing. This attribute is commonly used for non-blocking scripts that
// do not depend on the order of execution or the current state of the HTML
// document.
//
// Example Usage:
// <script src="script.js" async></script>
func Async() html.Node {
	return html.NewBoolAttribute("async")
}

// AutoFocus constructs an html.Node for the `autofocus` attribute.
//
// `autofocus` is an HTML attribute used to automatically focus on a specific
// element when a web page loads. This attribute is commonly used on input
// fields, such as text boxes or search bars, to make it more convenient for
// users to start typing immediately without having to manually click on the
// input field. Only one element on a web page should have the `autofocus`
// attribute, and it is typically used on the most important or commonly used
// input field to enhance user experience.
//
// Example Usage:
// <input type="text" autofocus>
// <button autofocus>Click me!</button>
func AutoFocus() html.Node {
	return html.NewBoolAttribute("autofocus")
}

// AutoPlay constructs an html.Node for the `autoplay` attribute.
//
// The `autoplay` attribute is used to specify that a media element (such as an
// audio or video) should start playing automatically when the page loads. This
// attribute is primarily used for creating a seamless and uninterrupted user
// experience, where the media content starts playing without requiring any
// user interaction.
//
// Example Usage:
// <video src="video.mp4" autoplay></video>
// <audio src="audio.mp3" autoplay></audio>
func AutoPlay() html.Node {
	return html.NewBoolAttribute("autoplay")
}

// Checked constructs an html.Node for the `checked` attribute.
//
// `checked` is used to specify that an input element should be pre-selected or
// pre-checked when the HTML document loads. It is primarily used with radio
// buttons and checkboxes to indicate that a specific option or choice is selected
// by default. When the `checked` attribute is present, the associated input
// element will be displayed with its checked state applied.
//
// Example Usage:
// <input type="radio" name="gender" value="male" checked> Male
// <input type="radio" name="gender" value="female"> Female
// <input type="checkbox" name="agree" value="yes" checked> I agree
//
// In the above example, the "Male" radio button and "I agree" checkbox will be
// selected by default when the page loads.
func Checked() html.Node {
	return html.NewBoolAttribute("checked")
}

// Controls constructs an html.Node for the `controls` attribute.
//
// The `controls` attribute is used to add audio or video controls to an HTML
// element. When applied to an `<audio>` or `<video>` tag, it displays a set of
// play, pause, and volume control buttons for the media content. This attribute
// enhances user experience by allowing them to easily interact with the media and
// control its playback.
//
// Example Usage:
// <video src="video.mp4" controls></video>
// <audio src="audio.mp3" controls></audio>
func Controls() html.Node {
	return html.NewBoolAttribute("controls")
}

// TODO this doc comment looks wrong
// Default constructs an html.Node for the `default` attribute.
//
// The `default` attribute is used to specify a predefined value or state for
// an input element in an HTML form. It allows developers to set a default
// value that will be displayed when the form is initially loaded or reset.
// This can be useful for providing suggested input or pre-filling form fields
// with commonly used values.
//
// Example Usage:
// <input type="text" name="username" value="JohnDoe" default="JohnDoe">
// <input type="checkbox" name="rememberMe" default="checked">
// <textarea name="message" default="Please enter your message here."></textarea>
func Default() html.Node {
	return html.NewBoolAttribute("default")
}

// Defer constructs an html.Node for the `defer` attribute.
//
// The `defer` attribute is used to indicate that a script should be executed
// after the HTML document has been parsed. When the browser encounters a
// script tag with the `defer` attribute, it will continue parsing the rest of
// the HTML document and then execute the script once the document has finished
// loading. This attribute is useful for improving page load speed, as it
// allows scripts to be loaded asynchronously without blocking the parsing and
// rendering of the HTML.
//
// Example Usage:
// <script defer src="script.js"></script>
func Defer() html.Node {
	return html.NewBoolAttribute("defer")
}

// Disabled constructs an html.Node for the `disabled` attribute.
//
// The `disabled` attribute is used to disable an HTML element, preventing user
// interaction or input. It is commonly used for form elements such as buttons,
// input fields, and checkboxes to indicate that they cannot be interacted with
// or modified. When an element is disabled, it appears greyed out and does not
// respond to user actions. This attribute is particularly useful for
// preventing users from submitting incomplete or incorrect data in forms.
//
// Example Usage:
// <button disabled>Submit</button>
// <input type="text" disabled>
// <input type="checkbox" disabled>
func Disabled() html.Node {
	return html.NewBoolAttribute("disabled")
}

// FormNoValidate constructs an html.Node for the `formnovalidate` attribute.
//
// The `formnovalidate` attribute is used to override the default form
// validation of an HTML form. When applied to a submit button or an input
// element with a type of "submit", it allows the form to be submitted without
// performing any client-side validation. This is particularly useful when you
// want to bypass validation for a specific button or input field.
//
// Example Usage:
// <input type="submit" value="Submit" formnovalidate>
func FormNoValidate() html.Node {
	return html.NewBoolAttribute("formnovalidate")
}

// Hidden constructs an html.Node for the `hidden` attribute.
//
// The `hidden` attribute is used to hide an HTML element from display on a
// webpage. When an element has the `hidden` attribute, it will not be visible
// on the page and will not take up any space in the layout. This attribute is
// commonly used to temporarily hide or reveal elements based on specific
// conditions or user interactions. It can be added to any HTML element, such
// as divs, buttons, or images.
//
// Example Usage:
// <div hidden>This element is hidden from display.</div>
// <button hidden>This button is not visible.</button>
func Hidden() html.Node {
	return html.NewBoolAttribute("hidden")
}

// Inert constructs an html.Node for the `inert` attribute.
//
// `inert` is used to indicate that an HTML element and its descendants should
// be non-interactive and inactive. This means that the element and its child
// elements will not respond to user interactions such as clicks or keyboard
// input. It allows developers to temporarily disable or "freeze" certain
// sections of a webpage, preventing any changes or actions from occurring
// within those elements.
//
// Example Usage:
// <div inert>This div and its contents are non-interactive.</div>
func Inert() html.Node {
	return html.NewBoolAttribute("inert")
}

// IsMap constructs an html.Node for the `ismap` attribute.
//
// The `ismap` attribute is used to specify that an image in an HTML document
// is a server-side image map. This attribute allows users to click on
// different regions of the image and be directed to different URLs based on
// their clicks. When the `ismap` attribute is present, the browser sends the
// coordinates of the clicked location to the server, which then determines the
// appropriate URL to load. This attribute is typically used in conjunction
// with the `<img>` tag.
//
// Example Usage:
// <img src="map.jpg" ismap>
func IsMap() html.Node {
	return html.NewBoolAttribute("ismap")
}

// ItemScope constructs an html.Node for the `itemscope` attribute.
//
// `itemscope` is used to define the scope of an item in the HTML document. It
// is primarily used in conjunction with the `itemtype` and `itemprop`
// attributes to markup structured data using microdata. The `itemscope`
// attribute indicates that the element represents an item, such as a person,
// event, or product, and that it contains properties and values related to
// that item. This attribute helps search engines and other web services
// understand the content and context of the data.
//
// Example Usage:
// <div itemscope itemtype="http://schema.org/Person">
// <span itemprop="name">John Doe</span>
// <span itemprop="jobTitle">Web Developer</span>
// </div>
func ItemScope() html.Node {
	return html.NewBoolAttribute("itemscope")
}

// Loop constructs an html.Node for the `loop` attribute.
//
// The `loop` attribute is used to specify whether an audio or video element
// should start playing again from the beginning once it reaches the end. When
// the `loop` attribute is included with a value of "true" or without any value
// at all, the media will loop indefinitely. However, if the `loop` attribute
// is set to "false", the media will play only once. This attribute is commonly
// used when creating background music or looping video animations.
//
// Example Usage:
// <video src="video.mp4" loop>
// Your browser does not support the video tag.
// </video>
func Loop() html.Node {
	return html.NewBoolAttribute("loop")
}

// Multiple constructs an html.Node for the `multiple` attribute.
//
// The `multiple` attribute is used to indicate that a user can select multiple
// options from a list or dropdown menu. It is commonly used in the form's select
// (element) to allow users to select multiple options simultaneously. When the
// `multiple` attribute is added to a select element, the user can hold down the
// Ctrl or Shift key (Windows) or the Command key (Mac) while clicking or
// dragging to select multiple options.
//
// Example Usage:
// <select multiple>
// <option value="option1">Option 1</option>
// <option value="option2">Option 2</option>
// <option value="option3">Option 3</option>
// </select>
func Multiple() html.Node {
	return html.NewBoolAttribute("multiple")
}

// Muted constructs an html.Node for the `muted` attribute.
//
// The `muted` attribute is used to specify that the audio or video element
// should be muted or without sound. It is often used when you want to start a
// video or audio file without any sound playing initially. This attribute can
// be added to the `<video>` and `<audio>` elements.
//
// Example usage:
// <video src="myVideo.mp4" muted></video>
// <audio src="myAudio.mp3" muted></audio>
func Muted() html.Node {
	return html.NewBoolAttribute("muted")
}

// NoModule constructs an html.Node for the `nomodule` attribute.
//
// `nomodule` is used to specify that a JavaScript module should not be
// executed in an HTML document if the browser supports JavaScript modules. This
// attribute is commonly used as a fallback for older browsers that do not support
// JavaScript modules, allowing alternative code or scripts to be executed instead.
// By including `nomodule` in the script tag with the module attribute, the browser
// will bypass executing the module script if it supports modules, and instead
// execute the fallback script specified within the `nomodule` attribute.
//
// Example Usage:
// <script type="module" src="main.js"></script>
// <script nomodule src="fallback.js"></script>
func NoModule() html.Node {
	return html.NewBoolAttribute("nomodule")
}

// NoValidate constructs an html.Node for the `novalidate` attribute.
//
// `novalidate` is used to disable the default HTML5 form validation in an HTML
// document. When this attribute is added to a form element, it tells the browser
// not to validate the form inputs before submission. This attribute is useful in
// situations where custom validation scripts are being used or when the form data
// is being processed on the server side.
//
// Example Usage:
// <form action="/submit" method="post" novalidate>
// <input type="text" required>
// <input type="submit" value="Submit">
// </form>
func NoValidate() html.Node {
	return html.NewBoolAttribute("novalidate")
}

// Open constructs an html.Node for the `open` attribute.
//
// The `open` attribute is used to specify whether a details element should be
// initially open or closed when the page loads. The details element is used to
// create an interactive widget that can be expanded or collapsed to reveal or
// hide additional content. When the `open` attribute is present, the details
// element is expanded by default. When the `open` attribute is not present,
// the details element is collapsed by default.
//
// Example Usage:
// <details open>
// <summary>Click here to expand the details</summary>
// <div>This content is initially visible because the 'open' attribute is present.</div>
// </details>
//
// <details>
// <summary>Click here to expand the details</summary>
// <div>This content is initially hidden because the 'open' attribute is not present.</div>
// </details>
func Open() html.Node {
	return html.NewBoolAttribute("open")
}

// PlaysInline constructs an html.Node for the `playsinline` attribute.
//
// The `playsinline` attribute is used to specify whether a video element
// should play inline or go fullscreen when played on iOS devices. By default,
// videos on iOS devices go fullscreen when played, covering the entire screen.
// However, by adding the `playsinline` attribute to the video element, the
// video will play inline within the webpage, allowing it to be displayed
// alongside other content. This attribute is particularly useful when
// embedding videos in a responsive design or when multiple videos need to be
// displayed simultaneously.
//
// Example Usage:
// <video src="video.mp4" playsinline></video>
func PlaysInline() html.Node {
	return html.NewBoolAttribute("playsinline")
}

// ReadOnly constructs an html.Node for the `readonly` attribute.
//
// `readonly` is used to specify that an input element is read-only, meaning
// that the user cannot edit its value directly. This attribute is particularly
// useful for displaying data that should not be modified by the user, such as
// displaying a user's username or a static value. The `readonly` attribute can be
// applied to input elements of type text, password, date, and more.
//
// Example Usage:
// <input type="text" value="Readonly value" readonly>
// <input type="password" value="********" readonly>
// <input type="date" value="2021-01-01" readonly>
func ReadOnly() html.Node {
	return html.NewBoolAttribute("readonly")
}

// Required constructs an html.Node for the `required` attribute.
//
// `required` is used to specify that an input field must be filled out before
// submitting a form. It is primarily used in form elements such as text
// fields, checkboxes, and radio buttons to ensure that certain fields are not
// left blank. When the `required` attribute is added to an input field, the
// browser will validate the form and display an error message if the field is
// empty when the form is submitted.
//
// Example Usage:
// <input type="text" required>
// <input type="checkbox" required>
// <input type="radio" required>
func Required() html.Node {
	return html.NewBoolAttribute("required")
}

// Reversed constructs an html.Node for the `reversed` attribute.
//
// The `reversed` attribute is used in an ordered list (`<ol>`) element to
// indicate that the order of the list items should be reversed. This means that
// the first list item will be displayed as the last, the second item as the
// second-to-last, and so on. This attribute is useful when presenting a list in
// a non-standard order, such as counting down from a certain number.
//
// Example Usage:
// <ol reversed>
// <li>Third item</li>
// <li>Second item</li>
// <li>First item</li>
// </ol>
func Reversed() html.Node {
	return html.NewBoolAttribute("reversed")
}

// Selected constructs an html.Node for the `selected` attribute.
//
// The `selected` attribute is used to pre-select an option in a dropdown list or
// select element. When this attribute is present in an option tag, that option
// will be displayed as the default selected option when the page loads or the
// form is reset. It allows developers to pre-fill a form field with a specific
// option, providing a default value that can be easily changed by the user if
// needed. The selected attribute is used in combination with the `<option>`
// element.
//
// Example Usage:
// <select>
// <option value="option1">Option 1</option>
// <option value="option2" selected>Option 2</option>
// <option value="option3">Option 3</option>
// </select>
//
// In this example, "Option 2" will be selected by default when the select list is
// rendered.
func Selected() html.Node {
	return html.NewBoolAttribute("selected")
}
