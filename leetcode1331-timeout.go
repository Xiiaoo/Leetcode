package main

import "fmt"

// 暴力解法，超时
// func arrayRankTransform(arr []int) []int {
// 	oldarr,newarr,rank := []int{},[]int{},1
// 	newarr = append(newarr, rank)
// 	n := len(arr)
// 	for i := 0; i < n; i++ {
// 		oldarr = append(oldarr, arr[i])
// 	}
// 	for i := n - 1; i > 0; i-- {
// 		for j := 0; j < i; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// 	for i := 1; i < n; i++ {
// 		if arr[i] == arr[i-1] {
// 			newarr = append(newarr, rank)
// 		} else {
// 			rank++
// 			newarr = append(newarr, rank)
// 		}
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if arr[j] == oldarr[i] {
// 				oldarr[i] = newarr[j]
// 				break
// 			}
// 		}
// 	}
// 	fmt.Println(newarr)
// 	return oldarr
// }

func arrayRankTransform(arr []int) []int {
	
}

func main() {
	fmt.Println(arrayRankTransform([]int{2,-11,24,15,26,-31,-48,-49,22,37,-36}))
}