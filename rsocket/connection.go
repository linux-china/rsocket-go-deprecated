package rsocket

import "github.com/reactivex/rxgo/observable"

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
	uri string
	RSocket
}

func (client *RSocketClient) Transport(uri string) *RSocketClient {
	client.uri = uri
	return client
}

func (client *RSocketClient) Start() *RSocketClient {
	return client
}

func (client *RSocketClient) onClose() observable.Observable {
	return observable.Just(client)
}
