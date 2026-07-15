package main

import "fmt"

func main() {
	arr := []int{-1, -100, 3, 99}
	rotate(arr, 2)
}

func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := range nums {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, tmp)
	fmt.Println(tmp)
}
