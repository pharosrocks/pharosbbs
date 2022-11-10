package bbs

import (
	"net"
	"testing"

	"github.com/PatrickRudolph/telnet"
	"github.com/PatrickRudolph/telnet/options"
)

func TestNewReader(t *testing.T) {
	tests := []struct {
		name string
		want chan []byte
	}{
		// TODO: Add test cases.
		{
			name: "server.NewReader",
			want: make(chan []byte), // an empty channel
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReader(); got.C == tt.want {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_Read(t *testing.T) {
	type fields struct {
		C chan []byte
	}
	type args struct {
		in *telnet.Connection
	}

	mockConn, _ := net.Pipe()
	options := []telnet.Option{options.NAWSOption}
	mockServer := telnet.NewConnection(mockConn, options)
	a := args{}
	a.in = mockServer

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "server.Reader.Read",
			args: a,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = net.Pipe()
			r := &Reader{
				C: tt.fields.C,
			}
			if err := r.Read(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
