package main

import "fmt"

func main() {
	fmt.Println(gcdOfOddEvenSums(5))
}

func gcdOfOddEvenSums(n int) int {
	sumOdd, sumEven := getSumOdd(n), getSumEven(n)
	maxDiv := max(sumOdd, sumEven)

	for i := range maxDiv {
		if i == 0 {
			continue
		}
		if sumOdd%i == 0 && sumEven%i == 0 {
			maxDiv = i
		}
	}

	return maxDiv
}

func getSumOdd(n int) int {
	sum := 0

	for i := range n * 2 {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	return sum
}
func getSumEven(n int) int {
	sum := 0

	for i := range n * 2 {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	return sum
}
