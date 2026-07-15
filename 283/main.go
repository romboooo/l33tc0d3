package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
}

func moveZeroes(nums []int) {
	zero := 0

	for scan := range nums {
		if nums[scan] != 0 {
			nums[zero], nums[scan] = nums[scan], nums[zero]
			zero++
		}
	}

	fmt.Println(nums)
}
