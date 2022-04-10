/*
Code Challenge: Implement BWMatching.

Input: A string BWT(Text), followed by a space-separated collection of Patterns.
Output: A space-separated list of integers, where the i-th integer corresponds to the number of substring matches of the i-th member of Patterns in Text.
*/

package main

import (
	"strconv"
)

// label repeated letter in BWT and First column with numbers
func LabelRepeatedLetters(BWT, FC []string) ([]string, []string) {
	// this map keeps track of how many times we've seen each letter
	repeatedLetters := make(map[string]int)
	for i := range BWT {
		if BWT[i] != "$" {
			repeatedLetters[BWT[i]]++
			n := repeatedLetters[BWT[i]]
			s := strconv.Itoa(n)
			BWT[i] += s
		}
	}

	repeatedLetters = make(map[string]int)

	for i := range FC {
		if FC[i] != "$" {
			repeatedLetters[FC[i]]++
			n := repeatedLetters[FC[i]]
			s := strconv.Itoa(n)
			FC[i] += s
		}
	}

	return BWT, FC
}

// construct a map such that the letter x in FC is mapped to the index of the same letter in BWT (LastColumn)
func MatchLastToFirst(BWT, FC []string) map[string]int {
	ltf := make(map[string]int)
	for i := range BWT {
		for j := range FC {
			if FC[j] == BWT[i] {
				ltf[BWT[i]] = j
			}
		}
	}
	return ltf
}

// counts the total number of matches of Pattern in Text
func BWMatching(LastColumn []string, Pattern string, LastToFirst map[string]int) (int, int) {
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
				return -1, -1
			}
		} else {
			return top, bottom
		}
	}
	return top, bottom
}
