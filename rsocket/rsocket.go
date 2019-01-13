package rsocket

func Connect() *RSocketClient {
	return &RSocketClient{}
}

func Receive() *RSocketServer {
	return &RSocketServer{}
}

func Handle() *RSocketServer {
	return &RSocketServer{}
}
