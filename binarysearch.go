package main

import "fmt"

// array目标数组  target目标数字 left头索引 right尾索引
func BinarySearch(array []int, target int, left, right int) int {
	if left > right {
		return -1
	}
	// 中位数索引
	mid := (left+right) / 2
	midNum := array[mid]
	if midNum == target {
		return mid
	} else if midNum > target {
		return (BinarySearch(array, target, 0, mid-1))
	} else {
		return (BinarySearch(array, target, mid+1, right))
	}
}

func main() {
	array := []int{1, 3, 5, 7, 8, 12, 13}
	target:=12
	fmt.Println(target,BinarySearch(array, target, 0, len(array)-1))
}
