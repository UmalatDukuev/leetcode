package main

import (
	"fmt"
	"sync"
)

func mergeChans(chs ...<-chan int) chan int {
	result := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)
		ch := ch
		go func() {
			for data := range ch {
				result <- data
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	ch1, ch2, ch3 := make(chan int, 5), make(chan int, 5), make(chan int, 5)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)

	}()
	go func() {
		ch2 <- 4
		ch2 <- 5
		ch2 <- 6
		close(ch2)

	}()
	go func() {
		ch3 <- 7
		ch3 <- 8
		ch3 <- 9
		close(ch3)

	}()
	for data := range mergeChans(ch1, ch2, ch3) {
		fmt.Println(data)

	}
}
