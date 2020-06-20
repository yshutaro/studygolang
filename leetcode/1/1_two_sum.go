package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	r := twoSum(nums, 9)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	var result []int
	result = make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return result
}
