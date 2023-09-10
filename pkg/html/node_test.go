package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeDefault(t *testing.T) {
	var n Node	
	require.Empty(t, n.Render())
	require.Empty(t, n.String())
}
