package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the column removal threshold, the alphabet, the alignment strings, and the list of columns that exceeds the column removal threshold
func ReadInput(input string) ([]string, map[string]int, []string, [][]string, []bool) {
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

	// group together consetive columns exceeding crt in GA
	GA := make([][]string, 0)
	flag := false
	GAflags := make([]bool, 0)
	for j := 0; j < len(A[0]); j++ {
		c := 0.0
		col := ""
		for i := 0; i < len(A); i++ {
			col += string(A[i][j])
			if A[i][j] == '-' {
				c++
			}
		}
		if c/float64(len(A)) > crt {
			// crtIndexArr = append(crtIndexArr, j)
			if flag == true {
				// there is a consecutive column exceeding crt
				GA[len(GA)-1] = append(GA[len(GA)-1], col)
			} else {
				GA = append(GA, make([]string, 1))
				GA[len(GA)-1][0] = col
				GAflags = append(GAflags, true)
			}
			flag = true
		} else {
			GA = append(GA, make([]string, 1))
			GA[len(GA)-1][0] = col
			flag = false
			GAflags = append(GAflags, false)
		}
	}

	return alphabetStr, alphabet, A, GA, GAflags
}

// finds the sequence of states followed by each string in alignment A
func StatesStrs(A []string, GA [][]string, GAflags []bool) ([][]string, [][]string, [][]string) {
	states := make([]string, len(A))
	symbols := make([]string, len(A))

	for i := 0; i < len(A); i++ {
		var currentStateIndex int

		if GAflags[0] == true {
			s := ""
			for k := 0; k < len(GA[0]); k++ {
				// 0, k, i
				if GA[0][k][i] != '-' {
					s += string(GA[0][k][i])
				}
			}
			if len(s) != 0 {
				states[i] += "I0 "
				symbols[i] += s + " "
			}
		} else {
			currentStateIndex++
			// M1 or D1
			if GA[0][0][i] == '-' {
				// D1
				states[i] += "D1 "
			} else {
				// M1
				states[i] += "M1 "
				symbols[i] += string(GA[0][0][i]) + " "
			}
		}

		for j := 1; j < len(GAflags); j++ {
			if GAflags[j] == true {
				// either I or nothing
				s := ""
				for k := 0; k < len(GA[j]); k++ {
					// 0, k, i
					if GA[j][k][i] != '-' {
						s += string(GA[j][k][i])
					}
				}
				if len(s) != 0 {
					// I what?
					stateIndexStr := strconv.Itoa(currentStateIndex)
					states[i] += "I" + stateIndexStr + " "
					symbols[i] += s + " "
				}
			} else {
				// the column was not emitted
				currentStateIndex++
				stateIndexStr := strconv.Itoa(currentStateIndex)
				// M1 or D1
				if GA[j][0][i] == '-' {
					// D what?
					states[i] += "D" + stateIndexStr + " "
				} else {
					// M what?
					states[i] += "M" + stateIndexStr + " "
					symbols[i] += string(GA[j][0][i]) + " "
				}
			}
		}
	}

	// convert the string representing states for a string in alignment into a list of states
	statesList := make([][]string, len(A))
	statesListNoD := make([][]string, len(A))
	for i, str := range states {
		fmt.Println(str)
		s := strings.Fields(str)
		statesList[i] = s

		for j := range s {
			if s[j][0] != 'D' {
				statesListNoD[i] = append(statesListNoD[i], s[j])
			}
		}
	}
	symbolsList := make([][]string, len(A))
	for i, str := range symbols {
		// fmt.Println(str)
		s := strings.Fields(str)
		symbolsList[i] = s
	}

	return statesList, statesListNoD, symbolsList
}

// returns the states and emission map for the alignment
func StateAndEmMap(GAflags []bool, statesList [][]string, statesListNoD [][]string, symbolsList [][]string) ([]string, map[string]int, map[string][]string, map[string][]string) {
	states := make([]string, 2)
	states[0] = "S"
	states[1] = "I0"
	// figure out all the states
	// the number of false flags in GAflags = the number of M and D states;
	numRetainedCols := 0
	for _, b := range GAflags {
		if !b {
			numRetainedCols++
		}
	}
	for i := 1; i <= numRetainedCols; i++ {
		I := strconv.Itoa(i)
		states = append(states, "M"+I)
		states = append(states, "D"+I)
		states = append(states, "I"+I)
	}
	states = append(states, "E")
	fmt.Println(states)

	// maps each state to an index in the transition matrix
	stateMap := make(map[string]int)
	for i, s := range states {
		stateMap[s] = i
	}

	// initialize the map that maps each state to a list of states it reached, and the map that maps each state to a list of symbols it emitted
	statesMap := make(map[string][]string)
	symbolsMap := make(map[string][]string)
	for _, s := range states {
		statesMap[s] = make([]string, 0)
		symbolsMap[s] = make([]string, 0)
	}
	// find out what states did the S state transition to
	for _, l := range statesList {
		statesMap["S"] = append(statesMap["S"], l[0])
	}
	for _, l := range statesList {
		for i := 0; i < len(l)-1; i++ {
			statesMap[l[i]] = append(statesMap[l[i]], l[i+1])
		}
		// add the exit state
		i := len(l) - 1
		statesMap[l[i]] = append(statesMap[l[i]], "E")
	}

	// fmt.Println(statesMap)
	// for i := range statesList {
	// 	fmt.Println(statesList[i])
	// }

	for i, l := range statesListNoD {
		for j, s := range l {
			switch s[0] {
			case 'M':
				// emits just one symbol
				symbol := symbolsList[i][j]
				symbolsMap[s] = append(symbolsMap[s], symbol)
			case 'I':
				// might emith multiple symbols
				symbols := symbolsList[i][j]
				for _, symbol := range symbols {
					symbolsMap[s] = append(symbolsMap[s], string(symbol))
				}
			}
		}
	}
	// fmt.Println(symbolsMap)

	return states, stateMap, statesMap, symbolsMap
}

// compute the transition matrix and emission matrix
func TEM(states []string, alphabet map[string]int, alphabetStr []string, stateMap map[string]int, statesMap map[string][]string, symbolsMap map[string][]string) {
	// initialize the transiitoin matrix and emission matrix (TM and EM)
	TM := make([][]float64, len(stateMap))
	EM := make([][]float64, len(stateMap))
	for i := range TM {
		TM[i] = make([]float64, len(stateMap))
		EM[i] = make([]float64, len(alphabet))
	}

	rowDenoms := make([]float64, len(TM))
	for s, nextStates := range statesMap {
		// index of s in TM:
		si := stateMap[s]
		for _, next := range nextStates {
			// index of next in TM:
			nexti := stateMap[next]
			TM[si][nexti]++
			// add 1 to the denominator of this row
			rowDenoms[si]++
		}
	}
	// divide each element in TM by the corresponding row denominator
	for i := range TM {
		for j := range TM[i] {
			if TM[i][j] != 0 {
				TM[i][j] /= rowDenoms[i]
			}
		}
	}

	PrintFloatMatrixWithLabels(TM, states, states)
	fmt.Println("--------")

	// calculate the emission proababilities
	rowDenoms = make([]float64, len(EM))
	for s, symbols := range symbolsMap {
		// index of s in EM:
		si := stateMap[s]
		for _, sym := range symbols {
			// index of symbol in EM:
			symi := alphabet[sym]
			EM[si][symi]++
			// add 1 to the denominator of this row
			rowDenoms[si]++
		}
	}
	// divide each element in TM by the corresponding row denominator
	for i := range EM {
		for j := range EM[i] {
			if EM[i][j] != 0 {
				EM[i][j] /= rowDenoms[i]
			}
		}
	}

	PrintFloatMatrixWithLabels(EM, states, alphabetStr)
}

func main() {
	input := "input.txt"
	alphabetStr, alphabet, A, GA, GAflags := ReadInput(input)
	statesList, statesListNoD, symbolsList := StatesStrs(A, GA, GAflags)
	states, stateMap, statesMap, symbolsMap := StateAndEmMap(GAflags, statesList, statesListNoD, symbolsList)
	TEM(states, alphabet, alphabetStr, stateMap, statesMap, symbolsMap)
}
