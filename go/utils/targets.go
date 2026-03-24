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
	lineReader := GetConstrainedScanner(targetFile)
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

func ParseTargetWithFallback(templateId string) (*ParsedTargets, error) {
	var parsed, err = ParseTargetPath("data/" + templateId + ".tsv")
	if err == nil {
		return parsed, err
	}
	return ParseTargetPath("data/default.tsv")
}

func GetTemplateFile(templateId string) (*os.File, error) {
	var templateFile, err = os.Open("data/" + templateId + ".json")
	if err != nil {
		return nil, err
	}
	return templateFile, nil
}

const tempFileName string = "generated.json"
func GetTemporaryFile() (*os.File, error) {
	var err0 = os.Remove("generated.json")
	if err0 != nil {
		return nil, err0
	}
	var file, err1 = os.OpenFile("generated.json", os.O_WRONLY | os.O_CREATE, 0640)
	if err1 != nil {
		return nil, err1
	}
	return file, nil
}

func (pt *ParsedTargets) Select(targetId string, reader *bufio.Scanner, writer *bufio.Writer) error {
	var fields, exists = pt.Entries[targetId]
	if (!exists) {
		return errors.New("Selected target does not exist.")
	}
	var replaceMap = make([]string, 0, len(pt.Fields) << 1)
	for i, key := range pt.Fields {
		replaceMap = append(replaceMap, "__" + strings.ToUpper(key) + "__", fields[i])
	}
	var replacer = strings.NewReplacer(replaceMap...)
	for reader.Scan() {
		var replacedText = replacer.Replace(reader.Text())
		//fmt.Println(replacedText)
		var _, err = writer.WriteString(replacedText + "\n")
		if err != nil {
			return errors.New("Failed to write to the temporary file.")
		}
		writer.Flush()
	}
	return nil
}
