package lookup

import (
	"log"
	"testing"
)

func TestLookup(t *testing.T) {
	out, err := Host("https://jvns.ca")
	if err != nil {
		t.Fatal(err)
	}

	log.Println(out.String())

}