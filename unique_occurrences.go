package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{1,2,2,1,1,3}
	result := uniqueOccurrences(arr)
	fmt.Println(result)
}

func uniqueOccurrences(arr []int) bool {
    occurences := make(map[int]int)
	for _, num := range arr {
		count, exists := occurences[num]
		if exists {
			occurences[num] = count + 1
		} else {
			occurences[num] = 1
		}
	}

	var counts []int

	for _, count := range occurences {
		if slices.Contains(counts, count) {
			return false
		}
		counts = append(counts, count)
	}
	return true
}