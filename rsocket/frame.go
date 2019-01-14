package rsocket

type Frame struct {
	streamId  int32
	FrameType byte
	Flags     int16
	Payload
}

func (frame *Frame) HasMetadata() bool {
	return true
}

func NewSetupFrame(metadataMimeType string, dataMimeType string, payload Payload) Frame {
	return Frame{}
}

func NewRequestFrame(streamId int32, payload Payload, initialRequestN int32) Frame {
	return Frame{}
}
