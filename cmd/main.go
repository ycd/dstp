package main

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/ycd/dstp/config"
	"github.com/ycd/dstp/pkg/dstp"
	"log"
)

var (
	dstpCmd = &cobra.Command{
		Use:   "dstp",
		Short: "Run bunch of networking tests against your site.",
	}
	Addr     string
	addrFlag string = "addr"
)

func init() {
	dstpCmd.PersistentFlags().StringVarP(&Addr, addrFlag, "a", "", "URL, domain name or a dns server to run tests against")
	dstpCmd.MarkPersistentFlagRequired(addrFlag)
}

func main() {
	ctx := context.Background()
	conf, err := executeCLI()
	if err != nil {
		log.Fatalf(err.Error())
	}

	dstp.RunAllTests(ctx, conf)
}

func executeCLI() (config.Config, error) {
	if err := dstpCmd.Execute(); err != nil {
		return config.Config{}, err
	}

	conf := config.Config{Addr: Addr}

	return conf, nil
}
