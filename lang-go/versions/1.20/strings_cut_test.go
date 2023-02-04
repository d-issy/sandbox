package go120_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCutPrefix(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		s, prefix string
		after     string
		found     bool
	}{
		{"abc", "a", "bc", true},
		{"abc", "abc", "", true},
		{"abc", "b", "abc", false},
	} {
		after, found := strings.CutPrefix(tt.s, tt.prefix)
		assert.Equal(t, tt.after, after)
		assert.Equal(t, tt.found, found)
	}
}

func TestCutSuffix(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		s, suffix string
		before    string
		found     bool
	}{
		{"abc", "c", "ab", true},
		{"abc", "abc", "", true},
		{"abc", "b", "abc", false},
	} {
		before, found := strings.CutSuffix(tt.s, tt.suffix)
		assert.Equal(t, tt.before, before)
		assert.Equal(t, tt.found, found)
	}
}
