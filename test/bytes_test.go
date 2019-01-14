package test

import (
	"bytes"
	"sync/atomic"
	"testing"
)

var counter uint32 = 0

func TestBytesBuffer(t *testing.T) {
	var content = bytes.NewBufferString("demo")
	print(content.String())
}

func TestAtomicCounter(t *testing.T) {
	atomic.AddUint32(&counter, 1)
	print(counter)
}
