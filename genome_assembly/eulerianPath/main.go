package main

import "fmt"

// given an adjacency list, find the unbalanced node
// returns the start and end
func FindEndPoints(adjList map[*Node][]*Node) (*Node, *Node) {
	// a balanced node is mapped to zero, since # in_edges = # out_edges
	// node -> 1: this node is the end
	// node -> -1: this node is the start
	balanceMap := make(map[*Node]int)
	var start *Node
	var end *Node

	// initialize balance map to have the same number of keys (nodes) as the adjList
	for n := range adjList {
		balanceMap[n] = 0
	}

	// cound in and out edges
	for n, neighbors := range adjList {
		// outgoing edges count as -1
		balanceMap[n] -= len(neighbors)
		for _, neighbor := range neighbors {
			// incoming edges count as +1
			balanceMap[neighbor]++
		}
	}

	// find unbalanced nodes:
	for k, v := range balanceMap {
		if v == -1 {
			start = k
		}
		if v == 1 {
			end = k
		}
	}

	fmt.Println(start.id, end.id)
	return start, end
}

// adds an edge from end to start
func ConnectEndToStart(start, end *Node) {
	if end.next == nil {
		end.next = make([]*Node, 1)
		end.next[0] = start
	} else {
		end.next = append(end.next, start)
	}
}

// find the eulerian path
func EulerianPath(start, end *Node) *LinkedList {
	ConnectEndToStart(start, end)
	cycle := EulerianCycle(start)

	NewListEnd := LocateNewListEnd(cycle, start, end)
	path := LinkedList{root: NewListEnd.next}

	// set the end
	NewListEnd.next = nil
	// the end of path should be connected to the second node in cycle
	previousEnd := path.GetLast()
	previousEnd.next = cycle.root.next

	TraverseCycle(&path)

	return &path
}

// in a linked list, find where the start and end are connected; they should not be connected in a eulerian path
// returns the ListNode that contains the end node
func LocateNewListEnd(cycle *LinkedList, start, end *Node) *ListNode {
	current := cycle.root
	var NewListEnd *ListNode
	for current.next != nil {
		if current.node == end && current.next.node == start {
			NewListEnd = current
		}
		current = current.next
	}
	if NewListEnd == nil {
		panic("did not find end node")
	}
	return NewListEnd
}

func main() {
	input := "input2.txt"
	adjList := ReadInput(input)
	start, end := FindEndPoints(adjList)
	EulerianPath(start, end)
}
