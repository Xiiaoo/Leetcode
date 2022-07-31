package main

import "fmt"

// 想复杂了
// func divingBoard(shorter int, longer int, k int) []int {
// 	nums, result := []int{}, []int{}
// 	if k == 0 {
// 		return nums
// 	} else {
// 		nums = append(nums, k*shorter)
// 		for i := 1; i <= k; i++ {
// 			nums = append(nums, i*longer+(k-i)*shorter)
// 		}

// 	}
// 	temp := map[int]struct{}{}
// 	for _, item := range nums {
// 		if _, ok := temp[item]; !ok {
// 			temp[item] = struct{}{}
// 			result = append(result, item)
// 		}
// 	}

// 	return result
// }

// 官解  板子数K为0时返回空，短板和长版一样时，返回k*shorter，不一样长时，省略
func divingBoard(shorter int, longer int, k int) []int {
    if k == 0 {
        return []int{}
    }
    if shorter == longer {
        return []int{shorter * k}
    }
    lengths := make([]int, k + 1)
    for i := 0; i <= k; i++ {
        lengths[i] = shorter * (k - i) + longer * i
    }
    return lengths
}




func main() {
	fmt.Println(divingBoard(1, 2, 4))
}
