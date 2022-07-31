package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	count, tmp := 0, 0
	if len(nums) == 1 && nums[0] == 1 {
		count++
		return count
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if count>tmp {
				tmp=count
			}
			count = 0
		}
	}
	if count>tmp {
		return count
	}else {
		return tmp
	}
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0,1}))
}
