package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTag(t *testing.T) {
	require.Equal(t, "<div></div>", NewTag("div").String())
	require.Equal(t,
		"<p id=\"foo-bar\"><span></span></p>",
		NewTag("p",
			NewAttribute("id", "foo-bar"),
			NewTag("span"),
		).String(),
	)
}

func TestNewVoidTag(t *testing.T) {
	require.Equal(t, "<col>", NewVoidTag("col").String())
	require.Equal(t, "<col style=\"width: 50%\">", NewVoidTag("col",
		NewAttribute("style", "width: 50%"),
		// Children of void tags are ignored.
		NewTag("ignored"),
	).String())
}
