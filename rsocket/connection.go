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

//  RSocket Server
type RSocketServer struct {
	uri           string
	acceptor      RSocketAccept
	handler       RSocket
	errorConsumer RSocketErrorConsumer
	connection    DuplexConnection
	RSocket
}

func (server *RSocketServer) Acceptor(acceptor RSocketAccept) *RSocketServer {
	server.acceptor = acceptor
	return server
}

func (server *RSocketServer) Transport(uri string) *RSocketServer {
	server.uri = uri
	return server
}

func (server *RSocketServer) ErrorConsumer(errorConsumer RSocketErrorConsumer) *RSocketServer {
	server.errorConsumer = errorConsumer
	return server
}

func (server *RSocketServer) Start() *RSocketServer {
	return server
}

//  Rsocket Client
type RSocketClient struct {
	uri              string
	metadataMimeType string
	dataMimeType     string
	setupPayload     Payload
	handler          RSocket
	errorConsumer    RSocketErrorConsumer
	connection       DuplexConnection
	RSocket
}

func (client *RSocketClient) MetadataMimeType(metadataMimeType string) *RSocketClient {
	client.metadataMimeType = metadataMimeType
	return client
}

func (client *RSocketClient) DataMimeType(dataMimeType string) *RSocketClient {
	client.dataMimeType = dataMimeType
	return client
}

func (client *RSocketClient) SetupPayload(setupPayload Payload) *RSocketClient {
	client.setupPayload = setupPayload
	return client
}

func (client *RSocketClient) Acceptor(handler RSocket) *RSocketClient {
	client.handler = handler
	return client
}

func (client *RSocketClient) Transport(uri string) *RSocketClient {
	client.uri = uri
	return client
}

func (client *RSocketClient) ErrorConsumer(errorConsumer RSocketErrorConsumer) *RSocketClient {
	client.errorConsumer = errorConsumer
	return client
}

func (client *RSocketClient) Start() *RSocketClient {
	return client
}

func (client *RSocketClient) Dispose() {

}
