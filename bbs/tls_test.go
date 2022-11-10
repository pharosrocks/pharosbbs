package bbs

import (
	"testing"
)

func TestGetCertificate(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetCertificate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetCertificate()
		})
	}
}
