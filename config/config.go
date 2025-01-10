package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Config struct {
	Addr            string
	Output          string
	PingCount       int
	Timeout         int
	ShowHelp        bool
	Port            string
	CustomDnsServer string
}

var usageStr = `Usage: dstp [OPTIONS] [ARGS]
Options:
	-a, --addr   <string>  The URL or the IP address to run tests against      [REQUIRED]
	-o, --out    <string>  The type of the output, either json or plaintext    [Default: plaintext] 
	-p           <int>     Number of ping packets                              [Default: 3]
	-t           <int>     Give up on ping after this many seconds             [Default: 2s per ping packet]
	--port       <string>  Port for testing TLS and HTTPS connectivity         [Default: 443]
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
	fs.StringVar(&opts.Port, "port", "", "Port for testing TLS and HTTPS connectivity")
	fs.IntVar(&opts.PingCount, "p", 3, "Number of ping packets")
	fs.IntVar(&opts.Timeout, "t", -1, "Give up on ping after this many seconds")
	fs.StringVar(&opts.CustomDnsServer, "dns", "", "Custom DNS server to use for DNS resolution")
	fs.BoolVar(&opts.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&opts.ShowHelp, "help", false, "Show help message")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	values := fs.Args()

	if opts.ShowHelp {
		HelpAndExit()
	}

	if !opts.ShowHelp && len(values) < 1 && opts.Addr == "" {
		HelpAndExit()
	}

	if opts.Addr == "" {
		if len(values) >= 1 {
			opts.Addr = values[0]
		} else {
			return nil, fmt.Errorf("address cannot be empty")
		}
	}

	return opts, nil
}
