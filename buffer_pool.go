package quic

import (
	"fmt"
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"sync"
)

var bufferPool sync.Pool

func getPacketBuffer() *[]byte {
	fmt.Printf("---- getting package buf -----\n")
	return bufferPool.Get().(*[]byte)
}

func putPacketBuffer(buf *[]byte) {
	if cap(*buf) != int(protocol.MaxReceivePacketSize) {
		panic("putPacketBuffer called with packet of wrong size!")
	}
	bufferPool.Put(buf)
}

func init() {
	bufferPool.New = func() interface{} {
		b := make([]byte, 0, protocol.MaxReceivePacketSize)
		return &b
	}
}
