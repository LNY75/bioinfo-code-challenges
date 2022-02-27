package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the column removal threshold, the alphabet, the alignment strings, and the list of columns that exceeds the column removal threshold
func ReadInput(input string) (float64, map[string]int, []string, []int) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// extract the column removal threshold
	columnRemovalThresholdStr := lines[0]
	crt, err := strconv.ParseFloat(columnRemovalThresholdStr, 64)
	if err != nil {
		panic("cannot convert string to float")
	}

	// extract alphabet
	alphabetStr := strings.Fields(lines[2])
	alphabet := make(map[string]int)
	for i, a := range alphabetStr {
		alphabet[a] = i
	}

	// extract the alignment strings
	A := make([]string, 0)
	for i := 4; i < len(lines); i++ {
		A = append(A, lines[i])
	}

	crtIndexArr := make([]int, 0)
	for j := 0; j < len(A[0]); j++ {
		c := 0.0
		for i := 0; i < len(A); i++ {
			if A[i][j] == '-' {
				c++
			}
		}
		if c/float64(len(A)) > crt {
			crtIndexArr = append(crtIndexArr, j)
		}
	}

	fmt.Println(crt)
	fmt.Println(alphabet)
	PrintAlignment(A)
	fmt.Println(crtIndexArr)

	return crt, alphabet, A, crtIndexArr
}

// combine columns in the alignment that exceeds the crt
func CombCols(A []string, crtIndexArr []int) [][]int {
	// GA stores consecutive indicies in one list
	GA := make([][]int, 1)
	GA[0] = make([]int, 0)
	GA[0] = append(GA[0], crtIndexArr[0])

	// group consecutive indecies
	for i := 1; i < len(crtIndexArr); i++ {
		curIndex := crtIndexArr[i]
		prevIndex := crtIndexArr[i-1]
		if curIndex-prevIndex == 1 {
			GA[len(GA)-1] = append(GA[len(GA)-1], curIndex)
		} else {
			GA = append(GA, make([]int, 0))
			GA[len(GA)-1] = append(GA[len(GA)-1], curIndex)
		}
	}
	fmt.Println(GA)

	// group strings in A according to GA

	return GA
}

// for each string in the alignment strings, figure out the states it followed and record the symbols emitted from each state
func StatesAndEms(A []string) {

}

func main() {
	input := "input.txt"
	_, _, A, crtIndexArr := ReadInput(input)
	CombCols(A, crtIndexArr)
}
