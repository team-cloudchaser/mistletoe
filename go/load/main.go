package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/team-cloudchaser/mistletoe/help"
	"github.com/team-cloudchaser/mistletoe/utils"
)

var trimmedArgs []string = os.Args[1:]

func main() {
	utils.PrintBanner()
	if len(trimmedArgs) == 0 {
		help.ShowUsage("./load <template> <targets>")
		help.ShowHelp("load")
		if utils.IsDir("data") {
			var dataFiles, _ = os.ReadDir(utils.GetPath("data"))
			for _, dataFile := range dataFiles {
				var fileName = dataFile.Name()
				var fileExt = filepath.Ext(fileName)
				if fileExt == ".json" {
					var cleanedName, _ = strings.CutSuffix(fileName, fileExt)
					print("- " + cleanedName + "\n")
				}
			}
		} else {
			utils.PrintLevel(utils.LogError, "The data directory does not exist. Please create and populate the directory first.")
		}
		os.Exit(1)
	} else if utils.IsFile("data/" + trimmedArgs[0] + ".json") {
		help.ShowUsage("./load " + trimmedArgs[0] + " <targets>")
		help.ShowHelp("loadTarget")
	} else {
		utils.PrintLevel(utils.LogError, "The specified template \"" + trimmedArgs[0] + "\" does not exist.")
	}
}
