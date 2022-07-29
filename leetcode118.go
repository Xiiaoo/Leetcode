package main

import "fmt"

// 本人解，较冗余
// func generate(numRows int) [][]int {
// 	subarr, arr := []int{}, [][]int{}
// 	if numRows <= 2 {
// 		for i := 0; i < numRows; i++ {
// 			subarr = append(subarr, 1)
// 			arr = append(arr, subarr)
// 		}
// 	}else {
// 		arr = [][]int{{1}, {1, 1}}
// 		for i := 2; i < numRows; i++ {
// 			subarr=[]int{}
// 			subarr = append(subarr, 1)
// 			for j := 1; j < i; j++ {
// 				subarr=append(subarr,arr[i-1][j-1]+arr[i-1][j])
// 			}
// 			subarr = append(subarr, 1)
// 			arr = append(arr, subarr)
// 		}
// 	}
// 	return arr
// }

// 官方解
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(generate((16)))
}
