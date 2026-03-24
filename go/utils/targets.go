// Copyright 2026 © Team Cloudchaser

package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParsedTargets struct {
	Path string
	Fields []string
	Entries map[string][]string
}

func ParseTargetPath(path string) (*ParsedTargets, error) {
	var targetFile, err = os.Open(path)
	if err != nil {
		return &ParsedTargets{Path: path}, err
	}
	defer targetFile.Close()
	lineReader := bufio.NewScanner(targetFile)
	var lineBuffer []byte = make([]byte, 16384)
	lineReader.Buffer(lineBuffer, len(lineBuffer))
	var lineNumber int = 0
	var definedFields []string
	var entriesMap map[string][]string = make(map[string][]string)
	for lineReader.Scan() {
		var lineText string = lineReader.Text()
		if lineNumber < 0 || len(lineText) <= 0 {
			lineNumber ++
			continue
		} else if lineNumber == 0 {
			definedFields = strings.Split(lineText, "\t")
			if definedFields[0] != "id" {
				return &ParsedTargets{Path: path, Fields: definedFields}, errors.New("The first field must be set to \"id\".")
			}
		} else {
			var lineArray []string = strings.Split(lineText, "\t")
			if len(lineArray) != len(definedFields) {
				fmt.Println(definedFields)
				return &ParsedTargets{Path: path, Fields: definedFields}, errors.New("Amount of fields do not match on line " + strconv.Itoa(lineNumber + 1) + ". Should be " + strconv.Itoa(len(definedFields)) + " but instead got " + strconv.Itoa(len(lineArray)) + ".")
			}
			entriesMap[lineArray[0]] = lineArray[1:]
		}
		lineNumber ++
	}
	return &ParsedTargets{Path: path, Fields: definedFields[1:], Entries: entriesMap}, nil
}

func ParseTargetWithFallback(id string) (*ParsedTargets, error) {
	var parsed, err = ParseTargetPath("data/" + id + ".tsv")
	if err == nil {
		return parsed, err
	}
	return ParseTargetPath("data/default.tsv")
}

func GetTemplateLineReader(id string) (*bufio.Scanner, error) {
	var templateFile, err = os.Open("data/" + id + ".json")
	if err != nil {
		return nil, err
	}
	lineReader := bufio.NewScanner(templateFile)
	var lineBuffer []byte = make([]byte, 16384)
	lineReader.Buffer(lineBuffer, len(lineBuffer))
	return lineReader, nil
}

const tempFileName string = "generated.json"
func GetTemporaryFile() (*os.File, error) {
	os.Remove("generated.json")
	var file, err = os.OpenFile("generated.json", os.O_WRONLY | os.O_CREATE, 0640)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (pt *ParsedTargets) Select(id string, reader *bufio.Scanner, writer *bufio.Writer) bool {
	return false
}
