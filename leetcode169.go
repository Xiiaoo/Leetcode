package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	count, index := 1, 0
	if len(nums) == 1 || len(nums) == 2 {
		return nums[0]
	} else {
		for i := 0; i < len(nums); i++ {
			if _, ok := m[nums[i]]; ok {
				m[nums[i]]++
				if m[nums[i]] > len(nums)/2 {
					index = i
					break
				}
			} else {
				m[nums[i]] = count
			}
		}
		return nums[index]
	}
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 2, 6, 4, 7, 2}))
}
