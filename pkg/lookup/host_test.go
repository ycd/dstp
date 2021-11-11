package lookup

import (
	"context"
	"log"
	"testing"
)

func TestLookup(t *testing.T) {
	out, err := Host(context.Background(), "jvns.ca")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(out.String())

}
