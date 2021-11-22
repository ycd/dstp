package lookup

import (
	"context"
	"github.com/ycd/dstp/pkg/common"
	"net"
	"strings"
	"sync"
)

func Host(ctx context.Context, wg *sync.WaitGroup, addr common.Address, result *common.Result) error {
	defer wg.Done()

	addrs, err := net.LookupHost(addr.String())
	if err != nil {
		return err
	}

	result.SystemDNS = "resolving " + strings.Join(addrs, ", ")

	return nil
}
