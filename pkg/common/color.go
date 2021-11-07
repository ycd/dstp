package common

import "github.com/fatih/color"

// Bunch of pre-defined coloring functions
var (
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Cyan    = color.New(color.FgCyan).SprintFunc()
	White   = color.New(color.FgHiWhite).SprintFunc()
)
