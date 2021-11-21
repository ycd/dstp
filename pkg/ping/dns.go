package ping

import (
	"context"
	"fmt"
	"github.com/ycd/dstp/pkg/common"
	"time"
)

func RunDNSTest(ctx context.Context, addr common.Address, count int, timeout int) (common.Output, error) {
	var output string

	pinger, err := createPinger(addr.String())
	if err != nil {
		return "", err
	}

	pinger.Count = count
	if timeout == -1 {
		pinger.Timeout = time.Duration(2*count) * time.Second
	} else {
		pinger.Timeout = time.Duration(timeout) * time.Second
	}
	err = pinger.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run ping: %v", err.Error())
	}

	output += joinS("resolving", pinger.IPAddr().String())
	return common.Output(output), nil
}
