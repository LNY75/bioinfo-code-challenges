/*
Inverse Burrows-Wheeler Transform Problem: Reconstruct a string from its Burrows-Wheeler transform.

Input: A string Transform (with a single "\$$" symbol).
Output: The string Text such that BWT(Text) = Transform.
Code Challenge: Solve the Inverse Burrows-Wheeler Transform Problem.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// returns BWT(Text) and the first column of M(text) as arrays of letters; M is the matrix of strings ordered according to the order of the suffix array of Text.
func ReadInput(input string) ([]string, []string) {
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

	return arrBWT, firstColumn
}

// label repeated letter in BWT and First column with numbers
func LabelRepeatedLetters(BWT, FC []string) ([]string, []string) {
	// this map keeps track of how many times we've seen each letter
	repeatedLetters := make(map[string]int)
	repeatedLetters["A"] = 0
	repeatedLetters["C"] = 0
	repeatedLetters["G"] = 0
	repeatedLetters["T"] = 0

	for i := range BWT {
		if BWT[i] != "$" {
			repeatedLetters[BWT[i]]++
			n := repeatedLetters[BWT[i]]
			s := strconv.Itoa(n)
			BWT[i] += s
		}
	}

	repeatedLetters["A"] = 0
	repeatedLetters["C"] = 0
	repeatedLetters["G"] = 0
	repeatedLetters["T"] = 0

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

// return Text from BWT(Text) and first column of M(Text)
func BuildText(BWT, FC []string) string {
	// construct a map such that the letter x in FC is mapped to the index of the same letter in BWT
	Atoi := make(map[string]int)
	for i := range FC {
		for j := range BWT {
			if FC[i] == BWT[j] {
				Atoi[FC[i]] = j
			}
		}
	}
	// fmt.Println(Atoi)

	Text := make([]string, len(BWT))
	Text[0] = "$"
	// Build Text
	for i := 1; i < len(Text); i++ {
		current := Text[i-1]
		Text[i] = FC[Atoi[current]]
	}

	// fmt.Println(Text)

	T := "" // string representation of Text
	for i := range Text {
		T += string(Text[i][0])
	}
	T = T[1:]
	T += "$"

	return T
}

func main() {
	input := "input1.txt"
	BWT, FC := ReadInput(input)
	BWT, FC = LabelRepeatedLetters(BWT, FC)
	// fmt.Println(FC, BWT)
	Text := BuildText(BWT, FC)
	fmt.Println(Text)
}
