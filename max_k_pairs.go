package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,2,3,4}
	k := 5
	result := maxOperations(nums, k)
	fmt.Println("The maximum number of operations is:", result)
}

func maxOperations(nums []int, k int) int {
	left := 0
	right := len(nums)-1
	var sum int
	operations := 0
	sort.Ints(nums)
	
	for left < right {
		sum = nums[left] + nums[right]
		if sum == k {
			operations++
			left++
			right--
		} else {
			if sum < k {
				left++
			} else {
				right--
			}
		}
	}
	return operations
}