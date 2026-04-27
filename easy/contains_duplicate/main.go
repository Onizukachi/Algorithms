package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // false
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true
}

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}

		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
