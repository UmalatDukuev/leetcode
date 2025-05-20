package main

import "fmt"

func factorial(n int, ch chan int) {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	ch <- sum

}

func main() {
	ch := make(chan int)
	go factorial(5, ch)
	fmt.Println(<-ch)
}
