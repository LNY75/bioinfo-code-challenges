package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the match reward, the mismatch and indel penalty, and the two strings for alignment
func ReadInput(input string) (int, int, int, string, string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// extract the match reward, mismatch and indel penalties:
	ns := strings.Split(lines[0], " ")
	match, err := strconv.Atoi(ns[0])
	if err != nil {
		panic("cannot convert to int")
	}
	mismatch, err := strconv.Atoi(ns[1])
	indel, err := strconv.Atoi(ns[2])
	return match, mismatch, indel, lines[1], lines[2]
}

// get backtracking pointers; returns the backtracking matrix and the score of alignment
func BackTrack(reward, mismatch, indel int, v, w string) ([][]string, int) {
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
	return B, s[len(s)-1][len(s[0])-1]
}

// output the string representation of the alignment
func OutputAlignment(B [][]string, v string, w string) (string, string) {
	i := len(B) - 1
	j := len(B[0]) - 1
	s1 := ""
	s2 := ""
	for i > 0 && j > 0 {
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

	for i > 0 {
		i--
		s1 = string(v[i]) + s1
		s2 = "-" + s2
	}

	for j > 0 {
		j--
		s1 = "-" + s1
		s2 = string(w[j]) + s2
	}

	return s1, s2
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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

func PrintIntMatrix(s [][]int) {
	for _, r := range s {
		for _, e := range r {
			fmt.Print(e, " ")
		}
		fmt.Println(" ")
	}
}

func main() {
	input := "input1.txt"
	match, mismatch, indel, s1, s2 := ReadInput(input)
	mismatch = -mismatch
	indel = -indel
	// fmt.Println(match, mismatch, indel, s1, s2)

	B, score := BackTrack(match, mismatch, indel, s1, s2)
	fmt.Println(score)

	m1, m2 := OutputAlignment(B, s1, s2)
	fmt.Println(m1)
	fmt.Println(m2)
}
