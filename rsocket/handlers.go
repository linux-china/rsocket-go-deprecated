package rsocket

import (
	"github.com/reactivex/rxgo/observable"
)

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

type RSocket struct {
	RequestResponse RSocketRequestResponse
	FireAndForget   RSocketFireAndForget
	RequestStream   RSocketRequestStream
	RequestChannel  RSocketRequestChannel
}
