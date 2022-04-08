/*
Inverse Burrows-Wheeler Transform Problem: Reconstruct a string from its Burrows-Wheeler transform.

Input: A string Transform (with a single "\$$" symbol).
Output: The string Text such that BWT(Text) = Transform.
Code Challenge: Solve the Inverse Burrows-Wheeler Transform Problem.
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
