package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForEachNil(t *testing.T) {
	node := ForEach(nil, func(model string) Node {
		panic("is never called")
	})
	require.Empty(t, node.String())
}

func TestForEachEmpty(t *testing.T) {
	node := ForEach([]string{}, func(model string) Node {
		panic("is never called")
	})
	require.Empty(t, node.String())
}

func TestForEachOneTag(t *testing.T) {
	node := ForEach([]string{
		"orange",
	}, func(model string) Node {
		return NewTag("li", InnerText(model))
	})
	require.Equal(t, node.String(), "<li>orange</li>")
	require.Equal(t, NewTag("ul", node).String(), "<ul><li>orange</li></ul>")
}

func TestForEachManyTag(t *testing.T) {
	node := ForEach([]string{
		"orange",
		"bannana",
		"apple",
	}, func(model string) Node {
		return NewTag("li", InnerText(model))
	})
	require.Equal(t, node.String(), "<li>orange</li><li>bannana</li><li>apple</li>")
	require.Equal(t, NewTag("ul", node).String(), "<ul><li>orange</li><li>bannana</li><li>apple</li></ul>")
}

func TestCombineNone(t *testing.T) {
	require.Empty(t, Combine().String())
}

func TestCombineMixed(t *testing.T) {
	node := Combine(
		NewAttribute("id", "id-value"),
		NewTag("div", InnerText("div content")),
		NewAttribute("class", "class-value"),
		NewTag("button", InnerText("button content")),
	)
	require.Equal(t, node.String(), "<div>div content</div><button>button content</button>")
	require.Equal(t, NewTag("span", node).String(),
		"<span id=\"id-value\" class=\"class-value\"><div>div content</div><button>button content</button></span>")
}
