/*
Code Challenge: Solve the Trie Construction Problem.

Input: A space-separated collection of strings Patterns.

Output: The adjacency list corresponding to Trie(Patterns), in the following format. If Trie(Patterns) has n nodes, first label the root with 0 and then label the remaining nodes with the integers 1 through n - 1 in any order you like. Each edge of the adjacency list of Trie(Patterns) will be encoded by a triple: the first two members of the triple must be the integers labeling the initial and terminal nodes of the edge, respectively; the third member of the triple must be the symbol labeling the edge.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	id         int
	edgeSymbol byte
	parent     *Node
	next       []*Node
}

func ReadInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	if err != nil {
		panic("cannot read n")
	}
	Patterns := strings.Fields(lines[0])

	return Patterns
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
			currentSymbol := Pattern[i]
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

func PrintTrie(Trie []*Node) {
	for _, node := range Trie {
		// fmt.Println(node.id, string(node.edgeSymbol))
		// fmt.Print("next: ")
		// for _, next := range node.next {
		// 	fmt.Print(next.id, ", ")
		// }
		// fmt.Println()
		// if node.parent != nil {
		// 	fmt.Println("parent: ", node.parent.id)
		// }
		for _, next := range node.next {
			fmt.Println(node.id, next.id, string(next.edgeSymbol))
		}
	}
}

func main() {
	input := "input1.txt"
	Patterns := ReadInput(input)
	Trie := TrieConstruction(Patterns)
	PrintTrie(Trie)
}
