package test

import (
	"bytes"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	var content = bytes.NewBufferString("demo")
	print(content.String())
}
