package main

import "fmt"

func PrintProfile(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], ", ")
		}
		fmt.Println(" ")
	}
}

func PrintCounts(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j], ", ")
		}
		fmt.Println(" ")
	}
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
