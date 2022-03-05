/*
Code Challenge: Solve the 2-Break Distance Problem.

Input: Genomes P and Q.
Output: The 2-break distance d(P, Q).
*/
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// returns the input permutation as an array of int
func ReadInput(input string) ([][]int, [][]int) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	pStr := make([][]string, 0)
	qStr := make([][]string, 0)

	re := regexp.MustCompile(`\+\d+|\-\d+`)

	lp := strings.Split(lines[0], ")(")
	for _, l := range lp {
		p := re.FindAllString(l, -1)
		pStr = append(pStr, p)
	}

	lq := strings.Split(lines[1], ")(")
	for _, l := range lq {
		q := re.FindAllString(l, -1)
		qStr = append(qStr, q)
	}

	// convert the string representation of synteny blocks to numbers
	P := make([][]int, len(pStr))
	Q := make([][]int, len(qStr))

	// +3 -4 -5 +4
	// -2 -1
	for i, p := range pStr {
		P[i] = make([]int, len(p))
		// +3 -4 -5 +4
		for j := range p {
			pjint, err := strconv.Atoi(p[j][1:])
			if err != nil {
				panic("canot convert string to int")
			}
			if p[j][0] == '-' {
				P[i][j] = -pjint
			} else {
				P[i][j] = pjint
			}
		}
	}

	for i, q := range qStr {
		Q[i] = make([]int, len(q))
		// +3 -4 -5 +4
		for j := range q {
			qjint, err := strconv.Atoi(q[j][1:])
			if err != nil {
				panic("canot convert string to int")
			}
			if q[j][0] == '-' {
				Q[i][j] = -qjint
			} else {
				Q[i][j] = qjint
			}
		}
	}

	return P, Q
}

// returns the number of blocks in chromosome P
func NumBlocks(P [][]int) int {
	count := 0
	for i := range P {
		count += len(P[i])
	}
	return count
}

/*
Input: A chromosome Chromosome containing n synteny blocks.
Output: The sequence Nodes of integers between 1 and 2n resulting from applying ChromosomeToCycle to Chromosome.*/
func ChromosomeToCycle(chromosome []int) []int {
	nodes := make([]int, 0)
	for _, block := range chromosome {
		if block < 0 {
			nodes = append(nodes, -block*2, -block*2-1)
		} else {
			nodes = append(nodes, block*2-1, block*2)
		}
	}
	return nodes
}

// // returns a set of edges (one edge is represented by a size-2 array) formed by the chromosome P
func ColoredEdges(P [][]int) [][]int {
	E := make([][]int, 0)
	// for each chromosome in P:
	for _, chromosome := range P {
		nodes := ChromosomeToCycle(chromosome)
		for j := 0; j < len(chromosome)-1; j++ {
			edge := []int{nodes[2*j+1], nodes[2*j+2]}
			E = append(E, edge)
		}
		// add the last edge
		E = append(E, []int{nodes[len(nodes)-1], nodes[0]})
	}
	return E
}

// compute the 2-break distance for chromosome P and Q
func TwoBreakDist(P, Q [][]int) [][]int {
	redEdges := ColoredEdges(P)
	blueEdges := ColoredEdges(Q)
	E := append(redEdges, blueEdges...)

	// fmt.Println(E)
	cycles := make([][]int, 0)
	for len(E) > 0 {
		// find one cycle
		cycle := make([]int, 0)
		cycle = append(cycle, E[0][0], E[0][1])
		currentNode := cycle[1]
		E = remove(E, 0)
		for cycle[len(cycle)-1] != cycle[0] {
			for i, e := range E {
				// fmt.Println(i, currentNode, len(E))
				// check whether either side of the edge matches with the current node
				if e[0] == currentNode {
					cycle = append(cycle, e[1])
					currentNode = e[1]
					E = remove(E, i)
					break
				} else if e[1] == currentNode {
					cycle = append(cycle, e[0])
					currentNode = e[0]
					E = remove(E, i)
					break
				}
			}
		}
		// fmt.Println(cycle)
		// fmt.Println(E)

		cycles = append(cycles, cycle)
	}

	return cycles
}

// find one cycle from the set of edges; returns the cycle and the remaining edges (having removed those edges that make the cycle)
func FindOneCycle(E [][]int) ([]int, [][]int) {
	// start with the first edge of E as the first edge of the cycle
	cycle := make([]int, 0)
	cycle = append(cycle, E[0][0], E[0][1])
	currentNode := cycle[1]
	E = remove(E, 0)
	for cycle[len(cycle)-1] != cycle[0] {
		for i, e := range E {
			// check whether either side of the edge matches with the current node
			if e[0] == currentNode {
				cycle = append(cycle, e[1])
				currentNode = e[1]
				E = remove(E, i)
			} else if e[1] == currentNode {
				cycle = append(cycle, e[0])
				currentNode = e[0]
				E = remove(E, i)
			}
		}
	}
	return cycle, E
}

// remove the element at indedx i from l
func remove(l [][]int, i int) [][]int {
	if i >= len(l) {
		fmt.Println(i, len(l))
		panic("cannot remove becasue index is out of range")
	}
	if i == 0 {
		return l[1:]
	} else if i == len(l)-1 {
		return l[:len(l)-1]
	} else {
		return append(l[:i], l[i+1:]...)
	}
}

func main() {
	input := "input.txt"
	P, Q := ReadInput(input)
	// fmt.Println(P)
	// fmt.Println(Q)

	// nodesP := ChromosomeToCycle(P[0])
	// nodesQ1 := ChromosomeToCycle(Q[0])
	// nodesQ2 := ChromosomeToCycle(Q[1])
	// fmt.Println(ColoredEdges(P))
	// fmt.Println(ColoredEdges(Q))

	cycles := TwoBreakDist(P, Q)

	fmt.Println(NumBlocks(P))
	fmt.Println(len(cycles))

	fmt.Println(NumBlocks(P) - len(cycles))
}
