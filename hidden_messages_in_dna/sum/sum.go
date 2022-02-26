package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	// read input
	input := "./dataset_647148_2.txt"
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	nums := strings.Fields(contentStr)

	// sum numbers
	a, err := strconv.Atoi(nums[0])
	b, err := strconv.Atoi(nums[1])
	sum := a + b

	// write to output file
	f, err := os.Create("output.txt")
	sumStr := strconv.Itoa(sum)
	f.WriteString(sumStr)
}
