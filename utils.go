package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	TypeUnknown = iota
	TypeCsv
	TypeJson
	TypeTxt
)

func readFile(filename string) *bufio.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(file)
}

func getFileType(filename string) int {
	if filename == "" {
		return TypeUnknown
	}

	var ext string
	split := strings.Split(filename, ".")
	if len(split) > 1 {
		ext = split[len(split)-1]
	}
	ext = strings.ToLower(ext)

	switch ext {
	case "csv":
		return TypeCsv
	case "json":
		return TypeJson
	case "txt":
		return TypeTxt
	default:
		return TypeUnknown
	}
}
