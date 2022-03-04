/*
Input: A permutation P.
Output: The sequence of permutations corresponding to applying GreedySorting to P, ending with the identity permutation.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the input permutation as an array of int
func ReadInput(input string) []int {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	pStr := strings.Fields(lines[0])
	p := make([]int, len(pStr))
	for i, s := range pStr {
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

// perform greedy sorting on the permutation
func GreedySorting(p []int) []int {
	steps := 0
	for i := range p {
		for p[i] != i+1 {
			if p[i] == -i-1 {
				p = Reverse(p, i)
				steps++
			} else {
				// find p[j] = i
				for j := i; j < len(p); j++ {
					if p[j] == i+1 || p[j] == -i-1 {
						p = ReverseSubList(p, i, j)
						steps++
						break
					}
				}
			}
		}
	}
	fmt.Println(steps)
	return p
}

// reverse swap the sublist in p from i to j (inclusive)
func ReverseSubList(p []int, i, j int) []int {
	for k := i; k <= (i+j)/2; k++ {
		p = ReverseSwap(p, k, j-(k-i))
	}
	PrintPermutation(p)
	return p
}

// Reverse swap two elements in the permutation p at index i and j
func ReverseSwap(p []int, i, j int) []int {
	tmp := p[i]
	p[i] = -p[j]
	p[j] = -tmp
	return p
}

// reverse the sign of the i-th element in permutaion p
func Reverse(p []int, i int) []int {
	p[i] = -p[i]
	PrintPermutation(p)
	return p
}

func main() {
	input := "input.txt"
	p := ReadInput(input)
	// PrintPermutation(p)
	// ReverseSwap(p, 1, 2)
	// Reverse(p, 3)
	GreedySorting(p)
}
