package rsocket

import (
	"github.com/reactivex/rxgo/observable"
)

type Availability interface {
	Availability() float64
}

type Closeable interface {
	OnClose() observable.Observable
}

type Disposable interface {
	Dispose()
	IsDisposed() bool
}

type (
	// setup function
	RSocketSetup func(payload Payload, socket RSocket) RSocket

	// request_response function
	RSocketRequestResponse func(payload Payload) observable.Observable

	// fire and forget function
	RSocketFireAndForget func(payload Payload) observable.Observable

	// request stream function
	RSocketRequestStream func(payload Payload) observable.Observable
	// request channel
	RSocketRequestChannel func(payloads observable.Observable) observable.Observable
)

type RSocket struct {
	RequestResponse RSocketRequestResponse
	FireAndForget   RSocketFireAndForget
	RequestStream   RSocketRequestStream
	RequestChannel  RSocketRequestChannel
}
