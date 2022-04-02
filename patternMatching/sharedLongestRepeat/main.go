/*
Longest Shared Substring Problem: Find the longest substring shared by two strings.

Input: Strings Text1 and Text2.
Output: The longest substring that occurs in both Text1 and Text2.

Code Challenge: Solve the Longest Shared Substring Problem. (Multiple solutions may exist, in which case you may return any one.)
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

// returns the Patterns (sufficies) of Text1 and Text2
// !!!Make sure Text1 and Text2 end in $ adn # respectively
func ReadInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")

	Text1 := lines[0]
	Text2 := lines[1]
	Patterns := make([]string, 0)
	for i := range Text1 {
		Patterns = append(Patterns, Text1[i:])
	}
	for i := range Text2 {
		Patterns = append(Patterns, Text2[i:])
	}

	return Patterns
}

func main() {
	input := "input1.txt"
	Patterns := ReadInput(input)
	Trie := TrieConstruction(Patterns)
	Tree := SuffixTreeConstruction(Trie)

	LeafParents := FindLeafParents(Tree)
	longestS := GetLongestRepeat(LeafParents, Tree[0])
	fmt.Println(longestS)
}
