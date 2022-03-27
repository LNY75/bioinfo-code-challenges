/*
Code Challenge: Implement SmallParsimony to solve the Small Parsimony Problem.

Input: An integer n followed by an adjacency list for a rooted binary tree with n leaves labeled by DNA strings.
Output: The minimum parsimony score of this tree, followed by the adjacency list of a tree corresponding to labeling internal nodes by DNA strings in order to minimize the parsimony score of the tree.  You may break ties however you like.

Note: Remember to run SmallParsimony on each individual index of the strings at the leaves of the tree.
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id          int
	seq         string
	child1      *Node
	child2      *Node
	parent      *Node
	lenToParent int
}

func ReadInput(input string) (int, []*Node) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	n, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("cannot read n")
	}
	T := make([]*Node, n+len(lines)/2)
	lines = lines[1:]

	// read leaves first
	// the first n lines contain sequences representing leaves, the rest are internal nodes
	for i := 0; i < n-1; i += 2 {
		adj1 := strings.Split(lines[i], "->")
		adj2 := strings.Split(lines[i+1], "->")
		parentID, err := strconv.Atoi(adj1[0])
		if err != nil {
			panic("cannot convert parent id string to int")
		}
		// make parent node:
		parent := &Node{id: parentID}
		T[parentID] = parent

		// make the children nodes:
		child1 := &Node{}
		child2 := &Node{}

		// child and child2 are leaves
		child1Seq := adj1[1]
		child2Seq := adj2[1]

		child1.seq = child1Seq
		child2.seq = child2Seq
		child1.id = i
		child2.id = i + 1
		T[i] = child1
		T[i+1] = child2

		parent.child1 = child1
		parent.child2 = child2
		child1.parent = parent
		child2.parent = parent
	}

	// read the rest of the internal nodes
	for i := n; i < len(lines)-1; i += 2 {
		adj1 := strings.Split(lines[i], "->")
		adj2 := strings.Split(lines[i+1], "->")
		parentID, err := strconv.Atoi(adj1[0])
		if err != nil {
			panic("cannot convert parent id string to int")
		}
		// make parent node:
		parent := &Node{id: parentID}
		T[parentID] = parent

		// child1 and child2 are internal nodes
		child1ID, err := strconv.Atoi(adj1[1])
		child2ID, err := strconv.Atoi(adj2[1])
		if err != nil {
			panic("cannot convert child1ID to int")
		}

		if T[child1ID] == nil {
			child1 := &Node{}
			child1.id = child1ID
			T[child1ID] = child1
		}
		if T[child2ID] == nil {
			child2 := &Node{}
			child2.id = child2ID
			T[child2ID] = child2
		}

		child1 := T[child1ID]
		child2 := T[child2ID]
		parent.child1 = child1
		parent.child2 = child2
		child1.parent = parent
		child2.parent = parent
	}

	// PrintNodes(T)
	return n, T
}

// I indicates the index of the characters in leaves that concerns this particular SmallParsinomy problem
// e.g. I=0; leaves have sequences: AGT, AGC, CGT, CGC; then the actual characters from all the leaves that this SmallParsinomy takes into consideration are: A, A, C, C
// by default, Character is always ['A', 'C', 'G', 'T']
// returns a matrix representing the s array for each node
func SmallParsimony(n int, T []*Node, Character []byte, I int) [][]int {
	// initialize Tags for every node:
	Tags := make([]bool, len(T))
	s := make([][]int, len(T))
	for i := range s {
		s[i] = make([]int, 4)
	}

	// loop through the leaves
	for i := 0; i < n; i++ {
		leaf := T[i]
		Tags[i] = true
		for k := range Character {
			if Character[k] == leaf.seq[I] {
				s[i][k] = 0
				// e.g. if a leaf with id 0 has leaf.seq[I] is 'A'; then s[0]['A'] = s[0][0] = 0
			} else {
				s[i][k] = math.MaxInt64
			}
		}
	}

	// PrintLeavesS(n, s)

	// while there are ripe nodes in T
	// We call an internal node of T ripe if its tag is 0 but its children’s tags are both 1. SmallParsimony works upward from the leaves, finding a ripe node v at which to compute sk(v) at each step.
	v := FindRipeNode(Tags, n, T)
	for v != nil {
		Tags[v.id] = true
		for k := range Character {
			// sk(v) = s[v.id][k]
			// sk(v) ← minimum_(all symbols i) {si(Daughter(v))+αi,k} + minimum_(all symbols j) {sj(Son(v))+αj,k}
			c1 := v.child1
			c2 := v.child2

			minScore := math.MaxInt64
			for i := range Character {
				for j := range Character {
					if s[c1.id][i] == math.MaxInt64 || s[c2.id][j] == math.MaxInt64 {
						continue
					}

					var alpha_ik int
					var alpha_jk int
					if k == i {
						alpha_ik = 0
					} else {
						alpha_ik = 1
					}
					if k == j {
						alpha_jk = 0
					} else {
						alpha_jk = 1
					}

					score := s[c1.id][i] + alpha_ik + s[c2.id][j] + alpha_jk
					if score < minScore {
						minScore = score
					}
				}
			}
			s[v.id][k] = minScore
		}
		v = FindRipeNode(Tags, n, T)
	}

	// PrintInternalNodeS(n, s)
	return s
}

// a node is ripe if its tag is 0 but its children’s tags are both 1. SmallParsimony works upward from the leaves, finding a ripe node v at which to compute sk(v) at each step.
// returns a pointer to the first-found ripe node; otherwise nil
func FindRipeNode(Tags []bool, n int, T []*Node) *Node {
	// start searching from n, becauase 0 ~ n-1 are leaves
	for i := n; i < len(Tags); i++ {
		node := T[i]
		// check if the node is ripe
		if !Tags[i] && Tags[node.child1.id] && Tags[node.child2.id] {
			return node
		}
	}
	return nil
}

// returns the root of the tree
func GetRoot(T []*Node) *Node {
	for _, node := range T {
		if node.parent == nil {
			return node
		}
	}
	panic("could not find root")
}

// returns an array of characters where each character corresponds to one element in s. The goal is to minimize the sum over all s(k), the second thing it returns is the score of the root
func GetSymbolsAndRootScore(root *Node, s [][]int, T []*Node, Character []byte) ([]int, int) {
	syms := make([]int, len(s))
	// start from the root, traverse the tree in a BFS manner
	queue := make([]*Node, 1)
	queue[0] = root
	// choose the character with smallest score in s[root] as root's symbol
	syms[root.id] = Min(s[root.id])[0]
	rootScore := s[root.id][syms[root.id]]

	for len(queue) > 0 {
		current := queue[0]
		if current.child1.child1 != nil {
			queue = append(queue, current.child1, current.child2)
			syms = SetSymbol(current, syms, s, Character)
		}
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			break
		}
	}

	return syms, rootScore
}

// given a node, determine its children's symbols (and then add to syms) based on s[child1] and s[child2]
func SetSymbol(node *Node, syms []int, s [][]int, Character []byte) []int {
	rootSym := syms[node.id]
	c1 := node.child1
	c2 := node.child2
	// possible choices of symbols for child1:
	possibleIs := Min(s[c1.id])
	// possible choices of symbols for child2:
	possibleJs := Min(s[c2.id])

	for _, i := range possibleIs {
		for _, j := range possibleJs {
			var alpha_ik int
			var alpha_jk int
			if i == rootSym {
				alpha_ik = 0
			} else {
				alpha_ik = 1
			}
			if j == rootSym {
				alpha_jk = 0
			} else {
				alpha_jk = 1
			}

			score := s[c1.id][i] + s[c2.id][j] + alpha_ik + alpha_jk
			if score == s[node.id][rootSym] {
				syms[c1.id] = i
				syms[c2.id] = j
				return syms
			}
		}
	}
	panic("what, cannot find symbols for both child")
}

// return the indicies of all minimum element in sv
func Min(sv []int) []int {
	mins := make([]int, 0)
	min := sv[0]
	for i := range sv {
		if sv[i] < min {
			min = sv[i]
		}
	}

	for i := range sv {
		if sv[i] == min {
			mins = append(mins, i)
		}
	}
	return mins
}

// for each internal node, get its sequence
func GetInternalSeqsAndScore(n int, root *Node, T []*Node, Character []byte) ([]string, int) {
	// store the sequence for each internal node
	internalSeqs := make([]string, len(T))
	// if the length of a sequence at a leaf is l, then we need to run GetSymbolsAndRootScore l times
	l := len(T[0].seq)
	var sumRootScore int
	for i := 0; i < l; i++ {
		s := SmallParsimony(n, T, Character, i)
		syms, rootScore := GetSymbolsAndRootScore(root, s, T, Character)
		sumRootScore += rootScore

		// loops (# of total nodes - n) times; adds one symbol to each internal node sequence
		for j := n; j < len(syms); j++ {
			char := Character[syms[j]]
			internalSeqs[j] += string(char)
		}
	}
	return internalSeqs, sumRootScore
}

// assign sequences to internal nodes
func AssignSeqToNode(internalSeqs []string, T []*Node) []*Node {
	for i, node := range T {
		node.seq += internalSeqs[i]
	}
	return T
}

// assign edge values
func AssignLenToParent(n int, root *Node, internalSeqs []string, T []*Node) []*Node {
	// fill in sequences for the first n elements in internalSeqs
	for i := 0; i < n; i++ {
		internalSeqs[i] = T[i].seq
	}

	queue := make([]*Node, 1)
	queue[0] = root
	for len(queue) > 0 {
		current := queue[0]
		c1 := current.child1
		c2 := current.child2
		if current.child1.child1 != nil {
			queue = append(queue, current.child1, current.child2)
		}
		s1 := internalSeqs[current.id]
		s2 := internalSeqs[c1.id]
		s3 := internalSeqs[c2.id]

		c1.lenToParent = DifferenceScore(s1, s2)
		c2.lenToParent = DifferenceScore(s1, s3)
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			break
		}
	}
	return T
}

// returns the difference score between two sequences
func DifferenceScore(s1, s2 string) int {
	score := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			score++
		}
	}
	return score
}

func main() {
	input := "input2.txt"
	n, T := ReadInput(input)
	Character := []byte{'A', 'C', 'G', 'T'}
	root := GetRoot(T)
	// s0 := SmallParsimony(n, T, Character, 1)
	// GetSymbolsAndRootScore(root, s0, T, Character)
	internalSeqs, sumScore := GetInternalSeqsAndScore(n, root, T, Character)
	fmt.Println(sumScore)
	T = AssignSeqToNode(internalSeqs, T)
	T = AssignLenToParent(n, root, internalSeqs, T)

	// PrintNodes(T)
	PrintAdj(T)
}
