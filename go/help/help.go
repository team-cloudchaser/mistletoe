package help

import (
	"embed"
	"io/fs"
)

//go:embed man/*.txt
var helpFiles embed.FS

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
