package main

import (
	"fmt"
	"os"
)

// output adjList to file
func OutputDeBruijnMap(dbMap map[string][]string, output string) {
	// write to output file
	f, err := os.Create(output)
	if err != nil {
		panic("cannot create output file for the deBruijn graph")
	}
	// do not write an empty line at the end
	counter := 0
	for k, v := range dbMap {
		f.WriteString(k)
		f.WriteString(" -> ")
		for i, n := range v {
			if i == len(v)-1 {
				f.WriteString(n)
			} else {
				f.WriteString(n)
				f.WriteString(",")
			}
		}
		if counter != len(dbMap)-1 {
			f.WriteString("\n")
		}
		counter++
	}
}

// converts a linkedList into string
func LinkedListToStr(path *LinkedList) string {
	current := path.root
	strs := make([]string, 0)
	for current != nil {
		strs = append(strs, current.node.id)
		current = current.next
	}
	str := FindPath(strs)
	fmt.Println(str)
	return str
}

func main() {
	input := "input2.txt"
	seqs := ReadDeBruijnInput(input)
	deBruijnMap := BuildDeBruijn(seqs)
	// output deBruijnMap to file, then use that as input for eulerianpath
	dbOutput := "deBruijnOutput.txt"
	OutputDeBruijnMap(deBruijnMap, dbOutput)

	adjList := ReadAdjListInput(dbOutput)
	start, end := FindEndPoints(adjList)
	path := EulerianPath(start, end)

	// convert linkedlist into string
	LinkedListToStr(path)
}
