package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= k; j++ {
			if i+j <= len(nums)-1 && nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
