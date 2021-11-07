package ping

import (
	"context"
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	out, err := RunTest(context.Background(), "8.8.8.8")
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println("ping: " + out)

}
