package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read the string, k, profile matrix from input file
func ReadEverything(input string) (string, int, [][]float64) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// line[0] -> string
	// line[1] -> k
	// line[2] and onward -> profile matrix
	// profile matrix: A, C, G, T
	str := lines[0]
	k, err := strconv.Atoi(lines[1])

	profileMatrix := make([][]float64, len(lines)-2)
	for i := 2; i < 6; i++ {
		line := strings.Split(lines[i], " ")
		fmt.Println(line)

		floatArr := make([]float64, len(line))
		for j := 0; j < len(line); j++ {
			n, err := strconv.ParseFloat(line[j], 64)
			if err != nil {
				panic("cannot convert string to float64")
			}
			floatArr[j] = n
		}
		profileMatrix[i-2] = floatArr
	}
	return str, k, profileMatrix
}

func PrintMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], ", ")
		}
		fmt.Println(" ")
	}
}

// given a profile matrix, compute the probability of observing a k-mer
func Prob(kmer string, m [][]float64) float64 {
	prob := 1.0
	for i := 0; i < len(kmer); i++ {
		switch kmer[i] {
		case 'A':
			prob = prob * m[0][i]
		case 'C':
			prob = prob * m[1][i]
		case 'G':
			prob = prob * m[2][i]
		case 'T':
			prob = prob * m[3][i]
		}
	}
	return prob
}

// find the most probable kmer
func FindBestKMer(s string, k int, m [][]float64) string {
	// look at all kmers in s:
	kmers := make([]string, len(s)-k+1)
	for i := 0; i < len(s)-k+1; i++ {
		kmer := s[i : i+k]
		kmers[i] = kmer
	}
	// fmt.Println(kmers)

	bestProb := 0.0
	bestKmerIndex := 0
	for i := 0; i < len(kmers); i++ {
		prob := Prob(kmers[i], m)
		if prob > bestProb {
			bestProb = prob
			bestKmerIndex = i
		}
	}
	return kmers[bestKmerIndex]
}

func main() {
	// input := "./sampleInput.txt"
	input := "d.txt"
	// input := "./sample2.txt"
	s, k, profileMatrix := ReadEverything(input)
	// fmt.Println(s, k, profileMatrix)
	bestKmer := FindBestKMer(s, k, profileMatrix)
	fmt.Println(bestKmer)
}
