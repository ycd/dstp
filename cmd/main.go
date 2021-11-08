package main

import (
	"context"
	"flag"
	"github.com/ycd/dstp/config"
	"github.com/ycd/dstp/pkg/dstp"
	"os"
	"path/filepath"
)

func main() {
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	// Configure the options from the flags/config file
	opts, err := config.ConfigureOptions(fs, os.Args[1:])
	if err != nil {
		config.UsageAndExit(err)
	}

	if opts.ShowHelp {
		config.HelpAndExit()
	}

	ctx := context.Background()

	dstp.RunAllTests(ctx, *opts)
}
