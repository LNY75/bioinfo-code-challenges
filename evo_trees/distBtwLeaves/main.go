/*
Distances Between Leaves Problem: Compute the distances between leaves in a weighted tree.

Input:  An integer n followed by the adjacency list of a weighted tree with n leaves.
Output: An n x n matrix (di,j), where di,j is the length of the path between leaves i and j.
Code Challenge: Solve the Distances Between Leaves Problem. The tree is given as an adjacency list of a graph whose leaves are integers between 0 and n - 1; the notation a->b:c means that node a is connected to node b by an edge of weight c. The matrix you return should be space-separated.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	e.g.: 4 -> [[0, 11], [1, 2], [5, 4]]
*/
func ReadInput(input string) (int, map[int][][]int) {
	m := make(map[int][][]int)
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	n, err := strconv.Atoi(lines[0])

	for _, line := range lines[1:] {
		l := strings.Split(line, "->")
		node, err := strconv.Atoi(l[0])
		if err != nil {
			panic("cannot convert string to integer")
		}

		edgeStr := strings.Split(l[1], ":")
		neighbor, err := strconv.Atoi(edgeStr[0])
		weight, err := strconv.Atoi(edgeStr[1])
		if err != nil {
			panic("cannot convert string to int")

		}
		edge := []int{neighbor, weight}

		if m[node] == nil {
			m[node] = make([][]int, 0)
		}
		m[node] = append(m[node], edge)
	}
	return n, m
}

/*
	given a leave node (integer), find the distances from it to all other leaves in the tree
*/
func BST(start int, adjL map[int][]int) {

}

func main() {
	input := "input.txt"
	n, m := ReadInput(input)
	fmt.Println(n, m)
}
