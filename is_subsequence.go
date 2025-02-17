package main

import "fmt"

func main() {
	t := "ahbgdc"
	s := "abc"
	result := isSubsequence(s, t)
	if result {
		fmt.Println(s, "is a subsequence of", t)
	} else {
		fmt.Println("Not a subsequence")
	}
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sp := 0
	for tp := range t {
		if s[sp] == t[tp] {
			sp++
		}
		if sp == len(s) {
			return true
		}
	}
	return false
}