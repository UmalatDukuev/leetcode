package main

import "fmt"

func maxArea(height []int) int {
	arrLen := len(height)
	// maxHeights := make([]int, 0)
	max := 0
	for i := 0; i < arrLen; i++ {
		left := 0
		right := arrLen - 1

		lowerBound := 0
		if height[left] <= height[right] {
			lowerBound = height[left]
			defer func() {
				left++
			}()
		} else {
			lowerBound = height[right]
			defer func() {
				right--
			}()
		}
		water := (right - left) * lowerBound
		if water > max {
			// fmt.Println(lowerBound)
			// // fmt.Println(right)
			// fmt.Println(left)
			// fmt.Println(right)

			max = water
		}
		// fmt.Println(lowerBound)
	}
	// for i := 0; i < arrLen; i++ {
	// 	for j := i; j < arrLen; j++ {
	// 		lowerBound := 0
	// 		if i != j {
	// 			if height[j] <= height[i] {
	// 				lowerBound = height[j]
	// 			} else {
	// 				lowerBound = height[i]
	// 			}
	// 			water := (j - i) * lowerBound
	// 			if water > max {
	// 				max = water
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(maxHeights)
	// for _, height := range maxHeights {
	// 	if height > max {
	// 		max = height
	// 	}
	// }

	return max
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
