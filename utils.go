package main

import (
	"regexp"
	"strings"
)

const (
	TypeUnknown = iota
	TypeCsv
	TypeJson
	TypeTxt
)

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

// countEmails takes a string and returns the number of email addresses in it.
func countEmails(s string) int {
	pattern := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
	re := regexp.MustCompile(pattern)
	return len(re.FindAllString(s, -1))
}
