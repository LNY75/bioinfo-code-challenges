/*
Number of Breakpoints Problem: Find the number of breakpoints in a permutation.

Input: A permutation.
Output: The number of breakpoints in this permutation.
*/
package main

import (
	"os"
	"strconv"
	"strings"
)

// returns the input permutation as an array of int
func ReadInput(input string) []int {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	pStr := strings.Fields(lines[0])
	p := make([]int, len(pStr))
	for i, s := range pStr {
		n, err := strconv.Atoi(string(s[1:]))
		if err != nil {
			panic("cannot convert this string to int")
		}
		if s[0] == '-' {
			p[i] = -n
		} else {
			p[i] = n
		}
	}
	return p
}
