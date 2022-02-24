package bbs

import (
	"log"

	"github.com/PatrickRudolph/telnet"
)

// Reader reads messages via the telnet.Connection and passes them to a channel.
type Reader struct {
	C chan []byte
}

// NewReader constructs a new Reader with a new channel C.
func NewReader() *Reader {
	return &Reader{C: make(chan []byte)}
}

// Read() reads data into buf
// unread bytes will be discard if buf is full
func (r *Reader) Read(in *telnet.Connection) (err error) {
	if r.C == nil {
		r.C = make(chan []byte)
	}
	defer close(r.C)
	for {
		buf := make([]byte, 256)
		in.Read(buf)
		log.Printf("Received byte: %v", buf)

		if n := len(buf); n > 0 {
			r.C <- buf
		}
	}

}
