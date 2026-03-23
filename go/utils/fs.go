package utils

import (
	"io/fs"
	"os"
)

var CwdFs = os.DirFS(".")

func IsDir(path string) bool {
	fileInfo, err := fs.Stat(CwdFs, path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
