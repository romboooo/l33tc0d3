package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcdefa"))
}

func lengthOfLongestSubstring(s string) int {

	maxLen := 0

	left := 0
	right := 1
	for range len(s) {

		currStr := s[left:right]

		if !isStringHasDublicate(currStr) {
			maxLen = len(currStr)
			right++
		} else {
			left++
			right++

		}

	}
	return maxLen
}

func isStringHasDublicate(s string) bool {
	seen := make(map[rune]bool)

	for _, char := range s {
		if seen[char] {
			return true
		}
		seen[char] = true
	}
	return false
}
