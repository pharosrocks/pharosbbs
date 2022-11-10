package bbs

import (
	"bytes"
	"net"
	"testing"

	"github.com/PatrickRudolph/telnet"
)

func TestTelnetInit(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	tests := []struct {
		name     string
		expected []byte
	}{
		// TODO: Add test cases.
		{
			name: "nonblocking",
			expected: []byte{
				telnet.IAC, telnet.WILL, telnet.TeloptTTYPE,
				// telnet.IAC, telnet.SB, telnet.TeloptTTYPE, telnet.TelQualSEND, telnet.IAC, telnet.SE,
				telnet.IAC, telnet.DO, telnet.TeloptNAWS,
				telnet.IAC, telnet.WILL, telnet.TeloptECHO,
				telnet.IAC, telnet.DONT, telnet.TeloptECHO,
				telnet.IAC, telnet.WILL, telnet.TeloptSGA,
				telnet.IAC, telnet.WILL, telnet.TeloptBINARY,
				telnet.IAC, telnet.DONT, telnet.TeloptBINARY,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := net.Pipe()

			buf.Reset()
			go func() {
				s := &telnet.Connection{Conn: server}
				s.Close()
			}()

			buf.ReadFrom(client)
			client.Close()

			if !bytes.Equal(tt.expected, buf.Bytes()) {
				t.Errorf("Expected %v, got %v", tt.expected, buf.Bytes())
			}
		})
	}
}
