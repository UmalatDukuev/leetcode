package main

func mergeChans(chs ...<-chan int) chan int{
	
}

func main() {
	ch1, ch2, ch3 := make(chan int, 5), make(chan int, 5), make(chan int, 5)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()
	go func() {
		ch2 <- 4
		ch2 <- 5
		ch2 <- 6
	}()
	go func() {
		ch3 <- 7
		ch3 <- 8
		ch3 <- 9
	}()
}
