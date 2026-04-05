// Copyright 2026 © Team Cloudchaser

package utils

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
)

var CwdPath, _ = os.Getwd()
var CwdFs = os.DirFS(".")

func GetPath(path string) string {
	return filepath.Join(CwdPath, path)
}

func IsFile(path string) bool {
	var fileInfo, err = fs.Stat(CwdFs, path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func IsDir(path string) bool {
	var fileInfo, err = fs.Stat(CwdFs, path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func GetConstrainedScanner(file *os.File) *bufio.Scanner {
	lineReader := bufio.NewScanner(file)
	var lineBuffer []byte = make([]byte, 1048576)
	lineReader.Buffer(lineBuffer, len(lineBuffer))
	return lineReader
}
