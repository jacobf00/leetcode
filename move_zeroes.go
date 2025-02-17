package main

import "fmt"


func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{2, 1}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	right := len(nums)-1

	// find the last non-zero element
	for nums[right] == 0 && right > 0 {
		right--
	}
	
	left := right

	for left > 0 {

		// move the left pointer back to the last non-zero element
		left = right
		
		// move left to the next zero moving left from the right pointer
		for left > 0 && nums[left] != 0 {
			left--
		}
		
		if nums[left] != 0 {
			break
		}

		// shift elements between left and right to the left
		for i := left; i < right; i++ {
			nums[i] = nums[i+1] 
		}
		nums[right] = 0

		// move right one to the left now that another zero is on the end
		if left < right {
			right--
		}

	}
	
	fmt.Printf("The new string is: %v\n", nums)
	
}