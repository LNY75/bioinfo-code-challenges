/*
Suffix Array Construction Problem: Construct the suffix array of a string.

Input: A string Text.
Output: SuffixArray(Text), as a space-separated collection of integers.

Code Challenge: Solve the Suffix Array Construction Problem.
*/
package main

import (
	"fmt"
	"sort"
)

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

// func main() {
// 	input := "input1.txt"
// 	sufficies := ReadInput(input)
// 	sorted := SortSufficies(sufficies)
// 	SuffixArray := GetSuffixArrayIndicies(sorted)
// 	fmt.Println(SuffixArray)
// }
