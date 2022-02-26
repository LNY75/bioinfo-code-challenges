package main

import "fmt"

func max(scores []int) int {
	max := scores[0]
	for _, s := range scores {
		if s > max {
			max = s
		}
	}
	return max
}

func PrintIntMatrix(s [][]int) {
	for _, r := range s {
		for _, e := range r {
			fmt.Print(e, " ")
		}
		fmt.Println(" ")
	}
}
