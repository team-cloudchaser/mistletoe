package utils

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/team-cloudchaser/mistletoe"
)

func VersionToString(version []uint8) string {
	var serialized []string = make([]string, len(version))
	for i0, e0 := range version {
		serialized[i0] = strconv.FormatUint(uint64(e0), 10)
	}
	return strings.Join(serialized, ".")
}

var VersionString string = VersionToString(mistletoe.Version)

func PrintBanner() {
	fmt.Printf("\x1b[1;37mMistletoe v%s\x1b[0m\n\n", VersionString)
}
