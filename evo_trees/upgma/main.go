/*
Code Challenge: Implement UPGMA.

Input: An integer n followed by a space separated n x n distance matrix.
Output: An adjacency list for the ultrametric tree returned by UPGMA. Edge weights should be accurate to two decimal places (answers in the sample dataset below are provided to three decimal places).
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

func UPGMA(D [][]float64, n int) map[int][][]float64 {
	clusters := make([]int, n) // the id of each cluster
	for i := range clusters {
		clusters[i] = i
	}
	clusterLens := make([]int, n) // the length of each cluster
	for i := range clusterLens {
		clusterLens[i] = 1
	}

	Age := make([]float64, n)

	T := make(map[int][][]float64, n)
	for i := 0; i < n; i++ {
		T[i] = make([][]float64, 0)
	}

	for len(clusters) > 1 {
		// find closest clusters ci and cj
		minDist := D[0][1]
		I := 0
		J := 1
		for i := 0; i < len(D); i++ {
			for j := i; j < len(D[i]); j++ {
				if D[i][j] != 0 && D[i][j] < minDist {
					minDist = D[i][j]
					I = i
					J = j
				}
			}
		}

		// merge Ci and Cj into a new cluster Cnew, with |Ci| + |Cj| elements
		Cnew := clusters[len(clusters)-1] + 1
		// add this Cnew to the tree T, and connect Cnew with Ci and Cj
		if T[Cnew] == nil {
			T[Cnew] = make([][]float64, 0)
		}
		T[Cnew] = append(T[Cnew], []float64{float64(clusters[I]), 0})
		T[Cnew] = append(T[Cnew], []float64{float64(clusters[J]), 0})
		T[clusters[I]] = append(T[clusters[I]], []float64{float64(Cnew), 0})
		T[clusters[J]] = append(T[clusters[J]], []float64{float64(Cnew), 0})

		// Age(Cnew) = D_ij/2
		Age = append(Age, 0)
		Age[Cnew] = float64(D[I][J]) / 2

		// add col and column to D by computing D(Cnew, C) for each C in Clusters
		// the distance between Cnew and another cluster Cm is equal to (DCi,Cm ·|Ci|+DCj,Cm ·|Cj|) / (|Ci|+|Cj|).
		newRow := make([]float64, 0)
		for i := range D {
			if i != I && i != J {
				Ilen := float64(clusterLens[I])
				Jlen := float64(clusterLens[J])
				d := (D[i][I]*Ilen + D[i][J]*Jlen) / (Ilen + Jlen)
				newRow = append(newRow, d)
			}
		}
		newRow = append(newRow, 0)

		// add Cnew to clusters and clusterLens
		clusters = append(clusters, Cnew)
		clusterLens = append(clusterLens, clusterLens[I]+clusterLens[J])

		// remove rows and columns of I and J
		D = removeRows(D, I, J)
		if len(D) > 0 {
			D = removeCols(D, I, J)
		}

		// remove Ci and Cj from clusters and clusterLens
		clusters = removeElements(clusters, I, J)
		clusterLens = removeElements(clusterLens, I, J)

		// append new row and column to D
		for i := range D {
			D[i] = append(D[i], newRow[i])
		}
		D = append(D, newRow)

		// PrintFloatMatrix(D)
		// fmt.Println("----------------")
	}

	// add edge values
	for k := range T {
		for _, l := range T[k] {
			w := k
			v := l[0]
			l[1] = math.Abs(Age[w] - Age[int(v)])
		}
	}
	// fmt.Println(T)

	return T
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

func main() {
	input := "input1.txt"
	n, D := ReadInput(input)
	fmt.Println(n)
	PrintFloatMatrix(D)
	fmt.Println("------------")

	T := UPGMA(D, n)
	PrintAdjList(T)
}
