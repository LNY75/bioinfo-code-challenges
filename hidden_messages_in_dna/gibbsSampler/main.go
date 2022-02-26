package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// read the input txt file, returns the number k (for k-mers to be analyzed later), t (the number of sequences in dna), N (number of iterations for randomly generating the k-mer based on the profile probabilities), dna (the set of dna seqs)
func ReadInput(input string) (int, int, int, []string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	// lines[0] -> k t; k -> length of k-mer, t -> number of sequences in the set dna
	// lines[1:1+t] -> dna
	ktN := strings.Fields(lines[0])
	k, err := strconv.Atoi(ktN[0])
	t, err := strconv.Atoi(ktN[1])
	N, err := strconv.Atoi(ktN[2])
	dna := lines[1 : 1+t]
	return k, t, N, dna
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
// note: we need to ignore, randomly, one seq from the motifs
// returns the profile matrix, the count matrix, and the index of the ignored row
func Profile(motifs []string, k, t int) ([][]float64, [][]int, int) {
	// decide which row from the motifs to ignore:
	luckyRowIndex := rand.Intn(t)
	// luckyRowIndex = 2

	var counts [][]int = make([][]int, 4)          // order: A, C, G, T
	var profile [][]float64 = make([][]float64, 4) // order: A, C, G, T
	// initialize the profile and counts matrix
	for i := 0; i < 4; i++ {
		profile[i] = make([]float64, k)
		counts[i] = make([]int, k)
		for j := 0; j < k; j++ {
			profile[i][j] = 1
			counts[i][j] = 1
		}
	}
	// add counts (ignoring the luckyRow)
	for i := 0; i < t; i++ {
		if i != luckyRowIndex {
			for j := 0; j < k; j++ {
				switch motifs[i][j] {
				case 'A':
					profile[0][j] += 1
					counts[0][j] += 1
				case 'C':
					profile[1][j] += 1
					counts[1][j] += 1
				case 'G':
					profile[2][j] += 1
					counts[2][j] += 1
				case 'T':
					profile[3][j] += 1
					counts[3][j] += 1
				}
			}

		}
	}
	// compute probability
	for j := 0; j < k; j++ {
		for i := 0; i < 4; i++ {
			profile[i][j] /= (float64(t-1) + 4.0)
		}
	}

	// PrintProfile(profile)
	// PrintCounts(counts)
	return profile, counts, luckyRowIndex
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

// given a counts matrix, compute the product of all counts for a given kmer
func CountProduct(kmer string, m [][]int) int {
	prob := 1
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

// select a kmer from the ignored row based on the probability of each kmer
func SelectKmer(luckyRow string, counts [][]int, k int) string {
	thresholds := make([]int, len(luckyRow)-k+1)
	sum := 0
	for i := 0; i < len(thresholds); i++ {
		thresholds[i] = CountProduct(luckyRow[i:i+k], counts)
		sum += thresholds[i]
	}

	for i := 1; i < len(thresholds); i++ {
		thresholds[i] += thresholds[i-1]
	}

	// for i := range thresholds {
	// 	fmt.Print(thresholds[i], " ")
	// }

	// fmt.Println(thresholds)

	randInt := rand.Intn(sum)

	// fmt.Println(randInt)

	kmerIndex := 0
	for i := range thresholds {
		if randInt < thresholds[i] {
			kmerIndex = i
			break
		}
	}

	// fmt.Println(luckyRow[kmerIndex : kmerIndex+k])
	return luckyRow[kmerIndex : kmerIndex+k]
}

// return a new set of k motifs
func Motifs(prevMotifs []string, luckyRowIndex int, dna []string, counts [][]int, k, t int) []string {
	// find the missing kmer
	luckyMotif := SelectKmer(dna[luckyRowIndex], counts, k)
	motifs := make([]string, t)
	for i := 0; i < t; i++ {
		if i == luckyRowIndex {
			motifs[i] = luckyMotif
		} else {
			motifs[i] = prevMotifs[i]
		}
	}
	return motifs
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

// perform Gibbs sampling
func GibbsSampler(dna []string, k, t, N int) []string {
	motifs := RandMotif(dna, k, t)
	bestMotifs := CopyMotifs(motifs)
	for i := 0; i < N; i++ {
		profile, counts, luckyRowIndex := Profile(motifs, k, t)
		motifs = Motifs(bestMotifs, luckyRowIndex, dna, counts, k, t)
		if Score(motifs, profile, k, t) > Score(bestMotifs, profile, k, t) {
			bestMotifs = CopyMotifs(motifs)
		}
	}
	return bestMotifs
}

func main() {
	input := "sample1.txt"
	k, t, N, dna := ReadInput(input)
	// fmt.Println(k, t, N, dna)

	// motifs := RandMotif(dna, k, t)
	// motifs = []string{"TAAC", "GTCT", "CCGG", "ACTA", "AGGT"}
	// _, counts, _ := Profile(motifs, k, t)
	// newMotifs := Motifs(motifs, 2, dna, counts, k, t)
	// fmt.Println(newMotifs)
	// bestMotifs := GibbsSampler(dna, k, t, N)
	// PrintMotifs(bestMotifs)

	// run Gibbs Sampling 20 times
	motifs := GibbsSampler(dna, k, t, N)
	bestMotifs := CopyMotifs(motifs)
	bestProfile, _, _ := Profile(bestMotifs, k, t)
	bestScore := Score(bestMotifs, bestProfile, k, t)
	for i := 0; i < 60; i++ {
		motifs = GibbsSampler(dna, k, t, N)
		profile, _, _ := Profile(motifs, k, t)
		score := Score(motifs, profile, k, t)

		if score > bestScore {
			bestScore = score
			bestMotifs = CopyMotifs(motifs)
		}
	}
	PrintMotifs(bestMotifs)
}
