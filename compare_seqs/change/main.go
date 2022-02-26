package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// read an int value for money and an array of coins
func ReadInput(input string) (int, []int) {
	// read input
	content, err := os.ReadFile(input)
	if err != nil {
		panic("cannot open file")
	}
	contentStr := string(content)
	lines := strings.Split(contentStr, "\n")
	moneyStr := lines[0]
	coinsStr := strings.Split(lines[1], " ")
	var coins []int = make([]int, len(coinsStr))

	money, err := strconv.Atoi(moneyStr)
	for i, c := range coinsStr {
		coin, err := strconv.Atoi(c)
		if err != nil {
			panic("cannot parse string to int")
		}
		coins[i] = coin
	}
	return money, coins
}

// returns the minimum number of coins with denominations Coins that changes money
// Coins: {coins1, coins2, ... }
func DPChange(money int, coins []int) int {
	MinNumCoins := make([]int, 1)
	MinNumCoins[0] = 0
	maxInt := math.MaxInt32
	for m := 1; m <= money; m++ {
		MinNumCoins = append(MinNumCoins, maxInt)
		for _, c := range coins {
			if m >= c {
				if MinNumCoins[m-c]+1 < MinNumCoins[m] {
					MinNumCoins[m] = MinNumCoins[m-c] + 1
				}
			}
		}
	}
	fmt.Println(MinNumCoins)
	return MinNumCoins[money]
}

func main() {
	input := "input1.txt"
	money, coins := ReadInput(input)
	fmt.Println(money, coins)

	numcoins := DPChange(money, coins)
	fmt.Println(numcoins)
}
