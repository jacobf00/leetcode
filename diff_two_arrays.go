package main

import (
	"fmt"
	"slices"
)

func main() {
	nums1 := []int{1,2,3}
	nums2 := []int{2,4,6}
	result := findDifference(nums1, nums2)
	fmt.Println("The resulting sets are:", result)
}

func findDifference(nums1 []int, nums2 []int) [][]int {
    var answer1 []int
    var answer2 []int
	
	for _, num := range nums1 {
		if !slices.Contains(answer1, num) && !slices.Contains(nums2, num) {
			answer1 = append(answer1, num)
		}
	}

	for _, num := range nums2 {
		if !slices.Contains(answer2, num) && !slices.Contains(nums1, num) {
			answer2 = append(answer2, num)
		}
	}
	return [][]int{answer1, answer2}
}