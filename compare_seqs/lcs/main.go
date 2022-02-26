package main

import (
	"fmt"
	"os"
	"strings"
)

// read two seqs from input
func ReadInput(input string) (string, string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	l1 := strings.Split(lines[0], " ")
	l2 := strings.Split(lines[1], " ")

	return l1[0], l2[0]
}

// get backtracking pointers; returns the backtracking matrix
func LCSBackTrack(v, w string) [][]string {
	s := make([][]int, len(v)+1)
	B := make([][]string, len(v)+1)
	for i := range s {
		s[i] = make([]int, len(w)+1)
		B[i] = make([]string, len(w)+1)
	}

	// fill B
	var match int
	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {
			match = 0
			if v[i-1] == w[j-1] {
				match = 1
			}
			s[i][j] = max(s[i-1][j], s[i][j-1], s[i-1][j-1]+match)
			if s[i][j] == s[i-1][j] {
				B[i][j] = "so" // south -> down
			} else if s[i][j] == s[i][j-1] {
				B[i][j] = "ea" // east -> right
			} else if s[i][j] == s[i-1][j-1]+match {
				B[i][j] = "se" // southeast -> downright
			}
		}
	}

	return B
}

// find the LCS from the backtracking matrix
// should initially invoke OutputLCS(backtrack, v, |v|, |w|)
func OutputLCS(B [][]string, v string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	}
	if B[i][j] == "so" {
		return OutputLCS(B, v, i-1, j)
	} else if B[i][j] == "ea" {
		return OutputLCS(B, v, i, j-1)
	} else {
		return OutputLCS(B, v, i-1, j-1) + string(v[i-1])
	}
}

func PrintBackTrack(B [][]string) {
	for _, r := range B {
		for _, e := range r {
			fmt.Print(e, " ")
		}
		fmt.Println(" ")
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

func main() {
	input := "input1.txt"
	s1, s2 := ReadInput(input)
	// fmt.Println(s1, s2)
	B := LCSBackTrack(s1, s2)
	// PrintBackTrack(B)
	lcs := OutputLCS(B, s1, len(s1), len(s2))
	// fmt.Println(len(lcs))
	fmt.Println(lcs)
}
