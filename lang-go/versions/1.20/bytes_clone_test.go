package go120_test

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestBytesClone(t *testing.T) {
	src := []byte("hello1")
	dst := bytes.Clone(src)
	assert.Equal(t, src, dst)
	assert.NotEqual(t, unsafe.Pointer(&src), unsafe.Pointer(&dst))

	src = []byte("hello2")
	assert.NotEqual(t, src, dst)

}
