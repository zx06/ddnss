package build

import (
	"fmt"
)

// Will be changed at build time via -ldflags
var Version = "debug"
var Commit = "-"
var BuildDate = "-"

func PrintBuildInfo()  {
	fmt.Println("=============Build info=============")
	fmt.Printf("  Version: %s\n", Version)
	fmt.Printf("  Commit: %s\n", Commit)
	fmt.Printf("  Build date: %s\n", BuildDate)
	fmt.Println("====================================")
}