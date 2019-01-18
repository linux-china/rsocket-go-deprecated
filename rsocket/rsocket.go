package rsocket

func Connect() *RSocketRequester {
	return &RSocketRequester{}
}

func Receive() *RSocketResponder {
	return &RSocketResponder{}
}
