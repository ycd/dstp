//go:build integration || darwin

package lookup

import (
	"context"
	"github.com/ycd/dstp/pkg/common"
	"sync"
	"testing"
)

func TestLookup(t *testing.T) {
	var wg sync.WaitGroup
	var result common.Result
	err := Host(context.Background(), &wg, "jvns.ca", &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.SystemDNS == "" {
		t.Fatal(err)
	}
}
