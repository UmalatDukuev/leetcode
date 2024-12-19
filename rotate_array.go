package main

import "fmt"

func rotate(nums []int, k int) {
	c := len(nums)
	k = k % c
	
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	num := 3
	rotate(arr, num)
	fmt.Println(arr)
}
