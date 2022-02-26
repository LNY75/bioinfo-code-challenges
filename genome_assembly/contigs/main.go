package main

import "fmt"

// return a list of non 1-in-1-out nodes
func FindNodes(adjList map[*Node][]*Node) map[*Node]int {
	nodes := make(map[*Node]int) // a node is mapped to 1 if it is a 1-in-1-out node; otherwise it is mapped to 0

	outMap := make(map[*Node]int) // maps a node to the number of its outgoing edges
	inMap := make(map[*Node]int)  // maps a node to the number of its incoming edges
	for node, neighbors := range adjList {
		outMap[node] = len(neighbors)
		for _, neighbor := range neighbors {
			inMap[neighbor]++
		}
	}

	for node := range adjList {
		if inMap[node] == 1 && outMap[node] == 1 {
			nodes[node] = 1
		} else {
			nodes[node] = 0
		}
	}
	// fmt.Println("traverse the following nodes: ")
	// for _, n := range nodes {
	// 	fmt.Println(n.id)
	// }

	return nodes
}

// find all non branching paths starting from the start node
func FindNonBranchingPaths(start *Node, nodes map[*Node]int) []*LinkedList {
	paths := make([]*LinkedList, 0)

	path := &LinkedList{root: &ListNode{node: start}}
	for _, w := range start.next {
		path.Extend(&ListNode{node: w})
		u := w
		for nodes[u] == 1 && len(u.next) > 0 {
			path.Extend(&ListNode{node: u.next[0]})
			u = u.next[0]
		}
		paths = append(paths, path)
		path = &LinkedList{root: &ListNode{node: start}}
	}

	return paths
}

// find all branching paths (contigs) from the graph
func Contigs(nodes map[*Node]int) []string {
	paths2D := make([][]*LinkedList, 0)
	for n, m := range nodes {
		if m == 0 {
			ps := FindNonBranchingPaths(n, nodes)
			paths2D = append(paths2D, ps)
		}
	}
	paths := make([]*LinkedList, 0)
	for _, ps := range paths2D {
		for _, p := range ps {
			paths = append(paths, p)
			// TraverseCycle(p)
		}
	}

	strPaths := make([]string, 0)
	for _, p := range paths {
		pStr := LinkedListToStr(p)
		strPaths = append(strPaths, pStr)
	}
	return strPaths
}

func main() {
	input := "input2.txt"
	seqs := ReadDeBruijnInput(input)
	dbMap := BuildDeBruijn(seqs)
	// convert dbMap to adj. list
	dbOutput := "deBruijnOutput.txt"
	OutputDeBruijnMap(dbMap, dbOutput)

	adjList := ReadAdjListInput(dbOutput)
	nodes := FindNodes(adjList)

	// var randomStartNode *Node
	// for k, v := range nodes {
	// 	if v == 0 {
	// 		randomStartNode = k
	// 		break
	// 	}
	// }
	// paths := FindNonBranchingPaths(randomStartNode, nodes)
	// for _, p := range paths {
	// 	TraverseCycle(p)
	// }

	contigs := Contigs(nodes)
	// print all contigs
	for _, c := range contigs {
		fmt.Print(c, " ")
	}

}
