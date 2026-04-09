package main

import (
	"sync"
)

type ZeroEvenOdd struct {
	size    int
	currNum int
	mu      sync.Mutex
	zeroCh  chan struct{}
	evenCh  chan struct{}
	oddCh   chan struct{}
}

func NewZeroEvenOdd(size int) *ZeroEvenOdd {

	zeo := &ZeroEvenOdd{
		size:    size,
		currNum: 1,
		zeroCh:  make(chan struct{}, 1),
		evenCh:  make(chan struct{}),
		oddCh:   make(chan struct{}),
	}
	zeo.zeroCh <- struct{}{}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {

	for i := 0; i < z.size; i++ {
		<-z.zeroCh
		printNumber(0)
		if i%2 == 0 {
			z.oddCh <- struct{}{}
		} else {
			z.evenCh <- struct{}{}

		}

	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {

	for num := 2; num <= z.size; num += 2 {
		<-z.evenCh
		printNumber(num)
		if num < z.size {
			z.zeroCh <- struct{}{}
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for num := 1; num <= z.size; num += 2 {
		<-z.oddCh
		printNumber(num)
		if (num) < z.size {
			z.zeroCh <- struct{}{}
		}
	}
}
