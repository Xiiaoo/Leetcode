package main

import "fmt"

func countOdds(low int, high int) int {
	count := 0
	if low == high && low%2 == 1 {
		return 1
	} else if low == high && low%2 == 0 {
		return 0
	}
	if low%2 == 1 && high%2 == 1 {
		count = (high-low)/2 + 1
	} else if low%2 == 1 && high%2 == 0 {
		count = (high-low-1)/2 + 1
	} else if low%2 == 0 && high%2 == 1 {
		count = (high-low-1)/2 + 1
	} else if low%2 == 0 && high%2 == 0 {
		count = (high-low)/2
	}
	return count
}

func main() {
	fmt.Println(countOdds(4, 9))
}
