package dstp

import (
	"context"
	"fmt"
	"github.com/ycd/dstp/config"
	"github.com/ycd/dstp/pkg/common"
	"github.com/ycd/dstp/pkg/ping"
	"reflect"
)

type Result struct {
	Ping      string
	DNS       string
	SystemDNS string
	TLS       string
	HTTPS     string
}

func (r Result) Output() string {
	var output string

	v := reflect.ValueOf(r)

	for i := 0; i < v.NumField(); i++ {
		output += fmt.Sprintf("%s: %v\n", common.White(v.Type().Field(i).Name), v.Field(i).Interface())
	}

	return output
}

// RunAllTests executes all the tests against the given domain, IP or DNS server.
func RunAllTests(ctx context.Context, config config.Config) error {
	var result Result

	out, err := ping.RunTest(ctx, common.Address(config.Addr))
	if err != nil {
		result.Ping = common.Red(out.String())
	} else {
		result.Ping = common.Green(out.String())
	}

	fmt.Println(result.Output())

	return nil
}
