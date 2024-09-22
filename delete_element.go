package main

import "fmt"

func removeElement(nums []int, val int) int {
	nums_size := len(nums)
	k, i, j := 0, 0, nums_size-1

	for i < j {
		if nums[i] == val {

			for i < j && nums[j] == val {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
			k++
		}
	}
	fmt.Println(k)
	for i := 0; i < k; i++ {
		fmt.Printf("%d ", nums[i])
	}

	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	removeElement(nums, val)
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Printf("%d ", nums[i])
	// }
}
