package test

import (
	"bytes"
	"sync/atomic"
	"testing"
	"time"
)

var counter uint32 = 0

func TestBytesBuffer(t *testing.T) {
	var content = bytes.NewBufferString("demo")
	print(content.String())
}

func TestAtomicCounter(t *testing.T) {
	for i := 0; i < 5000; i++ {
		go func() {
			atomic.AddUint32(&counter, 1)
		}()
	}
	time.Sleep(time.Second)
	print(counter)
}
