package dstp

import (
	"context"
	"crypto/tls"
	"fmt"
	"math"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/ycd/dstp/config"
	"github.com/ycd/dstp/pkg/common"
	"github.com/ycd/dstp/pkg/lookup"
	"github.com/ycd/dstp/pkg/ping"
)

// RunAllTests executes all the tests against the given domain, IP or DNS server.
func RunAllTests(ctx context.Context, config config.Config) error {
	var result common.Result

	addr, err := getAddr(config.Addr)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(5)

	go ping.RunTest(ctx, &wg, common.Address(addr), config.PingCount, config.Timeout, &result)

	go ping.RunDNSTest(ctx, &wg, common.Address(addr), config.PingCount, config.Timeout, &result)

	go lookup.Host(ctx, &wg, common.Address(addr), &result)

	go testTLS(ctx, &wg, common.Address(addr), config.Timeout, config.Port, &result)

	go testHTTPS(ctx, &wg, common.Address(addr), config.Timeout, config.Port, &result)
	wg.Wait()

	s := result.Output(config.Output)
	s += "\n"

	printWithColor(s)

	return nil
}

func testTLS(ctx context.Context, wg *sync.WaitGroup, address common.Address, t int, port string, result *common.Result) error {
	var output string
	defer wg.Done()

	p := "443"

	if port != "" {
		p = port
	}

	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: time.Duration(t) * time.Second}, "tcp", fmt.Sprintf("%s:%s", string(address), p), nil)
	if err != nil {
		result.Mu.Lock()
		result.TLS = err.Error()
		result.Mu.Unlock()
		return err
	}
	err = conn.VerifyHostname(string(address))
	if err != nil {
		result.Mu.Lock()
		result.TLS = err.Error()
		result.Mu.Unlock()
		return err
	}

	notAfter := conn.ConnectionState().PeerCertificates[0].NotAfter
	expiresAfter := time.Until(notAfter)
	expiry := math.Round(expiresAfter.Hours() / 24)
	if expiry > 0 {
		output += fmt.Sprintf("certificate is valid for %v more days", expiry)
	} else {
		output += fmt.Sprintf("the certificate expired %v days ago", -expiry)
	}

	result.Mu.Lock()
	result.TLS = output
	result.Mu.Unlock()

	return nil
}

func testHTTPS(ctx context.Context, wg *sync.WaitGroup, address common.Address, t int, port string, result *common.Result) error {
	defer wg.Done()

	url := fmt.Sprintf("https://%s", address.String())
	if port != "" {
		url += fmt.Sprintf(":%s", port)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		result.Mu.Lock()
		result.HTTPS = err.Error()
		result.Mu.Unlock()
		return err
	}

	client := http.Client{
		Timeout: time.Second * time.Duration(t),
	}

	resp, err := client.Do(req)
	if err != nil {
		result.Mu.Lock()
		result.HTTPS = err.Error()
		result.Mu.Unlock()
		return err
	}

	result.Mu.Lock()
	result.HTTPS = fmt.Sprintf("got %s", resp.Status)
	result.Mu.Unlock()

	return nil
}
