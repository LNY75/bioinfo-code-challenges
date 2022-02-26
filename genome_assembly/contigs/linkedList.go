package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	id   string
	next []*Node
}

type LinkedList struct {
	root *ListNode
}

type ListNode struct {
	node *Node
	next *ListNode
}

func (L *LinkedList) Extend(n *ListNode) {
	end := L.GetLast()
	end.next = n
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
	// fmt.Println(str)
	return str
}

// solve the genome path problem
func FindPath(strings []string) string {
	path := strings[0]
	for i := 1; i < len(strings); i++ {
		path += strings[i][len(strings[i])-1 : len(strings[i])]
	}
	return path
}

// extract the adjacency list
func ReadAdjListInput(input string) map[*Node][]*Node {
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

// merge two cycles
func MergeCycles(c1, c2 *LinkedList) *LinkedList {
	mergedCycle := LinkedList{root: c2.root}
	// the new start node is the root of c2
	// find the end of c2
	end2 := c2.GetLast()

	// find which node in c1 has the same node as the end of c2
	sameNode := c1.FindNode(end2.node)
	// the end of c2 is connected to the next node of the same node from c1
	end2.next = sameNode.next
	// fmt.Println("the end of c2 is now connected to: ", end2.next.node.id)
	// the same node from c1 is the end of the merged cycle
	sameNode.next = nil

	// the end of the merged cycle should be connected to the next node of c1's root
	endMerged := mergedCycle.GetLast()
	endMerged.next = c1.root.next

	return &mergedCycle
}

// find the listnode from L that contains node n
func (L *LinkedList) FindNode(n *Node) *ListNode {
	current := L.root
	for current != nil {
		if current.node == n {
			return current
		}
		current = current.next
	}
	return nil
}

// return the last node of the linked list L
func (L *LinkedList) GetLast() *ListNode {
	current := L.root
	for current.next != nil {
		current = current.next
	}
	return current
}

func TraverseCycle(cycle *LinkedList) {
	current := cycle.root
	for current != nil {
		if current.next != nil {
			fmt.Print(current.node.id, "->")
		} else {
			fmt.Print(current.node.id)
		}
		current = current.next
	}
	fmt.Println(" ")
}

// find a unvisited node from the source nodes' neighbors
func NextUnvisitedNode(source *Node) *Node {
	if len(source.next) != 0 {
		return source.next[0]
	} else {
		return nil
	}
}

// pick a node from the adjacency list
func PickNode(adjList map[*Node][]*Node) *Node {
	var node *Node
	for k := range adjList {
		node = k
		break
	}
	return node
}

// print the adjacency list nicely
func PrintAdjList(l map[*Node][]*Node) {
	for k := range l {
		fmt.Print(k.id, " -> ")
		for i := range k.next {
			fmt.Print(l[k][i].id, " ")
		}
		fmt.Println(" ")
	}
}

func PrintEdges(edges map[*Node]int) {
	for k, v := range edges {
		fmt.Println(k.id, " : ", v)
	}
}
