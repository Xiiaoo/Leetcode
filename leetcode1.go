package main

// 1.暴力解 时间复杂度：O(N^2)  空间复杂度：O(1)
// func twoSum(nums []int, target int) []int {
// 	re := []int{}
// 	for i := 0; i <= len(nums)-2; i++ {
// 		for right:= i + 1;right<len(nums);right++{
// 			if nums[i]+nums[right] == target{
// 				re := []int{i, right}
// 				fmt.Println(re)
// 				return re
// 			}
// 		}
// 	}
// 	return re
// }

// 1.暴力解leetcode 时间复杂度：O(N^2)  空间复杂度：O(1)
// func twoSum(nums []int, target int) []int {
//     for i, x := range nums {
//         for j := i + 1; j < len(nums); j++ {
//             if x+nums[j] == target {
//                 return []int{i, j}
//             }
//         }
//     }
//     return nil
// }

// 2.哈希表 时间复杂度：O(N) 空间复杂度O(N)
func twoSum(nums []int, target int) []int {
    hashTable := map[int]int{}
    for i, x := range nums {
        if p, ok := hashTable[target-x]; ok {
            return []int{p, i}
        }
        hashTable[x] = i
    }
    return nil
}

func main() {
	twoSum([]int{3,2,4}, 6)
}