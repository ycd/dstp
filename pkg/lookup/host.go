package lookup

import (
	"context"
	"net"
	"strings"
	"sync"

	"github.com/ycd/dstp/pkg/common"
)

func Host(ctx context.Context, wg *sync.WaitGroup, addr common.Address, customDnsServer string, result *common.Result) error {
	defer wg.Done()

	part := common.ResultPart{}

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", customDnsServer)
		},
	}

	addrs, err := r.LookupHost(ctx, addr.String())
	if err != nil {
		part.Error = err
		result.SystemDNS = part
		return err
	}

	part.Content = "resolving " + strings.Join(addrs, ", ")
	result.SystemDNS = part

	return nil
}
