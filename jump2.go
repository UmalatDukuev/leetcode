package main

import "fmt"

func canJump(nums []int) int {
	num_len := len(nums)
	goal := num_len - 1
	cnt := 0
	for goal > 0 {
		i := 0
		for i <= goal {
			if nums[i] >= goal-i {
				goal = i
				cnt += 1
			}
			i++
		}
	}
	return cnt
}

func main() {
	nums := []int{2, 3, 0, 1, 4}
	fmt.Println(canJump(nums))
}
