/*
2-Break Sorting Problem: Find a shortest transformation of one genome into another by 2-breaks.

Input: Two genomes with circular chromosomes on the same set of synteny blocks.
Output: The sequence of genomes resulting from applying a shortest sequence of 2-breaks transforming one genome into the other.
*/
package main

import (
	"os"
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

	pstr := lines[0][1 : len(lines[0])-1]
	qstr := lines[1][1 : len(lines[1])-1]

	pstrlist := strings.Fields(pstr)
	qstrlist := strings.Fields(qstr)

	P := make([][]int, 1)
	Q := make([][]int, 1)
	for _, blockStr := range pstrlist {
		block, err := strconv.Atoi(blockStr[1:])
		if err != nil {
			panic("cannot convert string to int")
		}
		if blockStr[0] == '-' {
			block = -block
		}
		P[0] = append(P[0], block)
	}

	for _, blockStr := range qstrlist {
		block, err := strconv.Atoi(blockStr[1:])
		if err != nil {
			panic("cannot convert string to int")
		}
		if blockStr[0] == '-' {
			block = -block
		}
		Q[0] = append(Q[0], block)
	}

	return P, Q
}

func TwoBDSort(P, Q [][]int) {
	// goal: sort P so that it matches with Q
	// general approach to solution:
	// keep track of the red edges of the cycle
	// select a pair of red edges and apply two-break
	// check if the resulting graph has one more cycle
	// terminate when the number of cycles is equal the number of blocks

	RE := ColoredEdges(P)
	BE := ColoredEdges(Q)
	BKE := BlackEdges(Q)

	cycles := TwoBreakDist(P, Q)
	// TBD: two break distance
	TBD := NumBlocks(P) - len(cycles)

	PrintGenome(P)
	// we need TBD number of two-break operations to make P match Q
	for i := 0; i < TBD; i++ {
		var i1, i2 int
		cycles := GetCycles(RE, BE)
		// get the group of red edges to choose from
		for _, cycle := range cycles {
			if len(cycle) > 3 {
				// fmt.Println(cycle)
				redEdgeIndicesOfCycle := GroupEdges(RE, cycle)
				i1 = redEdgeIndicesOfCycle[0]
				i2 = redEdgeIndicesOfCycle[1]
				break
			}
		}
		// perform twobreak operation
		RE = TwoBreakOnGenome(RE, BE, BKE, i1, i2)

		genome := GraphToGenome(RE, BKE)
		PrintGenome(genome)
	}
}

// returns the list of indices of edges in E that form the cycle
func GroupEdges(E [][]int, cycle []int) []int {
	edgeIndices := make([]int, 0)
	for i := 0; i < len(cycle)-1; i++ {
		e := []int{cycle[i], cycle[i+1]}

		edgeIndex := ContainsEdge(E, e)
		if edgeIndex != -1 {
			edgeIndices = append(edgeIndices, edgeIndex)
		}
	}
	return edgeIndices
}

// checks whether E contains e
func ContainsEdge(E [][]int, e []int) int {
	for i, v := range E {
		if (e[0] == v[0] && e[1] == v[1]) || (e[1] == v[0] && e[0] == v[1]) {
			return i
		}
	}
	return -1
}

/*
Input: The red edges of a genome graph GenomeGraph, followed by indices of two edges: i1 i2
Output: The red edges of the genome graph resulting from applying the 2-break operation 2-BreakOnGenomeGraph(GenomeGraph, i1 , i2).
*/
func TwoBreakOnGenome(RedE, BlueE, BlackE [][]int, i1, i2 int) [][]int {
	numCycles := len(GetCycles(RedE, BlueE))

	// 1,2 3,4 -> 1,3 2,4 | 1,4 2,3
	tmp := RedE[i1][1]
	RedE[i1][1] = RedE[i2][0]
	RedE[i2][0] = tmp

	newNumCycles := len(GetCycles(RedE, BlueE))

	if newNumCycles > numCycles {
		// print the new genome  * we need the black edges
		// genome := GraphToGenome(RedE, BlackE)
		// PrintGenome(genome)
		return RedE
	} else {
		// switch back first
		tmp := RedE[i1][1]
		RedE[i1][1] = RedE[i2][0]
		RedE[i2][0] = tmp
		// switch the other way
		tmp = RedE[i1][1]
		RedE[i1][1] = RedE[i2][1]
		RedE[i2][1] = tmp

		// genome := GraphToGenome(RedE, BlackE)
		// PrintGenome(genome)
		return RedE
	}
}

// obtain the black edges
func BlackEdges(P [][]int) [][]int {
	E := make([][]int, 0)
	for i := range P {
		for _, block := range P[i] {
			e := make([]int, 2)
			if block > 0 {
				e[0] = 2*block - 1
				e[1] = 2 * block
			} else {
				block = -block
				e[0] = 2 * block
				e[1] = 2*block - 1
			}
			E = append(E, e)
		}
	}

	return E
}

func main() {
	input := "input1.txt"
	P, Q := ReadInput(input)
	// fmt.Println(P, Q)

	// blueEdges := [][]int{[]int{2, 3}, []int{4, 8}, []int{7, 6}, []int{5, 1}}
	// blueEdges2 := [][]int{[]int{1, 3}, []int{4, 8}, []int{7, 6}, []int{5, 2}}
	// n1 := EdgesToNodes(blueEdges)
	// n2 := EdgesToNodes(blueEdges2)
	// n1 = ReshapeCycle(n1)
	// n2 = ReshapeCycle(n2)
	// chr1 := CycleToChromosome(n1)
	// chr2 := CycleToChromosome(n2)
	// fmt.Println(chr1, chr2)

	// RE := [][]int{[]int{2, 4}, []int{3, 6}, []int{1, 5}, []int{7, 8}}
	// BE := [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}
	// chromosomes := GraphToGenome(RE, BE)
	// fmt.Println(chromosomes)

	// fmt.Println(BlackEdges(Q))

	// RE := [][]int{[]int{2, 4}, []int{3, 6}, []int{1, 8}, []int{5, 7}}
	// BE := [][]int{[]int{2, 3}, []int{4, 8}, []int{1, 5}, []int{6, 7}}
	// BKE := BlackEdges(Q)

	// fmt.Println(TwoBreakOnGenome(RE, BE, BKE, 0, 1))

	TwoBDSort(P, Q)
}
