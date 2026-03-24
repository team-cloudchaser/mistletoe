// Copyright 2026 © Team Cloudchaser

package main

import (
	"bufio"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/team-cloudchaser/mistletoe/help"
	"github.com/team-cloudchaser/mistletoe/utils"
)

var trimmedArgs []string = os.Args[1:]

func main() {
	utils.PrintBanner()
	if len(trimmedArgs) == 0 || trimmedArgs[0] == "help" {
		help.ShowUsage("./load <template> <targets>")
		help.ShowHelp("load")
		if len(trimmedArgs) != 0 {
			print("Skipped listing.\n")
			os.Exit(0)
		} else if utils.IsDir("data") {
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
		if !utils.IsDir("conf") {
			utils.PrintLevel(utils.LogError, "The config directory does not exist. Please create and populate the directory first.")
		}
		os.Exit(1)
	} else if trimmedArgs[0] == "version" {
		help.ShowHelp("copyright")
		os.Exit(0)
	} else if utils.IsFile("data/" + trimmedArgs[0] + ".json") {
		if len(trimmedArgs) < 2 {
			help.ShowUsage("./load " + trimmedArgs[0] + " <targets>")
			help.ShowHelp("loadTarget")
			var parsedTargets, parseError = utils.ParseTargetWithFallback(trimmedArgs[0])
			print(" (" + parsedTargets.Path + ")\n")
			if parseError != nil {
				utils.PrintLevel(utils.LogError, parseError.Error())
				os.Exit(1)
			}
			var targetsForSorting = make([]string, 0, len(parsedTargets.Entries))
			for target, _ := range parsedTargets.Entries {
				targetsForSorting = append(targetsForSorting, target)
			}
			slices.Sort(targetsForSorting)
			for _, target := range targetsForSorting {
				print("- " + target + "\n")
			}
			os.Exit(1)
		} else {
			var parsedTargets, parseError = utils.ParseTargetWithFallback(trimmedArgs[0])
			if parseError != nil {
				utils.PrintLevel(utils.LogError, parseError.Error())
				os.Exit(1)
			}
			var tmpFile, err0 = utils.GetTemporaryFile()
			if err0 != nil {
				utils.PrintLevel(utils.LogError, err0.Error())
				os.Exit(1)
			}
			var templateFile, err1 = utils.GetTemplateFile(trimmedArgs[0])
			if err1 != nil {
				utils.PrintLevel(utils.LogError, err1.Error())
				os.Exit(1)
			}
			var tmpFileWriter = bufio.NewWriter(tmpFile)
			var templateReader = utils.GetConstrainedScanner(templateFile)
			var writeError = parsedTargets.Select(trimmedArgs[1], templateReader, tmpFileWriter)
			tmpFile.Close()
			templateFile.Close()
			if writeError == nil {
				// Run the dialer
				var dialConfig, dialConfigError = utils.ParseDialConfigPath("conf/default.json")
				if dialConfigError != nil {
					utils.PrintLevel(utils.LogError, dialConfigError.Error())
					os.Exit(1)
				}
				utils.PrintLevel(utils.LogDebug, "%o", dialConfig)
				print("WIP")
			} else {
				utils.PrintLevel(utils.LogError, writeError.Error())
			}
			tmpFile.Sync()
		}
	} else {
		utils.PrintLevel(utils.LogError, "The specified template \"" + trimmedArgs[0] + "\" does not exist.")
		os.Exit(1)
	}
}
