/*
Code Challenge: Solve the Multiple Approximate Pattern Matching Problem.

Input: A string Text, followed by a collection of space-separated strings Patterns, and an integer d.
Output: For each string Pattern in Patterns, the string Pattern followed by a colon, followed by a space-separated collection of all positions where Pattern appears as a substring of Text with at most d mismatches.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// returns 1) BWT(Text), 2) the first column of M(text) as arrays of letters, 3) Patterns, 4) d - the maximum number of mismatches;  M is the matrix of strings ordered according to the order of the suffix array of Text.
func ReadInput(input string) ([]string, []string, []string, []int, int) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	Text := lines[0] + "$"
	// convert Text to BWT(Text)
	sufficies := make([]string, 0)
	for i := range Text {
		sufficies = append(sufficies, Text[i:]+Text[:i])
	}
	sorted := SortSufficies(sufficies)
	BWT := GetBWTransform(sorted)

	// get the suffix array:
	truncatedSufficies := make([]string, 0)
	for i := range Text {
		truncatedSufficies = append(truncatedSufficies, Text[i:])
	}
	sortedTS := SortSufficies(truncatedSufficies)
	suffixArr := GetSuffixArrayIndicies(sortedTS)

	arrBWT := make([]string, len(BWT)) // last column of M(text)
	firstColumn := make([]string, len(BWT))
	for i := range firstColumn {
		firstColumn[i] = string(BWT[i])
		arrBWT[i] = string(BWT[i])
	}
	sort.Strings(firstColumn)

	// extract patterns
	Patterns := strings.Fields(lines[1])

	// extract d:
	d, err := strconv.Atoi(lines[2])
	if err != nil {
		panic("cannot convert string to int")
	}

	return arrBWT, firstColumn, Patterns, suffixArr, d
}

func MatchPattern(LastColumn []string, Pattern string, LastToFirst map[string]int, SuffixArr []int) []int {
	top, bottom := BWMatching(LastColumn, Pattern, LastToFirst)
	matchedIndicies := make([]int, 0)
	if top == -1 {
		return matchedIndicies
	}
	for i := top; i <= bottom; i++ {
		matchedIndex := SuffixArr[i]
		matchedIndicies = append(matchedIndicies, matchedIndex)
	}
	sort.Ints(matchedIndicies)
	return matchedIndicies
}

func main() {
	input := "input.txt"
	BWT, FC, Patterns, SuffixArr, d := ReadInput(input)
	fmt.Println("d = ", d)

	BWT, FC = LabelRepeatedLetters(BWT, FC)
	LastToFirst := MatchLastToFirst(BWT, FC)
	for _, Pattern := range Patterns {
		MatchedIndicies := MatchPattern(BWT, Pattern, LastToFirst, SuffixArr)
		fmt.Print(Pattern, ": ")
		if len(MatchedIndicies) > 0 {
			for i := range MatchedIndicies {
				if i == len(MatchedIndicies)-1 {
					fmt.Print(MatchedIndicies[i])
				} else {
					fmt.Print(MatchedIndicies[i], " ")
				}
			}
		}
		fmt.Println()
	}
}
