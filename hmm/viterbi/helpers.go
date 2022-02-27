package main

import "fmt"

func PrintIntMatrix(s [][]int) {
	for _, r := range s {
		for _, e := range r {
			fmt.Print(e, ", ")
			// fmt.Printf(" %10.10f ", e)
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

func PrintFloatMatrix(s [][]float64) {
	for _, r := range s {
		for _, e := range r {
			fmt.Print(e, ", ")
			// fmt.Printf(" %10.10f ", e)
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

func max(l []float64) (float64, int) {
	max := l[0]
	maxi := 0
	for i, e := range l {
		if e > max {
			max = e
			maxi = i
		}
	}
	return max, maxi
}
