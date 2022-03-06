/*
Code Challenge: Solve the 2-Break Distance Problem.

Input: Genomes P and Q.
Output: The 2-break distance d(P, Q).
*/
package main

import "fmt"

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

// returns the cycles formed by the input edges E
func GetCycles(RedE, BlueE [][]int) [][]int {
	RE := CopyEdges(RedE)
	BE := CopyEdges(BlueE)

	cycles := make([][]int, 0)
	for len(RE) > 0 && len(BE) > 0 {

		// find one cycle
		cycle := make([]int, 0)
		cycle = append(cycle, BE[0][0], BE[0][1])
		currentNode := cycle[1]
		BE = remove(BE, 0)

		for cycle[len(cycle)-1] != cycle[0] {
			for i, e := range RE {
				// check whether either side of the edge matches with the current node
				if e[0] == currentNode {
					cycle = append(cycle, e[1])
					currentNode = e[1]
					RE = remove(RE, i)
					break
				} else if e[1] == currentNode {
					cycle = append(cycle, e[0])
					currentNode = e[0]
					RE = remove(RE, i)
					break
				}
			}

			for i, e := range BE {
				// check whether either side of the edge matches with the current node
				if e[0] == currentNode {
					cycle = append(cycle, e[1])
					currentNode = e[1]
					BE = remove(BE, i)
					break
				} else if e[1] == currentNode {
					cycle = append(cycle, e[0])
					currentNode = e[0]
					BE = remove(BE, i)
					break
				}
			}
		}
		cycles = append(cycles, cycle)
	}
	return cycles
}

/*
// Input: The colored edges ColoredEdges of a genome graph. (consider only the red and black edges: RE & BKE)
// Output: The genome P corresponding to this genome graph.
*/
func GraphToGenome(RedE, BlackE [][]int) [][]int {
	RE := CopyEdges(RedE)
	BKE := CopyEdges(BlackE)

	// find all cycles
	// for each cycle, run CycleToChromosome

	cycles := make([][]int, 0)
	for len(RE) > 0 && len(BKE) > 0 {

		// find one cycle
		cycle := make([]int, 0)
		cycle = append(cycle, BKE[0][0], BKE[0][1])
		currentNode := cycle[1]
		BKE = remove(BKE, 0)

		for cycle[len(cycle)-1] != cycle[0] {
			for i, e := range RE {
				// check whether either side of the edge matches with the current node
				if e[0] == currentNode {
					cycle = append(cycle, e[1])
					currentNode = e[1]
					RE = remove(RE, i)
					break
				} else if e[1] == currentNode {
					cycle = append(cycle, e[0])
					currentNode = e[0]
					RE = remove(RE, i)
					break
				}
			}

			for i, e := range BKE {
				// check whether either side of the edge matches with the current node
				if e[0] == currentNode {
					cycle = append(cycle, e[1])
					currentNode = e[1]
					BKE = remove(BKE, i)
					break
				} else if e[1] == currentNode {
					cycle = append(cycle, e[0])
					currentNode = e[0]
					BKE = remove(BKE, i)
					break
				}
			}
		}
		cycles = append(cycles, cycle)
	}

	chromosomes := make([][]int, 0)
	for _, cycle := range cycles {
		chromosomes = append(chromosomes, CycleToChromosome(cycle))
	}

	return chromosomes
}

func CopyEdges(E [][]int) [][]int {
	copy := make([][]int, len(E))
	for i := range E {
		copy[i] = make([]int, len(E[i]))
		for j := range E[i] {
			copy[i][j] = E[i][j]
		}
	}
	return copy
}

/*
input: a set of edges
output: a sequence of nodes
*/
func EdgesToNodes(E [][]int) []int {
	e := make([]int, 0)
	for _, edge := range E {
		for _, node := range edge {
			e = append(e, node)
		}
	}
	return e
}

/*
Input: A sequence Nodes of integers between 1 and 2n
Output: The chromosome Chromosome containing n synteny blocks resulting from applying CycleToChromosome to Nodes.
*/
func CycleToChromosome(n []int) []int {
	chromosome := make([]int, len(n)/2)

	for i := 0; i < len(chromosome); i++ {
		// compare 2i and 2i+1
		if n[2*i+1] > n[2*i] {
			chromosome[i] = n[2*i+1] / 2
		} else {
			chromosome[i] = -n[2*i] / 2
		}
	}
	return chromosome
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

// returns the cycles in P, Q
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

func PrintGenome(P [][]int) {
	for i := range P {
		fmt.Print("(")
		for j := range P[i] {
			if P[i][j] > 0 {
				fmt.Print("+", P[i][j])
			} else {
				fmt.Print(P[i][j])
			}
			if j != len(P[i])-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print(")")
	}
	fmt.Println()
}

// func main() {
// 	input := "input.txt"
// 	P, Q := ReadInput(input)
// 	// fmt.Println(P)
// 	// fmt.Println(Q)

// 	// nodesP := ChromosomeToCycle(P[0])
// 	// nodesQ1 := ChromosomeToCycle(Q[0])
// 	// nodesQ2 := ChromosomeToCycle(Q[1])
// 	// fmt.Println(ColoredEdges(P))
// 	// fmt.Println(ColoredEdges(Q))

// 	cycles := TwoBreakDist(P, Q)

// 	fmt.Println(NumBlocks(P))
// 	fmt.Println(len(cycles))

// 	fmt.Println(NumBlocks(P) - len(cycles))
// }
