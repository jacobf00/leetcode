package main

import "fmt"

func main() {
	nums := []int{1,7,3,6,5,6}
	result := pivotIndex(nums)
	fmt.Println("The pivot index is:", result)
}

func pivotIndex(nums []int) int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	sum := 0
	if len(nums) == 0 {
		return -1
	}
	// Calculate leftSum
	for i := 1; i < len(nums); i++ {
		// previous sum plus num at index before
		sum += nums[i-1]
		leftSum[i] = sum
	}

	// reset sum
	sum = 0

	rightSum[len(nums)-1] = 0
	for j := len(nums)-2; j >= 0; j-- {
		// previous sum plus num at index before
		sum += nums[j+1]
		rightSum[j] = sum
	}
	
	for i, _ := range nums {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}
	return -1
}