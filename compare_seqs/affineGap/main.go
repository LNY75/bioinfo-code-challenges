package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the match reward, the mismatch and indel penalty, and the two strings for alignment
func ReadInput(input string) (int, int, int, int, string, string) {
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
	indelEx, err := strconv.Atoi(ns[3])
	return match, mismatch, indel, indelEx, lines[1], lines[2]
}

// get backtracking pointers; returns the backtracking matrix and the score of alignment
func BackTrack(reward, mismatch, indel, indelEx int, v, w string) ([][][]string, int) {
	// initialize scoring and backtracking matricies for upper, middle and lower manhattan
	su := make([][]int, len(v)+1)
	sm := make([][]int, len(v)+1)
	sl := make([][]int, len(v)+1)

	Bu := make([][]string, len(v)+1)
	Bm := make([][]string, len(v)+1)
	Bl := make([][]string, len(v)+1)

	for i := range sm {
		su[i] = make([]int, len(w)+1)
		Bu[i] = make([]string, len(w)+1)

		sm[i] = make([]int, len(w)+1)
		Bm[i] = make([]string, len(w)+1)

		sl[i] = make([]int, len(w)+1)
		Bl[i] = make([]string, len(w)+1)
	}

	// fill first element in the first column and row
	su[1][0], sm[1][0], sl[1][0] = indel, indel, indel
	su[0][1], sm[0][1], sl[0][1] = indel, indel, indel
	// fill the rest elements in the first column and row
	for i := 2; i < len(su); i++ {
		su[i][0] = su[i-1][0] + indelEx
		sm[i][0] = sm[i-1][0] + indelEx
		sl[i][0] = sl[i-1][0] + indelEx
	}
	for i := 2; i < len(su[0]); i++ {
		su[0][i] = su[0][i-1] + indelEx
		sm[0][i] = sm[0][i-1] + indelEx
		sl[0][i] = sl[0][i-1] + indelEx
	}

	// fill Backtracking matricies
	var match int
	for i := 1; i <= len(v); i++ {
		for j := 1; j <= len(w); j++ {
			// lower level sl_i,j
			sl[i][j] = max2(sl[i-1][j]+indelEx, sm[i-1][j]+indel)

			// upper level su_i,j
			su[i][j] = max2(su[i][j-1]+indelEx, sm[i][j-1]+indel)

			match = mismatch
			if v[i-1] == w[j-1] {
				match = reward
			}
			// middle level sm_i,j
			sm[i][j] = max3(su[i][j], sl[i][j], sm[i-1][j-1]+match)

			// fill backtracking matrices
			if sl[i][j] == sl[i-1][j]+indelEx {
				Bl[i][j] = "so" // go south in current level
			} else {
				Bl[i][j] = "su" // go south, then up a level
			}

			if su[i][j] == su[i][j-1]+indelEx {
				Bu[i][j] = "ea" // go east in current level
			} else {
				Bu[i][j] = "ed" // go east, go down a level
			}

			if sm[i][j] == su[i][j] {
				Bm[i][j] = "da" // go down a level
			} else if sm[i][j] == sl[i][j] {
				Bm[i][j] = "up" // go up a level
			} else {
				Bm[i][j] = "se" // go southeast in current level
			}

		}
	}

	maxScore := sm[len(v)][len(w)]

	Bs := make([][][]string, 3)
	Bs[0] = Bl
	Bs[1] = Bm
	Bs[2] = Bu

	fmt.Println(maxScore)

	// PrintStrMatrix(Bu)
	// PrintStrMatrix(Bm)
	// PrintStrMatrix(Bl)

	// PrintIntMatrix(su)
	// PrintIntMatrix(sm)
	// PrintIntMatrix(sl)

	return Bs, maxScore
}

// output the string representation of the alignment
// Bs: a collection of three backtracking matricies
// Bi: index of the backtracking matrix in Bs that gave the max score
func OutputAlignment(Bs [][][]string, v, w string) (string, string) {
	i := len(v)
	j := len(w)
	s1 := ""
	s2 := ""

	currentB := Bs[1]
	currentBi := 1
	for i > 0 && j > 0 {
		// fmt.Println(i, j, currentBi, currentB[i][j])

		switch currentB[i][j] {
		case "so": // we are in Bl
			i--
			s1 = string(v[i]) + s1
			s2 = "-" + s2
		case "su":
			i--
			s1 = string(v[i]) + s1
			s2 = "-" + s2
			currentBi++
		case "ea": // we are in Bu
			j--
			s1 = "-" + s1
			s2 = string(w[j]) + s2
		case "ed":
			j--
			s1 = "-" + s1
			s2 = string(w[j]) + s2
			currentBi--
		case "da": // we are in Bm
			currentBi++
		case "up":
			currentBi--
		case "se":
			i--
			j--
			s1 = string(v[i]) + s1
			s2 = string(w[j]) + s2
		}
		currentB = Bs[currentBi]
	}

	// i := len(B) - 1
	// j := len(B[0]) - 1
	// s1 := ""
	// s2 := ""
	// for i > 0 && j > 0 {
	// 	if B[i][j] == "so" {
	// 		i--
	// 		s1 = string(v[i]) + s1
	// 		s2 = "-" + s2
	// 	} else if B[i][j] == "ea" {
	// 		j--
	// 		s1 = "-" + s1
	// 		s2 = string(w[j]) + s2
	// 	} else {
	// 		i--
	// 		j--
	// 		s1 = string(v[i]) + s1
	// 		s2 = string(w[j]) + s2
	// 	}
	// }

	fmt.Println(s1)
	fmt.Println(s2)
	return s1, s2
}

func main() {
	input := "input.txt"
	match, mismatch, indel, indelEx, s1, s2 := ReadInput(input)
	mismatch = -mismatch
	indel = -indel
	indelEx = -indelEx
	// fmt.Println(match, mismatch, indel, indelEx, s1, s2)

	Bs, _ := BackTrack(match, mismatch, indel, indelEx, s1, s2)
	OutputAlignment(Bs, s1, s2)
}
