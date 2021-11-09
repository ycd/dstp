package ping

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	_, err := RunTest(context.Background(), "8.8.8.8")
	if err != nil {
		t.Fatalf(err.Error())
	}

}
