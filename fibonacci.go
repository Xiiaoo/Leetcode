package main

import "fmt"

func F(n, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	// n=n-1  a1=a2  a2=a1+a2 尾递归
	return (F(n-1, a2, a1+a2))
}

func main() {
	fmt.Println(F(4, 1, 1))
}