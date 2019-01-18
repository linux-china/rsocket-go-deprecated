# rsocket-go

[![Join the chat at https://gitter.im/RSocket/reactivesocket-go](https://badges.gitter.im/RSocket/reactivesocket-go.svg)](https://gitter.im/ReactiveSocket/reactivesocket-go)

Golang implementation of [RSocket](http://rsocket.io)


# How to use?


* RSocket Server: 

```
package main

import (
	"github.com/linux-china/rsocket-go/rsocket"
	"github.com/reactivex/rxgo/observable"
)

func main() {
	var handler = rsocket.RSocket{}
	handler.RequestResponse = func(payload rsocket.Payload) observable.Observable {
		return observable.Just(payload)
	};
	rsocket.Receive().Transport("tcp://0.0.0.0:42252").Acceptor(handler).Start()
}


```

* RSocket Client: 

```
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

```


### Convention

* More type alias for functions than interfaces

```

type (
	// accept for setup function
	RSocketAccept func(payload Payload, socket RSocket) RSocket

	// request_response function
	RSocketRequestResponse func(payload Payload) observable.Observable

	// fire and forget function
	RSocketFireAndForget func(payload Payload) observable.Observable

	// request stream function
	RSocketRequestStream func(payload Payload) observable.Observable
	// request channel
	RSocketRequestChannel func(payloads observable.Observable) observable.Observable

	RSocketErrorConsumer func(exception interface{})
)

```


# References

* Reactive Extensions for the Go language: https://github.com/ReactiveX/RxGo
* ReactiveX: http://reactivex.io/