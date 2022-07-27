package main

// import (
// 	"fmt"
// 	"strconv"
// )

// 1.进制转换 O(N)
// func isPowerOfTwo(n int) bool {
// 	n2:=strconv.FormatInt(int64(n), 2)
// 	n3:=string(n2)
// 	var x=1
// 	for i:=1;i<len(n3);i++{
// 		x*=2
// 	}
// 	if n==x{
// 		fmt.Println("true")
// 		return true
// 	}else{
// 		fmt.Println("false")
// 		return false
// 	}
// }

// 1.二进制转换leetcode 时间复杂度：O(1) 空间复杂度：O(1)
// func isPowerOfTwo(n int) bool {
//     return n > 0 && n&(n-1) == 0
// }

// 2.判断是否为最大 22 的幂的约数  时间复杂度：O(1) 空间复杂度：O(1)
func isPowerOfTwo(n int) bool {
	const big = 1 << 30
	return n > 0 && big%n == 0
}

func main() {
	isPowerOfTwo(18)
}
