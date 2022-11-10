package bbs

import (
	"testing"
)

func TestServer_ListenAndServe(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ListenAndServe",
			args: args{
				addr: ":8888",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			if err := s.ListenAndServe(tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("Server.ListenAndServe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
