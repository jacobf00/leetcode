package main

import "fmt"

func main() {
	// nums := []int{1,1,1,0,0,0,1,1,1,1,0}
	nums := []int{1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0}
	k := 2
	result := longestOnes(nums, k)
	fmt.Println("The max num of consecutive ones is:", result)
}

func longestOnes(nums []int, k int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		right++
		if k < 0 {
			if nums[left] == 0{
				k++
			}
			left++
		}
		fmt.Println("left:", left)
		fmt.Println("right:", right)
		fmt.Println("k:", k)
	}
	return right - left 
}