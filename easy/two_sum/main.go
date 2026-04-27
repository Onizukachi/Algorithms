package main

import "fmt"

func main() {
	arr := []int{3, 3}
	fmt.Println(twoSum(arr, 6))
}

func twoSum(nums []int, target int) []int {
	values := make(map[int]int)

	for i := range nums {
		if i == 0 {
			values[nums[i]] = i
			continue
		}

		diff := target - nums[i]
		if idx, ok := values[diff]; ok {
			return []int{idx, i}
		}

		values[nums[i]] = i
	}

	return nums
}
