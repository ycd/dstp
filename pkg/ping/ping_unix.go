//go:build !windows

package ping

import (
	"github.com/go-ping/ping"
)

func createPinger(addr string) (*ping.Pinger, error) {
	p, err := ping.NewPinger(addr)

	return p, err
}
