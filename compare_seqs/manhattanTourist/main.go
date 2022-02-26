package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read input to the Down and Right matrices
func ReadInput(input string) (int, int, [][]int, [][]int) {
	// read input
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	nm := strings.Split(lines[0], " ")
	n, err := strconv.Atoi(nm[0])
	m, err := strconv.Atoi(nm[1])
	// dim of Down: n * m+1; Right: n+1 * m
	// Down[i][j] stores the weight of the edge entering node_i,j
	// Right[i][j] sotores the weigt of the edge entering node_i,j
	Down := make([][]int, n)
	Right := make([][]int, n+1)

	// extract the Down matrix
	for i := 1; i < 1+n; i++ {
		row := strings.Split(lines[i], " ")
		Down[i-1] = make([]int, 0)
		for _, str := range row {
			e, err := strconv.Atoi(str)
			if err != nil {
				panic("cannot convert string to integer")
			}
			Down[i-1] = append(Down[i-1], e)
		}
	}

	// extract the Right matrix
	counter := 0
	for i := 2 + n; i < 2*n+3; i++ {
		row := strings.Split(lines[i], " ")
		Right[counter] = make([]int, 0)
		for _, str := range row {
			e, err := strconv.Atoi(str)
			if err != nil {
				panic("cannot convert string to integer")
			}
			Right[counter] = append(Right[counter], e)
		}
		counter++
	}

	return n, m, Down, Right
}

func Print2DMatrix(m [][]int) {
	for _, row := range m {
		for _, e := range row {
			fmt.Print(e, " ")
		}
		fmt.Println(" ")
	}
}

// returns the longest path in the Manhattan Tourist problem
func Manhattan(n, m int, Down, Right [][]int) int {
	// initialize the matrix
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, m+1)
	}
	// fill in max path lengths for the first row
	for i := 1; i < m+1; i++ {
		s[0][i] = s[0][i-1] + Right[0][i-1]
	}
	// fill in max path lengths for the first column
	for i := 1; i < n+1; i++ {
		s[i][0] = s[i-1][0] + Down[i-1][0]
	}
	// fill in max path lengths for all other nodes
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			s[i][j] = Max(s[i-1][j]+Down[i-1][j], s[i][j-1]+Right[i][j-1])
		}
	}
	fmt.Println(s)
	return s[n][m]
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	input := "input1.txt"
	n, m, Down, Right := ReadInput(input)
	// Print2DMatrix(Down)
	// Print2DMatrix(Right)
	l := Manhattan(n, m, Down, Right)
	fmt.Println(l)
}
