package main

import (
	"github.com/linux-china/rsocket-go/rsocket"
	"github.com/reactivex/rxgo/observable"
)

func main() {
	var handler = rsocket.RSocket{}

	handler.RequestResponse = func(payload rsocket.Payload) observable.Observable {
		return observable.Just(payload)
	}
	rsocket.Receive().Transport("tcp://0.0.0.0:42252").Acceptor(handler).Start()
}
