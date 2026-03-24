// Copyright 2026 © Team Cloudchaser

package utils

import (
	"fmt"
)

const LogFatal uint8 = 0
const LogError uint8 = 1
const LogWarn uint8 = 2
const LogInfo uint8 = 3
const LogDebug uint8 = 4

func PrintLevel(level uint8, format string, args ...any) {
	var baseForm string = "\x1b[1;37mUnknown\x1b[0m"
	switch level {
		case LogFatal:
			baseForm = "\x1b[1;31mFatal\x1b[0m"
		case LogError:
			baseForm = "\x1b[1;31mError\x1b[0m"
		case LogWarn:
			baseForm = "\x1b[1;33mWarning\x1b[0m"
		case LogInfo:
			baseForm = "\x1b[1;32mInfo\x1b[0m"
		case LogDebug:
			baseForm = "\x1b[1;34mDebug\x1b[0m"
	}
	fmt.Printf(baseForm + ": " + format + "\n", args...)
}
