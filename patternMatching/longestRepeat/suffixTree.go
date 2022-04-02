/*
Code Challenge: Solve the Suffix Tree Construction Problem.

Input: A string Text.
Output: A space-separated list of the edge labels of SuffixTree(Text). You may return these strings in any order.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	id         int
	edgeSymbol string
	parent     *Node
	next       []*Node
}

// returns the Text and the Pattern
func ReadInput(input string) (string, []string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	Text := lines[0]
	Pattern := make([]string, 0)
	for i := range Text {
		Pattern = append(Pattern, Text[i:])
	}

	return Text, Pattern
}

/*
Trie Construction Problem: Construct a trie from a set of patterns.

Input: A collection of strings Patterns.
Output: Trie(Patterns).
*/
func TrieConstruction(Patterns []string) []*Node {
	// the trie is a list of nodes
	Trie := make([]*Node, 1)
	root := &Node{id: 0, next: make([]*Node, 0)}
	Trie[0] = root

	for _, Pattern := range Patterns {
		currentNode := Trie[0]
		for i := range Pattern {
			currentSymbol := string(Pattern[i])
			// if there is an outgoing edge from currentNode with label currentSymbol
			// examine the neighbors of currentNode
			var endingNode *Node
			for _, next := range currentNode.next {
				if next.edgeSymbol == currentSymbol {
					// currentNode is the ending node of this edge
					endingNode = next
					currentNode = endingNode
				}
			}
			if endingNode == nil {
				// add a new node newNode to Trie
				newNode := &Node{id: len(Trie), parent: currentNode, edgeSymbol: currentSymbol, next: make([]*Node, 0)}
				// add a new edge from currentNode to newNode with label currentSymbol
				currentNode.next = append(currentNode.next, newNode)
				Trie = append(Trie, newNode)
				// currentNode ← newNode
				currentNode = newNode
			}
		}
	}

	return Trie
}

// construct a suffix tree from a Trie: merge all non-branching nodes
func SuffixTreeConstruction(Trie []*Node) []*Node {
	root := Trie[0]
	// while there is a non-branching internal node:
	fnb := GetFirstNonBranchingInternalNode(root) // first non-branching internal node
	for fnb != nil {
		// find the path that we want to merge:
		path := make([]*Node, 1)
		path[0] = fnb
		currentNode := fnb
		for len(currentNode.next) == 1 {
			currentNode = currentNode.next[0]
			path = append(path, currentNode)
		}
		fn := path[len(path)-1]
		// merge nodes in the path
		mergedNode := &Node{id: fnb.id, parent: fnb.parent, next: fn.next}
		// remove fnb from fnb.parent.next
		fnb.parent = RemoveNeighbor(fnb.parent, fnb)
		// add the mergedNode to the neighbors of fnb.parent
		fnb.parent.next = append(fnb.parent.next, mergedNode)
		// all of fn's children should consider mergedNode as its new parent
		for _, node := range fn.next {
			node.parent = mergedNode
		}

		// obtain the edge symbol of the mergedNode:
		edgeSymbol := ""
		for _, node := range path {
			edgeSymbol += node.edgeSymbol
		}
		mergedNode.edgeSymbol = edgeSymbol

		fnb = GetFirstNonBranchingInternalNode(root)
	}

	// Add nodes to the Trie
	Trie = AddNodesToArray(root)

	return Trie
}

// in a BFS manner, traverse and add all nodes (including merged ones) to the new array representing Trie
func AddNodesToArray(root *Node) []*Node {
	Trie := make([]*Node, 0)
	queue := make([]*Node, 1)
	queue[0] = root
	for len(queue) >= 1 {
		currentNode := queue[0]
		Trie = append(Trie, currentNode)
		for _, node := range currentNode.next {
			queue = append(queue, node)
		}
		if len(queue) <= 1 {
			break
		}
		queue = queue[1:]
	}
	return Trie
}

// find the first non-branching internal node in Trie in a BFS manner
func GetFirstNonBranchingInternalNode(root *Node) *Node {
	queue := make([]*Node, 1)
	queue[0] = root
	for len(queue) >= 1 {
		currentNode := queue[0]
		if len(currentNode.next) == 1 {
			return currentNode
		}
		for _, node := range currentNode.next {
			queue = append(queue, node)
		}
		if len(queue) <= 1 {
			break
		}
		queue = queue[1:]
	}
	return nil
}

// remove neighbot from parent
func RemoveNeighbor(parent *Node, neighbor *Node) *Node {
	for i, node := range parent.next {
		if node == neighbor {
			parent.next = append(parent.next[:i], parent.next[i+1:]...)
		}
	}
	return parent
}

// print edges in the suffix tree
func PrintEdges(Trie []*Node) {
	for i := 1; i < len(Trie); i++ {
		fmt.Println(Trie[i].edgeSymbol)
	}
}

func PrintTrie(Trie []*Node) {
	for _, node := range Trie {
		for _, next := range node.next {
			fmt.Println(node.id, next.id, string(next.edgeSymbol))
		}
	}
}

func PrintPath(path []*Node) {
	fmt.Println(" ")
	for i := range path {
		fmt.Print(path[i].edgeSymbol)
	}
	fmt.Println(" ")
}
