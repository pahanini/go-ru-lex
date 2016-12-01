package rulex

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestHasAnyPrefix(t *testing.T) {
	r, p := HasAnyPrefix("индия", []string{"инд", "исп"})
	require.Equal(t, "инд", p)
	require.True(t, r)

	r, p = HasAnyPrefix("англия", []string{"инд", "исп"})
	require.Equal(t, "", p)
	require.False(t, r)
}
