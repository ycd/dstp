package ping

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ycd/dstp/pkg/common"
)

func RunDNSTest(ctx context.Context, wg *sync.WaitGroup, addr common.Address, count int, timeout int, result *common.Result) error {
	defer wg.Done()

	part := common.ResultPart{}
	pinger, err := createPinger(addr.String())
	if err != nil {
		part.Error = err
		result.Mu.Lock()
		result.DNS = part
		result.Mu.Unlock()
		return err
	}

	pinger.Count = count
	pinger.Timeout = time.Duration(timeout) * time.Second

	err = pinger.Run()
	if err != nil {
		part.Error = fmt.Errorf("failed to run ping: %v", err.Error())
		result.Mu.Lock()
		result.DNS = part
		result.Mu.Unlock()
		return err
	}

	part.Content = "resolving " + pinger.IPAddr().String()
	result.Mu.Lock()
	result.DNS = part
	result.Mu.Unlock()
	return nil
}
