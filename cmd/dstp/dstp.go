package main

import (
	"context"

	flag "github.com/spf13/pflag"

	"fmt"
	"os"
	"path/filepath"

	"github.com/ycd/dstp/config"
	"github.com/ycd/dstp/pkg/dstp"
)

func main() {
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	// Configure the options from the flags/config file
	opts, err := config.ConfigureOptions(fs, os.Args[1:])
	if err != nil {
		config.UsageAndExit(err)
	}

	ctx := context.Background()

	err = dstp.RunAllTests(ctx, *opts)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
