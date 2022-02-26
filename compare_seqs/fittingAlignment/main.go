package main

import (
	"fmt"
)

// get backtracking pointers; returns the backtracking matrix, the score matrix, score of alignment, the row index of where the best local alignment ends
func BackTrack(lim map[string]int, scores [][]int, v, w string) ([][]string, [][]int, int, int) {
	indel := -1

	s := make([][]int, len(v)+1)
	B := make([][]string, len(v)+1)
	for i := range s {
		s[i] = make([]int, len(w)+1)
		B[i] = make([]string, len(w)+1)
	}

	// fill first column with 0
	for i := 1; i < len(s); i++ {
		s[i][0] = 0
	}
	// fill first row
	for i := 1; i < len(s[0]); i++ {
		s[0][i] = s[0][i-1] + indel
	}

	// fill B
	var match int
	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {
			match = Score(string(v[i-1]), string(w[j-1]), lim, scores)

			s[i][j] = max3(s[i-1][j]+indel, s[i][j-1]+indel, s[i-1][j-1]+match)

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
	// find best score from the last column
	max := 0
	maxi := 0
	for i, row := range s {
		score := row[len(row)-1]
		if score > max {
			max = score
			maxi = i
		}
	}
	s[len(s)-1][len(s[0])-1] = max

	// PrintIntMatrix(s)
	// fmt.Println(maxj)
	return B, s, s[len(s)-1][len(s[0])-1], maxi
}

// output the string representation of the local alignment
func OutputAlignment(B [][]string, s [][]int, v string, w string, endi int) (string, string) {
	i := endi
	j := len(B[0]) - 1
	s1 := ""
	s2 := ""

	for j > 0 {
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
	}

	return s1, s2
}

func main() {
	LetterIndexMap := InitLetterIndexMap()
	// fmt.Println(LetterIndexMap)

	scores := ReadBLOSUM62()
	// PrintIntMatrix(scores)

	input := "input.txt"
	s1, s2 := ReadInput(input)
	// fmt.Println(s1, s2)
	B, s, score, endi := BackTrack(LetterIndexMap, scores, s1, s2)
	fmt.Println(score)

	a1, a2 := OutputAlignment(B, s, s1, s2, endi)
	fmt.Println(a1)
	fmt.Println(a2)
}
