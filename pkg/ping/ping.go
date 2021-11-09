package ping

import (
	"context"
	"fmt"
	"github.com/ycd/dstp/pkg/common"
	"strings"
)

func RunTest(ctx context.Context, addr common.Address) (common.Output, error) {
	var output string

	pinger, err := createPinger(addr.String())
	if err != nil {
		return "", err
	}

	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run ping: %v", err.Error())
	}

	stats := pinger.Statistics()
	output += joinS(joinC(stats.AvgRtt.String()))

	return common.Output(output), nil
}

func joinC(args ...string) string {
	return strings.Join(args, ",")
}

func joinS(args ...string) string {
	return strings.Join(args, " ")
}

func RunDNSTest(ctx context.Context, addr common.Address) (common.Output, error) {
	var output string

	pinger, err := createPinger(addr.String())
	if err != nil {
		return "", err
	}

	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run ping: %v", err.Error())
	}

	output += joinS("resolving", pinger.IPAddr().String())
	return common.Output(output), nil
}
