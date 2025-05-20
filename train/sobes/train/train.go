package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go printMessage("Hello from goroutine")
	printMessage("Hello from main")
}
