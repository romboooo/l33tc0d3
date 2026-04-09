package main

import (
	"fmt"
	"sync"
)

func printNumber(n int) {
	fmt.Print(n)
}

func main() {
	z := NewZeroEvenOdd(500)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		z.Zero(printNumber)
	}()
	go func() {
		defer wg.Done()
		z.Even(printNumber)
	}()
	go func() {
		defer wg.Done()
		z.Odd(printNumber)
	}()

	wg.Wait()
	fmt.Println()
}
