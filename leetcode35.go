package main

import "fmt"

func searchInsert(nums []int, target int) int {
	mid := nums[(len(nums)-1)/2]
	index := 0
	if len(nums) == 1 {
		if nums[0] >= target {
			index = 0
		} else {
			index = 1
		}
	} else {
		if mid == target {
			index = (len(nums) - 1) / 2
		} else if target < mid {
			fmt.Println(1)
			for i := 0; i <= (len(nums)-1)/2; i++ {
				if nums[i] == target {
					index = i
				} else if target > nums[i] && target < nums[i+1] {
					index = i + 1
				} else if target > nums[i+1] {
					index = i + 2
				}
			}
		} else if target > mid {
			for i := (len(nums) - 1) / 2; i < len(nums)-1; i++ {
				if nums[i+1] == target {
					index = i + 1
				} else if target > nums[i] && target < nums[i+1] {
					index = i + 1
				} else if target > nums[i+1] {
					index = i + 2
				}
			}
		}
	}
	return index
}

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 3))
}
