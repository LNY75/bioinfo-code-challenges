package main

import "fmt"

type Node struct {
	id   string
	next map[*Node]int // node -> edge weight
	pred *Node         // predecessor
}

// check if the node with the id is already seen
func NodeExists(id string, nodes map[string]*Node) bool {
	if nodes[id] != nil {
		return true
	}
	return false
}

func PrintMap(nodes map[string]*Node) {
	for k, v := range nodes {
		fmt.Println(k, ": ")
		fmt.Println("next: ")
		for n, e := range v.next {
			fmt.Println(n.id, e)
		}
		fmt.Println("----")
	}
}

func PrintNode(n *Node) {
	fmt.Println(n.id, ": ")
	fmt.Println("next: ")
	for node, edge := range n.next {
		fmt.Println(node.id, edge)
	}
	fmt.Println("  ")
}

// returns true if the queue has n, false otherwise
func IsNodeInQueue(queue []*Node, n *Node) bool {
	for _, v := range queue {
		if v == n {
			return true
		}
	}
	return false
}

func PrintPathLens(pathLens map[*Node]int) {
	for n, i := range pathLens {
		fmt.Println(n.id, i)
	}
}
