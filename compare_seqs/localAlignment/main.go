package main

import (
	"fmt"
)

// get backtracking pointers; returns the backtracking matrix, the score matrix, score of alignment, the row and column indices of where the best local alignment ends
func BackTrack(lim map[string]int, scores [][]int, v, w string) ([][]string, [][]int, int, int, int) {
	indel := -5

	s := make([][]int, len(v)+1)
	B := make([][]string, len(v)+1)
	for i := range s {
		s[i] = make([]int, len(w)+1)
		B[i] = make([]string, len(w)+1)
	}

	// fill first column and row
	for i := 1; i < len(s); i++ {
		s[i][0] = s[i-1][0] + indel
	}
	for i := 1; i < len(s[0]); i++ {
		s[0][i] = s[0][i-1] + indel
	}

	// fill B
	var match int
	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {
			match = Score(string(v[i-1]), string(w[j-1]), lim, scores)

			s[i][j] = max4(s[i-1][j]+indel, s[i][j-1]+indel, s[i-1][j-1]+match, 0)

			if s[i][j] == s[i-1][j]+indel {
				B[i][j] = "so" // south -> down
			} else if s[i][j] == s[i][j-1]+indel {
				B[i][j] = "ea" // east -> right
			} else if s[i][j] == s[i-1][j-1]+match {
				B[i][j] = "se" // southeast -> downright
			}
		}
	}

	// figure out s_n,m
	max := 0
	maxi := 0
	maxj := 0
	for i, r := range s {
		for j, score := range r {
			if score > max {
				max = score
				maxi = i
				maxj = j
			}
		}
	}
	s[len(s)-1][len(s[0])-1] = max

	// PrintIntMatrix(s)
	return B, s, s[len(s)-1][len(s[0])-1], maxi, maxj
}

// output the string representation of the local alignment
func OutputAlignment(B [][]string, s [][]int, v string, w string, endi, endj int) (string, string) {
	i := endi
	j := endj
	s1 := ""
	s2 := ""

	currentScore := s[len(s)-1][len(s[0])-1]
	for currentScore > 0 {
		if B[i][j] == "so" {
			i--
			s1 = string(v[i]) + s1
			s2 = "-" + s2
		} else if B[i][j] == "ea" {
			j--
			s1 = "-" + s1
			s2 = string(w[j]) + s2
		} else {
			i--
			j--
			s1 = string(v[i]) + s1
			s2 = string(w[j]) + s2
		}
		currentScore = s[i][j]
	}

	return s1, s2
}

func main() {
	LetterIndexMap := InitLetterIndexMap()
	// fmt.Println(LetterIndexMap)

	scores := ReadPAM250()
	// PrintIntMatrix(scores)

	input := "input.txt"
	s1, s2 := ReadInput(input)
	// fmt.Println(s1, s2)
	B, s, score, endi, endj := BackTrack(LetterIndexMap, scores, s1, s2)
	fmt.Println(score)

	a1, a2 := OutputAlignment(B, s, s1, s2, endi, endj)
	fmt.Println(a1)
	fmt.Println(a2)
}
