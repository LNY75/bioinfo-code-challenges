package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// records which index each letter corresponds to
func InitLetterIndexMap() map[string]int {
	lettersStr := "A  C  D  E  F  G  H  I  K  L  M  N  P  Q  R  S  T  V  W  Y"
	letters := strings.Fields(lettersStr)
	var LetterIndexMap map[string]int = make(map[string]int)
	for i, l := range letters {
		LetterIndexMap[l] = i
	}
	return LetterIndexMap
}

// extract the PAM250 scoring matrix
func ReadBLOSUM62() [][]int {
	content, err := os.ReadFile("BLOSUM62.txt")
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	s := make([][]int, 20)
	for i := range s {
		s[i] = make([]int, 20)
	}

	for i := 1; i < len(lines); i++ {
		line := strings.Fields(lines[i])
		for j := 1; j < len(line); j++ {
			score, err := strconv.Atoi(line[j])
			if err != nil {
				panic("cannot convert string to int")
			}
			s[i-1][j-1] = score
		}
	}
	return s
}

// returns the 2 input strings
func ReadInput(input string) (string, string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	return lines[0], lines[1]
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max3(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func PrintIntMatrix(s [][]int) {
	for _, r := range s {
		for _, e := range r {
			fmt.Printf(" %3d ", e)
		}
		fmt.Println(" ")
	}
}

// find the score of the letter pair according to the PAM250 scoring matrix
func Score(a, b string, lim map[string]int, scores [][]int) int {
	ai := lim[a]
	bi := lim[b]
	return scores[ai][bi]
}
