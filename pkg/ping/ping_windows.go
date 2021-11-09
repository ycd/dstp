//go:build windows
// +build windows

package ping

import (
	"github.com/go-ping/ping"
)

func createPinger(addr string) (*ping.Pinger, error) {
	p, err := ping.NewPinger(addr)

	// https://pkg.go.dev/github.com/go-ping/ping#readme-windows
	p.SetPrivileged(true)

	return p, err
}
