package help

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed man/*.txt
var helpFiles embed.FS

func ShowUsage(format string, args ...any) {
	fmt.Printf("\x1b[1;36mUsage\x1b[0m: " + format + "\n", args...)
}
func ShowHelp(topic string) {
	var intendedPath string = "man/" + topic + ".txt"
	if fs.ValidPath(intendedPath) {
		contents, _ := helpFiles.ReadFile(intendedPath)
		print(string(contents))
	} else {
		contents, _ := helpFiles.ReadFile("man/fail.txt")
		print(string(contents))
	}
}
