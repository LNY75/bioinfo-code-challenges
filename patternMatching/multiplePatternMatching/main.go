/*
Code Challenge: Solve the Multiple Pattern Matching Problem. You can use a suffix array based approach or a Burrows-Wheeler approach if you like, but you may find the latter very challenging.

Input: A string Text followed by a space-separated collection of strings Patterns.
Output: For each string Pattern in Patterns, the string Pattern followed by a colon, followed by a space-separated list of all starting positions in Text where Pattern appears as a substring.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// returns 1) BWT(Text), 2) the first column of M(text) as arrays of letters, 3) Patterns; M is the matrix of strings ordered according to the order of the suffix array of Text.
func ReadInput(input string) ([]string, []string, []string, []int) {
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

	return arrBWT, firstColumn, Patterns, suffixArr
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
	input := "input1.txt"
	BWT, FC, Patterns, SuffixArr := ReadInput(input)
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
