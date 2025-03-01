package main

import (
	"fmt"
	"os"
)

func main() {
	nums := []int{1,1,1,0,0,0,1,1,1,1,0}
	k := 2
	result := longestOnes(nums, k)
}

func longestOnes(nums []int, k int) int {
	start := 0
	end := k-1
	numConsecutiveOnes := 0
	maxNumConsecutiveOnes := 0
	numFlipped := 0
	flippedArr := make([]bool, len(nums))
	
	for end < len(nums) {
		for numFlipped < 5 {
			
		}
		if start == 0 {
			for i := start; i <= end; i++ {
				if nums[i] == 0 {
					numFlipped++
					flippedArr[i] = true
					numConsecutiveOnes++
				} else if nums[i] == 1 {
					numConsecutiveOnes++
				} else {
					fmt.Println("Error! Value is neither 1 or 0")
					os.Exit(1)
				}
			}
		} else {
			if nums[end] == 0 && numFlipped < k {
				numFlipped++
				flippedArr[end] = true
				numConsecutiveOnes++
			} else if nums[end] == 1 {
				numConsecutiveOnes++
			} else if nums[end] == 0 && numFlipped >= k {
				
			}
			
		}
		if numConsecutiveOnes > maxNumConsecutiveOnes {
			maxNumConsecutiveOnes = numConsecutiveOnes
		}

		start++
		end++
	}
}