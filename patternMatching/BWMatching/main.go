/*
Code Challenge: Implement BWMatching.

Input: A string BWT(Text), followed by a space-separated collection of Patterns.
Output: A space-separated list of integers, where the i-th integer corresponds to the number of substring matches of the i-th member of Patterns in Text.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// returns 1) BWT(Text), 2) the first column of M(text) as arrays of letters, 3) Patterns; M is the matrix of strings ordered according to the order of the suffix array of Text.
func ReadInput(input string) ([]string, []string, []string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	BWT := lines[0]
	arrBWT := make([]string, len(BWT))

	firstColumn := make([]string, len(BWT))
	for i := range firstColumn {
		firstColumn[i] = string(BWT[i])
		arrBWT[i] = string(BWT[i])
	}
	sort.Strings(firstColumn)

	// extract patterns
	Patterns := strings.Fields(lines[1])

	return arrBWT, firstColumn, Patterns
}

// counts the total number of matches of Pattern in Text
func BWMatching(LastColumn []string, Pattern string, LastToFirst map[string]int) int {
	top := 0
	bottom := len(LastColumn) - 1
	for top <= bottom {
		if len(Pattern) > 0 {
			symbol := Pattern[len(Pattern)-1] // last letter in Pattern
			// remove last letter from Pattern
			Pattern = Pattern[:len(Pattern)-1]

			positions := make([]int, 0) // contain positions from top to bottom in LastColumn contain an occurrence of symbol
			for i := top; i <= bottom; i++ {
				if LastColumn[i][0] == symbol {
					positions = append(positions, i)
				}
			}
			if len(positions) > 0 {
				topIndex := positions[0]                   // first position of symbol among positions from top to bottom in LastColumn
				bottomIndex := positions[len(positions)-1] // last position of symbol among positions from top to bottom in LastColumn
				top = LastToFirst[LastColumn[topIndex]]
				bottom = LastToFirst[LastColumn[bottomIndex]]
			} else {
				return 0
			}
		} else {
			return bottom - top + 1
		}
	}
	return bottom - top + 1
}

func main() {
	input := "input.txt"
	BWT, FC, Patterns := ReadInput(input)
	BWT, FC = LabelRepeatedLetters(BWT, FC)
	LastToFirst := MatchLastToFirst(BWT, FC)
	for _, Pattern := range Patterns {
		numMatched := BWMatching(BWT, Pattern, LastToFirst)
		fmt.Print(numMatched, " ")
	}
}
