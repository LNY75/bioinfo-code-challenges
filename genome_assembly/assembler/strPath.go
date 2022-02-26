package main

import (
	"os"
	"strings"
)

// read the input strings from the input file
// caution: the file should not contain empty lines or be an empty file
func ReadStrPathInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	return lines
}

// solve the genome path problem
func FindPath(strings []string) string {
	path := strings[0]
	for i := 1; i < len(strings); i++ {
		path += strings[i][len(strings[i])-1 : len(strings[i])]
	}
	return path
}
