package main

// returns a list of nodes that are parents of all leaves
// for sharedLongestRepeat problem, a valie parent node should contain at least one child with edgeSymbol ending in # and another with edge symbol ending in $
func FindLeafParents(Tree []*Node) []*Node {
	// store internal nodes in a map as keys so that we don't need to worry about duplicates
	LeafParentsM := make(map[*Node]*Node)
	for _, node := range Tree {
		if len(node.next) == 0 && node.parent.id != 0 && LeafParentsM[node.parent] == nil {
			// if node is a leaf, add its parent to map if it is not already in the map; Do not add if the parent is root
			LeafParentsM[node.parent] = node
		}
	}
	LeafParents := make([]*Node, 0)
	for k := range LeafParentsM {
		if ValidParent(k) {
			LeafParents = append(LeafParents, k)
		}
	}
	return LeafParents
}

// a valid parent contains at least on child with edgeSymbol ending in # and another ending in $
func ValidParent(node *Node) bool {
	containsText1Suffix := false
	containsText2Suffix := false
	for _, child := range node.next {
		l := len(child.edgeSymbol) - 1
		if child.edgeSymbol[l] == '$' {
			containsText1Suffix = true
		}
		if child.edgeSymbol[l] == '#' {
			containsText2Suffix = true
		}
	}
	return containsText1Suffix && containsText2Suffix
}

// returns the string representation of the path from root to node
func FindPathToRoot(node *Node, root *Node) string {
	s := ""
	currentNode := node
	for currentNode != root {
		s = currentNode.edgeSymbol + s
		currentNode = currentNode.parent
	}
	return s
}

func GetLongestRepeat(LeafParents []*Node, root *Node) string {
	longestS := ""
	S := make([]string, 0)
	for _, node := range LeafParents {
		s := FindPathToRoot(node, root)
		if len(s) > len(longestS) {
			longestS = s
		}
		S = append(S, s)
	}

	return longestS
}
