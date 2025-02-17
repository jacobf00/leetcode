package main

import "fmt"

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	result := max_area(height)
	fmt.Println("The max area is:", result)
}

func max_area(height []int) int {
	left := 0
	right := len(height)-1
	max_area := 0
	var area int

	for left < right {
		
		if height[left] < height[right] {
			area = height[left] * (right - left)
		} else {
			area = height[right] * (right - left)
		}

		if area > max_area {
			max_area = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max_area
}