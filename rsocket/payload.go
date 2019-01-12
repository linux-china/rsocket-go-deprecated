package rsocket

import "bytes"

type Payload struct {
	Metadata *bytes.Buffer
	Data     *bytes.Buffer
}
