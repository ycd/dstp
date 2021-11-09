//go:build !windows

package dstp

import (
	"fmt"
)

func printWithColor(s string) {
	fmt.Print(s)
}
