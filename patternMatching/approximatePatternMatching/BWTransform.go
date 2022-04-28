/*
Burrows-Wheeler Transform Construction Problem: Construct the Burrows-Wheeler transform of a string.

Input: A string Text.
Output: BWT(Text).
*/

package main

// returns the BWT(Text) given a SuffixArray in its string form
func GetBWTransform(sorted []string) string {
	bwt := ""
	for _, s := range sorted {
		l := len(s)
		bwt += string(s[l-1])
	}
	return bwt
}

// func main() {
// 	input := "inputg.txt"
// 	sufficies := ReadInput(input)
// 	sorted := SortSufficies(sufficies)
// 	bwt := GetBWTransform(sorted)
// 	fmt.Println(bwt)
// }
