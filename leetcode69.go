package main

import "fmt"

func mySqrt(x int) int {
	i := 1
	for i*i <= x {
		i++
	}
	return i - 1
}

func main() {
	fmt.Println(mySqrt(26))
}
