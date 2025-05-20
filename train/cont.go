package main

import (
	"fmt"
	"sync"
)

func merge(chs []chan int) chan int {
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for _, ch_ := range chs {
			wg.Add(1)
			go func(ch_ chan int) {
				defer wg.Done()
				for data := range ch_ {
					ch <- data
				}
			}(ch_)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()
	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)

	}()
	go func() {
		ch3 <- 5
		ch3 <- 6
		close(ch3)

	}()
	chs := []chan int{ch1, ch2, ch3}
	for ch := range merge(chs) {
		fmt.Println(ch)
	}
}
