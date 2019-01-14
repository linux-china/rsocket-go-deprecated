package rsocket

import (
	"github.com/reactivex/rxgo/observable"
)

type RSocketServer struct {
	uri     string
	handler RSocket
	RSocket
}

func (server *RSocketServer) Acceptor(handler RSocket) *RSocketServer {
	server.handler = handler
	return server
}

func (server *RSocketServer) Transport(uri string) *RSocketServer {
	server.uri = uri
	return server
}

func (server *RSocketServer) Start() *RSocketServer {
	return server
}

func (server *RSocketServer) onClose() observable.Observable {
	return observable.Just(server)
}

type RSocketClient struct {
	uri              string
	metadataMimeType string
	dataMimeType     string
	setupPayload     Payload
	handler          RSocket
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

func (client *RSocketClient) Start() observable.Observable {
	return observable.Just(client)
}

func (client *RSocketClient) Dispose() {

}
