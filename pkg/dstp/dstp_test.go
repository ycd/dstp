package dstp

import (
	"context"
	"github.com/ycd/dstp/config"
	"testing"
)

func TestRunAllTests(t *testing.T) {
	ctx := context.Background()
	c := config.Config{
		Addr:       "8.8.8.8",
		Output:     "plaintext",
		ShowHelp:   false,
		Concurrent: false,
	}

	if err := RunAllTests(ctx, c); err != nil {
		t.Fatal(err.Error())
	}
}
