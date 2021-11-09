package dstp

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ycd/dstp/config"
	"github.com/ycd/dstp/pkg/common"
	"github.com/ycd/dstp/pkg/ping"
	"math"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Result struct {
	Ping  string `json:"ping"`
	DNS   string `json:"dns"`
	TLS   string `json:"tls"`
	HTTPS string `json:"https"`
}

func (r Result) Output(outputType string) string {
	var output string

	switch outputType {
	case "plaintext":
		v := reflect.ValueOf(r)
		for i := 0; i < v.NumField(); i++ {
			output += fmt.Sprintf("%s: %v\n", common.White(v.Type().Field(i).Name), common.Green(v.Field(i).Interface()))
		}
	case "json":
		// SAFETY: we are sure that this never fails
		byt, _ := json.MarshalIndent(r, "", "  ")
		output += string(byt)
	}

	return output
}

// RunAllTests executes all the tests against the given domain, IP or DNS server.
func RunAllTests(ctx context.Context, config config.Config) error {
	var result Result

	addr := config.Addr

	for _, prefix := range []string{"https://", "http://"} {
		if strings.HasPrefix(addr, prefix) {
			addr = strings.ReplaceAll(addr, prefix, "")
		}
	}

	if out, err := ping.RunTest(ctx, common.Address(addr)); err != nil {
		result.Ping = err.Error()
	} else {
		result.Ping = out.String()
	}

	if out, err := ping.RunDNSTest(ctx, common.Address(addr)); err != nil {
		result.DNS = err.Error()
	} else {
		result.DNS = out.String()
	}

	if out, err := testTLS(ctx, common.Address(addr)); err != nil {
		result.TLS = err.Error()
	} else {
		result.TLS = out.String()
	}

	if out, err := testHTTPS(ctx, common.Address(addr)); err != nil {
		result.HTTPS = err.Error()
	} else {
		result.HTTPS = out.String()
	}

	fmt.Println(result.Output(config.Output))

	return nil
}

func testTLS(ctx context.Context, address common.Address) (common.Output, error) {
	var output string

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", string(address)), nil)
	if err != nil {
		return "", err
	}
	err = conn.VerifyHostname(string(address))
	if err != nil {
		return "", err
	}

	notAfter := conn.ConnectionState().PeerCertificates[0].NotAfter
	expiresAfter := time.Until(notAfter)
	expiry := math.Round(expiresAfter.Hours() / 24)
	output += fmt.Sprintf("certificate is valid for %v more days", expiry)

	return common.Output(output), nil
}

func testHTTPS(ctx context.Context, address common.Address) (common.Output, error) {
	var output string

	resp, err := http.Get(fmt.Sprintf("https://%s", address))
	if err != nil {
		return "", err
	}

	output += fmt.Sprintf("got %s", resp.Status)
	return common.Output(output), nil
}
