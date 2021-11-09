package config

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

type Config struct {
	Addr       string
	Output     string
	ShowHelp   bool
	Concurrent bool
}

var usageStr = `
Usage: dstp [OPTIONS] [ARGS]
Options:
	-a, --addr   <string>  The URL or the IP address to run tests against      [REQUIRED]
	-o, --out    <string>  The type of the output, either json or plaintext    [Default: plaintext] 
	-c           <bool>    Run all the tests concurrently.                     [Default: false]
	-h, --help             Show this message and exit.
`

// UsageAndExit prints usage and exists the program.
func UsageAndExit(err error) {
	color.Red(err.Error())
	fmt.Printf(usageStr)
	os.Exit(1)
}

// HelpAndExit , prints helps and exists the program.
func HelpAndExit() {
	fmt.Printf(usageStr)
	os.Exit(0)
}

// ConfigureOptions is a helper function for parsing options
func ConfigureOptions(fs *flag.FlagSet, args []string) (*Config, error) {
	opts := &Config{}

	// Define flags
	fs.StringVar(&opts.Addr, "a", "", "The URL or the IP address to run tests against")
	fs.StringVar(&opts.Addr, "addr", "", "The URL or the IP address to run tests against")
	fs.StringVar(&opts.Output, "o", "plaintext", "The type of the output")
	fs.StringVar(&opts.Output, "out", "plaintext", "The type of the output")
	fs.BoolVar(&opts.Concurrent, "c", false, "Run all the tests concurrently")
	fs.BoolVar(&opts.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&opts.ShowHelp, "help", false, "Show help message")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	values := fs.Args()

	if !opts.ShowHelp && len(values) < 1 {
		return nil, fmt.Errorf("address is required")
	}

	if len(values) > 1 {
		return nil, fmt.Errorf("expected one argument, got %v", len(values))
	}

	if opts.Addr == "" && len(values) >= 1 {
		opts.Addr = values[0]
	} else {
		return nil, fmt.Errorf("address cannot be empty")
	}

	return opts, nil
}
