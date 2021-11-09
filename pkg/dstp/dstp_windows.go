package dstp

import (
	"fmt"

	"github.com/fatih/color"
)

func printWithColor(s string) {
	// https://pkg.go.dev/github.com/fatih/color@v1.13.0#readme-insert-into-noncolor-strings-sprintfunc
	fmt.Fprint(color.Output, s)
}
