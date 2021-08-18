package common

import (
	"testing"
)

func TestGetIPV4(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ipv4",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetIPV4()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIPV4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
