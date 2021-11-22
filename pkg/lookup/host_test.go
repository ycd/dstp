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
	wg.Add(1)
	err := Host(context.Background(), &wg, "jvns.ca", &result)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()

	if result.SystemDNS == "" {
		t.Fatal(err)
	}
}
