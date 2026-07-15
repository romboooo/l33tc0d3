package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSubstringsInPartition("abababaccddb"))
}

func minimumSubstringsInPartition(s string) int {

	indexToSolutionsMap := map[int]int{}
	var solve func(i int) int
	solve = func(i int) int {
		answer := 10000000
		if i == len(s) {
			return 0
		}
		if ans, ok := indexToSolutionsMap[i]; ok {
			return ans
		}
		freq := [26]int{}
		for j := i; j < len(s); j++ {

			freq[s[j]-'a']++
			if isBalanced(freq) {
				candidate := 1 + solve(j+1)
				answer = min(answer, candidate)
				indexToSolutionsMap[i] = answer
			}
		}
		return answer
	}

	return solve(0)
}

func isBalanced(freq [26]int) bool {
	reference := 0

	for _, count := range freq {
		if count == 0 {
			continue
		}

		if reference == 0 {
			reference = count
			continue
		}

		if count != reference {
			return false
		}
	}

	return true
}
