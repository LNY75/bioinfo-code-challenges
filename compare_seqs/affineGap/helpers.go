package main

import "fmt"

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max3(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
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

func PrintStrMatrix(s [][]string) {
	for _, r := range s {
		for _, e := range r {
			if e == "" {
				fmt.Printf("%3s", "--")
			} else {
				fmt.Printf("%3s", e)
			}
		}
		fmt.Println(" ")
	}
}
