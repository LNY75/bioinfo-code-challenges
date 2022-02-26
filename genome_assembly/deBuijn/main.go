package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// read the input file to get a list of strings, which we will use later to construct the de Bruijn graph
func ReadInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	return lines
}

// takes a list of seqs and returns the de Bruijn represented by a map of string to a list of strings
func BuildDeBruijn(seqs []string) map[string][]string {
	deBruijnGraph := make(map[string][]string)
	for i := range seqs {
		prefix := seqs[i][:len(seqs[i])-1]
		suffix := seqs[i][1:]
		if deBruijnGraph[prefix] == nil {
			deBruijnGraph[prefix] = make([]string, 1)
			deBruijnGraph[prefix][0] = suffix
		} else {
			deBruijnGraph[prefix] = append(deBruijnGraph[prefix], suffix)
		}
	}
	// sort values
	for k := range deBruijnGraph {
		sort.Strings(deBruijnGraph[k])
	}
	return deBruijnGraph
}

func PrintMap(dbg map[string][]string) {
	// get a list of keys in alphabetical order
	keys := make([]string, 0)
	for k := range dbg {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Print(k, " -> ")
		for i := range dbg[k] {
			if i == len(dbg[k])-1 {
				fmt.Print(dbg[k][i])
			} else {
				fmt.Print(dbg[k][i], ",")
			}
		}
		fmt.Println(" ")
	}
}

func main() {
	input := "input2.txt"
	seqs := ReadInput(input)
	dbg := BuildDeBruijn(seqs)
	PrintMap(dbg)
}
