//go:build integration || darwin

package lookup

import (
	"context"
	"sync"
	"testing"

	"github.com/ycd/dstp/pkg/common"
)

func TestLookup(t *testing.T) {
	var wg sync.WaitGroup
	var result common.Result
	wg.Add(1)
	err := Host(context.Background(), &wg, common.Address("jvns.ca"), "8.8.8.8", &result)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()

	if result.DNS.Error != nil {
		t.Fatal(result.DNS.Error)
	}

	if result.SystemDNS.Content == "" {
		t.Fatal("System DNS resolution failed")
	}

}
