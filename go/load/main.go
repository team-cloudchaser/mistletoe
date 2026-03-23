package main

import (
	"os"
	"github.com/team-cloudchaser/mistletoe/utils"
	"github.com/team-cloudchaser/mistletoe/help"
)

var trimmedArgs []string = os.Args[1:]

func main() {
	utils.PrintBanner()
	if len(trimmedArgs) == 0 {
		help.ShowHelp("load")
		os.Exit(1)
	}
}
