package ping

import (
	"context"
	"fmt"
	"github.com/go-ping/ping"
	"github.com/ycd/dstp/pkg/common"
	"strings"
)

func RunTest(ctx context.Context, addr common.Address) (common.Output, error) {
	var output string

	pinger, err := ping.NewPinger(addr.String())
	if err != nil {
		return "", err
	}

	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run ping: %v", err.Error())
	}

	stats := pinger.Statistics()
	output += joinS(common.Green(joinC(stats.AvgRtt.String())), "packet loss:", common.Red("%"+fmt.Sprintf("%v", stats.PacketLoss)))

	return common.Output(output), nil
}

func joinC(args ...string) string {
	return strings.Join(args, ",")
}

func joinS(args ...string) string {
	return strings.Join(args, " ")
}
