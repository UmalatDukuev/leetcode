package main

import "fmt"

type user struct {
	balance int64
}

func main() {
	users := []user{
		{balance: 1000},
		{balance: 2000},
	}
	for _, u := range users {
		fmt.Println(u)
		u.balance += 1000
	}
	fmt.Println(users)
}
