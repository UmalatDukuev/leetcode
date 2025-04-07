package main

import "fmt"

func maxArea(height []int) int {
	arrLen := len(height)
	max := 0
	left := 0
	right := arrLen - 1

	for left < right {
		lowerBound := 0
		if height[left] <= height[right] {
			lowerBound = height[left]

		} else {
			lowerBound = height[right]
		}
		water := (right - left) * lowerBound
		if water > max {
			max = water
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func main() {
	height := []int{1, 10000, 6, 2, 5, 4, 10000, 3, 7}
	fmt.Println(maxArea(height))
}
