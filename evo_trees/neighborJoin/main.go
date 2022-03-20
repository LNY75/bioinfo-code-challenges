/*
Code Challenge: Implement NeighborJoining.

Input: An integer n, followed by an n x n distance matrix.
Output: An adjacency list for the tree resulting from applying the neighbor-joining algorithm. Edge-weights should be accurate to two decimal places (they are provided to three decimal places in the sample output below).

Note on formatting: The adjacency list must have consecutive integer node labels starting from 0. The n leaves must be labeled 0, 1, ..., n - 1 in order of their appearance in the distance matrix. Labels for internal nodes may be labeled in any order but must start from n and increase consecutively.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns n, j and the distance matrix
func ReadInput(input string) (int, [][]float64) {
	D := make([][]float64, 0)
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	n, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("cannot convert string to int")
	}

	for i := 1; i < len(lines); i++ {
		l := strings.Fields(lines[i])
		Drow := make([]float64, len(l))
		for i, s := range l {
			d, err := strconv.ParseFloat(s, 64)
			if err != nil {
				panic("cannot convert string into int d")
			}
			Drow[i] = d
		}
		D = append(D, Drow)
	}

	return n, D
}

func PrintFloatMatrix(s [][]float64) {
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
func PrintAdjList(T map[int][][]float64) {
	for k, v := range T {
		for i := range v {
			fmt.Printf("%d->%.0f:%.3f", k, v[i][0], v[i][1])
			fmt.Println()
		}
	}
}

// remove the element at i, j from list L
func removeElements(L []int, i, j int) []int {
	if len(L) == 2 {
		return make([]int, 0)
	}
	if i > j {
		L = append(L[:i], L[i+1:]...)
		L = append(L[:j], L[j+1:]...)
	} else {
		L = append(L[:j], L[j+1:]...)
		L = append(L[:i], L[i+1:]...)
	}
	return L
}

// remove the rows i, j from D
func removeRows(D [][]float64, i, j int) [][]float64 {
	if len(D) <= 2 {
		return make([][]float64, 0)
	}

	if i > j {
		D = append(D[:i], D[i+1:]...)
		D = append(D[:j], D[j+1:]...)
	} else {
		D = append(D[:j], D[j+1:]...)
		D = append(D[:i], D[i+1:]...)
	}
	return D
}

// remove the col i from D
func removeCols(D [][]float64, i, j int) [][]float64 {
	if len(D[0]) <= 2 {
		return make([][]float64, 0)
	}

	if i > j {
		for k := range D {
			D[k] = append(D[k][:i], D[k][i+1:]...)
		}
		for k := range D {
			D[k] = append(D[k][:j], D[k][j+1:]...)
		}
	} else {
		for k := range D {
			D[k] = append(D[k][:j], D[k][j+1:]...)
		}
		for k := range D {
			D[k] = append(D[k][:i], D[k][i+1:]...)
		}
	}

	return D
}

func NeighborJoin(n int, D [][]float64) map[int][][]float64 {
	T := make(map[int][][]float64)

	nodes := make([]int, n) // corresponds to the row and column of D that represent what nodes are left in D
	for i := range nodes {
		nodes[i] = i
	}
	currentNumberOfNodes := n // keep track of the id of the last node added into the tree T

	for len(D) > 2 {

		// construct Dstar from D
		Dstar := make([][]float64, len(D))
		for i := range Dstar {
			Dstar[i] = make([]float64, len(D[i]))
		}
		for i := range Dstar {
			for j := range Dstar[i] {
				if i != j {
					Dstar[i][j] = float64(len(D)-2)*D[i][j] - TotalDistance(D, i) - TotalDistance(D, j)
				}
			}
		}

		// find I, J such that Dstar[I][J] is a minimum non-diagonal element element of Dstar
		I := 0
		J := 1
		for i := 0; i < len(Dstar); i++ {
			for j := 0; j < len(Dstar[0]); j++ {
				if j != i {
					if Dstar[i][j] < Dstar[I][J] {
						I = i
						J = j
					}
				}
			}
		}

		// get delta: (TotalDistance(i)-TotalDistance(j)) / (n-2)
		delta := (TotalDistance(D, I) - TotalDistance(D, J)) / (float64(len(D)) - 2)
		// get limb lengths for I, J
		limbI := 0.5 * (D[I][J] + delta)
		limbJ := 0.5 * (D[I][J] - delta)

		PrintFloatMatrix(Dstar)
		fmt.Println(I, J, limbI, limbJ)

		// and new row and column m into D s.t. D_k,m = D_m,k = (1/2)(D_k,i + D_k,j - D_i,j) for any k;
		// That is, elements in other rows and columns needs to be changed according to m as well

		// add new row and column m to D
		newRowM := make([]float64, 0)
		for k := range D {
			if k != I && k != J {
				d := 0.5 * (D[k][I] + D[k][J] - D[I][J])
				newRowM = append(newRowM, d)
			}
		}
		newRowM = append(newRowM, 0.0)
		nodem := currentNumberOfNodes
		nodes = append(nodes, currentNumberOfNodes)
		currentNumberOfNodes++

		// add two new limbs (connecting node m with leaves i and j) to the tree T
		// assign length limbLengthi to Limb(i)
		// assign length limbLengthj to Limb(j)
		nodei := nodes[I]
		nodej := nodes[J]
		if T[nodei] == nil {
			T[nodei] = make([][]float64, 0)
		}
		if T[nodej] == nil {
			T[nodej] = make([][]float64, 0)
		}
		if T[nodem] == nil {
			T[nodem] = make([][]float64, 0)
		}
		T[nodei] = append(T[nodei], []float64{float64(nodem), limbI})
		T[nodej] = append(T[nodej], []float64{float64(nodem), limbJ})
		T[nodem] = append(T[nodem], []float64{float64(nodei), limbI})
		T[nodem] = append(T[nodem], []float64{float64(nodej), limbJ})

		fmt.Println("the tree: ")
		PrintAdjList(T)
		fmt.Println("=======================")

		// remove rows and columns of I, J from D
		D = removeRows(D, I, J)
		D = removeCols(D, I, J)
		nodes = removeElements(nodes, I, J)

		for i := range D {
			D[i] = append(D[i], newRowM[i])
		}
		D = append(D, newRowM)

		PrintFloatMatrix(D)
		fmt.Println(nodes)
	}

	// add the edge between node D[0] and D[1]
	i := nodes[0]
	j := nodes[1]
	if T[i] == nil {
		T[i] = make([][]float64, 0)
	}
	if T[j] == nil {
		T[j] = make([][]float64, 0)
	}
	T[i] = append(T[i], []float64{float64(j), D[0][1]})
	T[j] = append(T[j], []float64{float64(i), D[0][1]})

	fmt.Println("the final tree: ")
	PrintAdjList(T)

	return T
}

// returns the total distance from i to all other nodes according to the distance matrix D
func TotalDistance(D [][]float64, i int) float64 {
	sum := 0.0
	for j := range D[i] {
		sum += D[i][j]
	}
	return sum
}

func main() {
	input := "input1.txt"
	n, D := ReadInput(input)
	NeighborJoin(n, D)
}
