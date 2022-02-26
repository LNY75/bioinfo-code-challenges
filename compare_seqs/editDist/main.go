package main

import (
	"fmt"
	"os"
	"strings"
)

// returns the match reward, the mismatch and indel penalty, and the two strings for alignment
func ReadInput(input string) (string, string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	return lines[0], lines[1]
}

// get backtracking pointers; returns the backtracking matrix and the score of alignment
func BackTrack(reward, mismatch, indel int, v, w string) [][]string {
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
			match = mismatch
			if v[i-1] == w[j-1] {
				match = reward
			}
			s[i][j] = max(s[i-1][j]+indel, s[i][j-1]+indel, s[i-1][j-1]+match)

			if s[i][j] == s[i-1][j]+indel {
				B[i][j] = "so" // south -> down
			} else if s[i][j] == s[i][j-1]+indel {
				B[i][j] = "ea" // east -> right
			} else if s[i][j] == s[i-1][j-1]+match {
				B[i][j] = "se" // southeast -> downright
			}
		}
	}
	// PrintIntMatrix(s)
	return B
}

// calculate the edit distance
func EditDist(B [][]string, v string, w string) (int, string, string) {
	ed := 0
	i := len(B) - 1
	j := len(B[0]) - 1
	s1 := ""
	s2 := ""
	for i > 0 && j > 0 {
		if B[i][j] == "so" {
			i--
			s1 = string(v[i]) + s1
			s2 = "-" + s2
			ed++
		} else if B[i][j] == "ea" {
			j--
			s1 = "-" + s1
			s2 = string(w[j]) + s2
			ed++
		} else {
			i--
			j--
			if v[i] != w[j] {
				ed++
			}
			s1 = string(v[i]) + s1
			s2 = string(w[j]) + s2
		}
	}

	return ed, s1, s2
}

func max(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func main() {
	input := "input.txt"
	s1, s2 := ReadInput(input)
	B := BackTrack(2, -1, -2, s1, s2)
	ed, _, _ := EditDist(B, s1, s2)
	fmt.Println(ed)
	// fmt.Println(a1)
	// fmt.Println(a2)
}
