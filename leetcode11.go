package main

import "fmt"

// 1.暴力法 FAIL
func maxArea(height []int) int {
	max, rong := 0, 0
	for left := 0; left < len(height)-1; left++ {
		for right := left+1; right < len(height); right++ {
			length := right - left
			if height[left] >= height[right] {
				rong = length * height[right]
			}else{
				rong = length * height[left]
			}
			if rong >= max {
				max = rong
			}
		}
	}
	fmt.Println(max)
	return max
}





func main() {
	maxArea([]int{1, 2})
}