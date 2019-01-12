package test

import (
	"bytes"
	"github.com/linux-china/rsocket-go/rsocket"
	"github.com/reactivex/rxgo/handlers"
	"github.com/reactivex/rxgo/observer"
	"testing"
)

func TestRSocketRequestResponse(t *testing.T) {
	var payload = rsocket.Payload{Metadata: bytes.NewBufferString("ip:192.168.0.2"), Data: bytes.NewBufferString("Hello World!")}
	var client = rsocket.Connect().Transport("tcp://127.0.0.1:42252").Start()
	handler := observer.New(handlers.NextFunc(func(payload interface{}) {
		println(payload)
	}))
	client.RequestResponse(payload).Subscribe(handler)
}
