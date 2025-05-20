package main

import (
	"fmt"
	"time"
)

func search(server string, query string) ([]string, error) {
	fmt.Println(1)
	return []string{}, nil
}

func main() {
	servers := []string{
		"server1.example.com",
		"server2.example.com",
		"server3.example.com",
		"server4.example.com",
		"server5.example.com",
	}

	query := "golang concurrency"
	for _, server := range servers {
		go search(server, query)
	}
	time.Sleep(time.Second)
}
