package uid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	ids := map[string]bool{}
	for i := 0; i < 1000; i++ {
		id := New("")
		require.True(t, len(id) > 0)
		_, ok := ids[id]
		require.False(t, ok, "generated id must be unique")
		ids[id] = true
	}
}

func TestNewWithPrefix(t *testing.T) {
	prefixes := []Prefix{
		NodePrefix,
	}

	ids := map[string]bool{}
	for _, prefix := range prefixes {
		for i := 0; i < 1000; i++ {
			id := New(string(prefix))
			require.True(t, len(id) > 0)
			_, ok := ids[id]
			require.False(t, ok, "generated id must be unique")
			ids[id] = true
		}
	}
}
