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

// duplex connection
type DuplexConnection struct {
	disposed bool
}

func (conn *DuplexConnection) SendOne(frame Frame) observable.Observable {
	return observable.Empty()
}

func (conn *DuplexConnection) Send(frames observable.Observable) observable.Observable {
	return observable.Empty()
}

func (conn *DuplexConnection) Receive() observable.Observable {
	return observable.Empty()
}

func (conn *DuplexConnection) Availability() float64 {
	return 1.0
}

func (conn *DuplexConnection) OnClose() observable.Observable {
	return observable.Empty()
}

func (conn *DuplexConnection) Dispose() {
	conn.disposed = true
}

func (conn *DuplexConnection) IsDisposed() bool {
	return conn.disposed
}

//  RSocket Responder
type RSocketResponder struct {
	uri           string
	acceptor      RSocketAccept
	handler       RSocket
	errorConsumer RSocketErrorConsumer
	connection    DuplexConnection
	RSocket
}

func (server *RSocketResponder) Acceptor(acceptor RSocketAccept) *RSocketResponder {
	server.acceptor = acceptor
	return server
}

func (server *RSocketResponder) Transport(uri string) *RSocketResponder {
	server.uri = uri
	return server
}

func (server *RSocketResponder) ErrorConsumer(errorConsumer RSocketErrorConsumer) *RSocketResponder {
	server.errorConsumer = errorConsumer
	return server
}

func (server *RSocketResponder) Start() *RSocketResponder {
	return server
}

//  Rsocket requester
type RSocketRequester struct {
	uri              string
	metadataMimeType string
	dataMimeType     string
	setupPayload     Payload
	handler          RSocket
	errorConsumer    RSocketErrorConsumer
	connection       DuplexConnection
	RSocket
}

func (client *RSocketRequester) MetadataMimeType(metadataMimeType string) *RSocketRequester {
	client.metadataMimeType = metadataMimeType
	return client
}

func (client *RSocketRequester) DataMimeType(dataMimeType string) *RSocketRequester {
	client.dataMimeType = dataMimeType
	return client
}

func (client *RSocketRequester) SetupPayload(setupPayload Payload) *RSocketRequester {
	client.setupPayload = setupPayload
	return client
}

func (client *RSocketRequester) Acceptor(handler RSocket) *RSocketRequester {
	client.handler = handler
	return client
}

func (client *RSocketRequester) Transport(uri string) *RSocketRequester {
	client.uri = uri
	return client
}

func (client *RSocketRequester) ErrorConsumer(errorConsumer RSocketErrorConsumer) *RSocketRequester {
	client.errorConsumer = errorConsumer
	return client
}

func (client *RSocketRequester) Start() *RSocketRequester {
	return client
}

func (client *RSocketRequester) Dispose() {

}

func (client *RSocketRequester) IsDisposed() bool {
	return client.connection.disposed
}
