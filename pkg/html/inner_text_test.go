package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInnerText(t *testing.T) {
	type testCase struct {
		argument string
		result   string
	}
	tests := []testCase{
		{"", ""},
		{"contains no html", "contains no html"},
		{"<div></div>", "&lt;div&gt;&lt;/div&gt;"},
	}
	for _, tc := range tests {
		node := InnerText(tc.argument)
		require.Equal(t, tc.result, node.String())
	}
}

func TestRawInnerText(t *testing.T) {
	type testCase struct {
		argument string
		result   string
	}
	tests := []testCase{
		{"", ""},
		{"contains no html", "contains no html"},
		{"<div></div>", "<div></div>"},
	}
	for _, tc := range tests {
		node := RawInnerText(tc.argument)
		require.Equal(t, tc.result, node.String())
	}
}
