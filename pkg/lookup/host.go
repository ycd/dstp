package lookup

import (
	"context"
	"github.com/ycd/dstp/pkg/common"
	"net"
	"strings"
)

func Host(ctx context.Context, addr common.Address) (common.Output, error) {
	addrs, err := net.LookupHost(addr.String())
	if err != nil {
		return "", err
	}

	output := "resolving " + strings.Join(addrs, ", ")

	return common.Output(output), nil
}
