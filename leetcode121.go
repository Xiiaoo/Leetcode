package main

import "fmt"

func maxProfit(prices []int) int {
	count, start, profit, end := 0, 0, 0, 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] >= prices[i+1] {
			count++
		} else {
			for j := i + 1; j < len(prices); j++ {
				if start <= prices[i] && end > prices[j] {
					profit = end - start
				}
			}
		}

	}
	if count == len(prices)-1 {
		return 0
	} else {
		return profit
	}
}

func main() {
	fmt.Println(maxProfit([]int{7, 5, 4, 5}))
}