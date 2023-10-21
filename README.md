# Sanity

*Sanity := Server Side Rendering - Templates*

Sanity is a Go library for server-side HTML rendering. Sanity is inspired by
the component-based architecture of modern UI toolkits like React. Components
are written as normal Go functions that return an `html.Node`. The advantage of
using Go over HTML templating is that Go is more flexible than most HTML
templating languages and has better tooling support for type checking and
auto-completion. It also allows the use of Go control flow instead of using a
mediocure language embedded in the template.

```
go get github.com/jeffswenson/sanity@latest
```

```go
package main

import (
	"fmt"

	"github.com/jeffswenson/sanity/pkg/attr"
	"github.com/jeffswenson/sanity/pkg/html"
	"github.com/jeffswenson/sanity/pkg/tag"
)

var fruit = []string { "apple", "orange", "banana" }

// Calling fruitView(fruit) produces the following html:
//
// <ul class="fruit-list">
//     <li>apple</li>
//     <li>orange</li>
//     <li>banana</li>
// </ul>
func fruitView(fruitList []string) html.Node {
	return tag.UL(
		attr.Class("fruit-list"), 
		html.ForEach(fruitList, func(fruit string) html.Node {
			return tag.LI(html.InnerText(fruit)) 
		}),
	)
}

func main() {
    node := fruitView(fruit)
    fmt.Print(node)
}
```

## API

The Sanity API is broken into three packages.

* `pkg/html`: contains the core implementation and utilities
* `tag`: contains a function for every HTML tag
* `attr`: contains a function for every HTML attribute

The `tag` and `attr` packages are implemented using public functions from
`html`. So it is possible to create tags and attributes that are not part of
the standard by using the functions declared in `html`.

The function header comments in `tag` and `attr` were written by Chat GPT, so
take them with a grain of salt. All other documentation and all code was
written the old fashioned way.

## Testing

Simple components can be tested using the `html.Node.String()` function. For more
complicated structures, use something like the goquery library to parse the
generated output.

## Performance

The [benchmarks](BENCHMARK.md) contained in the sanity package suggest sanity
is more cpu efficient than `text/html` for templates with many dynamic (`{{}}`)
blocks. Using sanity to generate the HTML is also much cheaper than compressing
the HTML with gzip, so sanity views are unlikely to be a significant portion of
an application's resource consumption footprint.
