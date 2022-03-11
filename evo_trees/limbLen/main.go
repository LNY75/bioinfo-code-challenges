/*
Code Challenge: Solve the Limb Length Problem.

Input: An integer n, followed by an integer j between 0 and n - 1, followed by a space-separated additive distance matrix D (whose elements are integers).

Output: The limb length of the leaf in Tree(D) corresponding to row j of this distance matrix (use 0-based indexing).

solution:
 compute LimbLength(j) by finding the minimum value of (D_{i,j} + D_{j,k} - D_{i,k})/2 over all pairs of leaves i and k.
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// returns n, j and the distance matrix
func ReadInput(input string) (int, int, [][]int) {
	D := make([][]int, 0)
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	n, err := strconv.Atoi(lines[0])
	j, err := strconv.Atoi(lines[1])
	if err != nil {
		panic("cannot convert string to int")
	}

	for i := 2; i < len(lines); i++ {
		l := strings.Fields(lines[i])
		Drow := make([]int, len(l))
		for i, s := range l {
			d, err := strconv.Atoi(s)
			if err != nil {
				panic("cannot convert string into int d")
			}
			Drow[i] = d
		}
		D = append(D, Drow)
	}

	return n, j, D
}

func PrintIntMatrix(s [][]int) {
	for _, r := range s {
		for i, e := range r {
			if i != len(r)-1 {
				fmt.Print(e, " ")
			} else {
				fmt.Print(e)
			}
			// fmt.Printf("%6d ", e)
		}
		fmt.Println("")
	}
}

// returns the limb length of the leaf in row j of D
func limbLen(j, n int, D [][]int) float64 {
	var ll float64 = math.MaxFloat64 // short for limb length

	// for all pairs of i and k
	for i := 0; i < n; i++ {
		if i != j {
			for k := i; k < n; k++ {
				if k != j {
					// (D_ij + D_jk - D_ik)/2
					newll := float64((D[j][i] + D[j][k] - D[i][k]) / 2)
					if newll < ll {
						ll = newll
					}
				}
			}
		}
	}

	return ll
}

func main() {
	input := "input1.txt"
	n, j, D := ReadInput(input)
	// fmt.Println(n, j)
	// PrintIntMatrix(D)

	ll := limbLen(j, n, D)
	fmt.Println(ll)
}
