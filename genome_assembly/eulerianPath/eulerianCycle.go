package main

import (
	"os"
	"strings"
)

// extract the adjacency list
func ReadInput(input string) map[*Node][]*Node {
	var adjList map[*Node][]*Node = make(map[*Node][]*Node)
	var nodes map[string]*Node = make(map[string]*Node) // used for finding a Node object from its id (string)

	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	for i := range lines {
		// reg, err := regexp.Compile("[^0-9]+")
		// if err != nil {
		// 	panic(err)
		// }

		splitFromArrow := strings.Split(lines[i], " -> ")
		splitFromComma := strings.Split(splitFromArrow[1], ",")

		// build adjacency list
		sourceStr := string(splitFromArrow[0])
		var source *Node
		// check if this node is already in the adjacency list
		if nodes[sourceStr] == nil {
			source = &Node{id: sourceStr}
			nodes[sourceStr] = source
		} else {
			source = nodes[sourceStr]
		}

		dests := make([]*Node, 0)
		for j := 0; j < len(splitFromComma); j++ {
			destStr := string(splitFromComma[j])
			// check if we've alreayd seen a node with this id before
			if nodes[destStr] != nil {
				dests = append(dests, nodes[destStr])
			} else {
				dest := Node{id: string(splitFromComma[j])}
				nodes[destStr] = &dest
				adjList[&dest] = make([]*Node, 0)
				dests = append(dests, &dest)
			}
		}
		source.next = dests
		adjList[source] = dests
	}

	return adjList
}

// from a node with unvisited edges, form a cycle
func FormCycle(source *Node) LinkedList {
	var cycle LinkedList
	var root ListNode = ListNode{node: source}
	cycle.root = &root
	// start traversing a cycle
	current := &root
	for true {
		// find next unvisited node from the source
		next := NextUnvisitedNode(current.node)
		if next != nil {
			// remove next from the neighbors of the current node
			current.node.next = current.node.next[1:]
			// make a new ListNode and add it to the cycle
			nextlistNode := &ListNode{node: next}
			current.next = nextlistNode
			current = nextlistNode
		} else {
			break
		}
	}
	return cycle
}

// Find the node in the cycle that has unvisited edges, if any
func GetNodeWithUnvistedEdge(cycle LinkedList) *Node {
	current := cycle.root
	for current != nil {
		if len(current.node.next) > 0 {
			// we've found a node with unvisited edges
			// fmt.Println(current.node.id)
			return current.node
		}
		current = current.next
	}
	return nil
}

func EulerianCycle(source *Node) *LinkedList {
	cycle := FormCycle(source)
	// fmt.Println("c1: ")
	// TraverseCycle(&cycle)

	newStart := GetNodeWithUnvistedEdge(cycle)
	for newStart != nil {
		cycle2 := FormCycle(newStart)
		// fmt.Println("c2: ")
		// TraverseCycle(&cycle2)

		cycle = *MergeCycles(&cycle, &cycle2)
		// fmt.Println("merged: ")
		// TraverseCycle(&cycle)

		newStart = GetNodeWithUnvistedEdge(cycle)
	}
	// TraverseCycle(&cycle)
	return &cycle
}

// func main() {
// 	input := "input2.txt"
// 	adjList := ReadInput(input)
// 	// PrintAdjList(adjList)

// 	startNode := PickNode(adjList)
// 	EulerianCycle(startNode)
// }
