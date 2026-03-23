package utils

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type ParsedTargets struct {
	fields []string
	entries map[string][]string
}

func ParseTargetPath(path string) (ParsedTargets, error) {
	var targetFile, err = os.Open(path)
	if err != nil {
		return ParsedTargets{}, err
	}
	defer targetFile.Close()
	lineReader := bufio.NewScanner(targetFile)
	var lineBuffer []byte = make([]byte, 16384)
	lineReader.Buffer(lineBuffer, len(lineBuffer))
	var lineNumber int = 0
	var definedFields []string
	var entriesMap map[string][]string
	for lineReader.Scan() {
		var lineText string = lineReader.Text()
		if lineNumber < 0 || len(lineText) <= 0 {
			lineNumber ++
			continue
		} else if lineNumber == 0 {
			definedFields = strings.Split(lineText, "\t")
			if definedFields[0] != "id" {
				return ParsedTargets{fields: definedFields}, errors.New("The first field must be set to \"id\".")
			}
		} else {
			var lineArray []string = strings.Split(lineText, "\t")
			if len(lineArray) != len(definedFields) {
				return ParsedTargets{fields: definedFields}, errors.New("Amount of fields do not match on line " + strconv.Itoa(lineNumber + 1) + ".")
			}
			entriesMap[definedFields[0]] = lineArray[1:]
		}
		lineNumber ++
	}
	return ParsedTargets{fields: definedFields[1:], entries: entriesMap}, nil
}
