package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// read the input txt file, returns the number k (for k-mers to be analyzed later), t (the number of dna seqs in the set dna) and the set of DNA sequences
func ReadInput(input string) (int, int, []string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// lines[0] -> k t; k -> length of k-mer, t -> number of sequences in the set dna
	// lines[1:1+t] -> dna
	kt := strings.Fields(lines[0])
	k, err := strconv.Atoi(kt[0])
	t, err := strconv.Atoi(kt[1])

	dna := lines[1 : 1+t]
	return k, t, dna
}

// randomly generates a motif set from dna
func RandMotif(dna []string, k, t int) []string {
	var motifs []string = make([]string, t)
	for i := 0; i < t; i++ {
		// number of k-mers in a sequence = n-k+1 (n=length of sequence)
		numberKmers := len(dna[i]) - k + 1
		// the starting position of a random k-mer; 0<=index<numberKmers
		index := rand.Intn(numberKmers)
		kmer := dna[i][index : index+k]
		motifs[i] = kmer
	}
	return motifs
}

// given a set of motifs, generate the corresponding profile matrix (pseudocounts are added)
func Profile(motifs []string, k, t int) [][]float64 {
	var profile [][]float64 = make([][]float64, 4) // order: A, C, G, T
	// initialize the profile matrix
	for i := 0; i < 4; i++ {
		profile[i] = make([]float64, k)
		for j := 0; j < k; j++ {
			profile[i][j] = 1
		}
	}
	// add counts
	for i := 0; i < t; i++ {
		for j := 0; j < k; j++ {
			switch motifs[i][j] {
			case 'A':
				profile[0][j] += 1
			case 'C':
				profile[1][j] += 1
			case 'G':
				profile[2][j] += 1
			case 'T':
				profile[3][j] += 1
			}
		}
	}

	// compute probability
	for j := 0; j < k; j++ {
		for i := 0; i < 4; i++ {
			profile[i][j] /= (float64(t) + 4.0)
		}
	}

	return profile
}

func PrintMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], ", ")
		}
		fmt.Println(" ")
	}
}

// given a profile, find the best set of motifs
func Motifs(profile [][]float64, motifs []string, dna []string, k, t int) []string {
	bestMotifs := motifs
	for i := 0; i < t; i++ {
		// for every possible k-mer in each seq from dna:
		for j := 0; j < len(dna[i])-k+1; j++ {
			currentMotif := dna[i][j : j+k]
			previousProb := Prob(bestMotifs[i], profile)
			currentProb := Prob(currentMotif, profile)
			if currentProb > previousProb {
				bestMotifs[i] = currentMotif
			}
		}
	}

	return bestMotifs
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

// find the consensus string, given a profile matrix
func Consensus(profile [][]float64, k int) string {
	var consensus string
	for j := 0; j < k; j++ {
		maxIndex := 0 // index of the most common nucleotide at position j
		for i := 0; i < 4; i++ {
			if profile[i][j] > profile[maxIndex][j] {
				maxIndex = i
			}
		}
		switch maxIndex {
		case 0:
			consensus += "A"
		case 1:
			consensus += "C"
		case 2:
			consensus += "G"
		case 3:
			consensus += "T"
		}
	}
	return consensus
}

// calculate the score of a set of motifs
func Score(motifs []string, profile [][]float64, k, t int) int {
	consensus := Consensus(profile, k)
	score := 0
	for i := 0; i < t; i++ {
		for j := 0; j < k; j++ {
			if motifs[i][j] == consensus[j] {
				score++
			}
		}
	}
	return score
}

// find the best motif
func RandomizedMotifSearch(dna []string, k, t int) []string {
	// randomly select motifs:
	motifs := RandMotif(dna, k, t)

	// motifs = []string{"CCTG", "ACAG", "TTGG", "CAGT"}

	bestMotifs := CopyMotifs(motifs)

	for i := 0; i < 10; i++ {
		profile := Profile(bestMotifs, k, t)

		motifs = Motifs(profile, motifs, dna, k, t)

		score := Score(motifs, profile, k, t)
		if score > Score(bestMotifs, profile, k, t) {
			// fmt.Println("lol")
			bestMotifs = motifs
		}
	}

	return bestMotifs
}

func CopyMotifs(m []string) []string {
	r := make([]string, len(m))
	for i := 0; i < len(m); i++ {
		r[i] = m[i]
	}
	return r
}

func PrintMotifs(m []string) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

func main() {
	k, t, dna := ReadInput("test.txt")

	// run RandomizedMotifSearch 1000 times
	motifs := RandomizedMotifSearch(dna, k, t)
	bestMotifs := CopyMotifs(motifs)
	bestProfile := Profile(bestMotifs, k, t)
	bestScore := Score(bestMotifs, bestProfile, k, t)
	for i := 0; i < 1000; i++ {
		motifs = RandomizedMotifSearch(dna, k, t)
		profile := Profile(motifs, k, t)
		score := Score(motifs, profile, k, t)

		if score > bestScore {
			bestScore = score
			bestMotifs = motifs
		}
	}
	PrintMotifs(bestMotifs)
}
