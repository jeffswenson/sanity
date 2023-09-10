package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAttribute(t *testing.T) {
	type testCase struct {
		attribute string
		value     string
		result    string
	}
	tests := []testCase{
		{"id", "", `<div id=""></div>`},
		{"class", "foo class", `<div class="foo class"></div>`},
		{"escape<test", "\"", `<div escape&lt;test="&#34;"></div>`},
	}
	for _, tc := range tests {
		node := NewAttribute(tc.attribute, tc.value)
		require.Equal(t, tc.result, NewTag("div", node).String())
	}
}

func TestBoolAttribute(t *testing.T) {
	type testCase struct {
		attribute string
		result    string
	}
	tests := []testCase{
		{"async", `<link async>`},
		{"default", `<link default>`},
		{"escape<test", `<link escape&lt;test>`},
	}
	for _, tc := range tests {
		node := NewBoolAttribute(tc.attribute)
		require.Equal(t, tc.result, NewVoidTag("link", node).String())
	}
}
