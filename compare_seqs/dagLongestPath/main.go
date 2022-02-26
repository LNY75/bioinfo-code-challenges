package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns the source, destination node, and a list of all nodes
func ReadInput(input string) (*Node, *Node, map[string]*Node) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	nodes := make(map[string]*Node)

	for i := 1; i < len(lines); i++ {
		info := strings.Split(lines[i], " ") // v, w, edge weight
		edgeWeight, err := strconv.Atoi(info[2])
		vid := info[0]
		wid := info[1]
		if err != nil {
			panic("cannot convert string to int")
		}

		if !NodeExists(vid, nodes) {
			nodes[vid] = &Node{id: vid, next: make(map[*Node]int)}
		}
		if !NodeExists(wid, nodes) {
			nodes[wid] = &Node{id: wid, next: make(map[*Node]int)}
		}
		v := nodes[vid]
		w := nodes[wid]

		// if multiple edges exist between two nodes, keep only the edge with max weight
		if v.next[w] != 0 {
			if v.next[w] < edgeWeight {
				v.next[w] = edgeWeight
			}
		} else {
			v.next[w] = edgeWeight
		}
	}

	// get source and destination
	sdid := strings.Split(lines[0], " ")
	sid := sdid[0]
	did := sdid[1]
	s := nodes[sid]
	d := nodes[did]

	return s, d, nodes
}

// for each node, find its predecessor in the longest path (null if the longest path does not include it)
func FindPath(s *Node, nodes map[string]*Node) map[*Node]int {
	pathLens := make(map[*Node]int) // node -> longest path length from source to node
	// add nodes into pathLens
	for _, v := range nodes {
		pathLens[v] = 0
	}

	queue := make([]*Node, 1)
	queue[0] = s
	for len(queue) != 0 {
		v := queue[0]
		for w, ew := range v.next {
			queue = append(queue, w)

			if pathLens[w] < pathLens[v]+ew {
				pathLens[w] = pathLens[v] + ew
				w.pred = v
			}

		}
		queue = queue[1:]
	}
	return pathLens
}

func BackTrackPath(s, d *Node, nodes map[string]*Node) {
	path := make([]*Node, 1)
	path[0] = d
	for path[len(path)-1] != s {
		path = append(path, path[len(path)-1].pred)
	}
	// reverse print
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i].id, " ")
	}
	fmt.Println(" ")
}

func main() {
	input := "input.txt"
	s, d, nodes := ReadInput(input)
	// PrintNode(s)
	// fmt.Println("-------------")
	// PrintNode(d)
	// fmt.Println("-------------")
	// PrintMap(nodes)
	// fmt.Println("-------------")

	pathLens := FindPath(s, nodes)
	fmt.Println(pathLens[d])
	// PrintPathLens(pathLens)
	BackTrackPath(s, d, nodes)
}
