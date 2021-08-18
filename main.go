package main

import (
	"github.com/zx06/ddnss/build"
	"github.com/zx06/ddnss/cmd"
)

func main() {
	build.PrintBuildInfo()
	cmd.RunCron()
}
