package main

import "fmt"

func main() {
	gain := []int{-5,1,5,0,-7}
	result := largestAltitude(gain)
	fmt.Println("The highest altitude is:", result)
}

func largestAltitude(gain []int) int {
	currentAltitude := 0
	maxAltitude := 0
	for i := 0; i < len(gain); i++ {
		currentAltitude += gain[i]
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}
	return maxAltitude
}