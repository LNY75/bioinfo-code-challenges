package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the observed outcome string x, the alphabet, states, transition matrix, and emission matrix for the HMM
func ReadInput(input string) (string, map[string]int, []string, [][]float64, [][]float64) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// extract the match reward, mismatch and indel penalties:
	x := lines[0]

	alphabetStr := strings.Fields(lines[2])
	alphabet := make(map[string]int)
	for i, a := range alphabetStr {
		alphabet[a] = i
	}

	states := strings.Fields(lines[4])

	stateTransM := make([][]float64, len(states))
	emissionM := make([][]float64, len(states))

	// the state transition probability matrix starts at line 7
	transitionsStr := lines[7 : 7+len(states)]
	for l, line := range transitionsStr {
		row := strings.Fields(line)[1:]
		floatRow := make([]float64, len(row))
		for i := range row {
			f, err := strconv.ParseFloat(row[i], 64)
			if err != nil {
				panic("cannot convert string to float")
			}
			floatRow[i] = f
		}
		stateTransM[l] = floatRow
	}

	emissionsStr := lines[7+len(states)+2 : 7+len(states)+2+len(states)]
	for l, line := range emissionsStr {
		row := strings.Fields(line)[1:]
		floatRow := make([]float64, len(row))
		for i := range row {
			f, err := strconv.ParseFloat(row[i], 64)
			if err != nil {
				panic("cannot convert string to float")
			}
			floatRow[i] = f
		}
		emissionM[l] = floatRow
	}

	return x, alphabet, states, stateTransM, emissionM
}

func ViterbiDP(x string, states []string, alphabet map[string]int, STM [][]float64, EMM [][]float64) ([][]float64, [][]int, int) {
	// initialize the score and backtrack matrices
	// row = # of states
	// column = length of x + 1
	s := make([][]float64, len(states))
	B := make([][]int, len(states))
	// initialize the first column to 1
	for i := range s {
		s[i] = make([]float64, len(x)+1)
		s[i][0] = 1
		s[i][1] = EMM[i][alphabet[string(x[0])]]
		B[i] = make([]int, len(x)+1)
	}

	for j := 2; j < len(s[0]); j++ {
		for i := 0; i < len(s); i++ {
			// consider the emission probability from this state
			pEM := EMM[i][alphabet[string(x[j-1])]]
			// consider all states in the previous column
			pSTs := make([]float64, len(s))
			for k := range pSTs {
				pSTs[k] = STM[k][i] * s[k][j-1]
			}
			// record the max transition probability and the column number
			p, maxi := max(pSTs)
			s[i][j] = p * pEM

			B[i][j] = maxi
		}
	}

	// find out the max probability from the last column:
	maxP := s[0][len(s[0])-1]
	maxi := 0
	for i := 0; i < len(s); i++ {
		if s[i][len(s[0])-1] > maxP {
			maxP = s[i][len(s[0])-1]
			maxi = i
		}
	}

	// PrintFloatMatrix(s)
	// PrintIntMatrix(B)

	return s, B, maxi
}

func Backtrack(startRowIndex int, B [][]int, states []string) {
	output := ""
	currentRowIndex := startRowIndex
	for j := len(B[0]) - 1; j > 0; j-- {
		output = states[currentRowIndex] + output
		nextRowIndex := B[currentRowIndex][j]
		currentRowIndex = nextRowIndex
	}
	fmt.Println(output)
}

func main() {
	input := "input.txt"
	x, alphabet, states, STM, EMM := ReadInput(input)
	fmt.Println(x)
	fmt.Println(alphabet)
	fmt.Println(states)
	PrintFloatMatrix(STM)
	PrintFloatMatrix(EMM)
	_, B, r := ViterbiDP(x, states, alphabet, STM, EMM)
	Backtrack(r, B, states)
}
