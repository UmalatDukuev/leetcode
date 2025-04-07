package main

import "fmt"

func canJump(nums []int) bool {
	num_len := len(nums)
	goal := num_len - 1
	i := num_len - 2
	for i >= 0 {
		if nums[i] >= goal-i {
			goal = i
		}
		i--
	}
	if goal == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	nums := []int{2, 3, 0, 1, 4}
	fmt.Println(canJump(nums))
}
