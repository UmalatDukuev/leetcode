package main

import "fmt"

func isPalindrome(str string) bool {
	for _, value := range str {
		fmt.Println(value)
	}
	return true
}

func main() {
	str := "qwerty1"
	fmt.Println(isPalindrome(str))
}
