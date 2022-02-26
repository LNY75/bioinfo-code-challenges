package main

import (
	"fmt"
	"os"
	"strings"
)

// returns the match reward, the mismatch and indel penalty, and the two strings for alignment
// score: if all 3 symbols of a column are the same, score is 1; otherwise the score is 0
func ReadInput(input string) (string, string, string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// extract the match reward, mismatch and indel penalties:
	return lines[0], lines[1], lines[2]
}

// get backtracking pointers; returns the backtracking matrix and the score of alignment
func BackTrack(s1, s2, s3 string) ([][][]string, int) {
	s := make([][][]int, len(s1)+1)
	B := make([][][]string, len(s1)+1)
	for i := range s {
		s[i] = make([][]int, len(s2)+1)
		B[i] = make([][]string, len(s2)+1)
	}
	for i := range s {
		for j := range s[i] {
			s[i][j] = make([]int, len(s3)+1)
			B[i][j] = make([]string, len(s3)+1)
		}
	}

	// fill B
	var match int = 0
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			for k := 1; k <= len(s3); k++ {
				match = 0
				if s1[i-1] == s2[j-1] && s2[j-1] == s3[k-1] {
					match = 1
				}
				scores := []int{
					s[i-1][j][k],
					s[i][j-1][k],
					s[i][j][k-1],
					s[i-1][j-1][k],
					s[i-1][j][k-1],
					s[i][j-1][k-1],
					s[i-1][j-1][k-1] + match,
				}
				s[i][j][k] = max(scores)

				switch s[i][j][k] {
				case s[i-1][j][k]:
					B[i][j][k] = "i"
				case s[i][j-1][k]:
					B[i][j][k] = "j"
				case s[i][j][k-1]:
					B[i][j][k] = "k"
				case s[i-1][j-1][k]:
					B[i][j][k] = "ij"
				case s[i-1][j][k-1]:
					B[i][j][k] = "ik"
				case s[i][j-1][k-1]:
					B[i][j][k] = "jk"
				case s[i-1][j-1][k-1] + match:
					B[i][j][k] = "ijk"
				}

			}
		}
	}

	// PrintIntMatrix(s)
	return B, s[len(s1)][len(s2)][len(s3)]
}

// output the string representation of the alignment
func OutputAlignment(B [][][]string, s1, s2, s3 string) (string, string, string) {
	i := len(B) - 1
	j := len(B[0]) - 1
	k := len(B[0][0]) - 1

	a1 := ""
	a2 := ""
	a3 := ""

	for i > 0 && j > 0 && k > 0 {
		switch B[i][j][k] {
		case "i":
			i--
			a1 = string(s1[i]) + a1
			a2 = "-" + a2
			a3 = "-" + a3
		case "j":
			j--
			a1 = "-" + a1
			a2 = string(s2[j]) + a2
			a3 = "-" + a3
		case "k":
			k--
			a1 = "-" + a1
			a2 = "-" + a2
			a3 = string(s3[k]) + a3
		case "ij":
			i--
			j--
			a1 = string(s1[i]) + a1
			a2 = string(s2[j]) + a2
			a3 = "-" + a3
		case "ik":
			i--
			k--
			a1 = string(s1[i]) + a1
			a2 = "-" + a2
			a3 = string(s3[k]) + a3
		case "jk":
			j--
			k--
			a1 = "-" + a1
			a2 = string(s2[j]) + a2
			a3 = string(s3[k]) + a3
		case "ijk":
			i--
			j--
			k--
			a1 = string(s1[i]) + a1
			a2 = string(s2[j]) + a2
			a3 = string(s3[k]) + a3
		}
	}

	for i > 0 && j > 0 {
		i--
		j--
		a1 = string(s1[i]) + a1
		a2 = string(s2[j]) + a2
		a3 = "-" + a3
	}

	for i > 0 && k > 0 {
		i--
		k--
		a1 = string(s1[i]) + a1
		a2 = "-" + a2
		a3 = string(s3[k]) + a3
	}

	for j > 0 && k > 0 {
		j--
		k--
		a1 = "-" + a1
		a2 = string(s2[j]) + a2
		a3 = string(s3[k]) + a3
	}

	for i > 0 {
		i--
		a1 = string(s1[i]) + a1
		a2 = "-" + a2
		a3 = "-" + a3
	}

	for j > 0 {
		j--
		a1 = "-" + a1
		a2 = string(s2[j]) + a2
		a3 = "-" + a3
	}

	for k > 0 {
		k--
		a1 = "-" + a1
		a2 = "-" + a2
		a3 = string(s3[k]) + a3
	}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	return a1, a2, a3
}

func main() {
	input := "input.txt"
	s1, s2, s3 := ReadInput(input)
	fmt.Println(s1, s2, s3)
	B, score := BackTrack(s1, s2, s3)
	fmt.Println(score)
	OutputAlignment(B, s1, s2, s3)
}
