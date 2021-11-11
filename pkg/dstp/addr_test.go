package dstp

import (
	"testing"
)

func TestGetAddr(t *testing.T) {

	tests := []struct {
		addrs []string
	}{
		{
			addrs: []string{
				"facebook.com",
				"facebook.com:80",
				"https://jvns.ca",
				"https://jvns.ca:443",
				"8.8.8.8",
				"2606:4700:3031::ac43:b35a",
				"meta.stackoverflow.com:443",
				"https://meta.stackoverflow.com/",
			},
		},
	}
	for _, tt := range tests {
		for _, addr := range tt.addrs {
			a, _ := getAddr(addr)
			if a == "" {
				t.Fatalf("address parsing failed for: %v", addr)
			}
		}
	}
}
