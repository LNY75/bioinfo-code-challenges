/*
Longest Repeat Problem: Find the longest repeat in a string.

Input: A string Text.
Output: A longest substring of Text that appears in Text more than once.

Code Challenge: Solve the Longest Repeat Problem. (Multiple solutions may exist, in which case you may return any one.)
*/

package main

import "fmt"

// returns a list of nodes that are parents of all leaves
func FindLeafParents(Tree []*Node) []*Node {
	// store internal nodes in a map as keys so that we don't need to worry about duplicates
	LeafParentsM := make(map[*Node]*Node)
	for _, node := range Tree {
		if len(node.next) == 0 && node.parent.id != 0 {
			// if node is a leaf, add its parent to map, unless the parent is root
			LeafParentsM[node.parent] = node
		}
	}
	LeafParents := make([]*Node, 0)
	for k := range LeafParentsM {
		LeafParents = append(LeafParents, k)
	}

	return LeafParents
}

// returns the string representation of the path from root to node
func FindPathToRoot(node *Node, Tree []*Node) string {
	s := ""
	currentNode := node
	for currentNode != Tree[0] {
		s = currentNode.edgeSymbol + s
		currentNode = currentNode.parent
	}
	return s
}

// !!!Make sure the input string ends in a $
func main() {
	input := "input1.txt"
	_, Patterns := ReadInput(input)
	Trie := TrieConstruction(Patterns)
	Tree := SuffixTreeConstruction(Trie)

	LeafParents := FindLeafParents(Tree)
	longestS := ""
	S := make([]string, 0)
	for _, node := range LeafParents {
		s := FindPathToRoot(node, Tree)
		if len(s) > len(longestS) {
			longestS = s
		}
		S = append(S, s)
	}

	fmt.Println(longestS)
}
