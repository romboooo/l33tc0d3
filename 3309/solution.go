package main

import (
	"fmt"
	"math/bits"
)

func maxGoodNumber(num []int) int {

	max := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 3; k++ {
				if i == k || k == j {
					continue
				}

				val := ((num[i]<<bits.Len(uint(num[j])))|num[j])<<bits.Len(uint(num[k])) | num[k]

				if val > max {
					max = val
				}
			}
		}
	}

	fmt.Println(max)
	return max
}
