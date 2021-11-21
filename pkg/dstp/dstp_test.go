//go:build integration

package dstp

import (
	"context"
	"github.com/ycd/dstp/config"
	"testing"
)

func TestRunAllTests(t *testing.T) {
	ctx := context.Background()

	c := config.Config{
		Addr:       "https://jvns.ca",
		Output:     "plaintext",
		ShowHelp:   false,
		Concurrent: false,
		Timeout:    3,
		PingCount:  3,
	}

	c1 := config.Config{
		Addr:       "8.8.8.8",
		Output:     "plaintext",
		ShowHelp:   false,
		Concurrent: false,
		Timeout:    3,
		PingCount:  3,
	}

	c2 := config.Config{
		Addr:       "facebook.com",
		Output:     "plaintext",
		ShowHelp:   false,
		Concurrent: false,
		Timeout:    3,
		PingCount:  3,
	}

	c3 := config.Config{
		Addr:       "https://meta.stackoverflow.com/",
		Output:     "plaintext",
		ShowHelp:   false,
		Concurrent: false,
		Timeout:    3,
		PingCount:  3,
	}

	c4 := config.Config{
		Addr:       "facebook.com:80",
		Output:     "plaintext",
		ShowHelp:   false,
		Concurrent: false,
		Timeout:    3,
		PingCount:  3,
	}

	for _, conf := range []config.Config{c, c1, c2, c3, c4} {
		if err := RunAllTests(ctx, conf); err != nil {
			t.Fatal(err.Error())
		}
	}
}
