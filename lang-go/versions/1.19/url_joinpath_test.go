package go119_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var urlJoinPathTests = []struct {
	base string
	elem []string
	want string
}{
	{
		base: "https://github.com/d-issy/",
		elem: []string{"hello"},
		want: "https://github.com/d-issy/hello",
	},
	{
		base: "https://github.com/d-issy/a/b/c",
		elem: []string{"../../../hello"},
		want: "https://github.com/d-issy/hello",
	},
	{
		base: "https://github.com/d-issy/a/b/c",
		elem: []string{"..", "..", "..", "hello"},
		want: "https://github.com/d-issy/hello",
	},
}

func TestUrlJoinpath(t *testing.T) {
	for _, tt := range urlJoinPathTests {
		result, err := url.JoinPath(tt.base, tt.elem...)
		assert.NoError(t, err)
		assert.Equal(t, result, tt.want)
	}
}
