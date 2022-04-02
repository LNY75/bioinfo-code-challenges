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
	// while there is a non-branching internal node:
	fnb := GetFirstNonBranchingInternalNode(Trie) // first non-branching internal node
	for fnb != nil {
		// find the path that we want to merge:
		path := make([]*Node, 1)
		path[0] = fnb
		currentNode := fnb
		for len(currentNode.next) == 1 {
			currentNode = currentNode.next[0]
			path = append(path, currentNode)
		}
		// merge nodes in the path
		mergedNode := &Node{id: fnb.id, parent: fnb.parent, next: path[len(path)-1].next}
		Trie = append(Trie, mergedNode)
		// remove fnb from fnb.parent.next
		fnb.parent = RemoveNeighbor(fnb.parent, fnb)
		// add the mergedNode to the neighbors of fnb.parent
		fnb.parent.next = append(fnb.parent.next, mergedNode)

		// obtain the edge symbol of the mergedNode:
		edgeSymbol := ""
		for _, node := range path {
			edgeSymbol += node.edgeSymbol
		}
		mergedNode.edgeSymbol = edgeSymbol

		// remove all node in path from Trie
		for _, node := range path {
			Trie = RemoveNodeFromTrie(Trie, node)
		}

		fnb = GetFirstNonBranchingInternalNode(Trie)

	}
	// PrintTrie(Trie)
	return Trie
}

// find the first non-branching internal node in Trie
func GetFirstNonBranchingInternalNode(Trie []*Node) *Node {
	for _, node := range Trie {
		if len(node.next) == 1 {
			return node
		}
	}
	return nil
}

func RemoveNodeFromTrie(Trie []*Node, node *Node) []*Node {
	for i, n := range Trie {
		if n == node {
			Trie = append(Trie[:i], Trie[i+1:]...)
		}
	}
	return Trie
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

func main() {
	input := "input1.txt"
	_, Pattern := ReadInput(input)
	// fmt.Println(Text, Pattern)
	Trie := TrieConstruction(Pattern)
	Trie = SuffixTreeConstruction(Trie)
	// PrintTrie(Trie)
	PrintEdges(Trie)
}
