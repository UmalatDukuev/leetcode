package main

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

func Search(server string, query string) ([]string, error) {
	if server == "server1" {
		time.Sleep(2 * time.Second)
		return nil, errors.New("server1 failed")
	} else if server == "server2" {
		time.Sleep(1 * time.Second)
		return []string{"doc1", "doc2"}, nil
	}
	return nil, errors.New("unknown server")
}

func search(servers []string, query string) ([]string, error) {

	response := make(chan []string)
	err := make(chan error)
	var errCounter atomic.Int32

	for _, server := range servers {
		srv := server
		go func() {
			docs, err := Search(srv, query)
			if err != nil {
				errCounter.Add(1)
				if errCounter == len(servers){
					
				}
			}
		}()
	}

}

func main() {
	servers := []string{
		"server1",
		"server21",
		// "server3.example.com",
		// "server4.example.com",
		// "server5.example.com",
	}

	query := "golang concurrency"
	fmt.Println(search(servers, query))

	// time.Sleep(time.Second)
}
