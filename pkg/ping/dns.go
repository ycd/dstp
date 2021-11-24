package ping

import (
	"context"
	"fmt"
	"github.com/ycd/dstp/pkg/common"
	"sync"
	"time"
)

func RunDNSTest(ctx context.Context, wg *sync.WaitGroup, addr common.Address, count int, timeout int, result *common.Result) error {
	defer wg.Done()

	pinger, err := createPinger(addr.String())
	if err != nil {
		return err
	}

	pinger.Count = count
	pinger.Timeout = time.Duration(timeout) * time.Second

	err = pinger.Run()
	if err != nil {
		return fmt.Errorf("failed to run ping: %v", err.Error())
	}

	result.Mu.Lock()
	result.DNS = joinS("resolving", pinger.IPAddr().String())
	result.Mu.Unlock()
	return nil
}
