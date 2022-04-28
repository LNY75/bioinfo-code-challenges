/*
Suffix Array Construction Problem: Construct the suffix array of a string.

Input: A string Text.
Output: SuffixArray(Text), as a space-separated collection of integers.

Code Challenge: Solve the Suffix Array Construction Problem.
*/
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// returns the Text and the Pattern
func ReadInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	Text := lines[0]
	Pattern := make([]string, 0)
	for i := range Text {
		Pattern = append(Pattern, Text[i:]+Text[:i])
	}

	return Pattern
}

func SortSufficies(Sufficies []string) []string {
	sort.Strings(Sufficies)
	return Sufficies
}

func PrintSufficies(sufficies []string) {
	for _, s := range sufficies {
		fmt.Println(s)
	}
}

// returns the actual suffix array, which should be an array of ints rather than strings
func GetSuffixArrayIndicies(SortedSufficies []string) []int {
	SuffixArray := make([]int, len(SortedSufficies))
	L := len(SortedSufficies) - 1
	for i := range SuffixArray {
		SuffixArray[i] = L - (len(SortedSufficies[i]) - 1)
	}
	return SuffixArray
}
