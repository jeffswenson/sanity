package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDocument(t *testing.T) {
	type testCase struct {
		node  Node
		result string
	}
	tests := []testCase {
		{
			Document(),
			"<!DOCTYPE html><html></html>",
		},
		{
			Document(NewAttribute("lang", "en"), InnerText("Hello World!")), 
			`<!DOCTYPE html><html lang="en">Hello World!</html>`,
		},
	}
	for _, tc := range tests {
		require.Equal(t, tc.result, tc.node.String())
	}

}
