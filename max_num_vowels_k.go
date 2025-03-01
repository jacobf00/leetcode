package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abciiidef"
	k := 3
	result := maxVowels(s, k)
	fmt.Println("The max number of vowels in k length subarray is:", result)
}


func maxVowels(s string, k int) int {
	vowels := "aeiou"
	vowelArr := make([]bool, len(s))
	numMaxVowels := 0
	numVowels := 0
	var isVowel bool
	start := 0
	end := k-1

	for end < len(s) {
		if start == 0 {
			for i := start; i <= end; i++ {
				isVowel = strings.Contains(vowels, string(s[i]))
				vowelArr[i] = isVowel
				if isVowel {
					numVowels++
				}
			}
		} else {
			vowelArr[end] = strings.Contains(vowels, string(s[end]))
			if vowelArr[start-1] {
				numVowels--
			}
			if vowelArr[end] {
				numVowels++
			}
		}
		if numVowels > numMaxVowels {
			numMaxVowels = numVowels
		}
		start++
		end++
	}
	return numMaxVowels
}