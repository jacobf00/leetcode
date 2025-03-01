package main

import "fmt"

func main() {
	nums := []int{1,1,0,1}
	result := longestSubarray(nums)
	fmt.Println("The longest subarray is:", result)
}

func longestSubarray(nums []int) int {
	left, right, k := 0, 0, 1
	
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
	return right - left - 1
}