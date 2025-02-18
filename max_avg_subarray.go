package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1,12,-5,-6,50,3}
	k := 4
	// nums := []int{-1}
	// k := 1
	result := findMaxAverage(nums, k)
	fmt.Println("The maximum contiguous subarray avg is:", result)
}

func findMaxAverage(nums []int, k int) float64 {
	start := 0
	end := k-1
	maxAverage := -math.MaxFloat64
	var average float64
	var sum int
	
	for end < len(nums) {
		if start == 0 {
			for i := start; i <= end; i++ {
				sum += nums[i]
			}
			average = float64(sum) / float64(k)
		} else {
			sum = sum - nums[start-1] + nums[end]
			average = float64(sum) / float64(k)
		}
		if average > maxAverage {
			maxAverage = average
		}
		start++
		end++
	}
	return maxAverage
}