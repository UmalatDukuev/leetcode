package main

import "fmt"

func quadros(nums []int) []int {

	for i, v := range nums {
		nums[i] = v * v
	}
	return nums
}

func main() {
	nums := []int{2, -3, 0, 1, 4}
	fmt.Println(quadros(nums))
}
