package rsocket

func Connect() *RSocketRequester {
	return &RSocketRequester{}
}

func Receive() *RSocketResponder {
	return &RSocketResponder{}
}

func Handle() *RSocketServer {
	return &RSocketServer{}
}
