package lookup

import (
	"context"
	"net"
	"strings"
	"sync"

	"github.com/ycd/dstp/pkg/common"
)

func Host(ctx context.Context, wg *sync.WaitGroup, addr common.Address, result *common.Result) error {
	defer wg.Done()

	part := common.ResultPart{}

	addrs, err := net.LookupHost(addr.String())
	if err != nil {
		part.Error = err
		result.SystemDNS = part
		return err
	}

	part.Content = "resolving " + strings.Join(addrs, ", ")
	result.SystemDNS = part

	return nil
}
