package main

import "fmt"

func rotate(nums []int, k int) {
	for j := 0; j < k; j++ {
		temp := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
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
