package main

func countGood(nums []int, k int) int64 {
	digitMap := make(map[int]int)
	left := 0
	totalPairs := 0
	count := 0

	for right, num := range nums {
		totalPairs += digitMap[num]
		digitMap[num]++

		for totalPairs >= k {
			count += len(nums) - right

			totalPairs -= digitMap[nums[left]] - 1
			digitMap[nums[left]]--
			left++
		}
	}

	return int64(count)
}
