package main

import "fmt"

// 分类讨论
// 1.如果个位不是9，那么直接把个位加1返回
// 2.如果个位是9，但数组长度为1，那么直接返回[1,0]
// 3.个位是9，长度大于1，那么倒序遍历一次数组
// 		3.1 如果索引i不是头索引而且对应的值不为9，那么digits[i]自加1，返回digits
// 		3.2 如果索引i是头索引而且对应的值为9，那么digits[i]=1，其他元素改为0，最后使用append()函数在数组后加一个0，返回digits
// func plusOne(digits []int) []int {
// 	if digits[len(digits)-1] != 9 {
// 		digits[len(digits)-1]++
// 		return digits
// 	}else if digits[len(digits)-1] == 9 && len(digits)==1 {
// 		digits=[]int{1,0}
// 		return digits
// 	}
// 	for i := len(digits) - 2; i >= 0; i-- {
// 		if digits[i] != 9 {
// 			digits[i]++
// 			digits[len(digits)-1] = 0
// 			return digits
// 		} else {
// 			if i == 0 {
// 				digits[i] = 1
// 				digits[len(digits)-1] = 0
// 				digits = append(digits, 0)
// 			} else {
// 				digits[i] = 0
// 			}
// 		}
// 	}
// 	return digits
// }

// leetcode官方解
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9}))
}
