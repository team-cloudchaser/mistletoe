// Copyright 2026 © Team Cloudchaser

package utils

import (
	"encoding/json"
	"errors"
	"os"
)

const TmpFileName string = "generated.json"
func GetTemporaryFile() (*os.File, error) {
	var err0 = os.Remove("generated.json")
	if err0 != nil {
		return nil, err0
	}
	var file, err1 = os.OpenFile(TmpFileName, os.O_WRONLY | os.O_CREATE, 0640)
	if err1 != nil {
		return nil, err1
	}
	return file, nil
}

type DialConfig struct {
	IsValidConfig bool
	Path string
	Run []string `json:"run,omitempty"`
}

func ParseDialConfigPath(path string) (*DialConfig, error) {
	var parsedDialConfig = DialConfig{Path: path, IsValidConfig: false}
	var dialConfigFile, err0 = os.Open(path)
	if err0 != nil {
		return &parsedDialConfig, err0
	}
	defer dialConfigFile.Close()
	var dialConfigFileInfo, _ = dialConfigFile.Stat()
	var dialConfigFileSize int64 = dialConfigFileInfo.Size()
	var dialConfigFileBuffer []byte = make([]byte, dialConfigFileSize)
	var _, err1 = dialConfigFile.Seek(0, 0)
	if err1 != nil {
		return &parsedDialConfig, err1
	}
	var dialConfigReadSize, err2 = dialConfigFile.Read(dialConfigFileBuffer)
	if err2 != nil {
		return &parsedDialConfig, err2
	} else if int64(dialConfigReadSize) != dialConfigFileSize {
		return &parsedDialConfig, errors.New("The dial config cannot be read in full.")
	}
	var jsonReadError error = json.Unmarshal(dialConfigFileBuffer, &parsedDialConfig)
	if jsonReadError == nil {
		parsedDialConfig.IsValidConfig = true
	} else {
		return &parsedDialConfig, jsonReadError
	}
	return &parsedDialConfig, nil
}
