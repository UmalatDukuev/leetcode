package main

import "fmt"

func rotate(nums []int, k int) {
	c := len(nums)
	k = k % c
	for j := 0; j < k; j++ {
		temp := nums[c-1]
		for i := c - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	num := 3
	rotate(arr, num)
	fmt.Println(arr)
}
