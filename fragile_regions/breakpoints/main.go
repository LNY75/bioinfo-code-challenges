/*
Number of Breakpoints Problem: Find the number of breakpoints in a permutation.

Input: A permutation.
Output: The number of breakpoints in this permutation.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns a modified input permutation as an array of int
// modified: 0 and len(p) are respectively added to the beginning and the end of the permutation
func ReadInput(input string) []int {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	pStr := strings.Fields(lines[0])
	p := make([]int, len(pStr)+2)
	p[0] = 0
	p[len(pStr)+1] = len(pStr) + 1
	for i := 1; i <= len(pStr); i++ {
		s := pStr[i-1]
		n, err := strconv.Atoi(string(s[1:]))
		if err != nil {
			panic("cannot convert this string to int")
		}
		if s[0] == '-' {
			p[i] = -n
		} else {
			p[i] = n
		}
	}
	return p
}

func PrintPermutation(p []int) {
	fmt.Print("")
	for i, n := range p {
		if i == len(p)-1 {
			if n > 0 {
				fmt.Print("+", n, "")
			} else {
				fmt.Print(n, "")
			}
			fmt.Println()
		} else {
			if n > 0 {
				fmt.Print("+", n, " ")
			} else {
				fmt.Print(n, " ")
			}
		}
	}
}

// count the number of break points in the permutation
func countBreakPoints(p []int) int {
	c := 0
	for i := 0; i < len(p)-1; i++ {
		if p[i+1]-p[i] != 1 {
			c++
		}
	}
	return c
}

func main() {
	input := "input2.txt"
	p := ReadInput(input)
	// PrintPermutation(p)
	fmt.Println(countBreakPoints(p))
}
