package main

import "fmt"

func PrintAlignment(A []string) {
	for _, a := range A {
		fmt.Println(a)
	}
}

func PrintFloatMatrix(s [][]float64) {
	for _, r := range s {
		for _, e := range r {
			// fmt.Print(e, ", ")
			fmt.Printf(" %6.5f ", e)
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

func PrintFloatMatrixWithLabels(s [][]float64, rowLables []string, colLables []string) {
	// print the column labels:
	fmt.Printf("  ")
	for _, l := range colLables {
		fmt.Print(l, " ")
	}
	fmt.Println("")
	for i, r := range s {
		fmt.Print(rowLables[i], " ")
		for _, e := range r {
			// fmt.Print(e, ", ")
			fmt.Print(e, " ")
		}
		fmt.Println(" ")
	}
}
